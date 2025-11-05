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

// IIdentityRegistryMetadataEntry is an auto generated low-level Go binding around an user-defined struct.
type IIdentityRegistryMetadataEntry struct {
	Key   string
	Value []byte
}

// IdentityRegistryMetaData contains all meta data concerning the IdentityRegistry contract.
var IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agentExists\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadata\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[],\"outputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"tokenURI_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"tuple[]\",\"internalType\":\"structIIdentityRegistry.MetadataEntry[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"tokenURI_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadata\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAgents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataSet\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"indexedKey\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Registered\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080346200032d576001600160401b03906040908082018381118282101762000317578252601881526020927f4552432d383030342054727573746c657373204167656e740000000000000000848301528251838101818110838211176200031757845260058152641051d1539560da1b85820152825190828211620003175760008054926001958685811c951680156200030c575b89861014620002f8578190601f95868111620002a5575b5089908683116001146200024157849262000235575b5050600019600383901b1c191690861b1781555b8151938411620002215784548581811c9116801562000216575b888210146200020257838111620001ba575b5086928411600114620001545783949596509262000148575b5050600019600383901b1c191690821b1781555b806007556008540160085551611b269081620003338239f35b0151905038806200011b565b9190601f1984169685845280842093905b888210620001a2575050838596971062000188575b505050811b0181556200012f565b015160001960f88460031b161c191690553880806200017a565b80878596829496860151815501950193019062000165565b8582528782208480870160051c8201928a8810620001f8575b0160051c019086905b828110620001ec57505062000102565b838155018690620001dc565b92508192620001d3565b634e487b7160e01b82526022600452602482fd5b90607f1690620000f0565b634e487b7160e01b81526041600452602490fd5b015190503880620000c2565b8480528a85208994509190601f198416865b8d8282106200028e575050841162000274575b505050811b018155620000d6565b015160001960f88460031b161c1916905538808062000266565b8385015186558c9790950194938401930162000253565b9091508380528984208680850160051c8201928c8610620002ee575b918a91869594930160051c01915b828110620002df575050620000ac565b8681558594508a9101620002cf565b92508192620002c1565b634e487b7160e01b83526022600452602483fd5b94607f169462000095565b634e487b7160e01b600052604160045260246000fd5b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c90816301ffc9a714610da15750806306fdde0314610cf9578063081812fc14610cda578063095ea7b314610b6c5780631aa3a00814610b3857806323b872dd14610b1357806342842e0e14610adf578063466648da1461090a5780636352211e146108da57806370a08231146108455780638ea42286146105b757806395d89b41146104e6578063a22cb46514610416578063b88d4fde14610388578063c505371214610349578063c87b56dd146102c6578063cb4799f2146101ed578063de99f157146101af578063e985e9c51461015d5763f2c298be1461010057600080fd5b3461015a57602036600319011261015a5782359067ffffffffffffffff821161015a575061014861013960209461014e93369101610ecd565b610141611580565b3691610f6b565b336116c3565b90600160075551908152f35b80fd5b5090346101ab57806003193601126101ab5760ff8160209361017d610e67565b610185610e82565b6001600160a01b0391821683526005875283832091168252855220549151911615158152f35b5080fd5b503461015a57602036600319011261015a57506101e460209235600052600260205260018060a01b0360406000205416151590565b90519015158152f35b5082903461015a578260031936011261015a5781359060243567ffffffffffffffff81116101ab576102229036908501610ecd565b60008481526002602052604090205491949093916001600160a01b03161561028c5761028861026f876102766102648989848a8a81526009602052209161160e565b82519384809261104e565b0383610f2d565b51918291602083526020830190610e27565b0390f35b606490602087519162461bcd60e51b835282015260146024820152731059d95b9d08191bd95cc81b9bdd08195e1a5cdd60621b6044820152fd5b50913461034557602036600319011261034557356000818152600260205260409020546102889361032093929091610308906001600160a01b03161515610fa2565b8152600660205261032782822083519485809261104e565b0384610f2d565b815161033281610efb565b5251918291602083526020830190610e27565b8280fd5b5090346101ab57816003193601126101ab57600854600019810192908311610375576020838351908152f35b634e487b7160e01b815260118452602490fd5b5082346101ab5760803660031901126101ab576103a3610e67565b906103ac610e82565b916044356064359367ffffffffffffffff85116104125736602386011215610412576103e761040a9486602461040f98369301359101610f6b565b926103fa6103f584336111fb565b611122565b6104058383836112c3565b6114f5565b6111d7565b80f35b8580fd5b50919034610345578060031936011261034557610431610e67565b90602435918215158093036104e2576001600160a01b0316923384146104a05750338452600560205280842083855260205280842060ff1981541660ff8416179055519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a380f35b6020606492519162461bcd60e51b8352820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152fd5b8480fd5b5090346101ab57816003193601126101ab5780519082600180549161050a83611014565b8086529282811690811561058f5750600114610533575b50505061027682610288940383610f2d565b94508085527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b828610610577575050506102768260206102889582010194610521565b8054602087870181019190915290950194810161055a565b61028897508693506020925061027694915060ff191682840152151560051b82010194610521565b5091903461034557806003193601126103455767ffffffffffffffff82358181116104e2576105e99036908501610ecd565b94909360249586359584871161084157366023880112156108415786830135918583116104e257888801978936918560051b0101116104e2576106329161014891610141611580565b9581610648575b60208787600160075551908152f35b835b8281106106575750610639565b61067761066e610668838686611a65565b80611a9d565b905015156115d6565b610682818484611a65565b61069160209182810190611a9d565b8a8852600983526106b28a89206106ac610668878a8a611a65565b9061160e565b9189821161082f5785838389938f958f97906106d16106d79254611014565b85611627565b8c908d601f84116001146107a2579261077194928192600080516020611ad1833981519152999a9b9592610797575b50508160011b916000199060031b1c19161790555b8b61075361072d6106688b8888611a65565b91909361074a8c610742610668828c8c611a65565b9a9099611a65565b90810190611a9d565b939092828b519384938437820190815203902097519485948561169c565b0390a360001981146107855760010161064a565b634e487b7160e01b8552601184528885fd5b013590503880610706565b91601f1984168584528a8420935b81811061080457509161077195939185600080516020611ad18339815191529a9b9c9694106107ea575b505050600190811b01905561071b565b0135600019600384901b60f8161c191690553880806107da565b9599509397509550935087600181928787013581550195019201938f9591938f97938b958d976107b0565b634e487b7160e01b8952604188528c89fd5b8380fd5b5082346101ab5760203660031901126101ab576001600160a01b03610868610e67565b169081156108855760208480858581526003845220549051908152f35b608490602085519162461bcd60e51b8352820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152fd5b503461015a57602036600319011261015a57506108f960209235610fee565b90516001600160a01b039091168152f35b5091346103455760603660031901126103455780359067ffffffffffffffff90602435828111610412576109419036908301610ecd565b929091604435828111610adb5761095b9036908301610ecd565b92909161096887336111fb565b15610aa7576109788615156115d6565b8689526020906009825261098f898b20888861160e565b928511610a9457506109ab846109a58454611014565b84611627565b8890601f8511600114610a195750918391600080516020611ad1833981519152969594610a08948b91610a0e575b508360011b906000198560031b1c19161790555b875185858237808681018b815203902097519485948561169c565b0390a380f35b9050820135386109d9565b90601f198516838b52828b20928b905b828210610a7c57505091859391600080516020611ad1833981519152989796610a08969410610a62575b5050600183811b0190556109ed565b830135600019600386901b60f8161c191690553880610a53565b80600185968294968a01358155019501930190610a29565b634e487b7160e01b8a5260419052602489fd5b606490602089519162461bcd60e51b8352820152600e60248201526d139bdd08185d5d1a1bdc9a5e995960921b6044820152fd5b8780fd5b5090346101ab5761040a61040f91610af636610e98565b91925192610b0384610efb565b8684526103fa6103f584336111fb565b503461015a5761040f610b2536610e98565b91610b336103f584336111fb565b6112c3565b5090346101ab57816003193601126101ab5761014e602092610b58611580565b825190610b6482610efb565b8152336116c3565b509134610345578160031936011261034557610b86610e67565b6024359290916001600160a01b0391908280610ba187610fee565b16941693808514610c8d57803314908115610c6e575b5015610c0657848652602052842080546001600160a01b03191683179055610bde83610fee565b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258480a480f35b6020608492519162461bcd60e51b8352820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152fd5b90508652600560205281862033875260205260ff828720541638610bb7565b506020608492519162461bcd60e51b8352820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152fd5b503461015a57602036600319011261015a57506108f9602092356110e4565b5090346101ab57816003193601126101ab57805190828054610d1a81611014565b8085529160019180831690811561058f5750600114610d455750505061027682610288940383610f2d565b80809650527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b828610610d89575050506102768260206102889582010194610521565b80546020878701810191909152909501948101610d6c565b90508334610345576020366003190112610345573563ffffffff60e01b81168091036103455760209250632483248360e11b8114908115610de4575b5015158152f35b6380ac58cd60e01b811491508115610e16575b8115610e05575b5083610ddd565b6301ffc9a760e01b14905083610dfe565b635b5e139f60e01b81149150610df7565b919082519283825260005b848110610e53575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610e32565b600435906001600160a01b0382168203610e7d57565b600080fd5b602435906001600160a01b0382168203610e7d57565b6060906003190112610e7d576001600160a01b03906004358281168103610e7d57916024359081168103610e7d579060443590565b9181601f84011215610e7d5782359167ffffffffffffffff8311610e7d5760208381860195010111610e7d57565b6020810190811067ffffffffffffffff821117610f1757604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff821117610f1757604052565b67ffffffffffffffff8111610f1757601f01601f191660200190565b929192610f7782610f4f565b91610f856040519384610f2d565b829481845281830111610e7d578281602093846000960137010152565b15610fa957565b60405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606490fd5b6000908152600260205260409020546001600160a01b0316611011811515610fa2565b90565b90600182811c92168015611044575b602083101461102e57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611023565b906000929180549161105f83611014565b9182825260019384811690816000146110c15750600114611081575b50505050565b90919394506000526020928360002092846000945b8386106110ad57505050500101903880808061107b565b805485870183015294019385908201611096565b9294505050602093945060ff191683830152151560051b0101903880808061107b565b600081815260026020526040902054611107906001600160a01b03161515610fa2565b6000908152600460205260409020546001600160a01b031690565b1561112957565b60405162461bcd60e51b815260206004820152602d60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201526c1c881bdc88185c1c1c9bdd9959609a1b6064820152608490fd5b60809060208152603260208201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b60608201520190565b156111de57565b60405162461bcd60e51b8152806111f760048201611184565b0390fd5b906001600160a01b03808061120f84610fee565b16931691838314938415611242575b50831561122c575b50505090565b611238919293506110e4565b1614388080611226565b909350600052600560205260406000208260005260205260ff60406000205416923861121e565b1561127057565b60405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608490fd5b906112eb916112d184610fee565b6001600160a01b0393918416928492909183168414611269565b1691821561138857816113089161130186610fee565b1614611269565b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60008481526004602052604081206bffffffffffffffffffffffff60a01b9081815416905583825260036020526040822060001981540190558482526040822060018154019055858252600260205284604083209182541617905580a4565b60405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b9192600092909190803b156114eb57611427946040518092630a85bd0160e11b9485835233600484015287602484015260448301526080606483015281878160209a8b966084830190610e27565b03926001600160a01b03165af18491816114ab575b5061149a575050503d600014611492573d61145681610f4f565b906114646040519283610f2d565b81528091833d92013e5b8051918261148f5760405162461bcd60e51b8152806111f760048201611184565b01fd5b50606061146e565b6001600160e01b0319161492509050565b9091508581813d83116114e4575b6114c38183610f2d565b810103126104e257516001600160e01b0319811681036104e257903861143c565b503d6114b9565b5050915050600190565b9293600093909291803b156115755794849161154f9660405180948193630a85bd0160e11b9788845233600485015260018060a01b0380921660248501526044840152608060648401528260209b8c976084830190610e27565b0393165af18491816114ab575061149a575050503d600014611492573d61145681610f4f565b505050915050600190565b600260075414611591576002600755565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b156115dd57565b60405162461bcd60e51b8152602060048201526009602482015268456d707479206b657960b81b6044820152606490fd5b6020919283604051948593843782019081520301902090565b90601f811161163557505050565b600091825260208220906020601f850160051c83019410611671575b601f0160051c01915b82811061166657505050565b81815560010161165a565b9092508290611651565b908060209392818452848401376000828201840152601f01601f1916010190565b92906116b590611011959360408652604086019161167b565b92602081850391015261167b565b9190916008549260019182850160085560408051936116e185610efb565b60008086526001600160a01b0384169590919086156119d6576000898152600260205260409020546117ab9161040a91611727906001600160a01b031615155b15611a19565b60008b81526002602052604090205461174a906001600160a01b03161515611721565b8885528a60209760038952878720868154019055818752600289528787208b6bffffffffffffffffffffffff60a01b825416179055818b887fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a46113d9565b84516117ef575b50505181815285927fca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a9282916117ea91830190610e27565b0390a3565b9290969594939161181587600052600260205260018060a01b0360406000205416151590565b1561197c57868852600682528088209383519067ffffffffffffffff82116119685788996118518361184b899a9b9c9954611014565b8a611627565b8490601f84116001146118e05792807fca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a98999381936117ea9796936118d5575b501b916000199060031b1c19161790555b7ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7838251888152a19294508193506117b2565b890151925038611891565b979290601f198216848a52868a20995b81811061195057509882916117ea969594937fca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a9a9b10611937575b5050811b0190556118a2565b88015160001960f88460031b161c19169055388061192b565b828901518b55998401998d99509187019187016118f0565b634e487b7160e01b8a52604160045260248afd5b60849250519062461bcd60e51b82526004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152fd5b6064845162461bcd60e51b815260206004820152602060248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152fd5b15611a2057565b60405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606490fd5b9190811015611a875760051b81013590603e1981360301821215610e7d570190565b634e487b7160e01b600052603260045260246000fd5b903590601e1981360301821215610e7d570180359067ffffffffffffffff8211610e7d57602001918136038313610e7d5756fe2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468ba26469706673582212203701f9b5943a9b44c256f49252b402075557cd86166e8eb9ed50f6d13812cc0364736f6c63430008130033",
}

// IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryMetaData.ABI instead.
var IdentityRegistryABI = IdentityRegistryMetaData.ABI

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityRegistryMetaData.Bin instead.
var IdentityRegistryBin = IdentityRegistryMetaData.Bin

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// AgentExists is a free data retrieval call binding the contract method 0xde99f157.
//
// Solidity: function agentExists(uint256 agentId) view returns(bool exists)
func (_IdentityRegistry *IdentityRegistryCaller) AgentExists(opts *bind.CallOpts, agentId *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "agentExists", agentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AgentExists is a free data retrieval call binding the contract method 0xde99f157.
//
// Solidity: function agentExists(uint256 agentId) view returns(bool exists)
func (_IdentityRegistry *IdentityRegistrySession) AgentExists(agentId *big.Int) (bool, error) {
	return _IdentityRegistry.Contract.AgentExists(&_IdentityRegistry.CallOpts, agentId)
}

// AgentExists is a free data retrieval call binding the contract method 0xde99f157.
//
// Solidity: function agentExists(uint256 agentId) view returns(bool exists)
func (_IdentityRegistry *IdentityRegistryCallerSession) AgentExists(agentId *big.Int) (bool, error) {
	return _IdentityRegistry.Contract.AgentExists(&_IdentityRegistry.CallOpts, agentId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdentityRegistry.Contract.BalanceOf(&_IdentityRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IdentityRegistry.Contract.BalanceOf(&_IdentityRegistry.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IdentityRegistry.Contract.GetApproved(&_IdentityRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IdentityRegistry.Contract.GetApproved(&_IdentityRegistry.CallOpts, tokenId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes value)
func (_IdentityRegistry *IdentityRegistryCaller) GetMetadata(opts *bind.CallOpts, agentId *big.Int, key string) ([]byte, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "getMetadata", agentId, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes value)
func (_IdentityRegistry *IdentityRegistrySession) GetMetadata(agentId *big.Int, key string) ([]byte, error) {
	return _IdentityRegistry.Contract.GetMetadata(&_IdentityRegistry.CallOpts, agentId, key)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string key) view returns(bytes value)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetMetadata(agentId *big.Int, key string) ([]byte, error) {
	return _IdentityRegistry.Contract.GetMetadata(&_IdentityRegistry.CallOpts, agentId, key)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsApprovedForAll(&_IdentityRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsApprovedForAll(&_IdentityRegistry.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdentityRegistry *IdentityRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdentityRegistry *IdentityRegistrySession) Name() (string, error) {
	return _IdentityRegistry.Contract.Name(&_IdentityRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IdentityRegistry *IdentityRegistryCallerSession) Name() (string, error) {
	return _IdentityRegistry.Contract.Name(&_IdentityRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityRegistry.Contract.SupportsInterface(&_IdentityRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityRegistry.Contract.SupportsInterface(&_IdentityRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdentityRegistry *IdentityRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdentityRegistry *IdentityRegistrySession) Symbol() (string, error) {
	return _IdentityRegistry.Contract.Symbol(&_IdentityRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IdentityRegistry *IdentityRegistryCallerSession) Symbol() (string, error) {
	return _IdentityRegistry.Contract.Symbol(&_IdentityRegistry.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IdentityRegistry *IdentityRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IdentityRegistry *IdentityRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _IdentityRegistry.Contract.TokenURI(&_IdentityRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IdentityRegistry *IdentityRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IdentityRegistry.Contract.TokenURI(&_IdentityRegistry.CallOpts, tokenId)
}

// TotalAgents is a free data retrieval call binding the contract method 0xc5053712.
//
// Solidity: function totalAgents() view returns(uint256 count)
func (_IdentityRegistry *IdentityRegistryCaller) TotalAgents(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "totalAgents")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAgents is a free data retrieval call binding the contract method 0xc5053712.
//
// Solidity: function totalAgents() view returns(uint256 count)
func (_IdentityRegistry *IdentityRegistrySession) TotalAgents() (*big.Int, error) {
	return _IdentityRegistry.Contract.TotalAgents(&_IdentityRegistry.CallOpts)
}

// TotalAgents is a free data retrieval call binding the contract method 0xc5053712.
//
// Solidity: function totalAgents() view returns(uint256 count)
func (_IdentityRegistry *IdentityRegistryCallerSession) TotalAgents() (*big.Int, error) {
	return _IdentityRegistry.Contract.TotalAgents(&_IdentityRegistry.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Approve(&_IdentityRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Approve(&_IdentityRegistry.TransactOpts, to, tokenId)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistrySession) Register() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactorSession) Register() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string tokenURI_, (string,bytes)[] metadata) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactor) Register0(opts *bind.TransactOpts, tokenURI_ string, metadata []IIdentityRegistryMetadataEntry) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "register0", tokenURI_, metadata)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string tokenURI_, (string,bytes)[] metadata) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistrySession) Register0(tokenURI_ string, metadata []IIdentityRegistryMetadataEntry) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register0(&_IdentityRegistry.TransactOpts, tokenURI_, metadata)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string tokenURI_, (string,bytes)[] metadata) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactorSession) Register0(tokenURI_ string, metadata []IIdentityRegistryMetadataEntry) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register0(&_IdentityRegistry.TransactOpts, tokenURI_, metadata)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string tokenURI_) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactor) Register1(opts *bind.TransactOpts, tokenURI_ string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "register1", tokenURI_)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string tokenURI_) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistrySession) Register1(tokenURI_ string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register1(&_IdentityRegistry.TransactOpts, tokenURI_)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string tokenURI_) returns(uint256 agentId)
func (_IdentityRegistry *IdentityRegistryTransactorSession) Register1(tokenURI_ string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register1(&_IdentityRegistry.TransactOpts, tokenURI_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SafeTransferFrom(&_IdentityRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SafeTransferFrom(&_IdentityRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IdentityRegistry *IdentityRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SafeTransferFrom0(&_IdentityRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SafeTransferFrom0(&_IdentityRegistry.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetApprovalForAll(&_IdentityRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetApprovalForAll(&_IdentityRegistry.TransactOpts, operator, approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetMetadata(opts *bind.TransactOpts, agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setMetadata", agentId, key, value)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetMetadata(agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetMetadata(&_IdentityRegistry.TransactOpts, agentId, key, value)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string key, bytes value) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetMetadata(agentId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetMetadata(&_IdentityRegistry.TransactOpts, agentId, key, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferFrom(&_IdentityRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferFrom(&_IdentityRegistry.TransactOpts, from, to, tokenId)
}

// IdentityRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IdentityRegistry contract.
type IdentityRegistryApprovalIterator struct {
	Event *IdentityRegistryApproval // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryApproval)
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
		it.Event = new(IdentityRegistryApproval)
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
func (it *IdentityRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryApproval represents a Approval event raised by the IdentityRegistry contract.
type IdentityRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IdentityRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryApprovalIterator{contract: _IdentityRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IdentityRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryApproval)
				if err := _IdentityRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseApproval(log types.Log) (*IdentityRegistryApproval, error) {
	event := new(IdentityRegistryApproval)
	if err := _IdentityRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IdentityRegistry contract.
type IdentityRegistryApprovalForAllIterator struct {
	Event *IdentityRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryApprovalForAll)
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
		it.Event = new(IdentityRegistryApprovalForAll)
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
func (it *IdentityRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryApprovalForAll represents a ApprovalForAll event raised by the IdentityRegistry contract.
type IdentityRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IdentityRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryApprovalForAllIterator{contract: _IdentityRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IdentityRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryApprovalForAll)
				if err := _IdentityRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseApprovalForAll(log types.Log) (*IdentityRegistryApprovalForAll, error) {
	event := new(IdentityRegistryApprovalForAll)
	if err := _IdentityRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the IdentityRegistry contract.
type IdentityRegistryBatchMetadataUpdateIterator struct {
	Event *IdentityRegistryBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryBatchMetadataUpdate)
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
		it.Event = new(IdentityRegistryBatchMetadataUpdate)
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
func (it *IdentityRegistryBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the IdentityRegistry contract.
type IdentityRegistryBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*IdentityRegistryBatchMetadataUpdateIterator, error) {

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryBatchMetadataUpdateIterator{contract: _IdentityRegistry.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *IdentityRegistryBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryBatchMetadataUpdate)
				if err := _IdentityRegistry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseBatchMetadataUpdate(log types.Log) (*IdentityRegistryBatchMetadataUpdate, error) {
	event := new(IdentityRegistryBatchMetadataUpdate)
	if err := _IdentityRegistry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryMetadataSetIterator is returned from FilterMetadataSet and is used to iterate over the raw logs and unpacked data for MetadataSet events raised by the IdentityRegistry contract.
type IdentityRegistryMetadataSetIterator struct {
	Event *IdentityRegistryMetadataSet // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryMetadataSet)
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
		it.Event = new(IdentityRegistryMetadataSet)
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
func (it *IdentityRegistryMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryMetadataSet represents a MetadataSet event raised by the IdentityRegistry contract.
type IdentityRegistryMetadataSet struct {
	AgentId    *big.Int
	IndexedKey common.Hash
	Key        string
	Value      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMetadataSet is a free log retrieval operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedKey, string key, bytes value)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterMetadataSet(opts *bind.FilterOpts, agentId []*big.Int, indexedKey []string) (*IdentityRegistryMetadataSetIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "MetadataSet", agentIdRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryMetadataSetIterator{contract: _IdentityRegistry.contract, event: "MetadataSet", logs: logs, sub: sub}, nil
}

// WatchMetadataSet is a free log subscription operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedKey, string key, bytes value)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchMetadataSet(opts *bind.WatchOpts, sink chan<- *IdentityRegistryMetadataSet, agentId []*big.Int, indexedKey []string) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "MetadataSet", agentIdRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryMetadataSet)
				if err := _IdentityRegistry.contract.UnpackLog(event, "MetadataSet", log); err != nil {
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

// ParseMetadataSet is a log parse operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedKey, string key, bytes value)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseMetadataSet(log types.Log) (*IdentityRegistryMetadataSet, error) {
	event := new(IdentityRegistryMetadataSet)
	if err := _IdentityRegistry.contract.UnpackLog(event, "MetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the IdentityRegistry contract.
type IdentityRegistryMetadataUpdateIterator struct {
	Event *IdentityRegistryMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryMetadataUpdate)
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
		it.Event = new(IdentityRegistryMetadataUpdate)
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
func (it *IdentityRegistryMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryMetadataUpdate represents a MetadataUpdate event raised by the IdentityRegistry contract.
type IdentityRegistryMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*IdentityRegistryMetadataUpdateIterator, error) {

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryMetadataUpdateIterator{contract: _IdentityRegistry.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *IdentityRegistryMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryMetadataUpdate)
				if err := _IdentityRegistry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseMetadataUpdate(log types.Log) (*IdentityRegistryMetadataUpdate, error) {
	event := new(IdentityRegistryMetadataUpdate)
	if err := _IdentityRegistry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IdentityRegistry contract.
type IdentityRegistryRegisteredIterator struct {
	Event *IdentityRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRegistered)
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
		it.Event = new(IdentityRegistryRegistered)
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
func (it *IdentityRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRegistered represents a Registered event raised by the IdentityRegistry contract.
type IdentityRegistryRegistered struct {
	AgentId  *big.Int
	TokenURI string
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string tokenURI, address indexed owner)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, agentId []*big.Int, owner []common.Address) (*IdentityRegistryRegisteredIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "Registered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRegisteredIterator{contract: _IdentityRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string tokenURI, address indexed owner)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRegistered, agentId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "Registered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRegistered)
				if err := _IdentityRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string tokenURI, address indexed owner)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseRegistered(log types.Log) (*IdentityRegistryRegistered, error) {
	event := new(IdentityRegistryRegistered)
	if err := _IdentityRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IdentityRegistry contract.
type IdentityRegistryTransferIterator struct {
	Event *IdentityRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryTransfer)
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
		it.Event = new(IdentityRegistryTransfer)
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
func (it *IdentityRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryTransfer represents a Transfer event raised by the IdentityRegistry contract.
type IdentityRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IdentityRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransferIterator{contract: _IdentityRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdentityRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryTransfer)
				if err := _IdentityRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseTransfer(log types.Log) (*IdentityRegistryTransfer, error) {
	event := new(IdentityRegistryTransfer)
	if err := _IdentityRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
