package services

import (
    "encoding/json"
    "net/http"
)

func JSAPIPay(w http.ResponseWriter, r *http.Request) {
    var m = make(map[string]string)
    m["accept"] = r.Header.Get("Accept")

    respData, err := json.Marshal(m)
    if err !=nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if _, err := w.Write(respData); err != nil {
        panic(err)
    }
}