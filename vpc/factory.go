package vpc

import "github.com/pulumi/pulumi/sdk/go/pulumi"

type Factory interface {
	BuildVpc(ctx *pulumi.Context, name string, cidr string, tags map[string]string) Vpc
	BuildSubnet(ctx *pulumi.Context, vpcId *pulumi.IDOutput, name string, cidr string, az string, tags map[string]string) Subnet
}
