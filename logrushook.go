package logruscls

import (
	"fmt"
	"time"

	"github.com/chuangbo/logruscls/pb"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

// Hook to send logs via tencent cls.
type Hook struct {
	topicID string

	// max number of one batch upload
	batch int
	// max delay
	delay time.Duration

	cls  CLSLogger
	logs chan *pb.Log
}

// CLSLogger provides a cls client which has Log function for hook to use
type CLSLogger interface {
	Log(log *pb.Log) error
}

// NewHook creates a hook to be added to an instance of logger
func NewHook(client CLSLogger) (*Hook, error) {
	hook := &Hook{
		cls: client,
	}
	return hook, nil
}

// Fire uploads logs to cls
func (hook *Hook) Fire(entry *logrus.Entry) error {
	log := entryToLog(entry)
	return hook.cls.Log(log)
}

// Levels returns support levels for the hook
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
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
