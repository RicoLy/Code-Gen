package parser

import (
	"code-gen/config"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// proto信息
type Message struct {
	MessageMeta  string        // 元数据
	MessageName  string        // 消息名
	ElementInfos []ElementInfo // proto字段
}

// 结构体信息
type StructInfo struct {
	StructName   string
	ElementInfos []ElementInfo // golang字段
}

// 字段信息
type ElementInfo struct {
	Name string            // 名称
	Type string            // golang 数据类型
	Tags map[string]string // 标签信息 tag | value  json:"pid" form:"pid"
}

// 解析文件
func ParseFileToMessage(fileName string) (messages []Message) {
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

	messages = make([]Message, 0)
	messages = ParseMessageInfos(contentStr)
	return
}

// 解析Messages
func ParseMessageInfos(contentStr string) (messages []Message) {
	messages = make([]Message, 0)

	nameMetaInfo := ParseMessageNameAndMetaInfo(contentStr)
	for name, info := range nameMetaInfo {
		message := Message{}
		message.MessageName = name
		message.MessageMeta = info

		message.ElementInfos = ParseMessageElements(info)
		messages = append(messages, message)
	}
	return messages
}

// 解析字段
func ParseMessageElements(metaInfo string) (elements []ElementInfo) {
	elements = make([]ElementInfo, 0)

	eleTags := ParseMessageElementsInfo(metaInfo)
	for ele, tags := range eleTags {
		eleInfo := ElementInfo{}
		eleInfo.Tags = make(map[string]string)

		kv := strings.Split(strings.TrimSpace(ele), " ")
		if kv[0] != "repeated" { // 是否是slice
			eleInfo.Name = kv[1]
			eleInfo.Type = config.ProtoTypeToGoType[strings.TrimSpace(kv[0])] // 转换成goType
		} else {
			eleInfo.Name = strings.TrimSpace(kv[2])
			eleInfo.Type = "[]" + config.ProtoTypeToGoType[strings.TrimSpace(kv[1])] // 转换成goType slice
		}

		sTag := strings.Split(strings.ReplaceAll(strings.ReplaceAll(tags, ":", " "), "\"", ""), ",")
		for _, s := range sTag {
			tkv := strings.Split(strings.TrimSpace(s), " ")
			eleInfo.Tags[strings.TrimSpace(tkv[0])] = tkv[1]
		}
		elements = append(elements, eleInfo)
	}

	return elements
}

// 解析messageName 和元数据 字符串
func ParseMessageNameAndMetaInfo(contentStr string) (nameMetaInfo map[string]string) {
	// 解析proto message ---> name | metadata
	ret := regexp.MustCompile(`message[\s]*([\S]*){[\s\S]*?}`)
	result := ret.FindAllStringSubmatch(contentStr, -1)
	nameMetaInfo = make(map[string]string, len(result))
	for _, str := range result {
		name := str[1]
		nameMetaInfo[name] = str[0]
	}
	return nameMetaInfo
}

// 解析标签和元素字符串
func ParseMessageElementsInfo(contentStr string) (eleTags map[string]string) {
	ret := regexp.MustCompile(`@inject_tag:( .*)[\s]*(.*)`)
	result := ret.FindAllStringSubmatch(contentStr, -1)
	eleTags = make(map[string]string)
	for _, str := range result {
		ele := str[2]
		eleTags[ele] = str[1]
	}
	return eleTags
}

// Message 转换成 Struct
func MessagesToStructInfos(messages []Message) (structInfos []StructInfo) {
	structInfos = make([]StructInfo, 0)
	for _, message := range messages {
		structInfo := StructInfo{}
		structInfo.StructName = message.MessageName
		structInfo.ElementInfos = message.ElementInfos
		structInfos = append(structInfos, structInfo)
	}

	return
}
