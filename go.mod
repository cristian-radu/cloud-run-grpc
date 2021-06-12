module github.com/cristian-radu/cloud-run-grpc-client

replace github.com/cristian-radu/cloud-run-grpc-client/pkg/pb/run => /home/cristian/work/cloud-run-grpc-client/pkg/pb/run

replace github.com/cristian-radu/cloud-run-grpc-client/pkg/client => /home/cristian/work/cloud-run-grpc-client/pkg/client

replace github.com/cristian-radu/cloud-run-grpc-client/pkg/test => /home/cristian/work/cloud-run-grpc-client/pkg/test

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/googleapis/gax-go/v2 v2.0.5
	google.golang.org/api v0.48.0
	google.golang.org/genproto v0.0.0-20210611144927-798beca9d670
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
)
