package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	klient "github.com/vevo/kinesis-client"
)

var (
	dataFile string
	data string
	stream string
	region string
	partitionKey string
)

func init() {
	flag.StringVar(&dataFile, "datafile", "", "file containing data to send")
	flag.StringVar(&data, "data", "", "data to send")
	flag.StringVar(&partitionKey, "partition", "RANDOM", "partition key")
	flag.StringVar(&stream, "stream", "", "stream name")
	flag.StringVar(&region, "region", "us-east-1", "aws region")

}
func readJsonFile() string {
	file, e := ioutil.ReadFile(dataFile)
	if e != nil {
		fmt.Printf("File read error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	return string(file)

}

func main() {
	flag.Parse()
	if stream == "" {
		fmt.Println("Required: -stream <name>, name cannot be empty")
		os.Exit(1)
	}
	svc := klient.New(region, partitionKey, stream)

	if dataFile != "" {
		data = readJsonFile()
		svc.Put(data)
	} else if data != "" {
		svc.Put(data)
	} else {
		fmt.Println("Subscribed to stream:", stream, " - Now publish some data to this stream for it to appear -")
		svc.Subscribe()
	}
}
