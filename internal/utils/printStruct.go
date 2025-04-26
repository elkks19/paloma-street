package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(s... interface{}) {
	for _, v := range s {
		j, _ := json.MarshalIndent(v, "", "\t")
		fmt.Println(string(j))
		fmt.Println()
	}
}

