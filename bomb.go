package main

import (
    //"fmt"
    //"encoding/json"

    "bomb/cmd"
    "bomb/core"
    "bomb/utils"
    "bomb/models"
)


func main(){
    cmd.Execute()
    var test models.API
    test.DataType = 0

    utils.LoadDatas()
    forever := make(chan struct{})
    utils.ReshapeData()
    go core.StartBomb()
    <- forever
}

