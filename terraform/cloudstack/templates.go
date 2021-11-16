// Code generated by go-bindata.
// sources:
// templates/base.tf
// templates/sec_group.tf
// DO NOT EDIT!

package cloudstack

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

var _templatesBaseTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6f\x6b\x23\xb7\x13\x7e\xbf\x9f\x62\x58\x7e\x2f\xe2\xc3\xd9\x38\xf9\x35\xe1\x2e\x60\x8e\xd2\x52\xe8\x9b\xb6\xb4\x70\x6f\xc2\x21\xe4\x5d\xd9\x56\x23\xaf\x84\xa4\xb5\x93\x1e\xfe\xee\x45\x2b\xad\x57\xda\xd5\xda\xde\x24\xb4\x47\xb9\x0b\x84\xcb\x6a\xe6\x99\xd1\x33\x8f\x46\x7f\xb6\x58\x52\xbc\x60\x04\xd2\x9c\xf1\xaa\x50\x1a\xe7\x8f\xe8\x2f\x5e\x92\x14\xbe\x24\x00\xfa\x59\x10\x98\x83\xd2\x92\x96\xab\x64\x9f\x24\x51\x7b\x52\x16\x82\xd3\x52\x8f\xf1\xc1\x82\xa2\x47\xf2\x3c\xc6\x45\x91\x5c\x12\x8d\x70\x9e\x13\xa5\xce\x72\x2e\x89\xde\x71\xf9\x88\xb6\x22\x47\x7c\xb9\x24\xc6\xc2\x73\x02\x68\xfd\x00\x0a\xb2\xc4\x15\xd3\x30\x87\xf4\x47\xfb\xdf\x9f\x15\x67\x58\x93\xe2\x17\x0b\xf3\xab\x43\xf8\x89\xcb\x4f\x22\x77\x1f\x55\x1a\x86\x7c\x61\x28\xf8\xf4\xdb\x0f\x70\xf0\x1b\xa4\x20\xe7\x1b\x51\x69\x32\x2a\x82\x5a\x63\x49\x8a\x8c\x61\xb9\x22\x1d\x64\x52\x6e\x11\x2d\x4e\xd2\xa8\xd6\x5c\x6a\x74\xa6\xb1\x21\x20\xa7\x85\x3c\x27\xb5\xeb\x59\x56\xff\x5c\x5d\xdf\x75\x32\x2b\x4a\xd5\x05\x60\x54\xe9\x0b\x8b\x32\x09\x60\x1e\x12\x63\x90\xbe\xcf\xea\x9f\x74\x9a\x00\x7c\xee\x4c\x80\xe4\x95\x74\x92\x6e\xfd\x96\x98\x29\x12\x1a\x52\xc5\x91\x22\xab\x0d\x69\xc4\x1c\xb1\x66\x3c\xc7\x4c\xd9\x51\x2a\x49\xae\xb9\x44\x25\xde\xd4\x59\xda\x4c\xd3\x05\x57\xeb\xcb\xff\x7d\xd9\x62\x99\x59\xd6\xf6\x69\x02\x40\x4b\x4d\x64\x89\x59\x4d\x4f\x6b\x6d\xfe\x52\xd5\xa2\x24\xfa\xc2\x38\x34\xfc\x4d\xe1\xfd\x14\x66\x13\xdf\x6f\xb5\x03\x08\xfd\xd6\x5c\xe9\x8b\x3a\xa1\x2c\x40\x9f\xc2\xb5\xf1\xfc\xb3\xda\x88\x05\x7f\x42\x87\x31\x2a\x4e\x7b\xde\x4e\xfc\x99\xf9\xae\xa7\x3c\xef\x26\x86\x1f\x21\xf9\x96\x16\x44\xfa\xc2\xb5\x64\x9a\x35\x5f\x49\x66\xd3\x37\x73\x8d\xf4\x10\x67\xf6\x48\x9e\xa3\x66\x6e\x2c\x01\x70\xdd\xc0\x18\xf6\xac\x7a\x9d\xc2\x08\x89\x6e\x08\xaf\xb4\xab\xd0\x87\xd9\xac\x56\x9c\x24\x8a\x57\x32\x27\x90\x6a\xa6\x90\x90\x74\x8b\x35\xb1\xbd\xa5\xae\x22\xda\x6e\x9c\x10\x31\x5b\x71\x49\xf5\x7a\x63\xdc\x7f\xff\xe3\x7b\x53\x51\xa9\x30\x5a\x50\xad\x0c\xe4\x77\xb3\x0f\x77\x21\xa2\x9f\x90\x5a\x1b\x50\x81\xa9\xec\x01\x7b\xda\x99\x43\x1a\xa8\x06\x1d\x2c\x13\x00\x51\x2d\x18\xcd\xdd\x84\x3b\xd9\x66\x8d\x61\xd6\x5a\x21\x2e\x48\xa9\xd4\x7a\x30\xa9\xad\xc8\xd3\x7a\xbd\xf6\xf2\xe8\x67\x72\x69\xcc\x12\x00\x5f\xba\x8e\xf7\x46\xb0\x09\x80\xdf\xfc\xbc\xd1\xe6\x53\x02\x60\xb6\x96\x8e\x7f\x67\xe3\x19\x4c\xb7\x69\xe6\x38\x67\x29\xa4\x98\x31\xbe\x43\x98\xb1\x20\xf9\xb9\x3f\xe0\x12\xa2\x85\x11\x6e\x30\x6d\x93\x55\x46\x8b\x73\x42\x21\x59\x31\x92\x42\x4a\xcb\x95\x34\x62\xea\xc4\x35\x16\xdd\x00\x9e\x77\x76\x30\x37\xe1\x8c\x62\x2a\x46\x6a\x47\xe3\xaa\x29\x2f\x5b\xbe\x6b\xd3\xb4\x1e\x32\x74\x22\xd3\xef\xec\x90\x6d\x70\x00\x69\xd3\x2e\x67\x75\x93\x33\x6d\xce\xfc\x16\x92\x6b\x9e\x73\xe6\x01\x59\x18\xc1\xa5\xd1\x66\x13\xe1\xc1\x9a\x6b\x89\x97\x4b\x9a\x23\xd7\xc3\x9b\x99\x19\x97\xfd\x18\x46\xc8\x7f\x95\x10\x72\x3e\x1f\xcd\x62\xb6\xfd\xdb\x32\x10\x2c\x91\x33\x5b\x7c\xb0\xf6\x82\xf5\xe7\xef\xbd\xfb\xcb\x7a\x63\x71\xc1\x5a\x79\x07\xa1\xa2\x42\x37\x0d\x5d\x09\x86\x9f\x91\x26\x4f\x3a\xd8\xa6\x5a\xb4\xa6\x4c\x9d\x15\x1c\x3b\x44\x75\x57\xf2\xf0\x6a\x3e\x28\xa2\x6b\x68\xf7\x64\xf8\x08\x84\x11\xb3\xdf\x5e\x18\xea\x78\x99\x63\x7d\x31\x20\x1d\x8f\x69\xd3\xe1\xd1\x4a\xf2\x4a\x64\xef\x32\x5a\x4c\xe1\x21\x4d\x3f\x4f\x8c\x06\x66\xe6\xd7\x04\xee\xcf\xd2\xdf\xe9\xea\x9a\x03\x17\x65\xd8\x08\xf3\x35\x45\xbe\x1e\x53\x64\x2f\xe6\xdb\xd5\x3a\x0e\xfa\xad\xe4\xb1\x92\x97\x5a\x72\x86\x04\xc3\xcd\x35\x68\x4c\xb5\xef\x46\x57\xbb\x0e\x77\x69\xc3\xbd\x45\xa1\x3b\x78\x5f\x7b\x8d\x03\xbe\xff\xb1\x2a\x17\x58\xe3\x57\x94\xf8\x66\x4c\x89\x4d\xac\x37\xab\x6f\x08\xf6\xb5\x17\xb7\xa5\xf9\x5f\xa8\x2c\xb2\x87\x61\x57\x60\x5e\x95\xba\x3f\x77\xef\xc6\x07\x1f\xe1\x1a\xee\x61\x36\x5e\x0d\xff\x7f\x99\x1a\x2e\x5d\x7e\x6f\x29\x0a\x0f\xf3\x9b\x36\x9c\x03\x15\xb8\x28\xea\x33\x1d\xa4\xcd\x95\x98\x50\x61\x75\x71\xe2\x8a\x70\x20\x66\xf0\xb2\xc2\x2b\x2d\x2a\x1d\x04\x74\xaf\x06\xf5\x2d\xcc\xc8\xc2\x45\xc2\xac\x22\xc7\xee\x78\x1e\x58\x70\x13\x6d\x9d\x2d\xdd\xc3\x97\xbf\xf6\x23\x12\x64\x53\xdf\x92\x4b\x45\x35\xdd\x9a\xb8\x5a\x56\x41\xc2\xe4\xe9\x70\xb5\x0f\x13\x8c\x51\x97\x79\xc4\x65\x54\x20\xf7\xd9\xc7\x6b\x2c\x2a\xc9\x7a\x13\x1e\x01\xb9\xbf\xbf\xb9\x09\xb8\x38\xbc\x45\x1c\xaa\x18\x80\xaf\xb5\x16\xea\xfe\xea\x6a\x6c\x90\xdb\xdb\xdb\xdb\x20\x4e\xa3\xab\xc1\x8a\x1d\x39\x8b\xc7\xd2\xed\xe3\xd8\x57\x93\xc0\xc0\xf7\x0c\x9e\x53\x62\x9e\x81\x41\xd4\x73\xb5\x3b\xea\xb7\xda\xc5\x0a\xe6\x3f\xf2\xc4\xdc\x23\xcf\x48\xd1\x09\x9f\xc2\x89\x3d\x2a\x0d\xac\xa0\xf8\x43\x67\x03\xd7\x59\x89\x5d\xe3\x01\xcc\xf6\x59\x7b\x00\xe7\xc8\x8a\x0e\x9f\xb8\x07\xfc\x0f\x6f\x58\x3e\x39\xa5\xea\xfb\x14\xa5\x8a\x56\xcf\x9d\x74\xeb\xdb\xee\x06\x0b\xd1\x9b\xb8\xbd\x28\x47\x1a\xaa\xd7\x4c\x33\x23\xab\x49\x7c\x23\x89\x58\xd7\x9b\xd8\x20\x6c\x70\x38\xf3\x91\xa3\xb0\xa1\xf5\x71\x64\xef\x68\x1f\x66\x1c\x45\xf6\xad\x3d\xdc\x66\xaf\x19\xdc\x68\xb2\xde\x69\x20\x7b\x57\x87\x9b\xb6\x8f\x08\x66\xab\x81\xd9\x64\x02\x27\xff\xcd\xe1\x65\x01\xed\x41\xa1\x1b\xf0\x18\xeb\xdd\x5b\xa7\xa3\x68\x80\xf5\x9e\x75\x43\xd1\xfe\x98\xcc\x56\xbb\x37\x13\xd9\x51\xce\x8e\x03\xad\xb0\x26\x3b\xfc\xfc\x02\x09\x8e\x0d\x1a\x02\x9d\x8c\x3b\x28\xd0\xb1\x71\x7d\xa0\x30\xea\xab\xe4\xdb\xea\x76\xee\x9f\xb9\x8e\x9c\xba\xa2\x68\x4d\x46\xed\x99\xcb\x9d\xba\x00\x5e\xa2\xd0\xf1\x45\xe9\x01\x79\x1c\x19\xfd\xfe\x1d\x00\x00\xff\xff\x7d\x2a\x95\x87\xa3\x1c\x00\x00")

func templatesBaseTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBaseTf,
		"templates/base.tf",
	)
}

func templatesBaseTf() (*asset, error) {
	bytes, err := templatesBaseTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/base.tf", size: 7331, mode: os.FileMode(420), modTime: time.Unix(1623236118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSec_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x98\xff\x6e\xdb\x20\x10\xc7\xff\xf7\x53\x9c\xfc\xd7\x26\x6d\x19\xf1\x8f\x8e\x4c\x8a\xf6\x20\x55\x85\x28\x26\xad\x55\x6a\x2c\xc0\xe9\xa6\xa9\xef\x3e\x61\x63\x3b\x69\xf1\x96\xa5\x24\xf5\x92\x28\x09\x02\xe7\xee\x7b\x9f\xc3\x77\xc8\x5b\xaa\x4a\x7a\x2b\x38\xc4\x6c\x43\xf8\x0f\xc3\x55\x45\x05\xa9\xa5\x32\x3a\x86\x5f\x11\x80\xf9\x59\x73\x00\x80\x35\x88\x52\x9b\x0f\xda\xa8\xb2\xba\xfb\x18\x01\x14\x7c\x43\x1b\x61\x60\x0d\xd7\x91\xbd\x00\xa3\x4f\xed\x6f\x96\xa5\xdd\x20\x49\x92\xa4\x9f\xea\xe6\x6e\xa2\xe7\x28\x1a\x3d\xde\x4a\x7d\x7f\x94\xab\xde\xee\x15\xbe\xc2\xce\x57\x9e\xe7\xce\xd7\xe0\x35\xc9\x93\x1c\x4d\xb9\x7d\x53\xa8\x7e\xff\x4e\x40\x9a\xe2\x55\x37\xc2\x38\xcb\xdc\xc8\x0f\x40\xdf\x53\xc5\x0b\x62\x58\x7d\x94\x8c\x25\x46\x4b\xe7\x13\x21\x47\x3f\xcd\xf2\xaf\x0e\x04\x42\x2e\x11\xd9\x72\x18\xbd\x4a\x89\xcd\x5b\x9f\x39\x8c\xb0\x5b\xc5\x69\x6f\x0f\xa7\xbd\x0f\x9c\xa2\x64\x3f\x9a\xbd\x08\x71\x3e\xcc\x8d\xf1\xaf\xdc\x7f\x57\xa8\xf7\xbb\x42\x49\xda\x8f\x56\x68\x18\x2d\xa7\xe9\x34\xc5\x71\x74\xbc\xca\xaf\xd0\xb8\x23\x14\xd7\xb2\x51\xcc\x6e\x7d\x21\x9b\x42\x1b\xca\x1e\x48\xc5\xcd\x93\x54\x0f\x84\x32\x11\xbb\xad\xa2\x9b\xdb\x8a\x1b\xa2\x39\x23\x77\x4a\x36\x75\xa7\x83\xc9\xa6\x32\x56\xc6\x96\xaa\x85\xe6\xac\x51\x1c\xbe\xc3\x12\xbe\x01\x8a\x00\x2a\xfa\xc8\x5b\x91\xbb\x26\xe2\x08\x60\x5b\x33\x52\x16\xb0\x86\x1d\x9f\xdb\x9a\x2d\xec\xa7\x2c\x0e\x91\x45\x54\x23\xf8\x94\x36\xc2\xef\x14\xd7\xfa\x10\x89\xd6\xd4\x4b\x25\x3b\x6e\x16\x5e\xfb\xd7\xe8\xc6\xca\x8c\x00\xac\x8a\xd6\x8b\xb5\x64\x4a\x59\x41\xf7\x5a\x43\x4c\x85\x90\x4f\x71\xbb\xc4\xca\x42\x11\x9b\xa8\x6e\xa9\xcb\x0c\x40\x8c\x16\xed\xfb\x0b\x8a\xbb\xc4\xdc\xb4\xdf\xb5\x92\x46\x32\x29\x76\x0c\x75\x66\xda\xfc\xc3\xe0\xe1\xba\xbb\xdc\x28\xba\xd9\x94\x8c\xb4\x5b\x62\x0d\xb1\x8b\x3d\x02\x78\x0e\x00\xd2\x15\x88\x0b\x41\x69\x58\xed\x45\x69\xe3\xf1\x14\x44\x2f\xdf\xb2\x0a\x09\xb8\x64\x8f\x07\xdd\x4a\xff\x01\xdc\x36\x94\x76\xc5\x8e\xc8\x58\xa2\x3e\x2f\xc7\x59\x26\x8b\xfd\xd9\xc3\x36\xf5\xbf\x42\x8f\x21\x66\xb2\x32\x4a\x0a\x52\x0b\x5a\xf1\x63\xeb\xd6\x9e\x91\xb0\x95\x6b\x42\x5f\xc0\xda\x35\xe1\xe1\x12\xab\xd7\x14\x4c\x7b\x5b\xcc\x1f\xa5\x95\x64\x37\x96\x5d\x7e\x53\xf9\x0a\x5f\xb5\xa6\xc0\xfa\x0f\x6d\x97\x0e\xf9\x65\xd4\xe7\x44\xfd\xe2\x04\x78\x09\xa8\x9b\xe2\xaf\xa8\x87\xa8\xcf\x82\x3a\x54\x37\x9e\x45\xe5\x9d\x59\x3f\x2e\xa8\xa1\x6f\x6b\xc6\xa3\x85\xb0\x9d\xd8\xa7\x2c\x60\x1b\xf6\x99\xbf\xc4\x1e\xec\xc5\x18\xaa\x01\x9f\x10\xe2\x8c\xbb\xaf\x17\xe9\x29\x5a\xef\xcc\xf1\x9e\xb4\xef\xfe\x09\x72\xd0\xa6\x3b\x03\xc8\xef\xd6\x71\xfd\x35\x36\xe0\xd3\x85\xf9\x56\x59\x1b\xce\xab\xc7\xca\xa7\x87\x1b\xea\x2c\xf3\xfe\x60\xcf\x79\x90\xf9\x1d\x00\x00\xff\xff\xfc\x09\xec\x89\x0a\x18\x00\x00")

func templatesSec_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesSec_groupTf,
		"templates/sec_group.tf",
	)
}

func templatesSec_groupTf() (*asset, error) {
	bytes, err := templatesSec_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/sec_group.tf", size: 6154, mode: os.FileMode(420), modTime: time.Unix(1623236083, 0)}
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
	"templates/base.tf":      templatesBaseTf,
	"templates/sec_group.tf": templatesSec_groupTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"base.tf":      &bintree{templatesBaseTf, map[string]*bintree{}},
		"sec_group.tf": &bintree{templatesSec_groupTf, map[string]*bintree{}},
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