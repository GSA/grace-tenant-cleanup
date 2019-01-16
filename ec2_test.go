package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const region = "us-west-2"
const instanceName = "GRACE cleanup test"

func terminatedInstances(svc *ec2.EC2) ([]*ec2.Instance, error) {
	// Only grab instances that are terminated
	filters := []*ec2.Filter{
		&ec2.Filter{
			Name:   aws.String("instance-state-name"),
			Values: []*string{aws.String("terminated")},
		},
	}
	request := &ec2.DescribeInstancesInput{Filters: filters}
	result, err := svc.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	if len(result.Reservations) > 0 {
		return result.Reservations[0].Instances, nil
	}
	return nil, nil
}

func TestFetchInstances(t *testing.T) {
	fmt.Printf("AWS Profile: %s\n", os.Getenv("AWS_PROFILE"))
	if !(*integration || *destructive) {
		t.Skip("skipping test in non-integration mode")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc := ec2.New(sess)
	ids, err := fetchInstances(svc)
	if err != nil {
		t.Fatalf("failed to fetch instances: %v", err)
	}
	terminated, err := terminatedInstances(svc)
	if err != nil {
		t.Fatalf("Failed to fetch terminated instances: %v", err)
	}
	if len(ids) != 1+len(terminated) {
		t.Fatalf("expected %v instance(s).  Got %v", 1+len(terminated), len(ids))
	}
	resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: ids,
	})
	if err != nil {
		t.Fatalf("Error describing instances: %v", err)
	}
	if *resp.Reservations[0].Instances[0].InstanceId != *ids[0] {
		t.Fatalf("Expected instance ID to equal %v.  Got %v", *resp.Reservations[0].Instances[0].InstanceId, *ids[0])
	}
	if *resp.Reservations[0].Instances[0].Tags[0].Key != "Name" {
		t.Fatalf("Expected first tag to be 'Name'.  Got '%v'", *resp.Reservations[0].Instances[0].Tags[0].Key)
	}
	if *resp.Reservations[0].Instances[0].Tags[0].Value != instanceName {
		t.Fatalf("Expected first tag to be '%v'.  Got '%v'", instanceName, *resp.Reservations[0].Instances[0].Tags[0].Value)
	}
}

func TestCleanupEc2Instances(t *testing.T) {
	if !(*destructive) {
		t.Skip("skipping test in non-desctructive mode")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	cleanupEc2Instances(sess)
}
