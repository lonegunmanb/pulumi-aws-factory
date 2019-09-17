package aws

import (
	"github.com/lonegunmanb/pulumi-aws-factory/vpc"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Factory struct {
}

func (f Factory) BuildVpc(ctx *pulumi.Context, name string, cidr string, tags map[string]string) vpc.Vpc {
	return NewAwsVpc(ctx, name, cidr, tags)
}

func (f Factory) BuildSubnet(ctx *pulumi.Context, vpcId *pulumi.IDOutput, name string, cidr string, az string, tags map[string]string) vpc.Subnet {
	return NewSubnet(ctx, vpcId, name, cidr, az, tags)
}
