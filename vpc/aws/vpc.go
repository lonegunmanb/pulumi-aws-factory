package aws

import (
	"github.com/pulumi/pulumi-aws/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AwsVpc struct {
	*ec2.Vpc
}

func NewAwsVpc(ctx *pulumi.Context, name string, cidr string, tags map[string]string) *AwsVpc {
	vpc, err := ec2.NewVpc(ctx, name, &ec2.VpcArgs{
		CidrBlock: cidr,
		Tags:      tags,
	})
	if err != nil {
		panic(err)
	}
	return &AwsVpc{
		Vpc: vpc,
	}
}

// func (vpc *AwsVpc) ID() *pulumi.IDOutput {
// 	return vpc.Vpc.ID()
// }
