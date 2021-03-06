// Code generated by go-bindata.
// sources:
// data/add.html
// data/index.html
// data/source.html
// DO NOT EDIT!

package main

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

var _dataAddHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xcd\x6a\x84\x30\x10\xbe\xf7\x29\x86\x79\x80\x0d\x74\x0f\x85\x12\xbd\x14\x4a\x6f\x2d\x6c\x5f\x20\x9b\x8c\x26\xa0\x89\x98\x49\xdb\xad\xf8\xee\x45\xb3\xff\xeb\x52\x4f\x33\xfa\xfd\x8c\x1f\x9f\xb4\xdc\x36\xe5\x03\x00\x80\xb4\xa4\x4c\x1e\xe7\x95\x1d\x37\x54\xbe\xf4\xa4\x98\xc0\xd3\x37\x68\xea\xd9\x55\x4e\x2b\x26\x29\xf2\xd7\x4c\x14\x27\xa6\xdc\x06\xb3\x3b\x13\xb1\xeb\xbb\x0a\x76\x7d\x86\xab\x42\xdf\x42\x4b\x6c\x83\x29\xf0\xe3\x7d\xf3\x89\xa0\x34\xbb\xe0\x0b\x14\xa9\x33\x8a\x09\x4f\xe8\x99\xe1\x7c\x97\x18\x78\xd7\x51\x81\xd6\x19\x43\x1e\xc1\xab\x96\x0a\xcc\x44\x84\x2f\xd5\x24\x2a\x50\xcf\xfe\x08\xe2\x4a\xa0\x2b\xdf\x42\x64\xe0\x00\x5b\xaa\x9d\x87\x56\x79\x55\x3b\x5f\x3f\x4b\xd1\xdd\x40\x2f\xec\x98\x7e\xf8\x60\x66\x43\xe4\xa3\x15\x42\x74\xbf\x54\xe0\xd3\x23\x82\x4a\x1c\xaa\xa0\x53\x2c\xf0\x38\x4e\x47\x2c\xa9\x6f\x42\xea\x35\x2d\x1a\x5f\xec\xf3\xbb\x48\x0d\x69\xde\xdb\xc7\x99\x89\xb7\xb0\xe9\x19\x06\xe8\x95\xaf\x09\x56\x19\x16\x61\x1c\x17\x81\xb3\x6e\xe8\xa6\xd8\xca\x61\x80\x15\x8c\xa3\x14\xfb\xfd\x9e\x32\x79\xb3\x24\x27\x45\xbe\xef\xea\x47\xfe\x8b\x34\xa6\x6d\xeb\x4e\x41\x6e\xf6\xeb\x42\x5c\xd3\x7d\x3a\xf6\xd5\xab\xa3\xe6\xe2\x02\x29\xa6\x12\x1d\x2a\x99\x7b\x28\x45\xee\xf7\x5f\x00\x00\x00\xff\xff\x9f\x64\x44\xd5\xe7\x02\x00\x00")

func dataAddHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataAddHtml,
		"data/add.html",
	)
}

func dataAddHtml() (*asset, error) {
	bytes, err := dataAddHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/add.html", size: 743, mode: os.FileMode(420), modTime: time.Unix(1509677887, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x8f\xdb\x36\x10\xbe\xf7\x57\x4c\xd9\x83\x6d\xc0\xb0\xb0\xcd\x21\xc1\xae\xa4\x62\xeb\x24\xc8\x02\x7d\x04\x51\x73\x2a\x7a\xa0\xc5\x91\xc4\x86\x22\x05\x72\x94\xd4\x70\xf5\xdf\x0b\x8a\x8a\x9f\xb2\xd7\xdd\x45\x74\xd1\x63\xe6\xe3\xcc\xf7\x71\x66\xc4\xb8\xa2\x5a\xa5\xdf\x01\x00\xc4\x15\x72\x11\x1e\xfb\x57\x92\xa4\x30\x5d\xa2\x25\x20\x63\x54\x1c\x85\x0f\x3b\x07\x97\x5b\xd9\x10\xd0\xba\xc1\x84\x11\xfe\x43\xd1\xdf\xfc\x33\x0f\x5f\xd9\xce\xcf\x5f\x45\xab\x73\x92\x46\x83\x30\x0f\xf4\x71\xca\x73\x9a\x43\xc3\xa9\x9a\xc1\xe6\xc0\xcf\x5f\xc2\xe4\x6d\x8d\x9a\x16\x25\xd2\x1b\x85\xfe\xf1\xe7\xf5\x83\x98\x32\x0f\x60\xb3\xc5\x67\xae\x5a\x84\xa4\xc7\xdf\x5d\x8f\xe6\x7d\x02\x7b\x78\x9e\xd3\xff\x80\x17\x6c\xb6\x70\xed\xaa\x96\x34\x9d\x9d\xc2\x2c\x52\x6b\x35\x14\x5c\x39\x3c\xb4\x76\xe7\x75\xc8\x9e\xa8\x43\xf6\x3c\x21\xb2\x67\x2a\x91\x3d\x5f\x8a\x38\x0a\x55\x32\x14\x5e\xb4\xab\xbc\x78\x65\xc4\x7a\xaf\xc6\x9a\xf4\x9d\x84\xcd\x06\x16\xad\x43\xbb\x78\x53\x73\xa9\xee\x85\xb0\xe8\x1c\x74\xdd\x1c\xd6\xa6\x05\x6e\x11\x94\x29\x4b\x14\x20\xf5\xf7\xf0\x27\xc4\x1c\x2a\x8b\x45\xc2\x22\x65\x4a\xd3\x12\x4b\x7f\xe9\xef\x30\xa5\x4a\xba\x59\x1c\xf1\x14\xfe\x3d\xf1\xfa\x29\x2f\x92\x9b\xad\xeb\xf2\x6d\xef\xf6\x57\x1c\x35\xbb\x6c\x36\x1b\xb0\x5c\x97\x08\x8b\x1a\x9d\xe3\x25\xfa\x24\x0e\x28\xc6\x0d\x38\x5a\x2b\x4c\x58\xc3\x85\x90\xba\xbc\xbd\xc1\xfa\x0e\x56\xc6\x0a\xb4\xc3\x33\xcf\x3f\x95\xd6\xb4\x5a\xdc\x82\x2d\x57\xd3\x1f\x5f\xbc\x9a\xc3\xcd\xab\x17\x73\xb8\x79\xf9\x72\x76\xc7\x52\xcf\x16\xba\xee\x38\x32\x6a\xb1\x1f\x2d\x2e\x8c\xad\x41\x8a\x84\x15\x0c\x6a\xa4\xca\x88\x84\xbd\xff\x3d\xfb\x83\x41\xd8\xdf\x84\x45\x6d\x23\x38\xe1\x51\x2b\xc6\x52\x37\x2d\xf5\xc8\xbe\x8e\x86\xfe\xad\xa4\x10\xa8\x19\x68\x5e\xe3\x57\x4b\x74\x16\x39\x94\xd0\x28\xf6\xab\xed\x08\xed\x69\xe5\xce\x16\x6f\x25\xaa\x43\x26\x91\xa7\x92\x8e\x30\xcb\xce\x53\x73\xa6\xb5\xf9\x05\x6a\xd9\xd3\xb9\x65\xdf\x94\x5c\xd3\x4f\x54\x59\xc8\x9c\x13\x3a\xa8\xb9\xe6\x25\x8a\xdb\x83\xdd\x8e\x89\xaf\x14\x0e\x55\x93\xb0\x70\x3f\xe6\x4a\x36\x3d\xe9\xbd\x98\xaa\xf4\x37\x5e\x63\x1c\x51\x35\x6e\x7d\xcd\xd7\x0e\x3e\x60\xcd\xa5\x96\xba\x3c\xef\x97\xf5\xfa\x9e\xb7\x2f\x2b\xae\x14\xea\xf2\x82\xcb\x7d\x2f\x95\x3b\x75\x88\xa3\xe3\xdc\x77\x8d\x95\xa3\xa5\x93\xae\x3a\xcb\x37\x18\x44\xba\xed\xe5\x8a\xa8\x71\xb7\x51\xe4\xf7\xc3\x0b\x01\x5d\x17\x1a\x6a\x78\xf1\x4d\x1d\x47\x24\xce\x2f\xe5\x9d\xbd\x48\x5b\x8d\x7a\xd4\x63\x88\xa5\x45\xf1\xae\x5d\xf9\xad\x5d\x04\xe9\xa0\xeb\xf6\x47\xd1\x0f\x0c\x8c\xce\x95\xcc\x3f\x25\x6c\x98\x90\xe1\x17\x30\x09\x95\x3c\x99\x4f\xfc\x3a\xef\x39\x55\xd0\x75\x13\x3f\x07\x96\x15\xef\xf5\x0d\x63\xe8\x52\x06\xa3\x86\x41\x57\x59\x1c\x66\xb7\xdd\xb8\x31\x91\x0f\x16\x6e\x2c\x9e\x70\xdb\xa2\x17\x0f\xda\x91\x6d\xc3\x16\x87\x69\x65\xf1\x7c\x22\xfe\x7a\x5c\x8d\x8f\xd3\x49\x6e\xea\x46\x21\x8d\xea\x31\x98\x82\x22\x97\x38\x1f\x8d\xca\x03\x56\xcf\x12\x32\xab\xcc\x97\xd7\xe8\x93\x80\xae\xbb\x8a\x8f\xc0\x33\x6c\xc2\x32\x81\xcb\xe5\x94\x8f\xe2\x7f\x40\x8d\x5f\xae\x0d\xcf\x5b\x32\x23\xc1\xfb\x35\x9e\x10\xfb\x57\xae\x5b\xae\xae\x0d\x5e\xf7\xde\x23\xe1\xc3\x32\xd7\xc4\x1f\xdf\xae\xd1\x01\x72\xfc\x7f\x8c\xfa\x21\xba\x73\xdb\x3f\x17\x70\x21\x58\x7a\x2f\x44\x48\x01\x00\x86\x73\x48\x38\x7c\xc4\x51\x38\x14\xff\x17\x00\x00\xff\xff\xa2\xfb\xf1\xed\x1c\x0b\x00\x00")

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 2844, mode: os.FileMode(420), modTime: time.Unix(1509677887, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSourceHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xc1\x6a\xc3\x30\x0c\xbd\xef\x2b\x84\x3f\xa0\x3a\xf4\x36\x9c\x5c\x06\xbb\x6e\xd0\xfd\x80\x1b\xab\xb3\x21\xb1\x4d\xac\x0c\x4a\xf0\xbf\x0f\xdb\xd9\xda\x34\x19\x2c\xa7\x48\x7e\xef\xe9\x49\x3c\x69\x78\xe8\xdb\x27\x00\x00\x69\x48\xe9\xfa\x5b\x4a\xb6\xdc\x53\xfb\x62\x94\xfb\x24\x88\x7e\x1a\x3b\x92\x58\x9b\x15\x8f\x37\x82\x3c\x7b\x7d\xbd\xe3\x9a\xe3\x23\xd1\x1c\xef\x9e\x2f\x7e\x1c\x60\x20\x36\x5e\x37\xe2\xfd\xed\xf4\x21\x40\x75\x6c\xbd\x6b\x04\x4e\x41\x2b\x26\x71\x43\x17\x86\x75\x61\x62\xe0\x6b\xa0\x46\x18\xab\x35\x39\x01\x4e\x0d\xd4\x88\x4a\x14\xf0\xa5\xfa\x89\x1a\x51\xe7\x09\xc0\x7f\x0b\x18\x1f\xf9\x97\x3e\xcf\x70\xc8\x0d\x48\x69\xab\x11\xd6\x4b\x01\xfb\x67\x89\x61\x03\x5a\xd5\xa5\x17\xa9\xa7\x8e\x97\x79\x8b\xc3\x2d\x2c\x7f\xf3\x0c\x63\x19\x71\xa8\xb0\x08\x29\xed\x02\x8b\xae\x0f\x79\xf7\x36\x9b\x86\x94\x24\x2e\xf5\x5f\xca\xe4\xf4\x9e\x9c\xc4\xea\xef\x61\x91\x9d\xcd\x56\x57\x8c\xd3\x79\xb0\xb7\xcb\x9d\x96\x12\xdb\x0d\x33\xfb\xeb\xe2\x78\x79\xb5\xd4\xaf\x1c\x48\xcc\x49\xf8\x89\x53\xcd\x90\xc4\x1a\xc9\xef\x00\x00\x00\xff\xff\x3b\x09\x10\x5f\x9a\x02\x00\x00")

func dataSourceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataSourceHtml,
		"data/source.html",
	)
}

func dataSourceHtml() (*asset, error) {
	bytes, err := dataSourceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/source.html", size: 666, mode: os.FileMode(420), modTime: time.Unix(1509677887, 0)}
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
	"data/add.html": dataAddHtml,
	"data/index.html": dataIndexHtml,
	"data/source.html": dataSourceHtml,
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
		"add.html": &bintree{dataAddHtml, map[string]*bintree{}},
		"index.html": &bintree{dataIndexHtml, map[string]*bintree{}},
		"source.html": &bintree{dataSourceHtml, map[string]*bintree{}},
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

