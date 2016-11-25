// Code generated by go-bindata.
// sources:
// templates/hello.tmpl
// templates/homepage.tmpl
// templates/main.tmpl
// templates/partials/footer.tmpl
// templates/partials/header.tmpl
// templates/productPage.tmpl
// DO NOT EDIT!

package assets

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

var _templatesHelloTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x30\xb4\xab\xae\x56\xd0\x73\x2f\x4a\x4d\x2d\xc9\xcc\x4b\xaf\xad\xb5\xd1\xcf\x30\xb4\x03\x04\x00\x00\xff\xff\xcd\xf4\xc7\xe6\x17\x00\x00\x00")

func templatesHelloTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHelloTmpl,
		"templates/hello.tmpl",
	)
}

func templatesHelloTmpl() (*asset, error) {
	bytes, err := templatesHelloTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/hello.tmpl", size: 23, mode: os.FileMode(420), modTime: time.Unix(1480070731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHomepageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x73\xdb\x36\xf6\x7d\xef\xa7\xb8\x83\x97\xdf\x4c\x66\x28\x5a\x8e\x7f\x69\xb3\x23\x33\x93\xc6\x4d\x9b\x26\x6b\x67\xe3\x38\xdd\xf6\xc5\x03\x91\x57\x24\x6a\x10\xc0\x02\xa0\x64\x45\xe3\xef\xbe\x03\x90\x94\x28\x8a\x54\xec\xd0\x4d\xdd\x6d\xfc\x62\x11\x7f\x0e\x2e\x2e\xce\xc1\xbd\x00\xf9\xcd\x24\x61\x73\x88\x39\x35\xe6\x98\x28\x9a\x62\xc0\x84\xd5\x12\xa6\x34\xbe\x4a\xb5\x2c\x44\x12\x04\xd4\x58\x4d\x39\x89\xbe\x01\x00\x68\xb6\x5f\x68\xaa\x14\xea\xaa\xa6\x5d\x1b\x4b\x1e\xb8\x16\x8d\xea\x9e\x66\x1d\x2d\x7c\xab\x6c\xbc\x6b\xd9\xe5\xa5\x65\x96\x23\x89\x7e\x41\x1e\xcb\x1c\xc1\x4a\xb0\x19\xc2\xd9\x6c\xc6\x62\x84\x99\xd4\x70\x4a\x2d\x93\x82\x72\x38\xb7\xd4\x32\x63\x59\x6c\x26\x61\x36\xee\x19\x44\x75\x8d\x11\x4b\x61\x51\x58\x12\xbd\xcf\x10\x2e\x5e\xff\x9f\x01\x4e\x75\x8a\xc6\x02\x13\x09\x2a\x14\x09\x0a\x0b\x4a\xcb\xa4\x88\x51\x83\x9c\x81\x74\xe3\x33\xca\xc1\xac\x07\x05\x2a\x12\x6f\x9b\xc6\x58\xa6\x82\x19\x4c\x40\xd4\xb6\xad\x9b\x51\x0e\x4c\x18\xcb\x6c\x61\xd1\x01\x59\x3f\xe2\x68\x12\xaa\x0e\xbf\x85\x09\x9b\x37\xbc\xbd\x79\xac\x7e\x56\xff\xb6\x56\xb5\xb9\x94\x29\xe5\x1c\xf5\x72\xef\x5a\x6e\xe0\xb3\xc3\xe8\x0d\xb5\x6e\xd6\xaa\x98\x72\x66\x32\x4c\x40\x23\x47\x6a\xd0\x79\xf4\xb0\x61\x49\xc1\x81\x25\xc7\xa4\xae\x0d\x38\x33\x96\xd4\xd8\xee\x21\x08\x04\x16\x8e\x47\x50\xf3\x02\x72\xaa\x53\x26\x02\x8e\x33\x1b\x04\x07\xf5\xe3\x54\x5a\x2b\xf3\x20\x38\x68\xb1\x62\xb5\xd2\x54\xa4\x08\xa3\x13\x6a\xe9\xe8\x5d\x35\xd0\xcd\x4d\xcb\x49\x13\xce\x1a\xcc\xf2\x83\x05\x3c\x0d\x0e\x9f\xd6\x96\xb7\xc7\x39\xac\x0b\x34\x4b\x33\x1b\x04\xe3\xb6\x5d\x4d\xff\x2d\x32\x66\x11\x14\x4d\x12\x26\xd2\xc0\x4a\xe5\xda\xd7\x8f\x6b\x80\xba\xa0\x1e\x62\x53\x52\x62\x8e\x3b\x09\x3f\xc9\x1e\xd7\x96\x57\xa6\xae\xb9\x3e\xa1\x90\x69\x9c\x1d\x93\xd5\x0a\x46\x17\xef\x5e\xdd\xdc\x90\xc8\xfd\x7c\xef\xaa\x6f\x6e\x26\x21\x8d\x26\x61\xf6\xb8\x13\x55\xed\x80\xae\xc9\xed\x20\x2a\x47\x9e\x50\xeb\x81\x76\x48\x37\x09\x39\x6b\x2f\x04\x8a\xa4\xe1\xf7\x49\x58\xf0\x16\x6b\x7e\x42\x9a\x70\x26\x10\x66\x2c\x2d\x74\x1f\x59\xb2\xaa\x55\x50\xb5\xba\x1b\x5f\xf6\xd2\xa3\x36\xe0\x65\x89\x7c\x5b\x96\x8c\x8f\x60\x6d\x54\x42\x2d\x7d\x28\xd4\x58\xad\xd8\x0c\x46\xe7\x8a\xea\x2b\x67\x9b\x9b\xe2\xce\x94\xaa\x89\x35\x24\xbd\x35\x95\xcb\x4b\x53\x77\x6f\x6a\x7c\x5f\xff\x75\x07\xf8\xdd\x04\x81\xc9\xe4\x82\x80\x83\x0a\xd6\x15\x8e\x90\x95\xdb\x5b\xc6\xc1\x6a\x15\xc0\xe8\x57\xf0\xbf\x3c\x5f\x7a\x36\xf9\xfa\x6f\xdd\x1f\x32\xd4\xd8\x6f\x5f\xb5\xc1\xf5\xd5\xaf\x56\xe1\x23\x38\x3d\x83\x9f\xcf\xe1\x51\xd8\xe3\x23\x8f\x23\xa4\x89\x35\x53\x76\xbf\x55\x13\x96\xa7\x3b\xee\x08\x58\x4e\xd3\xd2\x29\x19\x4b\x90\x80\xd1\x71\x43\x9a\xe1\xba\xa1\x6f\xb7\x67\xde\x93\x70\xbf\x15\xfd\x93\x6d\x6b\xb0\xd1\x67\xb3\x8b\xb4\xd6\xff\x3e\xf7\x92\x16\x74\x29\xe0\xde\x99\x96\xf4\x2d\x43\x49\xa9\xc8\xd1\x5b\x8d\x17\x82\xd9\x9b\x1b\x37\x74\x77\xcd\x9e\x49\x96\xa0\xad\x8e\xe5\xbf\x3d\xed\x77\x8c\xe8\xb1\xa0\x2e\xee\x75\x71\x57\x68\x6e\x7a\x67\xc6\x0b\x93\xdd\xd7\xf6\x0a\x77\xdf\xae\xfc\xe6\xf3\x40\xe2\xda\x67\xf1\xf1\x8d\x94\x57\x4c\xa4\x2e\xb7\x4b\xe4\x42\x70\x49\x13\xb0\x2c\x47\x30\xa8\x19\x9a\x67\xfd\x24\xdd\x59\x84\xf7\x7a\x09\xb2\xd0\xcd\xee\x80\xd7\x8a\x4b\x8d\x7a\x67\x3d\xb6\x57\xa3\x23\xb0\x9d\xd9\x0c\x35\x9c\x9d\x9e\xfb\xdc\xae\x23\xe7\xeb\x8e\x74\xd2\x75\xbb\x87\x9c\xe8\x49\xcb\xc1\x2d\x5e\x64\xe8\x57\xca\xb8\xac\xe6\xf1\x36\x4b\x7c\x45\x9e\xf8\x8a\xea\x89\xa7\x41\x70\xf4\x5d\x07\x0d\x9a\x2c\x3a\xe8\xb1\xa8\x4d\xaa\x75\x36\xf5\x89\x6c\xbf\x45\x98\x7e\x8e\x55\x1c\x6e\x53\xae\x0b\xbf\x3d\xc6\x54\x5e\xaf\x93\x1c\x68\x3e\x04\x41\x26\x73\x74\x79\xfe\xd6\x30\xce\x11\xe3\x6f\xb7\xb9\xec\x7c\x38\x3e\xd8\x2e\x73\xee\x1b\xff\x7f\xd3\x99\xe3\xa7\x5b\xce\x3c\xda\x6e\xef\xca\xda\x69\x4a\x8f\x34\x3c\x41\x2a\x49\xec\xd8\x76\xd8\x10\x49\x98\x21\x57\x21\x97\x31\xe5\x1b\xd2\x6d\x04\xe3\x8e\x3f\xbe\xb2\x41\xc9\x67\x13\xa3\xa8\xa8\x47\xf2\x31\x29\xc8\x24\x4f\x50\xc3\x82\x25\x36\xf3\x73\x7d\x5a\xfd\xf6\xb3\x1a\x57\x0f\xde\x31\x6e\x74\x17\x05\x7d\x90\x1b\x8d\xc2\xd1\x28\x64\x79\x1a\xda\x71\xe0\x07\x1a\x29\x91\x12\xa0\xdc\x1e\x93\x35\xb7\x85\x0c\xa6\x52\x27\xee\x34\x31\x71\xd1\x50\x44\x5b\x26\x94\x4b\xc2\x59\x7c\x45\xa7\x7e\x07\xa8\xda\xec\x89\x3c\x6d\x71\x37\x28\xea\x8d\x24\xd1\x73\xc8\xa8\x48\x96\x90\x16\x2c\xf1\x87\x42\x8e\x16\x96\xb2\x80\x2b\x21\x17\xb0\x70\x49\x85\x2b\x9d\x31\x91\xec\x78\xa8\xfb\xb0\x05\xbb\x07\xae\x3d\xc5\xbb\x1b\xf9\xad\xc4\x79\xd4\x16\x67\xaf\x1c\x6b\x01\xee\xd1\xe8\xff\x8a\x28\xef\xa4\xc0\xbb\x0b\xac\x29\xa7\x18\x85\x29\x0c\x89\x5e\xf8\xff\xf7\x44\xd3\xbb\xea\xed\x71\xaf\xc4\x4a\xf3\x02\x2e\x53\xd9\x10\x5a\x69\x2c\xb8\xd2\x3d\x9a\xbb\x85\x8c\x2c\xe3\x68\x1a\x07\xc2\x13\x66\x62\x39\x47\x0d\x99\x5c\xf8\xa8\x59\x8e\xdf\xbc\xd3\x70\xfb\x0f\x28\xca\x84\x05\x0a\x8a\xc5\xb6\xd0\xeb\x6b\x8b\xf2\x72\xc3\x87\x46\x07\xb0\x40\xe0\x6c\x8e\x5f\xd5\xf5\x77\x55\x57\x66\xad\xfa\x47\x18\xce\x99\x29\x28\x1f\x49\x61\x46\xa9\x9c\x8f\x8a\x2b\x02\x96\xea\x14\xed\x31\xb9\x9c\x72\x2a\xae\x48\xf4\xa1\x6c\x72\x76\x7a\xfe\xe0\x24\x58\x5a\x1f\x38\xeb\x37\x0a\xfc\xe1\x9a\xe6\x8a\x7b\xe2\x33\x61\x51\xd3\xd8\xb2\x39\xfa\xd3\x31\x94\x1d\x98\xf1\x62\xd8\xd5\x67\x35\x2a\x2d\xac\xac\x5d\x1d\x1c\x3e\x19\xa2\xda\x8d\xf3\x80\x19\xa0\xb0\xc0\xa9\x71\x2c\x2e\xd3\x5c\x97\x16\x08\x5c\x00\x55\x4a\x4b\x1a\x67\x68\x5c\x1c\xcc\xa9\xcf\x17\x5c\x9f\xe6\x85\x65\x1c\xa3\x31\x6c\xca\xd1\x8b\x58\x23\xc7\x39\x15\xd6\x75\xa0\xce\x6e\x2c\x2f\x02\x63\xa0\x45\xc2\x50\xc4\x5f\x46\xdb\x5e\xcd\x79\x12\x1c\x7d\xdb\xbc\xce\x6b\x28\xfe\xe8\xbb\xdb\xca\xfc\xcf\xd2\x70\xf7\xb5\xce\x9d\x35\x3c\x2c\xd6\x61\x2c\x85\xcc\x97\x21\x13\x33\x5e\x5e\x43\x8b\x44\x69\x16\x23\x13\x09\x8b\xd1\x84\xd3\x82\x73\xb4\x4c\x98\x30\x96\xc2\x14\x39\xea\xaa\xba\x6a\x1f\x72\x7f\x4c\x26\xd1\x89\xe3\xf9\x79\xec\x29\x00\x2f\x68\xae\xee\x2d\x74\x76\x65\x78\x5e\xa0\x9f\xb8\x3e\xfa\xbe\x60\xdc\x79\x1c\x16\x52\xf3\x24\xe0\x48\xfd\x13\x5e\x2b\xd4\x96\x19\x04\x26\x7c\x88\x62\x42\xc8\x39\xf5\x62\xa5\x4a\x71\x16\x97\x21\x4b\xce\x4a\xed\x9a\xf8\xcb\xd3\xba\xf3\xa8\x56\x05\xb2\xcd\x0e\xf1\x95\xd4\xdd\xa4\x56\x28\x15\x47\x25\x55\xb1\xe6\x74\x2c\xf3\xbc\x10\xcc\x2e\xc3\xad\xe2\x9c\xa5\xba\xa4\xf1\xa6\x18\x8d\x65\xb9\x23\x75\x83\xfb\x54\x88\x82\xf2\x9c\x25\x4b\xa4\xba\xab\x69\xad\x82\x1f\x51\xa6\x9a\xaa\x6c\x59\xbd\x08\xb2\x7f\xb8\x08\xa0\x75\xee\x39\xf0\x9b\x3f\xb3\x3e\x57\x93\x0a\x05\xa4\x1b\x9b\xa4\xb6\x94\x83\x3b\x10\x32\x71\x85\x49\x45\x70\x66\xf1\xcb\x9c\x77\xbe\xd2\xfb\xaf\x4e\xef\xb3\x0f\xa3\x8b\xd7\x7f\x02\xa5\xcb\x4b\xb6\xae\x17\xaa\xd5\x0b\xd7\x04\x58\xae\xa8\xb6\x8c\x72\xbe\xf4\x99\xca\x4c\x23\xc2\x4c\xcb\x1c\x94\xe4\xac\x7e\xab\x3a\xe3\xc5\x1f\xb3\x9b\xb7\x6e\x06\x57\xab\xf0\xd1\xf6\x3d\xe1\x4b\xa4\xee\x60\x94\x40\x45\x87\xee\x6b\xc1\x59\xd5\xea\x33\x6e\x06\xf7\x5f\x04\x36\xd2\xa4\x07\xf0\xb6\x73\xa3\xf2\xa3\xee\x0b\xe2\xce\x3b\xe0\x8b\xd7\xf0\x16\xb5\x51\xe8\x53\x6b\x03\x87\x07\xe3\x27\xdd\x44\xbb\xdd\x56\x54\xee\x3f\x7f\x59\x77\x9c\x34\xd2\x13\x88\xcb\x94\xeb\xef\xeb\x8d\x57\x8d\xaf\x30\x34\xce\x19\x2e\x5c\x0a\x77\xf1\x1a\xca\x1c\x97\xc5\x5b\x57\xf2\x9f\x76\x53\xef\x55\xff\x1d\xaf\xf8\x5d\xab\xcf\xf9\xf8\x61\xef\xcb\x6c\x6f\xc9\x1d\x5e\x61\x3f\xf8\x55\xed\x5f\xd7\x5b\xbc\x88\xbc\xed\x07\x09\xf5\xd3\xfa\x35\xf0\xf6\x27\x32\x7d\x5f\xc8\xdc\xe5\x63\xa7\xec\x30\x7a\xee\xce\xc6\xbf\x39\xf2\x35\x3f\xe8\x59\x47\xdb\x5d\xae\x54\x68\x34\xf8\xe8\x69\xd2\x5e\xac\xb6\xbb\x1a\xab\x5c\x77\x09\x98\xc5\xbc\x99\x23\x50\x2b\x3f\x3e\xa3\x1f\x8f\x29\x89\x9e\x97\x5e\xdb\xb3\x03\x7c\x12\x65\x4a\xa2\xef\x87\xa3\xc4\x24\x7a\x31\x1c\x25\x21\xd1\xc9\x70\x14\x24\xd1\x0f\xc3\x51\x66\x24\x7a\x39\x1c\x25\x25\xd1\x8f\xc3\x51\x32\x12\xfd\x34\x1c\x85\x91\xe8\xd5\x70\x94\xdf\x49\xf4\xf3\x70\x94\x2b\x12\xbd\x1e\x8e\xc2\x49\xf4\x66\x38\x4a\x4e\xa2\x7f\x0e\x47\x11\x24\x3a\x1d\x8e\x22\x49\x74\x36\x1c\x45\x91\xe8\xed\x70\x94\xff\x90\xe8\x5f\xc3\x51\x34\x89\xde\x0d\x47\x31\x24\x3a\x1f\x8e\x62\x49\xf4\x7e\x38\x4a\x41\xa2\x8b\xe1\x28\x73\x12\x7d\x18\x8e\xb2\x20\xd1\x2f\xc3\x51\xae\x49\xf4\xef\xe1\x28\x4b\x12\xfd\x3a\x1c\xe5\x23\x89\x7e\xdb\x45\xd9\xc4\xf8\xed\xb0\xfe\xdf\x00\x00\x00\xff\xff\xfb\x0f\xef\xf8\xd0\x2c\x00\x00")

func templatesHomepageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomepageTmpl,
		"templates/homepage.tmpl",
	)
}

func templatesHomepageTmpl() (*asset, error) {
	bytes, err := templatesHomepageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/homepage.tmpl", size: 11472, mode: os.FileMode(420), modTime: time.Unix(1479996327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x97\xe1\x6e\xe3\x36\x12\xc7\x3f\x37\x4f\x31\x55\x81\x93\xbd\x11\x25\xab\xb9\xe4\xbc\x1b\x2b\x40\xb6\x1b\xf4\xf6\xae\xbb\x59\xdc\xa6\xb8\x3b\x04\x01\x4a\x4b\x23\x89\x29\x45\xaa\xe4\x28\xae\xa1\xe8\xdd\x0b\x4a\x76\xec\x24\x6e\x36\x45\xb0\x5f\x24\x4b\x9c\xf9\xfd\x67\x86\x43\x8a\x9e\x7d\xfb\xee\xfc\x87\x8b\xff\x7f\x3a\x83\x92\x2a\x79\xb2\x37\x5b\xdf\x90\x67\x27\x7b\x00\x00\x33\x12\x24\xf1\xa4\x6d\x21\xfc\x80\xc4\x33\x4e\x3c\xbc\x70\xaf\xa0\xeb\x80\xc1\x79\x9e\x8b\x14\x21\xd7\x06\x3e\x72\x12\x5a\x71\x09\x9f\x89\x93\xb0\x24\x52\x3b\x8b\x06\xef\x81\x54\x21\x71\x50\xbc\xc2\xc4\xcb\xd0\xa6\x46\xd4\xce\xc1\x83\x54\x2b\x42\x45\x89\x77\x4f\xe4\xdd\xc6\xa4\xeb\xbc\xc7\x88\x5f\x71\xb9\xd0\x26\xb3\xf7\xfc\x0d\x57\x05\x6e\x41\xfe\xbd\x32\xea\x3a\x68\x5b\x06\x61\xd7\x05\xd0\xb6\xa8\xb2\x07\xc8\x92\xa8\x66\xf8\x5b\x23\x6e\x12\xef\x7f\xec\xe7\x53\xf6\x83\xae\x6a\x4e\x62\x2e\x71\x8b\xff\xfe\x2c\xc1\xac\xc0\x20\x2d\x8d\xae\x30\x89\x3d\x88\xb6\x21\x69\xc9\x8d\x45\x4a\xbc\x86\x72\x36\xf5\xee\x8f\xad\x19\x0b\x91\x51\x99\x64\x78\x23\x52\x64\xfd\x43\x20\x94\x20\xc1\x25\xb3\x29\x97\x98\xc4\xe1\x24\x68\x2c\x9a\xfe\x91\xcf\x65\xaf\x33\xa4\x7c\x23\x70\x51\x6b\x43\x3b\xaa\x91\x6b\x53\x71\x62\x19\x12\xa6\x0f\xaa\x4a\x28\xb1\x2e\xb5\xc2\x44\xe9\x1d\x9e\x54\x62\x85\x2c\xd5\x52\x9b\x2d\xa7\xef\x0e\xa7\x87\xaf\x0f\xdf\xee\xb0\xe7\x75\x2d\x91\x55\x7a\x2e\x24\xb2\x05\xce\x19\xaf\x6b\x66\x89\x53\x63\xd9\x9c\x1b\x66\x69\x79\xaf\x68\x1b\xd2\xde\x37\xdf\xb4\xed\x2f\xb3\x6f\x19\xbb\x14\x39\x48\x42\x78\x7f\x06\xd3\xab\x93\x5f\xe0\x16\x2c\xcf\xf1\x9f\x17\x1f\x7e\xea\xba\x5e\x10\x60\x26\x85\xfa\x15\x0c\xca\xc4\xeb\x89\xb6\x44\x24\x0f\x4a\x83\xb9\x9b\xe8\xf0\x13\x27\x42\xa3\x7e\x12\x73\xc3\xcd\xf2\xd4\x5a\x24\xfb\x89\x53\xd9\x75\x51\x6a\x6d\xa4\x65\xc6\x04\x86\xa9\xb5\xde\xc9\x5a\xf7\x12\x55\x26\xf2\x2b\xc6\x76\x29\x6e\x45\xf6\xfe\x0c\x5e\x7f\x9d\xa8\x04\xb2\xd7\xab\x98\xee\x34\xff\x34\xaa\x7b\xd5\x2a\x68\x15\x96\x7b\xf1\x55\x62\xab\xb8\x50\x0f\x63\x63\xec\x89\xf8\x66\xd1\xb0\x4b\xcc\xe6\x3a\x5b\x42\x2a\xb9\xb5\xc3\x12\xbe\x58\xd6\x6e\x77\x70\x33\x3e\x90\x80\xb0\xaa\x25\x27\x04\xaf\xe6\xc6\xf5\xba\xed\x7d\xd1\x78\x10\x42\xd7\xed\xad\x9a\x8c\x0b\x05\x22\x4b\x3c\xf7\xc3\x03\xa3\x25\xae\x7f\x13\x9f\x0b\x95\xe1\xef\x89\xc7\xe2\x55\x80\x03\x7a\x29\x50\x66\xab\x12\xcc\xfa\x1c\x9e\x54\xcd\xb5\x26\xa7\xba\xd6\x7c\x54\xe1\xe9\x53\x15\x6e\xdb\xe8\x15\x5c\x9c\xbf\x3b\x07\x06\x4d\x9d\x39\xb4\x56\x20\xc5\x0d\xc2\xbc\x11\x32\x03\xd2\xd0\x58\x04\x5a\x68\xf8\xd7\x67\xc8\x85\x44\x1b\x80\xd2\x04\x1c\xac\x50\x85\x74\xf6\x08\xaf\xa2\xcd\x9c\x0d\x7b\x1c\xd0\xb2\x76\x4b\x11\x7f\xa7\xe8\x9a\xdf\xf0\xe1\xad\x07\xd6\xa4\x5f\x9a\xba\xeb\xd5\xcc\x5d\x5b\xef\x64\x16\x0d\x8e\x27\x7f\x81\xee\x00\xbc\xae\x77\xf8\x7f\xb9\x03\xf6\xb6\x44\xd6\x9a\xa3\xbc\x51\xfd\x0e\x34\x12\x81\x0d\x74\x50\x04\x26\xe0\x41\x35\x6e\xc5\xa5\xff\xa3\xd6\x85\xc4\x53\xc5\xe5\xd2\x7d\x1d\xce\xe7\xd7\x98\x92\x7f\x95\x98\x63\x71\x69\xae\x12\x77\xb9\xbd\xbd\xf3\x1f\xb7\x6b\xa4\x1b\x08\x7f\x4b\x86\xdb\xed\xed\xe5\xd5\x38\xac\x1b\x5b\x8e\xb8\x29\x9a\x0a\x15\xd9\x71\x17\xf4\x83\x32\x89\x5f\x29\x5c\xc0\x3b\x4e\x38\x1a\x1f\xf3\xc4\x86\xa9\x41\x4e\x78\x26\xd1\x19\x8e\xf4\x38\x58\x41\xab\xc4\x86\x05\xd2\x6a\xc0\xbe\x5d\x5e\xf0\xe2\x23\xaf\x70\xa4\xc7\x97\x93\xab\x63\x1e\x72\xbb\x54\x69\x12\x1f\xf3\xd0\xd5\xa9\x38\xae\xc2\x9a\x1b\x54\xf4\x51\x67\x18\x0a\x65\xd1\xd0\x5b\xcc\xb5\xc1\x91\x4b\x6f\x45\xed\xc6\xa3\x85\x50\x99\x5e\x04\x99\x4e\xfb\xd8\x02\x7f\xa8\x8f\x1f\xf8\x51\xb4\x58\x2c\xc2\xa2\x2f\x02\xe3\xeb\x2a\x84\xa9\xae\xa2\xcd\xd3\xb5\xf5\x03\xbf\xe0\xfe\xf8\x78\x85\x2c\xf8\xc8\x1f\x92\xf0\x03\xf0\x7f\x3e\x65\x87\x47\xd3\xd7\xdf\x4f\x0e\xfe\xc1\x0e\xfc\x00\x5a\x9f\x4b\xa9\x17\xa7\x2a\x2d\xb5\xf1\xdf\x00\x99\x06\xbb\x7b\xbe\x16\x55\xe6\x3c\x6b\x5e\xa0\xfb\x80\xf4\x4e\xee\xc1\x7f\x03\x52\xa7\xfd\x47\x3b\xac\x39\x95\x6e\x7b\x87\x7d\x28\x90\x3e\x23\x37\x69\x39\x1a\xc3\xfe\xc6\xa2\xe4\xb6\xdc\x80\xd7\x93\xb4\x6d\xdd\xde\x2d\xcb\xfb\x2b\xe5\xbf\x25\x2a\xb0\xbd\x11\x08\x0b\x95\xbe\xc1\x7e\xa5\x18\x54\x19\x1a\x34\xb0\x40\x5f\x4a\x50\x38\xbc\xe6\x59\xb6\xb6\x26\x34\x15\x08\x45\x1a\x5c\xbc\xf0\xe3\x29\x18\xb4\xb5\x56\x16\xb7\xa4\xa2\x08\x44\x3e\x7a\x9c\x49\x92\x24\xe0\x47\x03\xc9\xbf\x17\x5c\x14\xf5\xb7\x1b\x6e\x40\x35\xd5\x1c\xcd\x79\xfe\x1f\xb4\x8d\x24\x0b\x09\xb4\xed\x77\x22\x77\x3a\x8d\xa4\xf0\xc1\x70\x7f\x8a\x78\x62\x08\xa5\xc5\xae\x83\x89\x4b\x5f\xe4\x5d\x77\xfc\x58\xd4\x20\x35\x46\x6d\xca\xba\xca\x74\x1f\xfc\xbf\x3d\x20\x26\x3e\xec\x3f\x8c\xef\x01\xb0\x03\xa7\x08\x3b\x72\xdb\x2d\xf3\xc8\x7d\x6b\x27\xda\x78\xad\x5b\x79\x6f\xef\xce\xf2\x03\x72\xdb\x18\x04\x12\x55\xbf\xe5\xf5\xd3\xc1\x20\x35\x98\x09\x72\x93\xe6\xce\x4f\x6f\xa2\xa8\x44\x59\x87\x77\xbd\xec\xce\x4a\x7d\x77\x0f\x2d\x7f\xf7\x3e\xaa\x7a\x9c\x50\x05\x73\x40\xa6\x15\x9b\xeb\x46\xa5\xc8\x1c\x37\x7a\xd8\x62\xce\xc6\xc4\xf1\x68\xdc\x6e\xb7\x33\xde\xa0\x22\xf7\xe3\x42\x54\x78\xae\x3e\xb9\x86\x0e\xc0\x8f\xfb\x4b\xcc\x0e\x26\x60\x31\xd5\x2a\xb3\xae\xdf\xc1\x57\x5a\xbd\x57\x84\x86\xf7\x4c\xff\x0d\xc4\xd0\x8d\x8f\xbb\x9d\x5a\x07\xcf\xd4\xfa\xde\x5d\x0e\x62\x76\xf4\x02\xad\xa3\x67\x6a\xb9\xb5\xee\x1f\xc5\x2c\x9e\xbe\x40\x2c\x9e\x3e\x53\xed\xef\x7d\x15\xa7\x2e\xb5\x97\xe4\x36\x79\xa6\xdc\x61\x9f\xdc\xa4\xcf\xee\x45\xe9\x3d\x57\xf0\x68\xc8\x6f\x12\xef\xff\x15\xb5\x67\x80\x27\xfd\x85\xc5\xcf\xca\x62\x85\xb5\x48\x0e\xa2\x1b\x1a\xad\x5a\x3d\x88\xe3\xc9\x64\xf2\xa7\x06\x07\x71\x70\xf0\xa4\xc1\x51\x1c\x1c\x3d\x69\x10\x4f\xe3\x20\x9e\x3e\xcd\x98\xc4\xc1\xd1\xe4\x0b\x94\x89\xc3\x6c\x1b\x6d\x4e\x10\xb3\xc8\x1d\x0b\xdd\x7d\xf8\x4b\xf9\x47\x00\x00\x00\xff\xff\x7c\xc5\xca\xe2\x6a\x0e\x00\x00")

func templatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainTmpl,
		"templates/main.tmpl",
	)
}

func templatesMainTmpl() (*asset, error) {
	bytes, err := templatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.tmpl", size: 3690, mode: os.FileMode(420), modTime: time.Unix(1480070731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x08\x9e\x2b\x33\xd8\x00\x3d\x14\xb2\x80\x20\x68\xb6\x8b\x06\xc9\x21\x5b\xf4\x68\x8c\xc9\x91\x34\x35\x35\x14\xc8\x91\x14\xf7\xeb\x0b\x5a\xf1\x3a\xad\xbd\xe8\xc6\xd9\xf6\x62\xcb\x33\x7c\x6f\xc6\x7c\x8f\x43\x95\x75\x08\x82\x51\x59\x0f\x29\xad\x74\x1f\x89\xa5\x28\x5a\x72\xa8\xab\x45\xd9\x7e\x38\x24\x46\x4a\x03\x78\xbf\x6b\xc9\x39\x64\x5d\xdd\xcd\x30\x4f\xbc\x4d\xa5\x69\x3f\x54\x8b\xd2\xd1\x78\x58\x3d\x93\xea\xbf\x07\xa7\x08\x7d\x3f\x47\x19\xc6\x73\x80\x82\x61\x54\x36\xf8\x22\x2f\xfd\x07\xda\x06\xbf\x4f\x15\xbe\x29\x02\x63\x21\x2d\x45\x37\x47\x3a\x77\x8c\xec\xbb\xbe\x3e\xa5\x5d\xaf\x5b\x04\x47\xdc\xe8\xea\x17\xf4\x7d\x69\xda\xeb\x6a\x51\x0e\xfe\xdc\x4a\x4f\x49\x32\x8f\xa7\x73\x59\x12\xec\x72\x16\x54\x1b\xb1\x5e\x69\xd3\xa2\xef\x0d\x58\x8b\x29\xd1\x86\x3c\xc9\x4e\x57\x37\xaf\x7f\x96\x06\xaa\x45\x69\x3c\xbd\x91\xd3\x86\xb0\x25\x4c\xc0\xae\x8f\x34\x82\xdd\xe9\xea\x76\x0e\x29\x60\xa7\x5e\x82\x97\xb2\x0b\xc6\x2e\x73\xdb\xc0\x8e\x84\x02\x27\x5d\x7d\xce\xb1\x3d\xf9\x31\xfa\x9a\xdf\x0c\x3e\x7f\x3a\x1a\xff\x2b\x71\x6e\x36\x61\x10\xf5\xf8\xf0\xf4\x3d\x15\x82\x4c\x3a\x24\x33\xb5\x20\x13\xba\xa0\xab\xdf\x5b\x10\x35\xa1\x72\xe1\x82\xed\x3b\xd0\x59\x88\x88\x31\xe9\xea\x76\x7e\x78\x0f\x55\x60\x01\x2b\x43\x26\x9b\x1f\xd5\x70\x09\x1f\xe3\x94\x74\xf5\x80\xd3\x7b\x9a\x91\x08\x9c\x7a\x88\xc8\x76\x07\xec\x9a\x30\x62\x64\x60\x8b\xa6\x8e\x88\x2e\x74\xa1\x26\xae\x43\xec\x20\x1b\xa4\x0e\xa4\xab\xbb\x39\xa1\x42\xad\x3e\x1d\x53\xff\xaf\x75\x6e\x03\x33\x5a\x51\x13\x49\xbb\xdf\xbd\xef\x66\xa0\x56\xa4\x4f\x3f\x19\x23\x13\x89\x60\x5c\xda\xd0\x99\xc7\x87\x27\x7d\xc0\x92\x0d\xfc\x32\x34\x95\x40\x6c\x50\x56\x7a\xbd\xf1\xc0\x5b\x5d\x7d\x9e\x31\x6f\x97\xe3\x50\x74\x9a\xa6\x65\x0d\x16\x37\x21\x6c\xdf\x54\xf9\xee\x05\xf4\xbe\xd2\x79\xc2\xa3\x23\xde\x97\xb6\xa1\xeb\x81\x77\x26\xd4\x35\x59\x2c\xea\x90\x09\xb2\xd2\xe0\x8b\x24\x20\x94\x84\x6c\xfa\xa6\xee\xee\xf7\xbc\x9f\xf8\xf2\xee\xfa\x61\xe3\xc9\x2e\x9b\x30\x3a\xf4\x34\x62\xdc\xed\x7b\x04\x6b\xc3\xc0\x92\xcc\x6f\xbf\x3e\x3e\x3c\x99\x34\x6c\x92\x8d\xb4\xc1\x98\xf2\xe9\xf8\xa6\xde\x7e\xee\x80\xbc\x02\x8f\x51\xbe\x3e\xfe\x0e\x5f\xf3\x5d\x76\x6a\xec\x57\xd7\xdd\xe9\x45\xe7\xc9\x22\xa7\xfd\x1d\x4b\x5d\x73\x3e\xb9\x5e\x53\xd7\x68\x05\x5e\x56\xfa\xf1\xe3\xbd\x56\x13\x39\x69\x57\xfa\xc7\x2b\xad\x52\xb4\x2b\x6d\xa8\x6b\x4c\x68\xfc\xb2\xcf\xf6\x5f\x94\xfd\xd7\x78\x04\x9f\x45\x75\x10\x1b\xe2\xc2\x63\x2d\x45\xea\x8a\xe2\x4a\x57\x8b\x1b\x9f\x0f\x1f\x0b\xb2\x28\x4a\x0a\x46\x20\x0f\x1b\x8f\x6a\x60\x87\x51\x49\x8b\xaa\x84\x73\x3b\x76\x54\xe2\xc5\x26\x07\x1b\x40\xb4\x2d\x8d\x98\xb2\x2c\xcb\x61\x6b\x5c\xb0\x26\xf4\xc8\xc5\x3c\x45\x3a\x64\x99\xfb\xb2\x68\x46\x8c\x89\x02\x9b\x6b\x73\x2a\xc1\x63\x8f\xac\x3e\x7e\xc1\xa8\xfb\x19\xa3\xc6\xeb\xe5\x55\x96\x44\x95\xa9\x07\x7e\xdd\x9a\xda\xf7\x87\xcf\x92\x87\x55\x9e\x27\xd4\xb4\xf9\x9f\x82\xf7\xba\x2a\x4d\x5e\x5e\xfd\xa0\xf0\xd9\x62\x2f\x6a\x6a\x31\xa2\x0a\xd2\x62\x9c\x28\xa1\xca\xe6\x45\xb7\x28\x4d\x7f\x2a\xf0\x51\x59\x72\xf9\x45\x08\xa7\x3e\xc4\xcc\xfc\xc5\x4b\x7f\xa4\xe2\x18\xa6\x3f\x31\xd7\xcb\xa8\x53\x50\xe7\x2e\x00\xf9\xe6\xdf\x41\x66\xd6\xbc\x5a\xfc\x15\x00\x00\xff\xff\xc2\xab\x14\xfd\xd1\x09\x00\x00")

func templatesPartialsFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsFooterTmpl,
		"templates/partials/footer.tmpl",
	)
}

func templatesPartialsFooterTmpl() (*asset, error) {
	bytes, err := templatesPartialsFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/footer.tmpl", size: 2513, mode: os.FileMode(420), modTime: time.Unix(1479996327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x58\x4d\x8f\xdb\x36\x13\xbe\xef\xaf\x18\xf0\xbd\x24\x07\xad\xb2\xc9\x8b\x7e\x41\x76\xd1\x6e\x02\x34\x40\x93\x16\xd9\xe4\xd0\x93\x31\x16\xc7\x32\x63\x8a\x54\x48\xca\xbb\x86\xe1\xff\x5e\x50\xdf\x96\x25\xd9\x6e\x73\x68\x7c\xb2\xc8\xe1\xf0\x99\x67\x66\x38\x43\x46\x6b\x42\x4e\x66\x7e\x03\x00\x10\x21\xc4\x12\xad\x9d\x31\xbb\x11\x99\x14\x6a\xc3\x60\x6d\x68\x35\x63\xff\x4b\x51\x28\x06\x0e\x97\x42\x71\x7a\x9a\xb1\x17\x6c\xfe\xb0\x11\x19\x38\x0d\x7e\x0a\x62\xad\x1c\x29\x17\x85\x58\xa9\xe2\x62\x0b\x82\xcf\x58\x86\x09\xfd\x89\x6e\xcd\x6a\xd5\x6b\xc1\x89\xcd\xf7\xfb\xdb\x4f\x1f\xde\x1e\x0e\x51\xc8\xc5\x76\x7e\x53\xac\xd9\xef\xc5\x0a\x6e\x1f\xc8\x6c\x45\x4c\xef\xc8\x5a\x4c\xe8\x70\x28\xa6\x1a\x95\x95\x92\x25\x39\x0c\x96\xa8\x14\x19\x36\x6f\x24\xfa\x52\x8f\x06\xb3\xec\x44\xa2\x90\xca\x3c\x82\xfe\x56\x51\x98\xf5\x94\x95\xe8\x06\x3e\xf7\x7b\x52\xfc\x70\xb8\xb9\x39\xbf\x67\x77\xb6\x64\x1b\x62\x2d\x03\x2f\x38\x81\x3d\xd6\xb2\x10\x0b\x64\x12\x68\x45\x81\x5b\x0b\xc3\xcb\x91\x94\xb7\x23\x43\xb6\x61\x41\xbc\xd4\x89\x0e\xba\x3e\x0c\x07\x64\x0b\x79\x91\x26\xf5\xae\x7e\x11\x03\x6b\xe2\x19\x0b\xd1\x5a\x72\x36\x14\x69\x12\x6a\x65\x03\x3f\x75\x6b\xb7\x09\x03\x94\x6e\xc6\xfe\x58\xad\x44\x4c\xb0\xd2\x06\xde\xa3\x13\x5a\xa1\x84\x07\x87\x4e\x58\x27\x62\x3b\x04\xab\x0e\x8d\x11\x72\xa7\x08\x70\x8f\xba\x34\xd7\x36\x0c\x74\x86\x7c\x44\x05\x81\x4d\x21\x33\x42\xb9\x20\x28\x23\xec\x14\x00\x97\x25\x31\xa8\x92\x1c\x13\x6a\x22\xb2\x19\x18\xe1\x87\xbb\xbe\xe4\x62\xe1\x84\x93\xc4\xe6\xf7\x6b\x54\x09\x41\x3d\xfe\x53\x14\x72\x37\xa6\x85\x9f\x6a\x11\x8e\xd2\x91\x5d\xa1\x4e\x08\xfa\x02\xb7\xbf\x57\x2b\x80\x91\x62\x87\x43\x9b\xa7\xad\x2a\xef\x69\x10\xb1\x56\x95\xfd\xb5\xd3\xc3\x78\x77\xeb\x43\x5d\x38\x7a\xad\x7d\xaa\x1e\x0e\x75\xee\xb1\xf9\xfd\x2e\x35\x48\x09\x3c\xbb\xff\xeb\xb9\xf7\x4f\x1d\xd5\x57\x00\x8a\x77\xd7\x02\x1a\x45\xf3\x46\x25\x52\xd8\x35\x3c\x7b\xf3\xfe\x2c\x9a\x28\xe4\x7c\x28\xc6\xb8\xbc\x2a\xc8\x2c\xc5\x5a\x71\x34\xbb\x40\xe1\x16\x2e\x0f\xb9\x33\x91\x96\x97\x91\xa6\x70\x1b\x34\x3b\xb0\xc1\x3d\x3d\x4f\xd6\xc1\x67\xeb\x3f\x82\x58\x6a\x45\xe5\xd0\x58\x30\x4a\x31\xa2\xe7\x4c\x2c\x75\xce\xf6\xfe\xfe\x6a\x03\x27\x00\x3a\xc7\x86\x21\x49\x68\x29\x46\x49\x7e\x19\x9b\x7f\x28\x07\xa0\x1e\x39\xc9\xec\x96\x7c\x29\xfe\x33\x66\xa4\xe4\xd6\x9a\x6b\xa9\x93\x1d\x9b\xbf\x6b\x3f\xbe\x0d\xf4\x8a\x1e\xad\x87\xcd\x05\x7e\x1b\x80\x71\xa9\x73\x97\x5b\x36\xff\xc5\xff\xb9\x12\x73\x14\xe6\xd3\x79\xdc\xf9\xec\xfe\xed\x64\x76\x66\x44\x5a\xe7\xf5\x70\xb6\x46\x0a\xfb\x27\x43\x2e\xeb\xe5\xde\xac\xc0\x37\x35\x46\xcb\xc1\x6a\xd6\x32\x7a\x24\x3a\xc9\x68\x84\x75\x3b\xe5\xd7\x54\x08\x59\x71\x54\xa4\xa4\xf2\xc0\xe9\x24\x91\xc4\x00\x8d\xc0\x46\x63\xb9\x41\x23\x3c\xbc\xa9\x5f\x3e\xe5\x46\x9b\xa1\x1a\x59\xea\xe8\xc9\xf9\xc8\x52\x79\x14\x7a\xb1\x31\x3f\x0d\x38\x70\xc4\x79\x93\xd4\xc0\x45\xdc\x58\x42\x13\xaf\x4b\x6a\xca\xff\x53\xe4\xd4\xd2\xc3\xbb\x56\xb3\xff\x86\x9d\x87\x42\xc5\x57\xe0\x67\x20\xb0\x4f\x7b\x47\xf0\x18\x7c\x71\x2c\x42\x96\x93\x62\x4d\x35\x69\xc2\xa0\x60\x81\x9e\x32\x54\x9c\xf8\x8c\xad\x50\xda\xb1\x4a\x74\x9a\x10\x97\x56\x98\xa3\x25\x85\xf3\x9a\x06\xe0\xd3\x87\xb7\xc0\x42\x76\x38\x9c\xc8\x04\x01\xc6\x4e\x6c\xa9\xaa\xde\xd5\x19\xd1\x36\x69\x5a\xf2\x40\x50\x10\x70\x61\x33\x89\xbb\x60\x29\x75\xbc\x19\xf5\x4e\x7b\x00\xf5\xf0\xab\x4d\x5b\xae\x53\x1e\x7c\xdf\x14\xee\x1f\x3b\xfd\xee\x6f\x3a\xa5\x7f\x70\x5c\xee\xf7\xa6\x68\xeb\x6e\x3f\xe2\x93\x56\x3a\xdd\x4d\xb4\x44\x93\x74\x55\xb6\x7f\xb6\x95\xab\x70\x29\xc7\xba\xcc\x8e\xc1\x25\xfc\xa6\x2f\xba\x8c\x80\x1f\x1a\x02\xee\x5e\x54\xe1\xb1\x46\x9b\xe9\x2c\xcf\x66\xcc\x99\xbc\xbc\x73\x7d\xf4\x6d\xab\xbf\xe8\x8c\x70\xd2\x32\xe0\xaf\x62\xf7\x6b\x21\xb9\x21\x35\x61\x7e\x03\x7b\x24\xd0\x62\xaf\x22\x28\x7a\x9c\x2e\xda\xbb\xef\x1a\xb8\x2f\x5f\x1c\xf3\xb3\x58\x54\x77\xc9\xba\xba\xd4\x39\x30\x18\xf2\xe5\xa8\xc4\x25\xc9\x19\xb3\xf9\xf2\xcc\x41\x78\x6c\x63\xe5\xe5\x2b\xcc\x6c\xcc\x1d\xf1\x7a\x69\x6e\xe1\xfb\x13\xab\xfc\xd4\x85\xd0\x9a\x6d\x2e\x0b\x86\x9a\x64\x5f\x7e\xdb\x2b\x7a\x70\x77\x95\xcb\x8f\xf6\x1d\x4d\x8d\xfe\xef\xdc\x8d\xa1\xd5\xd8\x3f\xf8\xae\xd5\x34\x95\xae\x13\xf7\x84\xf3\xe9\x79\x51\xfb\x73\x9a\x7c\x13\xd9\x57\x9d\x3f\x36\x37\x5b\xda\x0d\x75\x0f\xdd\xdf\x47\xdc\x08\x95\x40\x86\xc6\x81\x50\x80\x50\x2e\xfb\x79\x82\x88\xaf\xdd\x4b\x35\x8d\xd0\x48\x2f\x55\x56\x50\x68\x4b\x6d\x5d\x98\x8e\x3a\xab\x4e\xb1\xfe\x15\xc7\x1e\x40\xea\xe7\x11\x30\x5a\x52\x2d\xde\x7f\x06\x59\x69\x93\x76\x9e\x01\x8a\xa7\x12\x28\x45\x17\x0b\x3f\xc9\xc0\xd7\x18\xad\x3c\xc9\x63\xd5\x3d\x2a\x0e\x85\x63\x13\x16\x8b\x6a\xb0\xe3\xba\x97\xaf\xda\xa3\xe8\xff\x0c\x56\xda\x1c\x35\x15\x55\xed\x2f\x1e\x3a\x10\x36\xb4\x7b\xd4\x86\x3f\xb3\xcf\x41\x1b\x70\x22\x25\xb0\x64\x04\x59\x78\xfb\x3a\x0a\x0b\xe5\x03\x48\x84\xca\x72\x07\x6e\x97\xb5\x16\x03\xe6\x4e\xc7\x3a\xcd\x24\x39\x9a\x31\xbd\x5a\xb1\x3e\xd4\x72\xd5\x11\xd4\xbb\x06\xea\xab\x97\xac\x73\xc5\x2c\x54\x8e\x05\x8c\xc2\x94\x66\xec\x0b\x83\x2d\xca\x9c\x66\x6c\x88\xab\x65\xee\x9c\x56\x35\xc4\x7c\x99\x0a\x77\x82\xa7\x92\xa9\xc1\xb4\xb4\xbd\xea\x43\x09\x2a\x0d\x23\x61\xda\x6d\xb7\xb6\xc2\xe6\x28\xe5\xae\x3a\xe4\x2f\xe9\xb4\xba\xcb\x45\xac\x55\xf9\xc8\x50\x07\xa6\x14\xc9\xda\xb1\xf9\x98\x86\x28\x2c\xcd\xe8\x67\x84\x8f\xaa\xa9\xeb\x45\x14\xd6\x0f\xa5\x7f\x07\x00\x00\xff\xff\xcb\x23\x40\x83\x32\x15\x00\x00")

func templatesPartialsHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsHeaderTmpl,
		"templates/partials/header.tmpl",
	)
}

func templatesPartialsHeaderTmpl() (*asset, error) {
	bytes, err := templatesPartialsHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/header.tmpl", size: 5426, mode: os.FileMode(420), modTime: time.Unix(1480070731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesProductpageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x4b\x8f\xa4\x36\x10\xbe\xef\xaf\x28\x71\xaf\x61\x98\x5e\xe5\x21\x31\x1d\x25\xd9\x43\x22\x25\x52\x14\x25\xe7\x96\x1b\x17\xb6\xd3\xc6\x46\xc6\xf4\x6c\x6b\xd4\xff\x3d\x82\x81\x06\xdb\xf4\x3e\xb4\x27\xfc\x28\x97\xeb\xfb\xea\xab\x32\x25\x57\x67\xa8\x34\xeb\xba\xe7\xac\x65\x82\x50\x19\xef\x2c\x1c\x59\x75\x12\xce\xf6\x86\x23\x0a\xa6\x35\xb9\x4b\xb6\x7f\x07\xb0\x36\x7f\x71\xac\x6d\xc9\x8d\xeb\xe1\x4e\x65\x35\x0e\xbb\xd3\x56\xb2\x79\x5b\x07\x28\x0d\x3b\x2f\xb3\xd0\xf2\xe8\x88\xf1\xca\xf5\xcd\x31\x5b\x9b\x00\x94\x56\x83\xe2\x81\x41\x7a\xe6\x70\xd0\xaa\xf3\xd1\x49\x80\xd7\x57\xc7\x8c\x20\x78\xf8\xe5\x66\x79\xbd\x46\x36\x00\xa5\x56\x5b\x1e\x95\xa7\x26\xf1\x38\xda\x33\x90\x8e\xea\xe7\xec\xf5\x15\x1e\xfe\xfd\xfb\xf7\xeb\xf5\x4e\x44\xe6\x94\xed\x07\x9b\x7f\x94\xd7\x74\xbd\x96\x39\x4b\xdd\x95\xb9\x56\x69\xd8\x64\x78\x14\x68\x99\x5b\x1d\x70\x97\x73\xb5\x22\xb3\xcc\x03\x6e\x47\x66\x07\xd6\x96\x34\x67\xab\x9c\xc0\x90\x34\x6c\x38\xbe\xff\xfe\x6d\xa8\x05\xee\x7e\x08\xc0\x96\xb2\x48\xa5\x72\x38\xf8\x01\xca\x1b\xaa\x3f\xc9\x33\xce\x3c\x5b\xe0\xc9\x22\xf0\xd0\x6e\x39\xa8\xac\xf1\x64\xe2\x54\x05\xfe\x3e\x50\x57\x39\xd5\x7a\x65\x4d\x40\x42\x99\xb7\x6b\xc0\x21\xfe\x95\x94\x0c\x3b\x63\x47\x95\x35\x9c\xb9\x0b\xe2\xd1\x3a\x4e\x0e\x35\xd5\x1e\xb5\x80\x3b\xf0\x8b\x1f\x6f\x43\x5b\xd7\x1d\x79\x7c\x82\x96\x71\xae\x8c\x18\x8f\x42\xc3\x9c\x50\x06\xbd\x6d\x07\x1b\xdc\x45\x6c\x3d\xcd\xb7\x47\x76\x45\xb6\xff\x4d\x09\xa9\x95\x90\x9e\x38\x4c\xf0\xcb\x5c\x3e\x05\xe7\xfb\x37\x91\xcb\xd9\xb4\xc3\x51\xd0\xb3\xd3\x61\x82\x68\xa8\xf7\x8e\xe9\x75\x28\x68\x48\x30\xaf\xce\x84\xd6\x10\xd6\xea\x63\x42\xec\x54\x01\x1f\x06\x66\x57\x91\xfc\xc5\x04\x75\x49\x31\x94\xa9\x18\x37\xf5\xfe\x19\x59\xa7\xa2\x4e\x25\x5d\xe6\xbd\xbe\x93\xce\xd5\xe4\x36\x9c\x06\xf3\x27\xee\x64\x13\xaf\x30\x25\x7b\xa4\x46\x39\x6b\xb0\x6b\xd2\xb5\x86\x7f\x53\x7f\x0b\x2b\x6d\xca\x5c\x25\x95\xe6\x8e\xcc\x27\xf2\x16\xa4\x26\x4c\xcc\xaf\xd3\xe9\xb8\xe8\x97\xce\xb4\x96\xad\x64\xba\x5e\xd4\x6a\x08\xbd\x54\x8e\x07\x7d\xbc\x21\x57\xf5\xee\x02\x92\x86\x74\x23\xee\x0a\x54\x86\x93\xf1\xc4\x91\xb4\x56\x6d\xa7\xba\x40\x47\x8f\xf3\x6c\xac\x93\x86\x23\x16\xf3\xca\xd1\x7a\x6f\x1b\x5c\x15\xc4\x74\x62\x9e\xba\xb7\x3b\x96\x85\xf9\xc4\x63\x50\x42\xc3\xfc\xbf\x0e\xa5\x3d\x93\xc3\x4a\xab\xea\x94\xb4\xd7\x35\xf1\x8d\x32\x38\x87\x5f\x3c\x6e\xbd\x52\x61\x3c\x4f\xd1\x65\x45\x1c\x5f\xb1\xd1\xcd\x4b\xb9\x9b\xef\xab\x75\xdf\xc9\x6c\xff\x05\x6a\x2f\x73\xb9\x4b\x25\x1f\xf4\xa3\x14\x4e\x10\x6b\x1a\x5b\xc2\x5d\x91\xc0\xf9\x1a\x6d\x47\x71\xb4\x11\xc8\x8d\x57\x6d\xc0\x78\xaf\xf3\xce\x10\xdb\x2f\x40\x1d\xd7\x7e\x5c\xf9\xeb\xba\x0f\x5e\xad\xa8\xe8\x3e\xa3\xf6\x6c\xe3\xb1\xd3\xb6\x62\x1a\x3b\xcf\xbc\xea\xbc\xaa\x86\x1e\x6a\x4e\xb7\x5a\x0c\xa5\xb7\xaa\x8c\xad\xc2\xc9\xee\xfe\xa7\xd8\x8f\x87\x83\x6a\x98\xa0\x50\x7d\xc5\x77\xf7\x53\x3a\x25\x30\xfe\xb1\x49\xa4\x17\x93\x3b\x2b\x31\x97\xa4\xdb\x7c\x44\xb7\x80\xcb\xf6\x7f\x58\x7b\x52\x46\x40\x6d\x1d\x8c\x9b\xb0\xec\xfe\xb4\xd5\x95\xbb\x96\x99\xf9\xca\x11\x02\x4a\xab\x39\xb9\x1b\x17\xc5\xfb\x6c\x5f\xaa\x46\x40\xe7\xaa\xe7\x2c\x57\x8d\xc8\x7d\x81\xa3\xef\x87\xd6\x88\x0c\x98\xf6\xcf\x59\x06\xf9\xbe\xcc\x07\x67\x71\xe6\xe3\xb2\x58\x74\x37\xd2\x36\x37\xe9\x6d\xb8\x3f\x83\x64\x86\x5f\x40\xf4\x8a\x13\x78\x0b\x9a\x3c\x5c\x6c\x0f\x27\x63\x5f\xe0\x45\x92\x1b\x57\x6b\x65\x78\x02\xf7\xe1\xdd\x27\x84\x9a\xfc\x2a\x7d\xd5\x53\xf3\x7f\x00\x00\x00\xff\xff\xad\x53\x4d\x1a\x35\x0b\x00\x00")

func templatesProductpageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesProductpageTmpl,
		"templates/productPage.tmpl",
	)
}

func templatesProductpageTmpl() (*asset, error) {
	bytes, err := templatesProductpageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/productPage.tmpl", size: 2869, mode: os.FileMode(420), modTime: time.Unix(1479996327, 0)}
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
	"templates/hello.tmpl": templatesHelloTmpl,
	"templates/homepage.tmpl": templatesHomepageTmpl,
	"templates/main.tmpl": templatesMainTmpl,
	"templates/partials/footer.tmpl": templatesPartialsFooterTmpl,
	"templates/partials/header.tmpl": templatesPartialsHeaderTmpl,
	"templates/productPage.tmpl": templatesProductpageTmpl,
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
		"hello.tmpl": &bintree{templatesHelloTmpl, map[string]*bintree{}},
		"homepage.tmpl": &bintree{templatesHomepageTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{templatesMainTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"footer.tmpl": &bintree{templatesPartialsFooterTmpl, map[string]*bintree{}},
			"header.tmpl": &bintree{templatesPartialsHeaderTmpl, map[string]*bintree{}},
		}},
		"productPage.tmpl": &bintree{templatesProductpageTmpl, map[string]*bintree{}},
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

