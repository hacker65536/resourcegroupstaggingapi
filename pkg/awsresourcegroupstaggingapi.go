package awsresroucegroupstaggingapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/hacker65536/resourcegroupstaggingapi/cmd"
)

type AwsRsourceGroupsTaggingAPI struct {
	svc    *resourcegroupstaggingapi.Client
	filter *cmd.FilterConfig
}

func New(f *cmd.FilterConfig) *AwsRsourceGroupsTaggingAPI {

	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	s := resourcegroupstaggingapi.NewFromConfig(cfg)
	return &AwsRsourceGroupsTaggingAPI{
		svc:    s,
		filter: f,
	}
}
func (a *AwsRsourceGroupsTaggingAPI) getResources() *resourcegroupstaggingapi.GetResourcesOutput {

	svc := a.svc
	tags := []*types.TagFilter{}

	for _, v := range a.filter {

	}
	input := &resourcegroupstaggingapi.GetResourcesInput{
		ResourceTypeFilters: aws.StringSlice(a.filter.ResourceTypes),
		TagFilters: []*types.TagFilter{
			{
				Key:    aws.String("testtag"),
				Values: aws.StringSlice([]string{"testvalue"}),
			},
		},
	}

	resp, err := svc.GetResources(context.Background(), input)

	if err != nil {
		panic("failed to describe autoscaling group, " + err.Error())
	}

	return resp
}
