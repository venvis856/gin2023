package filesystem

import (
	"context"
	"errors"
	"os"
)

type Driver interface {
	//Copy(oldFile, newFile string) error
	//Delete(file ...string) error
	//DeleteDirectory(directory string) error
	//Directories(path string) ([]string, error)
	Exists(ctx context.Context, file string) bool
	//Files(path string) ([]string, error)
	//Get(file string) (string, error)
	//LastModified(file string) (time.Time, error)
	//MakeDirectory(directory string) error
	//MimeType(file string) (string, error)
	//Missing(file string) bool
	//Move(oldFile, newFile string) error
	//Path(file string) string
	PutFile(ctx context.Context, dist string, src *os.File) error
	Put(ctx context.Context, dist, src string) error
	//PutFile(path string, source File) (string, error)
	//PutFileAs(path string, source File, name string) (string, error)
	//Size(file string) (int64, error)
	//TemporaryUrl(file string, time time.Time) (string, error)
	//WithContext(ctx context.Context) Filesystem
	Url(file string) string
}

type Config struct {
	Default string `help:"默认存错磁盘" devDefault:"cos" default:"cos"`
	Disks   struct {
		Cos   CosConfig
		Local LocalConfig
	}
}

func (c *Config) NewDriver(disk string) (Driver, error) {
	switch disk {
	case "cos":
		return NewCos(c.Disks.Cos)
	case "local":
		return NewLocal(c.Disks.Local)
	}
	return nil, errors.New("文件驱动失败")
}
