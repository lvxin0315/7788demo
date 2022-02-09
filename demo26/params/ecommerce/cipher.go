package ecommerce

import (
	"fmt"
	"github.com/go-pay/gopay/wechat/v3"
)

// 字段需进行加密处理
type CipherText string

func V3EncryptText(client *wechat.ClientV3, text string) CipherText {
	cipherText, err := client.V3EncryptText(text)
	if err != nil {
		fmt.Println("v3EncryptText is error: ", err)
		return ""
	}
	return CipherText(cipherText)
}
