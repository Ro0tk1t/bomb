package models


type Type int
const (
    String Type = iota
    Json
    Xml
)

var tmap = []string {"string", "json", "xml"}

type API struct {
    Name string
    Gap int
    Url string
    Headers map[string]string
    DataType Type
    Data string
    NeedAreaCode bool
    AreaCode string
}


func (a *API) GetDataType ()(string) {
    return tmap[a.DataType]
}

