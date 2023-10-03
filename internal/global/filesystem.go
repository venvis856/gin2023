package global

import (
	"gin/internal/library/filesystem"
)

var Filesystem *filesystem.Storage

func InitFilesystem(conf *filesystem.Config) (err error) {
	Filesystem, err = filesystem.NewStorage(conf)
	return
}
