package signers

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mcuadros/go-defaults"
	"github.com/openweb3/web3go/interfaces"
)

type SignerManager struct {
	signerMap map[common.Address]interfaces.Signer
	signers   []interfaces.Signer
}

func NewSignerManager(signers []interfaces.Signer) *SignerManager {

	sm := &SignerManager{
		signerMap: make(map[common.Address]interfaces.Signer),
		signers:   signers,
	}

	for _, signer := range signers {
		sm.signerMap[signer.Address()] = signer
	}

	return sm
}

func NewSignerManagerByPrivateKeyStrings(privateKeys []string) (*SignerManager, error) {
	signers := make([]interfaces.Signer, len(privateKeys))
	for i, p := range privateKeys {
		s, err := NewPrivateKeySignerByString(p)
		if err != nil {
			return nil, err
		}
		signers[i] = s
	}
	return NewSignerManager(signers), nil
}

func NewSignerManagerByMnemonic(mnemonic string, option *MnemonicOption) (*SignerManager, error) {
	defaults.SetDefaults(option)
	signers := make([]interfaces.Signer, option.Number)
	for i := 0; i < option.Number; i++ {
		s, err := NewPrivateKeySignerByMnemonic(mnemonic, fmt.Sprintf("%s/%v", option.DerivePath, i), option.Password)
		if err != nil {
			return nil, err
		}
		signers[i] = s
	}
	return NewSignerManager(signers), nil
}

func (s *SignerManager) Add(signer interfaces.Signer) {
	s.signers = append(s.signers, signer)
	s.signerMap[signer.Address()] = signer
}

func (s *SignerManager) Remove(addr common.Address) error {
	if _, err := s.Get(addr); err != nil {
		return err
	}

	delete(s.signerMap, addr)
	for i, signer := range s.signers {
		if signer.Address() == addr {
			s.signers = append(s.signers[:i], s.signers[i+1:]...)
			break
		}
	}
	return nil
}

func (s *SignerManager) Get(addr common.Address) (interfaces.Signer, error) {
	if s.signerMap[addr] == nil {
		return nil, errors.New("signer not found")
	}
	return s.signerMap[addr], nil
}

func (s *SignerManager) List() []interfaces.Signer {
	return s.signers
}
