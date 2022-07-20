package signers

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/pkg/errors"
)

var (
	ErrInvalidMnemonic error = errors.New("invalid mnemonic")
)

type PrivateKeySigner struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	pubKey     *ecdsa.PublicKey
}

func NewPrivateKeySigner(privateKey *ecdsa.PrivateKey) *PrivateKeySigner {
	p := &PrivateKeySigner{
		privateKey: privateKey,
		pubKey:     &privateKey.PublicKey,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}

	return p
}

func NewPrivateKeySignerByString(keyString string) (*PrivateKeySigner, error) {
	key, err := privatekeyhelper.NewFromKeyString(keyString)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key), nil
}

func NewRandomPrivateKeySigner() (*PrivateKeySigner, error) {
	key, err := privatekeyhelper.NewRandom()
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key), nil
}

func NewPrivateKeySignerByMnemonic(mnemonic string, addressIndex int, option *privatekeyhelper.MnemonicOption) (*PrivateKeySigner, error) {
	key, err := privatekeyhelper.NewFromMnemonic(mnemonic, addressIndex, option)
	if err != nil {
		return nil, err
	}
	return NewPrivateKeySigner(key), nil
}

func MustNewPrivateKeySignerByString(keyString string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByString(keyString)
	if err != nil {
		panic(err)
	}
	return signer
}

func MustNewRandomPrivateKeySigner() *PrivateKeySigner {
	signer, err := NewRandomPrivateKeySigner()
	if err != nil {
		panic(err)
	}
	return signer
}

func MustNewPrivateKeySignerByMnemonic(mnemonic string, addressIndex int, option *privatekeyhelper.MnemonicOption) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByMnemonic(mnemonic, addressIndex, option)
	if err != nil {
		panic(err)
	}
	return signer
}

func (p *PrivateKeySigner) Address() common.Address {
	return p.address
}

func (p *PrivateKeySigner) PrivateKey() *ecdsa.PrivateKey {
	return p.privateKey
}

func (p *PrivateKeySigner) PrivateKeyString() string {
	privKeyBytes := (crypto.FromECDSA(p.privateKey))
	privKeyStr := hexutil.Encode(privKeyBytes)
	return privKeyStr
}

func (p *PrivateKeySigner) PublicKey() *ecdsa.PublicKey {
	return p.pubKey
}

func (p *PrivateKeySigner) PublicKeyString() string {
	pubKeyBytes := crypto.FromECDSAPub(p.pubKey)
	pubKeyStr := hexutil.Encode(pubKeyBytes[1:])
	return pubKeyStr
}

func (p *PrivateKeySigner) SignTransaction(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(chainID)
	return types.SignTx(tx, signer, p.privateKey)
}

func (p *PrivateKeySigner) SignMessage(text []byte) ([]byte, error) {
	hash := accounts.TextHash(text)
	return crypto.Sign(hash, p.privateKey)
}

func (p PrivateKeySigner) String() string {
	return fmt.Sprintf("address: %v, publicKey: %v", p.address.Hex(), p.PublicKeyString())
}
