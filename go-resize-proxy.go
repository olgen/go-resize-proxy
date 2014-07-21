package main

import (
    "net/http"
    "os"
    "log"
    "github.com/blloon/go-resize-proxy/resizeproxy"
)

func main(){
    handler := resizeproxy.NewProxyHandler(originSetting())
    http.Handle("/", handler)

    port := portSetting()
    log.Printf("Handler listening on port:%v", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func portSetting() string {
    port := os.Getenv("PORT")
    if port == "" {
        panic("No PORT env-var given!")
    }
    return ":" + port
}

func originSetting() string {
    origin := os.Getenv("ORIGIN")
    if origin == "" {
        panic("No ORIGIN env-var given!")
    }
    return origin
}
