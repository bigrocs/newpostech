package common

import (
	"strings"

	"github.com/micro/go-micro/v2/util/log"
	uuid "github.com/satori/go.uuid"

	"github.com/bigrocs/newpostech/config"
	"github.com/bigrocs/newpostech/requests"
	"github.com/bigrocs/newpostech/responses"
	"github.com/bigrocs/newpostech/util"
)

// Common 公共封装
type Common struct {
	Config   *config.Config
	Requests *requests.CommonRequest
}

// Action 创建新的公共连接
func (c *Common) Action(response *responses.CommonResponse) (err error) {
	return c.Request(response)
}

// APIBaseURL 默认 API 网关
func (c *Common) APIBaseURL() string { // TODO(): 后期做容灾功能
	con := c.Config
	if con.Sandbox { // 沙盒模式
		return "https://horn.newpostech.com:26057/horn/pushmsg"
	}
	return "http://its.newpostech.cn:26080/horn/pushmsg"
}

// Request 执行请求
// AppCode           string `json:"app_code"`             //API编码
// AppId             string `json:"app_id"`               //应用ID
// UniqueNo          string `json:"unique_no"`            //私钥
// PrivateKey        string `json:"private_key"`          //私钥
// newpostechPublicKey string `json:"lin_shang_public_key"` //临商银行公钥
// MsgId             string `json:"msg_id"`               //消息通讯唯一编号，每次调用独立生成，APP级唯一
// Signature         string `json:"Signature"`            //签名值
// Timestamp         string `json:"timestamp"`            //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
// NotifyUrl         string `json:"notify_url"`           //工商银行服务器主动通知商户服务器里指定的页面http/https路径。
// BizContent        string `json:"biz_content"`          //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
// Sandbox           bool   `json:"sandbox"`              // 沙盒
func (c *Common) Request(response *responses.CommonResponse) (err error) {
	con := c.Config
	req := c.Requests
	// 构建配置参数
	params := map[string]interface{}{
		"access_key_id": con.AccessKeyId,
		"request_id":    strings.Replace(uuid.NewV4().String(), "-", "", -1),
		"nonce":         strings.Replace(uuid.NewV4().String(), "-", "", -1),
	}
	for k, v := range req.BizContent {
		params[k] = v
	}
	sign := util.HmacSha1(util.EncodeSignParams(params), con.AccessKey) // 开发签名
	if err != nil {
		return err
	}
	params["sign"] = sign
	log.Info("newpostech:PostJSON", c.APIBaseURL(), params)
	res, err := util.PostJSON(c.APIBaseURL(), params)
	log.Info("newpostech:PostJSON:res", string(res), err)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "string")
	return
}
