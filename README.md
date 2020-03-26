# Ipgetter

This module is designed to fetch your external IP address from the internet.

## Usage

```go
package main

import (
	"fmt"
	"github.com/kejunmao/ipgetter"
)

func main()  {
	fmt.Println(ipgetter.Myip())
}
```

## Command Usage

```shell script
$ go install github.com/kejunmao/ipgetter/...
$ ipgetter
8.8.8.8
```
