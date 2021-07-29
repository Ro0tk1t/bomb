package utils

import (
    "os"
    "fmt"
    "io/ioutil"

    "gopkg.in/yaml.v3"

    "bomb/conf"
    //"bomb/models"
)

func LoadDatas(){
    fmt.Println("Hello, World!")
    if fd, err := os.Open("data.yml"); err == nil {
        data, _ := ioutil.ReadAll(fd)
        yaml.Unmarshal(data, &conf.Datas)
    }

    fmt.Println(conf.Datas)
}


func Reshape() {
    // TODO
}
