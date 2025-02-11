package service

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/kihyun1998/go-aws/config"
)

type EC2Instance struct {
	InstanceID   string
	InstanceType string
	State        string
	Tags         map[string]string
	PublicIP     string
	PrivateIP    string
}

type EC2Service interface {
	ListInstances(ctx context.Context) ([]EC2Instance, error)
}

type AWSEC2Service struct {
	client *ec2.Client
}

func NewAWSEC2Service(cfg *config.AWSConfig) (EC2Service, error) {
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID,
			cfg.SecretAccessKey,
			"",
		)),
	)

	if err != nil {
		return nil, err
	}

	return &AWSEC2Service{
		client: ec2.NewFromConfig(awsCfg),
	}, nil
}

func (s *AWSEC2Service) ListInstances(ctx context.Context) ([]EC2Instance, error) {
	input := &ec2.DescribeInstancesInput{}
	result, err := s.client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, err
	}

	var instances []EC2Instance
	for _, reservation := range result.Reservations {
		for _, instance := range *&reservation.Instances {
			tags := make(map[string]string)
			for _, tag := range instance.Tags {
				tags[*tag.Key] = *tag.Value
			}

			inst := EC2Instance{
				InstanceID:   *instance.InstanceId,
				InstanceType: string(*&instance.InstanceType),
				State:        string(instance.State.Name),
				Tags:         tags,
			}

			if instance.PublicIpAddress != nil {
				inst.PublicIP = *instance.PublicIpAddress
			}
			if instance.PrivateIpAddress != nil {
				inst.PrivateIP = *instance.PrivateIpAddress
			}

			instances = append(instances, inst)
		}
	}

	return instances, nil

}
