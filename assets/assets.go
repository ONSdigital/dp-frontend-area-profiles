// assets in locale and templates folders are converted into Go source code with go-bindata
// the data.go file in this package is auto generated through the generate-debug and generate-prod tasks in the Makefile
package assets

import "io/fs"

func Asset(name string) ([]byte, error) {
	return []byte{}, nil
}

func AssetNames() []string {
	return []string{}
}

func AssetDir(path string) ([]string, error) {
	return []string{}, nil
}

func AssetInfo(path string) (fs.FileInfo, error) {
	return nil, nil
}
