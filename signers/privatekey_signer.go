package signers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
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
	if len(keyString) >= 2 && keyString[0:2] == "0x" {
		keyString = keyString[2:]
	}

	privateKey, err := crypto.HexToECDSA(keyString)

	if err != nil {
		return nil, errors.Wrap(err, "invalid HEX format of private key")
	}

	return NewPrivateKeySigner(privateKey), nil
}

func MustNewPrivateKeySignerByString(keyString string) *PrivateKeySigner {
	signer, err := NewPrivateKeySignerByString(keyString)
	if err != nil {
		panic(err)
	}
	return signer
}

func NewRandomPrivateKeySigner() (*PrivateKeySigner, error) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return NewPrivateKeySigner(privateKey), nil
}

func MustNewRandomPrivateKeySigner() *PrivateKeySigner {
	signer, err := NewRandomPrivateKeySigner()
	if err != nil {
		panic(err)
	}
	return signer
}

func NewPrivateKeySignerByMnemonic(mnemonic string, addressIndex int, option *MnemonicOption) (*PrivateKeySigner, error) {

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, ErrInvalidMnemonic
	}

	seed := bip39.NewSeed(mnemonic, option.Password)

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	_path, err := hdwallet.ParseDerivationPath(fmt.Sprintf("%v/%v", option.BaseDerivePath, addressIndex))
	if err != nil {
		return nil, err
	}

	account, err := wallet.Derive(_path, false)
	if err != nil {
		log.Fatal(err)
	}

	key, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}

	return &PrivateKeySigner{
		privateKey: key,
		address:    account.Address,
	}, nil
}

func MustNewPrivateKeySignerByMnemonic(mnemonic string, addressIndex int, option *MnemonicOption) *PrivateKeySigner {
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
