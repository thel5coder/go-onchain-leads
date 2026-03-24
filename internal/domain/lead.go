package domain

type Lead struct {
	DeveloperWallet string
	ENSName         string
	Twitter         string
	Email           string
	GitHub          string
	ContractAddress string
	TransactionHash string
	GasUsed         uint64
}

type Identity struct {
	Name    string
	Twitter string
	Email   string
	GitHub  string
}
