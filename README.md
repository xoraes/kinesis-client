# kinesis-client

Ensure you have run aws configure on your machine such that your aws access and secret are in ~/.aws/credentials. 

kinesis-client --help
Usage of kinesis-client:
  -data string
    	data to send
  -datafile string
    	file containing data to send
  -partition string
    	partition key (default "RANDOM")
  -region string
    	aws region (default "us-east-1")
  -stream string
    	stream name
