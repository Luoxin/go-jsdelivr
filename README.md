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
	nu, err := go_jsdelivr.Cover("")
	if err != nil {
		fmt.Println(err)
	} else {
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
