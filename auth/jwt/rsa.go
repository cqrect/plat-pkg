package jwt

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/smallstep/cli/crypto/pemutil"
)

var (
	ErrInvalidPrivateKeyFile = errors.New("invalid private key file")
)

// LoadRSAPrivateKeyFromPEM 从PEM私钥文件 keyFile 与密码文件 passFile 中加载 RSA 私钥
func LoadRSAPrivateKeyFromPEM(keyFile, passFile string) (*rsa.PrivateKey, error) {
	key, err := pemutil.Read(keyFile, pemutil.WithPasswordFile(passFile))
	if err != nil {
		return nil, err
	}
	if sk, ok := key.(*rsa.PrivateKey); ok {
		return sk, nil
	}

	return nil, ErrInvalidPrivateKeyFile
}

// LoadRSAPrivateKey 从私钥的字节序列中加载 RSA 私钥
func LoadRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM(key)
}

// LoadRSAPublicKeyFromPEM 从PEM公钥文件 file 中加载 RSA 公钥
func LoadRSAPublicKeyFromPEM(file string) (*rsa.PublicKey, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadRSAPublicKey(f)
}

// LoadRSAPublicKey 从字节序列中加载 RSA 公钥
func LoadRSAPublicKey(key []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(key)
}
