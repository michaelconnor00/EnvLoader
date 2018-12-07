package main

import (
	credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	ssm "github.com/aws/aws-sdk-go/service/ssm"
)

ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
SECRET_KEY = os.Getenv("AWS_ACCESS_SECRET")
REGION_NAME = os.Getenv("AWS_REGION")

func ssm_init() *ssm.SSM {
	cred := credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, "")
	cfg := aws.NewConfig().WithRegion(REGION_NAME_DASH).WithCredentials(cred)
	return ssm.New(session.New(cfg))
}

func fetch_params(ssm_client *ssm.SSM) {
	
}
