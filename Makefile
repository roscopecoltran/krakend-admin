# Snippets fro shell/bash:
# gox -verbose -os="darwin" -arch="amd64" -output="{{.Dir}}" $(glide novendor)

default: server

generate:
	go generate
	# go generate ./data
	# go install

all:
	@gox -verbose -os="darwin" -arch="amd64" -output="{{.Dir}}" $(shell glide novendor)

server:
	gox -verbose -os="darwin" -arch="amd64" -output="{{.Dir}}" ./cmd/server/...

dev:
	@go build -v $(shell glide novendor) && ./convert 

run:
	@go run main.go 

prepare: glide-gen gormgen-gen

glide-get:
	@go install -v github.com/roscopecoltran/glide

glide-gen:
	@if [ ! -f glide.yaml ]; then yes no | glide create; fi
	@if [ ! -f glide.lock ]; then yes no | glide install --strip-vendor; else glide update --strip-vendor; fi

gormgen-gen:
	@go generate $(shell glide novendor)

gormgen-get:
	@go install -v github.com/MohamedBassem/gormgen