package signers

import (
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/openweb3/web3go/interfaces"
)

type SignerManager struct {
	signerMap map[common.Address]interfaces.Signer
	signers   []interfaces.Signer
	mutex     sync.Mutex
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

func MustNewSignerManagerByPrivateKeyStrings(privateKeys []string) *SignerManager {
	sm, err := NewSignerManagerByPrivateKeyStrings(privateKeys)
	if err != nil {
		panic(err)
	}
	return sm
}

func NewSignerManagerByMnemonic(mnemonic string, addressNumber int, option *privatekeyhelper.MnemonicOption) (*SignerManager, error) {
	signers := make([]interfaces.Signer, addressNumber)
	for i := 0; i < addressNumber; i++ {
		s, err := NewPrivateKeySignerByMnemonic(mnemonic, i, option)
		if err != nil {
			return nil, err
		}
		signers[i] = s
	}
	return NewSignerManager(signers), nil
}

func MustNewSignerManagerByMnemonic(mnemonic string, addressNumber int, option *privatekeyhelper.MnemonicOption) *SignerManager {
	sm, err := NewSignerManagerByMnemonic(mnemonic, addressNumber, option)
	if err != nil {
		panic(err)
	}
	return sm
}

func (s *SignerManager) Add(signer interfaces.Signer) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.signerMap[signer.Address()]; ok {
		return errors.New("signer already exists")
	}
	s.signers = append(s.signers, signer)
	s.signerMap[signer.Address()] = signer
	return nil
}

func (s *SignerManager) Remove(addr common.Address) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	if _, ok := s.signerMap[addr]; !ok {
		return nil, errors.New("signer not found")
	}
	return s.signerMap[addr], nil
}

func (s *SignerManager) List() []interfaces.Signer {
	return s.signers
}
