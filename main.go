package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/caarlos0/env"
)

// config ... struct for holding environment variables
type config struct {
	Regions []string `env:"regions" envSeparator:"," envDefault:""`
}

// fetchRegion ... Gets all AWS regions using DescribeRegions
func fetchRegion() ([]string, error) {
	awsSession := session.Must(session.NewSession(&aws.Config{}))

	svc := ec2.New(awsSession)
	awsRegions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}

	regions := make([]string, 0, len(awsRegions.Regions))
	for _, region := range awsRegions.Regions {
		regions = append(regions, *region.RegionName)
	}

	return regions, nil
}

func main() {
	l := log.New(os.Stderr, "", 0)
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		l.Println(err)
		os.Exit(1)
	}
	regions := cfg.Regions
	if len(regions) == 0 {
		var err error
		regions, err = fetchRegion()
		if err != nil {
			l.Println(err)
			os.Exit(1)
		}
	}
	for _, region := range regions {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
		cleanupEc2Instances(sess)
	}
}
