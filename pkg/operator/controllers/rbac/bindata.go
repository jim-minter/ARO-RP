// Code generated for package rbac by go-bindata DO NOT EDIT. (@generated)
// sources:
// staticresources/clusterrole.yaml
// staticresources/clusterrolebinding.yaml
package rbac

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x4f\xaf\xe3\xb8\x0d\xbf\xe7\x53\x18\xdb\xc3\x02\x05\x5e\x06\x45\x2f\xc5\xeb\x71\xb7\x28\x0a\x14\x5d\x60\x30\xed\x9d\x91\x19\x9b\x1b\x59\xd4\x48\x54\xde\xa4\x9f\xbe\x90\x2c\x39\x71\xe2\xc4\x8e\xd3\x99\x53\x2c\x8a\xe2\x8f\x94\x28\xfe\x51\xfe\x50\xfd\xc2\x35\x56\x0d\x1a\x74\x20\x58\x57\xbb\x53\xd5\x82\x3a\x7c\x6a\xd0\xd4\xe4\x15\x1f\xd1\x9d\x14\xa8\x16\xff\x5a\xfd\xfa\x5b\xf5\xaf\xdf\xbe\x54\x7f\xfb\xf5\x1f\x5f\xb6\x1b\xb0\xf4\x1f\x74\x9e\xd8\xbc\x57\x6e\x07\x6a\x0b\x41\x5a\x76\xf4\x5f\x10\x62\xb3\x3d\xfc\xc5\x6f\x89\x3f\x1d\xff\xb4\x39\x90\xa9\xdf\xab\x5f\x74\xf0\x82\xee\x33\x6b\xdc\x74\x28\x50\x83\xc0\xfb\xa6\xaa\x94\xc3\xb4\xe0\x0b\x75\xe8\x05\x3a\xfb\x5e\x99\xa0\xf5\xa6\xaa\x0c\x74\xf8\x5e\xf9\x93\x17\xec\xde\xc1\xf1\x9b\x77\xb8\x71\x41\xa3\x7f\xdf\xbc\x55\x60\xe9\xef\x8e\x83\xf5\x51\xc8\x5b\xf5\xd3\x4f\x9b\xaa\x72\xe8\x39\x38\x85\x99\xa6\xb8\xb3\x6c\xd0\x88\x17\x90\xe0\xd1\x6f\xaa\xea\x88\x6e\x97\xa7\x1b\x94\xf4\xab\xc9\xcb\x52\x81\x66\x4f\x4d\x07\xd6\xa7\x21\x9a\xda\x32\x19\xc9\xa3\x23\x96\x4f\x4d\x1d\x89\x03\xd3\x60\x3f\x8e\x96\x78\x0b\xaa\x0c\xb9\xce\x5f\x36\x6e\xa0\x17\x34\x72\x64\x1d\x3a\x54\x1a\xa8\x9b\x9e\xca\x54\xae\x87\x0f\xc1\xce\x6a\x90\x3c\xe3\xd0\x6a\x52\x69\x2b\x15\x1b\x71\xac\x35\xba\x32\xd5\x5b\xf1\x35\xb0\x40\x4f\xf2\xe8\x8e\xa4\x10\x94\xe2\x50\xb4\xce\xb4\x47\xbb\x14\x3f\x3e\x40\x54\xbb\x6c\xbf\xa2\xb6\x9f\x34\x37\xb7\x12\x6f\x96\x43\xdd\x91\x8f\xce\xe4\xb0\x21\x2f\xee\xd2\x89\x6e\x05\x77\x41\x40\xc8\x34\x1f\xb8\x6b\x99\x0f\xfd\xb9\x84\x7e\x51\x6f\xcc\x11\x34\xd5\x0f\x79\x56\xd8\x08\x96\xf0\x9b\xa0\x89\x7a\xfa\xbb\xca\xa9\xe0\x85\xbb\x42\xac\x71\x4f\x86\x5e\x03\x5d\xb4\x27\x60\xe9\xb5\x13\x04\x9b\xfc\xfa\xd6\xe7\xb3\x33\x39\x3c\x92\x1f\x36\xb8\x06\xec\xd8\x78\xcc\xce\x53\xa3\xd5\x7c\xea\x86\x2b\x90\xdd\x71\x98\x8f\x57\x10\xf7\x41\x67\xc2\x4a\xf5\xb6\x6c\xd1\xf8\x96\xf6\x32\xb9\x05\x67\x25\xfa\xd3\xfe\x81\x48\x4b\xdd\xdc\xf1\x8c\x64\xd5\xc7\xc9\xb5\xaa\x07\x69\xd1\x48\x0e\x04\x77\x7d\x45\xf8\x80\x26\x9e\x27\x7e\x5c\x01\xa5\x70\x8c\xd3\x82\xaf\x83\xfb\xad\x5c\x8f\x7a\xef\xc3\xee\x77\x54\x02\x4a\xa1\xf7\x67\x8c\xd1\x64\x8a\xe2\xa3\xb9\xe9\x45\x4f\x2b\xb6\x68\x6f\x1d\x6b\xdc\x91\xa9\xc9\x34\xfe\x9a\x9e\xbd\xf7\x9a\xa3\x4c\x2d\x4e\x1f\xcf\xa8\x55\x86\x13\x5b\xf6\x43\xb6\xe5\xc2\x5a\x87\x5e\x1c\xa9\x57\xc2\x55\x10\xf6\x0a\x34\x99\xe6\x16\x29\xa9\xc4\x46\x40\x5b\xae\x0b\xe7\x2b\xce\x5e\xa0\x96\x1d\xfc\x18\xf1\xad\xea\x40\xb5\x64\xf0\x65\x45\x76\x89\x7c\x8b\xea\xd8\xfc\xce\xbb\x1e\x2b\x7f\xac\x91\x1e\x48\xd7\x33\x06\x26\x9e\x73\xd0\xcb\x84\xef\x0d\xb8\x34\xea\x29\x74\x42\xfb\x18\x94\xf0\x41\xda\xbc\x60\xa2\xc6\x24\x67\xfc\x1a\xd0\xaf\xcd\x17\x4a\x73\xa8\x95\xc3\x3a\xc6\x43\xd0\x73\x1e\x32\x30\xfa\x17\x61\xd3\x21\xcc\xa0\xe5\x5c\x5d\xfc\x70\x1c\xb5\xaf\x0e\x70\xf0\x5e\xb6\xb1\x42\x67\x37\x22\x1e\xfb\x12\xdc\x17\x6c\x3f\xc4\xb0\xda\xf8\xfc\xb5\x47\x90\xe0\xb0\x19\x6a\x45\xea\xa0\x94\xa6\x64\xf6\x0e\xbc\xb8\xa0\x22\x4b\xa1\xc5\x38\x50\x56\x1b\x94\x0f\x76\x87\x7e\xc0\x51\xd5\xfc\x99\xd5\x69\x43\xf6\x70\xeb\x38\x06\xab\x61\xf0\x8d\xb2\x04\xaf\x5a\xac\xc3\xfa\xeb\x95\xcd\x9a\x3b\xc1\x9e\x4b\x69\xaa\xf9\xc3\x68\x86\x7a\xb4\x29\xb1\x70\x73\x06\xb4\xe6\x46\x93\x39\x8c\xe6\x6e\x08\x86\xb3\x2b\x5e\x6f\xed\x09\x3a\xed\xa1\xb3\x8f\xb3\xc1\x43\x5b\xd8\xd5\x64\x1e\xa7\x51\x8d\xf0\xb8\x5b\x79\x00\x30\xf4\x6c\x77\xa5\x0f\x6d\x8b\x5e\x5f\x2b\xf6\xcd\xce\x7d\x88\xd2\x0b\xad\x11\x3d\x54\xd8\xb7\x72\x2f\x3d\x73\x85\xe8\xbd\xe6\x8f\x5c\xcf\x6e\x87\x2b\x78\xd7\x88\xc8\x1d\x5d\xb7\x83\xe2\xd2\xc4\x8e\xe4\xa4\xf1\x88\xfa\xff\xd1\x4e\xb4\xa8\xbb\x19\xa7\x8e\x2c\xaa\x05\x27\x0e\x2d\x7b\x12\x76\xb4\xd6\xf8\x74\xe9\x67\xe0\x2e\x03\x43\xfa\x14\x87\xd0\x7d\x77\xc0\x84\x32\x60\xcf\x65\x93\x27\xe5\x0a\x34\x17\x16\xe5\xd1\xe2\x3a\x2e\x2d\xca\xfd\xd7\x69\x5b\x62\xde\x7c\x28\x1a\x12\x71\x12\x60\x5d\x30\xab\xc3\x5f\x76\xfa\xa5\xe0\xb5\xf1\x0e\x15\xbb\xb5\xd9\x3f\x5e\x07\x65\x68\xab\x8c\xda\x4f\x02\xe4\x84\xf0\x06\x22\xa0\xda\xd8\x03\xbd\xbd\xdc\xe4\xe6\x4a\x6c\xc6\xb4\xcc\xd5\x22\x68\x69\x55\x8b\xea\x30\xaa\xe3\xc6\x83\xd5\x8d\x66\x16\x30\xba\xe1\xf3\x07\x2e\x40\x06\x9d\x0b\x46\xa8\xc3\x4b\x07\x38\x37\xd0\x97\xd4\x43\xd8\xa1\x46\xb9\x24\x8d\x70\x2d\xb3\x9e\x20\xaf\x35\x09\x05\xf4\x9f\xa7\x0b\x39\x70\x98\xa6\x5b\xf6\xe7\xdc\xdd\x77\xfa\xb9\xfd\x59\x07\xe8\x48\xdd\x4f\x10\x17\x4f\x60\xfc\xd0\x51\x6f\x25\x53\x33\xf3\x10\xe2\x85\x5d\xba\xfd\x43\xbd\x93\x29\xb9\x56\x1a\x24\xac\xb5\x8d\x4d\x0a\xc4\xa6\xd9\x2a\x76\xc8\x7e\xab\xb8\x9b\xa8\xf1\x34\x3a\xe9\xc0\xc4\xf8\x31\x98\x9a\xd7\x0e\xfb\xbc\xc3\xe1\xb3\x43\x69\x31\xf8\x1b\x42\xea\x95\x7b\x33\xfa\xe7\x9d\x91\x0c\x69\xc1\x70\xe2\x59\x1b\x5e\xf2\x75\x5e\xd6\x41\x8d\x8a\x41\x4c\x71\x29\x93\x2c\x6b\x52\xa5\xe2\x4b\xbe\x14\x76\xa6\xbc\xfe\x18\x94\xd1\x1b\xe8\x6b\x6a\x2e\x89\x82\x85\xcd\x1e\xe8\x35\xc0\x78\xcc\xf7\x3c\x2d\x47\x66\xa5\x61\xa8\x96\x27\x6b\xe7\x8b\xbd\x59\xa3\x08\xd7\x78\x57\x85\x12\x70\x06\x15\x56\x00\x2c\xdc\xd2\xa9\x2e\xe5\xaa\xc9\x1a\xb5\x26\xca\x53\xed\x68\x68\x72\xae\x62\xe2\xb9\x51\x51\x9e\xbc\x01\xeb\x5b\x96\xeb\x97\xeb\x73\x0f\x83\xa2\xea\x8b\x84\x1a\x19\x63\x89\x99\xf4\x1b\xb9\x5e\x39\x93\x2b\x49\x31\xde\x5e\x75\x5d\x91\x74\x66\x1b\xdd\xd4\x38\x35\xea\x5a\x32\x69\x22\x8e\x94\x8b\x38\x6e\x93\xca\x46\x5e\x41\x0e\xf4\x3b\xb8\xf9\x86\x2b\x18\xc5\xad\x17\x8f\xd5\x3f\x0c\x53\x0a\x04\x34\x37\x99\x76\x79\x7e\x59\x99\x51\x87\x49\xc6\x0b\x68\x6d\x35\x98\x71\x17\xd8\x24\xe0\x31\x6e\x6f\x43\xd8\x79\xe5\xc8\xbe\x10\x70\x2d\xa8\x43\xdc\x86\xed\x32\x8b\x32\x7b\x07\x86\xf6\x33\x7d\xfc\x2d\x54\x74\xa5\xd3\xe4\x3f\x19\x35\x79\x17\x92\x19\xbb\x50\x37\x25\xb2\xc5\xf4\x85\x2a\xc4\xbe\xe0\xb5\x5b\x9e\xbb\xe7\x99\x3b\x98\xb9\xa6\x9f\x28\xee\x98\xf4\x8c\xe0\x95\xca\xa7\x7f\x96\x66\xdf\x3c\xac\x26\xac\xcb\x5b\xeb\xf5\x7f\x52\x8b\x8f\x68\x09\xd6\xb3\x20\x0f\x4c\xbb\xfb\xd7\xe6\x8f\x7c\x5e\x7e\xa4\x1f\x07\x99\xab\x9b\x13\xcf\x4a\xf9\xf9\xb5\xe3\xf9\xd6\x67\x0d\x56\x1f\x72\x1f\xa5\xdc\xd2\x84\xbf\x94\xf0\xca\x8d\xdd\x92\xe9\x1f\x85\xe6\xf6\x0f\x4c\x83\xa0\x35\xab\x57\xea\xc6\x01\xf5\x69\xb0\xf3\xda\x94\xf7\xbe\xc5\xf4\xe1\xc5\x01\xad\x7e\x64\x29\x09\x77\x9b\x53\xcc\xdd\xfd\xce\x7f\x3e\x97\xfc\x7c\x51\xe9\x5c\xcd\xe4\x84\x3c\x31\xb5\x56\xc5\x19\xcd\xae\xcb\x0b\x4f\xe7\xa6\x22\xaf\xbd\x55\xf7\xdc\xb3\xae\xd4\xaa\xfc\xe9\x3e\xf7\x3c\xee\xf8\x80\xae\x30\xa7\xac\x69\x4a\x7a\x7d\x4c\x5d\xab\x57\x30\x38\xf7\x66\x6f\x1d\xef\xa9\xc4\xa0\xb4\x60\x25\x58\xf0\x38\x17\x07\x2e\xea\x01\x4a\xc5\xa1\x94\x02\x2d\x2e\xfe\x7e\xb8\x91\x25\x03\x9e\x3a\xb0\x76\xba\x85\xbd\x91\xfc\xd1\xa2\x43\xd8\x71\x90\x99\xc7\x0f\xb2\xe7\xce\x9c\x8f\xe8\x74\x8f\x91\x2e\x2d\x59\x87\xb1\x68\x7a\x2a\x48\x18\x36\x9f\x33\xc4\xbf\x3f\xff\x33\x73\xff\xfc\xc7\x9f\x6f\x97\xff\x2f\x00\x00\xff\xff\xdb\x34\x03\x1c\xaa\x23\x00\x00")

func clusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterroleYaml,
		"clusterrole.yaml",
	)
}

func clusterroleYaml() (*asset, error) {
	bytes, err := clusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\xbd\x0a\xc2\x40\x0c\x07\xf0\xfd\x9e\x22\x2f\xd0\x8a\x9b\xdc\xa8\x83\x7b\x41\xf7\xb4\x8d\x1a\xdb\x26\x47\x92\x13\xf4\xe9\x45\x70\x93\x3a\xff\x3f\x7e\x58\xf8\x4c\xe6\xac\x92\xc1\x7a\x1c\x5a\xac\x71\x53\xe3\x17\x06\xab\xb4\xd3\xce\x5b\xd6\xcd\x63\x9b\x26\x96\x31\xc3\x61\xae\x1e\x64\x9d\xce\xb4\x67\x19\x59\xae\x69\xa1\xc0\x11\x03\x73\x02\x10\x5c\x28\x83\x3f\x3d\x68\xc9\x68\xda\xb8\x51\x32\x9d\xa9\xa3\xcb\x27\xc7\xc2\x47\xd3\x5a\xfe\x58\x09\xe0\x87\x5a\x7b\xf6\xda\xdf\x69\x08\xcf\xa9\xf9\x8e\x4e\x4e\xb6\xd6\x7e\x07\x00\x00\xff\xff\xc4\xb6\x1b\x05\xeb\x00\x00\x00")

func clusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterrolebindingYaml,
		"clusterrolebinding.yaml",
	)
}

func clusterrolebindingYaml() (*asset, error) {
	bytes, err := clusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"clusterrole.yaml":        clusterroleYaml,
	"clusterrolebinding.yaml": clusterrolebindingYaml,
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
	"clusterrole.yaml":        {clusterroleYaml, map[string]*bintree{}},
	"clusterrolebinding.yaml": {clusterrolebindingYaml, map[string]*bintree{}},
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
