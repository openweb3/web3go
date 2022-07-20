package signers

import (
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
)

func NewPrivateKeySignerByKeystoreFile(filePath string, auth string) (*PrivateKeySigner, error) {
	key, err := privatekeyhelper.NewFromKeystoreFile(filePath, auth)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key), nil
}

func NewPrivateKeySignerByKeystore(keyjson []byte, auth string) (*PrivateKeySigner, error) {
	key, err := privatekeyhelper.NewFromKeystore(keyjson, auth)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key), nil
}

func MustNewPrivateKeySignerByKeystoreFile(filePath string, auth string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByKeystoreFile(filePath, auth)
	if err != nil {
		panic(err)
	}
	return signer
}

func MustNewPrivateKeySignerByKeystore(keyjson []byte, auth string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByKeystore(keyjson, auth)
	if err != nil {
		panic(err)
	}
	return signer
}

func (p *PrivateKeySigner) ToKeystore(auth string) ([]byte, error) {
	return privatekeyhelper.ToKeystore(p.privateKey, auth)
}

func (p *PrivateKeySigner) SaveKeystore(dirPath string, auth string) error {
	return privatekeyhelper.SaveAsKeystore(p.privateKey, dirPath, auth)
}
