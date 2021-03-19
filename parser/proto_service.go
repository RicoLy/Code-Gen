package parser

import (
	"code-gen/micro/models"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)


// 解析文件
func ParseFileToServiceInfo(fileName string) (protoService models.ProtoService) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentStr := string(contentByte)

	protoService = ParseServiceInfo(contentStr)
	return
}

// 解析服务信息
func ParseServiceInfo(contentStr string) (protoService models.ProtoService) {
	serviceName := ParseServiceName(contentStr)
	protoService.ServiceName = serviceName
	protoService.Methods = ParseServiceMethods(contentStr)
	return
}

// 解析方法列表
func ParseServiceMethods(contentStr string) (methods []models.Method) {
	ret := regexp.MustCompile(`rpc[\s]*([\w]*)\((\w*)\)\s*\w*\s*\((\w*)\)`)
	result := ret.FindAllStringSubmatch(contentStr, -1)
	methods = make([]models.Method, 0)

	for _, str := range result {
		method := models.Method{}
		method.MethodName = str[1]
		method.Param = str[2]
		method.Returns = str[3]
		methods = append(methods, method)
	}

	return
}

// 解析服务名
func ParseServiceName(contentStr string) (serviceName string) {
	// 解析serviceName
	ret := regexp.MustCompile(`service\s*(\w*)\s*{`)
	result := ret.FindAllStringSubmatch(contentStr, -1)
	serviceName = result[0][1]
	return
}
