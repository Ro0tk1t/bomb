package models


type Type int
const (
    String Type = iota
    Json
    Xml
)

type method int
const (
    Get method = iota
    Post
    Put
)

var tmap = []string {"string", "json", "xml"}
var mmap = []string {"get", "post", "put"}


type API struct {
    Name string
    Gap int
    Url string
    Headers map[string]string
    DataType Type
    Data string
    NeedAreaCode bool
    AreaCode string
    Method method
}

type Sender interface {
    Init()
    SendData(API)
    //Recv()
    Run(API)
}


func (a *API) GetDataType ()(string) {
    return tmap[a.DataType]
}

func (a *API) GetMethod ()(string) {
    return mmap[a.Method]
}
