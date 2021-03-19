package parser

import (
	"fmt"
	"testing"
)

func TestParseFileToMessage(t *testing.T) {
	messages := ParseFileToMessage("protos/ai_service.proto")
	for _, message := range messages {
		fmt.Printf("%+v \n", message)
	}
	fmt.Println("---------------")
	structInfos := MessagesToStructInfos(messages)
	for _, info := range structInfos {
		fmt.Printf("%+v \n", info)
	}
}

func TestParseFileToStructMap(t *testing.T) {
	GMap := ParseFileToGlobalStructMap("protos/ai_service.proto")
	for name, info := range GMap {
		fmt.Printf("%+v \n", name)
		info.HasWrite = true
		GMap[name] = info
	}

	for _, info := range GMap {
		fmt.Printf("%+v \n", info)
	}
}