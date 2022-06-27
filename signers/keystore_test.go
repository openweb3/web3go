package signers

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewPrivateKeySignerByKeystore(t *testing.T) {
	keyjson := []byte(`{"address":"c41899f4588e58f76bbfb07ccca4e4fafccbe1ae","crypto":{"cipher":"aes-128-ctr","ciphertext":"99e29c806b220e98c74516ecfc590a1b46bb7e1a0b3538d53b72eed15434f5c1","cipherparams":{"iv":"63d7fc36514c80f3ae7d57a02133dc24"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"62c112c1c1f853a2ad413e103fc2acc9a2be261b24441893951a0d1db19b6267"},"mac":"27dd38452370d8413c68dde25ba0af17e788f7b2e62f4196146352fdf9f3c66f"},"id":"6cd87026-3b85-46d2-a229-e25164c75d21","version":3}`)
	signer, err := NewPrivateKeySignerByKeystore(keyjson, "foo")
	assert.NoError(t, err)

	assert.Equal(t, common.HexToAddress("0xc41899f4588e58f76bbfb07ccca4e4fafccbe1ae"), signer.Address())
}

func TestToKeystore(t *testing.T) {
	signer := MustNewPrivateKeySignerByString("0x0d501c86786789f8cb068c3ed04c0c010db2e526d864bedf2f5a68419daa8e90")
	keyjson, err := signer.ToKeystore("foo")
	assert.NoError(t, err)

	signer2, err := NewPrivateKeySignerByKeystore(keyjson, "foo")
	assert.NoError(t, err)

	assert.Equal(t, signer.Address(), signer2.Address())
}

func TestStoreToKeystore(t *testing.T) {
	err := os.RemoveAll("keys")
	assert.NoError(t, err)
	keyjson := []byte(`{"address":"c41899f4588e58f76bbfb07ccca4e4fafccbe1ae","crypto":{"cipher":"aes-128-ctr","ciphertext":"99e29c806b220e98c74516ecfc590a1b46bb7e1a0b3538d53b72eed15434f5c1","cipherparams":{"iv":"63d7fc36514c80f3ae7d57a02133dc24"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"62c112c1c1f853a2ad413e103fc2acc9a2be261b24441893951a0d1db19b6267"},"mac":"27dd38452370d8413c68dde25ba0af17e788f7b2e62f4196146352fdf9f3c66f"},"id":"6cd87026-3b85-46d2-a229-e25164c75d21","version":3}`)
	err = MustNewPrivateKeySignerByKeystore(keyjson, "foo").SaveKeystore("keys", "foo")
	assert.NoError(t, err)
	err = MustNewPrivateKeySignerByKeystore(keyjson, "foo").SaveKeystore("keys", "foo")
	assert.Error(t, err)

	err = MustNewRandomPrivateKeySigner().SaveKeystore("keys", "foo")
	assert.NoError(t, err)
}
