// Code generated for package noauth by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package noauth

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

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\x0a\x02\x31\x10\x45\xfb\x39\xc5\x4f\xa7\x57\x48\x67\x23\x58\x28\x88\xa5\x58\x0c\xeb\xec\x1a\xd8\x24\x4b\x66\x14\x17\xd9\xbb\x0b\x81\x88\x62\x37\xbc\x79\xbc\x2f\x4f\x93\x74\x85\xcd\x93\xe0\x78\x97\x32\xe3\x45\x00\x17\x0b\x3d\x77\xa6\xab\x76\x1d\x38\x8a\xc7\xc9\x4a\x48\x83\x5b\x7b\x6c\x9a\xb1\x4b\x7d\x76\xb4\x10\xd5\xc4\x0f\xae\xa9\x60\x12\xd5\xe3\xdc\x3e\xee\xf2\x6f\x57\xf1\x21\x45\x43\x4e\x9f\x11\x02\xba\x1b\xa7\x41\xc6\x3c\x7c\x43\x0b\x51\xd4\x38\x4e\x7b\xf5\xd8\x8e\x99\xcd\xd1\xf2\x0e\x00\x00\xff\xff\xc2\xab\x64\xad\xc7\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 199, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xdf\x6f\xe3\xb8\x11\x7e\xf7\x5f\x31\x7b\xfb\x70\x09\x10\x2c\x0e\x45\xef\x50\xf8\xa9\x3a\xdb\x7b\xab\x26\x71\xdc\xd8\xd9\xed\xa1\x58\x2c\x68\x71\x6c\x11\x91\x48\x1d\x49\x39\x71\x8b\xfd\xdf\x8b\x21\x29\x89\xb4\x95\xdb\xe6\xda\x37\x8b\x3f\xbe\xf9\x66\x38\xfc\x38\xa4\xf1\xd9\xa2\xe4\x60\x8f\x0d\xc2\xdf\x5b\xd4\x47\xf8\xf7\x04\xa0\x35\xa8\xa7\xf0\x60\x50\xe7\x72\xa7\xde\x4c\x00\x94\xde\x4f\xe1\x4e\xef\xbb\x6f\x1a\xb1\x46\x6b\x85\xdc\x1b\x3f\xb2\xfb\xea\x7a\x33\x6b\xb5\xd8\xb6\x16\x43\xff\xf0\x1d\xf0\xa8\xd1\x4c\xe1\x9f\xbd\x99\xcf\xd4\x51\x54\xad\xb1\xa8\x2f\x04\x9f\x42\x3e\x7f\x73\x39\x85\x99\x6f\xe9\x2c\x87\x01\x3f\x1f\x97\xac\xc6\x0b\xc9\x6a\x9c\xc2\xda\x6a\x21\xf7\x2f\x0f\x26\x33\x71\x4f\x6c\x69\xa6\xa4\xc4\xc2\x0a\x25\xcf\x6d\x0e\x7d\x03\xa0\xc8\xb4\x15\x3b\x56\xd8\x0b\x16\x7e\x6c\x8e\x0d\x4e\x21\x8b\xbe\x1c\xc4\x4d\xde\x35\xd1\x44\xd6\x5a\x55\xa8\xba\xa9\xd0\xe2\x85\x90\x4d\x6b\x3b\xda\x57\x50\xb4\xda\x28\xbd\x52\x66\x0a\xb9\xb4\x57\xc0\x9c\xc9\x29\x64\xd1\x9c\xcc\xb5\x11\xf8\x55\xc7\xfc\x21\x9f\x77\x18\x97\xe9\xe0\x7b\x34\x6d\x75\x66\xf6\xbd\xc0\x8a\x9f\xda\xde\x51\x63\xf0\x20\x1a\xbb\x90\x56\xd8\xe3\xb5\x90\xfc\x6a\x02\x00\xa0\xf1\xb7\x56\x68\xe4\x99\xde\xd3\x60\x0a\xe8\xf8\xf0\xcf\x2f\xd0\x4b\x86\xaf\xdb\xfd\x1e\x0d\x39\xf4\x79\x32\x01\x78\x0b\xeb\x42\x8b\xc6\xd6\x7b\x0d\x28\x79\xa3\x84\xb4\xe6\x0a\x34\xee\x50\x83\x55\xc0\x55\x61\x40\x48\x28\x2a\xd5\x72\xd6\x88\x77\x8d\x56\x56\x4d\x00\x2a\x71\xc0\x8f\x02\x9f\x88\xce\x4d\xf8\x7d\x8b\x96\x71\x66\x99\x5f\xe4\x6e\xc4\x4c\x49\x8b\xd2\x9a\x68\x8d\x6f\x4e\xba\x68\xb8\x71\x3c\x08\xce\x33\x4a\xc1\x7c\xef\x08\xd4\x3a\xe9\x78\xe3\x7d\x9a\x63\x53\xa9\x23\x3c\xe2\xd1\x4c\x00\xb8\xfb\xaa\x51\xda\x6b\x3c\x92\x81\x79\xdc\x90\xda\x49\xc6\x46\x66\x92\x29\xc1\x4a\xb6\xca\x3b\x13\xac\x11\x01\x3b\x5b\xe5\x67\xa0\xbe\x37\x42\xf3\x83\xde\x4c\xbe\x4e\x26\xb1\x0a\xdc\xb6\x96\xd1\xca\x38\x21\x98\x69\x64\x16\xc3\x6e\x48\x76\x17\xfc\x95\x63\xa3\xb1\x60\x16\xf9\x85\x46\x66\x28\x61\xbf\x0b\x03\x0c\x30\x8d\x20\xd5\x13\x14\x0e\x80\xc3\x41\x30\x68\x9e\x83\x67\xdf\x5d\x4e\x00\x1e\x1a\xce\x2c\x7e\x14\xff\x12\x6e\x9f\xed\xc4\xfe\x22\x24\x0e\xe5\x4d\x3e\x7f\x73\x05\x87\xa8\x73\x0a\x0b\x2e\x2c\xdb\x56\xc9\x94\x91\x2d\xef\x29\x27\xa1\x3a\x8b\x1c\xc0\x1c\x29\x0f\xe7\x2f\x04\xfa\x67\xa5\x2a\x64\x72\x80\xf3\xb1\x1a\x62\xd6\x01\xf8\xef\xf1\x99\xde\xc1\x58\x1a\x2f\x4c\xaf\x98\x9d\x33\x89\x72\x5e\x9e\x2b\xe9\x1a\x6d\x2a\x9e\x17\x2c\xd2\xd5\x18\x25\xd2\xd7\xcb\x31\xc5\xcd\xe5\x41\x78\x3a\x17\x58\x33\x51\xf5\xaa\x49\x1a\xa0\x8d\x5d\xc6\x4a\x7a\x05\x15\x3b\x69\xba\xec\x0e\x04\x82\x49\xfd\x5b\xa1\xae\x85\x31\x42\x49\x73\x41\xd2\xdf\x2f\x60\x9b\x76\xa6\x84\xa3\x8e\x01\xdc\xaf\xa1\x87\xbe\xd3\xfb\x3e\x72\x4a\xef\x7b\x54\x35\xb4\x0f\x88\xd1\x60\x42\xeb\x8f\xaa\xaf\x93\x89\x4b\xeb\x0e\xde\xa5\x75\x58\xaf\x09\x40\x72\x7e\x4c\x00\xd2\xd0\x4c\x00\x1a\x51\xd8\x56\x27\x63\x94\xde\x2f\x4f\xa6\x05\x7a\x43\x83\x30\x59\xd3\x68\x75\x40\x1e\xe5\x44\xc7\x25\x90\xfb\x16\x15\x49\x7e\x79\x18\x56\x99\x11\x9c\x38\x59\x1c\x18\x93\xac\x3a\x5a\x51\x98\xbb\xc6\x2a\x52\xf9\x28\x1d\x3b\x43\xf1\xe4\x21\x43\xdc\x74\xab\x5a\xbd\x46\x94\x2f\xcd\x73\x47\xc7\x0b\x49\x37\x0e\x30\x3e\xeb\xbf\xe2\xdc\x13\x4d\xc5\xec\x24\x68\x41\x5f\x32\x7b\x6b\xa6\xf0\xbe\x52\xcc\x7a\x01\x35\xc5\x10\xca\x14\xe8\x04\xe0\x91\xb6\xf5\x10\xf4\xd7\xe0\x8d\x2a\xf8\xff\xc0\x2f\xc1\xfb\xbf\xd0\x44\xd9\xd6\x23\xc7\xfa\xda\x32\x8b\xce\x40\xb6\x58\x7f\x79\x58\x5e\x2f\xef\x3e\x2d\xc3\xd7\x6a\xb1\x9c\xe7\xcb\x5f\xc2\xd7\xfd\xc3\x72\x39\x7c\xbd\xcf\xf2\x9b\xc5\x3c\x7c\x6c\x16\xf7\xb7\xf9\x32\xdb\x2c\xe6\xa3\x96\x86\x7a\xc5\x1b\xca\x36\x91\xa1\xb7\x90\x49\x40\x2e\x6c\x28\x75\x40\x15\x54\x03\x81\xd8\x01\x73\xa2\x01\x25\x33\x50\x2b\x2e\x76\x02\x39\xd8\x12\xc1\x67\x91\xc5\x67\x0b\xdb\x23\x08\x69\x50\x53\x0e\x81\xd2\xc0\x49\x8a\xe9\x77\x51\x32\xcd\x0a\x3a\x7f\xde\x39\x23\x9b\x52\x50\xdd\x50\x54\x2d\x47\x43\xa7\x9b\x9b\x20\x1d\xde\x23\x1e\xb7\x8a\x69\x0e\x4c\x72\x68\x98\xf1\x00\xaa\xae\x99\xe4\x6e\x3a\x31\x5e\xcc\xf3\x8d\xa7\x0b\x06\x2b\x2c\x06\xbe\xb2\x3a\x8e\x93\x2e\x4a\x65\x50\x02\x93\x49\xe9\x05\xa6\xaf\x78\xde\x75\xb4\xb8\xa0\xc3\xd3\x80\xab\x64\xde\x3a\x52\xc9\x14\x5b\x32\x0b\xc2\x82\x29\x55\x5b\x71\xa8\xd5\x01\xdd\x20\x32\xf5\xbd\x09\x45\x23\x95\x47\xd4\x28\x29\x30\x8c\xb4\xa2\xd1\x82\x56\xd7\xb2\x6d\xe7\xc5\x7a\x71\xb3\x98\x6d\x7e\x27\x1f\xa8\x6e\x0b\xe9\x70\x9d\xa4\xc3\xf5\x97\xd5\xdd\x3c\xfc\x5a\x7f\x9c\x75\xbf\x66\xf7\xf9\x6a\x13\x3e\x96\xd9\xed\x62\xbd\xca\x66\x8b\x61\x9b\x8d\x16\x7a\x0e\xff\x51\x48\xfe\x52\x9d\x79\xa2\x80\x21\x9d\xa9\xae\x72\xb5\x70\xdf\x5a\x33\x5b\x94\xc8\x73\xc9\xf1\xd9\xd5\xa1\xb9\xb4\x9f\xa9\x38\xa3\xa4\x1e\x03\x77\xd9\xde\xb3\xdb\xb0\xed\x09\x29\xca\x13\xca\x2f\x8e\xcf\xa0\x76\x2e\x9a\x96\x6d\x7d\xf8\x6d\x89\x26\x5e\x3c\x5f\xd8\xec\x94\xa6\xd8\x5a\xb6\x75\x2c\x5c\xd5\xee\x80\x3e\x95\x68\x4b\xd4\x21\x59\x28\xa3\x58\x34\x99\xe6\x81\xa5\xc5\x27\x7c\x6f\xf0\x49\x54\x15\xd4\xec\xd1\x2f\x6d\xc8\x3f\xc0\x67\x2c\x5a\x27\x97\x64\x67\xf8\xca\x76\x96\xd4\x93\xc0\x07\x9d\x84\x98\xdf\xef\x14\xda\x63\xeb\xe3\x2f\x0a\x51\x18\x76\x4a\xd7\xcc\x52\xc5\xe6\x37\x1c\x91\xed\x77\x9f\x09\x77\x86\xa7\x52\x14\xa5\xcb\xf6\x2d\xa2\x84\x86\x69\x83\x9c\xb6\xe5\x79\x0e\xab\x3e\xd1\x7d\x92\xb3\xed\xda\xaa\x06\x1a\x65\x84\xe3\x4b\xfe\xf5\x36\xf3\xf8\x6a\x92\x04\xf4\x94\x03\xf1\x62\x70\x60\x95\xe0\x57\x51\x7c\xba\x00\xbe\x73\x27\xf0\xa2\x6f\x8f\x83\xf5\x16\xb2\xaa\x4a\x96\x94\x96\x05\x59\x51\x46\xab\x4f\x24\x4d\x58\xe3\x75\x12\xdd\x24\x7f\x86\xa0\x52\xed\xcf\x84\x44\x4d\xd9\xd6\xfa\x93\xed\xf4\x40\x1f\x17\xed\x90\xb7\xc3\xb0\x1a\x8d\x61\xfb\xa4\xa9\xab\xaf\xe3\x16\x63\x99\xb6\x33\xd5\x4a\xeb\xf2\x6f\x38\x46\xae\xff\x62\x16\x07\x94\x7e\x55\x47\xc0\x5c\xb5\xb7\x11\x35\x26\x34\xa8\xde\x3b\x69\xec\x00\x57\x8a\xff\x21\xaf\x5a\xf3\x6a\xb7\x8a\x2e\x8c\xee\xc6\x9e\xc6\xd4\x5f\x63\x90\x5c\xa3\xde\xce\x4d\xdf\x3c\x1e\x0f\xa7\x77\xe1\x7e\x10\xb9\xe0\x53\x9d\xe3\x8e\x51\xf2\xbb\x05\x20\x15\x97\xca\x96\x21\xb7\x1e\xa5\x7a\x92\xb4\xfe\xb3\x75\x72\x6c\xd1\xbc\x30\xde\x40\x89\xac\xb2\xe5\x91\xa6\x96\xc8\xb4\xdd\x22\xb3\x5e\x20\x34\x16\x28\x0e\xc8\xe9\xb0\xd1\xb8\x6f\x2b\xa6\x41\x48\x8b\x9a\x0a\x39\x77\xe2\xd8\xd2\x6f\x88\x70\xe5\x21\x38\x8d\xa6\x51\x92\x13\x03\xab\xdc\x85\x1b\x8d\x35\x81\xc4\x87\x45\x76\xb3\xf9\xf0\xeb\x39\x89\x56\x46\x34\x9c\x86\x0c\x88\x85\x7f\xbe\xa0\x13\x54\xc1\x4a\x3c\x0b\x84\x19\x5d\xa1\x1d\x03\x61\x80\x2a\x4b\xc1\xbb\xbd\x36\xf8\x70\x05\x5b\xb7\xf5\xe5\xf7\x16\x7e\x6b\x51\x1f\xdd\xde\xa2\x6d\x62\x54\x8d\x61\xd9\xc2\x39\xa6\xd1\x60\xbd\xad\xd0\xc0\x87\xcd\x66\xf5\xbd\x81\x1f\x7f\xf8\x21\xac\x7e\x1f\xbf\x71\xf2\x4e\xfa\xf6\xca\x5d\xf0\x85\x19\xb8\x06\x3f\x7e\xb9\x5f\xcd\x3a\x0f\x48\x3c\xb7\x1a\xd9\xa3\x79\xe7\x00\x4a\xd5\xa0\x97\x26\x66\xfb\xc3\xb3\x73\xdc\xe1\x16\x44\x74\xcb\x8a\x47\x3a\xaa\x85\x44\xe7\xb2\x46\xd3\xd6\x24\x24\x10\x18\x79\x26\x81\xe7\x3c\x5f\xcf\xee\x96\xcb\xc5\x6c\xe3\x6a\x9c\xd3\x38\xd3\x95\x84\xd6\xe6\xa9\x44\x79\x1a\x68\xe1\x5b\x1a\xad\x0a\x34\x86\x74\xa4\x1b\xde\xc5\x60\x35\xcf\x36\xbe\x90\xf2\xb8\xfe\x6a\xeb\x2b\x86\xce\x73\x1f\x76\x6a\x92\xca\x82\xa1\x2d\xcc\xe4\x11\x94\x53\xc0\x5d\xab\xfd\xd1\xe2\xd3\xd8\xe1\xa3\x01\xb6\x55\xad\x0f\xc1\x53\x90\x4a\x61\xe3\xdc\x54\xfa\x94\xca\xb9\x8f\x81\xcb\x13\x33\x60\xf5\x31\xe4\x9f\x37\xe0\x29\xed\x98\xa8\xb0\xcf\x1a\xba\xd6\x0b\x09\x0c\xb6\x8c\x27\x01\x74\x4e\x2e\xba\x2a\xb1\x53\x8f\xf8\xba\xee\x76\x5f\xc3\x8c\xb1\xa5\x56\xed\xbe\x5c\xb8\x2b\xce\xd8\x15\x29\x7e\x69\x48\x2b\xe1\x4e\x59\x92\x6d\xdd\x29\xd8\x87\x2e\x87\x13\x31\x4a\xdf\x11\x92\xf7\x83\xbe\xf7\x23\x6a\x23\x4e\xc4\xc8\x5b\x78\xb9\xe7\xf4\x16\xd8\x68\xb4\xf6\x38\x1b\xef\x3c\x7f\x15\xeb\x04\x4f\xab\x6a\x55\x31\x89\xbd\xce\xba\xb2\xa6\xff\xf2\x02\xd7\xef\xf3\x39\xb3\xec\xdb\xc3\x65\x5b\x2f\x15\x47\x13\xb4\xd0\x35\xe4\xd2\x58\xdd\xd2\xfd\x02\x79\xda\xe9\x63\x7a\x7b\xae\xd0\x8d\xc6\x83\x50\xad\x59\x8f\x05\xfd\xac\x3f\x39\x3f\x4e\x97\x32\x7d\x4b\xf5\x8b\xda\x64\x9c\x6b\x34\xc9\x39\x61\xd5\x23\xca\xf3\xcb\xd1\xf0\xf6\xe0\xa6\x9e\xdd\xd3\x85\xeb\xbb\x11\xf2\x31\x99\xfb\x16\xee\xbf\xf1\x8a\xe8\xd0\x4f\x1f\x0f\xbf\x71\x35\x3f\xbb\x68\xbd\xd2\x4c\xf7\x52\x18\x8e\x68\x6f\x73\x7a\xc6\xc2\xad\xc0\x73\xd5\x8d\x8e\x19\x1c\x84\xf9\xdb\xfa\x6e\xf9\x47\x48\xa4\x2f\x9b\xaf\xf2\x14\x48\x9d\x3a\x96\xe9\xae\x7d\x95\xf1\x17\xfc\x3f\x79\x73\x0d\xdb\x23\x75\xbd\xbf\xc5\x44\xcf\xed\x0e\x06\x20\xb9\x62\xba\xcf\x9b\x7c\xf9\xf0\x8f\x2f\xd9\xed\xfc\xa7\x3f\x77\x4d\xf3\xec\xfe\x53\xbe\x4c\xdb\x66\x77\xcb\x4d\x96\x2f\x17\xf7\x5f\xd6\x8b\xcd\x97\x5f\xb3\xdb\x9b\xf5\x78\xd7\x08\x5e\x3a\x60\xb3\xb8\x5d\xdd\x90\x08\x7a\x90\x7e\x0b\x0c\xff\x05\xf8\xff\x57\x74\x92\xbb\xa6\x64\x7f\xfa\xf1\xa7\xc4\xc7\xf4\xd1\xe4\x35\x1a\x3a\xfe\xe4\x12\x3d\xb6\xf9\x15\x3f\x7f\x9f\x3a\x9f\x18\xbd\xa9\xf9\x4d\xf7\xc2\x8b\xd4\xe4\xeb\x7f\x02\x00\x00\xff\xff\x02\x14\xca\x88\x47\x1a\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 6727, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
