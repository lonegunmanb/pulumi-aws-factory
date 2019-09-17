package vpc

import "github.com/pulumi/pulumi/sdk/go/pulumi"

type Vpc interface {
	ID() *pulumi.IDOutput
}

type Subnet interface {
	ID() *pulumi.IDOutput
}
