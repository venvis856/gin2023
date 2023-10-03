package helper

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// IsExist 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// FileExist 判断文件是否存在
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// Mkdir 创建目录
func Mkdir(path string, perm os.FileMode) error {
	if path == "" {
		return errors.New("path empty")
	}
	// recordID 要过虑非正常路径，如&之类的字符
	filterStr := []string{"..", "&", ":", ";", "|", "$", "%", "?", "\r", "\n", "`", ",", " "}
	for _, str := range filterStr {
		if strings.Contains(path, str) {
			return errors.New("目标路径存在特殊字符: " + str)
		}
	}
	err := os.MkdirAll(path, perm)
	return err
}

// Copy 复制文件
func Copy(src, dst string) (int64, error) {

	// dst 要过虑非正常路径，如&之类的字符
	filterStr := []string{"..", "&", ":", ";", "|", "$", "%", "?", "\r", "\n", "`", ","}
	for _, str := range filterStr {
		if strings.Contains(dst, str) {
			return 0, errors.New("目标路径存在特殊字符: " + str)
		}
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// 递归获取目录树
// isDir 是否只要目录
func ReadDirTree(dirPath string, onlyDir bool, showPath string) ([]map[string]interface{}, error) {
	treeData := []map[string]interface{}{}
	flist, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return treeData, err
	}
	for _, f := range flist {
		if f.IsDir() {
			tmp := map[string]interface{}{
				"label":    f.Name(),
				"value":    dirPath + "/" + f.Name(),
				"size":     f.Size(),
				"modtime":  f.ModTime(),
				"isDir":    f.IsDir(),
				"showPath": showPath + "/" + f.Name(),
			}
			child, _ := ReadDirTree(dirPath+"/"+f.Name(), onlyDir, showPath+"/"+f.Name())
			if len(child) > 0 {
				tmp["children"] = child
			}
			treeData = append(treeData, tmp)
		} else if !onlyDir {
			treeData = append(treeData, map[string]interface{}{
				"label":   f.Name(),
				"value":   dirPath + "/" + f.Name(),
				"size":    f.Size(),
				"modtime": f.ModTime(),
				"isDir":   f.IsDir(),
			})
		}
	}
	return treeData, nil
}

// 获取目录下所有文件
func ReadDirFiles(dirPath string) ([]map[string]interface{}, error) {
	treeData := []map[string]interface{}{}
	flist, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return treeData, err
	}
	for _, f := range flist {
		if !f.IsDir() {
			treeData = append(treeData, map[string]interface{}{
				"label":   f.Name(),
				"value":   dirPath + "/" + f.Name(),
				"size":    f.Size(),
				"modtime": f.ModTime(),
				"isDir":   f.IsDir(),
			})
		}
	}
	return treeData, nil
}

// 获取文件内容的类型os.open()后调用
func GetFileContentType(out *os.File) (string, error) {
	// 只需前512 个字节即可
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	t := http.DetectContentType(buffer)
	return t, nil
}

// FormatFileSize 字节的单位转换 保留两位小数  用于硬盘数据监控
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// IsNotExistMkDir 如果目录不存在则创建
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir 创建目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// CheckNotExist 检查目录是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

/**创建文件**/
func CreateFile(src string) error {
	err := os.MkdirAll(path.Dir(src), os.ModePerm)
	if err != nil {
		return err
	}
	//创建文件
	f, err := os.Create(src)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

// 打开文件并写入内容
func WriteFile(src string, content string) error {
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}

// ApplicationDir returns best base directory for specific OS.
func ApplicationDir(subDir ...string) string {
	cas := cases.Title(language.Dutch, cases.NoLower)
	for i := range subDir {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			cas.String(subDir[i])
		} else {
			subDir[i] = strings.ToLower(subDir[i])
		}
	}
	var appDir string
	home := os.Getenv("HOME")
	switch runtime.GOOS {
	case "windows":
		// Windows standards: https://msdn.microsoft.com/en-us/library/windows/apps/hh465094.aspx?f=255&MSPPError=-2147217396
		for _, env := range []string{"AppData", "AppDataLocal", "UserProfile", "Home"} {
			val := os.Getenv(env)
			if val != "" {
				appDir = val
				break
			}
		}
	case "darwin":
		// Mac standards: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html
		appDir = filepath.Join(home, "Library", "Application Support")
	case "linux":
		fallthrough
	default:
		// Linux standards: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
		appDir = os.Getenv("XDG_DATA_HOME")
		if appDir == "" && home != "" {
			appDir = filepath.Join(home, ".local", "share")
		}
	}
	return filepath.Join(append([]string{appDir}, subDir...)...)
}

func ApplicationDefaultFileName(filename string) string {
	return filepath.Join(ApplicationDir("zkgj_video"), filename)
}

func ApplicationAbsFileDir(fileName string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return ""
	}
	currentFile := filepath.Join(currentDir, fileName)
	return currentFile
}
