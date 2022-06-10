package signers

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewSingerManagerByPrivateKeys(t *testing.T) {
	sm, err := NewSignerManagerByPrivateKeyStrings([]string{
		"9ec393923a14eeb557600010ea05d635c667a6995418f8a8f4bdecc63dfe0bb9",
		"1ab8ec2627e19007d2c62145df6acf51f16b8fd93b0a27c01dae4eb271aadee1",
	})
	if err != nil {
		t.Fatal(err)
	}
	a := assert.New(t)
	list := sm.List()
	a.Equal(2, len(list))
	a.Equal(list[0].Address(), common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"))
	a.Equal(list[1].Address(), common.HexToAddress("0x3a3347C42705C5328012dE9a38b030128eee4F83"))

	signer, err := sm.Get(common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"))
	a.NoError(err)
	a.Equal(signer.Address(), common.HexToAddress("0xe6D148D8398c4cb456196C776D2d9093Dd62C9B0"))
}
