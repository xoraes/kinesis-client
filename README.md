# kinesis-client 
- This program uses aws kinesis stream api to publish and consume kinesis data. It can be used to test kinesis real time stream data moving in your kinesis pipelines.

### New to GoLang ?
    $ brew install go (for Ubuntu see : https://github.com/golang/go/wiki/Ubuntu)
    $ mkdir -p ${HOME}/gocode
    $ export GOPATH=${HOME}/gocode
    $ export PATH="${PATH}:${GOPATH}/bin"
    $ go get github.com/xoraes/kinesis-client
    $ cd ${GOPATH}/src/github.com/xoraes/kinesis-client; go install ./...
    
    
Ensure you have run "aws configure" on your machine such that your aws access and secret are in ~/.aws/credentials.

    $ kinesis-client --help
    Usage of kinesis-client:
      -data string (Optional: if -data and -datafile are not sent, the program will act as subscriber)
         data to send
      -datafile string (Optional: if -data and -datafile are not sent, the program will act as subscriber)
        file containing data to send
      -partition string
        partition key (default "RANDOM", used to send data to a particular shard based on hash of this partition key)
      -region string
        aws region (default "us-east-1")
      -stream string (Required)
        stream name
### Publisher Example:
    Send json data in hello.json to stream called "Hello-Stream" in us-east-1
    $ kinesis-client -datafile ~/hello.json -stream "Hello-Stream" 
    
### Subscriber Example:
    Listen and print real time data in stream called "Hello-Stream" in us-east-1
    $ kinesis-client -stream "Hello-Stream" 
