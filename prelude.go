package prelude

import (
    "fmt"
    "strings"
)

type Arr []interface{}
type Dict map[string]interface{}

// fmt.Println( Concat(1, 2, []int{4,5,6}, "foo", "bar") )
func Concat(args ...interface{}) string {
    var parts []string
    for _, arg := range args {
        parts = append(parts, fmt.Sprintf("%v", arg))
    }
    return strings.Join(parts, "")
}

