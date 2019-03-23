// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package aws

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x9c\xdd\x6e\xdb\xbc\x19\xc7\xcf\x73\x15\x82\x8f\x1b\x2d\x92\x3f\xf2\x71\x56\x64\xe9\x50\x74\x45\x8b\xb9\xdd\x30\x0c\x45\x40\x4b\xb4\xad\x45\x12\x55\x4a\x4e\xe2\x16\xb9\x9a\x5d\xca\x6e\xec\x85\x15\x5b\x96\xc9\x47\xfc\x8b\x60\xd5\xf7\x3d\x78\x1b\x50\x7f\x49\x3f\x92\xcf\x17\x49\xdb\x3f\xcf\x3c\x6f\x94\xb3\x8c\x8f\x6e\xbc\x11\x7b\x2a\x47\x6f\x76\x0d\x3c\x7f\x2c\x47\x37\xde\x7f\xce\x3c\xcf\xf3\x46\x31\x7f\xac\x9b\x3d\x6f\xf4\x9d\x1d\xfe\x2a\xa4\x88\x47\x67\x9e\xf7\xad\xbe\x41\xf2\x55\x22\xf2\xe3\x3d\x3f\xeb\xff\x7b\xde\x28\x15\x11\xab\x12\x91\xef\x1e\xff\x45\x3c\x6c\xc5\xfe\xfe\xe6\x9e\xfa\xbd\xc5\x79\x2e\x64\xb5\xe6\xac\xac\xce\x83\xa3\xe2\x87\xc8\xf9\xf1\x99\x75\xd3\xa9\x94\x35\x5a\xed\x52\xd4\x7d\xa9\xe6\xde\xfd\xf7\xad\xfe\xf7\xe5\x4d\x37\x32\xa4\x0d\xfb\xd3\x86\xdd\xb4\x61\xd4\x1f\xe9\xe3\x26\x5b\xb0\xa4\x03\xac\x14\x9b\x6a\x8d\x87\x70\x2f\x53\x81\xf6\xcd\x8b\xfe\x30\xf3\x24\x5f\xb1\x42\x48\x6e\xe2\xe9\x37\xad\x2d\x29\xc9\xf5\x7a\x69\xd1\x7d\xc9\x62\x0c\xe7\xdb\x38\xe7\x5b\xc8\x0c\x27\xb7\x25\xed\x66\x0e\xbb\x99\x6d\xe6\x9d\xa2\x8d\xd8\x79\xc4\xf3\x4a\xb2\x14\x8c\x6f\x5b\x78\x42\xda\xbe\x60\x31\xef\xef\x24\xcb\x1f\x96\x1b\x59\x51\x54\x7c\xd3\x93\xaa\x2d\x3c\xa1\x6a\x5f\x58\x74\x5d\xb0\x18\xbb\xf7\x92\xa7\x2c\x8f\x3b\x60\x9f\x38\xb4\xcf\x46\xa5\x62\xbe\xb6\xaa\x8c\xaf\xad\x16\x80\x7f\x17\x79\x2c\x72\x13\x9f\xd9\x16\x1b\x15\xc9\x17\x92\x7c\xae\xc6\x77\x78\xce\xb8\x17\xd9\x98\x24\x1b\x93\x64\x63\x1b\x57\xfe\xff\xff\x84\xf7\x99\x6d\x52\x32\xb3\x94\xec\xbc\x47\xf4\x69\x54\x27\x8c\x4d\xeb\x82\x6c\xb5\x60\xfc\x67\x22\x57\x49\x9e\x30\x8a\x70\x53\xf6\x21\x6c\x54\x27\x84\x4d\xeb\x82\x6c\x8d\xc8\xd6\x98\x6c\xe5\x64\xeb\xb2\x7f\x1f\x3f\xad\x13\x72\x06\x0e\xcf\x32\xdb\x6f\xa3\x22\xfb\x17\x92\xfd\xb3\xb1\xdf\x5b\x96\x26\x4b\x21\xbb\xe7\xa0\x47\x0c\x68\x54\x2a\xe3\x21\x06\xf4\x1f\x2b\xc9\x57\xb4\xb7\x1f\x9e\x06\x47\x8b\xf0\xf6\xa6\x55\x1d\x2d\xda\xdb\x9b\xda\x2d\xc9\xcb\x8a\xe5\x11\xff\xb2\x2d\x38\x51\xc1\x95\x0f\x9b\x3a\xd5\x04\x7e\xc6\xe3\x64\x93\x1d\xc9\x62\x5e\x46\x32\x29\x0e\xbd\x22\x14\x11\xab\xf8\x4a\xc8\x6d\x3d\x05\x22\x2b\x36\x15\xf7\x44\x51\x25\x59\xf2\x83\xb7\xa2\x71\x54\xec\x5e\x11\x1e\xc7\x83\x65\xa3\x1b\x2f\xf0\x2f\xc9\x91\x3c\x02\x3d\xa7\x4c\xae\xb8\x09\x48\x55\x58\x01\x5d\x29\x40\x66\x9c\xb1\x1f\x22\x1e\x42\xe2\x04\x14\x4c\x01\xd1\x04\x13\x69\x12\x2b\xa2\x60\xa6\x20\x8d\x2f\x00\xd2\x15\x46\xd2\x24\x56\x48\x63\xd5\x8e\x66\x08\x09\x02\xb9\xe0\xa8\x34\x63\xff\x12\xcd\x1a\x1e\x21\xa7\x01\x9a\xa8\x66\xed\x9b\x81\x26\xd8\xb0\x09\xc9\x90\x86\x3d\xc1\x86\x4d\x48\x06\x35\xec\x09\x36\x6c\x42\x62\x67\xd8\x2a\x12\x30\xec\x09\x32\x6c\x4d\x30\xac\x61\x4f\xa0\x61\xeb\x8a\x41\x0d\x7b\xea\x07\x70\xd2\x28\x8d\x15\xd4\xa5\x96\xd6\x26\x13\x40\x05\xdd\x8d\x90\xb8\xb9\xdb\x0c\x10\x41\x77\x23\x24\x8e\xee\x16\x02\xa4\x6b\x8c\xa4\x49\xdc\xdc\xed\x12\x21\x41\xa0\x5f\xe9\x6e\xc8\x88\xf0\xf0\xfc\x52\x5f\xbb\x32\xe2\x44\x21\x0e\x8f\x94\xc6\x35\xf1\x83\x00\xb0\x0a\xb0\x65\x53\x9a\x36\xd6\xdf\x3e\x7f\xf5\x0e\x15\x74\x09\xcd\x3a\x0c\x01\x92\x0c\xf0\x48\x51\x9a\x36\xd2\x47\x9e\x09\xb9\xb5\x18\xa8\xd0\x1c\x92\xe2\x10\x86\x24\x4a\xd2\x46\x9a\x57\x42\xb2\x95\x45\x48\x9a\x05\x80\x08\x4d\x1c\x25\xb1\x22\xd2\xe6\x2e\x08\x8d\x01\x20\xc6\x36\x4e\x49\xac\x98\xb4\x98\x84\x67\x0e\x23\x39\x11\xa9\x71\x60\x0c\x7c\x6e\x19\xf8\xc1\x0c\x30\x91\x9a\x36\xd5\xdb\x28\xe2\x29\x97\xac\xe2\xb1\x17\xd5\x61\x21\xc9\x57\x2a\xd9\x4c\x45\xbb\xbe\x34\xa6\xb9\x65\x00\xad\x9c\x92\x58\x73\x69\xc9\xd7\x6c\x56\x2b\xec\x7c\x94\x44\x0d\x51\x9f\x7e\x4d\xe9\xbd\xc2\x46\x4e\x49\x7a\xd3\x58\x2e\xde\x56\x63\x6c\x4d\xa4\xc6\xdd\x9a\x26\x57\xc6\x94\xb7\xc2\x8b\x6f\x4a\x62\xcd\x65\x1b\xa5\x56\x78\x05\x4e\x49\xac\xb9\x6c\x73\xcc\xba\x47\x5c\x20\x35\x56\xd1\x4a\x9b\xc5\x70\x6a\x8c\x09\x6b\x1c\x13\x28\x89\x53\xe6\x33\x57\xbe\x6b\x5c\xb2\x50\x12\xb7\xcc\x37\x43\x53\x87\x4c\x8a\x92\xd8\x25\x3e\x6d\x15\x15\x1a\xdd\x6f\x5d\xf6\x80\xa2\x34\x8e\x25\x42\x60\xdc\x23\x4c\x70\x2c\xa7\x24\x03\x16\x52\x09\x2e\xa4\x28\xc9\xa0\x85\x54\x82\x73\x0c\x25\x71\xb3\x27\x10\x9e\x12\x58\x48\x11\x8a\x41\x0b\xa9\xa4\x47\xea\x23\x35\x6e\x01\x13\xa4\xbd\x04\xef\x82\x53\x92\x21\x2d\x1c\x27\x62\x4a\x32\xac\x85\xe3\x24\x4c\x49\x86\xb5\x70\xb4\x0f\xae\x0b\xac\x78\xb4\x00\x3e\xf5\x43\x64\xe0\x19\xaf\x58\x6a\x02\x52\x04\x6d\xa0\xaf\xf9\x43\x2e\x9e\x72\x15\xe3\x3c\x50\x38\xa6\x01\x9a\x2a\x3c\x53\xbf\xd5\xf1\xb3\x00\x4c\x94\x2e\x38\x29\xbf\x79\xce\x25\x4b\xbd\x62\x23\x0b\x51\x72\x34\x4d\x60\x0f\x35\x83\xa7\x82\x84\xc2\x02\x47\x9d\x2d\xb4\xc9\x9c\x05\x7e\x99\xb1\xb4\xdb\x68\x74\x81\x03\x0d\x38\xa1\xcc\xe0\x09\x25\xa1\xb0\xc0\x51\x2d\xc7\xbc\x74\xcb\x70\xf1\x41\x49\x6c\x36\x96\x34\x53\x9e\xf8\x46\xe7\xca\x70\xf5\x41\x49\x6c\x90\xb4\x64\x71\xe5\x1b\xc3\x60\x06\x13\x3d\xa1\xb0\x01\xd2\xe2\xe0\xa5\x6f\xcc\x5f\x19\xce\xa8\x94\xc4\xc2\x8c\xb4\x05\x88\x71\xc9\x9d\xa1\x3c\xa1\x0b\x06\x0c\x3f\x63\x18\x7e\x74\xc5\x90\xe1\x07\x66\x0b\x42\x31\x9c\xc7\x4f\xfc\xe0\x02\xf1\x50\x1a\x1b\xa2\x0b\x15\xc9\xbc\x63\xb3\x7b\x1f\x2a\x5b\x49\x8d\x05\x93\xed\x2a\x3f\xc3\x27\xdc\x94\xc4\xc5\xc5\xcc\x61\x11\x9f\x6f\x53\x12\x1b\xbb\xb6\x5b\xe1\x67\xe8\x28\x59\x17\x38\xf8\xbc\x71\x81\x91\xc1\x53\x64\x42\xe1\xe2\x61\x66\xcb\x99\xfa\x01\x34\x1d\x4a\x63\x43\xa4\xed\xd0\x5e\x9b\xad\x67\xea\x87\xd0\x7c\x28\x8d\x05\xd3\xb5\x76\x5c\x7b\x65\x36\x20\x7c\xa8\x4d\x49\x86\xf3\x30\x7c\xa4\x4d\x49\x06\xf4\x30\x74\x7a\xac\x0b\x06\xf3\x30\x78\x76\x4c\x28\x06\xf3\xb0\x22\xc4\xf9\x82\xd4\xb8\xef\xf0\x5f\x9a\x6d\xa8\xc0\xdb\x54\x94\xc4\x7d\x27\x1d\x6c\xc1\x14\xb0\x82\x25\x14\xd6\x54\xea\x60\x99\x37\x61\x8a\x1e\x9b\x55\xa4\x66\xf0\x73\x9a\x02\x97\xd7\x94\xc4\xf9\xd4\x0f\x8e\x17\xb4\xad\x3f\xe3\x94\x46\xe2\xd1\xa2\x24\x4e\x0b\x36\xe3\x40\x49\xbc\xbb\x47\x49\x6c\x80\x6c\x37\xf7\x24\x9e\x3b\x4a\x32\xe4\x07\x38\x24\x5a\xb3\xe9\x02\xa7\x25\x2d\xda\xda\x93\x70\x99\x44\x28\x9c\x36\x22\xc0\x9e\x9a\xec\xb1\x2a\x21\x35\x36\x50\xb6\xa1\x49\xe2\x65\x09\x25\x19\xd0\xd9\xf0\xba\x84\x92\x0c\xea\x6c\xf8\x73\xb7\x94\x64\x50\x67\x43\x8b\x25\x5d\x30\xac\xb3\xc1\x15\x13\xa1\x18\xd2\xd9\xaa\xc0\xcf\x92\x48\x8a\x4e\x1e\x5d\x70\x82\xa3\x5c\x22\x37\x6a\x2e\xfc\x59\x30\x36\x42\xe0\xdd\x50\x4a\x32\xd8\x8a\xa4\x0a\x81\xd5\xe8\x82\xa1\x16\x00\x55\x88\x36\xd5\x08\x85\x03\x8c\xd1\x9b\x76\xaf\x32\x1b\x8b\x26\x70\xd9\xd0\x07\x28\x39\xcb\x8d\x24\xa7\xd7\x1d\x40\x90\x0b\x85\xe0\x98\x43\x17\x38\xc0\x20\xc3\xc5\x7e\xf4\xdb\x96\x8b\xcf\x3d\x3e\x46\x44\x6a\x9c\x12\x39\xf8\x64\xe1\x73\xe0\x8f\x51\xb0\x21\x35\x56\x99\x33\xd4\xf7\x89\xa6\xc6\x99\x7b\x0e\x78\x9f\xc1\xa2\x44\x4e\xa3\xd5\x87\x0b\x0f\x17\xa1\x71\xa9\x7c\x40\x52\xdf\xbd\xaf\xc7\x1c\x52\x22\xb7\x49\x1c\x5f\x5f\x40\x30\x54\x94\x91\x1a\xa7\xaa\x0c\x54\xae\xbb\x17\xa2\xb2\x8c\xd4\x38\xd5\x65\xd0\x0b\x39\x8a\x54\x94\xc4\xa5\x12\x6a\x8a\xd7\xe6\xbb\xb6\x91\xe4\x31\xcf\xab\x84\xa5\xc4\x37\x6d\x0b\x29\x1e\x93\x98\xcb\x7a\xf9\xfe\xaf\x79\x0b\x33\x29\x8b\x94\x6d\xdf\x09\x99\xb1\x6a\x77\x75\x99\xf0\xb4\xf5\x76\x96\xe7\xa2\xaa\xbf\x5d\xbc\x7b\xea\xe1\x79\xbb\x27\xae\x99\xcc\xb8\xf4\x59\x51\x94\x91\x88\xb9\x1f\x89\xec\x2f\x51\xba\x29\x2b\x2e\xcf\x8f\x2c\xa3\xf6\x37\xfc\xbb\x6e\x8b\xf3\xd2\xf6\x96\xf2\xf5\xb3\x0f\xea\x6d\xfb\xbb\x5e\x1a\xfe\xba\x3b\xa7\xdf\x71\x3e\x76\xe2\xf5\xe7\x68\x22\x91\x2f\x93\xd5\x7e\x64\xee\xdf\xde\xde\xde\xcd\xe7\xf7\x1f\xee\xfe\x7d\xff\xfe\xaf\x2d\x8c\xdd\xb3\x84\xdc\x0d\xfd\x88\x3d\x95\xf7\x2c\x8a\x78\x59\xde\x3f\xf0\xed\x7d\x12\x9f\xca\xfe\x5b\xee\x7f\xef\xa3\x96\x7c\xe0\x5b\xf5\x39\x29\x5b\xf0\xf4\xb0\x91\x52\x96\xde\x07\xbe\xf5\xde\x2b\x0f\x49\xf2\x62\x53\xcf\x47\xc5\x9f\xab\x51\x73\xe5\xe5\x4d\xcf\x6e\xcc\xef\x6e\xff\x71\xf7\xa5\xd5\x9b\xee\xae\x94\x3c\x92\xbc\x6a\xf5\x88\xee\xce\xab\xec\xed\xa1\x53\x1d\x5d\x9a\xd7\x2a\xef\xd8\xb3\x8e\x6e\x15\xac\x2c\x9f\x84\x8c\x5b\x5d\xdb\xff\xa5\x7e\x87\xfc\x61\xb3\xe0\x32\xe7\x15\xf5\x05\xf2\x47\x2e\xcb\xbd\x83\x05\xfe\x95\x7f\x71\xb4\xdb\xfd\xcf\x0c\xb5\x0c\x36\xe6\x8f\xa3\x1b\xaf\x92\x1b\xde\xb6\x2d\x29\x62\xbd\xf5\x3b\xdb\xb7\x9d\xb5\xd1\x54\xe7\x3f\x79\xf9\x75\x9f\x97\x2f\x59\x5a\x12\x6f\x57\x9b\xeb\xd7\xd7\x8d\xa7\xef\xdf\x0d\xc9\xd9\xcb\x1f\x01\x00\x00\xff\xff\xa3\x73\x18\xdf\x57\x49\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 18775, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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