package parser

import (
	"fmt"
	"testing"
)

func TestParseFileToServiceInfo(t *testing.T) {
	protoService := ParseFileToServiceInfo("protos/ai_service.proto")
	fmt.Printf("%+v", protoService)
}
