package utils

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"path"
)

const (
	endpoint        = ""
	accessKeyId     = ""
	accessKeySecret = ""
	bucketName      = ""
	bucketUrl       = ""
	bucketDir       = ""
)

const ossStyleUri = "?x-oss-process=style/index_banner"

func NewOssClient() *OssClient {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}

	return &OssClient{
		client: client,
		bucket: bucket,
	}
}

type ossClientReader struct {
	fileName string
	reader   io.Reader
}

type OssClient struct {
	client *oss.Client
	bucket *oss.Bucket
}

func (o *OssClient) Upload(filePath string) (ossUrl string, err error) {
	fileName := path.Base(filePath)
	err = o.bucket.PutObjectFromFile(bucketDir+"/"+fileName, filePath)
	if err != nil {
		return
	}
	ossUrl = bucketUrl + "/" + bucketDir + "/" + fileName + ossStyleUri
	return
}

func (o *OssClient) UploadObject(reader ossClientReader) (ossUrl string, err error) {
	err = o.bucket.PutObject(bucketDir+"/"+reader.fileName, reader.reader)
	if err != nil {
		return
	}
	ossUrl = bucketUrl + "/" + bucketDir + "/" + reader.fileName + ossStyleUri
	return
}

func (o *OssClient) urlToReader(url string) (reader ossClientReader, err error) {
	// 下载远程文件
	body, err := HttpHandle(url, http.MethodGet, nil, nil, nil)
	if err != nil {
		return
	}
	// 远程文件名
	ext := path.Ext(url)
	//if ext == "" {
	//	ext = ".jpg"
	//}
	// 文件随机名称, 暂时用UUID
	u := uuid.NewV4()
	reader.fileName = u.String() + ext
	reader.reader = bytes.NewReader(body)
	return
}

func (o *OssClient) UploadUrlFile(url string) (ossUrl string, err error) {
	// 获取远程文件
	reader, err := o.urlToReader(url)
	if err != nil {
		return
	}
	// 开始上传oss
	ossUrl, err = o.UploadObject(reader)
	if err != nil {
		return
	}

	return
}
