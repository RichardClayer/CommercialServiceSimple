package request

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type Params map[string]string

// GetParams 获取请求参数
func GetParams(r *http.Request) (p Params, err error) {
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil, err
    }

    p = make(Params)

    if err = json.Unmarshal(b, &p); err != nil {
        return nil, err
    }

    return
}
