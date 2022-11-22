GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=./proto/user --go_out=:./proto/user proto/user/*.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o web *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t gateway:latest .

.PHONY: graphql-init
graphql-init:
	@echo "generate gqlgen.yml"
	@echo '# Where are all the schema files located? globs are supported eg  src/**/*.graphqls\nschema:\n  - graph/*.graphqls\n\n# Where should the generated server code go?\nexec:\n  filename: graph/generated/generated.go\n  package: generated\n\n# Uncomment to enable federation\n# federation:\n#   filename: graph/generated/federation.go\n#   package: generated\n\n# Where should any generated models go?\nmodel:\n  filename: graph/model/models_gen.go\n  package: model\n\n# Where should the resolver implementations go?\nresolver:\n  layout: follow-schema\n  dir: graph\n  package: graph\n\n# Optional: turn on use ` + "`" + `gqlgen:"fieldName"` + "`" + ` tags in your models\n# struct_tag: json\n\n# Optional: turn on to use []Thing instead of []*Thing\n# omit_slice_element_pointers: false\n\n# Optional: turn off to make struct-type struct fields not use pointers\n# e.g. type Thing struct { FieldA OtherThing } instead of { FieldA *OtherThing }\n# struct_fields_always_pointers: true\n\n# Optional: turn off to make resolvers return values instead of pointers for structs\n# resolvers_always_return_pointers: true\n\n# Optional: set to speed up generation time by not performing a final validation pass.\n# skip_validation: true\n\n# gqlgen will search for any type names in the schema in these go packages\n# if they match it will use them, otherwise it will generate them.\nautobind:\n#  - "web/graph/model"\n\n# This section declares type mapping between the GraphQL and go type systems\n#\n# The first line in each type will be used as defaults for resolver arguments and\n# modelgen, the others will be allowed when binding to fields. Configure them to\n# your liking\nmodels:\n  ID:\n    model:\n      - github.com/99designs/gqlgen/graphql.ID\n      - github.com/99designs/gqlgen/graphql.Int\n      - github.com/99designs/gqlgen/graphql.Int64\n      - github.com/99designs/gqlgen/graphql.Int32\n  Int:\n    model:\n      - github.com/99designs/gqlgen/graphql.Int\n      - github.com/99designs/gqlgen/graphql.Int64\n      - github.com/99designs/gqlgen/graphql.Int32' > gqlgen.yml

.PHONY: graphql-generate
graphql-generate:
	@go run github.com/99designs/gqlgen generate
