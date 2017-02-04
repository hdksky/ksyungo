# KsyunGo: Go SDK for ksyun Services

This is an unofficial Go SDK for Ksyun Services


## Package Structure

*  kec: [云服务器](https://docs.ksyun.com/read/latest/52/_book/index.html)
*  slb: [负载均衡](https://docs.ksyun.com/read/latest/55/_book/index.html)
*  eip: [弹性IP](https://docs.ksyun.com/read/latest/57/_book/index.html)
*  vpc: [虚拟私有网络](https://docs.ksyun.com/read/latest/56/_book/index.html)
*  common: Common libary of ksyun Go SDK (Borrowed from AliyunGo)
*  util: Utility helpers



## Quick Start

```go
package main

import (
	"fmt"

	"github.com/hdksky/ksyungo/kec"
)

const ACCESS_KEY_ID = "<YOUR_ID>"
const ACCESS_KEY_SECRET = "<****>"

func main() {
	client := kec.NewClient(ACCESS_KEY_ID, ACCESS_KEY_SECRET,"cn-beijing-6")
	fmt.Print(client.DescribeInstances(&DescribeInstancesArgs{}))
}

```


## Build and Install

go get:

```sh
go get github.com/hdksky/ksyungo
```

## License
This project is licensed under the Apache License, Version 2.0.
