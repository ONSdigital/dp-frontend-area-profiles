package utils

import (
	assets "github.com/ONSdigital/dp-frontend-area-profiles/assets"
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

// AssetDIR similiar to assetFS() but requires a directory name
func AssetDIR(dir string) *assetfs.AssetFS {
	return &assetfs.AssetFS{Asset: assets.Asset, AssetDir: assets.AssetDir, AssetInfo: assets.AssetInfo, Prefix: dir}
}
