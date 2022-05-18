package parser

import (
	"code-gen/config"
	"code-gen/micro/models"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)


// 解析文件
func ParseFileToMessage(fileName string) (messages []models.Message) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentStr := string(contentByte)

	messages = make([]models.Message, 0)
	messages = ParseMessageInfos(contentStr)
	return
}

// 解析Messages
func ParseMessageInfos(contentStr string) (messages []models.Message) {
	messages = make([]models.Message, 0)

	nameMetaInfo := ParseMessageNameAndMetaInfo(contentStr)
	for name, info := range nameMetaInfo {
		message := models.Message{}
		message.MessageName = name
		message.MessageMeta = info

		message.ElementInfos = ParseMessageElements(info)
		messages = append(messages, message)
	}
	return messages
}

// 解析字段
func ParseMessageElements(metaInfo string) (elements []models.ElementInfo) {
	elements = make([]models.ElementInfo, 0)

	eleTags := ParseMessageElementsInfo(metaInfo)
	for ele, tags := range eleTags {
		eleInfo := models.ElementInfo{}
		eleInfo.Tags = make(map[string]string)

		kv := strings.Split(strings.TrimSpace(ele), " ")
		if kv[0] != "repeated" { // 是否是slice
			eleInfo.Name = kv[1]
			eleInfo.Type = config.ProtoTypeToGoType[strings.TrimSpace(kv[0])] // 转换成goType
		} else {
			eleInfo.Name = strings.TrimSpace(kv[2])
			if _, ok := config.ProtoTypeToGoType[strings.TrimSpace(kv[1])]; ok {
				eleInfo.Type = "[]" + config.ProtoTypeToGoType[strings.TrimSpace(kv[1])] // 有则转换成goType slice
			} else {
				eleInfo.Type = "[]" + strings.TrimSpace(kv[1]) //没有则转换成自定义类型 slice
			}
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
	ret := regexp.MustCompile(`message[\s]*([\S]*)[\s]*{[\s\S]*?}`)
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
func MessagesToStructInfos(messages []models.Message) (structInfos []models.StructInfo) {
	structInfos = make([]models.StructInfo, 0)
	for _, message := range messages {
		structInfo := models.StructInfo{}
		structInfo.StructName = message.MessageName // 结构体名
		structInfo.ElementInfos = message.ElementInfos // 赋值元素信息
		structInfo.HasWrite = false  // 默认未写入
		structInfos = append(structInfos, structInfo)
	}

	return
}

// 解析文件为全局structInfo map
func ParseFileToGlobalStructMap(fileName string) (structMap map[string]models.StructInfo) {
	structMap = make(map[string]models.StructInfo)
	messages := ParseFileToMessage(fileName)
	structInfos := MessagesToStructInfos(messages)

	for _, info := range structInfos {
		structMap[info.StructName] = info
	}

	return
}