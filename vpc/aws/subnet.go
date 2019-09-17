package aws

import (
	"github.com/pulumi/pulumi-aws/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Subnet struct {
	*ec2.Subnet
}

func NewSubnet(ctx *pulumi.Context, vpcId *pulumi.IDOutput, name string, cidr string, az string, tags map[string]string) *Subnet {
	subnet, err := ec2.NewSubnet(ctx, name, &ec2.SubnetArgs{
		AvailabilityZone: az,
		CidrBlock:        cidr,
		Tags:             tags,
		VpcId:            vpcId,
	})
	if err != nil {
		panic(err)
	}
	return &Subnet{
		Subnet: subnet,
	}
}

// func (s *Subnet) ID() *pulumi.IDOutput  {
// 	return s.Subnet.ID()
// }
