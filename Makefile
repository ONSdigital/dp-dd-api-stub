build: generate
	go build -tags 'production' -o build/dp-frontend-renderer

debug: generate
	go build -tags 'debug' -o build/dp-frontend-renderer
	HUMAN_LOG=1 DEBUG=1 ./build/dp-frontend-renderer

generate: ${GOPATH}/bin/go-bindata
	# build the production version
	go generate ./...
	{ echo "// +build production"; cat stub/data.go; } > stub/data.go.new
	mv stub/data.go.new stub/data.go

	# build the dev version
	cd stub; go-bindata -debug -o debug.go -pkg stub data/...
	{ echo "// +build debug"; cat stub/debug.go; } > stub/debug.go.new
	mv stub/debug.go.new stub/debug.go

${GOPATH}/bin/go-bindata:
	go get -u github.com/jteeuwen/go-bindata/go-bindata

.PHONY: build debug generate