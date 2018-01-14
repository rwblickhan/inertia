// Code generated by go-bindata.
// sources:
// client/bootstrap/daemon-down.sh
// client/bootstrap/daemon-up.sh
// client/bootstrap/docker.sh
// client/bootstrap/keygen.sh
// client/bootstrap/token.sh
// DO NOT EDIT!

package client

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

var _clientBootstrapDaemonDownSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x4a\xc3\x40\x10\xc6\xf1\xfb\x3c\xc5\x47\xdb\xeb\xe6\x0d\x72\x88\x36\x88\x60\x57\x28\x7a\xf0\xd4\xae\xd9\xd9\x74\x48\x33\x5b\x67\xb7\xa8\x6f\x2f\x22\x42\xe8\xfd\xe3\xff\xfb\xd6\xb8\x0b\x45\x06\x94\xc1\xe4\x52\x91\xb2\xe1\xdd\x44\x47\xd1\x11\x31\x7f\x2a\xea\x89\x11\x03\xcf\x59\x1b\xa2\xc2\x15\x8e\x89\xb6\x5d\xbf\x7b\xf6\x07\xdf\xed\xfa\x56\x94\xad\x4a\x70\x7f\x23\xa2\x35\xee\x4f\x3c\x4c\x90\x84\x70\x36\x0e\xf1\x1b\x76\x55\x15\x1d\x1b\xea\x9e\xf6\x7d\xb7\x7d\x3b\xec\x5f\xbd\x7f\xf4\x0f\xed\xb1\x5c\x63\x46\xcc\xc3\xc4\x86\x4b\x81\xfb\x80\x73\x49\xce\x95\x0d\x2b\x0d\x33\xb7\x9b\x05\xb5\x3a\xfe\xd6\x5f\xc2\xc4\xe0\x2f\x29\xf5\xff\x63\x43\xcb\x8c\xcd\x70\x09\x9b\x1b\x8a\x7e\x02\x00\x00\xff\xff\x8a\x49\xa7\x95\xe9\x00\x00\x00")

func clientBootstrapDaemonDownShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonDownSh,
		"client/bootstrap/daemon-down.sh",
	)
}

func clientBootstrapDaemonDownSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonDownShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-down.sh", size: 233, mode: os.FileMode(420), modTime: time.Unix(1515908421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDaemonUpSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x5d\x6f\xdb\x46\x10\x7c\xbf\x5f\xb1\xa5\x54\xf8\x25\xc7\x73\x5a\x14\x48\x1d\xe8\x41\x4d\x84\x58\x48\x45\x19\x92\x8a\x22\x68\x0b\xe5\x74\x5c\x89\x17\x91\x77\xec\xed\x52\x8c\xfd\xeb\x8b\xa3\x3e\x28\xbb\x7e\x92\xb8\xdc\x9d\x99\x1d\xce\x0e\xe0\x37\x4d\xd6\x00\x99\x60\x6b\x86\xad\x0f\xb0\x09\xd6\xed\xac\xdb\x41\x53\x03\x17\x08\xb9\xc6\xca\xbb\x54\x08\x42\x06\x89\x42\x3c\xcc\x17\xab\xd1\x8f\x24\x3e\x8e\x27\xb3\x79\xb6\xce\xc6\xb3\xc9\xc8\x3a\x0c\x6c\xb5\x3c\x36\x8b\x0f\xf3\x6c\x35\x9e\x66\x93\xc5\xba\x6b\x7e\x77\xfb\xee\xad\x98\xce\xc6\x9f\x26\xeb\xc5\xe4\x61\xbe\x9c\xae\xe6\x8b\x2f\xa3\x66\x63\x4a\xdd\x38\x53\xd4\x3a\x57\x27\x00\x21\x06\xf0\xa1\x40\xb3\x07\xbb\x05\x5d\x06\xd4\xf9\x23\x84\xc6\x39\xeb\x76\xa9\x18\xff\xbe\x98\x8c\x3f\x7e\x59\x2f\xfe\xc8\xb2\x69\xf6\x69\xf4\x95\x9a\xdc\x43\xee\xcd\x1e\x03\xd4\x04\xf2\x5f\x90\x72\x6b\x4b\xc6\x00\x89\xd3\x15\x8e\x86\x57\x22\x93\xaf\x11\x7d\xa5\xf7\x08\xf8\xdd\x12\xc7\x1d\x73\xdf\xba\x54\xd8\x2d\xfc\x05\x3f\x80\x7c\x82\x64\xf8\x82\x23\x81\x7f\xde\x47\x1b\x9c\x00\x00\x40\x53\x78\x48\x3e\xdb\xb2\x8c\xc3\x17\x14\xe3\x1d\xeb\xb8\x41\xd2\x75\x5d\xab\x0a\x15\xc8\x2d\xbc\x44\x15\x5b\xfb\x3e\x8a\x79\x68\xca\xb2\x33\xb9\xd4\x8c\xc4\x70\x72\xe1\xe2\xf9\xb3\xfd\x62\xef\xf0\xa5\x8b\x11\x65\x16\x57\x7a\x08\xfe\x1b\x1a\x86\xdc\x06\x34\xec\xc3\xa3\xa8\xf6\xb9\x0d\x20\x6b\x18\xde\xcf\x67\x13\x55\x1f\x1b\xe2\xc0\xa2\x71\xbd\x66\x68\x2d\x17\xa0\x8d\x41\x22\x60\xdf\xc9\x29\x3c\xf1\x99\x97\xe2\x0f\x83\x76\x39\x04\x2c\xf1\xa0\x5d\x4f\x62\x91\x40\x8a\x01\x70\x61\x09\x2c\x81\xc3\x88\xa2\xc3\x23\x6c\xd0\xe8\x86\x10\x5a\x84\x36\x4e\xf4\x49\x8a\x1c\x1b\x04\xbd\x29\x11\x88\x75\x60\x31\xe8\xc0\x89\x7d\xdd\xab\x22\x88\x9d\x27\x29\x29\x4c\xf9\x86\x40\x97\xe4\xbb\x8e\xe0\x0f\x18\xc8\xea\xf2\x8d\x18\x40\xc1\x5c\xd3\x9d\x52\x6d\xdb\xa6\xe5\xa1\x48\xad\x57\xb5\x27\x26\x95\x7b\xc7\x12\xbf\xd7\x9e\x50\x72\x81\xf2\xb8\x8f\x3c\xee\x23\x9d\x67\x89\x07\x74\x92\xbd\xd4\xf2\xc2\x9b\x16\x5c\x95\x62\x70\x45\x18\xd0\xf8\xaa\x42\x97\x63\x7e\x4d\xf7\xad\x46\xd6\x4f\x4f\x3e\xdd\x59\x2e\x9a\x4d\xa4\xfd\xe9\xf6\xed\x2f\xea\xf6\x57\x75\xfb\xb3\xca\x7d\x47\xd0\xd0\x85\xd6\xba\xf3\xbf\xad\x0f\xd2\x58\x25\x06\x30\x26\xd0\x10\x90\x9a\x92\xdf\x1c\x3d\xec\xbf\x4a\xa1\x09\x82\xf7\x7c\xfe\x32\x27\x3b\x02\x56\x9e\x11\x0e\x35\x3d\x4f\x47\x68\x1c\xc8\x1c\xa4\x0c\x15\xfc\xdd\xa5\x50\xd6\x90\x0c\xe3\xf9\x25\x77\xf1\xfe\xce\xd5\x03\xa8\x83\x0e\x2a\x34\x4e\x1d\x47\xd3\x68\xc8\xdd\x6b\xc5\x7e\xa4\x4b\xd0\x9d\xd2\x75\xad\xba\x68\x9c\x5e\x20\x1c\xcf\x6b\x94\x70\x68\x30\xe9\xcb\xb1\x7d\xd4\x0d\xf5\xb5\xe5\xf2\x7e\xfd\x39\x9b\xff\x99\xad\xef\xe7\xcb\xd5\x72\x74\x73\x81\x53\x29\x51\xa1\xf6\xce\xb7\x6e\x1d\x9f\xe9\xe6\x3c\x25\xe3\x11\x43\xf2\xec\x8a\x4f\xef\x92\xff\x9d\x42\x12\xa3\xfd\x8a\xd8\xca\x37\x8e\xe9\x94\xe9\xf3\x75\x00\xf9\xeb\x4c\x1a\xed\xce\x46\x5b\x16\xff\x05\x00\x00\xff\xff\xe2\x8b\xf9\x8a\x16\x05\x00\x00")

func clientBootstrapDaemonUpShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonUpSh,
		"client/bootstrap/daemon-up.sh",
	)
}

func clientBootstrapDaemonUpSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonUpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-up.sh", size: 1302, mode: os.FileMode(493), modTime: time.Unix(1515908524, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDockerSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x5b\x6f\xd3\x40\x10\x85\xdf\xf7\x57\x9c\x36\x51\x2f\x52\x6d\xb7\x11\xe2\x81\xaa\x48\x90\x46\x08\x51\x29\x52\x2e\xcf\x68\xe3\x9d\x8d\x57\x75\x76\xcd\xce\x98\x10\x21\xfe\x3b\xb2\x63\x1b\x0a\xe5\xe2\xa7\xf5\xce\x99\xd9\x6f\xce\xcc\xe8\x24\xdb\x38\x9f\x71\xa1\xd4\x08\x6f\x43\x10\x96\xa8\x2b\x86\xc6\x4e\xe7\x85\xf3\x04\x1b\x22\x4c\xc8\x1f\x29\x42\x7b\xd3\x1d\x93\x3c\xec\xaa\xc0\x94\xaa\x11\x56\x85\x63\x38\x46\x15\x49\xe4\x00\x5d\x56\x85\x4e\xf1\xde\xb3\xe8\xb2\x64\xe4\x75\x2c\x11\x7c\x79\x80\xb3\x70\x02\x13\x88\xfd\xb9\x80\xbe\x38\x96\x36\x9f\x58\xc8\x20\x78\xac\x37\xb5\x97\x1a\x37\x2f\xd3\xeb\x17\x57\x88\xf4\xa9\x76\x91\x18\x5c\x9b\xd0\x62\x68\x58\xda\x83\x85\x2a\x4e\x95\x62\x12\x24\xa4\xd4\xfd\x7c\xfa\x61\xb6\xf8\xb8\x9c\xaf\x17\xd3\xd9\xdd\x96\x24\x3d\x32\xa6\x79\xd8\xf5\xc1\xfb\xd9\x72\x75\x77\x9e\xc9\xae\xca\xb6\x24\x49\x27\xe0\xe2\xbc\x69\x7b\x5a\x50\xfe\xd8\xd0\xf5\x6d\x46\xd2\xe6\x00\x77\xec\x80\x4c\xaa\x9c\x45\xa1\xb9\xe8\x05\x93\xd7\x99\xa1\xcf\x99\xaf\xcb\xf2\x16\x52\x90\x57\x00\x9a\x86\x04\xd7\xca\xba\x5b\xa5\x2c\x49\x5e\x58\x57\xd2\xc5\x25\xbe\xb6\xd1\x11\xde\xc4\x2d\xbf\xea\xce\xc0\xf8\x06\x1c\xea\x98\x13\xd6\x8b\x87\x1f\xb7\x13\x18\x62\x71\x5e\x8b\x0b\x1e\x4d\x85\xb4\x0d\xf6\x04\xad\x9b\xcf\xbf\xdf\x7c\xad\x55\xad\x26\xb1\xbc\x7c\x68\x5e\x49\x02\x4e\xc7\x93\xd3\x23\x62\xd9\xd7\xd9\x6f\x49\xfe\x55\xa7\xd5\x24\xf3\x36\x1d\xe3\x9b\xae\x02\xd3\xa0\x8a\x24\x75\xf4\x38\x46\x9a\xbe\xbf\x35\x76\xbe\x23\xe9\x8d\x7a\x76\xe2\xce\xe2\xe4\xe2\x0f\x76\x5e\xfe\xc4\x31\xc2\x2a\x1e\x20\x01\x26\xec\x7d\x19\xb4\x41\xcd\xce\x6f\xbb\x85\x8a\x2d\xde\x55\xa7\xdc\x90\x0d\x91\x10\x89\x43\x94\x46\x24\xa1\x9f\x5f\x9f\x32\xf8\x38\xcc\x06\xe3\x27\xab\x33\xfc\x36\xcb\xf2\xab\x1f\xc5\x93\xe8\xef\x56\xe8\x4a\x92\xc6\xae\xba\x32\x5a\x08\x67\x67\xc3\x4d\x32\x6c\x52\x8b\x31\x64\xfc\x17\xc6\x5f\x01\xac\x53\xd6\xa9\xef\x01\x00\x00\xff\xff\x70\x6c\x7b\xfa\xc2\x03\x00\x00")

func clientBootstrapDockerShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDockerSh,
		"client/bootstrap/docker.sh",
	)
}

func clientBootstrapDockerSh() (*asset, error) {
	bytes, err := clientBootstrapDockerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/docker.sh", size: 962, mode: os.FileMode(493), modTime: time.Unix(1515908421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapKeygenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xdc\x30\x10\x46\xef\xf3\x2b\xbe\xb0\x81\x5c\x2a\xef\xbd\xa1\x0b\x6d\x13\xda\x3d\x74\x13\x68\x7a\x2a\xc5\x68\xad\x71\x34\xd8\x91\x8c\x66\xdc\xad\x2f\xfd\xed\xc5\xeb\xa5\x24\x64\x2f\xd1\x55\x6f\x78\x4f\x9a\xd5\xc5\x7a\x2f\x69\xad\x91\x68\x85\xfb\x92\xc3\xd8\xb0\xc2\x63\x18\xf7\xbd\x34\x6e\x28\xf2\xdb\x1b\xa3\xe3\xc9\x0d\x5e\x0a\x7c\x0a\xc8\xa3\x0d\xa3\x29\x2c\xf2\x89\x9b\xef\x2b\x22\x65\x83\x63\xa2\xed\x4d\x7d\x73\xfb\xfd\x61\xbb\xfb\xf8\xb0\xbd\xdb\x7d\xb8\xfc\x7a\xf7\xed\x76\x5d\xa9\xc6\xb5\x84\xba\xa8\xaf\x25\x71\x31\xf1\x75\xe0\xa1\xcf\x13\xdd\xff\xf8\x54\xbf\x71\xa6\x1a\xc6\xfd\x9c\xfc\x39\x72\xd3\x41\x5a\x04\x56\x93\xe4\x4d\x72\x42\x2b\x3d\xc3\xf7\x85\x7d\x98\xc0\x7f\x44\x4d\xdf\x93\xb4\xf8\x09\xd7\xe2\xf2\xa5\x09\xbf\xae\xe7\x87\x24\x02\x80\x23\x73\x71\xa4\x5e\x37\xbd\x20\xe7\xb3\xc2\xb6\x7d\xf6\x01\x08\x99\x35\xd9\x22\x7c\x87\x27\xdf\x31\xc4\xaa\xff\xb8\x6a\x74\x1d\x4f\x8f\x9c\xe0\xa6\x73\x25\x9b\x73\xd6\xe3\x78\x2b\xd7\xc4\xbd\x32\x2d\xda\x2f\x9c\xb8\x9c\xd6\x82\x83\x58\x44\xca\x18\xbc\xea\x21\x97\xb0\x08\x9f\xcb\x5e\x9b\x9c\xa1\xa8\x87\xdb\xe1\xea\x8a\x5a\x21\x3a\xe1\xda\xf8\x84\x47\xb1\x38\xee\xab\x26\x3f\x61\xb3\xc1\xdf\x65\x09\x5d\xca\x87\x54\xc7\xac\xa6\x44\x8d\xb7\xb3\xa9\xff\x02\x00\x00\xff\xff\x37\x00\x91\x4b\x4d\x02\x00\x00")

func clientBootstrapKeygenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapKeygenSh,
		"client/bootstrap/keygen.sh",
	)
}

func clientBootstrapKeygenSh() (*asset, error) {
	bytes, err := clientBootstrapKeygenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/keygen.sh", size: 589, mode: os.FileMode(493), modTime: time.Unix(1515908421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapTokenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x9e\x54\xe8\x69\xbb\x77\x21\x07\x11\x31\x45\x6d\x84\x08\x5e\x84\xb0\x4d\x46\x77\x89\x9d\x8d\x3b\xb3\x16\xff\xbd\x44\x1a\xed\x71\xe6\x9b\xef\x31\x6f\x75\xe1\xf6\x91\x9d\x04\x63\x84\x14\x96\x8c\x59\xe1\x8e\x98\xb2\x57\x12\x78\x46\xdb\xd6\xd0\x34\x12\xe3\x2d\x65\x14\x21\x1c\xa3\x06\x5c\x3f\x6d\x91\xe9\xb3\x90\xa8\x6c\xce\x1c\x78\x0c\x9e\x0e\x89\x4f\x52\x91\xc8\xef\xb8\x79\xd8\x6e\x8c\x94\x21\x61\x48\xfd\x48\x19\xb9\x30\xac\xcd\x07\xbc\x1a\x00\xb0\x5f\xb8\xac\x9b\xc7\xdb\x2b\xe7\xa7\xc9\x85\x24\xba\x00\x9a\x1f\xe8\xee\x77\xcd\xcb\xae\xab\x9b\xf6\xb9\xad\xd6\x7f\x37\x6e\x23\x12\xdc\xc8\xe9\xc8\xdd\x3c\xcb\xfa\xdf\x9a\xd3\xaa\xdf\xcc\x65\x67\x89\x35\x7f\x4f\x29\xb2\x56\x91\x29\x6b\xf4\x27\x54\xf6\xfd\x87\x2f\xdc\x87\xc9\x0f\x6e\x41\xe7\x2d\xcc\x4f\x00\x00\x00\xff\xff\x5f\x07\x1a\x29\x27\x01\x00\x00")

func clientBootstrapTokenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapTokenSh,
		"client/bootstrap/token.sh",
	)
}

func clientBootstrapTokenSh() (*asset, error) {
	bytes, err := clientBootstrapTokenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/token.sh", size: 295, mode: os.FileMode(493), modTime: time.Unix(1515908421, 0)}
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
	"client/bootstrap/daemon-down.sh": clientBootstrapDaemonDownSh,
	"client/bootstrap/daemon-up.sh": clientBootstrapDaemonUpSh,
	"client/bootstrap/docker.sh": clientBootstrapDockerSh,
	"client/bootstrap/keygen.sh": clientBootstrapKeygenSh,
	"client/bootstrap/token.sh": clientBootstrapTokenSh,
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
	"client": &bintree{nil, map[string]*bintree{
		"bootstrap": &bintree{nil, map[string]*bintree{
			"daemon-down.sh": &bintree{clientBootstrapDaemonDownSh, map[string]*bintree{}},
			"daemon-up.sh": &bintree{clientBootstrapDaemonUpSh, map[string]*bintree{}},
			"docker.sh": &bintree{clientBootstrapDockerSh, map[string]*bintree{}},
			"keygen.sh": &bintree{clientBootstrapKeygenSh, map[string]*bintree{}},
			"token.sh": &bintree{clientBootstrapTokenSh, map[string]*bintree{}},
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
