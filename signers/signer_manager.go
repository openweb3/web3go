package signers

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/interfaces"
)

type SignerManager struct {
	signerMap map[common.Address]interfaces.Signer
	signers   []interfaces.Signer
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
