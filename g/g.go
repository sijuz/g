package g

import (
    "encoding/json"
    "errors"
    "github.com/imdario/mergo"
)

func Has(jsons []byte, path []string) bool {
    if _,err := Get(jsons, path); err != nil {
        return false
    }
    return true
}

func Get(jsons []byte, path []string) ([]byte, error) { 
    var dat map[string]interface{}
    json.Unmarshal(jsons, &dat)
    for cur, _ := range path {
        if (cur == len(path)-1) {
            rt, err := json.Marshal(dat[path[cur]])
            return rt, err
        } else {
            dat = dat[path[cur]].(map[string]interface{})
        }
    }
    return nil, errors.New("JSON ERROR")
}

func Set(jsons []byte, path []string, value []byte) ([]byte, error) {
    var dat,next map[string]interface{}
    
    if err := json.Unmarshal(jsons, &dat); err != nil {
        return nil, err
    }
    next = dat
    
    
    for j := 0; j < len(path); j++ {
        if (j == len(path)-1) {
            if next1, ok := next[path[j]]; ok {
                _ = next1
                rt, err := json.Marshal(next1)
                if err != nil {
                    return nil, err
                }
                _ = rt
                next[path[j]] = string(value)
            } else {
                return nil, errors.New("Not Found")
            }
            
        } else {
            if next2, ok := next[path[j]].(map[string]interface{}); ok {
                next = next2
                _ = next2
            } else {
                return nil, errors.New("Not Found")
            }
        }
    }
    
    for j := len(path)-2; j > 0; j-- {
        result1,err := Get(jsons, path[0:j])
        _ = err
        var dat2 map[string]interface{}
    
        if err := json.Unmarshal(result1, &dat2); err != nil {
            return nil, err
        }
        dat2[path[j]] = next
        next = dat2
    }
    dat[path[0]] = next
    next = dat
    
    rt1, err1 := json.Marshal(next)
    if err1 != nil {
        return nil, err1
    }
    
    return rt1, nil
}

//func SetTEST(jsons []byte, path []string, value []byte) ([]byte, error) {
//    var dat,cc map[string]interface{}
//    json.Unmarshal(jsons, &dat)
//    for cur, _ := range path {
//        if (cur == len(path)-1) {
//            //_, err := json.Marshal(dat[path[cur]])
//            dat[path[cur]] = string(value)
//            
//            for j := len(path)-2; j > 0; j-- {
//
//                var local map[string]interface{}
//                ret, _ := Get(jsons, path[0:j])
//                fmt.Println(string(ret))
//                json.Unmarshal(ret, &local)
//                local[path[j]] = local
//                cc = local
//              
//                
//            }
//            dat[path[0]] = cc
//            fmt.Println(dat)
//            rt, err := json.Marshal(dat)
//            if err != nil {
//                return nil, errors.New("JSON ERROR")
//            }
//            return rt, nil
//            
//            
//        } else {
//            dat = dat[path[cur]].(map[string]interface{})
//        }
//    }
//    
//    
//    return nil, errors.New("JSON ERROR")
//    
//}

func Remove(jsons []byte, path []string) ([]byte, error) {
    var dat,next,dat3 map[string]interface{}
    
    if err := json.Unmarshal(jsons, &dat); err != nil {
        return nil, err
    }
    next = dat
    
    
    for j := 0; j < len(path); j++ {
        if (j == len(path)-1) {
            if next1, ok := next[path[j]]; ok {
                _ = next1
                rt, err := json.Marshal(next1)
                if err != nil {
                    return nil, err
                }
                _ = rt
                nul := []byte(`{}`)
                if err4 := json.Unmarshal(nul, &dat3); err != nil {
                    return nil, err4
                }
                next = dat3
            } else {
                return nil, errors.New("Not Found")
            }
            
        } else {
            if next2, ok := next[path[j]].(map[string]interface{}); ok {
                next = next2
                _ = next2
            } else {
                return nil, errors.New("Not Found")
            }
        }
    }
    
    for j := len(path)-2; j > 0; j-- {
        result1,err := Get(jsons, path[0:j])
        _ = err
        var dat2 map[string]interface{}
    
        if err := json.Unmarshal(result1, &dat2); err != nil {
            return nil, err
        }
        dat2[path[j]] = next
        next = dat2
    }
    dat[path[0]] = next
    next = dat
    
    rt1, err1 := json.Marshal(next)
    if err1 != nil {
        return nil, err1
    }
    
    return rt1, nil
}

func Merge(jsons []byte, jsons2 []byte) ([]byte, error) {
    var dat,dat2 map[string]interface{}
    
    json.Unmarshal(jsons, &dat)
    json.Unmarshal(jsons2, &dat2)
    
    mergo.Merge(&dat, dat2)
    rt, err := json.Marshal(dat)
    return rt, err
}