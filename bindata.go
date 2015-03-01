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

var _config_repos_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\x41\x6f\xc2\x30\x0c\x85\xcf\xe3\x57\x20\xce\x13\x48\xdc\xf7\x4b\x26\x0e\x26\xb1\xd2\x08\x27\x8e\x1c\xb7\xb4\xff\x7e\x8c\x69\xa2\x52\x4c\x7a\x7d\xef\xd3\xab\xeb\xe7\x78\xbe\x67\x62\xf0\xfb\xaf\xfd\xf7\xee\xe3\x10\x98\x20\x87\x23\x4b\x38\xcd\xa7\x2b\x66\x37\x24\x90\x5b\x3d\x7c\x36\x1e\x71\x30\xd4\x31\x92\x6f\x65\x27\x4b\x51\x6e\x75\x8f\xd7\xd1\x48\xc1\x19\x52\x21\xb4\x8c\xd2\x8a\x31\x41\x30\xd8\xc4\xd7\x68\x65\x64\xd4\x56\x64\x18\x75\x38\xb7\x7a\x21\x58\x82\xf0\x98\x8d\x9f\x12\x9c\x22\xde\x5b\xbd\x2e\xc6\xba\x14\xc8\xda\xa2\xe2\x6c\x4c\xa3\xcc\xf4\x0f\x73\x20\x3c\xae\x5c\x28\xf1\x8d\xe3\x88\xc7\xdf\x31\x2f\xbb\x5d\xcc\xf5\xf1\x41\xb2\x3a\x7d\x66\x9f\x5c\xf2\x7f\xed\xba\x64\x6c\xf4\xc5\xb8\x47\x4a\x10\x28\x43\x17\xe2\x09\xa5\x07\xf8\xb8\x99\x81\xc6\x19\xbc\xdc\xc0\x1e\xe7\x0d\x80\x5d\x1f\x88\xa9\xb0\xa8\xd5\xc1\x0a\x4a\x53\xb9\x6d\x4c\x22\x98\x21\x19\x77\xb5\x66\x74\x29\x5d\x62\xd0\x44\x67\x10\x8d\xce\xba\xd0\x17\xc7\x02\x1b\x44\x11\xac\x98\xdf\xdd\xd0\x13\xa9\x15\xfc\xd8\xaf\xb9\xaa\xc4\x1c\xfa\x25\x6a\x2c\x9b\x5b\x9e\x9e\x6f\xeb\xf2\x13\x00\x00\xff\xff\xa1\x60\x7d\xf2\x54\x04\x00\x00")

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

	info := bindata_file_info{name: "config/repos.toml", size: 1108, mode: os.FileMode(508), modTime: time.Unix(1424709265, 0)}
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

