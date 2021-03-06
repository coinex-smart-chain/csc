// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (11.27kB)

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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x93\xdb\xb6\x92\x7e\x96\x7f\x45\x87\x3b\x89\xa8\x9d\x21\x29\x79\x72\x2b\x89\xd4\x96\x77\xe2\x4d\xcd\x56\x6d\x92\x5a\x3b\xb5\x7b\x2a\xc9\x03\x44\xb4\x44\x78\x40\x80\x01\x40\x69\x14\x95\xfe\xfb\x29\x80\x17\x91\x94\x66\x62\xc7\x3e\x75\xea\xf8\x61\x4c\x02\x7d\xf9\xd0\xdd\x40\x77\x83\x8a\x3f\xfb\xee\xc7\xbb\xb7\x7f\xfb\xe9\x35\x64\x26\xe7\xcb\x17\xb1\xfd\x0f\x38\x11\x9b\xc4\x43\xe1\x2d\x5f\x8c\xe2\x0c\x09\x5d\xbe\x18\x8d\xe2\x1c\x0d\x81\x34\x23\x4a\xa3\x49\xbc\xd2\xac\x83\x6f\xbd\xd3\x44\x66\x4c\x11\xe0\xef\x25\xdb\x26\xde\xff\x07\x3f\xbf\x0a\xee\x64\x5e\x10\xc3\x56\x1c\x3d\x48\xa5\x30\x28\x4c\xe2\xdd\xbf\x4e\x90\x6e\xb0\xc3\x27\x48\x8e\x89\xb7\x65\xb8\x2b\xa4\x32\x1d\xd2\x1d\xa3\x26\x4b\x28\x6e\x59\x8a\x81\x7b\xb9\x01\x26\x98\x61\x84\x07\x3a\x25\x1c\x93\x99\xb7\x7c\x61\xe5\x18\x66\x38\x2e\x0f\x87\xf0\x07\x34\x3b\xa9\x1e\x8e\xc7\x39\xbc\x2a\x4d\x86\xc2\xb0\x94\x18\xa4\xf0\x5f\xa4\x4c\xd1\xc4\x51\x45\xe9\x98\x38\x13\x0f\x90\x29\x5c\x27\x9e\x85\xae\xe7\x51\x94\x52\xf1\x4e\x87\x29\x97\x25\x5d\x73\xa2\x30\x4c\x65\x1e\x91\x77\xe4\x31\xe2\x6c\xa5\x23\xb3\x63\xc6\xa0\x0a\x56\x52\x1a\x6d\x14\x29\xa2\xdb\xf0\x36\xfc\x26\x4a\xb5\x8e\xda\xb1\x30\x67\x22\x4c\xb5\xf6\x40\x21\x4f\x3c\x6d\xf6\x1c\x75\x86\x68\x3c\x88\x96\x7f\x4d\xef\x5a\x0a\x13\x90\x1d\x6a\x99\x63\xf4\x65\xf8\x4d\x38\x75\x2a\xbb\xc3\xcf\x6b\xb5\x6a\x75\xaa\x58\x61\x40\xab\xf4\xbd\xf5\xbe\xfb\xbd\x44\xb5\x8f\x6e\xc3\x59\x38\xab\x5f\x9c\x9e\x77\xda\x5b\xc6\x51\x25\x70\xf9\x51\xb2\x03\x21\xcd\x3e\x7a\x19\x7e\x19\xce\xa2\x82\xa4\x0f\x64\x83\xb4\xd1\x64\xa7\xc2\x66\xf0\x93\xe9\x7d\xca\x87\xef\x86\x2e\xfc\x14\xca\x72\x99\xa3\x30\xe1\x3b\x1d\xbd\x0c\x67\xdf\x86\xd3\x66\xe0\x5c\xbe\x53\x60\x9d\x66\x55\x8d\xc2\x2d\x2a\x1b\xb9\x3c\x48\x51\x18\x54\x70\xb0\xa3\xa3\x9c\x89\x20\x43\xb6\xc9\xcc\x1c\x66\xd3\xe9\xe7\x8b\x4b\xa3\xdb\xac\x1a\xa6\x4c\x17\x9c\xec\xe7\xb0\xe6\xf8\x58\x0d\x11\xce\x36\x22\x60\x06\x73\x3d\x87\x4a\xb2\x9b\x38\x3a\x9d\x85\x92\x1b\x85\x5a\xd7\xca\x0a\xa9\x99\x61\x52\xcc\x6d\x44\x11\xc3\xb6\x78\x89\x56\x17\x44\x9c\x31\x90\x95\x96\xbc\x34\x38\x00\xb2\xe2\x32\x7d\xa8\xc6\xdc\x6e\xee\x2e\x22\x95\x5c\xaa\x39\xec\x32\x56\xb3\x81\x53\x04\x85\xc2\x5a\x3c\x14\x84\x52\x26\x36\x73\xf8\xba\xa8\xd7\x03\x39\x51\x1b\x26\xe6\x30\x3d\xb1\xc4\x51\x63\xc6\x38\xaa\x0e\xae\x17\xa3\x78\x25\xe9\xde\xf9\x90\xb2\x2d\xa4\x9c\x68\x9d\x78\x03\x13\xbb\x03\xa9\x47\x60\xcf\x21\xc2\x44\x33\xd5\x9b\x53\x72\xe7\x81\x53\x94\x78\x15\x88\x60\x25\x8d\x91\xf9\x1c\x66\x16\x5e\xcd\x32\x90\xc7\x03\xbe\x09\x66\x2f\x9b\xc9\x51\x9c\xcd\x1a\x21\x06\x1f\x4d\xe0\xfc\xd3\x7a\xc6\x5b\xc6\xac\xe1\x5d\x13\x58\x93\x60\x45\x4c\xe6\x01\x51\x8c\x04\x19\xa3\x14\x45\xe2\x19\x55\xa2\x8d\x23\xb6\x84\xee\xf1\xf7\xc4\xe9\x97\xcd\x1a\x5c\x11\x65\xdb\x7a\x59\x9d\xc7\xc1\x0a\x9f\x5e\xc4\xb7\x50\x3f\xc8\xf5\x5a\xa3\x09\x3a\x6b\xea\x10\x33\x51\x94\x26\xd8\x28\x59\x16\xed\xfc\x28\x76\xa3\xc0\x68\xe2\x95\x8a\x7b\xf5\xf1\xef\x1e\xcd\xbe\xa8\x4d\xe1\xb5\x0b\x97\x2a\x0f\xac\x27\x94\xe4\x1e\x14\x9c\xa4\x98\x49\x4e\x51\x25\xde\xbd\x93\xb3\x97\xa5\x82\x3b\xc9\x04\x3e\xc2\x9b\x9c\x28\x03\x77\x19\x61\x02\x08\xa5\x36\x40\xc3\x30\xf4\xa2\x93\x6e\x17\xae\xe7\xe8\x82\x95\x11\x27\x84\xa3\x78\x55\x1a\x23\x5b\xc2\x95\x11\xb0\x32\x22\xa0\xb8\x26\x25\x37\x40\x95\x2c\xa8\xdc\x89\xc0\xc8\xcd\xc6\x26\xb7\x0a\x77\xc5\xe4\x01\x25\x86\xd4\x53\x89\xd7\xd0\x36\x6e\x23\xba\x90\x45\x59\xd4\x8e\xab\x06\xf1\xb1\x20\x82\x22\xb5\x6e\xe6\x1a\xbd\xe5\xf7\x6c\x8b\x90\x23\x98\xbb\xd7\x6f\x87\x31\x90\x12\x85\x26\xe8\x8a\x3c\x8b\x84\x38\xaa\xa0\x54\x0b\x82\xfa\x5f\x5c\xf2\x46\x52\xbb\x80\x1c\x45\x09\xbd\xb7\x40\xd9\x83\xc4\x5b\x1e\x0e\x8a\x88\x0d\xc2\x15\xa3\x8f\x37\x70\x45\x72\x59\x0a\x03\xf3\x04\xc2\x57\xee\x51\x1f\x8f\x3d\xe9\x00\x31\x67\xcb\x98\x3c\x17\xcf\x20\x45\xca\x59\xfa\x90\x78\x86\xa1\x4a\x0e\x07\x2b\xfc\x78\x5c\x80\xde\xe7\x2b\xc9\x93\xb1\x5d\xef\x78\x01\x87\x03\x5b\xc3\x55\xf8\xbf\x98\x92\xc2\xa4\x19\x39\x1e\x37\xaa\x79\x0e\xf1\x11\xd3\xd2\xa0\x3f\x39\x1c\x90\x6b\x3c\x1e\x75\xb9\xca\x99\xf1\x1b\x69\x76\x5c\xd0\xe3\xd1\x2e\xa1\x86\x7d\x3c\x42\x64\x85\x0a\x8a\x8f\x70\x15\xfe\x84\x8a\x49\xaa\xa1\xa2\x8f\x23\xb2\x8c\x23\xce\x96\x35\x5f\xdf\x66\x51\xc9\x4f\xc1\x13\xd9\xe8\x69\xe3\xfc\x5f\x23\x94\x52\x95\xbe\x9c\x82\x91\x0f\x28\xf4\x3f\x29\x94\xa0\x8d\xa5\xca\xcf\x37\x70\xe5\x50\xdd\x8b\xb5\x74\x11\x75\x67\xdf\xde\x5a\x88\x76\xa8\x76\x81\x33\xde\x07\x84\x54\x1d\x43\x87\x43\xad\xc5\xc6\xd5\x5f\x0d\xa4\x46\x42\x2f\x96\x5a\xcc\xf5\x0e\x78\x63\xd4\xf1\x08\x1d\xea\xbf\x1c\x49\xee\x00\x76\x58\xbb\x50\x2f\x9c\xa7\x9b\xa0\x85\x5f\xc7\x86\x66\x06\x1f\x70\x9f\x78\x87\x43\x97\xb7\x9e\x4d\x09\xe7\x2b\xe2\xac\xe3\xd6\xd6\x32\xfd\x81\x36\x66\xb7\x4c\xbb\xe2\x7c\xd9\x20\x38\xc1\x7e\xcf\x04\x31\x48\x81\x46\x16\x73\xb8\x7d\xd9\xc9\x7f\x97\x72\xc7\xd7\x83\xdc\x71\x7b\x91\xb8\x20\x02\x39\xb8\xbf\x81\xce\x09\x6f\x9e\xeb\x9d\xd3\xc9\x27\x43\xa6\xc0\x66\xfb\x16\x5a\x5b\x35\x4c\x17\x20\xb7\xa8\xd6\x5c\xee\xe6\x40\x4a\x23\x17\x90\x93\xc7\xb6\x72\xba\x9d\x4e\xbb\xb8\x6d\x53\x41\x56\x1c\x5d\x9e\x52\xf8\x7b\x89\xda\xe8\x36\x2b\x55\x53\xee\xaf\x4d\x4e\x14\x85\x46\x3a\xb0\x86\xd5\x68\x4d\xeb\xa8\x3a\xae\x6f\x8d\x79\x11\xfb\x5a\xca\xb6\x18\xe9\xc2\xa8\x45\x77\xea\x26\x6f\x19\x1b\x75\xa2\x1b\xc5\x86\x7e\x50\x31\xa1\x6c\xb3\xf0\x54\x2d\x51\x9d\x6e\x76\xed\x05\xa2\xaa\x2a\x55\x1b\xb2\xe0\x5e\xe3\xc8\xd0\x8f\xd0\x6c\x83\x70\x45\x34\xbe\x8f\x7a\x57\x33\x9e\xd4\xbb\xd7\x8f\xd5\x9f\x21\x51\x66\x85\xc4\xbc\x0f\x80\x75\x29\x68\x67\xfd\x36\x45\x7d\xac\xfe\x52\xb0\x2d\x2a\xcd\xcc\xfe\x7d\x01\x20\x3d\x21\xa8\xde\xfb\x10\xe2\xc8\xa8\xe7\x43\xad\xfb\x72\x71\x6f\xb7\x4f\xf5\xc3\xa9\xe3\x71\xd3\x51\x04\xdf\x73\xb9\x22\x1c\xb6\x16\xf2\x8a\xa3\x06\x23\xc1\xd6\x62\x60\x32\x84\xb4\x54\x0a\x85\x01\x6d\x88\x29\x35\xc8\xb5\x1b\x5d\xbb\xba\xd3\xf2\x6f\x89\x02\x62\x0c\xe6\x85\x81\xa4\xae\xd7\xed\x98\x46\xb5\xad\xbb\x10\xfb\x6a\x0b\x83\xfe\xbc\x3b\x5c\x21\x01\xcf\x6b\xc7\x9a\xed\x08\x09\xfc\xf2\xdb\xe2\x45\x0d\xef\x3b\x5c\x33\x81\x40\xac\x81\x52\xdb\x87\x80\xc9\x88\x81\x54\x21\x31\xa8\x21\xe5\x52\x97\xaa\x42\x6d\xd3\x14\x58\xe4\x8d\xa4\x46\xb2\x9d\x28\x1c\x82\x46\x88\x9f\x11\x9d\x4d\xea\x16\x44\xa1\x29\x95\x38\xcd\x35\xe3\xa3\xb5\x54\xe0\x5b\x01\x2c\x99\x2e\x80\xc5\x8d\xdc\x90\xa3\xd8\x98\x6c\x01\xec\xfa\xba\x25\x1e\xb1\x35\xf8\x0d\xc5\x2f\xec\xb7\xd0\x3c\x86\x56\x0b\x24\x09\x74\xb5\x39\x85\xb5\x1c\x5d\x70\x96\xa2\xcf\x6e\x60\x36\x59\x34\xb3\x2b\x85\xe4\xa1\x79\xab\xcf\xed\xea\x3f\xf7\xf7\xb8\xe8\x5b\xc6\x39\xa4\x67\x9b\x2a\x2b\x68\x20\xb0\x61\xda\x40\xa9\xb8\xb5\x8e\xa5\xab\xdc\xd2\x3a\xc1\xd1\x75\xad\x72\x96\xad\xea\x87\x3a\x87\x34\x4b\xa8\xc4\x84\x1a\x05\xf5\xff\xfb\xcd\x8f\x3f\x84\xda\x28\x26\x36\x6c\xbd\xf7\x0f\xa5\xe2\x73\xb8\xf2\xbd\x7f\xb3\xa5\xff\xe4\x97\xe9\x6f\xe1\x96\xf0\x12\x6f\x6a\x97\xcf\xa1\xa9\x16\x6c\x4c\xcc\xdd\xdf\x33\xad\x37\x50\x3f\xce\xa1\x0f\xe0\x38\x99\x2c\x2e\x67\xd4\x4e\x05\xa0\x50\xa3\xf1\x2d\x61\x9b\xf8\x86\x36\x23\x90\xa3\xc9\x24\xb5\x76\x51\x98\x4a\x21\x30\x35\x50\x16\x52\xd4\x26\x02\x2e\xb5\x3e\x05\x66\x43\x91\x9c\x07\x49\x4d\x9f\x80\xc0\x1d\xfc\x1f\xae\xde\xc8\xf4\x01\x8d\xef\xfb\x3b\x26\xa8\xdc\x85\x5c\xa6\xc4\x32\xd8\xde\xda\xc8\xd4\x06\x7d\x92\x40\x7d\xd3\xe0\x4d\xe0\x3f\xc0\xdb\x69\x3d\x8f\x22\x0f\xe6\xf6\xd1\x3e\x4d\xe0\x1a\x86\xec\x99\xd4\x06\xae\xc1\x8b\x48\xc1\xbc\x49\xb5\x3d\x1a\x47\x48\x91\xa3\xd6\x64\x83\x5d\x80\xb8\x45\x61\xda\xa0\xb3\xeb\xc8\xf5\x06\x12\x70\x0e\x2b\x88\xd2\x58\x91\x84\xf6\xdc\x6e\xa2\xcf\xc6\xb0\x23\x4b\x12\x10\x25\xe7\xa7\xa0\xad\x36\xc9\xa2\x09\xc7\x1e\x79\xe8\x4e\x53\xf8\x2c\x49\xc0\x9e\x62\xd6\xc4\xf4\xc4\x69\x83\xa1\x3a\x6e\x27\xa1\x3d\x48\x4f\x1c\x93\x45\x37\xba\x7b\xd2\x90\xfe\x99\x38\xa4\x43\x79\x48\x9f\x10\xe8\xb2\xdb\x73\xf2\xaa\x6c\xd8\x11\xe7\x06\x9e\x90\x26\xca\x7c\x85\xea\x39\x71\x55\x76\xab\xc5\x39\x53\xdf\x0b\xd3\xe1\xbd\x81\xd9\xd7\x93\x27\xa4\xa3\x52\xf2\x49\xe1\x42\x9a\xbd\x7f\xe0\x64\x2f\x4b\x33\x87\xb1\x91\xc5\x9d\xcb\x46\xe3\x1b\xb0\xba\xe6\xd0\x4a\xb8\x71\x2d\xc7\x1c\xc6\xee\xcd\xce\xb3\x1c\x1d\xd7\x57\xd3\xe9\xf4\x06\x9a\x9b\x9e\xff\x24\x76\x13\xaa\x12\x8f\x4f\xe0\xd1\x65\x9a\xa2\x7e\xd2\x7a\xef\x85\xa8\x96\xd1\x62\xaa\xdf\x3f\x02\x55\x9b\x2b\x7a\xb0\xe0\x8b\x2f\xe0\x6c\xb6\x1f\xc6\x51\x04\xff\x43\xd4\x03\xb8\xca\x53\xe1\x96\xc9\x52\x9f\x32\x4f\xce\xb4\x66\x62\x03\x44\x03\x95\x02\x6b\x9e\x0f\x4b\x03\x67\x18\x6b\x32\x58\xc2\x74\x08\xd0\x1e\x8f\x9d\x34\x71\x21\x7b\x74\xe4\xf6\x13\xc3\xe8\xd8\xd5\xd7\xe3\x64\x39\xc2\x67\x36\xb5\x76\x99\xcf\x28\xda\xdc\x3b\xaa\x4e\x12\xf3\xb6\xf2\x85\x5f\x67\xcb\x4b\xb9\x6c\x72\x63\x0b\xea\xe9\xe4\x0c\xc4\xf1\x64\xde\x57\x45\x81\x82\x02\x11\x7b\x77\x24\xb6\xb6\x65\xc2\x48\x90\xa5\x3d\x5a\x53\xc2\x6d\xb3\xc0\xd1\x9d\x54\x35\xab\x35\x70\x2a\xf3\x5c\x0a\x48\x20\x98\x2d\x2e\x64\xd5\x8e\x25\x3b\x4b\x1b\xba\xe7\x82\xed\x87\x2e\xea\xdb\x6c\x40\x1c\xcc\x7a\x4e\xe9\xf9\xeb\xb2\x63\x46\x2d\x6e\x76\xb2\xe8\xc0\x5d\x27\x7f\x0d\x6d\xd6\xc1\x5f\xc9\xb9\x9e\xbd\xe7\x32\xda\xe9\xa2\xd4\x99\x3f\x00\x3a\x59\x9c\xfb\xe6\xde\xa0\x22\x06\x5d\xc7\xe4\x7c\x81\xc2\x30\x85\x67\x2e\x01\x22\x6c\x15\x15\x28\x14\x14\x55\x53\x62\xd8\x86\xab\xea\x8e\x7a\x2e\x73\x9f\x75\x7a\xe1\xf4\x81\x1b\xc6\x95\x68\x52\xa0\x6d\xa8\x07\x9b\xc0\x05\x6a\x2f\x52\x2d\x31\x72\x52\x68\xa4\x90\x40\x75\xf1\xee\x4f\xc2\x52\xb0\x47\x7f\x12\xd4\xef\x43\x19\xcd\x7c\x9d\x36\x9d\xc7\x2a\xd8\xd7\x09\x78\xb1\x51\xb6\x24\x1f\x7b\x70\x7d\x69\x0b\xda\xac\x3b\x5e\x9e\x10\x74\x59\x01\x62\x43\x97\xae\xdf\xab\x5a\x85\x5f\x3d\xdb\x99\x6f\x94\x2c\x05\x9d\xdb\xd2\xcb\x3f\x13\x4b\xb6\xc4\x10\xe5\xa4\x4e\x16\x70\x22\x77\x0d\xfc\x1c\x52\xeb\x9c\x05\x54\x2d\xa1\xeb\xba\xa1\x6d\x66\xdd\xdb\x4a\x2a\x8a\x2a\x50\x84\xb2\x52\xcf\xe1\xcb\xe2\x71\xf1\x6b\xd3\xec\xbb\xfe\xe1\x59\xa8\x85\xc2\xe5\x19\xa2\x34\x75\xf7\x80\xd7\xe0\xc5\x91\x25\xf8\x33\x31\xed\x62\xbb\x17\xfe\x70\xa1\x4b\x82\xf6\x3a\xbe\x1e\xcf\x19\xa5\x1c\x2d\xe0\x93\x78\xbb\x19\xad\xff\xbb\x5b\xaa\xaf\x12\xea\xf6\xe8\xc4\x73\x04\xe4\x1a\x9f\x61\x68\x3b\xad\xb1\x0d\x80\xc0\x2e\x99\x39\x9b\xd7\x4d\x9b\x1b\x56\x63\x67\x8b\xfa\xf3\x0d\x2d\x95\xab\xb5\xfc\xa0\x0e\xb0\x1b\x18\x6b\x5b\xfb\x51\x3d\x9e\x84\x59\x99\x13\xc1\xfe\x40\xdf\xe6\xa5\x49\x65\x2b\xd7\xba\x79\xe7\x47\xf2\x19\x98\xd3\x85\xc0\xb8\xc9\x71\xe3\xda\x88\xe3\xc6\xbb\xd6\x91\x9d\x4f\x1f\xe3\x0f\xb4\xd0\x65\x2d\xc1\x8a\x28\xe8\xbe\x04\x4d\xf2\x05\x25\xad\xf6\x66\x6e\x45\xd4\xb8\x6a\x5b\x5d\xbd\x2e\xe4\x2e\x19\xdf\x4e\x5b\x90\x95\xa3\x9d\x9f\xc7\x75\xac\x9d\x39\xc3\xa2\x6c\xb6\xe6\x12\x6e\xa7\x9f\x02\x2d\x25\x62\x83\xc3\x15\x18\xc5\x0a\xa4\x40\x52\xc3\xb6\xf8\x0f\x58\xc8\x27\x30\xf2\x07\x43\xb4\x71\xd8\x18\xcf\x85\x69\x0f\xaf\x9d\x6d\x6d\xfb\xef\x76\xbf\x41\xe4\x2c\x7c\x0d\xde\xc5\x85\x3c\x19\x89\x03\xc2\xc1\xd6\x7e\x7a\xdf\xbb\xbb\x08\x6f\x98\x53\x6c\xb5\xdb\x5e\xa3\x4d\xc2\xcc\xe4\xdc\xf7\x62\xe3\x3e\xcc\x59\xcc\xad\x04\x27\xa0\x1a\xee\x97\x74\xc7\x7e\x23\x63\xfb\x79\x1c\xf4\x59\xd0\x29\x4e\xda\x5e\xac\xa9\x44\xe0\x78\xfa\x7e\x19\x45\xf0\xc6\x10\x65\x80\xc0\xcf\xf7\x50\x16\x94\x18\x9b\xbd\x24\xd8\xfc\xe8\xb2\x58\xfb\x81\x73\x45\x94\x86\xb5\x54\x3b\xa2\x28\x94\xc2\x30\x6e\xe7\xf7\x40\x14\xb6\xa5\x9f\x46\x73\x6f\x4f\xb1\x2d\xe1\xfe\x59\xdf\x77\xe5\x8f\xc3\xae\xcb\xc7\x93\x10\x49\x9a\x9d\x13\xba\x8c\xd5\xea\x4d\xe0\x07\xd7\x02\xf8\x57\xbe\xc9\x98\x9e\x84\xc4\x18\xe5\x8f\x7b\xc1\x30\x9e\x58\xbf\xce\x3a\x2d\x59\xcb\x1e\xf7\xb6\xd5\x73\x32\x4e\xc5\x74\x5b\x08\x34\xe4\xa9\xd6\x7e\x15\x57\xe3\x9b\x8e\xec\x7e\x58\x8d\x3f\x1f\xb7\x8e\x3a\x6d\xef\xd3\x3a\x92\x8b\x48\x7a\xa2\xc7\x76\x97\x8d\xcf\xd4\x13\x4a\xef\xec\xfe\xf1\xbd\x0b\x3b\x7d\x18\x1d\x93\xd6\xd8\xd5\x79\xfd\xac\x95\xab\x4f\x41\x4f\x98\x98\xd1\xf1\x24\xd4\xe5\xaa\xba\xab\xf0\xbf\x6a\x1b\xb0\x86\xcc\x05\xef\x30\x15\x9c\x15\x14\x56\x45\xbf\xa8\x08\x06\x45\xc8\x33\x59\xa3\x56\x59\xad\xea\x78\x63\x0d\x3e\x9d\xb4\x57\x5d\xaf\xb5\x2d\xae\x98\xce\x80\xc0\x0e\x57\xda\xdd\x24\x40\x1d\xef\xee\x76\xa7\xba\xc5\x79\xf5\xd3\x7d\xe7\x26\xa7\xdd\x11\xbe\x93\xde\xfe\xf6\xe0\xd2\x3d\xc9\xc5\x1f\x3b\xec\x76\xbb\x70\x23\xe5\x86\x57\x3f\x73\x68\x2f\x52\x22\x52\xb0\xf0\x9d\xf6\x80\xe8\xbd\x48\x81\xe2\x1a\xd5\xb2\x23\xbe\xbe\x5d\x89\xa3\xea\x33\x7c\x1c\x55\xbf\x34\xfa\x7b\x00\x00\x00\xff\xff\x0d\x85\x24\x20\x7a\x24\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0xa2, 0x98, 0x44, 0x4b, 0x50, 0xf8, 0xa1, 0xac, 0x4a, 0x76, 0x2e, 0xcc, 0x3d, 0xcb, 0x81, 0x9e, 0x2a, 0xaa, 0x87, 0xf5, 0x9d, 0x53, 0x4, 0x8a, 0xdd, 0x5a, 0xfe, 0xd3, 0xc3, 0xf, 0x11}}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
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
