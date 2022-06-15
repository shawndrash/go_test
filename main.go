package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	value := ""
	value, _ = sjson.Set(value, "name", "travis")
	value, _ = sjson.Set(value, "status", 1)
	value, _ = sjson.Set(value, "user_info.ttt", 0.3)
	value, _ = sjson.Set(value, "user_info.sss", 56)
	fmt.Println(value)
}
