// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\xcd\x72\xe3\xb8\x11\xbe\xf3\x29\x60\xb9\xe6\x36\xd9\xda\xb3\x6f\x1c\x8a\xb6\x99\x91\x48\x16\x49\x79\xd7\xb9\xb0\x20\xb2\x25\x21\xa6\x00\x16\x00\xda\xd1\xde\xf2\x5e\x79\xa7\xbc\x42\xaa\x01\xf0\x4f\x92\x67\x34\x95\xd3\x8c\x6d\xf4\x7f\xf7\xd7\x5f\xf3\xbe\x12\xc7\xa3\xe0\x5e\xec\xaf\xc3\x32\xfc\x33\xca\x8b\xfc\x81\x2c\x62\x7a\x04\x42\x1b\x09\xb4\x3e\x11\xf8\x17\x53\x5a\x2d\xbc\x28\x2d\xe3\xa4\x18\x1f\xa5\x0d\x50\x05\x64\xc7\x9a\x86\x30\x4e\xf4\x01\x88\x84\x3d\x53\x5a\x9e\x48\x94\x12\x61\x7f\xa5\x4e\x4a\xc3\x91\x28\xd0\x9a\xf1\x3d\x69\xe9\x1e\x16\x9e\xe7\xdd\x57\x4d\xa7\x34\x48\x2f\x58\x6d\xf2\x22\xcc\xca\x65\xb8\x0a\x8b\xb0\x7c\xf4\xa3\x55\xb8\x7c\x20\x8b\x8a\x72\xc2\x85\x26\x35\x34\xa0\x81\xb8\xe7\x68\xa8\xea\xa4\x04\xae\x89\xd2\x54\xc3\x62\x50\x10\xe5\xc6\xbd\x6c\x13\xc7\x51\xfc\xf4\x40\x16\xc5\x61\x22\xa6\x8c\x32\xd9\x71\xce\xf8\xfe\x42\x68\x95\x04\xfe\xea\x81\x2c\xa2\x63\x2b\xa4\x1e\xa4\x2a\xca\x51\x6a\x0b\xa4\x6b\xf7\x92\xd6\x50\x5b\xd7\x25\xd4\xc0\x35\xa3\x8d\x37\xf3\xba\xcc\xc2\x3c\xd9\x64\x41\xf8\x40\x16\x8f\x94\x35\x50\x13\x2d\x5c\x00\x77\xa4\x38\x80\x04\x74\x84\x72\x42\x95\x12\x15\xa3\x1a\x6a\x72\x10\x4a\x93\x8e\xd7\x20\x89\x3e\x30\x45\xde\xe0\xb4\xf8\x44\x6d\xf9\x8f\x24\xfe\x25\xdd\x7f\x09\x0e\x57\x74\x3f\xfa\x9b\x55\x51\x06\x59\xb8\x0c\xe3\x22\xf2\x57\x65\xe0\xc7\x26\x0d\xd6\xec\x03\x59\x2c\x61\x47\xbb\x46\x93\x31\xd2\x49\x2e\xac\xd1\x7a\x61\x7b\x26\x78\x0e\x83\xef\x63\xd9\x4c\xd2\x47\x29\x8e\x8d\x34\x8a\x9a\x86\x30\xbd\xa5\xcc\xff\x3b\x05\xd2\xbc\x59\x78\xa9\x9f\xe7\x7f\x24\xd9\x72\x70\x26\xde\xac\xb0\x24\x2d\x55\xea\x43\xc8\x9a\xf4\x0d\xb1\x05\xb2\x6d\x28\x7f\xfb\xef\x7f\xfe\xbd\xf0\xd2\x2c\x7a\xf1\x8b\xb0\xfc\x1e\xbe\x9e\x0b\xa2\x27\xad\x64\xef\x54\x03\x06\x3e\xf1\x62\x14\xf7\xee\x31\xfd\xde\x73\x92\x17\xa5\xbf\xca\x42\x7f\xf9\x3a\xf6\xf7\x33\x56\xe6\x7c\x08\x5c\x65\x8c\xc4\x10\xf4\xd5\x82\xd8\xca\x62\x4d\x9c\x8a\x49\x61\x3e\x98\x3e\x98\x04\xb8\x4e\xbb\xa6\xb7\xfc\xf6\x5a\xa6\x59\xf2\xf7\x30\x28\xfe\x2f\x13\xad\x14\xff\x84\x4a\x2f\xbc\xfc\x35\x2f\xc2\x75\xe9\xc6\xf8\x31\xd9\xc4\xcb\xeb\x53\xdc\x88\x8a\x36\x38\xc2\x3b\x26\x95\x36\x89\x8a\x13\x94\xf3\x5f\xfc\x68\xe5\x7f\x5b\x61\x8b\xc4\x82\x44\x2d\xa1\xef\x94\x35\x74\xdb\xc0\xc2\x8b\x72\x3b\x46\x26\x86\xc9\x00\x33\x3b\x53\x56\x29\x3a\xbc\xb0\xf9\x8e\xd6\x69\x92\x15\x65\x98\x65\x49\xd6\xd7\x2c\x16\xa4\xa6\x9a\x62\x98\x4e\xec\x83\x2a\xb2\x13\x1d\xaf\xef\x88\xf3\xb4\x3a\x40\xf5\x66\xfc\x74\x4f\x76\xac\x81\xbb\xb9\x52\x54\x57\xbe\xf8\xab\x0d\x7a\x1a\x1e\x5b\x7d\xb2\x7a\x05\x27\x0d\xe3\x40\xbe\xa8\xb3\xf7\x49\x51\x06\xc9\x3a\x35\x35\xe8\xe5\x22\x5e\x89\x63\x6b\xd0\xe7\x87\xc2\x7f\x64\x49\xfc\x54\x3e\x26\xd9\xda\x2f\x9c\x98\x94\x50\x69\x62\xbd\x13\xf2\x48\xf5\xa7\xc2\x93\x29\x9c\x56\x25\x98\x8c\x90\xd0\x36\x03\x9f\xea\x70\x5d\x32\x2f\xab\xad\xfa\x0d\xd2\xae\xdb\xe2\xcd\xfa\x81\x2c\x7c\xa2\x85\xa6\x0d\x11\x3b\xf2\x45\x11\x29\x3e\x14\xfe\xd7\x84\x4f\x25\x10\xba\xe5\x18\x4e\xf3\x95\xa8\x37\xd6\x3a\x3d\x03\x3c\xf5\x7d\x1b\xc5\xf3\xa1\xd8\x32\xee\x70\x4e\x82\x12\x9d\xac\xd0\x89\xaf\x44\x02\x55\x82\x3f\x7c\xe2\x4f\xee\xbf\xcc\xb1\x4e\xd1\x77\xd7\xf0\x67\xc2\x9e\x77\x8f\x38\x32\x22\x08\xe6\x61\xed\x17\xc1\x73\x8f\x02\x3d\x84\x30\x45\x58\x5f\x9d\x85\x97\x64\xd1\x53\x14\xbb\xc4\x4f\xdf\x0b\xc9\xf6\x8c\xd3\xe6\x33\xc1\x4d\x3e\x6e\x0e\x3f\x28\x22\xe3\x68\xd1\xc3\x99\x5b\x35\xc0\x71\x2c\x26\x6d\x2b\xb8\xa6\x95\x36\x8d\x4b\xeb\x23\xe3\xb8\x28\xa9\x16\xf2\xce\x29\x9c\x56\x2f\x16\x44\x75\xd5\xc1\x28\x34\xf3\xe7\x2f\xd7\x51\x7c\x89\xd3\x68\xb4\x76\x58\x6d\x94\x5a\x17\x2e\xb0\xfa\x6e\xee\x74\x16\xae\xfc\x22\x5c\x4e\xe0\x65\x83\x62\x07\x8a\xae\x4f\x41\xc4\x61\x87\x71\x61\xb5\xf4\xd3\xc1\x83\x4d\xba\xf4\x07\x0f\x9a\x9a\xb6\xe7\x86\xa1\x66\xd6\xee\x4b\x98\x45\x8f\xaf\x65\x90\x2c\x27\xcb\xfd\x05\x24\xdb\xb1\x8a\x6a\x26\x38\xa9\x44\x0d\x04\xa4\x14\x72\xe1\x85\x6b\x3f\x5a\x95\xcb\x28\x77\x28\xb3\xa6\xac\xe9\xb9\x83\x32\x2d\x58\x33\x75\x63\x62\x7b\x6d\xd3\xf2\x86\x47\x54\x78\xa4\xba\x3a\x90\x9d\x69\x2d\x0b\x6f\xb8\xc9\x86\xfe\xc9\xf1\xa7\xc1\x57\x4c\xcd\x0f\xd6\x58\xdf\x23\xe7\x4a\x0c\xae\x3d\x90\xc5\x87\x14\x7c\x3f\x2e\x3a\x22\xe4\x44\xc4\x3a\x68\x36\xce\xe0\xdc\xf9\xc6\xf1\xee\x91\x55\x09\xde\xaf\x88\x2c\x7c\x8a\x92\xf8\x56\xca\x41\xac\xf0\xcf\x96\x04\x32\x05\x34\x85\xff\xf6\x86\x90\x6d\xdc\x6c\xc6\x50\x8d\x9f\x6d\xa2\x86\xf2\x39\xf5\xb2\xa8\x1f\xd8\xc4\xee\x41\x4f\x77\xe2\x15\xc0\xaf\x04\xdf\xb1\x7d\x27\x4d\xdf\x98\xc2\x45\x6b\xff\x29\xfc\x5c\x15\x3b\xd2\x3d\xdc\xa6\x28\x2d\xf3\xe7\x24\xb3\x00\xae\xba\xdd\x8e\x55\x0c\x49\x66\xd4\x62\x5a\x44\x0b\x5c\x69\x5a\xbd\x79\x4f\x61\xd1\x57\xa0\xaf\x70\x2c\xfa\x24\x1b\xa0\xc5\xf7\x6e\x6e\xd6\x70\xdc\x82\x1c\x46\xcf\x5f\x62\x3f\x7d\x51\x64\x98\xb6\x2d\x00\x27\xb4\x36\xc4\x72\x3a\xa0\x3d\x0e\x7c\x51\x33\x4c\x31\xfa\x1d\x71\x71\x26\x06\x3a\xd7\x2f\x81\x4f\xb9\x9c\x13\xb8\x46\xe4\x7a\xd9\x67\x3f\x2f\x5d\x79\x26\x2b\x64\x52\xca\xa1\x34\xc1\x15\x84\xf1\xee\xb9\xa8\xc1\x8b\x71\xd2\x7b\x32\xe5\xd8\x78\x59\xf8\xf9\x77\x5c\x2f\x75\x4d\xf0\x11\x4e\x81\x23\xf6\xe6\xc7\xbe\x6b\x1c\x3f\xff\xda\xda\x82\x7d\x50\xa6\x09\xd3\xa4\x16\x1c\x7e\x43\x03\x5b\x5a\xbd\x75\xad\x5f\x55\xa2\xe3\xda\x4b\xfd\xcc\x5f\x97\xe1\x3a\x2d\x5e\xcf\xcb\xd6\x52\x49\x8f\xa0\x41\x2a\xe4\x2e\x45\x99\x6f\xd2\xd4\x56\x77\xc3\x55\xd7\xe2\x66\xc6\x1e\x3e\xb5\x78\x40\xcc\x19\xec\x0c\x9b\x2c\x46\x0c\xf4\xec\x9b\x1f\x7c\xdf\xa4\xa5\x1f\x04\xc9\x26\xfe\x15\xa2\x36\x73\xfc\x66\xc6\xe6\xdd\xe3\xc8\xf4\xc6\xd3\x95\x1f\x5f\xe3\x9c\xd6\x47\x6b\x07\xdf\x8f\x9d\xd5\x29\xa8\x7b\x52\xd7\xdf\x34\xfd\x39\x30\x31\x33\xbf\x36\x6e\x88\xc6\x58\xb9\x3d\x08\x67\xfa\x9b\xc9\xc1\x59\x2a\x7f\x12\x8f\xcd\x1b\xa1\x2e\x71\xbf\x1a\x59\x6f\x24\x5a\x85\xf9\xf4\x36\x70\x80\xe0\x5a\x50\x8f\x86\x90\x4b\x92\x2d\xec\x84\x04\xa2\x3e\x98\xae\x0e\x78\xb4\xea\x0b\x4f\x66\x30\x66\xad\x5c\x5e\x9c\x5b\x40\x61\x14\x84\x9a\x74\xad\x19\xdc\x89\x58\x16\xe6\x45\x92\x85\x97\x72\x12\x94\x16\x92\xf1\xfd\x7c\xd4\x33\x47\x9e\x2e\xeb\x35\x09\xf3\xa7\xc1\x8d\x9c\xfe\xfa\xc9\x31\xce\xff\x70\x60\xf4\x65\xde\x42\x23\x70\x0d\x6b\x31\xc7\xe9\x02\x6f\x4b\xd1\x82\x74\xfb\x7c\xc0\x86\x16\x24\xd2\x5f\x83\x0e\x96\xe2\x5d\xc0\xdb\xb3\xbb\x60\x06\x78\x5b\x78\x03\x9f\xb4\x90\x19\xba\x97\x7d\xf8\xa6\xf1\x2c\x64\x9a\x89\xb8\xd4\x99\xba\xf6\x9c\xe8\x3c\x1b\xdc\x4b\x99\x6f\xe7\x03\x3a\x11\xf6\xee\x91\xe1\x58\x02\x34\xa0\xb4\x6d\xa4\xfc\xc4\xab\x83\x14\x9c\xfd\x85\x25\x56\x20\x2d\x49\xf9\xdd\xd1\xa5\x55\xf2\x14\xc5\xe7\x32\x9b\x29\x4b\xc4\x25\x7f\xe7\x5e\x8f\xb4\xa7\x18\x3f\x9c\xb4\x52\x1c\xd8\x96\x69\x45\xf0\x8d\xb3\xb1\x93\xe2\x48\x1a\xb1\xdf\x63\x83\x31\xfe\xdb\x2d\x24\xd3\xbb\xaf\x98\xf2\x82\x28\x37\x60\x7c\x8e\xd0\x78\xb0\x31\x45\x34\x55\x6f\xe7\x68\x8c\xa2\xef\xc7\xc0\xac\x4c\xef\x65\x5d\x06\x49\xfc\x18\x3d\x8d\x77\x72\x30\x5d\xa6\x17\xf4\x65\x14\x38\xff\xc2\x53\x9c\x2f\xe2\xcf\x1a\xad\x86\xb6\x11\xa7\xa3\xc1\xf5\x86\xf2\x1b\x1b\xce\xbb\x67\x2d\xae\xf5\xc1\x4f\xb4\x07\x5c\x83\x84\x1a\x8f\x5b\x05\x7b\xa3\x12\x5d\x68\x58\xa5\xd5\x88\x5d\xc6\x77\x4c\xed\xf8\xec\x2b\x69\xe7\x87\x32\xdd\x53\xc6\x87\x8f\x62\xd3\xcb\x38\x4a\xf1\x76\x42\x0f\xab\xaa\x6b\x19\xd4\x84\xf2\x7a\xe2\xa4\x04\xa3\xa9\x36\xc2\x51\xfc\xe2\xaf\x22\x4c\x47\xd4\xda\x33\xe3\x9d\x36\xac\xee\x39\xc9\xe4\x5b\x06\xef\x90\x4c\xe0\x39\xb6\x07\x8e\xa1\xdb\x30\x68\x5d\x4b\x50\x0a\x8c\xc5\xdf\x7f\xbb\xe4\x3a\x4a\x53\x69\x82\x71\x2f\x8d\x37\xaa\xdb\x72\xc0\xc6\x33\x69\xfa\x5b\x2b\x44\x83\xe6\xd2\x24\x59\x5d\xad\x53\x94\x12\x7c\x33\x21\x2d\x57\x20\x7f\xf8\x18\x60\xd9\xe0\x3c\xea\x81\x6c\x58\x3e\xab\xb4\x3c\x79\x48\xa4\xf2\x22\x7b\xbd\xfc\xf6\x52\x4c\x3f\x26\x8a\x9d\xfd\x7c\x45\x65\x75\x60\x1a\x2a\xdd\x49\x30\xbb\xfb\x0a\x57\x76\xa0\x34\x60\x65\x8f\xb7\x69\x96\xbc\x44\xcb\x30\x1b\x58\xdb\x14\x73\x2b\x09\x26\x12\x2c\x6b\xa7\xc5\x91\x6a\x56\x91\x23\x12\x13\xe7\xff\x91\xf2\x8e\x36\xcd\x09\x7f\xc9\x76\xa7\xf9\x31\xab\x26\x90\x55\xbc\xa6\xe1\xcc\x84\x41\x2a\xeb\xac\x1b\xf9\x81\x7e\xdc\x91\x84\x37\xa7\xfe\x67\x45\x10\x21\xbf\x92\x39\x52\xfd\x9c\x79\x8c\x78\xbd\xfb\x01\xf3\x20\x06\x07\x60\x7f\x6d\x57\xe3\xa6\x9a\xaf\xd9\x61\x77\xd1\x4a\x9b\x0f\x68\x76\x9e\x14\x28\x85\x87\x48\x1e\xe6\x39\x12\xe0\xef\xe1\xeb\x0c\x46\xdd\xdf\xcd\xd7\x36\xd7\xc7\x8e\xd6\xa6\x59\x82\xcb\x69\xd2\xed\xfd\x5b\x7b\x36\x0a\xfe\x0e\x52\x4d\x09\x97\x11\x43\x3e\x1a\x27\xd3\xe3\xa3\x9b\xdc\xa9\x7d\xfe\x4d\x13\xe3\xc6\xa4\x7b\xf0\x4c\x9d\xd1\x3b\x2c\xf5\x9f\x79\x84\x24\x26\xb7\x7f\x43\x28\x7d\x67\xc8\x11\xce\xda\xe6\x7f\x01\x00\x00\xff\xff\x37\x8b\x6f\x59\x04\x17\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 5892, mode: os.FileMode(420), modTime: time.Unix(1628233279, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x58\x4b\x53\xdb\x58\x16\xde\xeb\x57\xb8\xec\xea\x5d\xcf\xd4\xac\xb3\x53\x6c\x01\x1a\x64\x49\x25\xc9\x74\x33\x1b\x15\x4d\x3c\x19\x26\x80\x29\x0c\x53\xd5\xb3\x8a\x79\xd9\x06\x1b\xbb\x13\xcc\x04\x70\x02\x26\x10\xdc\x10\x3f\x80\x04\x3b\x7e\xc0\x9f\xd1\xbd\x92\x57\xfc\x85\xa9\x7b\xaf\x24\x5f\xcb\x81\xb0\xa1\x0a\xfb\xbc\xee\x79\x7c\xe7\x3b\x0e\x4c\xc7\xe6\xe6\x62\xf3\x8c\xc8\x86\x39\x9d\xfb\x95\x57\x35\xf5\x99\xcf\x0f\xf2\x59\xf3\xac\x0e\x1a\x57\xa0\xf2\x0e\x14\xcb\x7e\x86\x97\x75\x51\xd2\xfa\x02\x56\xad\x01\x8a\x65\xf3\xba\x6d\xb6\x0f\xad\xea\xad\xd9\xad\xf6\x4a\x5f\x7b\xef\x8f\x41\xe9\x02\x6c\xec\x19\xed\xb7\xa0\xf5\x96\x97\xfd\x0c\x13\x98\x9e\x5d\x8e\x2f\x45\x17\x99\xa0\x10\x51\x35\x4e\xd1\x43\x9c\xc0\x69\x9c\x3e\xc2\xf2\x02\x17\x7a\xe6\xf3\xc3\xff\x1d\xc1\xeb\x02\x48\x1d\xf5\xf6\x4e\x40\xf7\x2d\x48\x67\xcd\xcd\x1b\xf8\x3a\x61\xee\xaf\xf5\x0e\x36\xcc\xdb\x13\xbf\xab\xca\xab\x38\x08\x25\x22\x8a\xbc\x38\xfa\xcc\xe7\x27\x02\x46\x33\x0b\x8a\x65\xeb\x2e\x6f\x95\x32\x46\xb3\x72\xdf\x49\x0c\xa9\x08\x52\x90\x15\xd0\xbb\x6a\x1d\xb0\x7e\x4a\xd4\x6c\xc7\xd9\xa4\xd9\x3a\xf3\x33\x28\xd2\xc5\xe8\x8b\xe8\xfc\xd2\xcc\xd4\x2c\x33\x10\xa4\xae\x70\xaa\x14\x51\x82\x1c\x32\x40\xe2\x3c\xb9\xb4\xbe\x9c\xde\x77\x12\x56\xed\xd4\x3c\x7b\xd7\x7b\x73\x6a\x34\xb7\x60\x31\x0d\xd6\xaf\xad\xc4\x8e\xd1\x6c\xc3\x62\xcb\xff\x80\x11\xfd\x1f\x92\xf8\x54\x4b\x20\x57\x33\x77\xca\x20\x83\x8d\x8d\xb0\x11\x41\xd3\x83\x0a\x17\xe2\x44\x8d\x67\x05\x3d\xc8\x8a\xf8\x71\xc4\x0f\x4a\x47\xfb\x9d\x55\x3d\x01\xc9\x0a\xcc\x56\x8d\x66\xd6\x5a\xed\x12\x27\x38\x23\xb8\xc0\xc1\x31\x2e\x38\xde\xcf\x3d\xf1\x48\x8a\x4d\x14\x8c\xe6\xb6\xb9\x53\x86\xa9\x06\xfa\xf0\xa0\x09\xf2\x19\x3f\x23\xb3\xaa\xfa\x8b\xa4\x84\x5c\x87\x62\x44\x20\xc9\xdc\x30\x8f\x12\x8e\x5e\xcb\xfc\xb3\x85\x1d\xc9\x0a\x3f\xc1\x6a\x9c\x3e\xce\x4d\x7a\x35\x9c\x17\x7a\x34\x18\x26\xf0\xaf\x58\x7c\x89\x19\x93\x54\x4d\x67\x05\x85\x63\x43\x93\xfd\x56\x23\xe9\xa4\x7a\xd1\xce\x2b\x96\x76\x9f\x32\x9c\x4e\x57\xcf\x6c\xe7\x48\x3a\x9d\x7e\x1a\x36\xa0\x3f\x9f\xd4\x65\x45\xfa\x3b\x17\xd4\x9e\x6a\xab\xf4\xcd\x3c\xa8\xe2\xf0\xd5\x49\x55\xe3\xc2\xba\x3d\x22\x23\x52\x44\x0c\xd9\x13\xb2\x9e\x22\xf3\x00\x8b\x9f\x61\xb1\xc5\xcb\xa4\x10\x12\x12\x65\x27\x58\x5e\x60\x9f\x0b\xa8\x6e\xbc\xec\xb3\xbe\xae\xc1\x56\x1e\x65\xe6\xe6\xda\xcf\xf0\x2a\xe9\x58\x1c\x22\xb6\x65\x47\x60\x34\xb7\xc9\x6c\xf9\x78\xd9\x07\x36\xae\xcc\xf3\xc4\x7d\x27\x03\x8b\xe7\xd6\x6a\x17\xa6\xf2\x60\xf3\x10\x36\xda\x60\xf3\xc8\x4f\x72\xc9\x87\x65\x49\xd1\x74\x4e\x51\x24\xc5\xa9\x01\x2c\x9e\xc3\xf4\x2d\x48\xd5\x41\xae\x46\xc6\xc1\xdc\x5f\x83\x85\x3a\xcc\x56\xf1\x5b\x1b\xf0\xe3\x6b\x78\x78\x4a\xbe\x82\xbb\x49\xa3\x7d\x83\xc3\xa6\x0d\x22\x53\xfa\x04\x2b\x44\x50\xf4\x3f\xc5\x7d\x56\x29\x03\x8b\x69\xf3\xcf\x16\xb1\xe3\x11\x96\x34\x3d\x28\x85\x65\x9c\x73\x8f\x12\x16\x47\xf3\x5b\xcd\xc0\xc2\x97\x41\xbd\x5f\x14\x49\x1c\xd5\x47\x24\x25\xcc\x6a\xae\x86\x79\x51\x03\xb9\x8f\xf0\xa8\x03\x3a\x39\xa3\x99\x85\x95\x8f\x66\xc9\xe3\x8f\x9a\x10\xba\x1e\xb6\xc7\xf4\x2d\x72\x97\xaa\x83\xda\x46\xef\xcd\xe9\xa0\xa6\xdd\x01\x8f\xa9\x91\xb2\x0f\xaa\xd9\x2d\x24\x46\xc2\xa8\x77\xd6\x2f\x7d\x9e\xc7\xc1\xca\x47\xd0\x6c\xde\x77\x32\x56\xe3\xda\xba\x4b\xda\xca\x2e\x1e\x38\x1d\xc8\x63\x7f\xa4\xce\xa4\x1d\x90\x21\xe3\xee\x3d\xea\x7b\xbb\x17\x33\x60\xfb\x10\x1c\x1c\xdd\x77\xf6\x7f\x8a\x7f\x37\x08\x95\x9d\xe0\x5c\x2b\x3f\xd2\x67\x98\xc0\x72\x3c\xba\xd8\x1f\x70\xf4\xf0\x30\xab\x05\xc7\xdc\xe9\xee\xed\xec\x59\xb5\x9a\x9f\x91\x14\x7e\x94\x17\xed\x94\xba\x22\xdb\x87\x83\x52\x11\xb5\x0f\xb9\x6c\x50\xe3\x71\x2c\x04\x4f\x60\xf1\x1c\xe4\x11\x9a\x91\x26\xb3\x12\x3b\x68\x79\x54\x4b\x66\x7e\x03\xfc\xf1\x0e\x77\x18\xd6\xa6\x73\x8f\x10\xba\x72\x42\xf4\xb1\x04\x1b\x0a\xf3\xe2\x43\xb8\xe7\x9b\x7a\x31\x37\x33\xef\x23\xe2\x04\x63\xac\xe3\x0b\x0a\x01\xe9\xe8\x14\x4e\x60\x35\x2e\x44\x0d\xbd\x1d\xe6\x55\xc9\x45\x5f\xa7\xd6\x42\x88\x95\x5d\xa7\x11\x39\xc4\x62\xa7\xe8\xd3\x01\x67\xc6\x5d\x15\xee\x7c\xc3\x9e\x26\x38\x85\x1f\x99\xd4\x83\x52\x88\x5a\x74\xbd\xf3\x8c\x55\x4b\x50\xd9\xe2\xc2\x2c\x2f\xe8\x21\x5e\xb5\x61\xa0\xb7\x52\x35\xda\x37\x64\x9b\x5a\xc7\x17\xe6\xa7\xc4\x43\xe9\x72\x74\xe9\x62\x10\x6d\x90\xf9\xd6\x5b\xcf\xba\xd0\x65\xc3\xbe\x5b\x60\x15\xfd\xd7\x87\x7f\x07\xe9\x5d\xec\x27\xe5\x74\x80\x7f\x50\x17\xa3\x08\xad\x05\x53\xbb\x83\xe5\x27\x41\x61\xe0\x26\x01\x99\xd5\x4b\x0a\xb7\x19\x26\xb0\x18\x7d\x39\x13\x9b\x77\x00\x58\xe1\x46\x79\x49\x7c\xd2\x96\x05\x99\x16\x38\x3c\xa4\x01\x98\xda\x8d\x4c\xe0\xbf\xb1\xf9\xa8\x63\x15\xed\xd7\xa7\xd9\x74\x2c\x0c\xe0\xfa\x6a\xd9\xec\x5e\x59\xd5\x12\x48\xbd\x19\xa4\x11\x04\x3e\xad\xed\x06\xc8\xed\xda\x60\x80\xf7\x09\x8d\x9a\xbd\xf5\xac\xd9\x25\x5b\x81\x0f\xb3\xa3\xdc\x43\x8a\x85\x22\x58\xcd\x3d\xa4\x28\xeb\xea\x98\xa4\xa0\x14\xf2\x0b\x3e\x67\x25\x30\x4c\x20\xb6\x10\x9d\x8f\x2f\x4d\x4d\xbf\x62\x46\x39\xcd\x49\x9e\x53\x95\x3e\xb0\xe1\x4c\xa1\xa4\x2c\x2c\xc6\xfe\x1d\x9d\x5e\x0a\x47\xe7\x7e\x8b\x2e\xba\xdd\xcf\x86\x6c\x58\xb3\xeb\x88\xdf\xee\x6c\x0d\x7a\x44\x28\x04\x74\x47\x98\x6c\x0e\xb2\xb4\x1d\xfb\x2e\x35\x71\xe0\xf3\x81\xf9\x24\xe3\x34\xc4\x4b\x1c\xad\x31\x56\xd5\xed\x74\x23\x15\x2c\x4c\xaf\xed\xfb\x4e\x62\x48\x97\x09\xcc\xc7\x5e\x44\x19\x11\xcd\x99\xc3\x1b\x6c\x76\xa8\x6b\xac\x3a\x8e\xf1\xf8\xc6\x68\xef\x5a\x9b\x2b\xe6\xca\x37\x58\xa8\xf7\x92\x39\xf8\x36\x6b\x74\x8b\x08\x94\x8b\x65\x98\x3e\xb3\x4a\x99\x9f\x7d\x56\xad\x61\x56\xd2\xe0\x76\x1d\x54\x57\x8d\xf6\x67\xf2\x31\x5a\x4a\xb5\xc2\x5f\x91\x9b\xdf\xa6\xa6\x5f\x2d\x2f\xb0\xd3\xd3\xb1\xe5\xf9\x25\x46\x66\x15\x36\xac\x73\x61\x59\x9b\x44\x1e\x72\x2b\xb0\x50\x77\xea\x84\x1e\xae\x46\x64\x99\x14\x10\xa1\xff\x4e\x0d\x66\x10\x9b\x35\x2f\xdb\xe0\xc3\x96\x9f\xf1\x70\x30\x78\x54\xea\x9d\x67\xa8\x89\xb5\xdb\xf8\x39\x1b\x1c\x8f\xc8\x3a\x1b\x0c\x4a\x11\xf1\xa9\x3c\x05\x9c\x24\x8d\x76\xd7\xfa\xf2\x09\xe4\x1a\x0f\xb0\x15\x26\xb0\x30\x3b\xe5\x8e\xa0\x2c\xb0\xe2\x63\x24\x8a\x1e\x07\x34\xca\xc5\xb2\x53\xc6\x2d\x87\x7c\x57\x8c\xee\x9d\xb9\x53\xf6\xf2\xdd\x1f\x44\xea\x31\xec\x89\x94\x3a\x1c\x9e\xe3\xd4\x7b\xb2\xf2\x48\xc4\x9e\x14\x3c\x1a\xb1\x63\x8d\x17\x38\x95\xe6\xa9\x36\x77\xb3\x6d\x23\x7b\x84\x08\x81\x8d\x2c\x48\x25\x61\xf6\x98\x76\x32\x80\x12\xc4\xa2\x7b\x9a\x90\x2e\x23\xd2\xdf\x39\x4d\x14\x4e\xd5\x24\x85\xf3\x88\xc3\xc4\x31\x38\xc9\x3a\xe2\xee\x9c\x29\xd1\x78\x6c\x79\x71\x3a\x3a\x9c\x67\xea\x19\x8f\x04\x4f\x37\x97\x87\xfa\xf6\xe7\x6e\x80\xe8\x5e\x7e\x30\x5a\xdb\x9e\xe9\xb3\xee\x0e\x10\xb7\xa9\x9c\x90\x31\xa2\xa8\xe1\x10\x6a\x38\x94\x35\xeb\x2c\x00\x97\xf2\x10\x18\xe2\x1c\x7e\x85\xc9\x0e\x68\x5c\x39\x28\x84\x9b\x72\xd8\x1e\xdd\x32\x94\x55\xcf\xa0\x0c\xeb\xd1\xc5\xa2\xf4\x98\xc0\xec\x8b\xa9\x05\xb2\xdd\x5d\xdc\xb3\x2f\x9b\x7c\x06\x56\x4e\x41\xaa\x8e\xc6\xd6\x5e\xf1\xad\xbf\xd9\x4c\x40\x90\x46\x79\xd1\xab\xe1\x12\x01\x62\x1d\xa7\x05\x4b\xf7\xd7\x3b\x39\x96\xcd\x4f\x09\x58\x39\x46\x5f\x11\x15\x73\xaf\xdd\xdb\xdb\xf0\x3d\x40\x89\x98\xc0\xf4\x4c\x9c\x09\xf2\x2a\x86\x33\x2f\xc6\xa1\x48\x9d\xf5\x0a\xd3\x17\x30\x97\x33\x9a\x15\x73\x7f\xcd\x68\xb7\xc1\x66\x09\x69\xff\x67\x2e\x18\x9b\xff\xe7\xcc\x4b\x66\x22\xac\x07\x25\x71\x84\x1f\xed\xdf\x55\x64\xe5\x50\xfb\xb9\x2f\xe3\xbd\xd1\x5d\xd1\x7e\x57\x50\xb5\x78\xb4\x37\x98\xc0\xcc\x02\x5a\x69\xfd\x9f\x16\xf0\x45\x64\xee\xaf\xf1\x32\xac\x7e\x45\x9c\xa3\x71\x05\x8b\x69\xfc\x9f\x7b\xd1\x58\xb5\x46\x2f\x99\x85\xbb\x75\x22\xed\xfe\xfe\xe0\xb9\x9a\x50\xf3\x1c\x5f\x80\xec\x11\x26\x4a\x19\x97\xf4\xf5\x92\x9b\x70\xe7\x16\x6b\xf1\xe2\x04\x2b\xf0\x21\x2c\x0f\x8a\x75\xf0\xfe\xb5\x7b\x36\x38\xdb\xd6\x2d\xe1\x21\x4c\xe5\x71\x60\x44\x90\xac\x0b\x54\x79\x9f\xbb\xa7\xad\xaf\x0d\x70\xb6\x45\xbe\x07\x6f\x32\xa0\x92\x37\xbb\x7f\xb8\x0f\xfd\xcb\x42\x2c\x36\x8b\x4c\xca\x92\x24\x0c\x65\x91\x5f\xf0\xc1\xcb\xa3\xef\x72\x18\x74\xc4\x51\xbf\x83\xf8\x6d\xaa\x14\x5f\x5a\xfc\x9d\x41\x8b\x5e\xd5\x94\xc9\xe1\xcb\xd8\xaa\x9d\xc2\x0f\x37\xf0\x83\xbd\x5a\x70\xe5\xd1\x71\x68\xbb\x70\xfb\xdc\x46\x52\x17\x40\x1c\x04\x92\x15\x69\x82\x0f\x71\x8a\x4b\x22\xac\xe4\x39\xd8\x2c\xc3\x72\x09\x74\x72\x20\x75\x00\xda\x2d\xf7\x57\x18\x92\x5b\x98\xde\x02\x9b\x65\xc2\x74\xe9\x4b\x85\x9a\x6d\x6d\x52\xe6\x5c\x83\xee\xfe\x23\xa3\xed\x6e\x41\xd4\x32\xed\x75\xf2\x95\x0f\xc1\xc7\xcf\x83\x53\xfc\xe3\x2d\xd8\x87\xac\x1f\xac\x00\x9c\x68\x73\x7f\x8d\x7c\x68\x56\x76\xcd\xc2\x29\xc2\x64\x7c\x8c\xf4\x41\x36\x1e\x8d\xc7\x11\x35\x55\x39\x55\x45\xbc\x6a\x9c\x9b\x1c\xc0\x11\xfb\x7b\xdf\xab\xe8\xef\x3e\x54\xaa\x42\xca\x66\x4b\xb2\x22\x21\xe0\xa5\x1a\xcd\x11\x25\x03\x6e\x75\x3f\x83\xcc\x2e\x09\xcf\x56\x41\x54\x47\x94\x68\x72\x4a\x1f\x1f\x4e\x46\x51\x50\x4b\xb1\xc5\xa9\x97\x51\x06\x17\x0a\x85\x85\x6a\xf5\xab\xca\x6b\xf4\xe4\xa3\xbf\x2b\x65\x98\xcb\x1b\xb7\x07\xa0\xb0\xe1\x67\xfe\x1f\x00\x00\xff\xff\xcd\x39\xef\xda\xd3\x13\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 5075, mode: os.FileMode(420), modTime: time.Unix(1628233279, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
