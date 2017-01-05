build:
	go build -o build/dp-api-spike

debug: build
	HUMAN_LOG=1 ./build/dp-api-spike

.PHONY: build debug