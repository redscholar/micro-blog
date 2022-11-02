module article

go 1.19

require go-micro.dev/v4 v4.8.0

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210920160938-87db9fbc61c7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xanzy/ssh-agent v0.3.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211020060615-d418f374d309 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20211022215931-8e5104632af7 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// Uncomment if you use etcd
// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
// replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace system => ./
