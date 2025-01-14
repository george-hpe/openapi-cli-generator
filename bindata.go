// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/commands.tmpl
// templates/main.tmpl

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTemplatesCommandsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x5b\x73\xdb\x36\xf6\x7f\x26\x3f\x05\xca\x49\x3a\x64\xac\x50\x69\xff\x9d\xff\x83\xb6\xda\x99\xc4\xb9\xd8\x33\x8d\xe3\xb5\x9d\xf6\xc1\xeb\xd9\x40\xe4\x91\x84\x9a\x04\x68\x10\xf4\x65\x54\x7e\xf7\x9d\x03\x80\x20\x29\x52\xb2\x9d\xdd\x97\xf5\x83\x2d\xf1\x80\x38\xf7\xdf\xb9\x78\x3a\x25\x87\x22\x05\xb2\x02\x0e\x92\x2a\x48\xc9\xe2\x81\x88\x02\x38\x2d\xd8\xeb\x24\x63\xaf\x2d\x41\xc8\x98\xbc\xff\x42\x4e\xbe\x5c\x90\x0f\xef\x8f\x2f\x62\x7f\x3a\x25\xe7\x00\x64\xad\x54\x51\xce\xa6\xd3\x15\x53\xeb\x6a\x11\x27\x22\x9f\xa6\x94\x33\xc8\x56\x8a\x3e\x64\x42\x4e\x47\xef\xf2\xfd\x82\x26\xd7\x74\x05\x24\xa7\x8c\xfb\x3e\xcb\x0b\x21\x15\x09\x7d\x6f\xb3\x21\x6c\x49\xe2\x63\xfd\xa0\x8c\x3f\xe6\x8a\xd4\x75\xb0\xcc\x55\xb0\xd9\x10\xe0\x29\xa9\xeb\xc1\xa1\x73\x25\x19\x5f\x95\x78\xb0\x34\x1f\xf7\x1c\xbe\x60\x39\xe0\x49\xc5\x72\xe8\x1e\x0b\x80\x27\x22\x65\x7c\x35\xfd\xb3\x14\x3c\xf0\x7d\x2f\x78\x9e\x4e\xd3\x24\x63\x41\xff\xad\xe2\x7a\x35\x05\x29\x85\x2c\xb7\x08\x25\x93\x55\x51\x02\x9f\x66\x62\x25\xab\x01\xb5\x58\xfe\xf4\x7f\xd3\x44\x2c\x24\x1d\xa5\xdc\xb2\x02\xa4\xa6\x88\xe2\x7a\x15\x33\x3e\x5d\xff\xcc\x05\x9f\xae\x80\xab\x0c\x72\xca\xe3\xdb\x9f\x03\x3f\xf2\xfd\xcd\x86\xa4\xb0\x64\x1c\x48\x50\x50\x49\xf3\x32\xb0\x16\x79\x4d\x24\xe5\x2b\x20\xf1\x97\x42\x31\xc1\x69\x76\xaa\xc9\x9a\xaa\xc9\x6c\x49\xe0\x86\xc4\x17\x0f\x05\x90\x60\x21\x44\x06\x94\x9b\x97\x3d\x2f\xc9\xd3\xf8\x63\x46\x57\x65\x18\xc5\xef\x84\xc8\x42\x34\x63\x7c\xf8\xdb\xf1\x09\x35\xa6\x9d\x90\x25\xcd\x4a\x98\x10\x4d\x78\x0f\x65\x22\x99\xe6\x83\xc4\xc8\x72\x80\xac\x84\x3e\x1b\xc6\xd5\xff\xff\x32\xc6\xe4\x18\x09\x23\x5c\xde\x3c\x97\xc3\x32\x13\x74\x07\x8f\x8f\x86\x34\xc6\x25\x7e\x0a\x9f\xe1\x8d\x26\x2a\x47\x2e\x0c\x82\x47\xee\x73\x81\xeb\x3e\xf7\x7c\xf6\x07\x65\x0a\xa4\x75\xd6\xd0\x19\x77\x94\xa9\xd7\x78\xbd\x39\xb7\xdb\x31\x96\x7e\xbe\xc6\xd4\x33\xfc\x7b\x2c\x93\x8c\xc5\xe7\xa0\x0e\xab\x52\x89\xdc\xf0\x48\xf2\x34\xf2\x7d\x8f\x2d\x49\x97\xef\x11\x2d\xed\x47\xb2\xf1\x3d\xcf\x84\x5a\xfc\x8e\xf1\xf4\xd4\xbd\xd6\x1c\x8e\x7c\xaf\xf6\xdb\xb4\xc3\x8f\x2f\x38\x8a\x37\x9b\x93\xd8\xca\xa9\x1f\xd2\x82\xe9\x67\x9f\xc4\xd6\xd3\xd3\x6a\x91\xb1\x44\xd3\xcc\xc7\xf6\x84\x7f\x4b\x25\x69\x5e\xae\xeb\xf3\x6a\x91\x88\x3c\xa7\x3c\x25\x18\xc2\xbe\xbf\xac\x78\xd2\xa5\x83\xbc\x05\x89\x62\x5f\x5e\xe5\xb4\xb8\x34\xe0\x71\x65\xfe\xa0\x2a\x12\x54\x25\xf9\x18\x75\xa3\x7d\x65\x3d\xf2\xa2\xd4\x17\x69\x91\xec\x9d\x36\x1e\x46\xdf\xf3\xbc\x20\x6d\x3d\x1f\xcc\xb4\x37\xec\x1d\xdb\x31\x31\x31\xe7\x2b\x99\x6d\x9d\xfb\x7a\xf6\x9b\xa3\xd7\x13\x23\x4d\x13\x38\xb5\x6f\x0c\x6b\xa5\x13\x05\x02\x14\x5e\x88\x02\x7e\x69\xbe\x19\x19\xa7\x53\xd2\xb7\x6b\x5d\x63\x6c\x38\x9b\x22\xd5\x85\x88\xef\x75\x2d\x38\xfe\x42\xe8\x18\xc7\x67\x70\x53\x31\x09\xa9\x03\x97\xfe\xcd\xc6\x22\x13\xe2\x24\x37\x81\x43\x5e\x69\x7c\x8b\x7f\xc7\xdf\x16\xbc\x0f\x29\x3f\xa2\xb7\xf0\x4e\xa4\x0f\xa4\xae\x27\x64\x81\x1f\xac\x45\x9b\xb7\x23\x12\xbe\x6a\x11\xf0\x0c\xca\x42\x70\x8c\x75\x64\x7a\xa6\x1d\xa9\x41\x00\x5f\xd7\x98\x6c\x82\x75\x4d\x79\x9a\x81\x3c\xa5\x6a\x8d\xe6\xd1\x89\x71\x64\x9e\x35\x59\xe3\x7b\x18\xf0\xa3\x61\xa5\xdd\xd9\xbd\xc2\xdc\x60\x22\xba\xae\x49\x40\x0e\x48\x87\xec\x7b\x1e\x26\x96\xd7\xc6\x8b\x51\xf5\x13\xa8\x06\x2d\x0c\x49\x63\x01\x5b\x12\x7b\x70\x3e\x27\x41\x60\x98\x35\x4f\xc6\xc2\xf8\xd2\xdd\x76\xcc\x55\x73\xd5\x6b\xc6\x53\xb8\x0f\xa2\xab\x4b\x1d\x45\x57\x8d\x0c\x95\xcc\x50\x00\x73\xe8\x40\xeb\xad\x35\x40\x85\x2d\x10\xd9\xf8\x61\x13\xf2\x42\xbb\x46\xc7\xcf\xc0\xa9\x28\x55\x5b\x32\xcc\xc9\xf8\x98\x63\xc9\x51\xeb\x06\x6a\x35\xbb\xb9\x75\x59\x19\x9f\x41\x91\xd1\x04\xc2\x4a\x66\x1a\x8c\xbe\x6d\xbe\xe9\xe8\xb0\x6f\x5b\xcb\x6f\x36\xdf\xea\x6f\x1a\xb6\x5a\x92\x0b\x9f\x09\xf9\x29\x6a\x58\x37\x81\xdf\xc7\x2f\x4f\xc2\x0d\x8a\x8c\x38\x76\x98\x31\xe0\x2a\x46\x2d\x3f\x83\x5a\x0b\x3c\x12\x46\x98\x45\x28\x43\xe4\xf7\xd2\xf9\x49\x0a\x0f\xf5\xbd\xa9\x40\x3e\x38\x85\x91\xf9\x9c\x48\xb8\x89\xdf\xa6\xe9\x3f\x90\x64\x4a\x41\x8b\xc5\x23\x4a\x59\x8d\xba\xa5\xab\xc3\x60\x0d\x34\x05\x39\xce\xe1\x48\xd3\x9e\xc3\xa2\xb5\x59\xc7\x64\x8f\xb4\x06\x5e\x3f\x83\x67\x73\x62\xc1\xfe\x13\x28\x24\xe9\x0c\xfb\x8b\x28\xa6\x32\x8d\x05\xdb\xc5\x4f\x33\x37\xc9\xd4\xb9\xe6\x07\x1d\xcd\xfa\xe5\x13\x96\x69\xd0\xd1\x0a\x76\x1a\x91\xa1\x7d\x9f\x60\xe0\x65\xae\xe2\xf3\x42\x32\xae\x96\x61\xf0\xf2\xd6\xd8\xa3\x63\x89\xc8\x71\xe9\xb6\x0a\x23\x96\x7e\x8a\xa9\x9f\xc1\xac\xb1\xbc\x37\x0c\xd9\x51\xbc\xb3\x46\xd3\xa0\xf7\x43\x0b\x05\xe3\x52\x1d\x0a\xae\x80\xab\xd7\x68\xcd\xa6\xd1\xf8\x0c\x29\xa3\x16\xfd\x02\xec\x13\xd2\x07\x8b\x37\x78\x67\xd4\x8a\xd2\x91\x04\xb3\xc6\x00\xe1\x3b\x58\x0a\x09\x61\x07\xc7\x26\xd6\xed\x13\x64\x1e\x99\x54\x2b\x0b\x8d\xab\x18\x13\x28\xd1\x7b\x11\x5a\x0c\xc3\x87\x3f\xcc\x09\x67\x99\x11\xdb\x16\x55\xce\xb2\x89\xf9\x65\x3a\xe4\xf8\x0f\x49\x8b\x10\xa4\x9c\x90\x00\x53\x0e\x4a\x45\x96\x94\x65\x90\xea\xa8\xd1\x32\x61\x79\x4f\x21\x11\x29\xa4\x43\x58\xf7\x0d\x3b\x94\x24\x3e\x57\x54\x55\xa5\x1e\x6c\x7e\x25\xbf\xbc\x79\x63\x38\x5b\x61\x2c\x24\x7c\xe5\x39\x95\xe5\x9a\x66\x4d\xa9\x08\x8d\x12\x3f\x5a\x0e\xd1\xdf\x06\xa2\x3f\x45\x76\x77\x6d\x86\x4d\x84\xb4\x77\x77\x55\xd1\xb6\xae\x4d\xcc\xed\xb5\xc8\x07\xfc\xb3\x0c\x83\xa3\x8b\x8b\x53\xf2\x32\x9d\x91\x97\x65\x30\xd9\x56\xd0\x3d\xd0\xfe\x8c\x9c\xad\xe8\x52\x81\xd3\xd5\x38\xf2\x2d\x3e\xda\xe5\x47\x54\xbd\xd1\xdc\x58\xd2\xdc\xd0\xd5\xbf\xb1\xfd\xdc\xd0\x4c\xb0\x72\xe8\x39\x02\xdb\x78\x90\x4b\x9a\xc0\xa6\xc6\x04\x8a\xc3\x81\xa7\xa2\x2e\xfc\x58\xa4\xd6\x16\xe8\x49\xa1\x6d\x31\xec\x16\x2d\x46\xdf\xe9\xe6\x55\x03\x74\xb7\x1f\x7e\x56\x7f\xe2\x5a\xa1\xff\x46\xa7\x12\x19\xaf\x69\x43\x51\xa5\x20\x2f\x14\x4a\xf7\xc6\xf7\xbc\xa5\x90\xa4\x79\xf4\xab\x96\xce\x48\x1f\xbf\x35\x0f\x4b\x07\x79\xf6\xd4\xc1\x81\x6f\xe2\xa2\x67\x0e\x1b\xbb\x63\xda\xb5\x9a\xfc\x27\x7a\x0e\x14\x8c\x3a\x59\x33\x96\x07\xc3\xe8\x3f\x14\x55\x96\x12\x2e\x14\x49\x68\x96\x11\xeb\x25\xd7\x81\x36\xf1\x8f\xbf\x31\x99\x69\xa2\x2a\x9a\x91\x4e\xc8\x34\x94\x9c\xaa\x64\x6d\xda\x76\xaf\x5b\x9b\xf5\x73\xeb\xf8\xcf\xe6\x73\x53\x9b\x3c\x73\x9b\x31\x94\x89\xfb\x4f\xa0\xf4\xa1\xdf\x69\x56\x99\xfc\x8e\x35\x3e\xde\x2b\x8b\x8c\xe7\x90\x41\xa2\x0c\x82\xdb\x52\xf6\x36\xcb\xce\x41\x29\xec\x53\xc2\xa8\x97\x13\xe3\xb6\x78\x8a\x31\x56\xa0\x48\x23\xf9\x2d\xca\x62\x0c\x61\x2d\xe1\x69\x52\x57\x6e\x2d\xb4\xa9\x30\x17\x08\x84\x5a\xbe\xcb\xab\xc5\x83\x02\x9d\x4e\x1f\xee\x0b\x48\x14\xa4\xe4\x2f\x62\x4a\x0e\x09\x5e\xde\x60\xb6\x45\x13\x6b\xd3\xef\x91\xf7\x0f\x2b\xa1\xb1\x3d\x22\x56\x25\x9d\xa4\xae\x46\x1a\xaa\xbd\xcb\xf5\x41\x1a\x91\x70\xc4\xb6\x6f\xb9\xd2\xb9\xc5\xae\xc1\x34\x93\xb4\x24\xa1\x1c\xed\x23\x81\x26\x6b\x92\x42\x89\xc1\x49\x4a\x7d\xd5\x02\x12\x5a\x95\x40\x5e\x96\x84\x95\x06\xfa\x06\x2e\xdb\x6f\x0b\x27\x62\x67\x42\xf7\x3c\x6f\x21\x81\x5e\xb7\x34\x57\x8d\xbd\xba\xdf\x1a\xe1\x37\xc5\x72\x88\xcf\x33\x80\x22\x34\x13\x7b\x46\xb1\x22\xbf\x32\xcf\x21\x11\x3c\x75\x88\x8b\x90\x69\xb3\xfc\xef\xf3\xbd\x69\xde\x37\xc9\x09\xdc\x85\xc1\x67\x7a\xcf\xf2\x2a\x6f\x6e\x28\x09\xdc\x27\x00\x69\xb7\xfa\xb5\x65\x62\x0b\x15\xb7\xc6\xda\x33\x58\xb1\x12\x91\xbe\xec\xcf\xbf\x7a\xe0\x91\x42\xa8\xa6\x2a\x9c\x09\xa1\x7c\x0f\xe7\x03\xae\x9f\x05\x84\x04\x66\xba\x2f\xfb\x23\x8e\x7e\x69\x4e\x7e\xd4\x2b\xa9\xf8\xd0\x50\xb4\x26\x5f\x4b\x98\xf5\x46\x1e\x33\x91\xea\x81\xd1\x10\xe2\x0b\xdb\x0c\x1a\xca\x6f\x82\xaf\x66\x36\xc6\xe5\x75\x2a\xee\x78\x38\xba\x0b\x99\xf8\xae\x27\x19\x8e\x5d\x73\xa2\x64\x05\x7e\xb7\x88\x36\xfa\xd8\x59\x75\xbe\xc5\xbb\x7b\x02\x45\x70\x79\xb6\x4f\x06\xdf\x33\x9b\x17\xdd\x92\xf5\xb6\x2e\xe8\x58\xb4\xd8\x4e\x8b\xe0\x81\xbe\x29\xf0\x7d\xa2\x2b\x01\x49\x40\x2a\xca\x38\x81\x5b\x34\xbc\x90\x2e\xe0\xb1\xcf\x22\xc6\xcd\x8c\xaf\xba\x06\x0b\xde\x65\x22\xb9\xc6\xa8\x80\xa4\xd2\x02\xa2\x1d\xaa\x12\x4a\x52\x08\xd3\x6a\x28\x41\x0a\x90\x4c\xa4\x0c\xa1\xf7\x81\x24\x6b\x48\xae\xbf\x83\x63\x6d\x1d\x8e\x4d\xa5\x55\x2c\x44\x75\xb6\x86\xa4\x1d\x05\xd8\x33\x25\xd8\xae\x82\x9a\x65\x50\x3b\xea\x62\xb0\x9b\xc4\x4c\xf2\x74\x87\x09\x3b\x61\x15\x7f\x2d\xdb\xd8\x71\x13\x41\xfc\x36\x63\x14\x75\x77\x39\x6d\x1f\xcc\xc8\x65\x6f\xd1\xe2\xf5\x26\x9a\xc1\x5b\x9e\xa7\x79\x74\x18\x6c\x77\xe8\xcd\x76\x65\x84\xd0\x8d\x71\xb7\x43\xb3\x67\x77\x45\xb9\x0e\xbd\x26\xbc\x51\x6c\xb9\x2a\x67\xc4\x58\xe0\x33\xe3\x88\x00\x27\xf8\x0c\xc1\x26\x03\xbe\xb7\x74\x37\x77\x9c\x55\x7c\x46\xd0\xe8\x21\x5a\xf4\x55\xcf\x9c\x13\x42\xe5\xaa\x74\x46\x69\x9c\xd2\x6d\x85\x9f\xd8\x2c\xbd\xb8\xef\x0d\xc5\x7b\xe4\x42\x8e\x97\x78\xeb\x3d\xa9\xeb\xab\x61\x4f\x31\xd2\x56\x7b\x9e\x67\x16\xe1\xf1\x47\xaa\x68\xb6\x0c\x03\x5d\x26\x74\xfb\xc0\xd0\x92\x7a\x9e\x02\x29\xa3\xc6\x27\x7e\xd7\x35\xe6\x1b\x1e\xee\xc6\xac\xdd\x55\x1a\x48\x47\x50\xcd\x74\x81\x6a\x16\xe1\xad\x0a\xcd\x30\x16\x46\xfd\x35\x5a\xb7\x12\xb4\x76\x90\x50\x8a\x4a\x26\x60\x77\x03\xe6\x8b\xdb\x9c\x77\xe8\x3d\xc4\xda\x83\x15\xf8\x13\x6c\xbd\xd9\x83\x0e\x42\x82\xf7\x40\x33\x72\xc7\xd4\x9a\x6c\x1d\x24\xcd\xe7\x78\x67\xee\xee\x14\x2a\xda\xa5\xe2\xde\x85\x61\x3f\xbf\xc7\xd3\xbb\x69\xe2\xe0\x9e\xe6\x45\x06\xa5\x6d\xa0\xfd\x7e\x2b\x07\xf7\xfa\xfe\x0f\xcd\x21\x9b\x58\xee\xa5\x03\x53\x91\xc8\x81\x2b\x56\x8d\xf1\x70\x76\x09\x23\x72\x40\x02\x6d\x0f\x27\xaf\x45\x0b\xfd\x10\x30\xfc\xfe\xc9\x83\x61\x4d\xdf\x03\x3c\x3b\x70\x67\x17\xec\xec\x44\x9d\xbd\xa0\x33\xc0\x9c\x6d\x64\xa9\x27\x23\xbb\x82\x7d\x78\xf3\x44\xb8\x69\xd4\x38\x62\x69\x0a\xdc\xb1\x33\x5f\x67\xba\x99\x72\xa4\x51\x11\xac\xab\x66\xce\xb1\xe6\xd4\xa3\x28\xb6\x0b\xbb\xbe\x07\xba\x1a\x25\x86\x2b\x12\xcf\x5b\x88\xf4\x61\xd2\x9d\xf1\x3f\x81\xc2\x03\xe1\x70\x0b\x62\xae\xbf\x1c\x97\xb0\xae\x67\x57\x16\x6a\x46\x3b\xe8\x6d\xb4\xfa\xca\xe9\x22\x03\xac\xc0\xd8\xe8\xa3\x14\x03\xc8\xaa\x07\x45\xa4\xdb\x40\xc7\x27\x00\x69\xd9\x2c\x22\x48\x5d\xe3\xa8\xd2\x36\xae\xff\x72\x01\xfc\xb4\x71\xf0\x71\xfc\x7e\x2e\x6a\xef\xd9\xc3\xb7\x0b\xf8\x67\x58\xcc\xe0\x7b\x62\xb7\x24\x2e\x81\x47\xcc\x66\xfe\xfe\x59\x0a\x7e\x06\xa5\xd3\x1a\xbf\x63\xb0\x97\x6b\x9a\x1d\xeb\x06\x36\x74\x86\x09\x82\x09\x31\x4d\xed\x53\x24\x6a\x06\x92\x76\x58\xab\x9a\x0d\x8e\xdb\xde\xec\x14\xcb\xde\x71\xcc\x97\x62\x19\x06\x7a\x46\xb1\x92\x46\x7d\xff\xf6\x5b\x24\xb3\xbb\x3b\x75\xd8\x39\x36\x79\xfa\x0e\x1c\x46\xd7\xd2\x1d\xd0\x1b\x5d\x50\xf7\x78\x5c\x06\x83\xdd\x7a\x70\x45\xe6\x2e\x05\x5e\x60\x7f\x7d\xd5\x32\xec\x45\xe8\x9e\x41\x7d\x34\x9e\xc7\xff\x4f\x69\xfd\xd0\xae\x8b\x1f\xfd\x67\x65\xdb\xac\x78\x77\xbb\xcb\x4c\x9f\xa7\x8e\xf2\x52\x0f\x88\xc2\xf4\xa6\x88\x43\xdd\x86\xef\xf9\x3b\x82\xf6\x3e\x0d\x1b\xce\xa8\x5b\xab\x81\xdd\x81\x36\x8c\xfe\x3d\x8b\x81\x61\xaf\xd3\x46\x1b\x8a\x2f\x57\x4d\x92\xa2\x7a\xc3\xa5\x73\x77\xfa\x1f\xf3\x4f\xff\x3f\x3b\xe9\xb6\xb9\x06\x31\xf4\xbf\x6e\x30\x1b\x3a\xf1\x39\x28\x5d\x05\x5e\xb0\xd4\xc2\xff\xa3\x86\xda\xdf\x23\xdb\x80\xdd\x05\xb5\x63\x41\xd8\x71\x5e\x07\x60\xef\xf6\xf6\xc5\xdb\x96\xb0\x7b\x13\xbd\x3a\x18\x51\xbd\xee\x37\xc8\x5b\xf5\xdb\x54\xf5\x66\xcd\xd1\x83\x90\x91\xfe\x70\xb4\x99\x1e\xef\xa5\xed\x2a\x37\xec\xb5\x94\xb5\xff\xef\x00\x00\x00\xff\xff\x1f\x43\x1b\xaa\x5d\x24\x00\x00")

func bindataTemplatesCommandsTmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesCommandsTmpl,
		"templates/commands.tmpl",
	)
}



func bindataTemplatesCommandsTmpl() (*asset, error) {
	bytes, err := bindataTemplatesCommandsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/commands.tmpl",
		size: 9309,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1667459835, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x41\x4b\xc3\x40\x10\x85\xcf\x3b\xbf\x62\xc8\x41\x12\xb0\x9b\x7a\xed\xad\x68\x0e\x5e\xac\x88\x78\x5f\x37\x93\xed\x60\x76\x66\xd9\x6c\x4a\x25\xe4\xbf\x4b\x2a\xe2\xed\xbd\xf7\xf1\x1e\x2f\x39\xff\xe5\x02\x61\x74\x2c\x00\x1c\x93\xe6\x82\x35\x98\x2a\x70\x39\xcf\x9f\xd6\x6b\x6c\x7b\x27\x4c\x63\x28\xee\x7b\xd4\xdc\x6a\x22\x71\x89\x77\x7e\xe4\x5d\x20\xa1\xec\x8a\xe6\xd6\x8f\x5c\x41\x03\x30\xcc\xe2\x6f\x63\x75\x83\x0b\x18\x3f\xb2\x7d\x16\x2e\xf5\xdd\xa6\x1e\x55\x06\x0e\x0b\x18\x73\x4c\xe9\xc5\x45\x3a\x20\x62\xb5\x2c\x68\x37\x83\xeb\x5a\xdd\x83\x31\x9d\x5c\x5e\x33\x0d\x7c\x3d\xfc\xb3\x4e\x2e\x7f\xf8\x83\xf2\xc4\x2a\xb7\xea\x83\xdd\xdb\xfd\x96\xae\x0d\x80\x69\x5b\x7c\x3f\x3d\x9d\x0e\x78\xec\x7b\xcc\x14\x78\x2a\x94\xd1\x6b\x8c\x4e\xfa\x09\xcf\x94\xc9\xc2\xef\xa7\x37\xd5\x62\xbb\x2b\xf9\xb9\x50\xdd\xc0\x0a\x3f\x01\x00\x00\xff\xff\xd7\x90\x9c\xb4\x08\x01\x00\x00")

func bindataTemplatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesMainTmpl,
		"templates/main.tmpl",
	)
}



func bindataTemplatesMainTmpl() (*asset, error) {
	bytes, err := bindataTemplatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/main.tmpl",
		size: 264,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1664358933, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"templates/commands.tmpl": bindataTemplatesCommandsTmpl,
	"templates/main.tmpl":     bindataTemplatesMainTmpl,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"templates": {Func: nil, Children: map[string]*bintree{
		"commands.tmpl": {Func: bindataTemplatesCommandsTmpl, Children: map[string]*bintree{}},
		"main.tmpl": {Func: bindataTemplatesMainTmpl, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
