# Tencent CLS Hook for logrus
[![GoDoc](https://godoc.org/github.com/chuangbo/logruscls?status.svg)](https://godoc.org/github.com/chuangbo/logruscls)
[![Go Report Card](https://goreportcard.com/badge/github.com/chuangbo/logruscls)](https://goreportcard.com/report/github.com/chuangbo/logruscls)

Asynchronise batch upload logs to [Tencent Cloud CLS](https://cloud.tencent.com/document/product/614) for logrus with configurable batch number and max wait time.

## Usage

```go
package main

import (
	"github.com/chuangbo/logruscls"
	log "github.com/sirupsen/logrus"
)

func main() {
	// NewCLSAsyncClient or NewCLSClient
	client, err := logruscls.NewCLSAsyncClient(
		"topicID",
		"region",
		"secretID",
		"secretKey",
		30, // max number of logs for batch upload
		time.Second*30, // max wait time before reach max number of logs
	)
	if err != nil {
		panic(err)
	}

	hook, err := logruscls.NewHook(client)
	if err != nil {
		panic(err)
	}
	log.AddHook(hook)
}
```

Of course the logrus-free golang cls client `CLSClient` and `CLSAsyncClient` are for you to use directly as well if logrus is not your thing.

```go
package main

import (
	"github.com/chuangbo/logruscls"
	"github.com/chuangbo/logruscls/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	client, err := logruscls.NewCLSAsyncClient(...)
	if err != nil {
		panic(err)
	}

	err = client.Log(&pb.Log{
		Time: proto.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
		Contents: []*pb.Log_Content{
			{
				Key:   proto.String("message"),
				Value: proto.String("Hello Logrus"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
```

For more API please refer to [godoc](https://godoc.org/github.com/chuangbo/logruscls) and [CLS protobuf definition](https://cloud.tencent.com/document/product/614/16873). But please keep in mind, this may never going to be a full-feature CLS client.

## Todos

* ~~signature for headers and query string~~ (md5 doesn't work so kind of pointless)
* ~~content md5~~ (seems upload using even wrong random md5 checksum work)
* ~~lz4 compress~~ (doesn't work)
* tests

## License

[MIT](http://opensource.org/licenses/MIT)

Copyright (c) 2019-present, Chuangbo Li
