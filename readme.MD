first:
```shell script
go get github.com/lonegunmanb/pulumi-aws-factory
cd $GOPATH/src/github.com/lonegunmanb/pulumi-aws-factory
go get ./...
go install github.com/lonegunmanb/pulumi-aws-factory
```
install and config pulumi and awscli as [get-started-aws](https://www.pulumi.com/docs/get-started/aws/)

install aws plugin first:
```shell script
pulumi plugin install resource aws v0.16.2
```

then create stack
```shell script
pulumi up
```

use destroy to cleanup
```shell script
pulumi destroy
```

we created abstract vpc and subnet interface in vpc package, and implementation struct in aws package. we assign aws factory to abstract factory interface variable, and it works.