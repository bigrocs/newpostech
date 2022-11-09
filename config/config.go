package config

// 服务器URL ： http://posp.newposp.com:26056/its/chapter2/section2-1.html

type Config struct {
	AccessKeyId string `json:"access_key_id"` // 访问密钥 ID
	AccessKey   string `json:"access_key"`    // 开发者密钥
	DeviceId    string `json:"device_id"`     // 产品KEY
	Sandbox     bool   `json:"sandbox"`       // 沙盒
}
