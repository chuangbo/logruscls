# Tencent CLS Hook for logrus
[![GoDoc](https://godoc.org/github.com/chuangbo/logruscls?status.svg)](https://godoc.org/github.com/chuangbo/logruscls)
[![Go Report Card](https://goreportcard.com/badge/github.com/chuangbo/logruscls)](https://goreportcard.com/report/github.com/chuangbo/logruscls)

Asynchronise batch upload logs to [Tencent Cloud CLS](https://cloud.tencent.com/document/product/614/16873) for logrus with configurable batch number and max wait time.

## Usage

```go
package main

import (
	"github.com/chuangbo/logruscls"
	log "github.com/sirupsen/logrus"
)

func main() {
	hook, err := logruscls.NewHook(
		"topicID",
		"region",
		"secretID",
		"secretKey",
		// max number of logs for batch upload
		30,
		// max wait time before reach max number of logs
		time.Second*30,
	)
	if err != nil {
		panic(err)
	}
	log.AddHook(hook)
}

```

Of course there is also a logrus-free golang cls client `CLSClient` for you to use directly. Please refer to [godoc](https://godoc.org/github.com/chuangbo/logruscls). But please keep in mind, this may never going to be a full-feature CLS client.

## Todos

* signature for headers and query string
* content md5
* tests

## License

[MIT](http://opensource.org/licenses/MIT)

Copyright (c) 2019-present, Chuangbo Li
