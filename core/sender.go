package core

import (
    "fmt"
    "bytes"
    "strings"
    "net/http"
    "io/ioutil"
   // "crypto/tls"

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

func (s *Send) SendData (api models.API) {
    fmt.Println(api.Url)
    // TODO
}

func (s Send) Recv () {
    //TODO
}

func (s *Send) Run (api models.API) {
    s.SendData(api)
}


func NewSender(way string) models.Sender {
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
    fmt.Println("************************")
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
                var (
                    req *http.Request
                    err error
                )
                if api.Data != "" {
                    req, err = http.NewRequest("GET", api.Url, strings.NewReader(api.Data))
                } else {
                    req, err = http.NewRequest("GET", api.Url, nil)
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
            body := []byte(api.Data)
            req, err := http.NewRequest("POST", api.Url, bytes.NewBuffer(body))
            client := &http.Client{}
            if err != nil {
                fmt.Printf("%v", err)
            } else {
                if len(api.Headers) > 0 {
                    for k, v := range api.Headers {
                        req.Header.Add(k, v)
                    }
                }
                resp, err := client.Do(req)
                if err == nil {
                    defer resp.Body.Close()
                    body, _ := ioutil.ReadAll(resp.Body)
                    fmt.Println(resp.Status, body)
                }
            }
        }
        case 2: {
            // TODO
        }
    }
}

func (h *HttpsSender) SendData (api models.API) {
    switch api.Method {
        case 0: {
            tr := &http.Transport{}
            client := &http.Client{Transport: tr}
            req, err := http.NewRequest("GET", api.Url, nil)
            if err != nil {
                fmt.Printf("%v", err)
            } else {
                if len(api.Headers) > 0 {
                    for k, v := range api.Headers {
                        req.Header.Add(k, v)
                    }
                }
                resp, err := client.Do(req)
                if err == nil {
                    defer resp.Body.Close()
                    body, _ := ioutil.ReadAll(resp.Body)
                    fmt.Println(resp.Status, body)
                }
            }
        }
        case 1: {
            tr := &http.Transport{}
            client := &http.Client{Transport: tr}
            req, err := http.NewRequest("POST", api.Url, nil)
            if err != nil {
                fmt.Printf("%v", err)
            } else {
                if len(api.Headers) > 0 {
                    for k, v := range api.Headers {
                        req.Header.Add(k, v)
                    }
                }
                resp, err := client.Do(req)
                if err == nil {
                    defer resp.Body.Close()
                    body, _ := ioutil.ReadAll(resp.Body)
                    fmt.Println(resp.Status, body)
                }
            }
        }
        case 2: {
            //
        }
    }
}
