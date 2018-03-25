package util

import (
	"github.com/pquerna/ffjson/ffjson"
)

/**
对象转换成json
*/
func ToJson(obj interface{}) (string, error) {
	jsobByte, err := ffjson.Marshal(obj)
	if err != nil {
		println("obj to json fail...\n", err)
		return "", err
	}
	return string(jsobByte), err
}
