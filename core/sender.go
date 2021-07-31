package core

import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"

    "bomb/models"
)

type Send struct {}

type HttpSender struct {
    *Send
}

type HttpsSender struct {
    *Send
}

type TcpSender struct {
    *Send
}

func (s Send) Init () {
    // TODO
}

func (s Send) SendData () {
    // TODO
}

func (s Send) Recv () {
    //TODO
}

func (s *Send) Run (api models.API) {
    // TODO
}


func NewSender(way string) models.Sender {
    fmt.Println("Hello, World!")
    switch way {
    case "http":
        return &HttpSender{}
    case "https": {
        return &HttpsSender{}
    };
    case "tcp": {
        return &TcpSender{}
    };
    default: {
        return &HttpSender{}
    }
    }
}


func (h *HttpSender) SendData (api models.API) {
    switch api.Method {
        case 0: {
            if len(api.Headers) <= 0 {
                resp, err := http.Get(api.Url)
                if err != nil {
                    fmt.Printf("bomb error: %v\n", err)
                } else {
                    defer resp.Body.Close()
                    body, _ := ioutil.ReadAll(resp.Body)
                    fmt.Println(resp.Status, body)
                }
            } else {
                client := http.Client{}
                if api.Data != "" {
                    req, err := http.NewRequest("GET", api.Url, strings.NewReader(api.Data))
                } else {
                    req, err := http.NewRequest("GET", api.Url, nil)
                }
                if err != nil {
                    fmt.Println(err)
                } else {
                    for k, v := range api.Headers {
                        req.Header.Add(k, v)
                    }
                    resp, err := client.Do(req)
                    if err == nil {
                        defer resp.Body.Close()
                        body, _ := ioutil.ReadAll(resp.Body)
                        fmt.Println(resp.Status, body)
                    }
                }
            }
        }
        case 1: {
            //
        }
    }
}
