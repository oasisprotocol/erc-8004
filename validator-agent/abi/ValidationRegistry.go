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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAgentValidations\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIdentityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSummary\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validatorAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avgResponse\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidationStatus\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"response\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRequests\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"requestHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestExists\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validationRequest\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validationResponse\",\"inputs\":[{\"name\":\"requestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"response\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"responseUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidationRequest\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationResponse\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"response\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"responseUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"tag\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60a0346100c557601f6111c738819003918201601f19168301916001600160401b038311848410176100ca578084926020946040528339810103126100c557516001600160a01b038116908190036100c5578015610080576080526040516110e690816100e18239608051818181610254015281816102f10152610e3d0152f35b60405162461bcd60e51b815260206004820152601860248201527f496e76616c6964207265676973747279206164647265737300000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163134e18f414610e2a575080631b74d04614610df957806330e5993a14610c255780634bf3158c14610bad5780638d5d0c2d14610b33578063a0aa15a1146108f7578063aaf400c414610283578063bc4d861b1461023e578063fb1e61ca146100f45763ff2febfc1461008d57600080fd5b346100ef5760203660031901126100ef57600435600052600160205260a06040600020600180831b038154169060018101549060ff60028201541660046003830154920154926040519485526020850152604084015260608301526080820152f35b600080fd5b346100ef576020806003193601126100ef5760043560009081528082526040902080546001600160a01b0316919061012d831515611061565b6001918282015460046002840193015492604051809560009083549361015285610fc2565b9485855287838216918260001461021c5750506001146101dd575b505061017f9250969492960385610f1d565b6040519485938452828401526080604084015283519182608085015260005b8381106101c657505060a0935060008483850101526060830152601f80199101168101030190f35b80860182015187820160a00152869450810161019e565b86925060005281600020906000915b85831061020457505061017f9350820101888061016d565b8054838b0185015289945087939092019181016101ec565b925093505061017f94915060ff191682840152151560051b820101888061016d565b346100ef5760003660031901126100ef576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100ef5760803660031901126100ef5761029c610e9a565b60449067ffffffffffffffff82358181116100ef576102bf903690600401610e6c565b916001600160a01b0391848316156108b357831561087b5760405163de99f15760e01b815260248035600483015292907f00000000000000000000000000000000000000000000000000000000000000008516906020818681855afa90811561078e5760009161085c575b5015610822576040516331a9108f60e11b8152843560048201526020818681855afa801561078e578691600091610803575b5016908482331491821561079a575b8215610725575b5050156106f157610387908588161415610f76565b610395338588161415610f76565b60643596871561069a575b87600052600460205260ff6040600020541661065757506040516103c381610eeb565b8487168152602081019184358352808711610642576040516103ef601f8901601f191660200182610f1d565b87815236888601116100ef578785602083013760006020898301015260408301908152896060840152426080840152896000526000602052604060002093878451166001600160601b0360a01b86541617855551600185015551805191821161062d5761045f6002850154610fc2565b601f81116105e6575b50602090601f8311600114610549577f530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a50599796959383610539969460049460809460009261053e575b50508160011b916000199060031b1c19161760028501555b606081015160038501550151910155823560005260026020526104ef886040600020610ffc565b8387166000526003602052610508886040600020610ffc565b8760005260046020526040600020600160ff1982541617905560405193849360208552359716956020840191611040565b0390a4005b015190508e806104b0565b906002850160005260206000209160005b601f19851681106105ce5750936001846004946080947f530436c3634a98e1e626b0898be2f1e9980cc1bd2a78c07a0aba52d0a48a50599c9b9a986105399a98601f198116106105b5575b505050811b0160028501556104c8565b015160001960f88460031b161c191690558e80806105a5565b9192602060018192868501518155019401920161055a565b600285016000526020600020601f840160051c810160208510610626575b601f830160051c8201811061061a575050610468565b60008155600101610604565b5080610604565b85634e487b7160e01b60005260416004526000fd5b84634e487b7160e01b60005260416004526000fd5b837f52657175657374206861736820616c7265616479206578697374730000000000606492601b6040519362461bcd60e51b855260206004860152840152820152fd5b965060405160208101906001600160601b03198860601b16825284356034820152868460548301376106e86088828981014260548201523360601b6074820152036068810184520182610f1d565b519020966103a0565b60405162461bcd60e51b815260206004820152600e818601526d139bdd08185d5d1a1bdc9a5e995960921b818a0152606490fd5b60405163020604bf60e21b8152823560048201529250602091839182905afa90811561078e5760009161075f575b5085163314848a610372565b610781915060203d602011610787575b6107798183610f1d565b810190610f57565b89610753565b503d61076f565b6040513d6000823e3d90fd5b91505060405163e985e9c560e01b815282600482015233868201526020818b81855afa801561078e5786916000916107d4575b509161036b565b6107f6915060203d6020116107fc575b6107ee8183610f1d565b810190610f3f565b8b6107cd565b503d6107e4565b61081c915060203d602011610787576107798183610f1d565b8a61035c565b60405162461bcd60e51b815260206004820152601481860152731059d95b9d08191bd95cc81b9bdd08195e1a5cdd60621b818a0152606490fd5b610875915060203d6020116107fc576107ee8183610f1d565b8961032a565b60405162461bcd60e51b8152602060048201526011602482015270456d70747920726571756573742055524960781b81880152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f7220616464726573730000000000000081880152606490fd5b346100ef5760603660031901126100ef5767ffffffffffffffff602435116100ef573660236024350112156100ef5767ffffffffffffffff60243560040135116100ef5736602480356004013560051b81350101116100ef57600435600052600260205260406000206040518082602082945493848152019060005260206000209260005b818110610b1a57505061099192500382610f1d565b600080805b8351821015610ac95760208260051b850101516000526001602052604060002060018060a01b038154168015610abc5760243560040135610a56575b506044351515604435610a46575b610a3a57600260ff910154168101809111610a245767ffffffffffffffff90921667ffffffffffffffff8114610a24576001610a1d9101916110a1565b9091610996565b634e487b7160e01b600052601160045260246000fd5b509190610a1d906110a1565b50604435600382015414156109e0565b60009060005b602435600401358110610a77575b505015610a3a57856109d2565b60248035600583901b0101356001600160a01b03811681036100ef576001600160a01b03168214610ab057610aab906110a1565b610a5c565b50505060018680610a6a565b50509190610a1d906110a1565b67ffffffffffffffff8316801580610b0c57610af65760ff808260409404165b8351928352166020820152f35b634e487b7160e01b600052601260045260246000fd5b506040915060ff6000610ae9565b845483526001948501948694506020909301920161097c565b346100ef576020806003193601126100ef5760043560005260028152604060002090604051908181845491828152019360005281600020916000905b828210610b9657610b9285610b8681890382610f1d565b60405191829182610eb0565b0390f35b835486529485019460019384019390910190610b6f565b346100ef576020806003193601126100ef576001600160a01b03610bcf610e9a565b1660005260038152604060002090604051908181845491828152019360005281600020916000905b828210610c0e57610b9285610b8681890382610f1d565b835486529485019460019384019390910190610bf7565b346100ef5760a03660031901126100ef5760043560243560ff81168091036100ef5760443567ffffffffffffffff81116100ef57610c67903690600401610e6c565b6084929192359060648311610dbb5784600052600060205260406000209260018060a01b03928385541694610c9d861515611061565b853303610d76577ff224d3d5ad74301be48e4d51ca5f1b24c7946875887327585becc59165297dcf94816001610d6393019760048b8a5460405193610ce185610eeb565b845260208401908152604084018981526060850191898352608086019342855260005260016020528760406000209651166001600160601b0360a01b87541617865551600186015560ff6002860191511660ff198254161790555160038401555191015554169554966040519485948552608060208601526080850191611040565b90606435604084015260608301520390a4005b60405162461bcd60e51b815260206004820152601860248201527f4e6f7420617574686f72697a65642076616c696461746f7200000000000000006044820152606490fd5b60405162461bcd60e51b81526020600482015260166024820152750526573706f6e7365206d75737420626520302d3130360541b6044820152606490fd5b346100ef5760203660031901126100ef576004356000526004602052602060ff604060002054166040519015158152f35b346100ef5760003660031901126100ef577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b9181601f840112156100ef5782359167ffffffffffffffff83116100ef57602083818601950101116100ef57565b600435906001600160a01b03821682036100ef57565b6020908160408183019282815285518094520193019160005b828110610ed7575050505090565b835185529381019392810192600101610ec9565b60a0810190811067ffffffffffffffff821117610f0757604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff821117610f0757604052565b908160209103126100ef575180151581036100ef5790565b908160209103126100ef57516001600160a01b03811681036100ef5790565b15610f7d57565b60405162461bcd60e51b815260206004820152601b60248201527f53656c662d76616c69646174696f6e206e6f7420616c6c6f77656400000000006044820152606490fd5b90600182811c92168015610ff2575b6020831014610fdc57565b634e487b7160e01b600052602260045260246000fd5b91607f1691610fd1565b80549068010000000000000000821015610f07576001820180825582101561102a5760005260206000200155565b634e487b7160e01b600052603260045260246000fd5b908060209392818452848401376000828201840152601f01601f1916010190565b1561106857565b60405162461bcd60e51b815260206004820152601160248201527014995c5d595cdd081b9bdd08199bdd5b99607a1b6044820152606490fd5b6000198114610a24576001019056fea26469706673582212204019deb01e8db88e51e77454ff8e8c4f15e4493107b78f0ed7ac38d7083ce6b864736f6c63430008130033",
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
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistryCaller) GetValidationStatus(opts *bind.CallOpts, requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	Tag              [32]byte
	LastUpdate       *big.Int
}, error) {
	var out []interface{}
	err := _ValidationRegistry.contract.Call(opts, &out, "getValidationStatus", requestHash)

	outstruct := new(struct {
		ValidatorAddress common.Address
		AgentId          *big.Int
		Response         uint8
		Tag              [32]byte
		LastUpdate       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AgentId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Response = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Tag = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.LastUpdate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidationStatus is a free data retrieval call binding the contract method 0xff2febfc.
//
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistrySession) GetValidationStatus(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
	Tag              [32]byte
	LastUpdate       *big.Int
}, error) {
	return _ValidationRegistry.Contract.GetValidationStatus(&_ValidationRegistry.CallOpts, requestHash)
}

// GetValidationStatus is a free data retrieval call binding the contract method 0xff2febfc.
//
// Solidity: function getValidationStatus(bytes32 requestHash) view returns(address validatorAddress, uint256 agentId, uint8 response, bytes32 tag, uint256 lastUpdate)
func (_ValidationRegistry *ValidationRegistryCallerSession) GetValidationStatus(requestHash [32]byte) (struct {
	ValidatorAddress common.Address
	AgentId          *big.Int
	Response         uint8
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
