package constant

import "github.com/meshplus/bitxhub-kit/types"

type BoltContractAddress string

const (
	InterchainContractAddr       BoltContractAddress = "0x000000000000000000000000000000000000000a"
	StoreContractAddr            BoltContractAddress = "0x000000000000000000000000000000000000000b"
	RuleManagerContractAddr      BoltContractAddress = "0x000000000000000000000000000000000000000c"
	RoleContractAddr             BoltContractAddress = "0x000000000000000000000000000000000000000d"
	AppchainMgrContractAddr      BoltContractAddress = "0x000000000000000000000000000000000000000e"
	TransactionMgrContractAddr   BoltContractAddress = "0x000000000000000000000000000000000000000f"
	AssetExchangeContractAddr    BoltContractAddress = "0x0000000000000000000000000000000000000010"
	ServiceMgrContractAddr       BoltContractAddress = "0x0000000000000000000000000000000000000011"
	DIDRegistryContractAddr      BoltContractAddress = "0x0000000000000000000000000000000000000012"
	MethodRegistryContractAddr   BoltContractAddress = "0x0000000000000000000000000000000000000013"
	InterRelayBrokerContractAddr BoltContractAddress = "0x0000000000000000000000000000000000000014"
	GovernanceContractAddr       BoltContractAddress = "0x0000000000000000000000000000000000000015"
	VCRegistryContractAddr       BoltContractAddress = "0x0000000000000000000000000000000000000016"
	EthHeaderMgrContractAddr     BoltContractAddress = "0x0000000000000000000000000000000000000017"
	NodeManagerContractAddr      BoltContractAddress = "0x0000000000000000000000000000000000000018"
	InterBrokerContractAddr      BoltContractAddress = "0x0000000000000000000000000000000000000019"
)

func (addr BoltContractAddress) Address() *types.Address {
	return types.NewAddressByStr(string(addr))
}

func (addr BoltContractAddress) String() string {
	return string(addr)
}
