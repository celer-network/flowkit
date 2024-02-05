// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	cadence "github.com/onflow/cadence"
	accounts "github.com/onflow/flowkit/v2/accounts"

	config "github.com/onflow/flowkit/v2/config"

	context "context"

	crypto "github.com/onflow/crypto"

	flow "github.com/onflow/flow-go-sdk"

	flowkit "github.com/onflow/flowkit/v2"

	gateway "github.com/onflow/flowkit/v2/gateway"

	mock "github.com/stretchr/testify/mock"

	output "github.com/onflow/flowkit/v2/output"

	project "github.com/onflow/flowkit/v2/project"

	transactions "github.com/onflow/flowkit/v2/transactions"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AddContract provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Services) AddContract(_a0 context.Context, _a1 *accounts.Account, _a2 flowkit.Script, _a3 flowkit.UpdateContract) (flow.Identifier, bool, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for AddContract")
	}

	var r0 flow.Identifier
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, flowkit.Script, flowkit.UpdateContract) (flow.Identifier, bool, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, flowkit.Script, flowkit.UpdateContract) flow.Identifier); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *accounts.Account, flowkit.Script, flowkit.UpdateContract) bool); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *accounts.Account, flowkit.Script, flowkit.UpdateContract) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BuildTransaction provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Services) BuildTransaction(_a0 context.Context, _a1 transactions.AddressesRoles, _a2 int, _a3 flowkit.Script, _a4 uint64) (*transactions.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for BuildTransaction")
	}

	var r0 *transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, transactions.AddressesRoles, int, flowkit.Script, uint64) (*transactions.Transaction, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(context.Context, transactions.AddressesRoles, int, flowkit.Script, uint64) *transactions.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, transactions.AddressesRoles, int, flowkit.Script, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccount provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) CreateAccount(_a0 context.Context, _a1 *accounts.Account, _a2 []accounts.PublicKey) (*flow.Account, flow.Identifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 *flow.Account
	var r1 flow.Identifier
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, []accounts.PublicKey) (*flow.Account, flow.Identifier, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, []accounts.PublicKey) *flow.Account); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *accounts.Account, []accounts.PublicKey) flow.Identifier); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(flow.Identifier)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *accounts.Account, []accounts.PublicKey) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeployProject provides a mock function with given fields: _a0, _a1
func (_m *Services) DeployProject(_a0 context.Context, _a1 flowkit.UpdateContract) ([]*project.Contract, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeployProject")
	}

	var r0 []*project.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.UpdateContract) ([]*project.Contract, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.UpdateContract) []*project.Contract); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*project.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flowkit.UpdateContract) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DerivePrivateKeyFromMnemonic provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Services) DerivePrivateKeyFromMnemonic(_a0 context.Context, _a1 string, _a2 crypto.SigningAlgorithm, _a3 string) (crypto.PrivateKey, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for DerivePrivateKeyFromMnemonic")
	}

	var r0 crypto.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, crypto.SigningAlgorithm, string) (crypto.PrivateKey, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, crypto.SigningAlgorithm, string) crypto.PrivateKey); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, crypto.SigningAlgorithm, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScript provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) ExecuteScript(_a0 context.Context, _a1 flowkit.Script, _a2 flowkit.ScriptQuery) (cadence.Value, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteScript")
	}

	var r0 cadence.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.Script, flowkit.ScriptQuery) (cadence.Value, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.Script, flowkit.ScriptQuery) cadence.Value); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flowkit.Script, flowkit.ScriptQuery) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Gateway provides a mock function with given fields:
func (_m *Services) Gateway() gateway.Gateway {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Gateway")
	}

	var r0 gateway.Gateway
	if rf, ok := ret.Get(0).(func() gateway.Gateway); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gateway.Gateway)
		}
	}

	return r0
}

// GenerateKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) GenerateKey(_a0 context.Context, _a1 crypto.SigningAlgorithm, _a2 string) (crypto.PrivateKey, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GenerateKey")
	}

	var r0 crypto.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, crypto.SigningAlgorithm, string) (crypto.PrivateKey, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, crypto.SigningAlgorithm, string) crypto.PrivateKey); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, crypto.SigningAlgorithm, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateMnemonicKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) GenerateMnemonicKey(_a0 context.Context, _a1 crypto.SigningAlgorithm, _a2 string) (crypto.PrivateKey, string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GenerateMnemonicKey")
	}

	var r0 crypto.PrivateKey
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, crypto.SigningAlgorithm, string) (crypto.PrivateKey, string, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, crypto.SigningAlgorithm, string) crypto.PrivateKey); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, crypto.SigningAlgorithm, string) string); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, crypto.SigningAlgorithm, string) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAccount provides a mock function with given fields: _a0, _a1
func (_m *Services) GetAccount(_a0 context.Context, _a1 flow.Address) (*flow.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 *flow.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, flow.Address) (*flow.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flow.Address) *flow.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flow.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: _a0, _a1
func (_m *Services) GetBlock(_a0 context.Context, _a1 flowkit.BlockQuery) (*flow.Block, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetBlock")
	}

	var r0 *flow.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.BlockQuery) (*flow.Block, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flowkit.BlockQuery) *flow.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flowkit.BlockQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollection provides a mock function with given fields: _a0, _a1
func (_m *Services) GetCollection(_a0 context.Context, _a1 flow.Identifier) (*flow.Collection, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCollection")
	}

	var r0 *flow.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) (*flow.Collection, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) *flow.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Services) GetEvents(_a0 context.Context, _a1 []string, _a2 uint64, _a3 uint64, _a4 *flowkit.EventWorker) ([]flow.BlockEvents, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for GetEvents")
	}

	var r0 []flow.BlockEvents
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, uint64, uint64, *flowkit.EventWorker) ([]flow.BlockEvents, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, uint64, uint64, *flowkit.EventWorker) []flow.BlockEvents); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.BlockEvents)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, uint64, uint64, *flowkit.EventWorker) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) GetTransactionByID(_a0 context.Context, _a1 flow.Identifier, _a2 bool) (*flow.Transaction, *flow.TransactionResult, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByID")
	}

	var r0 *flow.Transaction
	var r1 *flow.TransactionResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier, bool) (*flow.Transaction, *flow.TransactionResult, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier, bool) *flow.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier, bool) *flow.TransactionResult); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*flow.TransactionResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, flow.Identifier, bool) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionsByBlockID provides a mock function with given fields: _a0, _a1
func (_m *Services) GetTransactionsByBlockID(_a0 context.Context, _a1 flow.Identifier) ([]*flow.Transaction, []*flow.TransactionResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionsByBlockID")
	}

	var r0 []*flow.Transaction
	var r1 []*flow.TransactionResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) ([]*flow.Transaction, []*flow.TransactionResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) []*flow.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*flow.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier) []*flow.TransactionResult); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*flow.TransactionResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, flow.Identifier) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Network provides a mock function with given fields:
func (_m *Services) Network() config.Network {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Network")
	}

	var r0 config.Network
	if rf, ok := ret.Get(0).(func() config.Network); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Network)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *Services) Ping() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveContract provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) RemoveContract(_a0 context.Context, _a1 *accounts.Account, _a2 string) (flow.Identifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for RemoveContract")
	}

	var r0 flow.Identifier
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, string) (flow.Identifier, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, string) flow.Identifier); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *accounts.Account, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendSignedTransaction provides a mock function with given fields: _a0, _a1
func (_m *Services) SendSignedTransaction(_a0 context.Context, _a1 *transactions.Transaction) (*flow.Transaction, *flow.TransactionResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SendSignedTransaction")
	}

	var r0 *flow.Transaction
	var r1 *flow.TransactionResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *transactions.Transaction) (*flow.Transaction, *flow.TransactionResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *transactions.Transaction) *flow.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *transactions.Transaction) *flow.TransactionResult); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*flow.TransactionResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *transactions.Transaction) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SendTransaction provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Services) SendTransaction(_a0 context.Context, _a1 transactions.AccountRoles, _a2 flowkit.Script, _a3 uint64) (*flow.Transaction, *flow.TransactionResult, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 *flow.Transaction
	var r1 *flow.TransactionResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, transactions.AccountRoles, flowkit.Script, uint64) (*flow.Transaction, *flow.TransactionResult, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, transactions.AccountRoles, flowkit.Script, uint64) *flow.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, transactions.AccountRoles, flowkit.Script, uint64) *flow.TransactionResult); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*flow.TransactionResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, transactions.AccountRoles, flowkit.Script, uint64) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetLogger provides a mock function with given fields: _a0
func (_m *Services) SetLogger(_a0 output.Logger) {
	_m.Called(_a0)
}

// SignTransactionPayload provides a mock function with given fields: _a0, _a1, _a2
func (_m *Services) SignTransactionPayload(_a0 context.Context, _a1 *accounts.Account, _a2 []byte) (*transactions.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for SignTransactionPayload")
	}

	var r0 *transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, []byte) (*transactions.Transaction, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account, []byte) *transactions.Transaction); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *accounts.Account, []byte) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
