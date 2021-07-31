package core

import (
    "bomb/conf"
    "bomb/models"
)

func StartBomb(){
    var bombs []models.Timer
    for _, api := range conf.Datas {
        timer := models.Timer{
            Gap: api.Gap,
        }
        bombs = append(bombs, timer)
    }
}


func bomb() {}


func send() {}
