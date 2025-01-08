package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"gopkg.in/yaml.v2"
)

// Config 表示 Prometheus 配置
type Config struct {
	PrometheusURL string `yaml:"prometheus_url"` // Prometheus 地址
}

// QueryResult 表示 Prometheus 查询结果
type QueryResult struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values [][]interface{}   `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

// loadConfig 从配置文件读取 Prometheus URL
func loadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return &config, nil
}

// executePromQL 根据 PromQL 查询 Prometheus 数据
func executePromQL(prometheusURL, promQL string) (*QueryResult, error) {
	// 创建 Prometheus 客户端
	client, err := api.NewClient(api.Config{Address: prometheusURL})
	if err != nil {
		return nil, fmt.Errorf("创建 Prometheus 客户端失败: %v", err)
	}

	// 创建 Prometheus API 实例
	api := v1.NewAPI(client)

	// 设置查询超时时间为默认 3 秒
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 执行查询
	result, warnings, err := api.Query(ctx, promQL, time.Now())
	if err != nil {
		return nil, fmt.Errorf("PromQL 查询失败: %v", err)
	}

	// 打印警告信息（如果有）
	if len(warnings) > 0 {
		log.Printf("警告: %v\n", warnings)
	}

	// 将查询结果转为 JSON 格式
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("查询结果编码失败: %v", err)
	}

	// 解码 JSON 为 QueryResult 结构
	var queryResult QueryResult
	if err := json.Unmarshal(resultBytes, &queryResult); err != nil {
		return nil, fmt.Errorf("查询结果解码失败: %v", err)
	}

	return &queryResult, nil
}
