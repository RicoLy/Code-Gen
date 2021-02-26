package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// 方法信息 rpc[\s]*([\w]*)\((\w*)\)\s*\w*\s*\((\w*)\)
type Method struct {
	MethodName string // 方法名
	Param      string // 参数
	Returns    string // 返回值
}

// 服务信息  service\s*(\w*)\s*{
type ProtoService struct {
	ServiceName string
	Methods     []Method
}

// 解析文件
func ParseFileToServiceInfo(fileName string) (protoService ProtoService) {
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
func ParseServiceInfo(contentStr string) (protoService ProtoService) {
	serviceName := ParseServiceName(contentStr)
	protoService.ServiceName = serviceName
	protoService.Methods = ParseServiceMethods(contentStr)
	return
}

// 解析方法列表
func ParseServiceMethods(contentStr string) (methods []Method) {
	ret := regexp.MustCompile(`rpc[\s]*([\w]*)\((\w*)\)\s*\w*\s*\((\w*)\)`)
	result := ret.FindAllStringSubmatch(contentStr, -1)
	methods = make([]Method, 0)

	for _, str := range result {
		method := Method{}
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
