package helper

import "strings"

// CheckPicFormat 检测文件是否是图片格式 filename文件名
func CheckPicFormat(filename string) bool {
	format := []string{"webp", "bmp", "pcx", "tif", "gif", "jpeg", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps", "ai", "png", "hdri", "paw", "wmf", "flic", "emf", "ico", "avif"}
	fileExt := filename[strings.Index(filename, ".")+1:]
	bol := InArray(strings.ToUpper(fileExt), format)
	return bol
}
