package assets

import "embed"

//go:embed config
var configFiles embed.FS

func GetConfigFiles() embed.FS {
	return configFiles
}