package newpostech

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/newpostech/requests"
)

func TestPlay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AccessKeyId = os.Getenv("newpostech_AccessKeyId")
	client.Config.AccessKey = os.Getenv("newpostech_AccessKey")
	// client.Config.DeviceId = "0047000021"
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "play"
	request.BizContent = map[string]interface{}{
		"device_id":     "H81220919000077",
		"request_data":  "微信收款10012.35元",
		"push_template": "00",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	r, err := response.GetVerifySignDataMap()
	fmt.Println("TestPlay", r, err)
	t.Log(r, err, "|||")
}
