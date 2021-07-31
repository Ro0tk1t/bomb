package utils

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"

    "gopkg.in/yaml.v3"

    "bomb/cmd"
    "bomb/conf"
    "bomb/models"
)

func LoadDatas(){
    if fd, err := os.Open("data.yml"); err == nil {
        data, _ := ioutil.ReadAll(fd)
        err = yaml.Unmarshal(data, &conf.Datas)
        if err != nil {
            fmt.Println(err)
        }
    } else {
        fmt.Println(err)
    }
    newDatas := []models.API{}
    for _, api := range conf.Datas {
        if !strings.Contains(api.Data, "<phone>") && !strings.Contains(api.Url, "<phone>"){
            fmt.Println("[-] data not contains <phone>: ", api.Name)
        } else {
            newDatas = append(newDatas, api)
        }
    }
    conf.Datas = newDatas

    fmt.Println(conf.Datas[0].GetMethod())
    fmt.Println(conf.Datas[0].GetDataType())
}


func ReshapeData() {
    for _, api := range conf.Datas {
        if api.NeedAreaCode {
            if cmd.AreaCode == "" {
                cmd.AreaCode = "+86"
            }
        }
        api.Data = strings.ReplaceAll(api.Data, "<phone>", cmd.Phone)
        api.Url = strings.ReplaceAll(api.Url, "<phone>", cmd.Phone)
    }
}
