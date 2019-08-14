package logruscls

import (
	"log"
	"time"

	"github.com/chuangbo/logruscls/pb"
)

// CLSAsyncClient to send logs to cls in batch on the background.
type CLSAsyncClient struct {
	// max number of one batch upload
	batch int
	// max delay
	delay time.Duration

	cls  *CLSClient
	logs chan *pb.Log
}

// NewCLSAsyncClient creates a async version of CLSClient
func NewCLSAsyncClient(region, secretID, secretKey, topicID string, batch int, delay time.Duration) (*CLSAsyncClient, error) {
	client, err := NewCLSClient(region, secretID, secretKey, topicID)
	if err != nil {
		return nil, err
	}

	asyncClient := &CLSAsyncClient{
		batch: batch,
		delay: delay,

		// double buffer
		logs: make(chan *pb.Log, batch*2),
		cls:  client,
	}

	// upload in batch in a goroutine
	go asyncClient.startSender()

	return asyncClient, nil
}

// Log put one log to the logs queue
func (c *CLSAsyncClient) Log(log *pb.Log) error {
	c.logs <- log
	return nil
}

// startSender
func (c *CLSAsyncClient) startSender() {
	for {
		logs := []*pb.Log{}
		t := time.NewTimer(c.delay)

	receiveLoop:
		// receive either maximum number of `batch` logs or wait for max `delay` time
		for {
			select {
			case <-t.C:
				if len(logs) > 0 {
					break receiveLoop
				}
				t.Reset(c.delay)
			case l := <-c.logs:
				logs = append(logs, l)
				if len(logs) >= c.batch {
					break receiveLoop
				}
			}
		}

		logGroupList := &pb.LogGroupList{
			LogGroupList: []*pb.LogGroup{
				{
					Logs: logs,
				},
			},
		}

		err := c.cls.UploadStructuredLog(logGroupList)
		if err != nil {
			log.Printf("could not upload to cls: %v", err)
		}
	}
}
