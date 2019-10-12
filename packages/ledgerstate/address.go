package ledgerstate

type Address struct {
	hash                   AddressHash
	unspentTransferOutputs TransferOutputs
}

func NewAddress(hash AddressHash) *Address {
	return &Address{
		hash:                   hash,
		unspentTransferOutputs: make(TransferOutputs),
	}
}

func (address *Address) GetHash() AddressHash {
	return address.hash
}

func (address *Address) AddTransferOutput(transferHash TransferHash, coloredBalance *ColoredBalance) *Address {
	unspentTransferOutput, exists := address.unspentTransferOutputs[transferHash]
	if !exists {
		unspentTransferOutput = NewTransferOutput(transferHash)
		address.unspentTransferOutputs[transferHash] = unspentTransferOutput
	}

	unspentTransferOutput.SetColoredBalance(coloredBalance.color, coloredBalance.balance)

	return address
}

func (address *Address) GetUnspentTransferOutputs() TransferOutputs {
	return address.unspentTransferOutputs
}

func (address *Address) Exists() bool {
	return address != nil
}
