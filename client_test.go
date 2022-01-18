package web3go

import (
	"testing"
)

func TestClient(t *testing.T) {
	client, err := NewClient("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Eth.ClientVersion()
	if err != nil {
		t.Fatal(err)
	}
}
