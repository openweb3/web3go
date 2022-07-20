package signers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/stretchr/testify/assert"
)

func TestNewPrivateKeySignerByString(t *testing.T) {
	a := assert.New(t)
	s, err := NewPrivateKeySignerByString("9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9")
	a.NoError(err)
	a.Equal(common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"), s.Address())

	s, err = NewPrivateKeySignerByString("0x9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9")
	a.NoError(err)
	a.Equal(common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"), s.Address())

	_, err = NewPrivateKeySignerByString("0x123")
	a.Error(err, "invalid HEX format of private key")

	_, err = NewPrivateKeySignerByString("")
	a.Error(err, "invalid HEX format of private key")
}

func TestNewPrivateSignerByMnemonic(t *testing.T) {
	a := assert.New(t)
	_, err := NewPrivateKeySignerByMnemonic("", 0, &privatekeyhelper.MnemonicOption{
		BaseDerivePath: "m/44'/60'/0'/0",
	})
	a.EqualError(err, ErrInvalidMnemonic.Error())

	m := "sister horse tag together deposit lazy wide trust stay vital six napkin"
	s, _ := NewPrivateKeySignerByMnemonic(m, 0, &privatekeyhelper.MnemonicOption{BaseDerivePath: "m/44'/60'/0'/0"})
	a.Equal(common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"), s.Address())

	s, _ = NewPrivateKeySignerByMnemonic(m, 1, &privatekeyhelper.MnemonicOption{BaseDerivePath: "m/44'/60'/0'/0"})
	a.Equal(common.HexToAddress("0x3a3347C42705C5328012dE9a38b030128eee4F83"), s.Address())
}

func TestMarshalText(t *testing.T) {
	a := assert.New(t)
	s, err := NewPrivateKeySignerByString("9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9")
	a.NoError(err)
	a.Equal(s.String(), "address: 0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0, publicKey: 0x08b68fb0ff6e47cf278f0bb0879b00b6bb00b69cea9e5c24348bf7bed07f2cc347e3aa1f73f0ae2c9a6ae9213e18938db88fcaef5f3c6c1f8b3cfdd985a19cad")
}

func TestSignLegacyTransaction(t *testing.T) {
	a := assert.New(t)
	signer, err := NewPrivateKeySignerByString("9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9")
	a.NoError(err)

	tx := types.NewTx(&types.LegacyTx{})
	tx, err = signer.SignTransaction(tx, nil)
	a.NoError(err)

	fmt.Printf("tx %s\n", MustJsonMarshalTx(tx))

	v, r, s := tx.RawSignatureValues()
	expectR, _ := big.NewInt(0).SetString("0xbd6870726e062f73f80bd6529fac257f391009c0d95134ac719876ebbb504631", 0)
	expectS, _ := big.NewInt(0).SetString("0x4b81c0ca9202115889ac6124e162f3ff304621d26ef156ccf54af861d5f29163", 0)
	a.Equal(v, big.NewInt(28))
	a.Equal(r, expectR)
	a.Equal(s, expectS)
}

func TestSignDynamicFeeTx(t *testing.T) {
	a := assert.New(t)
	signer, err := NewPrivateKeySignerByString("9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9")
	a.NoError(err)

	from := signer.Address()

	tx := types.NewTx(&types.DynamicFeeTx{
		To:        &from,
		Value:     big.NewInt(1),
		GasFeeCap: big.NewInt(2),
		GasTipCap: big.NewInt(1),
		Nonce:     0,
		Gas:       21000,
		ChainID:   big.NewInt(3),
	})
	tx, err = signer.SignTransaction(tx, big.NewInt(3))
	a.NoError(err)

	fmt.Printf("tx %s\n", MustJsonMarshalTx(tx))

	marshaled, err := tx.MarshalBinary()
	a.NoError(err)
	expectMarshaled := "02f8620380010282520894e6d148d8398c4cb456196c776d2d9093dd62c9b00180c001a055e7caac90c8e0135575c851ae8e18693017b42da586b78c8b11d221b0192b77a0663d198301e510c88158375d9fcb953fb49b586fa0c48bdee207f95f038b5281"
	a.Equal(fmt.Sprintf("%x", marshaled), expectMarshaled)
}

func MustJsonMarshalTx(tx *types.Transaction) []byte {
	j, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return j
}
