package main

import (
	"fmt"
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
	"github.com/q191201771/naza/pkg/nazalog"
	"os"
	"time"
)

// 间隔时间
var sleepTime = 10 * time.Minute

// 几点开始
var beginHour = 6

// 几点结束
var endHour = 18

const (
	U1 = "rtmp://xxxx/xxx/xxx"
)

func pull(url string, filename string) {
	fmt.Println("url:", url)
	fmt.Println("filename:", filename, " 开始")
	var (
		w   httpflv.FLVFileWriter
		err error
	)

	if filename != "" {
		err = w.Open(filename)
		nazalog.Assert(nil, err)
		defer w.Dispose()
		err = w.WriteRaw(httpflv.FLVHeader)
		nazalog.Assert(nil, err)
	}

	session := rtmp.NewPullSession(func(option *rtmp.PullSessionOption) {
		option.PullTimeoutMS = 10000
		option.ReadAVTimeoutMS = 10000
	})

	err = session.Pull(url, func(msg base.RTMPMsg) {
		if filename != "" {
			tag := remux.RTMPMsg2FLVTag(msg)
			err := w.WriteTag(*tag)
			nazalog.Assert(nil, err)
		}
	})
	if err != nil {
		nazalog.Errorf("pull failed. err=%+v", err)
		return
	}

	go func() {
		time.Sleep(sleepTime)
		session.Dispose()
		return
	}()
	err = <-session.WaitChan()
	nazalog.Debugf("< session.WaitChan. [%s] err=%+v", session.UniqueKey(), err)
}

func start(name string, url string) {
	for {
		//判断是否上班
		workTime()
		nowDayStr := time.Now().Format("20060102")
		nowTimeStr := time.Now().Format("15点04分05秒")
		autoDir(fmt.Sprintf("%s/%s/", name, nowDayStr))
		pull(url, fmt.Sprintf("%s/%s/%s.flv", name, nowDayStr, nowTimeStr))
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func autoDir(path string) {
	if exist, _ := pathExists(path); exist {
		return
	}
	// 创建
	err := os.MkdirAll(path, 0777)
	fmt.Println(err)
}

func workTime() {
	for {
		nowHour := time.Now().Hour()
		if nowHour >= beginHour && nowHour <= endHour {
			return
		}
		// 没到点
		time.Sleep(10 * time.Minute)
	}
}

func main() {
	go start("U1", U1)
	select {}
}
