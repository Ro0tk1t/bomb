package core

import (
    "fmt"
    "net/http"

    "bomb/models"
)

func Init(s *models.Sender) {
    // TODO
}

func NewSender(way int) models.Sender {
    fmt.Println("Hello, World!")
    switch way {
    case "http":
        return &models.HttpSender{}
    case "https": {
        return &models.HttpsSender{}
    };
    case "tcp": {
        return &models.TcpSender{}
    };
    default: {
        return &models.HttpSender{}
    };
    }
}

