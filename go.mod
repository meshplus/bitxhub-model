module github.com/meshplus/bitxhub-model

go 1.13

require (
	github.com/cbergoon/merkletree v0.2.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/meshplus/bitxhub-kit v1.2.1-0.20210723100713-b8d99c166281
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.33.1
)

replace github.com/ultramesh/crypto-gm => git.hyperchain.cn/dmlab/crypto-gm v0.2.14
