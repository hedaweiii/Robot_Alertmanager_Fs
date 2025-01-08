package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// 定义字段结构
type Field struct {
	CustomField1 string
	CustomField2 string
}

type Data struct {
	Fields       []Field
	FieldsLength int
}

// GenerateJSON 函数，读取模板并生成 JSON 字符串
func GenerateJSON(fields []Field) (string, error) {
	// 1. 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前工作目录失败: %v", err)
	}

	// 2. 构造模板文件的相对路径
	filePath := wd + "/../../resources/card.json"

	// 3. 读取模板文件
	tmplData, err := ioutil.ReadFile(filePath) // 读取模板文件
	if err != nil {
		return "", fmt.Errorf("读取模板文件失败: %v", err)
	}

	// 4. 创建模板并解析
	tmpl, err := template.New("jsonTemplate").Funcs(template.FuncMap{
		"last": func(index int, length int) bool {
			// 判断是否是最后一个元素
			return index == length-1
		},
	}).Parse(string(tmplData))
	if err != nil {
		return "", fmt.Errorf("解析模板失败: %v", err)
	}

	// 5. 渲染模板
	var buf bytes.Buffer

	data := Data{
		Fields:       fields,
		FieldsLength: len(fields),
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("执行模板失败: %v", err)
	}

	// 6. 返回生成的 JSON 字符串
	return buf.String(), nil
}
