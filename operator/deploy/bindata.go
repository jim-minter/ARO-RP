// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package deploy generated by go-bindata.// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/role.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
package deploy

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xcd\x72\xdb\x38\x12\xbe\xeb\x29\xba\x3c\x07\xcf\xd4\x5a\xd2\x26\x7b\xd9\xd5\x2d\xe5\xcc\x4c\x69\x77\x36\x71\xc5\xde\x5c\x92\x1c\x5a\x60\x8b\xec\x35\x08\x70\x81\xa6\x14\x65\x6a\xde\x7d\xab\x01\x52\x7f\x26\x63\x3b\x55\x33\xc3\x8b\x2d\x80\xe8\x3f\x74\x7f\xfd\x01\x9c\x4c\xa7\xd3\x09\x36\xfc\x9e\x42\x64\xef\x16\x80\x0d\xd3\x67\x21\xa7\xbf\xe2\xec\xfe\xef\x71\xc6\x7e\xbe\x79\xb1\x22\xc1\x17\x93\x7b\x76\xc5\x02\xae\xdb\x28\xbe\x7e\x47\xd1\xb7\xc1\xd0\x6b\x5a\xb3\x63\x61\xef\x26\x35\x09\x16\x28\xb8\x98\x00\xa0\x73\x5e\x50\x87\xa3\xfe\x04\x30\xde\x49\xf0\xd6\x52\x98\x96\xe4\x66\xf7\xed\x8a\x56\x2d\xdb\x82\x42\xd2\xd0\xeb\xff\xbe\x75\xf7\xce\x6f\xdd\x0f\x13\x00\x13\x28\x49\xb8\xe3\x9a\xa2\x60\xdd\x2c\xc0\xb5\xd6\x4e\x00\x1c\xd6\xb4\x00\x63\xdb\x28\x14\xe2\x0c\x83\x9f\xf9\x86\x5c\xac\x78\x2d\x33\xf6\x93\xd8\x90\x51\xb5\x65\xf0\x6d\xb3\x80\x07\xf3\x59\x42\x67\x59\xe7\x55\x16\x96\x46\x2c\x47\xf9\xd7\xf1\xe8\x2f\x1c\x25\xcd\x34\xb6\x0d\x68\x0f\xaa\xd3\x60\x64\x57\xb6\x16\xc3\x7e\x78\x02\x10\x8d\x6f\xe8\x58\x6a\x6c\x57\xa1\x0b\x59\xa7\x37\x0a\x4a\x1b\x17\xf0\xeb\x6f\x13\x80\x0d\x5a\x2e\x92\xb7\x79\x52\xcd\x7d\x75\xb3\x7c\xff\xb7\x5b\x53\x51\x8d\x79\x10\xa0\xa0\x68\x02\x37\xe9\xbd\x5e\x38\x70\x04\xa9\x08\xf2\x9b\xb0\xf6\x21\xfd\xec\x4d\x84\x57\x37\xcb\x6e\x75\x13\x7c\x43\x41\xb8\xb7\x40\x9f\xa3\xcd\xdf\x8f\x9d\xe9\xb9\x54\x43\xf2\x3b\x50\xe8\x76\x53\x56\xd8\x6d\x1a\x15\x10\xb3\x6a\xbf\x06\xa9\x38\x42\xa0\x26\x50\x24\x97\x13\xe0\x48\x2c\xe8\x2b\xe8\xc0\xaf\xfe\x4b\x46\x66\x70\x4b\x41\x85\x40\xac\x7c\x6b\x0b\xcd\x91\x0d\x05\x81\x40\xc6\x97\x8e\xbf\xec\x25\x47\x10\x9f\x54\x5a\x14\xea\xb6\xa2\x7f\xd8\x09\x05\x87\x56\x43\xd8\xd2\x15\xa0\x2b\xa0\xc6\x1d\x04\x52\x1d\xd0\xba\x23\x69\xe9\x95\x38\x83\x7f\xfb\x40\xc0\x6e\xed\x17\x50\x89\x34\x71\x31\x9f\x97\x2c\x7d\xba\x1b\x5f\xd7\xad\x63\xd9\xcd\x53\xd2\xf2\xaa\x15\x1f\xe2\xbc\xa0\x0d\xd9\x79\xe4\x72\x8a\xc1\x54\x2c\x64\xa4\x0d\x34\xc7\x86\xa7\xc9\x70\x97\xb2\x7d\x56\x17\xdf\xed\x37\xfa\xf2\xc8\x52\xd9\x69\x42\x44\x09\xec\xca\xfd\x70\xca\xbd\xd1\xb8\x6b\x0e\xea\xee\x62\xb7\x2c\xdb\x7f\x08\xaf\x0e\x69\x54\xde\xfd\x78\x7b\x07\xbd\xd2\xb4\x05\xa7\x31\x4f\xd1\x3e\x2c\x8b\x87\xc0\x6b\xa0\xd8\xad\x29\xe4\x8d\x5b\x07\x5f\x27\x89\xe4\x8a\xc6\xb3\x93\x2e\x93\x98\xdc\x69\xd0\x63\xbb\xaa\x59\x74\xa7\xff\xd7\x52\x14\xdd\x9f\x19\x5c\xa7\xa2\x87\x15\x41\xdb\x14\x28\x54\xcc\x60\xe9\xe0\x1a\x6b\xb2\xd7\x18\xe9\x77\x0f\xbb\x46\x38\x4e\x35\xa4\x8f\x07\xfe\x18\xab\x4e\x5f\xcc\xd1\xda\x0f\xf7\x50\x32\xb8\x43\x5d\x05\xde\x36\x64\x4e\x2a\xa3\xa0\xc8\x41\xb3\x57\x50\x48\x73\xfe\x18\x5d\xc6\x6b\x51\x9f\x92\x1c\x6d\xf0\x17\x5f\x96\xec\xca\xd3\xa9\xf1\x45\xfa\xa0\x09\x6f\x14\x16\x1f\x4c\x8c\x44\xa0\x7f\x8c\x77\x6b\x2e\x07\x30\x60\xaf\x14\x45\x2b\x6c\x01\x97\x1f\xfe\x3a\xfd\xc7\xa7\xbf\xcc\xf2\x9f\xcb\xe7\x2a\xaa\xbd\x63\xf1\x3a\xf5\xf3\xf5\xed\x8f\x6e\xc3\xc1\xbb\x9a\x9c\x0c\xe9\x24\xd7\xd6\x43\xe3\x53\x78\xcd\x58\x3a\x1f\x85\x4d\xbc\x09\xbe\x18\x7c\xe7\xee\x1c\x23\x9e\x6b\xdd\x3b\x2a\x47\x82\xf1\x44\x19\x77\xe4\x70\xd8\xb3\xaf\x0a\x48\x5d\xa9\x41\xf3\xdc\x5d\x1c\xcc\x5c\xd8\x83\x23\xc9\x75\x45\xe6\x9e\xc2\x73\xf2\xc9\xa4\x36\x7f\xcb\x32\x34\x09\xc0\x42\xf5\xe0\xc4\x23\x2e\xf6\xd3\x18\x02\xee\x9e\xea\x46\x8f\x6d\xcb\xe2\x5c\xe3\x49\x39\xf6\x9c\x64\xf9\xba\xef\x89\xaf\xbe\xb4\x81\x8e\x96\xe7\x16\x45\x47\x8d\xfa\x49\x86\x47\x0a\x1b\x36\x74\x13\xd8\x19\x6e\xd0\x3e\x27\x8e\x35\x26\x88\x68\x57\x8e\xe4\xa1\xfd\x8f\xc6\x6b\xeb\xc3\xfd\x61\xf9\x9f\xbd\x15\xc3\x10\x99\x89\xcc\x63\x20\x99\xde\x3a\x81\x49\xbf\xd2\xc0\x7e\x1b\x4e\x1a\xef\x0a\x3e\xe2\x97\x63\xca\xf7\xaf\x75\x8d\x94\x24\xe9\xe9\x87\x81\x5d\x14\x74\x86\xe2\x6c\xf2\xa4\xa8\x9e\x48\xbf\x38\xc8\x39\x74\xd7\x4c\x70\xd4\xb3\x44\x7f\x4e\x28\xcf\x65\xcc\xbe\x9e\x2b\xd3\xe7\xc8\x54\x0c\xa4\x6b\xf6\x4c\x1c\x6a\x32\x15\x3a\x8e\x75\xaa\x68\x57\x50\xa1\x8c\x48\x3b\x6d\xa4\x02\xb6\x15\x39\x0d\xe8\x80\xd0\x82\x04\xd9\xc6\xbd\x11\x07\xb3\x54\x87\xb6\x6b\x84\x26\xb0\x0f\x0c\x89\x77\x83\x0f\xb0\x4d\x64\x2c\xcd\x35\x8d\x3d\xcf\x8c\x94\x06\x1e\xd0\xda\x43\xec\x92\x78\x28\x79\x43\x0e\x94\xb4\xcc\xe0\xa3\x3b\xf6\xa7\xe3\x77\x2b\x02\x2c\x0a\x1a\x02\x6d\xf1\x40\x9f\x1b\xcb\x86\xc5\xee\x32\x11\xdc\x1d\xed\x3d\x48\x85\xa2\xce\x86\x98\x08\x9e\xf1\x75\xe3\x5d\x8a\xb6\x49\xc1\x5a\xf9\x76\x08\xf2\x03\x4a\x95\xc8\x0d\xba\xc4\x55\x38\x64\xce\xe4\x23\x9d\x48\x4f\xb1\x4c\x44\x48\xdb\x76\xa2\x41\x5e\x57\x0e\x88\x3c\x8a\x61\x9c\xc1\x5b\x67\xa8\xcb\xe9\xe2\x2a\x25\x75\x4d\xe8\x54\x49\x0a\xc9\x21\x3f\x0c\x3a\xc8\xec\x68\x40\xa6\x6e\x6e\x49\x05\x60\x58\xb1\x04\x0c\x6c\x77\x30\x05\xd6\xb7\x8d\xaf\x29\x42\x83\x41\x7a\xec\x7a\x75\xb3\xcc\x2c\xb7\xc2\x5c\x46\x11\xeb\x21\xa1\x2b\x34\xf7\x5b\x0c\x45\x9c\xa6\xb7\xd7\x3e\xe4\x5f\x1a\x3b\x14\x5e\xb1\x65\x49\xa1\x36\x14\x5c\x97\x21\xbb\xec\x76\xd2\x37\xe4\xfb\xde\x82\xd9\xc5\x83\xe9\xaf\x81\x20\x80\xc5\x28\x77\x01\x5d\xe4\xfe\x48\x37\x8c\x58\x6b\x1f\x6a\x94\x05\x28\x81\x9c\x0a\x0f\x7a\xf6\x28\xae\xd5\x14\x23\x96\x23\x1a\x1e\x59\x1b\x08\xe3\x70\xef\x1f\x83\x96\x77\x69\x85\xe2\xcb\x59\x71\x22\x78\x47\xd3\xad\x0f\xc5\xd5\x81\x06\x0f\x0a\x86\xb3\x33\xd3\xbe\x4f\xa1\x50\xe9\xc3\x4e\x7f\x1b\x6c\x23\xed\x27\xda\x10\xc8\x49\x87\xbd\x43\x70\xa2\xcf\x52\x06\xac\x4a\x90\xc1\x2e\xed\x3c\xab\xc4\x56\x9a\x56\xae\x20\xb6\xa6\x02\x8c\xc9\x66\xcb\x6e\xcc\x50\x3d\xb9\x1b\xb1\x50\x2a\x92\x76\x4b\x35\xbf\xd8\x41\x6c\xeb\x1a\x03\x7f\x49\xe9\x6f\xb2\x89\x1d\x3a\x24\xe3\x47\xec\x7c\x64\x43\x1e\xb6\x97\x27\x2f\x4d\xd3\x8f\xef\xe4\x01\xc6\xef\x76\x0d\xf5\xdc\x41\x17\xef\xc3\xbd\xaf\xe3\xe4\x6a\x1c\x89\x8d\xec\x1a\x36\x68\xed\x4e\x4b\xbf\xdf\xf0\x42\x7b\x78\xa1\xc0\x1a\x2b\x1f\x04\x9a\x2a\xa4\xd3\xd0\x31\x44\x26\x65\x63\x52\x3b\xf4\x64\x57\xb0\xe6\x43\xd7\x2d\x39\x41\x3e\x7c\xbc\xc0\x95\xd3\x9a\xb1\x53\x09\x2d\x7d\xbc\x80\xc6\x5b\x0c\x2c\xbb\x19\xfc\xe4\x87\x00\x4c\x1f\xfa\x8c\x75\x63\xe9\x0a\xf8\xdc\xbf\x5e\x4b\xcc\x5d\x05\x55\x1c\x9b\x5d\xce\xa3\x74\x4b\x71\x35\xe6\x7c\xb2\x86\x63\xbe\xcb\xf8\x78\x01\x06\x63\x0a\x66\x13\xfc\x0a\x57\x76\x97\xde\x50\x5b\xaf\x20\xfa\x53\xb5\x5f\xf7\x7c\xa5\x85\x60\x2d\x15\xf0\xf1\x62\xe9\x3a\xf1\x03\x08\x04\x8f\x65\x44\x6e\x01\x34\xc0\xc2\xa6\x5d\x9a\x0d\x4c\xa8\xc4\x07\xc3\xa3\x54\x75\x9c\x54\x29\x3a\xa2\xf8\x30\x72\xd4\x1a\x35\x3c\x90\xd5\xc3\xf4\xdb\xa4\xeb\x41\x1d\x3c\x85\xa6\xe4\xa5\xef\x68\x4d\xa9\x20\xd3\x5d\x1c\xb2\x8b\x40\xce\xb7\x65\x95\x8e\xe4\x8a\xba\x29\x11\x3d\x58\x12\xd8\xf9\x76\x88\x69\x3a\x3d\x0e\x8b\xe6\x72\xed\x0b\x5e\xe7\x2d\x0d\xd4\xf5\xcd\xee\x5a\xe7\x99\x9d\x61\xf8\x06\x6a\xc4\x95\x57\x37\xcb\xfe\xde\xa9\xaf\xcd\x90\xfd\x1a\xd0\xfb\xd5\xb0\xe6\x67\xcd\x64\x8b\x1b\x94\xea\x09\xba\x2f\x97\xeb\xce\xd7\x44\x22\xbc\x16\x07\x93\xa1\x13\x86\x97\x48\x11\xa1\x9e\x34\x46\x92\x5a\x89\x9d\x13\x0e\xd4\xad\xb8\xca\x77\x2f\xdd\x15\xcf\xe1\x2a\x4c\xb7\x08\x30\x57\x13\xfc\xf3\xf6\xed\x9b\xf9\xcf\x7e\x44\x64\xf2\x02\xd0\x18\x8a\x1d\xc1\xd4\x83\xf5\x01\xd2\xbb\xfb\x88\xdb\x44\x3d\x6b\x74\xbc\xa6\x28\xb3\x4e\x07\x85\xf8\xe1\xe5\xa7\xb1\x16\xf2\x93\x0f\x0f\xd0\x62\x7f\xa1\xd4\x27\x14\xc7\x1c\x8e\xbd\x44\xd8\xb2\x54\x3c\x56\xd6\x0a\x2a\x45\xe7\x76\x26\x9b\x82\xf7\x04\xbe\x73\xb7\x25\xb0\x7c\x4f\x0b\xb8\xd0\x6c\x3b\x32\xf3\x57\x3d\x1f\xff\x36\x5c\xf7\x00\xdf\x6f\x2b\x0a\x04\x17\xfa\xd2\x45\x36\x6e\x7f\x6f\xa8\x63\x47\x58\xde\x19\x99\x48\xa5\x04\x2e\x4b\x0a\x83\xac\x14\x3a\x58\xa3\x0d\x39\xf9\x41\xd3\x9e\xd7\xe0\xfc\x91\x88\x24\x58\x77\xaf\x21\xc3\x6b\xa6\xe2\x81\xd1\x1f\x5e\x7e\x1a\xb5\xf8\x34\x5e\x8a\xbd\xf4\x19\x5e\x66\x68\x57\xe0\xf4\xc5\x0f\x33\xb8\x4b\xd9\xb1\x73\x82\x9f\x55\x93\x51\xf2\x3a\x16\x59\xef\x14\x66\x3d\x54\xb8\x21\x88\xbe\x26\xd8\x92\xb5\xd3\x8e\xa0\xc2\x16\x13\x93\xe8\x37\x4e\xf3\x0d\x7b\x6e\x39\x9e\xad\xfd\x6d\xed\xdd\xdb\xd7\x6f\x17\xd9\x32\x4d\xa8\x32\x51\x1e\xe5\xb4\x6b\x76\x68\x53\x67\xcc\xb7\x88\x29\x1b\x47\x9b\x64\x6c\x73\xfa\x88\xef\x78\x6f\xdf\xca\xd6\xad\xb4\x81\x66\x43\x17\x4b\x8f\xd6\xf1\xf9\x45\xea\xe1\x19\xb8\x52\x3d\x07\x8e\x3f\xe9\x62\xf2\xc9\xce\xb9\x91\x7b\xbd\x73\xe7\xde\x1c\x65\xf9\x57\x9d\x53\x0e\x17\x1c\x09\x25\xff\x0a\x6f\xe2\x3c\x9d\x03\x1a\x89\x73\xbf\xa1\xb0\x61\xda\xce\xb7\x3e\xdc\xb3\x2b\xa7\x9a\x9a\xd3\x9c\x03\x71\x9e\x6e\xa8\xe6\xdf\xa5\x3f\xdf\xec\xcb\xe8\x15\xd7\x90\x43\xe9\xe5\x3f\xc2\x2b\xd5\x13\xe7\xdf\xe4\x54\x7f\xc1\xf4\xf4\x3e\x76\x79\x9b\x01\xc3\x9c\xaf\xd5\xb2\xd8\x56\x6c\xaa\xfe\x53\x4a\x87\xb1\x23\xc5\xc4\x11\x6a\x2c\x32\x34\xa3\xdb\xfd\xee\xa9\xac\x01\xcd\xbc\x7e\x37\xed\x3e\xe9\x4d\xd1\x15\xfa\x7f\xe4\x28\x3a\xfe\x4d\x11\x6c\xf9\x49\xe5\xfb\x9f\xe5\xeb\x3f\x26\xc1\x5b\xfe\xa6\x5a\x7d\x36\x2d\x1c\x58\x70\x36\xb4\xff\x34\xba\x79\x81\xb6\xa9\xf0\xc5\x61\x2c\xd1\xa9\x69\xf7\x35\xf4\x68\x3a\x5f\x58\x52\xb1\x00\x3d\x12\xe4\x01\xf1\x41\x4f\xc4\x79\xe4\x70\xa4\x52\xce\xd0\x08\x15\x6f\xce\xbf\x87\x5e\xe4\xa6\xd5\x7f\xf0\x4c\x3f\x8f\xee\xdc\xe0\xc3\xa7\x49\x96\x4a\xc5\xfb\xde\x1a\x1d\xfc\x7f\x00\x00\x00\xff\xff\x86\xd0\x68\xea\x52\x1e\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\x31\x4f\x3b\x31\x0c\xc5\xf7\x7c\x8a\xa8\x7b\xae\xfa\x6f\x7f\xdd\xca\xc0\x8e\x10\xbb\x9b\x73\x5b\xab\x49\x1c\xd9\x4e\x81\x7e\x7a\x94\x2b\x15\xa8\x05\xd1\x4a\x0c\x4c\x67\x3f\x9d\xdf\xef\x39\x89\x0b\x21\x38\xa8\xf4\x84\xa2\xc4\x65\xf4\xb2\x82\x38\x40\xb3\x2d\x0b\x1d\xc0\x88\xcb\xb0\xfb\xaf\x03\xf1\x72\xff\xcf\xed\xa8\x4c\xa3\xbf\x4b\x4d\x0d\xe5\x81\x13\xba\x8c\x06\x13\x18\x8c\xce\xfb\x28\x38\x0f\x3c\x52\x46\x35\xc8\x75\xf4\xa5\xa5\xe4\xbc\x2f\x90\x71\xf4\x20\x1c\xb8\xa2\x80\xb1\x84\x0c\xdd\xc3\x49\x4b\xa8\xa3\x0b\x1e\x2a\xdd\x0b\xb7\xaa\xdd\x29\xf8\xc5\xc2\x79\x2f\xa8\xdc\x24\xe2\xbb\x16\xb9\xac\x69\x93\xa1\xea\xdc\x76\x53\xad\x10\xf1\xd8\x2a\xca\x9e\x22\x42\x8c\xdc\x8a\x75\x6d\x8f\xb2\x3a\x8d\xf6\x68\x38\x97\x1b\xb4\xf9\xdb\xea\xd4\xa5\xab\xc8\x8a\x51\xf0\x0a\xcf\x44\x7a\x2c\x2a\x58\xdc\x7e\xc6\xf4\xf2\x79\x16\x2f\x80\x50\xe7\x85\xce\x90\x13\x60\xe6\xa2\x37\x51\xaf\x60\x09\x0f\x5c\xb1\xe8\x96\xd6\x36\x10\x7f\x71\xc8\xc7\xcb\xfd\x9e\x3a\x61\xc2\x5f\x59\xfb\x96\x28\x1f\xcd\x72\x4d\x05\x12\x1d\xfe\x56\xc4\xa5\x1a\x58\x3b\x4b\x74\x62\x5f\x20\x2f\x40\x8a\xb1\x09\xd9\xeb\x0f\xb4\xd3\x6f\x91\x8b\xe1\x8b\x45\x2e\x6a\x02\x74\xc3\x6b\x7f\x0b\x00\x00\xff\xff\xb6\x8c\xf2\xc1\xef\x03\x00\x00")

func masterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRoleYaml,
		"master/role.yaml",
	)
}

func masterRoleYaml() (*asset, error) {
	bytes, err := masterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\xb1\x4e\xc5\x30\x0c\x45\x77\x7f\x85\x7f\x20\x45\x6c\x28\x1b\x30\xb0\x3f\x24\x76\xbf\xd4\xa5\xa6\xad\x1d\x39\x4e\x87\x7e\x3d\xaa\x40\x2c\x48\x9d\xef\x39\xf7\x40\x4a\x09\xa8\xca\x07\x7b\x13\xd3\x8c\x7e\xa7\x32\x50\x8f\xd9\x5c\x0e\x0a\x31\x1d\x96\xa7\x36\x88\x3d\xec\x8f\xb0\x88\x8e\x19\x5f\xd7\xde\x82\xfd\x66\x2b\xbf\x88\x8e\xa2\x9f\xb0\x71\xd0\x48\x41\x19\x10\x95\x36\xce\x48\x6e\xc9\x2a\x3b\x85\x79\xda\xe8\x14\xc0\x6d\xe5\x1b\x4f\x27\x44\x55\xde\xdc\x7a\xbd\x08\x02\xe2\xbf\xde\xe5\x7d\xeb\xf7\x2f\x2e\xd1\x32\xa4\x5f\xf3\x9d\x7d\x97\xc2\xcf\xa5\x58\xd7\xb8\x94\x7f\xb6\x56\xa9\x70\x46\xab\xac\x6d\x96\x29\x12\x1d\xdd\xf9\x0f\x86\xef\x00\x00\x00\xff\xff\xea\x5c\x27\x5f\x2f\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\xde\xa7\xf0\x05\x5c\x6c\x9b\x6e\xcf\x80\x44\xff\x95\xf9\x88\x08\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x27\x66\x26\x58\xe3\xce\xdc\x23\xbc\xe9\xf5\x27\xaf\xe1\x47\xd3\x1b\xf3\x1a\x9d\xff\xbd\xc7\xe9\x25\x93\x85\x03\x85\x26\xaa\x8e\xc9\xa6\xc8\xb0\x58\x4c\x54\xa4\x4d\xec\x62\x7e\x6d\x2f\x74\x36\x8d\x45\xdf\xcf\xf1\x28\xc3\xfb\x4c\xfe\xb2\x7c\x02\x00\x00\xff\xff\x5b\x98\x41\x31\x75\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xca\xb1\x0d\xc2\x40\x0c\x05\xd0\xde\x53\x78\x01\x17\xb4\x37\x04\x25\xfd\x27\xf9\x11\x56\x88\xef\x64\x5f\x28\x32\x3d\x4a\xf9\xa4\x27\x66\x26\x18\xfe\x62\x96\xf7\x68\xfa\x7b\xc8\xee\xb1\x36\x7d\xe2\x60\x0d\x2c\x94\x83\x13\x2b\x26\x9a\xa8\x06\x0e\x36\xed\x83\x51\x1f\xdf\xa6\xe1\x3a\x93\xd6\x07\x13\xb3\xa7\xd4\xe0\x72\xb7\xcd\x03\x5f\xbf\x98\x75\xcb\x74\x3f\xdf\xcc\xe0\x64\xc9\x3f\x00\x00\xff\xff\x44\x6f\xf6\xda\x72\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\x31\x4e\x04\x31\x0c\x45\x7b\x9f\xc2\x17\xc8\xac\xe8\x50\x5a\x0a\x7a\x84\xe8\xbd\x33\x86\xb5\x26\x13\x47\xb6\xb3\x2b\x71\x7a\x34\xd9\x9d\x06\x4a\xaa\x3c\x59\xff\xe7\x3f\x48\x29\x01\x35\xf9\x60\x73\xd1\x9a\xd1\xce\x34\x4f\xd4\xe3\xa2\x26\xdf\x14\xa2\x75\x5a\x9f\x7d\x12\x3d\x5d\x9f\x60\x95\xba\x64\x7c\x29\xdd\x83\xed\x4d\x0b\xc3\xc6\x41\x0b\x05\x65\x40\x9c\x8d\x47\xe1\x5d\x36\xf6\xa0\xad\x65\xac\xbd\x14\x40\xac\xb4\x71\x46\x32\x4d\xda\xd8\x28\xd4\xd2\x4d\x6d\x65\x03\xeb\x85\x3d\x43\x42\x6a\xf2\x6a\xda\x9b\xef\x3f\xa5\x3d\x3b\x69\xe3\xea\x17\xf9\x8c\x49\x14\x10\x8d\x5d\xbb\xcd\xfc\x48\xcc\x77\x0b\x07\xc4\x2b\xdb\xf9\xb8\xee\x0e\x3c\x70\xe1\xc2\x0f\xfc\xe2\x18\x6f\x11\xbf\x43\xa3\x98\x2f\x83\x7a\x5b\x8e\xc2\x6d\x1c\xff\xa1\x72\xf2\xa0\xe8\xbf\x8c\x8e\xed\x3f\x93\x3f\x01\x00\x00\xff\xff\x32\xe1\x82\x0f\x7b\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x4e\xc6\x30\x0c\x46\x77\x9f\xc2\x17\x48\x11\x1b\xca\x06\x0c\xec\x3f\x12\xbb\x9b\xba\xd4\xb4\xb5\x23\xc7\x29\x52\x4f\x8f\x10\x88\x05\xa9\xf3\xf7\xde\xf7\x20\xa5\x04\x54\xe5\x8d\xbd\x89\x69\x46\x1f\xa9\x0c\xd4\x63\x31\x97\x93\x42\x4c\x87\xf5\xa1\x0d\x62\x77\xc7\x3d\xac\xa2\x53\xc6\xe7\xad\xb7\x60\xbf\xd9\xc6\x4f\xa2\x93\xe8\x3b\xec\x1c\x34\x51\x50\x06\x44\xa5\x9d\x33\x92\x5b\xb2\xca\x4e\x61\x9e\x3e\xcd\x57\x76\x70\xdb\xf8\xc6\xf3\x37\x44\x55\x5e\xdc\x7a\xbd\x08\x02\xe2\xbf\xde\xe5\x7d\xeb\xe3\x07\x97\x68\x19\xd2\xaf\xf9\xca\x7e\x48\xe1\xc7\x52\xac\x6b\x5c\xca\x3f\x5b\xab\x54\x38\xa3\x55\xd6\xb6\xc8\x1c\x89\xce\xee\xfc\x07\xc3\x57\x00\x00\x00\xff\xff\x21\x49\xf8\xf0\x2f\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\xaa\x03\x41\x08\x06\xe0\xde\x53\x78\x01\x8b\xd7\xda\xbd\x33\x04\xd2\xcb\xec\x1f\x22\xcb\xea\xe0\xb8\x1b\xc8\xe9\xd3\x84\xd4\xdf\x47\x22\x42\x36\xfd\x8e\x5a\x9e\xa1\x7c\xfd\xd1\xee\xb1\x29\xdf\x50\x97\x0f\xfc\x8f\x91\x67\x34\x1d\x68\xdb\xac\x4d\x89\x39\xec\x80\xb2\x55\x4a\x4e\x94\x75\x96\xbc\xb2\x76\xd4\xd7\xd6\xb4\x01\xe5\x9c\x88\xf5\xf4\x47\x8b\xbd\xcf\xc2\x2f\xd3\x27\x00\x00\xff\xff\x5c\x51\x06\x72\x75\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aro.openshift.io_clusters.yaml": aroOpenshiftIo_clustersYaml,
	"master/role.yaml":               masterRoleYaml,
	"master/rolebinding.yaml":        masterRolebindingYaml,
	"master/serviceaccount.yaml":     masterServiceaccountYaml,
	"namespace.yaml":                 namespaceYaml,
	"worker/role.yaml":               workerRoleYaml,
	"worker/rolebinding.yaml":        workerRolebindingYaml,
	"worker/serviceaccount.yaml":     workerServiceaccountYaml,
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
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"role.yaml":           {masterRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
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
