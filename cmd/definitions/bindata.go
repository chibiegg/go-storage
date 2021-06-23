// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cmd/definitions/tmpl/function.tmpl (567B)
// cmd/definitions/tmpl/info.tmpl (1.702kB)
// cmd/definitions/tmpl/object.tmpl (2.177kB)
// cmd/definitions/tmpl/operation.tmpl (1.698kB)
// cmd/definitions/tmpl/pair.tmpl (502B)
// cmd/definitions/tmpl/service.tmpl (12.729kB)

// +build tools

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x8f\x25\x87\xae\xec\xe6\x07\x08\x9e\x8a\x82\xb0\xc8\xa2\xde\x25\x64\xa7\x6b\xb0\x99\x94\x66\x5a\x17\x62\xfe\xbb\xa4\x55\x17\x11\x0f\xde\x3c\x25\xcc\xcc\x7b\xf3\x3e\x26\xa5\x2d\x06\xc3\x47\x82\x7a\xda\x40\x4d\xb8\xbc\x82\xd2\x37\x23\xdb\x88\x9c\xab\xd2\x76\x2d\x38\x08\xd4\xa4\x6f\x7d\xdf\x91\x27\x16\x3a\x7c\x36\x55\xcb\x2f\xb3\x66\xd2\x77\xc6\x13\xde\x20\xa1\x31\x9e\xba\x65\xa0\x88\xd5\xa4\x77\xc1\x9a\xb9\xd2\x8e\x6c\x51\x47\x5c\xa4\x04\x75\x56\xec\x4d\x5c\x06\xd6\x48\xa9\x58\xe6\x5c\xa7\xa4\x26\xbd\x37\x83\xf1\x51\x3f\x0e\xce\xef\x4c\x14\xfd\x20\x83\xe3\xe3\x35\x1f\xe2\xab\x93\xe7\x26\x78\x6f\x72\x46\xe8\x05\xbd\x71\xc3\x2f\xa6\xa5\x5c\x62\x7e\xdf\xb4\x2c\xb8\xa7\x38\x76\x12\x3f\x8c\xe7\x00\x15\x00\xf4\x86\x9d\xad\x57\x05\xdc\x9d\xa9\x57\xeb\x6a\xa6\xa2\x2e\xd2\x1f\x71\xac\x9c\x60\x03\x0b\x9d\x44\x37\xcb\xbb\xc1\x3f\x66\xdc\x82\xf8\xeb\xca\x3f\xbe\xef\x01\x00\x00\xff\xff\x41\xc8\xae\xbf\x37\x02\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0xfc, 0xb, 0xf4, 0xfb, 0xf, 0x94, 0xca, 0x64, 0x7a, 0xb8, 0x48, 0xf9, 0x69, 0xf1, 0x88, 0x46, 0x59, 0xd3, 0xbc, 0x0, 0x4e, 0xe7, 0x8, 0x1e, 0xdf, 0x27, 0x59, 0x81, 0x75, 0x52, 0xa2}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\x0e\x55\x85\x1a\xd4\x35\x9b\x84\x78\x80\xe5\x69\x1b\x68\x42\xdb\x90\x36\xf1\x82\x10\x72\x92\x9b\xca\x34\xb6\x23\xdb\x89\x56\x32\x7f\x77\xe4\x24\x74\x4d\xf7\x07\x8a\x78\xe1\xcd\x37\xf6\xbd\xe7\x77\x8f\xaf\x13\xc7\x38\xd1\x39\x61\x49\x8a\x0c\x77\x94\x23\x5d\x63\xa9\x37\x31\x32\x99\xc7\x39\x15\x42\x09\x27\xb4\xb2\xef\x70\x7a\x85\xcb\xab\x1b\x9c\x9d\x9e\xdf\x2c\x58\xc5\xb3\x15\x5f\x12\xdc\xba\x22\xcb\x98\x90\x95\x36\x0e\x33\x06\x00\x93\x42\xba\x09\x8b\x18\x6b\xdb\x03\x18\xae\x96\x84\xe9\x6a\x8e\xa9\x50\x85\xb6\x78\x9b\x60\x71\x1e\x56\x17\xbc\x82\xf7\xac\x6d\x31\xb5\x64\x1a\x91\xd1\x25\x97\x14\xf6\xa7\x2b\xdc\xc1\xe9\x13\x2e\xa9\x0c\x47\x58\x1c\xe3\xbd\xa0\x32\x87\x50\x39\xdd\x42\x28\xb4\xed\x76\x92\xf7\x48\x85\x63\x99\x56\x36\x40\xec\xe8\x36\x5d\xcd\x5e\xdd\xfb\x0e\x71\x37\xfd\x3c\xd4\x0d\x24\xcd\xa2\x83\x08\xf2\x9f\xb8\xcd\x78\xd0\x47\x82\xa3\xe3\xe3\xb0\xbb\xea\x81\x0f\x40\x2a\x0f\xcb\x88\xb1\x60\x00\x76\x7b\x18\xa7\x5b\x67\xea\xcc\xa1\x1d\x94\x37\x6c\xdf\x9e\x62\x0b\x18\x37\xeb\xaa\xaf\xe5\xfd\xd6\x97\xfb\x33\x1b\x86\x2e\x8e\xe3\x60\x00\x6a\x4b\x39\xb8\x05\x0f\x91\xe4\x15\x0a\x6d\xa0\xd3\xef\x94\x39\x34\xbc\xac\x69\x8e\x43\x48\xe2\xca\x42\x69\x07\x4b\x6e\x8e\xa3\xe1\x83\x25\xd7\x95\xea\xea\x08\xe5\xde\xbc\xee\x42\x09\xc9\xab\x2f\xd6\x19\xa1\x96\x5f\x85\x72\x64\x0a\x9e\x51\xeb\xd9\xa0\xfc\xbc\xd7\x61\x57\x14\x81\xfe\xec\xb6\x9b\x10\xef\x59\x51\xab\x0c\x33\x89\x57\xcf\xba\x16\xe1\x03\xb9\xbe\xf1\x53\x61\xab\x92\xaf\x07\x37\x66\xd1\xd8\x8f\xc1\x57\x43\xae\x36\x0a\x72\xf1\xc0\xbe\x40\xfa\xa7\x9a\xd7\x4f\x68\x36\x63\xcd\xe8\x37\x85\x06\xa6\x47\x60\x90\xa0\x19\xf1\xb2\x61\xa0\x4a\xdb\xb1\xfe\x03\x77\x66\x23\xd4\x39\x52\xad\xcb\x68\x20\x12\x05\xe4\x22\xdc\xf0\xcb\x3d\x9f\xc0\x8b\x04\x87\x43\x8d\xe7\xdd\x9e\xc3\x99\x9a\xba\x83\x7e\xbb\xd1\x2d\xa8\x3b\xfc\x20\xa3\x3f\x87\x79\xec\x12\x0a\x5e\x5a\xda\xe7\x96\x2e\x6a\xeb\xf6\x9b\x8e\xbf\xee\x3b\x19\xf7\x5d\x71\x25\xb2\x59\x21\xdd\xe2\xba\x32\x42\xb9\x62\x36\x79\x8c\xf5\x23\xa5\x3c\xbd\x7f\xb9\xbf\xee\x5e\x6c\x9e\xdd\x24\x8a\x1e\x5a\xf4\xdf\xcc\x6e\x6f\xe5\x5d\xb2\x9f\x97\x8f\x8e\x7d\xff\x0f\xdb\xf9\xa5\xdd\x2f\x7f\x06\x00\x00\xff\xff\x26\xf5\xd0\xb3\xa6\x06\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x9d, 0xdc, 0x29, 0x17, 0xc8, 0x8a, 0x14, 0x25, 0x13, 0xa5, 0x85, 0x13, 0x6, 0xf, 0xb3, 0x4a, 0x21, 0xd1, 0x20, 0x4e, 0x6f, 0x92, 0x60, 0xe2, 0x6e, 0x29, 0x96, 0x4a, 0xee, 0x99, 0x66}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x51\x73\xe3\x34\x10\xc7\x9f\xab\x4f\xf1\x27\xd3\x01\x9b\x49\xed\x3b\x60\x78\x28\x97\x07\xe6\x52\xe0\x66\x68\xc3\x4c\x0a\xef\x8a\xbc\x4e\x45\x6d\xc9\x23\xad\x4d\x42\xea\xef\xce\xc8\x76\x72\x4e\x9b\x83\xdc\x30\x7d\x93\xac\xdd\xbf\x76\x7f\xbb\x2b\xa7\x29\xde\xdb\x8c\xb0\x26\x43\x4e\x32\x65\x58\x6d\xb1\xb6\x87\x3d\x54\x99\xa5\x19\xe5\xda\x68\xd6\xd6\xf8\x1f\x30\x5f\xe0\x6e\x71\x8f\x9b\xf9\x87\xfb\x44\x54\x52\x3d\xca\x35\x81\xb7\x15\x79\x21\x74\x59\x59\xc7\x88\x04\x00\x4c\xf2\x92\x27\xfd\x8a\x75\x49\xc3\xd2\x6f\x8d\x9a\x88\x58\x88\x34\xc5\x4f\x9a\x8a\x0c\xda\x64\xb4\x81\x36\xb0\xab\x3f\x49\x31\x56\x9a\x85\xb2\xc6\x07\x9d\xdd\xee\x0a\x4e\x9a\x35\xe1\xf2\x71\x8a\xcb\x06\xd7\x33\x24\x8b\xce\xee\x96\x58\xa2\x6d\x3b\xd5\xde\xf3\x43\x10\xda\xed\x70\xd9\x24\x77\xb2\x24\x3c\x81\xed\x6f\xd2\x2b\x59\xa0\x6d\x51\x6b\xc3\xdf\x7f\x87\x19\xde\xbe\x7b\x17\x8c\x1e\x83\x73\xd0\x27\x93\x85\x65\x1f\x52\xaf\x0d\xed\xc1\x0f\x04\x5f\xca\xa2\x20\xcf\xa8\x8d\xe6\x10\xe2\xda\x5e\x79\xb6\x4e\xae\x29\x11\x69\x1a\x1c\xee\x16\xf7\x37\xcb\xeb\xb0\x02\xae\x06\xf7\xaf\x3c\xf2\x90\x9a\xc7\xf2\x97\xc5\xef\xbf\xce\x61\x2c\x63\x45\x50\x0f\x21\x95\x0c\xb6\x66\xaf\x33\x82\x27\xd7\x68\x45\x3e\x39\x76\xc7\xfb\x1f\xef\x02\xe2\xe0\x61\x2b\x4d\xd9\xb3\x63\xed\xa1\xac\x51\xb5\x73\x64\x18\x5e\xe6\x94\x88\x50\x80\xfd\xb9\x67\x57\x2b\xc6\xee\x5c\x7a\xc1\x4c\xe7\x01\xdb\x9c\xbc\x72\xba\x0a\x95\xfe\x78\xf8\xc9\x83\x03\xba\x8f\x76\xf7\xdb\x8a\x3a\xf6\x6d\x3b\xfa\xf2\x8c\xb4\xb8\x48\x53\xa8\x42\x87\xe8\x07\xd0\xfb\x9d\xc1\x5f\x0f\x5a\x3d\x8c\x32\x95\x85\x6e\x28\x11\x17\x83\xc5\xb2\xa7\xef\x44\x77\x69\x9a\x86\x6e\x41\xed\x29\x83\xf4\x90\x61\x57\xca\x0a\xb9\x75\xfb\x6e\x6a\x64\x51\xd3\x14\x6f\x50\x92\x34\xbe\xab\x84\x27\x9e\xe2\xed\xf0\xc1\x13\x77\x52\x9d\x4e\xd7\x22\xe2\x22\xb3\x86\xba\xcd\xb7\xdf\x88\x8b\x32\x9c\x86\xae\x4d\x6e\x6b\xa6\x8d\x68\x85\x38\x07\x6c\xc8\x5e\xfb\x39\x55\x8e\x54\x98\xa3\xeb\x59\xc8\xee\x19\xca\x49\x9a\xe2\x60\x92\x5d\x4f\xf6\xa0\xfa\x6a\xdc\x6c\xba\x69\x6a\x5b\x91\xd7\x46\x21\xb2\xf8\xba\xbf\x23\xc6\xcf\xc4\x43\x61\xb4\xaf\x0a\xb9\x1d\x98\x47\xf1\x31\x75\xec\xba\xdc\x1c\x71\xed\x0c\x6c\xf2\xa2\x48\x21\x9b\x17\xe2\xcb\x4f\x88\x37\xc7\xe2\xf1\xde\x63\xb8\xe5\x84\x3c\x66\x68\x8e\x22\x10\x43\x23\x14\xbe\xbb\x3d\x50\x0a\xb9\x8e\x41\xb5\xad\x38\xa6\x82\x5b\xdb\x50\x06\xb6\x18\x46\x0f\x25\xb1\x4c\xc6\x1d\xf5\x19\x7c\xa2\xa3\x1c\xa6\x58\x59\x5b\xc4\x87\x0c\x3c\x4b\x8e\xe2\xbe\xb9\x74\x0e\x9b\x84\xb6\xf8\xf2\xac\xf7\xe5\x8b\x19\xde\x0c\x3a\xff\xce\x7c\x0a\x76\x35\x75\x86\xad\x18\xd3\x19\x45\xf6\x84\xbf\xc9\xd9\x3f\x42\xef\x76\x1e\xb9\x2c\x3c\x89\x57\x04\x76\x5b\x7b\xfe\xbc\xa6\xfa\x5f\xb0\x66\xc7\xb0\x2a\x69\xb4\x8a\xf2\x92\x93\x65\xe5\xb4\xe1\x3c\x9a\x0c\xe3\x3b\x52\x68\xdb\xf0\x1e\x0c\x03\x3c\x89\xe3\x81\xe0\x7f\x37\xf8\x6b\x31\x7b\x85\x39\xe9\x19\x3e\xcd\xce\x81\x78\x72\xb0\xfa\x50\xc7\x8f\xed\x8b\xb0\x55\x61\x0d\x45\x9b\xd1\x97\xdd\xe1\x41\x3f\xe7\x5f\x7b\x32\xfc\xcd\x29\xf6\xcf\x7e\x13\xa3\x0c\x3b\x87\xf0\xa3\x6f\xff\x09\x00\x00\xff\xff\x74\x56\xed\x90\x81\x08\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfd, 0xb0, 0xe7, 0x29, 0x9a, 0x5f, 0xc1, 0x32, 0x7, 0xac, 0x52, 0x8e, 0x9c, 0x2b, 0x1f, 0xb5, 0x69, 0x93, 0x58, 0xd7, 0x7, 0x6f, 0x97, 0x44, 0xd0, 0x62, 0x7d, 0x78, 0xa3, 0x58, 0x7d, 0x5b}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x8f\xd3\x30\x10\xbd\xfb\x57\x8c\xac\x1e\x12\x69\x49\xef\x48\x7b\x82\x45\x5a\x09\x75\x2b\x10\xe2\x88\x5c\x67\xda\xb5\x48\x6c\x33\x9e\xec\x87\x82\xff\x3b\xb2\x93\x76\xd3\xaa\x10\x15\x10\xe2\x16\x7b\x66\xde\x9b\xf7\xc6\x13\xaf\xf4\x57\xb5\x43\xe0\x67\x8f\x41\x08\xd3\x7a\x47\x0c\x85\x00\x00\x90\xda\x59\xc6\x27\x96\xc3\xc9\x38\x29\x4a\x21\xfa\xfe\x15\x90\xb2\x3b\x84\xc5\x97\x2b\x58\x18\x78\x7d\x0d\xd5\xad\x65\xa4\xad\xd2\x18\x20\x46\xd1\xf7\xb0\x30\xd5\x5b\x0c\x9a\x8c\x67\xe3\x6c\xba\x4c\x0c\x30\x46\x4c\xf0\x8d\x7a\x5e\xa9\x16\x21\x46\x30\xfb\x62\xe8\x33\x53\x62\x30\x5b\x70\x04\x05\x7e\x4b\xf9\x39\x51\x06\xa4\x07\xa3\x91\x64\x79\x72\xcf\x8e\xd4\x2e\xdd\xc7\x98\xeb\x3f\x32\x19\xbb\x2b\x4a\x08\xf9\xe3\x80\x89\xb6\x4e\x8d\x8c\xe7\xa9\x08\xe7\x93\x8a\x85\xa9\xee\x7c\x16\x90\x32\x96\xcb\xdc\xad\xf3\x03\xcd\x77\x60\xb7\x56\x41\xab\x26\xb5\x3c\x46\x4e\x24\x8e\xc0\xe7\x6b\x8a\x31\xf2\xce\x51\xab\x78\xad\x48\xb5\x89\xab\x84\xe3\xc0\x07\x0c\x5d\xc3\xe1\xb3\xe1\xfb\xf5\x30\x9b\x23\x95\x32\x95\x4c\x5d\xb2\x8e\x73\xf5\x7b\x37\xd0\xcc\xf6\x9e\x90\xdf\x0c\x83\xfd\x1d\x19\x93\xf2\x42\xf3\x13\x8c\x6f\xa4\x1a\xef\xae\xfe\xba\xca\xfd\xd8\x8e\x4f\xf9\xd8\x76\x81\x6f\xda\x0d\xd6\x9f\xac\x69\x7d\x83\x2d\x5a\xc6\xfa\xdc\x1b\x2b\x4a\x11\x85\x58\x2e\x61\x36\x33\x83\xc2\x06\x01\x13\x70\x8d\x35\xb0\x83\x7b\xf5\x80\xb0\x75\xf4\xa8\xa8\x06\xed\x5a\xaf\xd8\x6c\x1a\x84\x03\x96\x4a\xde\x85\x6a\x78\xe4\xf3\x1c\x81\xa9\xd3\x0c\x7d\x14\x62\xdb\x59\x0d\x45\x98\x2f\x2a\x2f\x92\x7b\x29\xf6\xc9\xce\x8c\x9b\x48\xc8\x1d\x59\x90\xb3\x00\x32\xd9\xfb\xeb\x9d\xba\xa0\x99\x7f\xb2\x44\xc7\x3f\x9b\xb3\x6b\x84\x44\x70\x0d\x2b\x7c\xbc\xf3\x48\x79\xc4\x2b\xc7\xb7\x2f\xdd\xdf\x10\x39\x2a\xe4\xb4\xdb\x18\x65\x79\xfa\xc3\x79\x71\x52\x44\xf1\x33\xbe\x0b\xec\xf9\x0f\x76\x73\x34\xef\x4f\x0c\x9a\x38\xb2\x37\x6a\xfa\x75\x70\xef\x47\x00\x00\x00\xff\xff\xbd\xe8\x37\x32\xa2\x06\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xaf, 0x9a, 0xb8, 0xf7, 0x9f, 0x1b, 0x53, 0xaa, 0x36, 0x95, 0x6d, 0xe6, 0x56, 0x7d, 0xb9, 0x8e, 0xd7, 0x4, 0x74, 0xce, 0x5b, 0x9, 0x4f, 0xb6, 0x92, 0xc, 0x1e, 0x49, 0x5, 0x4c, 0x4c}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\xe3\x43\x02\x89\x75\xd9\x53\x96\x3d\x6d\xf6\xb0\x14\x92\x1c\x42\x7b\x2c\x8a\x3c\x51\x44\x6c\x49\xc8\x63\xb7\xc6\xf5\x7f\x2f\xb2\x93\x40\xe9\xa5\x3a\xcd\x7c\x33\xef\xf1\x46\x52\xe2\xaf\x2f\x09\x86\x1c\x45\xc5\x54\xe2\xd4\xc3\xf8\x47\x0f\x5d\x97\xb2\xa4\xb3\x75\x96\xad\x77\xcd\x6f\x6c\xf7\xd8\xed\x8f\xf8\xb7\xfd\x7f\x2c\x44\x50\xfa\xaa\x0c\x21\x28\x1b\x1b\x21\x6c\x1d\x7c\x64\x2c\x04\x00\x64\xda\x3b\xa6\x77\xce\xc4\xdc\x1a\xcb\x97\xf6\x54\x68\x5f\xcb\x13\xf5\xde\x95\x0d\xfb\xa8\x0c\x49\xe3\xd7\xf7\xb2\xfb\x25\xc3\xd5\xc8\x0b\x73\xd0\x95\x25\xc7\xd9\xa4\x2d\x7e\xac\xe6\x3e\x50\x93\x09\xb1\x14\x62\x18\xd6\x88\xca\x19\x42\xfe\xba\x42\xde\x61\xf3\x07\xc5\x21\x05\xc5\x38\x4e\xd3\x3c\x38\x55\x53\xe2\x79\x57\xec\x52\xf9\x01\xf6\x07\xd5\x68\x55\xa5\x1d\x29\xf1\x62\xf9\x32\x0c\xf7\xcd\x71\xc4\x9b\xad\x2a\xa8\x10\xaa\x1e\x89\xdf\x74\xe3\x88\x4e\x55\x2d\x81\x3d\xf6\x61\xfa\xa9\x42\x48\x29\xe6\x95\x2d\x35\x3a\xda\x09\x27\xdb\x73\xeb\xf4\x37\xe3\x45\x77\xf3\x3b\xf6\x21\xf5\x4b\xa4\xa8\x18\xa6\xfb\x23\x71\x1b\xdd\x44\x66\x90\xde\x13\xf5\x1b\x64\x5f\x42\x64\xab\xc7\xf4\x39\xc5\xd9\xa0\x9b\xc9\x28\xe6\x8b\xc9\x95\x29\xc1\x67\x00\x00\x00\xff\xff\x81\x9d\x53\x8d\xf6\x01\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x1a, 0xbf, 0xf0, 0xe9, 0x47, 0x35, 0xe6, 0xa7, 0xc, 0xe6, 0xba, 0xc0, 0x5e, 0x39, 0xda, 0x2d, 0xb9, 0xd0, 0xee, 0xbc, 0x1, 0xc0, 0xd1, 0x5e, 0xa2, 0xbf, 0x4d, 0xcd, 0x4, 0xe9, 0xc0}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdf\x6f\xdb\x38\xf2\x7f\xf7\x5f\x31\x2b\x04\x85\x5d\xb8\xd2\xf7\x0b\xdc\x53\x0e\x79\xd8\x4b\xbb\xbd\x60\xfb\x23\x68\xba\xb7\x0f\xdd\x45\xc0\x48\x23\x9b\x17\x89\x54\x49\xda\x69\xe0\xea\x7f\x3f\x0c\x49\xfd\x96\xfc\xa3\xdb\xbb\xed\x1d\x9a\x97\xd8\xe4\x70\x38\x33\x9c\xf9\xcc\x0c\xcd\x28\x82\x4b\x99\x20\xac\x50\xa0\x62\x06\x13\xb8\x7b\x84\x95\xac\xbf\xc3\x96\x33\x88\xf3\x24\x4a\x30\xe5\x82\x1b\x2e\x85\xfe\x2b\x3c\x7f\x0b\x6f\xde\xbe\x87\x17\xcf\xaf\xde\x87\xb3\x82\xc5\xf7\x6c\x85\xb0\xdb\x41\xf8\x86\xe5\x08\x65\x39\x9b\xf1\xbc\x90\xca\xc0\x7c\x06\x00\x10\xc4\x52\x18\xfc\x64\x02\xf7\x8d\xcb\x60\xe6\x3e\xad\xb8\x59\x6f\xee\xc2\x58\xe6\xd1\x1d\x3e\x4a\x91\x68\x23\x15\x5b\x61\xb4\x92\xcf\xaa\x8f\xdb\xbf\x44\xc5\xfd\x2a\x42\x91\x14\x92\x8b\x8a\xc7\x09\x2b\x63\x85\x09\x0a\xc3\x59\x76\xfa\xda\xb5\x31\x45\x9c\x71\x3c\x75\x5f\x8d\x6a\xcb\x63\xd4\x6e\x55\x78\xf4\x3a\xf3\x58\xd0\xa2\xc5\x6c\xb6\x65\x0a\x6e\xa1\x91\x3d\xbc\x56\x72\xcb\x13\x54\x7e\xa6\xb2\x47\xf8\x0f\x96\x6d\xd0\x0f\xde\x38\x4e\x15\x4d\x25\x45\x78\xe3\x3e\xbc\x50\x4a\x56\x73\x8d\x66\xe1\xdb\xc2\x1e\xeb\x6c\x16\x45\xf0\xfe\xb1\x40\xe0\x1a\xcc\x1a\x81\x84\x81\x54\xaa\xce\xc9\xc6\x52\x68\xe3\xc8\x2e\x20\x68\xcd\x04\x76\xfd\xdb\xbb\x7f\x62\x6c\x5e\xa3\x61\x09\x33\x0c\x48\x35\xd4\x95\x20\x90\x57\xe3\xc4\x55\x5a\xd2\x70\x66\xb7\x19\xac\x53\x9b\xd8\xc0\x6e\xb6\xdb\x3d\x03\xc5\xc4\x0a\xe1\xec\x76\x09\x67\x5b\x38\xbf\x80\xf0\x4a\xa4\x52\x93\x30\x64\x5c\xa2\xe0\x29\xe0\x47\x38\xdb\x86\x37\xb1\x2c\x10\x02\xc7\x3a\xe8\x91\x08\x69\x88\xe6\x65\x26\xef\x58\xd6\x9e\x3b\x2b\x04\xe9\x70\x7e\x41\xd3\x56\x9d\xcf\x60\xe4\x35\xd3\x71\x97\x8e\xa7\x44\xf0\x9c\xeb\x22\x63\x8f\x95\x41\xc0\xff\xb5\x18\x5d\x4c\x90\x11\x09\x8a\xa4\xf9\x6a\xe9\x50\xc7\x8a\xdb\x23\x68\x4f\x38\x4e\x65\xe9\xa9\xac\xc1\xa7\xd8\xd4\x5f\x5b\x1f\x4b\x7b\x1c\x2f\xd1\xf4\x2c\xfb\xc0\xb3\x0c\x56\x68\xfa\x16\x4f\x95\xcc\xfd\x58\x38\x8b\x22\x5a\xfc\x0c\xde\xaf\xb9\x86\x74\x23\x62\x2b\x9d\x5e\xcb\x4d\x96\x58\x33\xde\x21\xc4\x2c\xcb\x1c\x5c\x54\xa7\xcb\xf3\x22\xc3\x1c\x85\x41\x15\x56\xeb\x11\x14\x9a\x8d\x12\x5c\xac\xfa\x3b\x72\x0d\x0a\x59\x02\x52\x64\x8f\xc0\x44\xd2\xe3\x9f\xcb\x84\xa7\x1c\x93\x70\x46\x02\x0c\x35\x99\x4b\x78\xea\x46\x16\x7d\xce\x3b\x6b\x17\x99\x2f\x41\xde\xd3\xa9\xca\xf0\x25\x1a\x1f\x03\xf5\xf2\x85\x25\xe2\x29\xd1\xec\xea\x53\x74\xd2\x82\xcc\xc3\x79\x97\xa9\x23\x77\x16\xf7\x44\x5d\x82\x5d\x65\x73\x3d\x6e\x73\x3d\xb4\x39\x17\x46\x1e\x67\x73\x6b\xa3\xc6\xe8\x52\xc4\xb8\x84\x22\x43\xa6\x11\x72\x76\x8f\xa0\x37\x0a\x81\x65\x19\x58\xc6\x6b\xa6\xe1\x0e\x51\xc0\x83\xe2\xc6\xa0\x80\x3b\x4c\xa5\x42\x92\xc1\x9b\x73\x20\x64\x63\xce\x25\xd4\x9e\x50\x2b\x5f\x99\x34\xbc\x19\x1a\x52\xe6\x0b\xaf\xb9\x87\x9f\x9b\x47\x6d\x30\x1f\x60\x80\x1d\xed\x42\x80\x47\x3e\x3b\xe8\x81\x60\x8a\xc7\x1f\xc2\x03\xbf\x4f\xd0\x8f\xd7\x29\x50\x18\x89\xe9\xbd\xe0\xd0\xe3\x39\x0d\x12\x27\x80\x45\x9b\xb4\x15\xed\xcd\xf0\x04\x78\xb4\x09\xf6\x80\xc8\xe9\x40\x32\x7e\x32\x35\x9e\xf4\xc6\x2d\x9e\xf8\x25\xaf\xed\xf1\x36\x0e\xde\x06\x85\x71\xae\x27\x63\xc3\x28\x9b\xb9\x86\xa7\x2d\x11\x16\x13\x9b\x39\xdf\xd6\x35\x5c\x68\x0b\x17\x5d\x4e\x7b\xd1\x42\xe7\xe1\x7c\x94\xf5\x08\x68\x8c\xd2\xb5\xb1\x63\x8f\x99\xf5\xd0\xcc\x16\x42\x26\xcc\xfc\x1f\xc2\x91\x63\x4c\xbf\x04\x9d\x8f\xab\x5e\x41\x8b\xb6\xd0\xd2\xe3\xe1\x90\x65\x34\xe6\xaf\x19\x57\xba\x72\xd5\xb1\x38\x3e\x26\x76\x9b\x78\xbd\x94\x22\xcd\x78\x6c\x68\x34\x8a\xe0\x39\x16\x0a\x63\xaa\x87\xcf\xe1\x17\x8d\x50\xd0\x6e\xe1\xaf\xdc\xac\x3b\x61\xc5\x85\x36\xc8\x12\x6b\xf4\x56\xc8\x44\x11\x0c\x48\xed\x01\xb2\xa2\xc8\x1e\x7d\x30\xfa\x58\x87\x2d\xd5\x6f\x40\x89\xc0\x95\x61\x9e\xd9\x48\x74\x5b\x83\xf7\x19\xcf\xb7\xdd\xe0\x5e\x00\x59\x66\xe8\xa3\x34\xda\x0c\xfe\x8c\x8f\xe7\xb6\x7a\x6b\x04\x09\x96\xf5\xac\x2d\x29\xcf\x61\xbb\xf4\xee\xdb\xc1\x83\xd6\x47\x5b\x4a\x92\x69\x5e\xb3\x02\x2e\x20\x67\xc5\x07\x6d\x14\x17\xab\xdf\xdd\x3f\x2f\xc5\xfe\xf3\x3b\x05\x67\x7b\x12\xd7\x2a\x78\xdd\xbd\x0a\x5d\xf0\x1a\xdd\x9d\x38\xe8\x82\xc5\xd8\x11\xc1\x60\x5e\x64\xd4\xf4\x04\x9c\x8a\x98\x94\xe6\x03\x1b\x17\x37\x19\x15\x38\xd5\xd6\x67\xdb\xf0\xaa\x26\x18\x67\x90\x22\x33\x1b\x35\xb9\xfc\xa7\x8d\x88\x27\x56\x92\x3d\x6f\x05\x3e\xb4\x57\xce\xc7\x6c\xb2\xb0\x83\xf8\x30\xcd\xe6\x0b\x76\xaf\x00\xe3\xe0\xd2\xb6\x1b\xd0\x67\xdb\x1f\xf6\x2c\xd7\x44\x21\x19\x9d\x8b\x04\x3f\x41\x08\xff\x57\x8f\x5b\x5a\xdd\x9e\xfb\x7f\x9a\x23\xaf\x9a\x4f\x7a\x4e\xb5\xca\x8b\x7f\xdb\x0e\xa7\x8e\xc7\xc0\x05\x3c\x71\x91\xd2\x1d\xdf\x0d\xb2\xdd\x62\x52\x9b\xe6\x18\xa7\x74\x19\x02\xca\x59\x2a\x86\x3a\xb9\xe2\xc6\x8b\x53\x96\x3f\x79\xbe\x4d\x61\x43\x22\xbd\x92\x52\xe3\xdb\x82\x1a\x6f\x2e\xc5\x8f\x59\x06\x77\x52\x66\xd3\x96\xa0\x8d\xbc\x19\xba\x4b\x27\x4d\xd2\xe1\x57\x29\x6c\x43\x9e\x2b\xb3\x61\xd9\x51\x9b\xa7\xa2\xbf\x7b\x85\xa3\xa9\x08\x6f\x78\xbe\xc9\xec\x45\x82\x9f\xeb\x73\x26\xd9\x52\x71\xac\x70\x7b\x64\x25\x04\xf9\x22\x31\x33\x14\x56\x04\xcf\xa7\x2b\xa7\xc5\xc9\x29\x11\x07\xdb\xf5\xa4\x1c\x39\xa1\x6a\x17\x38\x61\x9b\x93\x8e\x6f\x80\x78\x53\xbe\xdc\x00\xcb\xa1\xb8\x4c\xc5\xd0\x7f\xfd\xc4\x7d\xa5\xd6\x08\x42\x53\xde\x2b\xbc\x5e\xd6\xcb\x9d\x82\xf7\x36\x53\xba\x2b\x85\x82\x29\x8d\x89\xf7\x7a\x17\x13\xbd\x15\xb4\xa0\x2c\xbb\x71\x61\xb1\x0c\x3e\xfc\x4e\x46\x73\x2e\x10\x45\xf0\x0e\x3f\x6e\xb8\xc2\xc4\xcd\x8e\x1d\x01\x4d\x54\xe2\xd6\xd4\xde\x62\x7f\x67\xda\x6e\xca\xb8\x1a\x33\x34\xb4\x4d\xbd\x8f\xac\x99\xde\xd3\xa1\x47\x91\x4f\xee\x2c\x3b\x4e\xda\x9a\xfa\x6b\x4b\x7b\x58\xde\xea\x18\x95\xc6\xeb\x89\xb3\xb4\xa5\x8c\xa5\x70\xb5\x86\xb6\x49\xc2\x56\xa2\x4f\x27\xce\xdf\xd5\x2e\xfb\xb8\xce\x65\x61\xaa\x33\x5e\xc0\x7c\x82\xcf\x12\x50\x29\xa9\xaa\xa2\x51\xa1\xde\x64\x86\xac\x36\x41\xdf\x54\x3c\xd6\xee\xe7\x40\xbb\x54\x65\x8d\xfd\x47\x6d\xe8\xed\x12\x6c\xb8\xba\xd3\xb0\x82\x34\x0b\xf5\x03\x37\xf1\x1a\xb6\xe1\xcf\xf8\xd8\x1a\x1e\xf7\xc0\x13\xbd\x90\xfe\x62\xaa\xbd\x83\xee\xb9\x51\x6d\xd3\x69\xe6\x78\xea\x75\x0d\x0f\xb8\xc2\x6e\xd0\x6a\xc6\x52\x18\x2e\x36\xd8\x99\xe8\xb6\x8a\xc7\xb1\xbe\x00\xa3\x7a\x6c\xfc\xc2\xfd\xab\xb6\xee\x8a\x32\x9c\xf7\x5d\x6f\x31\xd5\x88\x4e\xc7\xcc\x89\x71\xf3\xdd\xbe\x93\xf6\x2d\xdb\x41\x70\x82\xcb\xf2\x14\x7e\x38\xd5\x54\xbe\x05\x99\x8a\xd1\x72\xd9\xdc\x54\x53\xf8\x57\x5b\xda\xeb\xea\x1d\xf5\x2a\xfa\xfc\x83\x6f\x29\x76\x23\x47\x09\x65\xad\x0d\x8c\x94\x0b\x7e\x77\x27\xf4\x12\x04\xcf\x0e\x24\x49\x1d\xf4\xfa\x92\x7d\x15\x5f\x4d\x66\x1b\xe6\x41\xda\xf4\xae\xfc\x1c\x53\xb6\xc9\x4c\xad\xbd\xeb\x7e\xb8\xa6\x7d\x69\xc2\xa7\x38\x7b\x27\x56\x60\xcc\x53\x1e\x03\xb3\x85\xb8\xe5\x60\x13\xe5\x38\x8f\x4e\xa2\x1c\xad\x41\xac\x60\x9d\x5b\x9f\x89\xea\xc2\x27\xd8\x11\x87\x19\x77\x93\x29\xf6\xfd\x36\xa4\xd3\x4a\x14\x02\xce\xb6\x7b\x2f\x9c\x06\xe7\xf1\x6d\x17\x2c\xf5\x8a\x03\x25\xcb\x9f\x51\x9c\x1c\x57\x9e\x7c\x9b\x85\xc8\x01\xc9\x7a\x65\xf5\xff\x48\x85\x34\xd7\xf0\xb4\x9e\x5c\xfc\x57\xd6\x4b\x51\x04\x5c\xff\x22\xf4\xa6\x28\xa4\x32\x98\x58\x03\x28\x8c\xa5\x4a\x34\x3c\xac\xd1\xac\x51\x41\xbc\x51\x0a\x85\x83\x3d\x0a\xb5\x4d\x43\x1f\xd6\x9c\x86\x6c\xce\x2f\x20\x65\x99\xc6\xd9\xa1\xea\xec\x7b\x0d\x76\x72\x8d\x30\x2a\xe7\x48\xe1\xf0\xbd\xfc\xfa\x33\x4c\xdb\x83\xbb\x6f\xc4\xb2\x51\x04\x57\x29\x6c\x34\x2a\x40\xc1\xee\x32\x74\x49\x73\xeb\x85\xb5\x0a\xf8\x2b\xb4\x25\x3c\x20\xc4\x8c\xaa\x40\x3d\x42\x65\xb1\x91\x9b\xb0\x2f\xbc\x0e\xab\x2b\xb8\xb0\x77\xed\xf3\xf9\xf3\xc4\x64\x83\x66\x47\x92\x9d\x66\x96\x2f\xf6\xa6\xaf\xed\x51\x47\x9e\xd1\x10\x45\x7b\xb2\x8d\x38\xa2\xaf\x4a\xcf\x4f\x63\x54\x36\xb0\x4c\x5d\xc2\x90\x7e\xb7\x3f\x22\x5a\xeb\x27\x1c\x2b\x93\x52\x53\xc2\xf1\x57\x89\x1d\xdf\xb2\x39\x97\xaf\x84\x54\x2e\xe9\xb6\x36\xb7\x8d\x44\x38\x1b\x77\xab\xe1\xa5\x6b\xd7\x65\x86\x37\xab\x3e\xc9\x1f\xd4\xa6\x39\xf4\x93\xda\x9f\xbe\xe0\x3b\x1a\x3c\xdf\x76\x2a\xf1\x28\x82\xcb\x35\xc6\xf7\xa0\x3a\x77\x10\xe1\xf7\x6e\xae\x7b\x7d\xdf\xfc\x9a\xf1\x55\x3b\xba\xa3\x1b\x21\xf7\xe3\xe2\x2b\xd9\x7b\x2b\xd0\xed\x91\x32\x9a\xbe\xad\x64\x9d\xec\x96\x6a\x6b\x64\x7a\xf0\x2b\xfe\xf0\x07\x9c\xc3\x6c\x0e\xfd\xf0\x3f\x6a\xc9\x2f\xed\xc5\x8e\x69\xc5\xda\xb1\xe5\xae\xc6\x7b\x3f\x80\xba\x9f\xb6\xbb\x3f\x6c\xdb\xb0\x8f\x15\x92\xf2\x0c\xfc\x43\x43\xb8\x7b\xac\x30\x2c\x6c\x17\xd6\x67\x85\x28\xcb\x45\x6b\xa3\xb9\xbd\x70\x0e\xaf\x99\x62\xb9\x0e\x6f\xac\x53\x12\x85\x1f\x7f\x67\xbd\xac\x3d\xe1\xc2\x20\x36\x9f\x48\x17\xbf\x5b\xf8\x37\x16\xdf\xaf\x94\xdc\x88\xc4\xbf\x13\xa8\xde\x05\x84\xcd\x4e\xbf\x72\xb3\xbe\x74\xf4\xf3\xd8\x7c\x5a\x42\x67\xe7\x4b\x96\x65\xa8\x08\xe1\xfb\xa6\x68\xad\x9b\xb0\xca\x3e\xfd\x7a\xbb\xd6\x12\xfb\xb1\x9e\x14\x47\xeb\x9f\x60\x8a\xca\x9e\xc1\x7c\xd1\x2d\xba\xcf\x0a\x66\xd6\xd5\x39\x7b\xb6\xd7\xcc\xac\x9d\x82\x23\xe1\xc1\x44\x02\x73\xfc\xe8\x17\x06\xc1\xc2\x7f\x13\x10\xf8\x87\x46\xc1\x62\xf4\x69\x0b\x91\x5f\x40\xb0\xfc\x2d\xf8\x2d\x18\x3c\xec\xe9\xe5\x33\x54\x94\xac\x74\x98\x4a\x95\x33\x63\x91\x68\x1e\x38\x15\xc9\x17\xcb\x32\xb0\x8d\x52\xc3\xb8\x2c\xc1\xbf\xdd\x98\x2f\x9a\x90\x9f\xfa\x89\xcb\x65\xab\xaa\x9e\x19\xa4\x27\x6a\x70\x84\x34\x3e\x93\x25\x4d\xb6\x62\xb1\x81\x8c\xdf\x23\x18\x72\xe9\x66\x9d\xa7\x6f\xde\xd3\xf9\xa6\x88\xd0\x7a\x58\xcc\x74\x72\xd7\x93\x27\x7b\x49\x46\xd3\x97\xb3\xce\x1b\x7c\xa8\xc9\xde\x48\x73\xd5\x6c\x5e\xdb\xab\x09\xde\xb2\x0c\x16\xbd\x84\xb0\x07\xb3\xab\x27\x1a\xe8\xba\x04\xf7\xb4\x4c\x26\x08\x41\x17\x97\x73\x1a\xab\x7a\x89\x86\x6a\x08\xcf\x64\x08\x19\xd2\x64\x78\x65\xb3\x96\x5d\x59\x96\xf3\xc5\x40\xad\x3a\x17\x35\x0c\xaf\xc4\x96\x65\xdc\x27\xa4\x17\x9f\x0a\x8c\xed\x3b\x0f\x9a\x6a\xf1\x5a\xc2\x8f\x31\xd9\xee\x1c\xdc\x4e\xfd\x7c\xbe\x47\x5f\x77\xf9\x73\x01\xac\x28\x50\x24\xb6\x37\xd7\x4b\xd0\xa1\x87\x24\x7b\x71\xd7\x82\x86\x30\x0c\x9d\x31\xb7\x4c\x51\x37\x3d\x95\x5a\x1d\x73\x59\x98\x65\xed\xd0\x7b\x6f\x08\xec\xbe\xf5\xc3\x25\x5a\xf2\xc3\x05\x25\xcc\x41\x2e\x6f\x57\x16\x3d\xe8\x6a\xb0\xfa\x92\xe5\x48\x07\x30\x82\x5e\xef\x15\xcf\x5f\x31\x6d\x3c\x8c\xbd\x10\x09\xb5\xe5\xeb\x4b\x99\xe7\xac\x2c\x49\xe2\xc5\x9e\x24\xdd\x4f\x7f\xfb\x32\x75\x7b\x6e\x2c\xd1\x54\xe8\x70\x20\xd9\xf8\xa0\x3d\x98\x70\x1c\x5d\x45\x7e\x6a\xe2\xa1\x65\xff\x86\xe4\xf3\x47\x3c\xec\x68\x2f\xab\x60\xcd\xd5\xd3\xf6\x12\x09\x1e\xd6\x3c\x43\x58\x33\x91\x64\x5c\xac\xc0\x9e\x1b\x29\xe8\x9f\x2d\x55\xcb\xac\x83\xde\x1e\xed\x9e\xfd\xba\xd2\xca\x3d\xe1\x7a\x27\x7b\x9d\x73\xec\xb6\xe7\xd9\x03\xe1\x82\x9b\x1a\x2b\x8e\x7c\x20\x44\x7f\x35\x98\xbc\xc3\x15\xd7\x06\xd5\xd4\x5d\xba\x9a\x53\xcb\xb6\x24\x50\x9d\x24\x59\x8c\xdd\x78\x0e\x76\xb8\x89\xd7\x98\x33\xcf\xce\xbf\xb5\xa2\x68\xfa\x57\x00\x00\x00\xff\xff\xcc\x47\xcf\xa7\xb9\x31\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xab, 0x4a, 0xd7, 0x84, 0x83, 0xf1, 0x2, 0xa7, 0xf3, 0x4b, 0x37, 0x44, 0x61, 0x62, 0x16, 0x3f, 0x1f, 0x1b, 0xe6, 0xe6, 0x63, 0x12, 0x6b, 0xde, 0x3b, 0x87, 0x86, 0x5a, 0x4a, 0xce, 0x7b}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cmd": {nil, map[string]*bintree{
		"definitions": {nil, map[string]*bintree{
			"tmpl": {nil, map[string]*bintree{
				"function.tmpl": {cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl": {cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl": {cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": {cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl": {cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl": {cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
