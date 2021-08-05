# bomb (一款电话短信轰炸器)

## 关于

看网上大部分轰炸器都是Python写的，并且还很多还都删库了，剩下的看着都不太好用，于是想用go写一款简单的轰炸器。  


## 编译

采用go1.16开发自测。 直接安装go后build就行  

Debian:
```bash
$ sudo apt install golang -y
$ git clone https://github.com/Ro0tk1t/bomb.git
$ cd bomb
$ go build
```


## 使用

```bash
$ ./bomb --help
Usage:
  bomb [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  version     show current version

Flags:
  -a, --areacode string   the phone number area code, example: +86
  -c, --config string     data file
  -d, --delay int         delay bomb between requests
  -h, --help              help for bomb
  -p, --phone string      phone number
```

`-d` 每个轰炸请求间隔  
`-p` 手机号  
`-a` 区号，如中国是 +86  
`-c` 指定配置文件  

example: ./bomb -d 3 -c data.yml -a +86 -p 1888888888


## Tips 

* 接口使用`yaml`数据格式
* 没太多时间去收集可用的接口，所以只有一个石墨的接口可用。 大家感兴趣的话可以提交合并请求。  
* 代码写的比较随意，有问题的地方欢迎提issue  


## 接口格式
```yaml
name: 接口名
data: 需提交的数据
datatype: 数据类型
gap: 轰炸间隔
headers: http头
method: http方法
needareacode: 区号
url: url
```

> 数据类型：0代表字符串，1代表json，2代表xml  
> http方法：0代表get，1代表post，2代表put  
