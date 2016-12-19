build:
	go build -o build/dp-api-spike.exe

debug: build
	HUMAN_LOG=1 ./build/dp-api-spike.exe

.PHONY: build debug
