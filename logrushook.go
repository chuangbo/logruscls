package logruscls

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/chuangbo/logruscls/pb"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

// Hook to send logs via syslog.
type Hook struct {
	topicID string

	// max number of one batch upload
	batch int
	// max delay
	delay time.Duration

	cls  *CLSClient
	logs chan *pb.Log
}

// NewHook creates a hook to be added to an instance of logger
func NewHook(topicID, region, secretID, secretKey string, batch int, delay time.Duration) (*Hook, error) {
	if topicID == "" || region == "" || secretID == "" || secretKey == "" {
		return nil, errors.New("please specific cls credentials")
	}

	hook := &Hook{
		topicID: topicID,

		batch: batch,
		delay: delay,

		// double buffer
		logs: make(chan *pb.Log, batch*2),
		cls:  NewCLSClient(region, secretID, secretKey),
	}

	// upload in batch in a goroutine
	go hook.startSender()

	return hook, nil
}

// Fire uploads logs to cls
func (hook *Hook) Fire(entry *logrus.Entry) error {
	log := entryToLog(entry)

	hook.logs <- log

	return nil
}

// Levels returns support levels for the hook
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// startSender
func (hook *Hook) startSender() {
	for {
		logs := []*pb.Log{}
		t := time.NewTimer(hook.delay)

	receiveLoop:
		// receive either maximum number of `batch` logs or wait for max `delay` time
		for {
			select {
			case <-t.C:
				if len(logs) > 0 {
					break receiveLoop
				}
				t.Reset(hook.delay)
			case l := <-hook.logs:
				logs = append(logs, l)
				if len(logs) >= hook.batch {
					break receiveLoop
				}
			}
		}

		logGroupList := &pb.LogGroupList{
			LogGroupList: []*pb.LogGroup{
				&pb.LogGroup{
					Logs: logs,
				},
			},
		}

		err := hook.cls.UploadStructuredLog(hook.topicID, logGroupList)
		if err != nil {
			log.Printf("could not upload to cls: %v", err)
		}
	}
}

func entryToLog(entry *logrus.Entry) *pb.Log {
	contents := []*pb.Log_Content{}

	// message
	contents = append(contents, &pb.Log_Content{
		Key:   proto.String("message"),
		Value: proto.String(entry.Message),
	})

	// level
	contents = append(contents, &pb.Log_Content{
		Key:   proto.String("level"),
		Value: proto.String(entry.Level.String()),
	})

	// fields
	for k, v := range entry.Data {
		contents = append(
			contents,
			&pb.Log_Content{
				Key:   proto.String(k),
				Value: proto.String(fmt.Sprintf("%v", v)),
			},
		)
	}

	return &pb.Log{
		Time:     proto.Int64(entry.Time.UnixNano() / int64(time.Millisecond)),
		Contents: contents,
	}
}
