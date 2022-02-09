package main

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/lvxin0315/7788demo/demo26/params/ecommerce"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"time"
)

func main() {
	var (
		mchID                      = "1572641601"                               // 商户号
		mchCertificateSerialNumber = "18F5E093713C6F200129EDD207F17D9461229159" // 商户证书序列号
		mchAPIv3Key                = "1sW10Ngu6h3z4zkKSDQXRUpk80c6efPL"         // 商户APIv3密钥
	)
	privateBytes, err := ioutil.ReadFile("./apiclient_key.pem")
	if err != nil {
		panic(err)
	}
	client, err := wechat.NewClientV3(mchID, mchCertificateSerialNumber, mchAPIv3Key, string(privateBytes))
	if err != nil {
		panic(err)
	}

	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		panic(err)
	}

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn

	ctx := context.Background()
	//// 上传身份证
	//fileContent, err := ioutil.ReadFile("./身份证-正面.jpeg")
	//h := sha256.New()
	//h.Write(fileContent)
	//sha256Str := hex.EncodeToString(h.Sum(nil))
	//
	//rsp, err := client.V3MediaUploadImage(ctx, "身份证-正面.jpeg", sha256Str,&util.File{
	//	Name:    "身份证-正面.jpeg",
	//	Content: fileContent,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("rsp.Response.MediaId: ", rsp.Response.MediaId)
	//
	//return
	bmParams := new(ecommerce.Ecommerce)
	bmParams.OutRequestNo = fmt.Sprintf("APPLYMENT_%d", time.Now().UnixNano())
	// 主体类型
	bmParams.OrganizationType = ecommerce.ORGANIZATION_TYPE_2401
	// 身份证
	bmParams.IdDocType = ecommerce.IDENTIFICATION_TYPE_MAINLAND_IDCARD
	bmParams.IdCardInfo.IdCardName = ecommerce.V3EncryptText(client, "吕鑫")
	bmParams.IdCardInfo.IdCardNumber = ecommerce.V3EncryptText(client, "210102198903055015")
	// 正面
	bmParams.IdCardInfo.IdCardCopy = "oSwAZ0XP7HstecLokC2Tn37fdsY9kVc8atlK60KJ5SF3cv39nH8lFIlzdqXvqEf-dPZMxoslE_91MOjKpqygghbhegtL_7UJGz9OApmci9I"
	// 背面
	bmParams.IdCardInfo.IdCardNational = "f_2Kl4T_QPWHZsHDPWEY_dOcey8Crky5EBVxFnBV1wQnF8MMAJDLDQ497ricTeYR6x8ERh8K2KbLs-yJg6u8yzr6sIEqejdS8GZqFEYDs1k"
	bmParams.IdCardInfo.IdCardValidTime = "2036-05-18"
	// 结算规则
	bmParams.SettlementInfo.QualificationType = "无" // 微信脑残，没有必须写个 "无"
	// 商户简称
	bmParams.MerchantShortname = "lvxin测试小店"
	// 商户信息
	bmParams.SalesSceneInfo.StoreName = "lvxin测试小店"
	bmParams.SalesSceneInfo.StoreUrl = "https://www.jicanchu.com"
	// 超级管理员
	bmParams.ContactInfo.ContactType = ecommerce.CONTACT_TYPE_65
	bmParams.ContactInfo.ContactName = ecommerce.V3EncryptText(client, "吕鑫")
	bmParams.ContactInfo.ContactIdCardNumber = ecommerce.V3EncryptText(client, "210102198903055015")
	bmParams.ContactInfo.MobilePhone = ecommerce.V3EncryptText(client, "18624088831")

	//fmt.Println(bmParams)
	//return

	// 生成结构
	bm := gopay.BodyMap{}
	err = mapstructure.Decode(bmParams, &bm)
	if err != nil {
		panic(err)
	}

	ecommerceApplyRsp, err := client.V3EcommerceApply(ctx, bm)
	if err != nil {
		panic(err)
	}
	fmt.Println("ecommerceApplyRsp: ", ecommerceApplyRsp)
	fmt.Println("OutRequestNo: ", ecommerceApplyRsp.Response.OutRequestNo)
	fmt.Println("ApplymentId: ", ecommerceApplyRsp.Response.ApplymentId)
}
