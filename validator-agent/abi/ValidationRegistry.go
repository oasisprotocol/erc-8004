// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidationRegistryMetaData contains all meta data concerning the ValidationRegistry contract.
var ValidationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAgentValidations\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIdentityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSummary\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validatorAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avgResponse\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidationStatus\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"response\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRequests\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"requestHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestExists\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validationRequest\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validationResponse\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"response\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"responseUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidationRequest\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationResponse\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"response\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"responseUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60a0346100c557601f6111dc38819003918201601f19168301916001600160401b038311848410176100ca578084926020946040528339810103126100c557516001600160a01b038116908190036100c5578015610080576080526040516110fb90816100e18239608051818181610260015281816102fd0152610e840152f35b60405162461bcd60e51b815260206004820152601860248201527f496e76616c6964207265676973747279206164647265737300000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163134e18f414610e71575080631b74d04614610e4057806330e5993a14610c3a5780634bf3158c14610bc25780638d5d0c2d14610b48578063a0aa15a11461090c578063aaf400c41461028f578063bc4d861b1461024a578063fb1e61ca146101005763ff2febfc1461008d57600080fd5b346100fb5760203660031901126100fb57600435600052600160205260c0604060002060018060a01b038154169060018101549060ff6002820154166003820154906005600484015493015493604051958652602086015260408501526060840152608083015260a0820152f35b600080fd5b346100fb576020806003193601126100fb5760043560009081528082526040902080546001600160a01b03169190610139831515611076565b6001918282015460046002840193015492604051809560009083549361015e85610fd7565b948585528783821691826000146102285750506001146101e9575b505061018b9250969492960385610f32565b6040519485938452828401526080604084015283519182608085015260005b8381106101d257505060a0935060008483850101526060830152601f80199101168101030190f35b80860182015187820160a0015286945081016101aa565b86925060005281600020906000915b85831061021057505061018b93508201018880610179565b8054838b0185015289945087939092019181016101f8565b925093505061018b94915060ff191682840152151560051b8201018880610179565b346100fb5760003660031901126100fb576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100fb5760803660031901126100fb576102a8610ee1565b60449067ffffffffffffffff82358181116100fb576102cb903690600401610eb3565b916001600160a01b0391848316156108c85783156108905760405163de99f15760e01b815260248035600483015292907f00000000000000000000000000000000000000000000000000000000000000008516906020818681855afa9081156107a357600091610871575b5015610837576040516331a9108f60e11b8152843560048201526020818681855afa80156107a3578691600091610818575b501690848233149182156107af575b821561073a575b50501561070657610393908588161415610f8b565b6103a1338588161415610f8b565b6064359687156106af575b87600052600460205260ff6040600020541661066c575060405160a081018181108382111761065757604052848716815260208101918435835280871161065757604051610404601f8901601f191660200182610f32565b87815236888601116100fb578785602083013760006020898301015260408301908152896060840152426080840152896000526000602052604060002093878451166001600160601b0360a01b865416178555516001850155518051918211610642576104746002850154610fd7565b601f81116105fb575b50602090601f831160011461055e577f530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a5059979695938361054e9694600494608094600092610553575b50508160011b916000199060031b1c19161760028501555b60608101516003850155015191015582356000526002602052610504886040600020611011565b838716600052600360205261051d886040600020611011565b8760005260046020526040600020600160ff1982541617905560405193849360208552359716956020840191611055565b0390a4005b015190508e806104c5565b906002850160005260206000209160005b601f19851681106105e35750936001846004946080947f530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a50599c9b9a9861054e9a98601f198116106105ca575b505050811b0160028501556104dd565b015160001960f88460031b161c191690558e80806105ba565b9192602060018192868501518155019401920161056f565b600285016000526020600020601f840160051c81016020851061063b575b601f830160051c8201811061062f57505061047d565b60008155600101610619565b5080610619565b85634e487b7160e01b60005260416004526000fd5b84634e487b7160e01b60005260416004526000fd5b837f52657175657374206861736820616c7265616479206578697374730000000000606492601b6040519362461bcd60e51b855260206004860152840152820152fd5b965060405160208101906001600160601b03198860601b16825284356034820152868460548301376106fd6088828981014260548201523360601b6074820152036068810184520182610f32565b519020966103ac565b60405162461bcd60e51b815260206004820152600e818601526d139bdd08185d5d1a1bdc9a5e995960921b818a0152606490fd5b60405163020604bf60e21b8152823560048201529250602091839182905afa9081156107a357600091610774575b5085163314848a61037e565b610796915060203d60201161079c575b61078e8183610f32565b810190610f6c565b89610768565b503d610784565b6040513d6000823e3d90fd5b91505060405163e985e9c560e01b815282600482015233868201526020818b81855afa80156107a35786916000916107e9575b5091610377565b61080b915060203d602011610811575b6108038183610f32565b810190610f54565b8b6107e2565b503d6107f9565b610831915060203d60201161079c5761078e8183610f32565b8a610368565b60405162461bcd60e51b815260206004820152601481860152731059d95b9d08191bd95cc81b9bdd08195e1a5cdd60621b818a0152606490fd5b61088a915060203d602011610811576108038183610f32565b89610336565b60405162461bcd60e51b8152602060048201526011602482015270456d70747920726571756573742055524960781b81880152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f7220616464726573730000000000000081880152606490fd5b346100fb5760603660031901126100fb5767ffffffffffffffff602435116100fb573660236024350112156100fb5767ffffffffffffffff60243560040135116100fb5736602480356004013560051b81350101116100fb57600435600052600260205260406000206040518082602082945493848152019060005260206000209260005b818110610b2f5750506109a692500382610f32565b600080805b8351821015610ade5760208260051b850101516000526001602052604060002060018060a01b038154168015610ad15760243560040135610a6b575b506044351515604435610a5b575b610a4f57600260ff910154168101809111610a395767ffffffffffffffff90921667ffffffffffffffff8114610a39576001610a329101916110b6565b90916109ab565b634e487b7160e01b600052601160045260246000fd5b509190610a32906110b6565b50604435600482015414156109f5565b60009060005b602435600401358110610a8c575b505015610a4f57856109e7565b60248035600583901b0101356001600160a01b03811681036100fb576001600160a01b03168214610ac557610ac0906110b6565b610a71565b50505060018680610a7f565b50509190610a32906110b6565b67ffffffffffffffff8316801580610b2157610b0b5760ff808260409404165b8351928352166020820152f35b634e487b7160e01b600052601260045260246000fd5b506040915060ff6000610afe565b8454835260019485019486945060209093019201610991565b346100fb576020806003193601126100fb5760043560005260028152604060002090604051908181845491828152019360005281600020916000905b828210610bab57610ba785610b9b81890382610f32565b60405191829182610ef7565b0390f35b835486529485019460019384019390910190610b84565b346100fb576020806003193601126100fb576001600160a01b03610be4610ee1565b1660005260038152604060002090604051908181845491828152019360005281600020916000905b828210610c2357610ba785610b9b81890382610f32565b835486529485019460019384019390910190610c0c565b346100fb5760a03660031901126100fb5760043560243560ff81168091036100fb5767ffffffffffffffff906044358281116100fb57610c7e903690600401610eb3565b606493919335926084359360648211610e025786600052600060205260406000209460018060a01b038087541690610cb7821515611076565b813303610dbd5760018801978854926040519060c0820199828b10908b1117610da7578c610d97956005927ff224d3d5ad74301be48e4d51ca5f1b24c7946875887327585becc59165297dcf9c604052845260208401908152604084018a815260608501908a825260808601928a845260a087019442865260005260016020528860406000209751166001600160601b0360a01b88541617875551600187015560ff6002870191511660ff198254161790555160038501555160048401555191015554169654976040519586958652608060208701526080860191611055565b91604084015260608301520390a4005b634e487b7160e01b600052604160045260246000fd5b60405162461bcd60e51b815260206004820152601860248201527f4e6f7420617574686f72697a65642076616c696461746f7200000000000000006044820152606490fd5b60405162461bcd60e51b81526020600482015260166024820152750526573706f6e7365206d75737420626520302d3130360541b6044820152606490fd5b346100fb5760203660031901126100fb576004356000526004602052602060ff604060002054166040519015158152f35b346100fb5760003660031901126100fb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156100fb5782359167ffffffffffffffff83116100fb57602083818601950101116100fb57565b600435906001600160a01b03821682036100fb57565b6020908160408183019282815285518094520193019160005b828110610f1e575050505090565b835185529381019392810192600101610f10565b90601f8019910116810190811067ffffffffffffffff821117610da757604052565b908160209103126100fb575180151581036100fb5790565b908160209103126100fb57516001600160a01b03811681036100fb5790565b15610f9257565b60405162461bcd60e51b815260206004820152601b60248201527f53656c662d76616c69646174696f6e206e6f7420616c6c6f77656400000000006044820152606490fd5b90600182811c92168015611007575b6020831014610ff157565b634e487b7160e01b600052602260045260246000fd5b91607f1691610fe6565b80549068010000000000000000821015610da7576001820180825582101561103f5760005260206000200155565b634e487b7160e01b600052603260045260246000fd5b908060209392818452848401376000828201840152601f01601f1916010190565b1561107d57565b60405162461bcd60e51b815260206004820152601160248201527014995c5d595cdd081b9bdd08199bdd5b99607a1b6044820152606490fd5b6000198114610a39576001019056fea2646970667358221220076946650a60e10f9ffb162ed4a86dc706e85caaf6d75971aee7d4549c99904a64736f6c63430008130033",
}

// ValidationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidationRegistryMetaData.ABI instead.
var ValidationRegistryABI = ValidationRegistryMetaData.ABI

// ValidationRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidationRegistryMetaData.Bin instead.
var ValidationRegistryBin = ValidationRegistryMetaData.Bin

// DeployValidationRegistry deploys a new Ethereum contract, binding an instance of ValidationRegistry to it.
func DeployValidationRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _identityRegistry common.Address) (common.Address, *types.Transaction, *ValidationRegistry, error) {
	parsed, err := ValidationRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidationRegistryBin), backend, _identityRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidationRegistry{ValidationRegistryCaller: ValidationRegistryCaller{contract: contract}, ValidationRegistryTransactor: ValidationRegistryTransactor{contract: contract}, ValidationRegistryFilterer: ValidationRegistryFilterer{contract: contract}}, nil
}

// ValidationRegistry is an auto generated Go binding around an Ethereum contract.
type ValidationRegistry struct {
	ValidationRegistryCaller     // Read-only binding to the contract
	ValidationRegistryTransactor // Write-only binding to the contract
	ValidationRegistryFilterer   // Log filterer for contract events
}

// ValidationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidationRegistrySession struct {
	Contract     *ValidationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ValidationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidationRegistryCallerSession struct {
	Contract *ValidationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ValidationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidationRegistryTransactorSession struct {
	Contract     *ValidationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ValidationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidationRegistryRaw struct {
	Contract *ValidationRegistry // Generic contract binding to access the raw methods on
}

// ValidationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidationRegistryCallerRaw struct {
	Contract *ValidationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidationRegistryTransactorRaw struct {
	Contract *ValidationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidationRegistry creates a new instance of ValidationRegistry, bound to a specific deployed contract.
func NewValidationRegistry(address common.Address, backend bind.ContractBackend) (*ValidationRegistry, error) {
	contract, err := bindValidationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistry{ValidationRegistryCaller: ValidationRegistryCaller{contract: contract}, ValidationRegistryTransactor: ValidationRegistryTransactor{contract: contract}, ValidationRegistryFilterer: ValidationRegistryFilterer{contract: contract}}, nil
}

// NewValidationRegistryCaller creates a new read-only instance of ValidationRegistry, bound to a specific deployed contract.
func NewValidationRegistryCaller(address common.Address, caller bind.ContractCaller) (*ValidationRegistryCaller, error) {
	contract, err := bindValidationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistryCaller{contract: contract}, nil
}

// NewValidationRegistryTransactor creates a new write-only instance of ValidationRegistry, bound to a specific deployed contract.
func NewValidationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidationRegistryTransactor, error) {
	contract, err := bindValidationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistryTransactor{contract: contract}, nil
}

// NewValidationRegistryFilterer creates a new log filterer instance of ValidationRegistry, bound to a specific deployed contract.
func NewValidationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidationRegistryFilterer, error) {
	contract, err := bindValidationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistryFilterer{contract: contract}, nil
}

// bindValidationRegistry binds a generic wrapper to an already deployed contract.
func bindValidationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRegistry *ValidationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRegistry.Contract.ValidationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRegistry *ValidationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRegistry *ValidationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRegistry *ValidationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRegistry *ValidationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRegistry *ValidationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAgentValidations is a free data retrieval call binding the contract method 0x8d5d0c2d.
//
// Solidity: function getAgentValidations(uint256 agentId) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistryCaller) GetAgentValidations(opts *bind.CallOpts, agentId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getAgentValidations", agentId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAgentValidations is a free data retrieval call binding the contract method 0x8d5d0c2d.
//
// Solidity: function getAgentValidations(uint256 agentId) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistrySession) GetAgentValidations(agentId *big.Int) ([][32]byte, error) {
	return _ValidationRegistry.Contract.GetAgentValidations(&_ValidationRegistry.CallOpts, agentId)
}

// GetAgentValidations is a free data retrieval call binding the contract method 0x8d5d0c2d.
//
// Solidity: function getAgentValidations(uint256 agentId) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetAgentValidations(agentId *big.Int) ([][32]byte, error) {
	return _ValidationRegistry.Contract.GetAgentValidations(&_ValidationRegistry.CallOpts, agentId)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ValidationRegistry *ValidationRegistryCaller) GetIdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getIdentityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ValidationRegistry *ValidationRegistrySession) GetIdentityRegistry() (common.Address, error) {
	return _ValidationRegistry.Contract.GetIdentityRegistry(&_ValidationRegistry.CallOpts)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetIdentityRegistry() (common.Address, error) {
	return _ValidationRegistry.Contract.GetIdentityRegistry(&_ValidationRegistry.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, string requestUri, uint256 timestamp)
func (_ValidationRegistry *ValidationRegistryCaller) GetRequest(opts *bind.CallOpts, requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestUri       string
	Timestamp        *big.Int
}, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getRequest", requestHash)

	outstruct := new(struct {
		ValidatorAddress common.Address
		AgentId          *big.Int
		RequestUri       string
		Timestamp        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AgentId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestUri = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, string requestUri, uint256 timestamp)
func (_ValidationRegistry *ValidationRegistrySession) GetRequest(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestUri       string
	Timestamp        *big.Int
}, error) {
	return _ValidationRegistry.Contract.GetRequest(&_ValidationRegistry.CallOpts, requestHash)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, string requestUri, uint256 timestamp)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetRequest(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestUri       string
	Timestamp        *big.Int
}, error) {
	return _ValidationRegistry.Contract.GetRequest(&_ValidationRegistry.CallOpts, requestHash)
}

// GetSummary is a free data retrieval call binding the contract method 0xa0aa15a1.
//
// Solidity: function getSummary(uint256 agentId, address[] validatorAddresses, bytes32 tag) view returns(uint64 count, uint8 avgResponse)
func (_ValidationRegistry *ValidationRegistryCaller) GetSummary(opts *bind.CallOpts, agentId *big.Int, validatorAddresses []common.Address, tag [32]byte) (struct {
	Count       uint64
	AvgResponse uint8
}, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getSummary", agentId, validatorAddresses, tag)

	outstruct := new(struct {
		Count       uint64
		AvgResponse uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AvgResponse = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSummary is a free data retrieval call binding the contract method 0xa0aa15a1.
//
// Solidity: function getSummary(uint256 agentId, address[] validatorAddresses, bytes32 tag) view returns(uint64 count, uint8 avgResponse)
func (_ValidationRegistry *ValidationRegistrySession) GetSummary(agentId *big.Int, validatorAddresses []common.Address, tag [32]byte) (struct {
	Count       uint64
	AvgResponse uint8
}, error) {
	return _ValidationRegistry.Contract.GetSummary(&_ValidationRegistry.CallOpts, agentId, validatorAddresses, tag)
}

// GetSummary is a free data retrieval call binding the contract method 0xa0aa15a1.
//
// Solidity: function getSummary(uint256 agentId, address[] validatorAddresses, bytes32 tag) view returns(uint64 count, uint8 avgResponse)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetSummary(agentId *big.Int, validatorAddresses []common.Address, tag [32]byte) (struct {
	Count       uint64
	AvgResponse uint8
}, error) {
	return _ValidationRegistry.Contract.GetSummary(&_ValidationRegistry.CallOpts, agentId, validatorAddresses, tag)
}

// GetValidationStatus is a free data retrieval call binding the contract method 0xff2febfc.
//
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 responseHash, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistryCaller) GetValidationStatus(opts *bind.CallOpts, requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	ResponseHash     [32]byte
	Tag              [32]byte
	LastUpdate       *big.Int
}, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getValidationStatus", requestHash)

	outstruct := new(struct {
		ValidatorAddress common.Address
		AgentId          *big.Int
		Response         uint8
		ResponseHash     [32]byte
		Tag              [32]byte
		LastUpdate       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AgentId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Response = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ResponseHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Tag = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.LastUpdate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidationStatus is a free data retrieval call binding the contract method 0xff2febfc.
//
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 responseHash, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistrySession) GetValidationStatus(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	ResponseHash     [32]byte
	Tag              [32]byte
	LastUpdate       *big.Int
}, error) {
	return _ValidationRegistry.Contract.GetValidationStatus(&_ValidationRegistry.CallOpts, requestHash)
}

// GetValidationStatus is a free data retrieval call binding the contract method 0xff2febfc.
//
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 responseHash, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetValidationStatus(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	ResponseHash     [32]byte
	Tag              [32]byte
	LastUpdate       *big.Int
}, error) {
	return _ValidationRegistry.Contract.GetValidationStatus(&_ValidationRegistry.CallOpts, requestHash)
}

// GetValidatorRequests is a free data retrieval call binding the contract method 0x4bf3158c.
//
// Solidity: function getValidatorRequests(address validatorAddress) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistryCaller) GetValidatorRequests(opts *bind.CallOpts, validatorAddress common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getValidatorRequests", validatorAddress)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetValidatorRequests is a free data retrieval call binding the contract method 0x4bf3158c.
//
// Solidity: function getValidatorRequests(address validatorAddress) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistrySession) GetValidatorRequests(validatorAddress common.Address) ([][32]byte, error) {
	return _ValidationRegistry.Contract.GetValidatorRequests(&_ValidationRegistry.CallOpts, validatorAddress)
}

// GetValidatorRequests is a free data retrieval call binding the contract method 0x4bf3158c.
//
// Solidity: function getValidatorRequests(address validatorAddress) view returns(bytes32[] requestHashes)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetValidatorRequests(validatorAddress common.Address) ([][32]byte, error) {
	return _ValidationRegistry.Contract.GetValidatorRequests(&_ValidationRegistry.CallOpts, validatorAddress)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ValidationRegistry *ValidationRegistryCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ValidationRegistry *ValidationRegistrySession) IdentityRegistry() (common.Address, error) {
	return _ValidationRegistry.Contract.IdentityRegistry(&_ValidationRegistry.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ValidationRegistry *ValidationRegistryCallerSession) IdentityRegistry() (common.Address, error) {
	return _ValidationRegistry.Contract.IdentityRegistry(&_ValidationRegistry.CallOpts)
}

// RequestExists is a free data retrieval call binding the contract method 0x1b74d046.
//
// Solidity: function requestExists(bytes32 requestHash) view returns(bool exists)
func (_ValidationRegistry *ValidationRegistryCaller) RequestExists(opts *bind.CallOpts, requestHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "requestExists", requestHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestExists is a free data retrieval call binding the contract method 0x1b74d046.
//
// Solidity: function requestExists(bytes32 requestHash) view returns(bool exists)
func (_ValidationRegistry *ValidationRegistrySession) RequestExists(requestHash [32]byte) (bool, error) {
	return _ValidationRegistry.Contract.RequestExists(&_ValidationRegistry.CallOpts, requestHash)
}

// RequestExists is a free data retrieval call binding the contract method 0x1b74d046.
//
// Solidity: function requestExists(bytes32 requestHash) view returns(bool exists)
func (_ValidationRegistry *ValidationRegistryCallerSession) RequestExists(requestHash [32]byte) (bool, error) {
	return _ValidationRegistry.Contract.RequestExists(&_ValidationRegistry.CallOpts, requestHash)
}

// ValidationRequest is a paid mutator transaction binding the contract method 0xaaf400c4.
//
// Solidity: function validationRequest(address validatorAddress, uint256 agentId, string requestUri, bytes32 requestHash) returns()
func (_ValidationRegistry *ValidationRegistryTransactor) ValidationRequest(opts *bind.TransactOpts, validatorAddress common.Address, agentId *big.Int, requestUri string, requestHash [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.contract.Transact(opts, "validationRequest", validatorAddress, agentId, requestUri, requestHash)
}

// ValidationRequest is a paid mutator transaction binding the contract method 0xaaf400c4.
//
// Solidity: function validationRequest(address validatorAddress, uint256 agentId, string requestUri, bytes32 requestHash) returns()
func (_ValidationRegistry *ValidationRegistrySession) ValidationRequest(validatorAddress common.Address, agentId *big.Int, requestUri string, requestHash [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationRequest(&_ValidationRegistry.TransactOpts, validatorAddress, agentId, requestUri, requestHash)
}

// ValidationRequest is a paid mutator transaction binding the contract method 0xaaf400c4.
//
// Solidity: function validationRequest(address validatorAddress, uint256 agentId, string requestUri, bytes32 requestHash) returns()
func (_ValidationRegistry *ValidationRegistryTransactorSession) ValidationRequest(validatorAddress common.Address, agentId *big.Int, requestUri string, requestHash [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationRequest(&_ValidationRegistry.TransactOpts, validatorAddress, agentId, requestUri, requestHash)
}

// ValidationResponse is a paid mutator transaction binding the contract method 0x30e5993a.
//
// Solidity: function validationResponse(bytes32 requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag) returns()
func (_ValidationRegistry *ValidationRegistryTransactor) ValidationResponse(opts *bind.TransactOpts, requestHash [32]byte, response uint8, responseUri string, responseHash [32]byte, tag [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.contract.Transact(opts, "validationResponse", requestHash, response, responseUri, responseHash, tag)
}

// ValidationResponse is a paid mutator transaction binding the contract method 0x30e5993a.
//
// Solidity: function validationResponse(bytes32 requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag) returns()
func (_ValidationRegistry *ValidationRegistrySession) ValidationResponse(requestHash [32]byte, response uint8, responseUri string, responseHash [32]byte, tag [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationResponse(&_ValidationRegistry.TransactOpts, requestHash, response, responseUri, responseHash, tag)
}

// ValidationResponse is a paid mutator transaction binding the contract method 0x30e5993a.
//
// Solidity: function validationResponse(bytes32 requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag) returns()
func (_ValidationRegistry *ValidationRegistryTransactorSession) ValidationResponse(requestHash [32]byte, response uint8, responseUri string, responseHash [32]byte, tag [32]byte) (*types.Transaction, error) {
	return _ValidationRegistry.Contract.ValidationResponse(&_ValidationRegistry.TransactOpts, requestHash, response, responseUri, responseHash, tag)
}

// ValidationRegistryValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the ValidationRegistry contract.
type ValidationRegistryValidationRequestIterator struct {
	Event *ValidationRegistryValidationRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidationRegistryValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRegistryValidationRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidationRegistryValidationRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidationRegistryValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRegistryValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRegistryValidationRequest represents a ValidationRequest event raised by the ValidationRegistry contract.
type ValidationRegistryValidationRequest struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestUri       string
	RequestHash      [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0x530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a5059.
//
// Solidity: event ValidationRequest(address indexed validatorAddress, uint256 indexed agentId, string requestUri, bytes32 indexed requestHash)
func (_ValidationRegistry *ValidationRegistryFilterer) FilterValidationRequest(opts *bind.FilterOpts, validatorAddress []common.Address, agentId []*big.Int, requestHash [][32]byte) (*ValidationRegistryValidationRequestIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _ValidationRegistry.contract.FilterLogs(opts, "ValidationRequest", validatorAddressRule, agentIdRule, requestHashRule)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistryValidationRequestIterator{contract: _ValidationRegistry.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0x530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a5059.
//
// Solidity: event ValidationRequest(address indexed validatorAddress, uint256 indexed agentId, string requestUri, bytes32 indexed requestHash)
func (_ValidationRegistry *ValidationRegistryFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *ValidationRegistryValidationRequest, validatorAddress []common.Address, agentId []*big.Int, requestHash [][32]byte) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _ValidationRegistry.contract.WatchLogs(opts, "ValidationRequest", validatorAddressRule, agentIdRule, requestHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRegistryValidationRequest)
				if err := _ValidationRegistry.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationRequest is a log parse operation binding the contract event 0x530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a5059.
//
// Solidity: event ValidationRequest(address indexed validatorAddress, uint256 indexed agentId, string requestUri, bytes32 indexed requestHash)
func (_ValidationRegistry *ValidationRegistryFilterer) ParseValidationRequest(log types.Log) (*ValidationRegistryValidationRequest, error) {
	event := new(ValidationRegistryValidationRequest)
	if err := _ValidationRegistry.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidationRegistryValidationResponseIterator is returned from FilterValidationResponse and is used to iterate over the raw logs and unpacked data for ValidationResponse events raised by the ValidationRegistry contract.
type ValidationRegistryValidationResponseIterator struct {
	Event *ValidationRegistryValidationResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidationRegistryValidationResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRegistryValidationResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidationRegistryValidationResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidationRegistryValidationResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRegistryValidationResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRegistryValidationResponse represents a ValidationResponse event raised by the ValidationRegistry contract.
type ValidationRegistryValidationResponse struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	RequestHash      [32]byte
	Response         uint8
	ResponseUri      string
	ResponseHash     [32]byte
	Tag              [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidationResponse is a free log retrieval operation binding the contract event 0xf224d3d5ad74301be48e4d51ca5f1b24c7946875887327585becc59165297dcf.
//
// Solidity: event ValidationResponse(address indexed validatorAddress, uint256 indexed agentId, bytes32 indexed requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag)
func (_ValidationRegistry *ValidationRegistryFilterer) FilterValidationResponse(opts *bind.FilterOpts, validatorAddress []common.Address, agentId []*big.Int, requestHash [][32]byte) (*ValidationRegistryValidationResponseIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _ValidationRegistry.contract.FilterLogs(opts, "ValidationResponse", validatorAddressRule, agentIdRule, requestHashRule)
	if err != nil {
		return nil, err
	}
	return &ValidationRegistryValidationResponseIterator{contract: _ValidationRegistry.contract, event: "ValidationResponse", logs: logs, sub: sub}, nil
}

// WatchValidationResponse is a free log subscription operation binding the contract event 0xf224d3d5ad74301be48e4d51ca5f1b24c7946875887327585becc59165297dcf.
//
// Solidity: event ValidationResponse(address indexed validatorAddress, uint256 indexed agentId, bytes32 indexed requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag)
func (_ValidationRegistry *ValidationRegistryFilterer) WatchValidationResponse(opts *bind.WatchOpts, sink chan<- *ValidationRegistryValidationResponse, validatorAddress []common.Address, agentId []*big.Int, requestHash [][32]byte) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _ValidationRegistry.contract.WatchLogs(opts, "ValidationResponse", validatorAddressRule, agentIdRule, requestHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRegistryValidationResponse)
				if err := _ValidationRegistry.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationResponse is a log parse operation binding the contract event 0xf224d3d5ad74301be48e4d51ca5f1b24c7946875887327585becc59165297dcf.
//
// Solidity: event ValidationResponse(address indexed validatorAddress, uint256 indexed agentId, bytes32 indexed requestHash, uint8 response, string responseUri, bytes32 responseHash, bytes32 tag)
func (_ValidationRegistry *ValidationRegistryFilterer) ParseValidationResponse(log types.Log) (*ValidationRegistryValidationResponse, error) {
	event := new(ValidationRegistryValidationResponse)
	if err := _ValidationRegistry.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
