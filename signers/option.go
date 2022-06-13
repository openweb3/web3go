package signers

type MnemonicOption struct {
	DerivePath string `default:"m/44'/60'/0'/0"`
	Password   string `default:""`
}

func (m *MnemonicOption) WithDerivePath(derivePath string) *MnemonicOption {
	m.DerivePath = derivePath
	return m
}

func (m *MnemonicOption) WithPassword(password string) *MnemonicOption {
	m.Password = password
	return m
}
