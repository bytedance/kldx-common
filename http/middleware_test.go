package http

import (
	"fmt"
	"testing"
)

func TestGetOpenapiToken(t *testing.T) {
	token, err := GetAppToken()
	fmt.Println(err)
	fmt.Println(token)
}
