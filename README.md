# kinesis-client

# New to GoLang ?
    $ brew install go
    $ mkdir -p ~/gocode
    $ export GOPATH=~/gocode
    $ export PATH="${PATH}:${GOPATH}/bin"
    $ go get https://github.com/xoraes/kinesis-client
    $ go install $GOPATH/src/github.com/xoraes/kinesis-client/...
    
    
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
