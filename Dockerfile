FROM ubuntu:16.04

WORKDIR /app/

COPY ./build/dp-dd-api-stub .

ENTRYPOINT ./dp-dd-api-stub
