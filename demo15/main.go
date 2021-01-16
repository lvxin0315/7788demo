package main

import (
	"flag"
	"fmt"
	"github.com/pingcap/errors"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

var host = flag.String("host", "127.0.0.1", "MySQL host")
var port = flag.Int("port", 3306, "MySQL port")
var user = flag.String("user", "root", "MySQL user, must have replication privilege")
var password = flag.String("password", "root", "MySQL password")
var tables = flag.String("tables", "", "Schema.table_name 使用逗号分隔")
var flavor = flag.String("flavor", "mysql", "Flavor: mysql or mariadb")
var rawMode = flag.Bool("raw", false, "Use raw mode")
var serverID = flag.Int("serverID", 111, "Slave serverID")

// table集合
var tableList []string

// 字段集合
var tableColumnMap sync.Map

// pos
var pos mysql.Position

func main() {
	// 解析参数
	flag.Parse()

	// 处理table集合
	tableList = strings.Split(*tables, ",")

	// 1.初始化表字段 TODO
	for _, t := range tableList {
		err := getTableFields(t)
		if err != nil {
			panic(err)
		}
	}

	// pos更新情况
	masterPos, err := getMasterPos()
	if err != nil {
		panic(err)
	}
	pos = masterPos

	go listenPos()
	// 2.binlog 监听
	for {
		fmt.Println("binlogSyncer is start")
		binlogSyncer()
		fmt.Println("binlogSyncer is over")
	}

}

/**
 * @Author lvxin0315@163.com
 * @Description binlog 起点
 * @Date 3:23 下午 2021/1/15
 * @Param
 * @return
 **/
func getMasterPos() (mysql.Position, error) {
	c, err := client.Connect(fmt.Sprintf("%s:%d", *host, *port), *user, *password, "")
	if err != nil {
		panic(err)
	}
	rr, err := c.Execute("SHOW MASTER STATUS")
	if err != nil {
		return mysql.Position{}, errors.Trace(err)
	}

	name, _ := rr.GetString(0, 0)
	pos, _ := rr.GetInt(0, 1)

	return mysql.Position{Name: name, Pos: uint32(pos)}, nil
}

/**
 * @Author lvxin0315@163.com
 * @Description 行事件的内容
 * @Date 3:23 下午 2021/1/15
 * @Param
 * @return
 **/
func rowEventInfo(ev *replication.BinlogEvent) {
	rowsEv := ev.Event.(*replication.RowsEvent)
	schema := rowsEv.Table.Schema
	table := rowsEv.Table.Table
	tableName := fmt.Sprintf("%s.%s", schema, table)

	if !inTableList(tableName) {
		return
	}
	columnNameList := getTableColumnList(tableName)
	//字段数量变化
	if len(columnNameList) != int(rowsEv.ColumnCount) {
		fmt.Println("refreshAndGetTableColumnList")
		columnNameList = refreshAndGetTableColumnList(tableName)
	}

	for _, dataList := range rowsEv.Rows {
		for index, data := range dataList {
			//TODO
			//fmt.Println( fmt.Sprintf("%s : %v", columnNameList.([]string)[index], data) )
			ioutil.WriteFile("./data.log", []byte(fmt.Sprintf("%s : %v", columnNameList[index], data)), 0777)
		}
	}

}

/**
 * @Author lvxin0315@163.com
 * @Description 获取表字段名称
 * @Date 3:22 下午 2021/1/15
 * @Param
 * @return
 **/
func getTableFields(schemaTableName string) error {
	query := fmt.Sprintf("SHOW COLUMNS FROM %s", schemaTableName)
	c, err := client.Connect(fmt.Sprintf("%s:%d", *host, *port), *user, *password, "")
	if err != nil {
		panic(err)
	}
	rr, err := c.Execute(query)
	if err != nil {
		return errors.Trace(err)
	}
	var fieldList []string

	for i := 0; i < rr.RowNumber(); i++ {
		colName, err := rr.GetString(i, 0)
		if err != nil {
			return err
		}
		fieldList = append(fieldList, colName)
	}
	tableColumnMap.Store(fmt.Sprintf("%s", schemaTableName), fieldList)
	return nil
}

/**
 * @Author lvxin0315@163.com
 * @Description //监听binlog文件变化
 * @Date 3:28 下午 2021/1/15
 * @Param
 * @return
 **/
func listenPos() {
	for {
		time.Sleep(20 * time.Second)
		newPos, err := getMasterPos()
		if err != nil {
			panic(err)
		}
		if pos.Name != newPos.Name {
			pos = newPos
			fmt.Println("listenPos change")
		}
	}
}

/**
 * @Author lvxin0315@163.com
 * @Description 监听解析binlog
 * @Date 3:23 下午 2021/1/15
 * @Param
 * @return
 **/
func binlogSyncer() {
	cfg := replication.BinlogSyncerConfig{
		ServerID:       uint32(*serverID),
		Flavor:         *flavor,
		Host:           *host,
		Port:           uint16(*port),
		User:           *user,
		Password:       *password,
		RawModeEnabled: *rawMode,
		UseDecimal:     true,
	}

	b := replication.NewBinlogSyncer(cfg)
	defer b.Close()

	s, err := b.StartSync(pos)
	if err != nil {
		fmt.Printf("Start sync error: %v\n", errors.ErrorStack(err))
		return
	}
	for {
		for _, ev := range s.DumpEvents() {
			switch ev.Header.EventType {
			case replication.WRITE_ROWS_EVENTv0:
			case replication.WRITE_ROWS_EVENTv1:
			case replication.WRITE_ROWS_EVENTv2:
				rowEventInfo(ev)

			case replication.UPDATE_ROWS_EVENTv0:
			case replication.UPDATE_ROWS_EVENTv1:
			case replication.UPDATE_ROWS_EVENTv2:
				rowEventInfo(ev)
			}
		}
	}

}

/**
 * @Author lvxin0315@163.com
 * @Description 判断是否被监控
 * @Date 6:03 下午 2021/1/15
 * @Param
 * @return
 **/
func inTableList(name string) bool {
	for _, t := range tableList {
		if t == name {
			return true
		}
	}
	return false
}

/**
 * @Author lvxin0315@163.com
 * @Description 获取table的字段列表
 * @Date 6:53 下午 2021/1/15
 * @Param
 * @return
 **/
func getTableColumnList(name string) []string {
	columnNameList, ok := tableColumnMap.Load(name)
	if !ok {
		// TODO 为毛会没有
		fmt.Println("getTableColumnList: 为毛没有")
		err := getTableFields(name)
		if err != nil {
			panic(err)
		}
		return getTableColumnList(name)
	}
	return columnNameList.([]string)
}

/**
 * @Author lvxin0315@163.com
 * @Description 刷新并获取table的字段列表
 * @Date 7:00 下午 2021/1/15
 * @Param
 * @return
 **/
func refreshAndGetTableColumnList(name string) []string {
	err := getTableFields(name)
	if err != nil {
		panic(err)
	}
	return getTableColumnList(name)
}
