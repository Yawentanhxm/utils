# Yawentanhxm/utils
A simple and easy go tools for slice to map, map to slice

## 1. Install
```shell
 go get -u github.com/Yawentanhxm/utils
```
  

## 2. Getting Started
Traditional Usage
```go
package conv

import (
	"fmt"
	"testing"
)

type TT struct {
	id   float64
	name string
}
func Test_Struct(t *testing.T) {
	data := make([]TT, 0)
	data = append(data, TT{
		id:   0.54,
		name: "数据一",
	})
	data = append(data, TT{
		id:   0.54,
		name: "数据二",
	})
	data = append(data, TT{
		id:   2.34,
		name: "数据三",
	})
	toMap := SliceToMap("id", data)
	fmt.Println(toMap)
}
```

## Features

* 支持Slice指定key转换成map
* 支持带有剔除能力的push