# hex2rgb [![Build Status](https://travis-ci.org/dlion/hex2rgb.svg?branch=master)](https://travis-ci.org/dlion/hex2rgb)

Convert HEX color to RGB

## How to use it

```go
package main

import (
	hex "github.com/dlion/hex2rgb"
	"fmt"
)

func main() {
	fmt.Println(hex.Convert("#A972B6"))
}
```

Result:

```
169 114 182
```

## Test

```
go test
```


## Author 

Domenico Luciani

https://domenicoluciani.com

## LICENSE

MIT

