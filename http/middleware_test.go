package http

import (
	"context"
	"fmt"
	"testing"
)

func TestGetOpenapiToken(t *testing.T) {
	token, err := GetAppToken(context.Background())
	fmt.Println(err)
	fmt.Println(token)
}
