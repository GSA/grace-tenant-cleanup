package main

import (
	"flag"
	"testing"
)

var expectedRegions = [...]string{
	"ap-south-1",
	"eu-west-3",
	"eu-north-1",
	"eu-west-2",
	"eu-west-1",
	"ap-northeast-2",
	"ap-northeast-1",
	"sa-east-1",
	"ca-central-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"eu-central-1",
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

var (
	integration = flag.Bool("integration", false, "run non-desctructive AWS integration tests")
	destructive = flag.Bool("destructive", false, "run destructive AWS integration tests")
)

func TestFetchRegion(t *testing.T) {
	if !(*integration || *destructive) {
		t.Skip("skipping test in non-integration mode")
	}
	regions, err := fetchRegion()
	if err != nil {
		t.Fatalf("failed to fetch regions: %v", err)
	}
	if len(regions) != len(expectedRegions) {
		t.Errorf("invalid regions, expected %v, got: %v", expectedRegions, regions)
	}
}
