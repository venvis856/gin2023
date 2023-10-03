package filesystem

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LocalConfig struct {
	RootDir string `help:"存储路径" default:"$ROOT/storage"`
	BaseUrl string `help:"访问的url" default:"http://localhost"`
}

type Local struct {
	rootDir string
	baseUrl string
}

func NewLocal(conf LocalConfig) (*Local, error) {
	dir, err := filepath.Abs(conf.RootDir)
	if err != nil {
		return nil, err
	}
	return &Local{rootDir: dir, baseUrl: strings.TrimSuffix(conf.BaseUrl, "/")}, nil
}

func (l *Local) PutFile(ctx context.Context, dist string, src *os.File) (err error) {
	dist = filepath.Join(l.rootDir, dist)
	dir := filepath.Dir(dist)
	err = os.MkdirAll(dir, 0766)
	if err != nil {
		return err
	}
	distFile, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer func() {
		if _err := distFile.Close(); err == nil && _err != nil {
			err = _err
		}
	}()
	n, err := io.Copy(distFile, src)
	fmt.Println(222222, n, src.Name(), src, err)
	return err
}

func (l *Local) Put(ctx context.Context, dist string, src string) (err error) {
	_, err = os.Stat(src)
	if err != nil {
		return
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	return l.PutFile(ctx, dist, srcFile)
}

func (l *Local) Url(fileName string) string {
	return l.baseUrl + "/" + strings.TrimPrefix(fileName, "/")
}

func (l *Local) Exists(ctx context.Context, file string) bool {
	_, err := os.Stat(l.path(file))
	if err != nil {
		return false
	}
	return true
}

func (l *Local) path(fileName string) string {
	return filepath.Join(l.rootDir, fileName)
}
