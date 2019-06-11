package main

import (
    "fmt"
    "./g"
)

func main() {
    fmt.Println("Start")
    jsons := []byte(`{"a": {"b": {"c": [1,2,3]}}}`)
    jsons2 := []byte(`{"a": {"c": {"c": [1,2,3]}}}`)
    value := []byte(`[3,2,1]`)
    strin :=  []string{"a", "b", "c"}
    
    fmt.Println("json>",string(jsons))
    fmt.Println("json2>",string(jsons2))
    fmt.Println("path>",strin)
    fmt.Println("value>",string(value))
    fmt.Println("==========")
    
    result := g.Has(jsons, strin)
    fmt.Println("Has func>",result)
    result1,err := g.Get(jsons, strin)
    
    fmt.Println("Get func>",string(result1),err)
    
    result2,err2 := g.Set(jsons, strin,value)
    
    fmt.Println("Set func>",string(result2),err2)
    
    result3,err3 := g.Remove(jsons, strin)
    
    fmt.Println("Remove func>",string(result3),err3)
    
    result4,err4 := g.Merge(jsons, jsons2)
    
    fmt.Println("Merge func>",string(result4),err4)
}