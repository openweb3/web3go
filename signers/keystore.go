package signers

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

func NewPrivateKeySignerByKeystoreFile(filePath string, auth string) (*PrivateKeySigner, error) {
	keyjson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySignerByKeystore(keyjson, auth)
}

func MustNewPrivateKeySignerByKeystoreFile(filePath string, auth string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByKeystoreFile(filePath, auth)
	if err != nil {
		panic(err)
	}
	return signer
}

func NewPrivateKeySignerByKeystore(keyjson []byte, auth string) (*PrivateKeySigner, error) {
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key.PrivateKey), nil
}

func MustNewPrivateKeySignerByKeystore(keyjson []byte, auth string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByKeystore(keyjson, auth)
	if err != nil {
		panic(err)
	}
	return signer
}

func (p *PrivateKeySigner) ToKeystore(auth string) ([]byte, error) {
	keyjson, _, err := p.toKeystore(auth)
	return keyjson, err
}

func (p *PrivateKeySigner) SaveKeystore(dirPath string, auth string) error {
	// load accounts from dir, return error if exists
	ks := keystore.NewKeyStore(dirPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if ks.HasAddress(p.Address()) {
		return keystore.ErrAccountAlreadyExists
	}

	keyjson, key, err := p.toKeystore(auth)
	if err != nil {
		return err
	}

	a := accounts.Account{
		Address: key.Address,
		URL:     accounts.URL{Scheme: keystore.KeyStoreScheme, Path: JoinPath(dirPath, keyFileName(key.Address))},
	}

	tmpName, err := writeTemporaryKeyFile(a.URL.Path, keyjson)
	if err != nil {
		return err
	}

	return os.Rename(tmpName, a.URL.Path)
}

func (p *PrivateKeySigner) toKeystore(auth string) ([]byte, *keystore.Key, error) {
	id, err := uuid.NewRandom()
	for err != nil {
		id, err = uuid.NewRandom()
	}

	key := &keystore.Key{
		Id:         id,
		Address:    p.Address(),
		PrivateKey: p.privateKey,
	}

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	keyjson, err := keystore.EncryptKey(key, auth, scryptN, scryptP)
	if err != nil {
		return nil, nil, err
	}
	return keyjson, key, nil
}

func writeTemporaryKeyFile(file string, content []byte) (string, error) {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return "", err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}
	f.Close()
	return f.Name(), nil
}

// keyFileName implements the naming convention for keyfiles:
// UTC--<created_at UTC ISO8601>-<address hex>
func keyFileName(keyAddr common.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}

func JoinPath(dirPath string, filename string) string {
	return filepath.Join(dirPath, filename)
}
