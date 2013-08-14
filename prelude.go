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

type Tuple2 struct {
    A, B interface{}
}
func (t Tuple2) Get() (a, b interface{}) {
    return t.A, t.B
}

type Tuple3 struct {
    A, B, C interface{}
}
func (t Tuple3) Get() (a, b, c interface{}) {
    return t.A, t.B, t.C
}

type Tuple4 struct {
    A, B, C, D interface{}
}
func (t Tuple4) Get() (a, b, c, d interface{}) {
    return t.A, t.B, t.C, t.D
}
