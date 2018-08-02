package testjson

import (
	"encoding/json"
	"fmt"
)

func testJson() {
    m := map[string] []string {"001":{"li", "male", "student"}, "002":{"wang", "female", "worker"}}

     data, err := json.Marshal(m)
     if err == nil {
        fmt.Printf("%s\n", data)
    }

    decodeM := map[string] []string{}
    if err := json.Unmarshal(data, &decodeM); err != nil {
        fmt.Println("Unmarshal failed")
    }

    fmt.Printf("%s", decodeM)
}