// Code generated by MockGen. DO NOT EDIT.
// Source: ../ethclient_gen.go
//
// Generated by this command:
//
//	mockgen -source ../ethclient_gen.go -package mock -destination ./mock_interfaces.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/omni-network/omni/lib/ethclient"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockClient) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address.
func (mr *MockClientMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockClient)(nil).Address))
}

// BalanceAt mocks base method.
func (m *MockClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceAt", ctx, account, blockNumber)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceAt indicates an expected call of BalanceAt.
func (mr *MockClientMockRecorder) BalanceAt(ctx, account, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceAt", reflect.TypeOf((*MockClient)(nil).BalanceAt), ctx, account, blockNumber)
}

// BlockByHash mocks base method.
func (m *MockClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockClientMockRecorder) BlockByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockClient)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockClientMockRecorder) BlockByNumber(ctx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockClient)(nil).BlockByNumber), ctx, number)
}

// BlockNumber mocks base method.
func (m *MockClient) BlockNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockClientMockRecorder) BlockNumber(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockClient)(nil).BlockNumber), ctx)
}

// CallContract mocks base method.
func (m *MockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, call, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockClientMockRecorder) CallContract(ctx, call, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockClient)(nil).CallContract), ctx, call, blockNumber)
}

// ChainID mocks base method.
func (m *MockClient) ChainID(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockClientMockRecorder) ChainID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockClient)(nil).ChainID), ctx)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// CodeAt mocks base method.
func (m *MockClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, account, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockClientMockRecorder) CodeAt(ctx, account, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockClient)(nil).CodeAt), ctx, account, blockNumber)
}

// EstimateGas mocks base method.
func (m *MockClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, call)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas.
func (mr *MockClientMockRecorder) EstimateGas(ctx, call any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockClient)(nil).EstimateGas), ctx, call)
}

// FilterLogs mocks base method.
func (m *MockClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", ctx, q)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs.
func (mr *MockClientMockRecorder) FilterLogs(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockClient)(nil).FilterLogs), ctx, q)
}

// HeaderByHash mocks base method.
func (m *MockClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHash indicates an expected call of HeaderByHash.
func (mr *MockClientMockRecorder) HeaderByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHash", reflect.TypeOf((*MockClient)(nil).HeaderByHash), ctx, hash)
}

// HeaderByNumber mocks base method.
func (m *MockClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", ctx, number)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockClientMockRecorder) HeaderByNumber(ctx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockClient)(nil).HeaderByNumber), ctx, number)
}

// HeaderByType mocks base method.
func (m *MockClient) HeaderByType(ctx context.Context, typ ethclient.HeadType) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByType", ctx, typ)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByType indicates an expected call of HeaderByType.
func (mr *MockClientMockRecorder) HeaderByType(ctx, typ any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByType", reflect.TypeOf((*MockClient)(nil).HeaderByType), ctx, typ)
}

// NonceAt mocks base method.
func (m *MockClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NonceAt", ctx, account, blockNumber)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NonceAt indicates an expected call of NonceAt.
func (mr *MockClientMockRecorder) NonceAt(ctx, account, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NonceAt", reflect.TypeOf((*MockClient)(nil).NonceAt), ctx, account, blockNumber)
}

// PeerCount mocks base method.
func (m *MockClient) PeerCount(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerCount", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerCount indicates an expected call of PeerCount.
func (mr *MockClientMockRecorder) PeerCount(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerCount", reflect.TypeOf((*MockClient)(nil).PeerCount), ctx)
}

// PendingBalanceAt mocks base method.
func (m *MockClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingBalanceAt", ctx, account)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingBalanceAt indicates an expected call of PendingBalanceAt.
func (mr *MockClientMockRecorder) PendingBalanceAt(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingBalanceAt", reflect.TypeOf((*MockClient)(nil).PendingBalanceAt), ctx, account)
}

// PendingCodeAt mocks base method.
func (m *MockClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCodeAt", ctx, account)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCodeAt indicates an expected call of PendingCodeAt.
func (mr *MockClientMockRecorder) PendingCodeAt(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCodeAt", reflect.TypeOf((*MockClient)(nil).PendingCodeAt), ctx, account)
}

// PendingNonceAt mocks base method.
func (m *MockClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingNonceAt", ctx, account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingNonceAt indicates an expected call of PendingNonceAt.
func (mr *MockClientMockRecorder) PendingNonceAt(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingNonceAt", reflect.TypeOf((*MockClient)(nil).PendingNonceAt), ctx, account)
}

// PendingStorageAt mocks base method.
func (m *MockClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingStorageAt", ctx, account, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingStorageAt indicates an expected call of PendingStorageAt.
func (mr *MockClientMockRecorder) PendingStorageAt(ctx, account, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingStorageAt", reflect.TypeOf((*MockClient)(nil).PendingStorageAt), ctx, account, key)
}

// PendingTransactionCount mocks base method.
func (m *MockClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingTransactionCount", ctx)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingTransactionCount indicates an expected call of PendingTransactionCount.
func (mr *MockClientMockRecorder) PendingTransactionCount(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingTransactionCount", reflect.TypeOf((*MockClient)(nil).PendingTransactionCount), ctx)
}

// SendTransaction mocks base method.
func (m *MockClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockClientMockRecorder) SendTransaction(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockClient)(nil).SendTransaction), ctx, tx)
}

// StorageAt mocks base method.
func (m *MockClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAt", ctx, account, key, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAt indicates an expected call of StorageAt.
func (mr *MockClientMockRecorder) StorageAt(ctx, account, key, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAt", reflect.TypeOf((*MockClient)(nil).StorageAt), ctx, account, key, blockNumber)
}

// SubscribeFilterLogs mocks base method.
func (m *MockClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeFilterLogs", ctx, q, ch)
	ret0, _ := ret[0].(ethereum.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeFilterLogs indicates an expected call of SubscribeFilterLogs.
func (mr *MockClientMockRecorder) SubscribeFilterLogs(ctx, q, ch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeFilterLogs", reflect.TypeOf((*MockClient)(nil).SubscribeFilterLogs), ctx, q, ch)
}

// SubscribeNewHead mocks base method.
func (m *MockClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewHead", ctx, ch)
	ret0, _ := ret[0].(ethereum.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewHead indicates an expected call of SubscribeNewHead.
func (mr *MockClientMockRecorder) SubscribeNewHead(ctx, ch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewHead", reflect.TypeOf((*MockClient)(nil).SubscribeNewHead), ctx, ch)
}

// SuggestGasPrice mocks base method.
func (m *MockClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice.
func (mr *MockClientMockRecorder) SuggestGasPrice(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockClient)(nil).SuggestGasPrice), ctx)
}

// SuggestGasTipCap mocks base method.
func (m *MockClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasTipCap", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasTipCap indicates an expected call of SuggestGasTipCap.
func (mr *MockClientMockRecorder) SuggestGasTipCap(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasTipCap", reflect.TypeOf((*MockClient)(nil).SuggestGasTipCap), ctx)
}

// SyncProgress mocks base method.
func (m *MockClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProgress", ctx)
	ret0, _ := ret[0].(*ethereum.SyncProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncProgress indicates an expected call of SyncProgress.
func (mr *MockClientMockRecorder) SyncProgress(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProgress", reflect.TypeOf((*MockClient)(nil).SyncProgress), ctx)
}

// TransactionByHash mocks base method.
func (m *MockClient) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, txHash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockClientMockRecorder) TransactionByHash(ctx, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockClient)(nil).TransactionByHash), ctx, txHash)
}

// TransactionCount mocks base method.
func (m *MockClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionCount", ctx, blockHash)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionCount indicates an expected call of TransactionCount.
func (mr *MockClientMockRecorder) TransactionCount(ctx, blockHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCount", reflect.TypeOf((*MockClient)(nil).TransactionCount), ctx, blockHash)
}

// TransactionInBlock mocks base method.
func (m *MockClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionInBlock", ctx, blockHash, index)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionInBlock indicates an expected call of TransactionInBlock.
func (mr *MockClientMockRecorder) TransactionInBlock(ctx, blockHash, index any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionInBlock", reflect.TypeOf((*MockClient)(nil).TransactionInBlock), ctx, blockHash, index)
}

// TransactionReceipt mocks base method.
func (m *MockClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockClientMockRecorder) TransactionReceipt(ctx, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockClient)(nil).TransactionReceipt), ctx, txHash)
}
