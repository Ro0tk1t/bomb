package core

import (
    "fmt"
//    "net/http"

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

func (s Send) Run () {
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
    };
    }
}

