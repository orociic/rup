package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _config_repos_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\xc1\x8e\xf3\x20\x0c\x84\xcf\x7f\x9e\xa2\xea\xf9\xd7\x46\xea\x7d\x9f\x64\xd5\x83\x03\x16\x41\x05\x8c\x8c\x93\x26\x6f\xbf\xdd\xec\x21\x95\xf0\xc2\x75\xfc\x65\xe4\xcc\x18\x4b\xcf\x14\x08\xec\xe5\xf3\xf2\x35\xfc\xbb\x3a\x0a\x90\xdc\x07\xb1\x1b\xb7\x71\xc2\x64\xe6\x08\xfc\x28\xd7\xff\xd5\x2c\x90\x53\xd4\xc5\x07\x5b\xcb\x86\xf7\x2c\x54\xeb\x16\xa7\x45\x71\xc1\x0d\x62\x0e\xa8\x0d\x72\x2d\xfa\x08\x4e\x61\x23\x4d\x5e\xf3\x48\x28\xb5\x48\xb0\xc8\x7c\xab\xf5\x1c\x60\x77\x4c\x4b\x52\x7e\x8a\x71\xf5\xf8\xac\xf5\xb2\x2b\x71\x09\x04\x2d\x45\xc1\x4d\xd9\x46\x88\xc2\x0f\x7c\x1f\x06\x9f\xca\xeb\xd3\xa0\xb5\x73\x50\xa3\x89\xf6\xb7\x27\x13\x95\x6c\x4e\xc6\xbc\x5c\x1c\x43\x9e\x9b\x10\xad\xc8\x2d\xc0\xfa\xae\x07\x2a\x85\x9e\x53\x47\x16\xb7\x0e\x40\xa6\x0d\xf8\x98\x89\x45\x4b\xf3\x0d\x8a\x6b\x7e\x74\x36\x61\x4c\x10\x95\x0b\x79\x67\x64\xcf\x4d\x62\x96\x18\x6e\xc0\xe2\x8d\x76\x6b\x27\x47\x0c\x1d\x22\x33\x16\x4c\x7f\x5d\xc3\x81\x94\x02\x76\x69\xd7\x5c\x84\x7d\x72\xed\x12\xc5\xe7\x6e\xca\xeb\xf1\x4a\xee\xdf\x01\x00\x00\xff\xff\xa2\x78\xac\xf9\x1e\x04\x00\x00")

func config_repos_toml_bytes() ([]byte, error) {
	return bindata_read(
		_config_repos_toml,
		"config/repos.toml",
	)
}

func config_repos_toml() (*asset, error) {
	bytes, err := config_repos_toml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "config/repos.toml", size: 1054, mode: os.FileMode(508), modTime: time.Unix(1425198201, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"config/repos.toml": config_repos_toml,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"config": &_bintree_t{nil, map[string]*_bintree_t{
		"repos.toml": &_bintree_t{config_repos_toml, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

