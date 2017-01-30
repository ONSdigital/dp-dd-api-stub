// +build production
// Code generated by go-bindata.
// sources:
// data/datasets/AF001EW/dimensions.json
// data/datasets/AF001EW.json
// data/datasets/CPI15/dimensions.json
// data/datasets/CPI15.json
// data/datasets.json
// data/hierarchies/CI_000641.json
// data/hierarchies/TIME_001.json
// data/hierarchies.json
// DO NOT EDIT!

package stub

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataDatasetsAf001ewDimensionsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x94\x5f\x8b\x9b\x40\x14\xc5\xdf\xfd\x14\x17\x9f\x17\x1c\xc7\x3f\xad\xbe\x05\x9a\x6d\xa1\x2d\x85\x6e\x21\x0f\xcb\x3e\xdc\xea\xad\x0e\x8c\x33\xc1\xb9\x59\x1a\x4a\xbf\x7b\x99\xac\xa1\x59\xab\x68\x21\x11\x44\x38\xea\x6f\xe6\x9e\x73\x98\xc7\x00\xe0\x57\x00\x00\x10\xaa\x3a\x2c\x21\x7c\x27\x84\x88\x65\x16\xde\xbd\x88\x06\x3b\xf2\xf2\x03\xfd\x3c\x4b\x7c\xdc\x9f\x24\xc7\x68\x6a\xec\xeb\xb3\x7e\xe8\xb5\x97\x5b\xe6\x7d\x19\x45\xda\x56\xa8\x5b\xeb\xb8\x94\x42\x14\x45\x54\x23\xa3\x23\x76\xd1\xe6\x5e\x88\x78\xbb\x8b\x6a\xd5\x91\x71\xca\x1a\x17\x8d\xd6\x6c\x15\xf5\xd8\x57\xad\xaa\xd0\x13\x7f\xa0\x76\x34\xbc\xb2\x7b\xf6\x7f\x84\x25\x3c\x9e\x84\xf3\xe6\x2f\x07\xf8\xe2\x69\x59\x32\xd0\x5e\x4d\xf1\x19\x35\x85\x83\xfc\xfb\x6e\x81\x90\x4e\x11\xee\xa9\xbb\x64\x9c\x9e\x4f\xc1\x40\x9b\x30\x32\x1d\x1b\xf9\x95\x9c\xaa\xc9\x54\x04\xdf\xbc\x8f\x37\xf6\x34\xbd\xa6\xa7\x79\x3c\xe5\xc8\x46\x6b\xa8\x90\xa9\xb1\xbd\x22\x57\xc2\x68\xc0\x95\x6e\xe7\x72\x8a\xfd\x49\x3d\x93\x03\x65\x00\xa1\xb5\x07\x47\xad\xd5\xf5\x6a\xe2\x64\x03\x2e\x88\x95\xed\xba\x83\x41\x0d\xe4\x18\xbf\x6b\xe5\xda\x8e\x0c\xaf\x8f\x36\x19\x47\xbb\x69\x6e\x9e\x67\x72\xc5\x3c\x65\x9e\xad\xc9\x73\xd3\x10\xc4\x39\xa0\xa9\xc1\x3e\x53\xbf\xd2\x7e\x99\xe7\x93\xf0\x17\x18\x5b\x90\xe9\x6a\xd2\x9b\x39\x92\xcc\x3c\x29\x59\x4f\x7a\x3b\x47\x4a\x4e\xa4\xb4\x58\x4d\x2a\xe6\x48\x99\xf8\xd7\xaa\xc5\x2a\xe5\xe3\x2a\x49\x11\xc7\xf0\xc0\xc8\xca\xb1\x8f\x18\xde\x93\x6d\x7a\xdc\xb7\x47\xf8\x30\x24\x7f\x1c\x77\xad\x39\x7f\x72\xad\xb2\xe5\x33\x65\xe3\xfe\xf0\x1f\x5d\xfb\x28\x52\xe1\xaf\xc9\xc3\x63\x6b\x1a\xed\xed\xf2\xf7\x0e\x35\xb9\xe5\x00\xb6\x85\x5c\xe4\x2d\x53\x76\x03\x65\xf2\x90\x7f\xbd\x93\xbf\xe9\x05\x4f\xc1\x9f\x00\x00\x00\xff\xff\x0d\xe0\xf1\x60\x30\x07\x00\x00")

func dataDatasetsAf001ewDimensionsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDatasetsAf001ewDimensionsJson,
		"data/datasets/AF001EW/dimensions.json",
	)
}

func dataDatasetsAf001ewDimensionsJson() (*asset, error) {
	bytes, err := dataDatasetsAf001ewDimensionsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/datasets/AF001EW/dimensions.json", size: 1840, mode: os.FileMode(420), modTime: time.Unix(1485792374, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDatasetsAf001ewJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x94\x4d\x6b\xe3\x48\x10\x86\xef\xfe\x15\x85\x4e\x09\x1b\x6c\x59\xd6\x26\xd8\xa7\xf5\x26\x9b\x9d\x39\x64\x18\x12\xc3\x1c\x86\x1c\xca\xdd\x65\xa9\x99\x56\xb7\xe8\x2a\x65\x2c\x86\xfc\xf7\xa1\x65\xcb\x1f\x78\x02\x09\x21\x27\x43\xf5\xab\xaa\xe7\xad\x0f\xff\x1a\x00\x24\x46\x27\x33\x48\xe6\xb7\x69\x3a\xfe\xef\x5b\x72\x11\x43\xaa\x61\xf1\x15\x85\x5b\x54\xc6\x15\x9f\x4f\x05\x62\xc4\xd2\x41\x14\xe0\x8e\xaa\x25\x05\x06\xbf\x02\x29\x09\xe6\xa1\x22\x0d\xb7\x3e\x28\x62\x58\xb6\x10\x88\x8d\x26\xa7\x08\xa4\xad\x29\x46\x98\xd6\xf1\x07\x0b\xda\xa4\x6c\x82\x8d\x09\x4b\x91\x7a\x36\x1a\x59\xaf\xd0\x96\x9e\x65\x96\xa5\xe9\x74\x3a\xd2\x28\xc8\x24\x3c\x3a\xc2\xa8\x48\x30\xbe\x24\x33\x88\x56\x00\x12\x4d\xac\x82\xa9\xc5\x78\x17\xb3\x2d\x4a\xc3\xb0\xfd\x16\xea\xe0\x9f\x8c\x26\x86\x2c\x1d\x8f\xe1\x9a\x1c\x37\x0c\xc4\x62\x2a\x14\x62\x90\x12\x05\x94\x45\x66\xb3\x6a\xa1\xe1\x06\x6d\x8f\x2d\x1c\x39\x35\x8c\x2f\x01\x9d\x06\xff\x44\x01\x7e\x96\x1e\x30\x10\x54\xc7\xc6\xb1\x33\xbe\x7a\xc9\xf8\x59\xe9\x1b\xa6\xd2\x5b\x0d\x3e\x80\xf2\x55\xd5\xb8\x83\x3a\xe7\x17\x7d\x6b\x62\x9d\x4d\x7b\x86\xb0\x28\xe9\x80\x33\x16\x45\x86\x08\xbb\xb1\xa0\xb1\xbd\x80\xec\x0a\xee\x30\xa8\xb2\x33\x37\xec\xda\x13\x07\xe9\x9d\xa0\x92\x5d\x7f\x00\x12\x87\xd5\x66\x72\x96\xd6\x08\xff\x06\xd4\x96\xda\xad\x1e\x20\xa1\x0a\x4d\x37\x88\xda\xd7\x43\xe3\x56\xfe\x1f\xef\x78\x58\xb0\x19\x16\xfe\x69\xd8\xfc\xd8\x2b\xeb\xd2\xbb\x2e\xd3\x5f\x79\x0e\x67\xe9\xf9\x78\x92\x4d\x21\xcf\xf3\xe9\x55\x96\x74\x9a\xe7\x2d\x44\x20\x4b\xc8\x74\x83\xd2\xc9\xb3\x09\xdc\x61\x1b\x39\xf3\x1e\xd3\xd1\x5a\xee\x8f\x55\x0b\x0f\x4b\x02\x74\xce\x37\x4e\x91\xde\x29\x31\x8e\x16\xed\x83\xa0\x18\x16\xa3\x38\x99\x81\x84\x86\xb6\xef\xc8\xec\x95\x41\x21\xfd\xb5\x59\x5a\xa3\x3a\x79\xd4\x7c\xef\xb1\x5f\xb3\x60\xd9\xde\xe5\x6b\xe4\x93\xb7\xc9\xf3\x4d\x7b\x1e\xb7\xc8\x15\x49\xe9\xb5\xb7\xbe\x68\x93\x19\xb8\xc6\xda\xed\x83\x50\xa8\x78\xee\xf4\xb5\x77\xda\xf4\x3e\x92\x2f\xa3\x79\xfc\xbe\x6b\x6e\xa2\x4d\x45\x8e\x8f\x2c\xee\x06\xbd\xb9\xeb\x9b\x34\x4d\xc7\xd9\xdf\x7b\xc2\x7e\xfe\x0f\xb4\xde\x07\xe3\x6e\xc6\x20\x0b\x3a\x8d\x41\xef\x5f\xde\x76\x96\xa3\x93\x72\xa5\xa1\x10\xf7\xd2\x28\x8c\x89\x56\x68\x99\x0e\xb7\xe3\x8f\xb8\xf9\x29\xee\xfd\xee\x8e\x16\x91\xf5\xa3\xc8\xf3\x77\x92\x4f\x4e\xc9\xe7\xc5\xc7\xe1\x4e\xde\x89\x7b\x79\x8a\xdb\xfd\x39\xee\xce\x0b\x2d\xfc\x4f\xbe\x08\x58\x97\x2d\x7c\xda\x96\x68\x4f\xfd\x14\xbd\xe8\x9d\x86\x2e\x5f\x34\x14\x8f\x7c\xe3\x67\x00\xf0\x38\x78\x1e\xfc\x0e\x00\x00\xff\xff\xf1\x31\xaf\xda\xbd\x06\x00\x00")

func dataDatasetsAf001ewJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDatasetsAf001ewJson,
		"data/datasets/AF001EW.json",
	)
}

func dataDatasetsAf001ewJson() (*asset, error) {
	bytes, err := dataDatasetsAf001ewJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/datasets/AF001EW.json", size: 1725, mode: os.FileMode(420), modTime: time.Unix(1485784386, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDatasetsCpi15DimensionsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x97\x5f\x4f\xdb\x3c\x14\xc6\xef\xfb\x29\xac\x5c\x81\x54\xd4\xb4\xfc\x79\x5f\x7a\xc7\xc6\x10\xa0\xb1\x54\xa2\x77\x88\x8b\x83\x73\x48\xac\x39\xb6\x67\x3b\x48\x68\xda\x77\x9f\x12\x9a\x36\x29\x66\x32\x3a\xb9\x42\x72\xea\x9f\x9f\xa7\xf6\x2f\xa6\x0f\x13\xc6\x7e\x4f\x18\x63\x2c\x11\x79\xb2\x64\xc9\xfd\x2a\x4d\xd3\x74\x9e\x4c\xdf\x06\x15\x54\xd8\x0e\x1b\xe4\x02\x24\x83\xa2\xb0\x58\x80\xc7\xee\x03\xfe\xd5\xb4\x1f\xe0\x12\x9c\x13\xcf\x82\x83\x17\x5a\x75\x4f\x4b\x81\x16\x2c\x2f\x05\x07\x99\x2c\x99\xb7\x35\x6e\x9e\xd4\xb6\x19\x48\x4a\xef\xcd\x72\x36\x93\x9a\x83\x2c\xb5\xf3\xcb\x45\x9a\x9e\x9f\xcf\x72\xf0\xe0\xd0\xbb\xd9\xd7\xd5\xcd\xfc\x74\x96\x8b\x0a\x95\x13\x5a\xb9\xd9\x5e\x3e\x6d\x9a\xe5\x5c\xb2\x64\x0f\xed\x40\xd7\xa6\xd7\xe8\x2a\xcb\x2e\x7b\x53\x06\xb5\x0e\xd2\xf9\x21\xbb\xd2\x3a\x67\xa0\x72\xa6\xb4\x3a\x02\xc9\x75\xa9\xa5\xe0\xec\x09\x5f\xd0\x42\x81\x2e\xd9\xcc\xfb\x33\xfd\x68\x89\x6c\xb5\x4e\x17\x21\x7e\x83\x8e\x9b\x7e\x1c\x9a\xfe\xc5\x22\xbc\x45\xe3\x68\x11\x64\x64\x94\x93\x10\xeb\x0e\xc1\xc7\x4d\x3f\x0d\x36\x11\xae\x8c\x9b\x7e\x16\x5c\x5d\xc8\x9f\x53\xc6\x4b\x44\x87\x6d\x21\x2c\x8a\xc8\x36\xff\x85\x78\x99\x90\xae\xe5\x3c\x83\x8f\xe4\xfc\x1f\xac\x65\x6b\x11\xf1\xb5\x5c\x7f\xbb\xf8\xbe\xbe\x4e\xd3\xf0\x2e\x1f\xa4\x8b\x43\x76\x8d\x20\x7d\xe4\x57\x14\xa4\xdc\x61\xde\x68\xc2\x8c\xd5\x79\xcd\xbd\x9b\x32\x30\x46\x0a\x50\x1c\xdf\xaa\xe2\xaf\x5a\x98\x0a\x55\xe4\x36\x06\x4f\xd4\xaa\x04\x5b\x01\xc7\xda\x0f\xd6\x22\x9c\xab\xcc\x97\x68\x59\xb5\x09\xdf\xe4\x6c\x06\xc0\xb4\x4b\x7c\x36\x73\xf0\xe8\x65\xb5\x3f\x32\xe0\x05\x2a\xcf\x1c\xda\x17\xc1\x63\x9d\x0c\x1f\xc5\x4d\xd4\x0e\xd5\x66\x36\x60\xa1\xda\x7b\x40\x38\x9e\x97\xa8\xfc\xa7\x49\xc1\x03\x7a\xa3\xf6\xbb\xc7\xc1\xce\xff\x55\x7d\xd0\xf8\x7d\xca\xf6\xef\xe3\x64\xb3\xc4\xe0\x7a\x58\x37\x6f\xd2\xf9\xbb\xeb\x61\x2d\xaa\x77\x37\x82\xef\x8d\x8d\x7f\x0f\xec\x05\x89\xb8\x07\x3e\x7e\x49\xe3\x93\xad\xc1\xbe\xb2\x45\x3a\x3f\xa6\x08\xdc\x34\xa4\x42\x2e\x8c\x15\x92\x9e\x84\x5c\xe6\xb6\x56\x48\x67\x48\x72\x8e\x8b\xba\xa8\x9d\xa7\x52\xee\xd1\x78\xac\x9e\xd0\x52\x41\x19\xf7\x7a\x04\xcc\x0f\xfd\x32\x4a\x9c\x4b\xe4\xa3\x70\x6e\x41\x75\x06\x9c\x50\xfe\xdd\xe9\x99\x44\xe1\x6c\x4d\xa2\x40\xb6\x26\xd1\x92\x90\xcb\x74\x26\xd1\x18\x92\x9c\x63\x67\x12\x85\x32\x30\x89\x02\xea\x99\x44\xc1\xf4\x4d\xa2\x70\xfa\x26\x91\x76\x6a\x67\xd2\xe9\x48\x26\x51\x38\x5b\x93\x28\x90\xad\x49\xb4\x24\xe4\x32\x9d\x49\x34\x86\x24\xe7\xd8\x99\x44\xa1\x0c\x4c\xa2\x80\x7a\x26\x51\x30\x7d\x93\x28\x9c\xbe\x49\xa4\x9d\xda\x99\x74\x36\x92\x49\x14\xce\xd6\x24\x0a\x64\x6b\x12\x2d\x09\xb9\x4c\x67\x12\x8d\x21\xc9\x39\x76\x26\x51\x28\x03\x93\x28\xa0\x9e\x49\x67\x81\x5f\x43\x93\xc7\xc9\xdf\x00\x00\x00\xff\xff\x41\x4f\x38\x4f\x30\x13\x00\x00")

func dataDatasetsCpi15DimensionsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDatasetsCpi15DimensionsJson,
		"data/datasets/CPI15/dimensions.json",
	)
}

func dataDatasetsCpi15DimensionsJson() (*asset, error) {
	bytes, err := dataDatasetsCpi15DimensionsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/datasets/CPI15/dimensions.json", size: 4912, mode: os.FileMode(420), modTime: time.Unix(1485792406, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDatasetsCpi15Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x4d\x8f\xe2\x46\x10\xbd\xcf\xaf\x28\xf9\x04\x0a\xc2\x66\x97\xdd\xd5\xf8\x94\x88\x68\x15\x4f\x0e\x63\xcd\x90\x53\x94\x43\xd1\x2e\x70\x69\xfa\xc3\xdb\xdd\x36\x8b\xa2\x91\xf2\x23\xf2\x0b\xf3\x4b\xa2\x6e\x37\xb0\xc0\xe4\x90\x3d\x21\x75\xbd\xe2\xbd\xaa\xf7\xca\x7f\xde\x01\x64\xdc\x64\x25\x64\xab\xba\x5a\x7c\xc8\x66\xe1\x41\xf4\xce\x1b\x45\xf6\x33\x0a\xd6\xbb\x2a\x96\xc7\x8a\x67\x2f\xe9\x84\x86\x95\xd1\xae\x57\x64\xa1\xb6\x2c\xc8\x41\xa5\x1b\xfa\x0a\x93\xd5\x63\xb5\x7a\xac\xa7\xf3\xb1\xa7\xb7\x32\x74\xb4\xde\x77\x65\x9e\x4b\x23\x50\xb6\xc6\xf9\xf2\x5d\x51\xdc\xdf\xe7\x0d\x7a\x74\xe4\x5d\xfe\x0d\xbf\x22\x8f\xe1\x3d\x2b\x21\x08\x04\xc8\x1a\x72\xc2\x72\xe7\xd9\xe8\xc8\x7e\xc1\x9b\x68\x9d\x47\xcf\xce\xb3\x70\xb0\x39\xc0\x9a\x15\x01\xea\x06\x9e\x3b\x12\x8c\x12\x7e\xda\xed\x2c\xed\xd0\x13\x4c\xfc\xa1\x23\x30\x5b\x20\x61\xb4\x51\x2c\x00\x85\xe7\x81\xfd\x61\x3a\x87\x75\xcb\x0e\x92\x28\x70\xad\xd9\x3b\xf0\x2d\x81\x32\x03\x29\xd2\x3e\xb4\x75\xe3\xb0\x66\x20\x1b\x6b\x12\x9d\x87\x2d\x0f\x04\x07\x42\xeb\x60\xcf\xbe\x65\x1d\x4b\xbf\xfd\x9a\x48\x0e\x33\xd8\x58\xf3\x42\x1a\x1a\xb3\xd7\x41\xa0\x32\xda\xb7\x51\xe1\x80\x96\x4d\xef\x40\x48\x74\x8e\xb7\x2c\x30\xcc\xe9\xde\x54\x18\x04\xd2\xf9\xf9\xba\x05\x2d\x41\x43\x96\x07\x6a\x60\x6b\x8d\x8a\x22\x46\x3b\x60\xb2\xba\x00\xc3\xe3\x36\x2c\x8e\x07\x6e\x7a\x94\xc9\xca\xb8\xe1\xa0\xae\xee\x6d\x67\x1c\x4d\x41\xb2\xf3\xa3\x91\x21\x18\x46\x7b\x14\xfe\xe4\x0b\x40\xa6\x51\xc5\x40\x3c\xa0\x22\x07\xeb\x5e\xbc\x90\x4d\x70\x80\x8c\x14\x72\x74\x5f\x74\xfc\xa3\xd1\x6e\xbe\x73\x3c\xdf\x99\x61\xde\xbf\x9c\x41\x5d\x6b\x74\xfc\x8f\x1f\x96\x4b\x98\x14\xd3\xc5\xc7\xf7\xef\x61\xf9\xe1\xe3\x7d\x51\x64\x11\xf3\x9a\xe8\x2d\x49\x42\x47\x3f\xa3\x8f\xf0\xc5\x27\x78\x40\xdd\xa3\x3d\xc0\xbb\x62\xf1\xe9\x28\x52\xd3\x57\xff\x74\x85\x5c\xc2\x67\xda\xd8\x5b\x68\xdc\x04\xca\xe7\x53\x74\xb2\x12\xbc\xed\x29\xd5\xd1\x39\x23\x18\x3d\x35\x75\xbf\x91\xc7\x2d\x67\x25\xe8\x5e\xca\x84\x51\xe4\x5b\xd3\x18\x69\x76\x87\xac\x84\xdf\xd3\x50\xc7\xfd\x5c\xdc\xcc\x75\x6a\xb7\x72\x74\x62\xc2\x5a\xc8\xbe\x21\x07\x28\x25\x2c\x81\x75\x13\x12\xf6\xcf\x5f\x7f\xaf\xea\x6a\x06\xab\xba\xfa\x65\x06\x4f\x75\x15\xd3\xf2\x54\x57\x0f\xd3\xd3\xf2\x2e\xef\xcb\x95\x79\xbe\xdf\xef\xe7\x71\xd3\x71\xcb\x79\x8a\x5f\xce\x47\x36\xd4\x4d\x0c\x70\x22\xc9\xbf\x28\x76\xb9\x48\xd2\x52\x25\x41\xbf\x28\xce\x12\xcf\x6b\xfc\xfd\xe3\x2e\x99\x91\x35\xac\x48\xbb\xb4\x8d\x71\xe8\x53\x24\xc6\xef\xc9\x73\x5d\x14\x45\xb1\x38\xdb\x7c\x4c\xca\xf1\x20\xf1\x78\x90\x67\x48\x38\xcc\x98\x96\x8b\xa4\x9e\xeb\xff\xe7\x53\x92\x9f\x25\xe6\x37\x5a\x5a\x26\x8b\x56\xb4\x2c\x50\x26\xcb\xbf\x4d\xda\xd5\x28\xeb\xd0\xbc\x78\x63\x94\xf0\x91\xb9\x55\xef\x2f\x5e\xbf\x57\xf3\x0d\xe9\x7f\x6a\xbe\x0b\xc6\xbc\xde\xfd\x1b\x00\x00\xff\xff\x77\xb1\xf0\x93\xcc\x05\x00\x00")

func dataDatasetsCpi15JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDatasetsCpi15Json,
		"data/datasets/CPI15.json",
	)
}

func dataDatasetsCpi15Json() (*asset, error) {
	bytes, err := dataDatasetsCpi15JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/datasets/CPI15.json", size: 1484, mode: os.FileMode(420), modTime: time.Unix(1485797239, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDatasetsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\xdb\x6e\x22\x47\x10\x7d\xe7\x2b\x4a\xf3\x04\x0a\xe1\x66\x76\x2d\x23\x45\x09\x21\x6b\x85\x8d\xbc\x1e\x79\x89\xa2\x28\xca\x43\x31\x53\x30\x2d\xf7\x65\xb6\xab\x07\x8c\x56\x96\xf2\x11\xf9\xc2\x7c\x49\xd4\xcd\x00\x33\x2c\x6b\x83\xf6\xc5\x16\xdd\xd5\x5d\xa7\x4e\xd5\x39\x3d\x9f\x1b\x00\x91\x70\xa4\x38\x1a\xc1\x5f\x0d\x00\x80\xcf\xe1\xaf\x5f\x4e\xa3\x11\x44\xe3\xdb\x5e\xaf\xff\xee\x8f\xa8\xbd\x5b\x4e\x0a\x76\x46\x91\xbd\xc5\x44\xe8\xe5\xf4\x74\x90\x13\x4e\x52\x65\x07\xe0\x8e\xd4\x9c\x2c\x83\x59\x80\xcb\x08\xc6\x56\x51\x0a\xb7\xc6\x26\xc4\x30\xdf\x80\x25\x16\x29\xe9\x84\xc0\x6d\x72\xf2\x2b\x4c\x4f\xfe\x1f\x2e\xe9\x70\x6d\x61\xa5\xbf\x34\x73\x2e\x1f\x75\xbb\xd2\x24\x28\x33\xc3\x6e\x34\xe8\xf5\x6e\x6e\xba\x29\x3a\x64\x72\xdc\xfd\x02\x8e\x22\x87\x7e\x37\x1a\xed\xcb\x03\x88\x52\xe2\xc4\x8a\xdc\x09\xa3\xfd\xad\xb3\x4c\x30\x94\x77\x40\x6e\xcd\x4a\xa4\xc4\x30\xe8\xf5\xfb\x30\x21\xcd\x05\x03\xb1\x13\x0a\x1d\x31\xb8\x0c\x1d\x24\x12\x99\xc5\x62\x03\x05\x17\x28\x77\x25\x38\xf6\x98\x53\xe8\xbf\x05\xd4\x29\x98\x15\x59\x58\x67\x06\xd0\x12\xa8\x3a\x09\x18\x48\x58\x7c\x8d\x84\x66\x66\x0a\xa6\xcc\xc8\x14\x8c\x85\xc4\x28\x55\xe8\x4a\x9e\x56\x7b\x47\x93\xcf\xb3\xa5\xaa\x03\xb3\x8c\x2a\x38\x7d\x52\x64\xf0\x60\xb7\x25\xa4\xb8\x69\xc3\xe0\x1a\xee\xd0\x26\x59\x28\xae\xb3\xa7\xc9\x37\xd7\x68\x87\x89\xab\xf1\x04\x10\x69\x54\xdb\x6e\x4a\x7a\x42\xf8\xd9\x62\x2a\x69\x53\x39\x07\x10\x91\x42\x11\x9a\x93\x9b\xbc\x23\xf4\xc2\xfc\x64\x34\x77\x96\x2c\x3a\x4b\xb3\xea\x14\x8f\xf5\xe8\x3c\x33\x3a\xdc\xf8\xdd\x70\x08\xcd\x5e\xab\x7f\x35\xb8\x81\xe1\x70\x78\x73\x3d\x88\xf6\x71\xcf\x15\x60\x96\x24\x21\xd3\x2f\xe8\xc2\xb1\xc1\x15\xdc\xe1\xc6\xe3\x1f\x56\xe1\x6b\x7a\x72\x0f\xf5\xc8\x99\x81\x39\x01\x6a\x6d\x0a\x9d\x50\x5a\x8b\x46\xdf\x7a\x94\x1f\x1d\x3a\xc1\x4e\x24\x5e\x04\xce\x16\x54\x89\x41\x66\x93\x08\x74\x94\xc6\xc5\x5c\x8a\x24\x1c\x39\x88\xa5\x8c\x3a\x67\x20\x07\x75\x06\xce\x39\x72\x75\xf9\x91\xe1\x81\xbe\xbf\x2b\x65\x28\x72\x99\x49\x8d\x34\xcb\x4d\x34\x02\x5d\x48\x59\xd9\x74\x64\x15\x8f\x75\x3a\x31\x3a\x15\xbb\xfa\xa2\x0f\xdd\xf1\xee\xae\xe7\x46\xa5\x1d\x47\xfe\x30\x89\xa7\xfd\x37\x2f\xbb\xc3\x51\xc8\xde\x1b\xc2\x3a\x4c\x8c\xe6\x42\x91\x85\xd8\x0a\xaf\x83\xa9\x4e\xe9\x09\x9a\x93\xfb\xe9\xe4\x3e\x6e\x75\x2e\x16\xff\x51\xb6\xb3\xa4\x5f\xc7\x50\x42\xe0\xfd\x58\x78\x71\xcd\x84\xa2\x20\xb4\x8f\x39\x25\x02\x25\x8c\x97\x4b\x4b\x4b\x74\x04\xcd\xa0\x56\xb3\x00\x4a\x8c\x36\x4a\x24\x80\x89\x13\x2b\xe1\x36\x2d\xaf\xc7\x8a\xab\x70\x66\xd6\x1c\xa4\xaf\xcc\x8a\x14\x69\xe7\x8f\xe5\xdb\xc2\x83\x53\xf8\x3d\x89\xec\x60\x21\x56\x04\x1b\x42\xcb\xb0\x16\x2e\x13\x3a\x6c\xfd\xfe\x5b\x99\x64\xd3\x86\xb9\x35\x8f\xa4\x21\x35\x6b\xed\x01\x2a\xa3\x5d\x16\x10\xae\xd0\x0a\x53\xf0\xce\x9f\x76\x43\x7b\x12\x61\x69\x18\xbb\xe5\xe3\x23\xde\x3f\x52\xb2\x62\xe5\x7d\xca\x1a\x15\x40\x6c\x5b\x03\xcd\x49\x2d\x18\xee\x17\x9e\x38\xb1\x12\xa9\x77\xc3\x2d\xa5\x81\x61\x8f\x2e\x2e\x6c\x6e\x98\x5a\x20\x05\xbb\x4b\x2c\xe7\x3d\x2a\x62\x98\x15\xc9\x23\xd9\xaf\x38\x4e\x92\x8b\xcb\xcc\xe6\xed\xd5\x15\x0c\xdf\xbc\xbd\xe9\xf5\xce\x32\x9b\xfe\x35\xbc\x47\x5d\xa0\x0d\x86\x73\xfd\x8a\xe1\xf4\x87\x70\x4b\x73\x7b\x3a\xfc\x5b\x1c\xe7\x48\xb5\x75\x49\x57\xed\xa8\xca\x63\x4d\x6f\xc7\x53\xbe\x90\xdb\xce\x35\x85\x4e\x64\xe1\x5f\x3b\x94\x12\x86\x20\x74\xea\x27\xf2\xbf\x7f\xfe\x9d\xc4\xd3\x36\x4c\xe2\xe9\xaf\x6d\x78\x88\xa7\x61\xba\x1e\xe2\xe9\xfb\x56\x8d\xe0\xba\x36\x79\xd4\xed\xae\xd7\xeb\x4e\xe8\x48\xe8\x46\xb7\x1c\xd9\xae\xd8\x65\x44\x9d\x86\xa1\x2f\x13\x75\x3f\x29\xc1\xdd\xa4\x84\x57\xee\x94\xa1\x9f\x94\x88\x2a\xb9\x9e\x0f\xf6\x76\x86\x39\x7d\xb8\xff\xfe\xee\xdd\x6c\xfc\xb2\x3d\x7d\x11\xb4\x27\xec\x16\x1f\x69\xaf\x5d\xaf\x42\x50\x82\x59\xe8\x25\xec\x5c\x05\x16\x82\x64\x7a\xb1\x47\xed\x52\xbe\x00\xfd\xdd\x5d\x3c\xfb\xf3\x0c\xf4\xa7\xe2\x5e\x28\x80\x54\xee\x36\x7b\xf8\x17\x03\x3f\x95\xed\xb4\xc3\x1e\x64\x7d\x34\xb7\x47\xe2\x7a\x45\x4d\xaf\xaa\x67\x81\x92\xbf\x5d\x3e\x7e\xb3\x3e\x4f\x8d\xf2\x01\x8d\x12\x53\x68\x5f\xc6\x30\xfc\x72\xc6\xa1\xdc\xff\x62\x87\xd6\x85\xd7\x22\x1a\x41\xaf\xbd\xff\x98\x8e\xc9\xc6\xfe\xc3\x75\x04\xfd\xed\xea\x42\x58\x76\x67\x10\xfc\x63\x8e\x4b\xfa\xa1\x1f\xaa\x8e\xfc\x3b\x70\xe9\x99\xbc\x4c\x7b\x00\xeb\x81\xf8\xfa\xfb\x8d\xe7\xc6\xff\x01\x00\x00\xff\xff\x76\x26\x1c\x9b\xef\x0b\x00\x00")

func dataDatasetsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDatasetsJson,
		"data/datasets.json",
	)
}

func dataDatasetsJson() (*asset, error) {
	bytes, err := dataDatasetsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/datasets.json", size: 3055, mode: os.FileMode(420), modTime: time.Unix(1485775541, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHierarchiesCi_000641Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x95\x51\x6f\xda\x30\x10\xc7\xdf\xf3\x29\x4e\x79\x6a\xa5\x0c\x25\x06\x02\xf4\xad\x45\x42\xea\x43\xb7\x49\x4c\xdb\xc3\x34\x4d\x87\x73\x09\x6e\x5d\x3b\xb2\x9d\x6d\xa8\xe2\xbb\x4f\x49\xb6\x0e\x42\x48\xd4\xa5\x7d\x41\xc1\x67\xdf\xef\x7c\xe7\xfb\xdf\x93\x07\xe0\x8b\xc4\xbf\x02\x7f\x79\xfb\x3d\x0c\xc3\x78\x12\xf9\x41\xb9\xa8\xf0\x91\xca\xe5\x75\x4e\x5c\xa0\x84\xeb\x2c\x33\x94\xa1\xa3\xda\xec\x76\x79\x65\xe6\x12\xad\x15\xa9\xe0\xe8\x84\x56\xb5\x4d\xe7\xe5\xb7\xf5\xaf\xe0\xab\x07\x00\xf0\x54\xfd\x36\x40\x13\x16\xc5\xd5\xf6\xca\xf2\x97\xb6\xfc\x78\x0b\x17\xfa\x07\x19\x94\x12\x84\x4a\xe8\xd7\xa5\x5f\xed\xd9\x07\x1d\xae\xc6\xe3\xe8\xd4\x55\x18\xc1\x4a\xeb\x04\x50\x25\xa0\xb4\x7a\x87\x92\xeb\xad\x96\x82\xc3\x86\x4a\x40\x46\xf6\xdf\xa1\x66\xc8\x87\xac\xb6\xd0\x67\xcf\x47\x1b\xcc\x51\x4d\x3d\x36\xb7\x79\x6f\x12\x5a\x28\x53\x76\xe4\xe6\x84\x34\x8a\xe0\xc6\x10\xd6\x57\xe4\x64\x08\xa5\xf5\x8f\x0e\xec\x83\x17\xf1\xc6\x3d\x3c\x06\x77\x84\x6e\x10\x62\xd2\x83\x18\xc3\x4a\xd8\xed\x20\xc4\xb4\x07\x31\x81\x3b\x21\x1f\x06\x21\xe2\x1e\xc4\x14\x3e\x08\x69\xab\xba\xa4\xe8\x86\x15\x65\xd6\xc3\x8a\x61\x65\x0a\x31\xac\x2a\xf3\x1e\xc6\x0c\x3e\x53\x46\x0e\x37\x92\x2c\x08\xc5\x65\x91\x08\x95\x41\xae\x1d\x3a\x4d\xf5\x45\x5d\xb1\x21\x33\xec\xaa\x8b\x9e\x30\xe6\xb0\x2e\x32\x34\x01\xdc\xe3\x63\x00\x76\x67\x8a\xdc\x06\xc0\xb7\x9a\x6b\x89\x8e\xea\x3e\xd0\x2a\x25\x5e\xf6\x1b\x99\xdd\x90\x68\xe2\xb0\x27\x9a\x45\xad\x2f\xb9\xd1\x49\xc1\x9d\x85\x0b\x45\xfc\xb2\x41\x3c\xf8\xf7\xcd\x6b\x89\xa3\x4b\x63\xe6\x67\x35\x86\xc1\xfb\x1e\x41\x83\xb7\x93\x1d\x36\x8a\x60\xa9\xd3\x94\x28\x00\x47\xf8\x27\xe9\x5c\xe3\x1b\x4a\x0f\x2b\xa5\x47\xa8\x72\x2c\xc0\x4f\x74\x64\x6c\x00\x56\xa7\x0e\x12\x23\xd4\x43\xfd\xfe\xee\x0b\xc1\xa9\xf9\xfe\xda\xf3\xef\x1d\xae\xf4\x4c\x16\xd6\x32\x59\x18\x5c\x9f\xe6\xbe\x6e\x02\xbd\x41\xce\xf5\xff\x0f\x96\xc5\x99\xa2\x97\x59\x6f\x81\xbe\x4a\xc1\xe3\xae\xe4\xb3\x6a\xce\xac\x73\x61\xc4\x30\x21\x8b\xbb\xa4\x9f\x55\xd3\xe5\x8b\x50\x04\x83\x18\x5d\xda\xcf\xaa\xf1\x72\x43\x64\x5e\xaf\x49\x59\x78\xb6\x5e\x0c\x3e\x35\x1e\x03\x0c\xa9\x51\xd7\xc8\x61\x23\x36\x0a\x9f\x71\x2f\xef\x00\xaf\xfc\xda\x7b\xbf\x03\x00\x00\xff\xff\x4d\xf5\x86\x06\x0d\x0a\x00\x00")

func dataHierarchiesCi_000641JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataHierarchiesCi_000641Json,
		"data/hierarchies/CI_000641.json",
	)
}

func dataHierarchiesCi_000641Json() (*asset, error) {
	bytes, err := dataHierarchiesCi_000641JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hierarchies/CI_000641.json", size: 2573, mode: os.FileMode(420), modTime: time.Unix(1485788508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHierarchiesTime_001Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x96\xcf\x6b\xc2\x30\x14\xc7\xef\xfd\x2b\x42\xce\x32\x5a\x37\xdf\xdb\xbc\x0d\xb6\xc1\x06\x6e\x87\x79\x1b\x63\xd4\x1a\x66\xa1\xbf\xa8\xa9\x20\xd2\xff\x7d\x18\xb5\xb4\x9a\xb0\xab\xdf\x5c\x24\x7c\x9b\xbe\xbc\x8f\x49\xc9\x67\x17\x08\x21\xd3\xa5\x9c\x0a\x39\x7f\x9d\x3d\xff\x84\x61\x24\x47\xfb\xac\x88\x73\x65\xd2\x34\x57\x87\x44\x6f\x2b\x93\xe8\x2e\x29\x2b\x9d\x96\xc5\x5a\x4e\xc5\x57\x20\x84\x10\x3b\xf3\xdb\x15\x1c\x87\xd1\xc4\x4c\x34\xd9\xa9\xe0\x30\xcd\xd4\x46\x65\xf3\x43\xe5\xd3\xeb\x42\xc8\xa4\x5c\x9a\xc9\x5b\x15\xd7\xdd\xe4\x5e\x91\xf3\xdc\x94\x91\x53\x11\x1e\xa3\xb6\x5b\xe0\xbc\xc7\x7e\x9f\x67\xbd\xde\x1c\xd9\x2f\x56\x7b\x8b\x8b\x26\xae\xb7\xc3\x87\xf6\xce\x07\xdd\xe7\x65\xa1\x57\x83\xb7\x7a\x45\xad\x0f\x4f\x1c\x51\x2f\x6e\xbb\x71\x3b\xfa\x1f\x61\x6c\x47\x78\x51\x8b\x1a\x86\xe1\xd6\xce\x30\x8b\xeb\x64\x05\x01\x70\x67\x07\x78\xac\xea\x34\x83\x00\x98\xb8\x76\x00\xe3\x00\x91\xe3\x3b\x6e\x0a\x05\xd1\x3f\xbb\xfa\xcf\x30\xfe\xff\x7b\xc7\xf9\x6f\x7e\x9b\xb5\x86\x20\x78\xb0\x13\x7c\xaa\x4a\xab\x7c\xa1\x6a\x04\x88\x28\xb4\x43\x7c\x24\xba\x44\x41\x70\xdc\xc8\xef\xe5\x06\x67\x1b\x1c\x57\xf2\x93\x4a\xae\x84\xe1\x38\xfa\x0e\x7a\x44\x97\x2e\x47\x56\x97\xa3\xab\x74\x39\xc2\x77\x39\xf2\xc0\xe5\x08\xdd\xe5\x08\xdd\xe5\x08\xdb\xe5\x08\xdc\xe5\x08\xdc\xe5\x08\xde\xe5\xc8\x07\x97\x23\x7c\x97\x23\x0f\x5c\x8e\xbc\x71\x39\xb6\xba\x1c\x5f\xa5\xcb\x31\xbe\xcb\xb1\x07\x2e\xc7\xe8\x2e\xc7\xe8\x2e\xc7\xd8\x2e\xc7\xe0\x2e\xc7\xe0\x2e\xc7\xf0\x2e\xc7\x3e\xb8\x1c\xe3\xbb\x1c\x7b\xe0\x72\x0c\xe7\x72\xc1\x7e\xd4\x06\x7f\x01\x00\x00\xff\xff\x5e\xd1\xba\xee\xb2\x1d\x00\x00")

func dataHierarchiesTime_001JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataHierarchiesTime_001Json,
		"data/hierarchies/TIME_001.json",
	)
}

func dataHierarchiesTime_001Json() (*asset, error) {
	bytes, err := dataHierarchiesTime_001JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hierarchies/TIME_001.json", size: 7602, mode: os.FileMode(420), modTime: time.Unix(1485796534, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHierarchiesJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x0a\xf1\xf4\x75\x8d\x37\x30\x30\x54\xd2\x81\x88\xe6\x25\xe6\xa6\x82\xc5\x33\x73\x53\x61\x62\x25\x95\x05\x60\xb1\x12\x90\x18\x97\x82\x42\xad\x0e\xba\x29\xce\x9e\xf1\x06\x06\x06\x66\x26\x18\xc6\x04\x17\xa4\x26\x67\x26\xe6\x28\x38\xa6\xa7\x17\xa5\xa6\x27\x96\x60\x98\x99\x9c\x93\x58\x5c\x9c\x99\x96\x99\x9c\x58\x92\x99\x9f\x87\xdd\x74\x77\x7f\x57\x7f\x6c\x6e\x74\x4f\xcd\x4f\x2f\x4a\x2c\xc8\xa8\x44\x37\x34\x1d\x2e\x01\x32\x8f\x2b\x96\x0b\x10\x00\x00\xff\xff\xae\x6a\x14\xb1\xf7\x00\x00\x00")

func dataHierarchiesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataHierarchiesJson,
		"data/hierarchies.json",
	)
}

func dataHierarchiesJson() (*asset, error) {
	bytes, err := dataHierarchiesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/hierarchies.json", size: 247, mode: os.FileMode(420), modTime: time.Unix(1485796515, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/datasets/AF001EW/dimensions.json": dataDatasetsAf001ewDimensionsJson,
	"data/datasets/AF001EW.json": dataDatasetsAf001ewJson,
	"data/datasets/CPI15/dimensions.json": dataDatasetsCpi15DimensionsJson,
	"data/datasets/CPI15.json": dataDatasetsCpi15Json,
	"data/datasets.json": dataDatasetsJson,
	"data/hierarchies/CI_000641.json": dataHierarchiesCi_000641Json,
	"data/hierarchies/TIME_001.json": dataHierarchiesTime_001Json,
	"data/hierarchies.json": dataHierarchiesJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"datasets": &bintree{nil, map[string]*bintree{
			"AF001EW": &bintree{nil, map[string]*bintree{
				"dimensions.json": &bintree{dataDatasetsAf001ewDimensionsJson, map[string]*bintree{}},
			}},
			"AF001EW.json": &bintree{dataDatasetsAf001ewJson, map[string]*bintree{}},
			"CPI15": &bintree{nil, map[string]*bintree{
				"dimensions.json": &bintree{dataDatasetsCpi15DimensionsJson, map[string]*bintree{}},
			}},
			"CPI15.json": &bintree{dataDatasetsCpi15Json, map[string]*bintree{}},
		}},
		"datasets.json": &bintree{dataDatasetsJson, map[string]*bintree{}},
		"hierarchies": &bintree{nil, map[string]*bintree{
			"CI_000641.json": &bintree{dataHierarchiesCi_000641Json, map[string]*bintree{}},
			"TIME_001.json": &bintree{dataHierarchiesTime_001Json, map[string]*bintree{}},
		}},
		"hierarchies.json": &bintree{dataHierarchiesJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

