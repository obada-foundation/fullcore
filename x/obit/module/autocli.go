package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	obitv1 "github.com/obada-foundation/fullcore/api/obadafoundation/obit/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:           obitv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           obitv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
