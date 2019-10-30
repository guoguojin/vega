// Package keys is code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// keys/jwt.priv
// keys/jwt.pub
package keys

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

var _keysJwtPriv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x97\x37\x0e\xac\x88\xa2\x44\x73\x56\x31\x39\xfa\xc2\xbb\x60\x02\xbc\xf7\x9e\x0c\x4f\x43\xe3\x7d\xaf\xfe\xeb\xce\x4d\x5f\xa5\x95\x1d\x95\x4a\x3a\xff\xf7\x27\x9c\x28\xab\xd6\x3f\x9e\xcf\xfe\xe3\x78\x6a\xc4\x06\xe2\x3f\xba\x98\xfe\xd7\x00\xa6\xaa\x6a\x7a\xab\x72\x2c\xab\xf3\x6c\x2b\xb2\x6f\x4d\x69\x69\xb3\xde\xba\x9c\xb9\x69\x69\xef\xdc\x73\x28\x33\x36\x0f\x4d\x78\xed\xb7\x96\x4e\x1b\x85\x6e\xe6\x33\xa2\x11\xa3\x6b\x05\x77\xb7\x80\x75\xf3\xbf\x91\xa9\xf6\x9e\xda\x6c\x47\x0a\x84\x95\x5a\x17\x50\x11\x5c\xe1\xbd\xfd\x47\xc6\x6d\xff\xfc\xf5\xc4\xe6\x1c\x36\x55\xd2\x49\x3b\xed\xfd\xa7\x95\xb1\xce\x33\x6b\x70\x73\x5d\xf5\xb0\x81\x8c\xc6\x2e\x81\x4d\xba\x49\xb7\x5f\x56\x08\xeb\xaf\x6e\xea\x46\x0e\x09\xd5\x28\x19\x45\x48\xb9\x0b\x86\xdf\x6a\xb5\xfb\x4f\x23\x42\x99\x32\x7d\xc0\xe7\xa2\x98\x61\xee\x82\x6f\xcb\xa8\x52\xaf\x03\x71\x57\x1a\x57\xde\x8b\xdd\xf8\x7a\x23\xb2\x5e\x15\xc7\xbe\xf5\x2c\x15\x2c\x8a\x67\xe3\xaf\x2c\x8b\xea\x49\xc5\xc0\x2f\x41\x37\x47\x33\xfb\x61\x73\x92\x1f\x94\xe0\x80\x4a\xc3\xab\xa3\x93\x36\x00\x56\x44\x74\x33\xad\x6a\xc2\x24\xef\x8c\x75\x07\xec\xbc\x13\x8d\x43\xa7\xa3\x34\x8f\x12\x8e\x18\xfe\xf8\x79\x0f\x78\xc5\xbf\x97\x1d\x0b\x23\xfd\x71\x9f\xed\x25\xf7\xf4\x72\x62\x6e\xc1\x8f\x58\x02\x6a\xa2\xde\xeb\x16\x4c\xeb\xad\x35\xf9\x48\x66\xa4\x1b\x73\x2d\xdc\x80\xd5\x23\xe0\x76\x42\x1a\x96\x27\xb1\x83\xeb\xe2\xe8\x1b\xcf\xa2\x99\x4d\x71\x89\x6d\x08\xbb\x89\x26\x92\xd5\xd8\x87\xa2\x80\x04\x3c\x57\x1f\x92\x2d\xb1\xe6\xce\x53\xb1\x62\x11\xec\xf5\x80\xf3\x28\x26\xb6\xbc\xd0\xdc\x55\x17\xe6\x3f\xb4\xa7\x42\x4d\x1b\xcb\xdc\xc3\x07\xce\xc8\x2e\xbc\x9b\x9d\xb8\xf1\x7b\xcc\xdc\x97\x01\x2c\x7c\x9b\x54\xb0\xa1\x8a\x56\xb9\xaa\x99\xd1\x8f\x2f\x8a\xdc\xf7\x05\x0d\x61\x82\xb5\xf3\x49\x13\xdb\x57\x61\x90\x4e\xc9\x24\xf0\xc2\xa9\x9e\x9e\xa2\xc7\x38\xbe\x4e\xac\xb3\x12\xea\x60\x75\x00\x89\xc6\x25\x86\x1e\xce\xa3\x52\x44\x15\x49\xa0\x8a\xc8\x50\xdc\xed\x4c\xe9\xe3\xa3\x53\x8c\xd8\x84\x72\xd7\x68\x7a\xe2\x6b\xda\xe2\x75\xdb\x7b\xa1\x94\x91\x6f\x58\x90\x08\x07\x56\x83\xc1\xd5\x80\x5e\xca\xce\x97\xd7\x30\x3b\x09\x33\x68\x58\x09\x99\x25\x1e\x07\x21\xb8\xda\xe4\x86\x8b\x28\x91\x7d\xf8\x5e\x2c\x1e\x86\xf8\xf1\xc6\xc8\x5e\xf7\x25\xcf\xab\xd3\x3b\x07\x94\x85\x98\xb9\x8c\xff\x00\x10\x84\x7e\xb2\x7b\x73\x37\x6d\x2c\x6c\x83\xe8\x24\x28\x09\x38\x3a\x83\xd8\x47\xa1\x28\xfb\x5b\x5c\x7e\x04\x27\xb0\x21\xa2\xdb\xb2\x52\x6c\x5d\x28\xc6\xc5\x69\x4b\x13\xa3\xbb\xc9\xb3\xb7\xc8\x02\xac\xfb\x77\xc2\x3b\xc2\x19\x23\x39\x85\x2f\xbd\x5f\x42\x2d\x2b\x58\x50\xcf\x6e\x49\xde\xb3\x6c\x38\x55\x1e\xb4\xb0\x33\xab\xa2\xe2\xb3\xeb\xcf\x89\x2e\x8e\x3c\x36\x7d\x02\x11\x9e\x0b\x00\x78\xc3\xa8\xc2\x07\x55\xb5\xd1\x69\x4e\x8a\x35\x83\x87\x3e\xe4\x64\xb6\x51\xcb\x59\x30\x7b\x9c\xf1\xf7\xc5\xfd\xe9\x5b\x7d\x06\x32\x25\x17\xe7\x47\xdc\x09\xbe\x11\xad\x38\xd3\xa2\xfc\x7a\x5a\x03\x90\xe9\xfb\x81\x20\x77\xfb\x92\x23\x84\x71\x13\x45\x70\x34\xb1\x61\xa6\x57\xd4\xcb\xae\xa3\xee\xab\x62\x27\x89\xbf\xf3\x22\x3b\x7b\x39\x5d\xce\x4e\x60\x83\xb1\x6d\xf2\x82\x2d\x78\x1b\x9f\xb9\x0c\x5c\x7e\xdd\x7f\xed\xc0\x2e\x75\x9c\xcb\x7c\x90\x65\xc6\x42\x2f\x7a\xa2\x16\x64\xee\xc3\xd2\x91\xba\xdd\x13\x49\x44\x27\xe2\x5b\x08\x91\x0c\x17\x65\xef\xf1\x2d\x4a\xcb\x8e\x4c\x8e\xb1\x4f\x75\x02\xc8\x32\xc3\xbd\x30\xf9\x5e\x8f\x0c\x92\x07\x4d\x0e\xcd\xd0\x27\x17\xec\xed\x5a\x7d\xe0\x31\x96\x14\x03\xae\x87\x4e\xb1\x8a\xe2\x6d\x98\x55\x1d\x88\xc8\x68\x1b\x91\x1d\xcb\x26\xf6\x12\xa7\x40\x02\x65\x45\x4b\x66\x20\xa5\x01\xef\x8b\x9c\x5b\xb9\x39\x89\x2e\x04\x47\xfd\x90\xfe\xf0\xd1\x43\xf7\x73\x35\xe7\xc7\x57\x4e\xb4\x57\x62\x8e\x31\x77\x0f\xc2\x66\xc2\xf3\x92\x05\xd9\x82\xf4\x92\x81\x00\xfe\x76\x43\xf9\xf2\x5b\x61\xfc\x22\xde\x48\xa3\x82\x0e\xeb\x4c\xe9\xab\x8e\xe6\x4f\x9b\xf9\x08\xc4\x5e\xe2\x37\x26\xc2\xf3\x9c\x86\x81\x1c\x9c\x78\x95\xb0\xb4\xe1\xfa\x05\x2b\x4f\xe6\xc2\xe1\x80\xc3\x39\xe8\x72\x14\x70\x38\xc0\xe1\x3a\x6a\x7c\x92\xeb\x95\x2f\x2a\xc2\xc3\x2d\xf6\x28\x53\x86\x56\x44\xf2\xbb\x74\x24\xc7\x13\x12\x95\x45\xaa\x76\xcd\x12\xa4\xea\x26\xe8\x63\x23\xc2\x35\x32\x60\x48\x6a\x6e\x6b\x4d\x64\x3b\xa8\x37\xfe\x7a\x5c\x6a\x90\xbd\x5a\xc2\x61\xc6\xfb\x5d\x32\x89\xcf\xee\x0c\x9c\xbe\x5e\xbd\xba\xac\xcc\x52\xac\x13\x49\x46\x0c\x4e\xe0\x2e\x2c\x42\xbb\xfb\x13\x07\x60\xd6\xd7\x8a\x19\xa7\x8f\x46\xe7\x0d\xba\x94\xbf\x4e\x62\x61\x68\x25\x3d\xed\xd7\x3e\xb7\x13\xe1\xde\x87\x1e\xe4\xaf\xba\xc9\x60\x3a\x23\x8c\x72\xd1\x64\xab\x69\x49\x29\xf3\xcf\xa3\xb9\x82\xaa\x02\xeb\xd3\x9f\x48\x1d\x61\xb0\x7f\xb2\x0b\x49\xbe\x8f\x6b\xa4\xba\xee\x05\x63\x50\x4f\xd2\x08\x39\x96\xc7\x0b\x76\x73\xe1\x0e\x26\x92\x6a\x37\x73\x0c\x2a\xf2\x6d\x2b\x72\xac\xb3\x3e\x4d\xc8\xa3\xc0\x37\x88\xe4\xa9\x90\x6e\xe3\x66\x09\x28\x26\x0e\x8d\x41\xdd\x30\x73\x8f\xfe\xd9\xed\x7e\xc8\x56\x07\xb2\xbe\xc4\x07\xcc\x0e\x76\x4b\x02\xa3\xc6\x0c\xd3\x19\x17\x68\x4b\x97\x8d\x33\xbe\x90\x85\x03\x3f\x25\xd6\x96\x1a\xc5\x70\xdf\x29\x58\xcb\xdf\x26\x46\x58\x4e\xfe\xee\x82\x79\xea\x53\xc7\x0a\x03\x9d\x0f\x79\x4b\xf2\x45\x35\x4f\x8c\x98\x3a\xd5\xb5\x8f\x25\x14\xab\xaa\xc7\x15\xb2\xc1\x9e\x01\xea\x22\x72\x99\x37\x9f\x76\x9e\xda\xe0\x69\x1d\xe3\x1b\xfe\x68\xb6\x96\xdb\x48\xc6\xea\xd9\xa1\x28\xfb\x23\x51\x54\xd8\x71\x5e\xc9\x1f\x4c\x9d\x8d\x7e\xce\x2a\x30\x83\x5e\xf2\x6f\x8c\xf4\xa0\x01\x2e\x53\xe3\x2a\x8c\x1a\xbf\x25\xde\x1d\x39\x34\xbb\xd3\x13\xbd\x3d\xd4\x76\x94\x64\x6e\x51\x98\x1a\x50\x35\xcc\x79\x6d\xc5\x32\xa8\x1f\xcf\xd3\xef\x98\x89\xbe\x9a\xa9\x4d\x0d\x97\xe9\x76\xd1\x01\xd1\xe4\x23\xd5\xa2\xd9\xc4\x37\x56\x2e\x0f\x3a\xf9\x9b\x97\x0b\x76\x2c\xb9\xf4\xc0\x5b\x0c\x4e\xdc\xd9\x73\x92\x2d\x50\x4b\x6f\x5b\xa8\xc6\xc7\x9a\xdb\x6c\x69\xf1\xe6\xb7\x5c\xc8\x37\x92\x1d\x80\x7c\x0c\x9f\x81\x3b\x63\x94\x5d\x0e\xeb\x17\xfe\x25\x6c\xe2\x88\xc8\x53\xb9\x8e\x41\xb2\x13\x9c\x51\x46\x5f\x4e\x8b\x53\xdc\xe6\x70\x91\xcd\x69\x87\xf8\xbd\x3e\xc6\x5e\x58\xf9\xcd\x2e\xc0\xe8\x30\x27\xf9\x6c\x59\x27\x69\xfe\xe1\x15\x5e\x92\x20\xe4\x2d\xc7\x98\xf8\xc4\xb6\xc6\xd5\x0e\xb6\x1c\x31\x5d\x47\x2d\x41\xf3\xae\x9f\xe3\x22\xfb\x3d\x05\xd1\xd0\xe8\xa5\x8c\x55\x57\xe4\x6f\x40\x60\x88\x4c\xf9\x45\xef\xc7\x80\x31\x93\xc0\xdf\xf2\xe2\x70\xa8\xd6\x71\x5a\x11\x3f\x66\xdb\x91\xca\xaf\x96\x38\xb3\x5c\xfa\x63\x81\x64\x1f\x54\xab\xf7\x0e\xe8\x22\xf7\x23\x5e\xf2\xc1\x31\xb3\x01\x0d\xd9\xa0\xc1\x19\x10\x43\xc4\x9b\x16\xfc\xcc\xc2\x2c\x2f\x2f\xb1\xdf\x07\xad\xd1\x85\x19\x29\x8e\x92\xc0\xb4\xa1\x6c\x58\xee\xeb\xa3\x5a\x66\x8b\x00\x3d\x5f\x27\x15\xae\x7c\x50\x48\x61\xee\x80\x04\x2a\xa4\xb9\x52\x53\x91\x8b\x3e\xdf\xea\x83\x23\x7b\xaa\xfe\xb2\x1a\x97\xc7\xdb\xb4\x2c\xd8\xb3\x0a\xf2\x7d\x3b\x2f\x72\xd1\x60\x89\xc4\x73\xcc\xf5\xfd\xfc\x68\x99\x19\xdb\x9b\xff\xce\x04\x09\x28\x7f\xae\x1b\x95\x91\x34\xc7\x40\x41\xe9\xb8\x4c\x7c\xb5\x40\x52\x29\x6a\x9e\x53\xbf\x65\x23\xd6\x6f\x7c\x67\xc8\x63\xda\x2b\xff\x02\xb6\xbb\xb3\xcb\xb0\x25\x7f\xa0\xc7\xe4\xd3\x0f\xf0\x8a\xe2\xa7\x7d\x26\xdb\xb2\xcb\xa8\xda\x50\x4c\x86\x0a\x47\xf5\xfc\x4f\x28\x15\x6c\xa7\x69\xcd\xe2\x48\x83\xfd\xce\xf5\x29\x83\xd6\xe5\x98\xbf\xf7\xc3\xef\x5d\x45\x0d\xed\x37\xa0\xf0\x94\x3e\x00\x43\xc3\x4d\x5c\x93\x2c\x72\xed\x4c\x70\xe8\x52\x27\xf0\xa2\x50\x60\x36\x7f\x7a\x8b\x85\x47\x14\x4b\x90\xa4\x11\x2f\x99\x24\x61\xf8\x53\xed\xbc\x5e\x8f\xec\x29\xb3\x1a\xec\x24\x6d\x9c\xe1\x24\xe0\xc3\xa3\x42\x33\x29\x5f\x3d\x0e\x2f\x30\x41\xdc\x23\x68\x96\xdd\xd8\xcd\xbc\x72\x5d\x78\x4a\xfd\x66\xe4\xeb\x30\x3d\x3b\x7f\x17\xf4\x0f\x05\x8f\xc5\xc4\x54\xe5\xc1\xc8\x58\x2b\x7a\x0b\xe3\x59\xc0\x2d\x7d\x23\xb5\x5a\xac\x7d\xa7\xf5\x1f\x25\x62\x78\x50\x81\xef\x6f\xcd\xf2\xb8\xde\xf6\x5d\xe9\x56\xe1\xd6\x91\x72\xd4\xc6\xbb\x4f\x67\x44\x50\xd1\x63\x32\x55\xd3\x9c\x48\xe8\x2a\x1a\x74\xcf\x80\xf9\xc3\xec\x78\x3d\x4d\x3f\x30\x8e\x2e\xa1\x76\x17\xb6\xfa\x29\x1d\x71\xe0\xf1\xc1\x49\xa4\x71\x9b\x6e\x72\x2a\x6a\xa8\xc4\xe6\xe7\x85\xa7\x1e\xe3\xb6\xb5\x20\xa8\x46\x21\xc9\x31\x32\x13\x3d\x00\xd8\x16\xb9\x6f\x0c\xfe\x43\x98\x55\x36\x9f\x54\x5e\xe8\xda\xc5\x1b\xf9\xb1\x6f\x5a\xfe\x66\xb5\xe8\xf2\x82\x77\x75\x6c\x2e\x29\x5d\xe5\xf4\xc0\x5d\xbe\xb6\xae\xc3\x9a\x38\x76\x3e\xc5\xcb\x15\x80\xe6\xdc\xd7\xb3\x4f\x3f\x40\xd9\x42\x10\x97\xe7\xe8\x04\xa8\xb5\x48\x78\x14\xe9\x46\x15\x5a\x95\x5c\x7e\xde\x0b\x52\x35\x71\x81\x9f\xb8\x56\x04\x38\x97\x82\x9c\x25\x9d\xe9\x8b\x6a\xc2\x8b\x2c\x00\x14\xec\xb5\x92\xa1\x4a\xdf\xc5\xb2\x6b\x81\x93\x35\xc8\x6e\x40\x0b\xc9\xfa\x62\xe6\x8c\xd4\x5f\xc2\x8b\x4e\x9e\xfc\x61\x9f\x56\xe1\xfd\xec\xd5\x47\xad\xe1\x26\x9a\xab\xfa\xba\x1b\xfb\x37\xe6\x01\x6c\xb2\x44\xba\x87\x77\x97\xf1\x7b\x86\x85\x7d\x95\xfd\x7d\xf5\x94\xb1\xf8\xb9\xce\x50\xc8\x7f\xdb\x4c\x60\x50\xbc\xb3\x19\x4e\x1f\xfd\xb3\xb5\xf6\x5b\x78\x0a\x93\xf4\xea\x95\x64\x3c\x77\x05\x01\x77\x14\xd0\x5f\x34\xb5\x64\x9a\xb5\xc2\xbd\x0e\x86\xf6\xc1\xcf\x89\x4c\x87\xa0\x75\x8b\xf8\xbd\x09\xcc\xf8\x85\xc9\x6b\x98\xf8\x72\x25\xe3\xbc\x72\x58\xfc\xaa\xcd\x0e\x35\x77\xa2\xc9\x18\x3f\x02\x5c\xab\x70\x23\xb1\x2e\x26\xd3\xbd\x0a\xcb\x83\x6b\x6e\xaf\x7e\x48\x25\xb9\x10\x37\xcd\x18\x21\xeb\xad\xf3\xac\x2b\xb2\xc8\x8d\x16\x71\xf0\x5a\x75\xd7\x68\xfe\x65\x8e\xf7\x11\x63\xcd\x9b\x2b\x80\x52\x2d\xdf\x2a\xab\xc5\xa2\x16\x0a\x93\xc0\x68\x77\xf6\xa7\x54\x4a\xbc\x33\x2a\xd1\x2f\x6d\x76\xce\x83\xb6\x8a\x5d\xe1\x5e\x8f\x7e\xc1\x1f\xfa\xda\x77\x06\x19\xfd\x2b\x9b\xcb\xbe\xc3\xa1\x9d\x00\x62\xdd\x11\xd8\x91\x0e\x98\xfe\x85\x44\x04\x92\x48\x48\x99\xb1\xe6\x8c\xc2\x55\x1e\x25\xe6\x73\x57\x36\x15\x0b\x93\x89\x1e\xfc\xf5\xe1\x96\x2d\xc9\x94\xa5\xcb\xca\x15\xa7\xf0\x42\xf7\x3e\xdb\x01\x2c\xb9\x7c\xbc\xb4\x72\xa8\x86\x61\x2b\xcd\xdd\x3a\xa9\xd2\xc6\x4a\x70\x33\x89\x93\xf5\xeb\x8e\xd0\xa4\x28\xef\x57\x6d\xa2\x15\xd3\x6b\x17\x22\xa6\xf0\xde\x33\xf7\x9a\x6c\xa3\x97\x3f\xf5\xb7\x04\x44\x65\xbc\x72\x30\x81\xd5\x2e\x02\xe5\xaa\xbd\x88\x40\xa3\xb6\xc2\xcb\xf6\xac\x28\x7e\x04\xe5\x88\x6a\x41\xdd\x90\x87\x31\x8b\xf2\x58\x70\x28\x77\xa9\xec\x45\xc5\x49\x36\xfa\x99\xf3\x42\xad\x02\x10\x72\x65\xe1\x24\x6e\x67\x6c\x93\x70\xb3\x57\x01\x5a\x1c\x20\xde\x2c\x7a\x31\xcd\x4a\xc4\x21\x43\x4a\xb1\xd2\xfb\x85\xed\xe7\xe9\x55\x0d\x0b\xc9\x4b\x15\x21\xf3\xca\x89\x2a\x80\xf6\xfb\xdf\x7f\x81\xff\xb4\x43\xb4\x84\xff\xad\x23\xff\x1f\x00\x00\xff\xff\x11\x98\x1d\x9c\xaf\x0c\x00\x00")

func keysJwtPrivBytes() ([]byte, error) {
	return bindataRead(
		_keysJwtPriv,
		"keys/jwt.priv",
	)
}

func keysJwtPriv() (*asset, error) {
	bytes, err := keysJwtPrivBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/jwt.priv", size: 3247, mode: os.FileMode(420), modTime: time.Unix(1558447430, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _keysJwtPub = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x93\xb7\xb2\xa3\x4a\x00\x05\x73\xbe\x62\x73\xea\x15\xc2\x09\x08\x36\x98\x19\xac\x70\xc2\x09\x93\x09\x18\x90\x40\x42\xc2\x9b\xaf\x7f\x75\x6f\xba\x27\x3c\x9d\x76\xff\xf7\x33\xa8\x68\x86\xf3\xe7\x1a\x41\xcb\x40\x7f\x4c\x25\xfd\x3d\x09\xdb\x30\x90\xd1\x00\x07\xd6\x6d\xff\x68\x9f\x9a\xb4\x9e\x20\xf0\x14\x15\x00\x17\x81\x5a\x04\x3f\x1c\xd5\x26\x02\xb5\x02\x76\x2c\x5c\xd2\xaa\x5f\x4d\x2d\xf3\xd2\xc2\x1d\xe1\x36\xe9\xc4\x87\xfd\xb4\x55\xb4\x8c\xeb\x25\xed\x06\x81\x19\xec\xed\xcd\xdc\x24\xf3\x92\xc3\xb5\x76\x56\x74\xbc\xa5\x72\x6c\x84\xc1\xbd\xaa\xa1\xdc\x0b\xfd\x97\xd4\x65\x4f\xde\xd7\x60\xd3\x38\x37\x98\x8f\x86\xe0\x87\xeb\xe4\x0a\x85\x98\xd4\xdd\xd8\x3c\x6b\x8d\x7d\xf8\x36\x26\x07\xcf\x33\x26\x37\x13\xd9\x45\x06\xc9\xa3\x33\xdd\x1d\xc8\x11\x7e\x99\xb6\x69\xdd\x29\xb9\x7c\xab\x56\x1e\x09\xde\x97\xe5\x56\xc2\x28\xc7\x60\xab\x14\x2a\xd3\xbb\x27\xb9\x2d\x82\xd4\x7e\x1e\xe1\xab\x96\x0c\xb5\x31\xe3\x47\x61\x2d\xf7\x46\x79\xbc\x77\xff\x4d\xf7\x4b\x09\xc1\x8e\x3f\x6a\x0e\x18\x2e\x7b\x1f\x45\x91\x97\x5b\x4a\x28\x61\x50\x90\xde\x9d\xc9\xdc\x0d\xdc\xcf\xa8\xd5\xc3\x89\x2a\x2c\x1f\xdf\x66\xd1\x02\x0a\x6d\xda\x69\x89\x79\xfb\xbc\x66\xc0\x6b\xd9\x79\x4d\x2e\x90\xe9\xa6\xc2\x9e\x8a\xd3\x4d\x42\xd3\xe1\x13\x1b\xb9\xc4\xc7\x0e\xde\xb9\x95\x1e\xf0\x39\xec\xe7\x31\x5d\xae\x31\xfc\x72\x53\xac\x62\x1e\x8f\xb8\x26\x53\x3c\xd4\x36\xba\x69\x92\xba\xb2\x9e\xc3\x59\x27\x63\x0a\xe1\xc8\xab\xed\x77\x4b\x5c\x22\x5c\x16\x28\xae\x5c\x76\xfb\x80\x94\x53\x41\xc5\xbb\xd5\xad\x3b\x83\x0b\xd8\x04\x21\x21\xe7\x3e\xa0\x34\x47\xc1\x70\x9e\x75\x27\x56\xc8\xc6\x0c\xa1\x2f\x48\xb1\xe3\x47\xf6\x68\x78\x27\xf4\x24\x44\xdf\xa0\xaa\x3a\xd6\xe0\x86\xc2\xeb\x1b\x7c\x91\x97\xcd\x9c\x75\x6c\xf6\x3d\xd0\xd8\x68\xaf\x52\xd9\xa5\x4a\xd1\x80\x65\xf5\x61\x9e\x81\xa2\xc0\xd7\x4e\x5a\x72\x77\xba\xcc\x5b\x9a\xb8\x81\x41\x9c\x48\x31\x3d\x27\xa1\x1f\x75\xb8\xdb\xf2\x86\x85\x08\x27\xce\x5c\xca\x38\xec\xaf\xaa\xc8\xa9\x92\xd8\xce\x53\xa9\x2b\x06\x9d\x50\x25\x9f\x31\x9c\xf7\xb0\xd5\x67\xc0\x74\x31\xed\xf2\xfa\x8a\x09\x26\x9d\xb9\x3e\xad\x39\x5c\x37\x7e\xa4\x66\xe7\x3d\xca\xcf\x34\x24\xcb\xd6\x82\xd8\x2c\xb4\xeb\x0b\x5d\x58\x37\x89\x32\xaa\xed\x79\x0d\xf0\xdb\x95\xe6\x21\xb6\x61\xbb\xf0\x05\x3d\xb6\xaf\x85\x00\x5c\x14\x71\xd3\x1e\xd3\x23\x6e\x0a\x84\x8c\x6e\xff\x84\x82\x43\xdb\x77\x8d\x3b\x48\x92\x3a\x34\x6f\x85\xab\x68\x7d\x41\x45\x9b\x67\x52\x95\x39\xe6\x43\xb2\x4f\x5d\x10\xdc\x57\xbe\x04\xb7\x13\x91\x9c\x2c\x85\x19\xbe\xbd\x00\x70\xae\x5b\x0b\xbc\x7c\xab\x98\x19\x6d\x04\x56\x05\x00\xef\xef\x5f\xe2\xd7\x78\xc5\x91\xff\x89\xe0\xff\x00\x00\x00\xff\xff\xf2\xf2\xc7\x4e\x20\x03\x00\x00")

func keysJwtPubBytes() ([]byte, error) {
	return bindataRead(
		_keysJwtPub,
		"keys/jwt.pub",
	)
}

func keysJwtPub() (*asset, error) {
	bytes, err := keysJwtPubBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/jwt.pub", size: 800, mode: os.FileMode(420), modTime: time.Unix(1558447447, 0)}
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
	"keys/jwt.priv": keysJwtPriv,
	"keys/jwt.pub":  keysJwtPub,
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
	"keys": &bintree{nil, map[string]*bintree{
		"jwt.priv": &bintree{keysJwtPriv, map[string]*bintree{}},
		"jwt.pub":  &bintree{keysJwtPub, map[string]*bintree{}},
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
