package main

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"time"
)

func main() {
	var (
		mchID                      = "*****"   // 商户号
		mchCertificateSerialNumber = "*******" // 商户证书序列号
		mchAPIv3Key                = "******"  // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	//svc := certificates.CertificatesApiService{Client: client}
	//resp, result, err := svc.DownloadCertificates(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)

	svc := native.NativeApiService{
		Client: client,
	}
	req := native.PrepayRequest{
		Appid:         core.String("wx280d7f5764e32755"),
		Mchid:         core.String("1605344844"),
		Description:   core.String("Image形象店-深圳腾大-QQ公仔"),
		OutTradeNo:    core.String("1217752501201407033233368019"),
		TimeExpire:    core.Time(time.Now()),
		Attach:        core.String("自定义数据说明"),
		NotifyUrl:     core.String("https://www.weixin.qq.com/wxpay/pay.php"),
		GoodsTag:      core.String("WXG"),
		SupportFapiao: core.Bool(false),
		Amount: &native.Amount{
			Currency: core.String("CNY"),
			Total:    core.Int64(1),
		},
		Detail: &native.Detail{
			CostPrice: core.Int64(608800),
			GoodsDetail: []native.GoodsDetail{native.GoodsDetail{
				GoodsName:        core.String("iPhoneX 256G"),
				MerchantGoodsId:  core.String("ABC"),
				Quantity:         core.Int64(1),
				UnitPrice:        core.Int64(828800),
				WechatpayGoodsId: core.String("1001"),
			}},
			InvoiceId: core.String("wx123"),
		},
		SettleInfo: &native.SettleInfo{
			ProfitSharing: core.Bool(false),
		},
		SceneInfo: &native.SceneInfo{
			DeviceId:      core.String("013467007045764"),
			PayerClientIp: core.String("14.23.150.211"),
			StoreInfo: &native.StoreInfo{
				Address:  core.String("广东省深圳市南山区科技中一道10000号"),
				AreaCode: core.String("440305"),
				Id:       core.String("0001"),
				Name:     core.String("腾讯大厦分店"),
			},
		},
	}
	resp, result, err := svc.Prepay(ctx, req)
	if err != nil {
		panic(err)
	}
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
}
