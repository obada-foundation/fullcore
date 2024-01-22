// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	nft "cosmossdk.io/x/nft"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockNftKeeper is a mock of NftKeeper interface.
type MockNftKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockNftKeeperMockRecorder
}

// MockNftKeeperMockRecorder is the mock recorder for MockNftKeeper.
type MockNftKeeperMockRecorder struct {
	mock *MockNftKeeper
}

// NewMockNftKeeper creates a new mock instance.
func NewMockNftKeeper(ctrl *gomock.Controller) *MockNftKeeper {
	mock := &MockNftKeeper{ctrl: ctrl}
	mock.recorder = &MockNftKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNftKeeper) EXPECT() *MockNftKeeperMockRecorder {
	return m.recorder
}

// BatchMint mocks base method.
func (m *MockNftKeeper) BatchMint(ctx types.Context, tokens []nft.NFT, receiver types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchMint", ctx, tokens, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchMint indicates an expected call of BatchMint.
func (mr *MockNftKeeperMockRecorder) BatchMint(ctx, tokens, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchMint", reflect.TypeOf((*MockNftKeeper)(nil).BatchMint), ctx, tokens, receiver)
}

// GetClasses mocks base method.
func (m *MockNftKeeper) GetClasses(ctx types.Context) []*nft.Class {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClasses", ctx)
	ret0, _ := ret[0].([]*nft.Class)
	return ret0
}

// GetClasses indicates an expected call of GetClasses.
func (mr *MockNftKeeperMockRecorder) GetClasses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClasses", reflect.TypeOf((*MockNftKeeper)(nil).GetClasses), ctx)
}

// GetNFT mocks base method.
func (m *MockNftKeeper) GetNFT(ctx types.Context, classID, nftID string) (nft.NFT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFT", ctx, classID, nftID)
	ret0, _ := ret[0].(nft.NFT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNFT indicates an expected call of GetNFT.
func (mr *MockNftKeeperMockRecorder) GetNFT(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFT", reflect.TypeOf((*MockNftKeeper)(nil).GetNFT), ctx, classID, nftID)
}

// GetNFTsOfClassByOwner mocks base method.
func (m *MockNftKeeper) GetNFTsOfClassByOwner(ctx types.Context, classID string, owner types.AccAddress) []nft.NFT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFTsOfClassByOwner", ctx, classID, owner)
	ret0, _ := ret[0].([]nft.NFT)
	return ret0
}

// GetNFTsOfClassByOwner indicates an expected call of GetNFTsOfClassByOwner.
func (mr *MockNftKeeperMockRecorder) GetNFTsOfClassByOwner(ctx, classID, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFTsOfClassByOwner", reflect.TypeOf((*MockNftKeeper)(nil).GetNFTsOfClassByOwner), ctx, classID, owner)
}

// GetOwner mocks base method.
func (m *MockNftKeeper) GetOwner(ctx types.Context, classID, nftID string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwner", ctx, classID, nftID)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetOwner indicates an expected call of GetOwner.
func (mr *MockNftKeeperMockRecorder) GetOwner(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwner", reflect.TypeOf((*MockNftKeeper)(nil).GetOwner), ctx, classID, nftID)
}

// HasClass mocks base method.
func (m *MockNftKeeper) HasClass(ctx types.Context, classID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasClass", ctx, classID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasClass indicates an expected call of HasClass.
func (mr *MockNftKeeperMockRecorder) HasClass(ctx, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasClass", reflect.TypeOf((*MockNftKeeper)(nil).HasClass), ctx, classID)
}

// Mint mocks base method.
func (m *MockNftKeeper) Mint(ctx types.Context, token nft.NFT, receiver types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mint", ctx, token, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mint indicates an expected call of Mint.
func (mr *MockNftKeeperMockRecorder) Mint(ctx, token, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mint", reflect.TypeOf((*MockNftKeeper)(nil).Mint), ctx, token, receiver)
}

// SaveClass mocks base method.
func (m *MockNftKeeper) SaveClass(ctx types.Context, class nft.Class) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveClass", ctx, class)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveClass indicates an expected call of SaveClass.
func (mr *MockNftKeeperMockRecorder) SaveClass(ctx, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveClass", reflect.TypeOf((*MockNftKeeper)(nil).SaveClass), ctx, class)
}

// Transfer mocks base method.
func (m *MockNftKeeper) Transfer(ctx types.Context, classID, nftID string, receiver types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, classID, nftID, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockNftKeeperMockRecorder) Transfer(ctx, classID, nftID, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockNftKeeper)(nil).Transfer), ctx, classID, nftID, receiver)
}

// Update mocks base method.
func (m *MockNftKeeper) Update(ctx types.Context, token nft.NFT) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNftKeeperMockRecorder) Update(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNftKeeper)(nil).Update), ctx, token)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}
