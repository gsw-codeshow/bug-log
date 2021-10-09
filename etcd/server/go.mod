module bug-log/etcd/server

go 1.14

require (
	bug-log/etcd/proto v0.0.0-00010101000000-000000000000
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d
	google.golang.org/grpc v1.35.0
)

replace bug-log/etcd/proto => ../proto

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
