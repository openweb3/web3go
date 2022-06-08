package signers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
)

type PrivateSigner struct {
	address    common.Address
	privateKey *ecdsa.PrivateKey
	pubKey     *ecdsa.PublicKey
}

func NewPrivateSigner(keyString string) (*PrivateSigner, error) {
	privateKey, err := crypto.HexToECDSA(keyString)

	if err != nil {
		return nil, errors.Wrap(err, "invalid HEX format of private key")
	}

	p := &PrivateSigner{
		privateKey: privateKey,
		pubKey:     &privateKey.PublicKey,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}

	return p, nil
}

func NewPrivateSignerByMnemonic(mnemonic string, password string, path string) (*PrivateSigner, error) {
	seed := bip39.NewSeed(mnemonic, password)

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	_path, err := hdwallet.ParseDerivationPath(path)
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

	return &PrivateSigner{
		privateKey: key,
		address:    account.Address,
	}, nil
}

func NewRandomPrivateSigner() (*PrivateSigner, error) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &PrivateSigner{
		privateKey: privateKey,
		pubKey:     &privateKey.PublicKey,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

func (p *PrivateSigner) Address() common.Address {
	return p.address
}

func (p *PrivateSigner) SignTransaction(tx *types.Transaction) (*types.Transaction, error) {
	chainID := tx.ChainId()
	signer := types.LatestSignerForChainID(chainID)
	return types.SignTx(tx, signer, p.privateKey)
}

func (p *PrivateSigner) SignMessage(text []byte) ([]byte, error) {
	hash := crypto.Keccak256(text)
	return crypto.Sign(hash, p.privateKey)
}
