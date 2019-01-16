package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func fetchInstances(svc *ec2.EC2) ([]*string, error) {
	input := &ec2.DescribeInstancesInput{}
	resp, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	var ids []*string
	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			ids = append(ids, i.InstanceId)
		}
	}
	return ids, nil
}

func cleanupEc2Instances(sess client.ConfigProvider) {
	svc := ec2.New(sess)
	ids, err := fetchInstances(svc)
	if err != nil {
		log.Println(err)
		return
	}
	if len(ids) > 0 {
		termResp, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
			InstanceIds: ids,
		})
		if err != nil {
			log.Print("Error terminating instances. ", err)
		}
		fmt.Println("Resp: ", termResp)
	}
}
