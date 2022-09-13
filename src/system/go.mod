module system

go 1.18

require go-micro.dev/v4 v4.8.0

require github.com/google/uuid v1.2.0 // indirect

// Uncomment if you use etcd
// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
// replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace system => ./
