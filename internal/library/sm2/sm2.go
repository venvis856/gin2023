package sm2

import (
	"github.com/tjfoc/gmsm/sm2"
)

// 国防加密算法sm2

//sm2 struct
type SM2Struct struct {
	PrivateKey *sm2.PrivateKey
	PublicKey  *sm2.PublicKey
}

//初始化sm2实例
func NewSm2(PublicKeyPem string, PrivateKeyPem string) (*SM2Struct, error) {
	//获取公钥实例
	public_key, err := sm2.ReadPublicKeyFromPem(PublicKeyPem, nil)
	if err != nil {
		return nil, err
	}
	//获取私钥实例
	private_key, err := sm2.ReadPrivateKeyFromPem(PrivateKeyPem, nil)
	if err != nil {
		return nil, err
	}
	//返回SM2Struct实例
	c := new(SM2Struct)
	c.PublicKey = public_key
	c.PrivateKey = private_key
	return c, nil
}

//sm2解密操作
func (c *SM2Struct) SM2Decrypt(retb []byte) ([]byte, error) {
	ret, err := c.PrivateKey.Decrypt(retb)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// sm2加密操作
func (c *SM2Struct) SM2Encrypt(retb []byte) ([]byte, error) {
	ret, err := c.PublicKey.Encrypt(retb)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
