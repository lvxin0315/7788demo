package utils

import "fmt"

func MapToParamsString(params map[string]interface{}) string {
	paramString := "?t=1"
	for k, v := range params {
		paramString += fmt.Sprintf("&%s=%v", k, v)
	}
	return paramString
}
