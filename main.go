package main

import (
	"github.com/lonegunmanb/pulumi-aws-factory/vpc"
	"github.com/lonegunmanb/pulumi-aws-factory/vpc/aws"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var factory vpc.Factory = aws.Factory{}

		v := factory.BuildVpc(ctx, "testvpc", "10.0.0.0/16", map[string]string{
			"app": "test",
		})
		subnet := factory.BuildSubnet(ctx, v.ID(), "test_subnet", "10.0.0.0/24", "ap-northeast-1a", map[string]string{
			"app": "test",
		})

		// Export the name of the bucket
		ctx.Export("vpc_id", v.ID())
		ctx.Export("subnet_id", subnet.ID())
		return nil
	})
}
