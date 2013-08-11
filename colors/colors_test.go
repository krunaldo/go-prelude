package colors

import (
    "fmt"
    "testing"
)

func TestColors(t *testing.T) {
    fmt.Println(Yellow("Inside yellow is ", Red("red text"), " and ", Blue("blue ", Cyan("(and cyan) "), "text"), ". This is still yellow."))
}
