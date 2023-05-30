package pkg

import (
	"fmt"
	"sync"
	"testing"
)

// 类型断言转换为原本类型
func GetValue(v any) (res string) {
	switch v.(type) {
	case int:
		res = fmt.Sprintf("%v", v.(int))
	case string:
		res = fmt.Sprintf("%v", v.(string))
	default:
		res = fmt.Sprintf("%v", "v is of unknown type")
	}
	return
}

func TestSyncMap(t *testing.T) {
	sMap := sync.Map{}
	k1 := "aaa"
	k2 := "ccc"
	v1 := "bbb"
	v2 := 2
	sMap.Store(k1, v1)
	sMap.Store(k2, v2)
	if res, ok := sMap.Load(k1); !ok {
		fmt.Println("cannot find key", v1)
	} else {
		fmt.Println(res)
	}

	sMap.Range(func(key, value any) bool {
		println(GetValue(key), GetValue(value))
		return true
	})

}
