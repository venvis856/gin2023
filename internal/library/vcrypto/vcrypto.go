package vcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"math"
	"strings"
)

// 加密类

// HexEnCrypt 字符串加密结果转16进制字符（注：后一个参数为加密函数）
func HexEnCrypt(str, key string, cryptCallback func(o, k []byte) ([]byte, error)) string {
	b, err := cryptCallback([]byte(str), []byte(key))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// HexDeCrypt 16进制字符转换并解密出字符串（注：后一个参数为解密函数）
func HexDeCrypt(str, key string, cryptCallback func(o, k []byte) ([]byte, error)) string {
	origData, err := hex.DecodeString(str)
	if err != nil {
		return ""
	}
	b, err := cryptCallback(origData, []byte(key))
	if err != nil {
		return ""
	}
	return string(b)
}

// DES-EDE3-CBC 解密 8位key
func DesCBCDecrypt(origData, key []byte) ([]byte, error) {

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	cipherText := make([]byte, len(origData))
	blockMode.CryptBlocks(cipherText, origData)
	cipherText = ZeroUnPadding(Pkcs5Unpadding(cipherText))

	return cipherText, nil
}

// DES-EDE3-CBC加密 8位key
func DesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	origData = Pkcs5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	cipherText := make([]byte, len(origData))
	blockMode.CryptBlocks(cipherText, origData)

	return cipherText, nil
}

// AesECBEncrypt AES-128-ECB加密  16位key
// AES-192-ECB加密  24位key
// AES-256-ECB加密  32位key
func AesECBEncrypt(origData, key []byte) ([]byte, error) {

	block, _ := aes.NewCipher(key)
	origData = PKCS7Padding(origData, block.BlockSize())
	decrypted := make([]byte, len(origData))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(origData); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], origData[bs:be])
	}

	return decrypted, nil
}

// AesECBDecrypt AES-128-ECB解密  16位key
// AES-192-ECB解密  24位key
// AES-256-ECB解密  32位key
func AesECBDecrypt(origData, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(origData))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(origData); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], origData[bs:be])
	}
	return PKCS7UnPadding(decrypted), nil
}

// Pkcs5Padding Pkcs5Padding
func Pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// Pkcs5Unpadding 去掉字符串后面的填充字符
func Pkcs5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PKCS7Padding PKCS7Padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding PKCS7UnPadding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// ZeroPadding ZeroPadding
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

// ZeroUnPadding ZeroUnPadding
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// StrPad 字符串长度不足，用指定的字符前后补全，多余的则删除 returns the input string padded on the left, right or both sides using padType to
// the specified padding length padLength.
//
// Example:
// input := "Codes";
// StrPad(input, 10, " ", "RIGHT")        // produces "Codes     "
// StrPad(input, 10, "-=", "LEFT")        // produces "=-=-=Codes"
// StrPad(input, 10, "_", "BOTH")         // produces "__Codes___"
// StrPad(input, 6, "___", "RIGHT")       // produces "Codes_"
// StrPad(input, 3, "*", "RIGHT")         // produces "Codes"
func StrPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input[0:padLength]
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input +
			strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}
