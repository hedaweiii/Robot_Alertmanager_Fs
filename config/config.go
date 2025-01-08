package config

import (
	"fmt"
	"io/ioutil"
	"os"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	"gopkg.in/yaml.v2"
)

// Config 结构体定义了配置文件的结构
type Config struct {
	AppId             string `yaml:"AppId"`             // 应用ID
	AppSecret         string `yaml:"AppSecret"`         // 应用密钥
	EncryptKey        string `yaml:"EncryptKey"`        // 加密密钥
	VerificationToken string `yaml:"VerificationToken"` // 验证令牌
	PrometheusURL     string `yaml:"PrometheusURL"`     // Prometheus URL
}

// Config 全局变量，保存加载的配置
var ConfigInstance Config

// LoadConfig 从指定的 YAML 配置文件中加载配置
func LoadConfig(configPath string) (Config, error) {
	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("配置文件不存在: %v", err)
	}

	// 读取配置文件
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 将读取的 YAML 数据解析到 Config 结构体中
	var c Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return Config{}, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 返回加载后的配置
	return c, nil
}

// Client 使用从配置文件加载的 AppId 和 AppSecret 创建一个新的 Lark 客户端
var Client *lark.Client

// init 初始化函数，在程序启动时调用，用于加载配置并创建客户端
func init() {
	// 加载配置并赋值给全局的 ConfigInstance 变量
	var err error
	ConfigInstance, err = LoadConfig("../../config/config.yaml")
	if err != nil {
		fmt.Printf("配置加载失败: %v\n", err)
		return
	}

	// 创建 Lark 客户端
	Client = lark.NewClient(ConfigInstance.AppId, ConfigInstance.AppSecret)
}
