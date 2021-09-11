# go-jsdelivr

try cover url to jsDelivr CDN url

## 代码引用

- 引入包
  `go get github.com/Luoxin/go-jsdelivr`
- 示例代码

```go
package main

import (
	go_jsdelivr "github.com/Luoxin/go-jsdelivr"
)

func main() {
	nu, err := go_jsdelivr.Cover("https://raw.githubusercontent.com/Luoxin/go-jsdelivr/master/README.md")
	if err != nil {
		fmt.Println(err)
	} else {
		// https://cdn.jsdelivr.net/gh/Luoxin/go-jsdelivr@master/README.md
		fmt.Println(nu)
	}
}
```

## 直接使用

### 源码使用

- `gh repo clone Luoxin/go-jsdelivr`
- `cd go-jsdelivr`
- `go mod downlaod`
- `go run ./cmd/jsdelivr.go <url>`

### 使用编译好的包

- 从页面 [https://github.com/Luoxin/go-jsdelivr/releases/latest](https://github.com/Luoxin/go-jsdelivr/releases/latest)
  获取最新版的二进制包下载并解压
- `jsdelivr <url>`
