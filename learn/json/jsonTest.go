package json

import (
	"encoding/json"
)

func main() {
	/*var obj interface{}
	humanStr := "aaa"
	err := json.Unmarshal([]byte(humanStr), &obj)
	objMap, ok := obj.(map[string]interface{})
	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}*/
}

type Human struct {
	Age  int
	name string
}

func unmarshal2Struct(humanStr string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(humanStr), &h)
	if err != nil {
		println(err)
	}
	return h
}

func marshal2JsonString(h Human) string {
	h.Age = 30
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(updatedBytes)
}
