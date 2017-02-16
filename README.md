dp-dd-api-stub
================

A stub for our APIs written in Go. 

### Getting started

Make sure [govendor](https://github.com/kardianos/govendor) is installed and
`GOPATH` and `GOROOT` environment variables are set correctly.
Clone the repository. Run `make debug`. 

The command will create/update `data.go` 
using [go-bindata](https://github.com/jteeuwen/go-bindata) based on the JSON files inside 
the `stub` directory. 

NOTE: there are still some handlers that have not been transformed to use stubbed JSON.

### Configuration

| Environment variable | Default | Description
| -------------------- | ------- | -----------
| BIND_ADDR            | :20099   | The host and port to bind to

### Examples

http://localhost:20099/datasets
http://localhost:20099/datasets/AF001EW
http://localhost:20099/datasets/AF001EW/editions/2016/versions/1
http://localhost:20099/datasets/AF001EW/editions/2016/versions/1/dimensions
http://localhost:20099/datasets/AF001EW/editions/2016/versions/1/dimensions
http://localhost:20099/hierarchies
http://localhost:20099/hierarchies/TIME_001

### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.