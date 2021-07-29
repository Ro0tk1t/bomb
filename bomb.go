package main

import (
    "fmt"
    //"encoding/json"

    "bomb/cmd"
    "bomb/utils"
    "bomb/models"
)


func main(){
    cmd.Execute()
    var test models.API
    test.DataType = 0

    go utils.LoadDatas()
    fmt.Println(cmd.CF)
}

