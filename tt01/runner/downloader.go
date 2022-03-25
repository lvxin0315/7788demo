package runner

import (
	"fmt"
	"github.com/lvxin0315/7788demo/tt01/config"
	"github.com/lvxin0315/7788demo/tt01/model"
	"io"
	"net/http"
	"os"
	"sync"
)

// BatchWork 执行批量下载文件
func BatchWork(res model.VideoListResponse) (err error, total int) {
	// 设置协程数
	wg := sync.WaitGroup{}
	wg.Add(len(res.ItemList))
	// 执行
	for _, item := range res.ItemList {
		go func(videoItem model.VideoItem) {
			err := downloadMP4File(videoItem)
			wg.Done()
			if err != nil {
				fmt.Println(fmt.Sprintf("ID:%s is error: %s", item.Id, err.Error()))
			} else {
				total++
			}
		}(item)
	}
	wg.Wait()
	return
}

// 下载mp4文件
func downloadMP4File(item model.VideoItem) error {
	resp, err := http.Get(item.Video.DownloadAddr)
	if err != nil {
		fmt.Println("downloadMP4File http.Get error: ", err)
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(config.VideoDir + string(os.PathSeparator) + item.Video.Id + ".mp4")
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
