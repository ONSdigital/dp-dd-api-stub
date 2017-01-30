// +build production
// Code generated by go-bindata.
// sources:
// data/datasets/AF001EW/dimensions.json
// data/datasets/AF001EW.json
// data/datasets/CPI15/dimensions.json
// data/datasets/CPI15.json
// data/datasets.json
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

var _dataDatasetsAf001ewDimensionsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x94\x5d\x8b\x9b\x40\x14\x86\xef\xfd\x15\x07\xaf\x17\x1c\xc7\x8f\x56\xef\x02\xcd\xb6\xd0\x96\x42\xb7\x90\x8b\x65\x2f\x4e\xf5\x54\x07\xc6\x99\xe0\x9c\x2c\x0d\xa5\xff\xbd\x4c\xd6\xd0\xac\x55\xb4\xec\x26\x10\x84\x57\x79\x3c\x3e\xef\x61\xee\x03\x80\x5f\x01\x00\x40\xa8\xea\xb0\x84\xf0\x9d\x10\x22\x96\x59\x78\xf3\x14\x1a\xec\xc8\xc7\x77\xf4\xf3\x1c\xf1\x71\x7f\x8a\x1c\xa3\xa9\xb1\xaf\xcf\xf9\xa1\xd7\x3e\x6e\x99\xf7\x65\x14\x69\x5b\xa1\x6e\xad\xe3\x52\x0a\x51\x14\x51\x8d\x8c\x8e\xd8\x45\x9b\x5b\x21\xe2\xed\x2e\x1a\xbd\xa8\x55\xd4\x63\x5f\xb5\xaa\x42\x8f\xf9\x81\xda\xd1\x70\xcb\xee\x59\x59\xe3\xc2\x12\xee\x4f\xc1\x79\xe2\xcb\xa9\xbf\x78\x5a\x96\x0c\xb4\x67\xa3\x7f\x46\x4d\xe1\x10\xff\xbe\x59\x20\xa4\x53\x84\x5b\xea\x2e\x19\xa7\xeb\x43\x30\xd0\x26\xec\xa5\x63\x7b\x5f\xc9\xa9\x9a\x4c\x45\xf0\xcd\xcb\xbb\x86\xc8\xf4\x35\x45\xe6\xf1\x94\x86\x8d\xd6\x50\x21\x53\x63\x7b\x45\xae\x84\xd1\x57\xad\x54\x9c\xcb\x29\xf6\x27\xf5\x48\x0e\x94\x01\x84\xd6\x1e\x1c\xb5\x56\xd7\xab\x89\x93\xb5\x5f\x10\x2b\xdb\x75\x07\x83\x1a\xc8\x31\x7e\xd7\xca\xb5\x1d\x19\x5e\xdf\x67\x32\xee\x73\xd3\x5c\xa7\xc4\xe4\x15\x4b\x94\x79\xb6\xa6\xc4\x4d\x43\x10\xe7\x80\xa6\x06\xfb\x48\xfd\x4a\xe7\x32\xcf\x27\xe1\x4f\x30\xb6\x20\xd3\xd5\xa4\x37\x73\x24\x99\x79\x52\xb2\x9e\xf4\x76\x8e\x94\x9c\x48\x69\xb1\x9a\x54\xcc\x91\x32\xf1\xaf\xaa\xc5\xfd\xc9\xc7\xfb\x23\x45\x1c\xc3\x1d\x23\x2b\xc7\xbe\x62\x78\x4f\xb6\xe9\x71\xdf\x1e\xe1\xc3\xd0\xfc\x71\xbc\x60\xcd\xf9\x91\x17\x6d\x58\x3e\xb3\x61\xdc\x1f\xfe\x63\xc1\x3e\x8a\x54\xf8\xdf\xe4\x31\xb1\x35\x8d\xf6\x8e\xfc\x7f\x87\x9a\xdc\xb2\xf5\x6d\x21\x17\x79\xcb\x94\xdd\x40\x99\x3c\xc3\x9f\x4f\xf2\xb7\xb2\xe0\x21\xf8\x13\x00\x00\xff\xff\x7e\xda\x1d\x2f\x04\x07\x00\x00")

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

	info := bindataFileInfo{name: "data/datasets/AF001EW/dimensions.json", size: 1796, mode: os.FileMode(420), modTime: time.Unix(1485785142, 0)}
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

var _dataDatasetsCpi15DimensionsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x97\x5f\x4f\xdb\x3c\x14\xc6\xef\xfb\x29\xac\x5c\x81\x54\x54\xb7\xfc\x79\x5f\x7a\xc7\xc6\x10\xa0\xb1\x54\xa2\x77\x88\x8b\x83\x7b\x48\xac\x39\xb6\x67\x3b\x48\x68\xda\x77\x9f\x12\x9a\x36\x29\x66\x32\x3a\xbd\x42\x72\xea\x9f\x9f\xa7\xf6\xcf\xa1\x0f\x23\xc6\x7e\x8f\x18\x63\x2c\x93\xab\x6c\xce\xb2\xfb\x05\xe7\x9c\x4f\xb3\xf1\xdb\xa0\x86\x0a\xdb\x61\x8b\x42\x82\x62\x50\x14\x0e\x0b\x08\xd8\x7d\x20\xbc\xda\xf6\x03\x42\x81\xf7\xf2\x59\x0a\x08\xd2\xe8\xee\x69\x29\xd1\x81\x13\xa5\x14\xa0\xb2\x39\x0b\xae\xc6\xf5\x93\xda\x35\x03\x59\x19\x82\x9d\x4f\x26\xca\x08\x50\xa5\xf1\x61\x3e\xe3\xfc\xfc\x7c\xb2\x82\x00\x1e\x83\x9f\x7c\x5d\xdc\x4c\x4f\x27\x3b\xa1\x8c\x6d\xd6\xf0\xd9\x9c\x3d\xb4\x03\x5d\x85\x5e\x8d\xab\x3c\xbf\xec\x4d\x19\x74\x39\xe0\xd3\x43\x76\x65\xcc\x8a\x81\x5e\x31\x6d\xf4\x11\x28\x61\x4a\xa3\xa4\x60\x4f\xf8\x82\x0e\x0a\xf4\xd9\x7a\xde\x9f\xf1\x47\x4b\xe4\x8b\x25\x9f\xc5\xf8\x0d\x3a\x6d\xfa\x71\x6c\xfa\x17\x87\xf0\x16\x4d\xa0\x43\x50\x89\x51\x4e\x62\xac\x3b\x84\x90\x36\xfd\x34\xda\x44\xfa\x32\x6d\xfa\x59\x74\x75\xa9\x7e\x8e\x99\x28\x11\x3d\xb6\x85\xb0\x28\x12\xdb\xfc\x17\xe3\xe5\x52\xf9\x96\xf3\x0c\x21\x91\xf3\x7f\xb4\x96\xab\x65\xc2\xd7\x72\xfd\xed\xe2\xfb\xf2\x9a\xf3\xf8\x2e\x1f\xf0\xd9\x21\xbb\x46\x50\x21\xf1\x2b\x8a\x52\xee\x70\xd5\xb8\xc1\xac\x33\xab\x5a\x04\x3f\x66\x60\xad\x92\xa0\x05\xbe\x55\xc5\x5f\xb5\xb4\x15\xea\xc4\x6d\x8c\x9e\xa8\x45\x09\xae\x02\x81\x75\x18\xac\x45\x38\x57\x79\x28\xd1\xb1\x6a\x1d\xbe\xc9\xd9\x0c\x80\x6d\x97\xf8\x6c\xe6\xe8\xd1\xcb\xeb\x70\x64\x21\x48\xd4\x81\x79\x74\x2f\x52\xa4\x3a\x19\x3f\x8a\xeb\xa8\x1d\xaa\xcd\x6c\xc1\x41\xb5\xf3\x80\x70\x3c\x2f\x51\x87\x4f\x93\xa2\x07\xf4\x46\xef\x76\x4f\x83\x9d\xff\xab\xfa\xa0\xf1\xfb\x94\xed\xdf\xc7\xd1\x7a\x89\xc1\x3b\x61\xd9\xdc\xa4\xd3\x77\xef\x84\xa5\xac\xde\xbd\x06\x42\x6f\x6c\x4f\x97\xff\xce\xea\x09\x97\xff\xc7\x37\x33\x3e\xb9\x1a\xdc\x2b\x9b\xf1\xe9\x31\xc5\xda\xa6\x16\x15\x72\x61\x9d\x54\xf4\x24\xe4\x32\xb7\xb5\x46\x3a\x43\x91\x73\x5c\xd4\x45\xed\x03\x95\x72\x8f\x36\x60\xf5\x84\x8e\x0a\xca\x45\x30\x7b\xc0\xfc\x30\x2f\x7b\x89\x73\x89\x62\x2f\x9c\x5b\xd0\x9d\x01\x27\x94\xff\x71\x7a\x26\x51\x38\x1b\x93\x28\x90\x8d\x49\xb4\x24\xe4\x32\x9d\x49\x34\x86\x22\xe7\xd8\x9a\x44\xa1\x0c\x4c\xa2\x80\x7a\x26\x51\x30\x7d\x93\x28\x9c\xbe\x49\xa4\x9d\xda\x9a\x74\xba\x27\x93\x28\x9c\x8d\x49\x14\xc8\xc6\x24\x5a\x12\x72\x99\xce\x24\x1a\x43\x91\x73\x6c\x4d\xa2\x50\x06\x26\x51\x40\x3d\x93\x28\x98\xbe\x49\x14\x4e\xdf\x24\xd2\x4e\x6d\x4d\x3a\xdb\x93\x49\x14\xce\xc6\x24\x0a\x64\x63\x12\x2d\x09\xb9\x4c\x67\x12\x8d\xa1\xc8\x39\xb6\x26\x51\x28\x03\x93\x28\xa0\x9e\x49\x67\x91\x9f\x40\xa3\xc7\xd1\xdf\x00\x00\x00\xff\xff\x37\xd4\xd0\xb4\x1a\x13\x00\x00")

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

	info := bindataFileInfo{name: "data/datasets/CPI15/dimensions.json", size: 4890, mode: os.FileMode(420), modTime: time.Unix(1485784245, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDatasetsCpi15Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x4d\x8f\xdb\x36\x10\xbd\xef\xaf\x18\xe8\x64\xa3\x86\x25\x27\x4e\x82\xd5\xa9\x85\x8b\xa0\xda\x1e\x56\xd8\x75\x4f\x45\x0f\x63\x6a\x6c\x0d\x96\x1f\x0a\x49\xc9\x31\x8a\x05\xfa\x23\xfa\x0b\xfb\x4b\x0a\x52\xb4\x1d\xdb\xc9\x61\x4f\x02\x38\x6f\xf4\x1e\xe7\xbd\xe1\xdf\x77\x00\x19\x37\x59\x09\xd9\xaa\xae\x16\x1f\xb2\x59\x38\x10\xbd\xf3\x46\x91\xfd\x8c\x82\xf5\xae\x8a\xe5\xb1\xe2\xd9\x4b\x3a\xa1\x61\x65\xb4\xeb\x15\x59\xa8\x2d\x0b\x72\x50\xe9\x86\xbe\xc2\x64\xf5\x58\xad\x1e\xeb\xe9\x7c\xec\xe9\xad\x0c\x1d\xad\xf7\x5d\x99\xe7\xd2\x08\x94\xad\x71\xbe\x7c\x57\x14\xf7\xf7\x79\x83\x1e\x1d\x79\x97\x7f\xc3\xaf\xc8\x63\x38\xcf\x4a\x08\x02\x01\xb2\x86\x9c\xb0\xdc\x79\x36\x3a\xb2\x5f\xf0\x26\x5a\xe7\xd1\xb3\xf3\x2c\x1c\x6c\x0e\xb0\x66\x45\x80\xba\x81\xe7\x8e\x04\xa3\x84\x5f\x76\x3b\x4b\x3b\xf4\x04\x13\x7f\xe8\x08\xcc\x16\x48\x18\x6d\x14\x0b\x40\xe1\x79\x60\x7f\x98\xce\x61\xdd\xb2\x83\x24\x0a\x5c\x6b\xf6\x0e\x7c\x4b\xa0\xcc\x40\x8a\xb4\x0f\x6d\xdd\x78\x59\x33\x90\x8d\x35\x89\xce\xc3\x96\x07\x82\x03\xa1\x75\xb0\x67\xdf\xb2\x8e\xa5\x3f\x7e\x4f\x24\x87\x19\x6c\xac\x79\x21\x0d\x8d\xd9\xeb\x20\x50\x19\xed\xdb\xa8\x70\x40\xcb\xa6\x77\x20\x24\x3a\xc7\x5b\x16\x18\xee\xe9\xbe\xab\x30\x08\xa4\xf3\xf1\x75\x0b\x5a\x82\x86\x2c\x0f\xd4\xc0\xd6\x1a\x15\x45\x8c\x76\xc0\x64\x75\x01\x86\xc7\x6d\x18\x1c\x0f\xdc\xf4\x28\x93\x95\x71\xc2\x41\x5d\xdd\xdb\xce\x38\x9a\x82\x64\xe7\x47\x23\x43\x30\x8c\xf6\x28\xfc\xc9\x17\x80\x4c\xa3\x8a\x81\x78\x40\x45\x0e\xd6\xbd\x78\x21\x9b\xe0\x00\x19\x29\xe4\xe8\xbe\xe8\xf8\x67\xa3\xdd\x7c\xe7\x78\xbe\x33\xc3\xbc\x7f\x39\x83\xba\xd6\xe8\xf8\x8f\x9f\x96\x4b\x98\x14\xd3\xc5\xc7\xf7\xef\x61\xf9\xe1\xe3\x7d\x51\x64\x11\xf3\x9a\xe8\x2d\x49\x42\x47\xbf\xa2\x8f\xf0\xc5\x27\x78\x40\xdd\xa3\x3d\xc0\xbb\x62\xf1\xe9\x28\x52\xd3\x57\xff\x74\x85\x5c\xc2\x67\xda\xd8\x5b\x68\x9c\x04\xca\xe7\x53\x74\xb2\x12\xbc\xed\x29\xd5\xd1\x39\x23\x18\x3d\x35\x75\xbf\x91\xc7\x29\x67\x25\xe8\x5e\xca\x84\x51\xe4\x5b\xd3\x18\x69\x76\x87\xac\x84\x3f\xd3\xa5\x8e\xf3\xb9\xd8\x99\xeb\xd4\x6e\xe5\xe8\xc4\x84\xb5\x90\x7d\x43\x0e\x50\x4a\x58\x02\xeb\x26\x24\xec\xbf\x7f\xfe\x5d\xd5\xd5\x0c\x56\x75\xf5\xdb\x0c\x9e\xea\x2a\xa6\xe5\xa9\xae\x1e\xa6\xa7\xe1\x5d\xee\x97\x2b\xf3\x7c\xbf\xdf\xcf\xe3\xa4\xe3\x94\xf3\x14\xbf\x9c\x8f\x6c\xa8\x9b\x18\xe0\x44\x92\x7f\x51\xec\x72\x91\xa4\xa5\x4a\x82\x7e\x51\x9c\x25\x9e\xd7\xf8\xfd\xeb\x2e\x99\x91\x35\xac\x48\xbb\x34\x8d\xf1\xd2\xa7\x48\x8c\xef\xc9\x73\x5d\x14\x45\xb1\x38\xdb\x7c\x4c\xca\x71\x21\xf1\xb8\x90\x67\x48\x58\xcc\x98\x96\x8b\xa4\x9e\xeb\x6f\x79\x4a\xf2\x1b\x01\x2d\x93\x45\x2b\x5a\x16\x28\x93\xcf\xdf\xc6\xeb\x4a\xff\x3a\x34\x2f\xbe\xa3\x3f\xbc\x2c\xb7\x92\x2f\x4f\xdf\x24\xf4\x86\xe9\x87\x42\xef\x82\x05\xaf\x77\xff\x07\x00\x00\xff\xff\x08\x10\x12\x5d\xb6\x05\x00\x00")

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

	info := bindataFileInfo{name: "data/datasets/CPI15.json", size: 1462, mode: os.FileMode(420), modTime: time.Unix(1485784429, 0)}
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

var _dataHierarchiesJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x0a\xc9\xcc\x4d\x55\xd2\x81\x88\xe4\x25\xe6\xa6\xa2\x8b\x95\x54\x16\x80\xc5\x4a\x40\x62\x5c\x0a\x0a\xb5\x3a\xe8\x26\x38\x7b\xc6\x1b\x18\x18\x98\x99\x18\xa2\x1b\x13\x5c\x90\x9a\x9c\x99\x98\xa3\xe0\x98\x9e\x5e\x94\x9a\x9e\x58\x82\x61\x66\x72\x4e\x62\x71\x71\x66\x5a\x66\x72\x62\x49\x66\x7e\x1e\xd8\x74\xae\x58\x2e\x40\x00\x00\x00\xff\xff\x2f\xc7\x5c\x4f\xa3\x00\x00\x00")

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

	info := bindataFileInfo{name: "data/hierarchies.json", size: 163, mode: os.FileMode(420), modTime: time.Unix(1485788104, 0)}
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

