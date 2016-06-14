package klient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"time"
	"os"
)

type Kclient struct {
	kinesis      *kinesis.Kinesis
	partitionKey string
	stream       string
}

func New(region, partitionKey, stream string) *Kclient {
	svc := kinesis.New(session.New(), aws.NewConfig().WithRegion(region))
	return &Kclient{kinesis: svc, partitionKey: partitionKey, stream: stream }
}

func (client *Kclient) getShards() ([]string, error) {
	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String(client.stream),
	}

	resp, err := client.kinesis.DescribeStream(params)
	if err != nil {
		return nil, err
	}

	shards := make([]string, 0, len(resp.StreamDescription.Shards))

	for _, s := range resp.StreamDescription.Shards {
		shards = append(shards, *s.ShardId)
	}
	return shards, nil
}

func (client *Kclient) getInitialShardIterator(shardId string) (*string, error) {
	params := &kinesis.GetShardIteratorInput{
		ShardId:           aws.String(shardId),
		ShardIteratorType: aws.String("LATEST"),
		StreamName:        aws.String(client.stream),
		Timestamp:         aws.Time(time.Now()),
	}

	resp, err := client.kinesis.GetShardIterator(params)

	if err != nil {
		return nil, err
	}

	return resp.ShardIterator, nil
}

func (client *Kclient) readStream(messages chan string, shardId string) {
	nextShardIterator, err := client.getInitialShardIterator(shardId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	for {
		getParams := &kinesis.GetRecordsInput{
			ShardIterator: nextShardIterator,
		}

		resp, err := client.kinesis.GetRecords(getParams)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		for _, e := range resp.Records {
			messages <- string(e.Data)
		}

		nextShardIterator = resp.NextShardIterator
		if len(resp.Records) == 0 {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
func (client *Kclient) Put(payload string) {
	params := &kinesis.PutRecordInput{
		Data:         []byte(payload),
		PartitionKey: aws.String(client.partitionKey),
		StreamName:   aws.String(client.stream),
	}
	resp, err := client.kinesis.PutRecord(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func (client *Kclient) Subscribe() {
	shardIds, err := client.getShards()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	messages := make(chan string, 10)

	for _, shardId := range shardIds {
		go client.readStream(messages, shardId)
	}

	for {
		fmt.Println(<-messages)
	}

}