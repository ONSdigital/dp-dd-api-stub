BINPATH ?= build

build: generate
	go build -tags 'production' -o $(BINPATH)/dp-dd-api-stub

debug: generate
	go build -tags 'debug' -o $(BINPATH)/dp-dd-api-stub
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-dd-api-stub

generate: ${GOPATH}/bin/go-bindata
	# build the production version
	cd stub; ${GOPATH}/bin/go-bindata -o data.go -pkg stub data/...
	{ echo "// +build production"; cat stub/data.go; } > stub/data.go.new
	mv stub/data.go.new stub/data.go

	# build the dev version
	cd stub; ${GOPATH}/bin/go-bindata -debug -o debug.go -pkg stub data/...
	{ echo "// +build debug"; cat stub/debug.go; } > stub/debug.go.new
	mv stub/debug.go.new stub/debug.go

${GOPATH}/bin/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/go-bindata

.PHONY: build debug generate
