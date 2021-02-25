package parser

import (
	"fmt"
	"testing"
)

func TestPassProto(t *testing.T) {

	messages := ParseFileToMessage("protos/models.proto")
	for _, message := range messages {
		fmt.Printf("%+v \n", message)
	}
	fmt.Println("---------------")
	structInfos := MessagesToStructInfos(messages)
	for _, info := range structInfos {
		fmt.Printf("%+v \n", info)
	}
}