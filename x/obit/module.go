package obit

import (
	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.HasGenesis     = AppModule{}

	_ appmodule.AppModule   = AppModule{}
	_ appmodule.HasServices = AppModule{}
)

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// NewAppModuleBasic creates a new AppModuleBasic
func NewAppModuleBasic(cdc codec.Codec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterCodec registers the capability module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// AppModule implements an application module for the obit module.
type AppModule struct {
	keeper *keeper.Keeper
}
