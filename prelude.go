package prelude

func Arr(a ...interface{}) []interface{} {
    return a
}

func Dict(a ...interface{}) map[string]interface{} {
    length := len(a)
    m := make(map[string]interface{}, length)
    for i:=0; i<length; i+=2 {
        m[a[i].(string)] = a[i+1]
    }
    return m
}
