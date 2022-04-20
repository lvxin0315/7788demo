package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const (
	endpoint        = "oss-cn-beijing.aliyuncs.com"
	accessKeyId     = "LTAI4G29MaaWJwTwQbTBdyVE"
	accessKeySecret = "V2YXPNnDFNJIYAiYvUFlRijzoicYYN"
	bucketName      = "go-edu"
	bucketUrl       = "https://go-edu.oss-cn-beijing.aliyuncs.com"
	bucketDir       = "goods"
)

const temp = "./tmp/"

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
	ossUrl = bucketUrl + "/" + fileName + ossStyleUri
	return
}

func (o *OssClient) UploadUrlFile(url string) (ossUrl string, err error) {
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
	localFilePath := temp + u.String() + ext
	err = ioutil.WriteFile(localFilePath, body, 0755)

	// 开始上传oss
	ossUrl, err = o.Upload(localFilePath)
	if err != nil {
		return
	}

	// 删除无用文件
	_ = os.Remove(localFilePath)
	return
}
