/*
Package logruscls provides logrus hook for tencent cls.

It has these top-level structs:
	Hook
	CLSClient
	CLSAsyncClient
*/
package logruscls

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chuangbo/logruscls/pb"

	"github.com/golang/protobuf/proto"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// CLSClient is cls client
type CLSClient struct {
	URL                 string
	SecretID, SecretKey string
	TopicID             string
}

// NewCLSClient create a cls client
func NewCLSClient(region, secretID, secretKey, topicID string) (*CLSClient, error) {
	if region == "" || secretID == "" || secretKey == "" || topicID == "" {
		return nil, errors.New("please specific cls configurations")
	}

	return &CLSClient{
		URL:       fmt.Sprintf("http://%s.cls.myqcloud.com/", region),
		SecretID:  secretID,
		SecretKey: secretKey,
		TopicID:   topicID,
	}, nil
}

// Log upload one log directly to cls
func (c *CLSClient) Log(log *pb.Log) error {
	return c.UploadStructuredLog(
		&pb.LogGroupList{
			LogGroupList: []*pb.LogGroup{
				{
					Logs: []*pb.Log{log},
				},
			},
		},
	)
}

// UploadStructuredLog upload structured log to tencent CLS
func (c *CLSClient) UploadStructuredLog(logGroupList *pb.LogGroupList) error {
	requestBody, err := proto.Marshal(logGroupList)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", fmt.Sprintf("%sstructuredlog?topic_id=%s", c.URL, c.TopicID), bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/x-protobuf")
	r.Header.Set("Authorization", sign("post", "/structuredlog", c.SecretID, c.SecretKey))

	resp, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func hmacSHA1(src, key string) string {
	var mac hash.Hash
	mac = hmac.New(sha1.New, []byte(key))

	mac.Write([]byte(src))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// sign 对 cls 请求进行签名，我们暂时不对 header 和 query 进行签名
// https://cloud.tencent.com/document/product/614/12445
func sign(method, uri, secretID, secretKey string) string {
	httpRequestInfo := fmt.Sprintf("%s\n%s\n\n\n", method, uri)
	httpRequestInfoSHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(httpRequestInfo)))

	startTime := time.Now().Unix()
	endTime := startTime + 300

	stringToSign := fmt.Sprintf("sha1\n%d;%d\n%s\n", startTime, endTime, httpRequestInfoSHA1)
	signKey := hmacSHA1(fmt.Sprintf("%d;%d", startTime, endTime), secretKey)
	signature := hmacSHA1(stringToSign, signKey)

	return fmt.Sprintf(
		"q-sign-algorithm=sha1&q-ak=%s&q-sign-time=%d;%d&q-key-time=%d;%d&q-header-list=&q-url-param-list=&q-signature=%s",
		secretID,
		startTime,
		endTime,
		startTime,
		endTime,
		signature,
	)
}
