package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const region = "us-west-2"
const instanceName = "GRACE cleanup test"

func terminatedInstances(svc *ec2.EC2) (int, error) {
	// Only grab instances that are terminated
	i := 0
	filters := []*ec2.Filter{
		&ec2.Filter{
			Name:   aws.String("instance-state-name"),
			Values: []*string{aws.String("terminated")},
		},
	}
	request := &ec2.DescribeInstancesInput{Filters: filters}
	result, err := svc.DescribeInstances(request)
	if err != nil {
		return i, err
	}
	for _, r := range result.Reservations {
		i = i + len(r.Instances)
	}
	return i, nil
}

func TestFetchInstances(t *testing.T) {
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
	if len(ids) != 1+terminated {
		t.Fatalf("expected %v instance(s).  Got %v", 1+terminated, len(ids))
	}
	err = checkInstances(t, ids, svc)
	if err != nil {
		t.Fatalf("Error checking instances: %v", err)
	}
}

func checkInstances(t *testing.T, ids []*string, svc *ec2.EC2) error {
	resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: ids,
	})
	if err != nil {
		t.Fatalf("Error describing instances: %v", err)
		return err
	}
	for _, r := range resp.Reservations {
		for _, i := range r.Instances {
			if *i.State.Name != "terminated" {
				if len(i.Tags) != 1 {
					t.Fatalf("Expected exactly one tag.  Got: %v", len(i.Tags))
				}
				if *i.Tags[0].Key != "Name" {
					t.Fatalf("Expected first tag to be 'Name'.  Got '%v'", *i.Tags[0].Key)
				}
				if *i.Tags[0].Value != instanceName {
					t.Fatalf("Expected first tag to be '%v'.  Got '%v'", instanceName, *i.Tags[0].Value)
				}
			}
		}
	}
	return nil
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
