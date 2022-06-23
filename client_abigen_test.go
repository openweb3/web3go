package web3go

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestClientForContractImplementInterfaces(t *testing.T) {
	var _ bind.ContractBackend = &ClientForContract{}
	var _ bind.DeployBackend = &ClientForContract{}
}
