# kinesis-client - 
- Uses aws kinesis stream api to publish and consume kinesis data. It should be used to test real time stream data in your kinesis pipeline. Assumes that a kinesis stream exists in your aws env. If the stream does not exist in your env/region, a 400 - stream not found error is shown. 
- This tool uses stream api, not the KCL - thus no checkpointing is done and the stream will available for re consumption by your KCL application. 

### Installation steps if you are new to GoLang !
    $ brew update;brew install go [for Ubuntu see : https://github.com/golang/go/wiki/Ubuntu]
    $ mkdir -p ${HOME}/gocode
    $ export GOPATH="${HOME}/gocode";export PATH="${PATH}:${GOPATH}/bin" [might want to stick this to .bashrc or .profile]
    $ go install github.com/xoraes/kinesis-client/...
    
    
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
    Send data in hello.json to the stream called "Hello-Stream" in us-east-1 region
    $ kinesis-client -datafile ~/hello.json -stream "Hello-Stream" 
    
    Send data to the stream called "Hello-Stream" in us-east-1 region
    $ kinesis-client -data "UPDATE /event/123" -stream "Hello-Stream" 
    
### Subscriber Example:
    Listen and print real time data in stream called "Hello-Stream" in us-east-1 region
    $ kinesis-client -stream "Hello-Stream" 
