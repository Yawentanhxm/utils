package conv

import (
	"fmt"
	"testing"
)

func Test_Struct(t *testing.T) {
	type TT struct {
		id   float64
		name string
	}
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

func Test_PtrStruct(t *testing.T) {
	type TT struct {
		id   int64
		name string
	}
	data := make([]*TT, 0)
	data = append(data, &TT{
		id:   0,
		name: "数据一",
	})
	data = append(data, &TT{
		id:   1,
		name: "数据二",
	})
	data = append(data, &TT{
		id:   2,
		name: "数据三",
	})
	toMap := SliceToMap("name", data)
	fmt.Println(toMap)
}
