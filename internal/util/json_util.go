package util

import (
	"encoding/json"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"
)

func ToString(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("序列化json字符串失败, error:%v, data:%v", err, obj)
	}
	return string(str)
}

func ToObject(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		log.Errorf("反序列化json失败, error:%v, data:%v", err, str)
	}
	return err
}
