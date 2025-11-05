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

// ReputationRegistryMetaData contains all meta data concerning the ReputationRegistry contract.
var ReputationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appendResponse\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"responseUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClients\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"clientList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIdentityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastIndex\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"lastIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResponseCount\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"responders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSummary\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tag1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"averageScore\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"giveFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fileuri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"filehash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"feedbackAuth\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readAllFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tag1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeRevoked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"clients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"scores\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"tag1s\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"tag2s\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"revokedStatuses\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"score\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tag2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isRevoked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeedbackRevoked\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewFeedback\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"tag1\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tag2\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"fileuri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"filehash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResponseAppended\",\"inputs\":[{\"name\":\"agentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"clientAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feedbackIndex\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"responder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"responseUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"responseHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60a0346100d957601f61239a38819003918201601f19168301916001600160401b038311848410176100de578084926020946040528339810103126100d957516001600160a01b038116908190036100d9578015610094576080526040516122a590816100f5823960805181818160dd015281816111060152818161119a015281816112530152818161151d01526116160152f35b60405162461bcd60e51b815260206004820152601860248201527f496e76616c6964207265676973747279206164647265737300000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063134e18f4146100c7578063155e5bbd146100c2578063232b0810146100bd57806331259cff146100b857806342dd519c146100b35780634ab3ca99146100ae5780636e04cacd146100a957806380b87594146100a4578063bc4d861b1461009f578063c2349ab21461009a5763f2d817591461009557600080fd5b610bae565b610a75565b6100c7565b610941565b6107ee565b6106c2565b610643565b6103f6565b6102b8565b610222565b3461010c57600036600319011261010c576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b600080fd5b9181601f8401121561010c578235916001600160401b03831161010c576020838186019501011161010c57565b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761016f57604052565b61013e565b61010081019081106001600160401b0382111761016f57604052565b90601f801991011681019081106001600160401b0382111761016f57604052565b604051906101be82610154565b565b6001600160401b03811161016f57601f01601f191660200190565b81601f8201121561010c578035906101f2826101c0565b926102006040519485610190565b8284526020838301011161010c57816000926020809301838601378301015290565b3461010c5760e036600319011261010c5760243560ff8116810361010c576001600160401b039060843582811161010c57610261903690600401610111565b60c49291923593841161010c5761027f6102949436906004016101db565b9260a4359260643590604435906004356110c9565b005b6001600160a01b0381160361010c57565b6001600160401b0381160361010c57565b3461010c57606036600319011261010c5761034660043561032e6024356102de81610296565b604435926102eb846102a7565b61030a6001600160401b038086168015159182610394575b50506118ce565b600052600060205260406000209060018060a01b0316600052602052604060002090565b906001600160401b0316600052602052604060002090565b60ff8154166103906001830154926103676003600283015492015460ff1690565b6040805160ff909516855260208501959095529383015291151560608201529081906080820190565b0390f35b60008581526001602090815260408083206001600160a01b038a1684529091529020919250905b541610153880610303565b9181601f8401121561010c578235916001600160401b03831161010c576020808501948460051b01011161010c57565b3461010c57608036600319011261010c576004356001600160401b0360243581811161010c5761042a9036906004016103c6565b6044359391606435919081156105cd57610445913691611a2e565b91935b60009283928015159183151590855b8151811015610588576104b76104aa61047a8c6000526001602052604060002090565b6104946104878587611a95565b516001600160a01b031690565b60018060a01b0316600052602052604060002090565b546001600160401b031690565b89600191165b808b831611156104d75750506104d290611a86565b610457565b9097986105058961032e859e956104946104876104fe896000526000602052604060002090565b928a611a95565b600381015460ff1661056057858061057a575b61056057878061056c575b6105605761054b6105579261054561053f610551945460ff1690565b60ff1690565b90611aa9565b9a611980565b98611980565b909a919a6104bd565b50989761055790611980565b508660028201541415610523565b508860018201541415610518565b8689811689600082156105c457506105a39161053f91611ab6565b905b604080516001600160401b0392909216825260ff929092166020820152f35b915050906105a5565b50506105eb6105e6836000526002602052604060002090565b6119bf565b9193610448565b90815180825260208080930193019160005b828110610612575050505090565b83516001600160a01b031685529381019392810192600101610604565b9060206106409281815201906105f2565b90565b3461010c5760208060031936011261010c5760043560005260028152604060002090604051908181845491828152019360005281600020916000905b8282106106a2576103908561069681890382610190565b6040519182918261062f565b83546001600160a01b03168652948501946001938401939091019061067f565b3461010c57604036600319011261010c576024356004356106e2826102a7565b61079060036107826001600160401b039461070a8682169687151590816107b9575b506118ce565b61075361074e61074a856107428561032e61072f8c6000526000602052604060002090565b3360009081526020919091526040902090565b015460ff1690565b1590565b61190a565b61032e3361076b876000526000602052604060002090565b9060018060a01b0316600052602052604060002090565b01805460ff19166001179055565b33907f25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d600080a4005b600088815260016020908152604080832033845290915290206107e5925054166001600160401b031690565b87111538610704565b3461010c57608036600319011261010c5760243561080b81610296565b604435610817816102a7565b6064356001600160401b039182821161010c5760209361083e6108499336906004016103c6565b929091600435611dd3565b60405191168152f35b8015150361010c57565b90815180825260208080930193019160005b82811061087c575050505090565b83518552938101939281019260010161086e565b929391906108a69060a0855260a08501906105f2565b936020948481038686015285808451928381520193019060005b81811061092a57505050816108e091856108ee959403604087015261085c565b90838203606085015261085c565b90608081830391015281808451928381520193019160005b828110610914575050505090565b8351151585529381019392810192600101610906565b825160ff16855293870193918701916001016108c0565b3461010c5760a036600319011261010c576004356024356001600160401b03811161010c576109749036906004016103c6565b909160443592606435906084359361098b85610852565b60008115610a035750916109f76109ab6109ef9361039096953691611a2e565b935b6109ba87848a8885611b08565b6109c381611ad6565b9889986109cf83611ad6565b9a8b916109db85611ad6565b9788946109e787611ad6565b9a8b97611ad6565b9b8c98611c12565b60405195869586610890565b809250849150526020906002825260408120906040518093808454928381520193835280832092905b828210610a555750505050916109f782610a4f6109ef9461039097960382610190565b936109ad565b83546001600160a01b031685529384019360019384019390910190610a2c565b3461010c5760a036600319011261010c57602435600435610a9582610296565b60443590610aa2826102a7565b6064356001600160401b039283821161010c57610afd610ae77fb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4933690600401610111565b9190958084168015159182610b835750506118ce565b610b08811515611948565b610b5d610b2a61072f8461032e8a61076b8a6000526004602052604060002090565b610b43610b3e82546001600160401b031690565b611980565b6001600160401b03166001600160401b0319825416179055565b60405133966001600160a01b03169590928392610b7e926084359285611998565b0390a4005b60008881526001602090815260408083206001600160a01b038e1684529091529020919250906103bb565b3461010c57604036600319011261010c5760206001600160401b03610bfe602435610bd881610296565b6004356000526001845260406000209060018060a01b0316600052602052604060002090565b5416604051908152f35b15610c0f57565b60405162461bcd60e51b8152602060048201526013602482015272053636f7265206d75737420626520302d31303606c1b6044820152606490fd5b9081602091031261010c575161064081610852565b6040513d6000823e3d90fd5b15610c7257565b60405162461bcd60e51b81526020600482015260146024820152731059d95b9d08191bd95cc81b9bdd08195e1a5cdd60621b6044820152606490fd5b15610cb557565b60405162461bcd60e51b815260206004820152601060248201526f082cecadce892c840dad2e6dac2e8c6d60831b6044820152606490fd5b15610cf457565b60405162461bcd60e51b8152602060048201526016602482015275086d8d2cadce882c8c8e4cae6e640dad2e6dac2e8c6d60531b6044820152606490fd5b15610d3957565b60405162461bcd60e51b815260206004820152601060248201526f086d0c2d2dc92c840dad2e6dac2e8c6d60831b6044820152606490fd5b15610d7857565b60405162461bcd60e51b81526020600482015260116024820152700a4caced2e6e8e4f240dad2e6dac2e8c6d607b1b6044820152606490fd5b15610db857565b60405162461bcd60e51b8152602060048201526015602482015274105d5d1a1bdc9a5e985d1a5bdb88195e1c1a5c9959605a1b6044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b9060016001600160401b0380931601918211610e2357565b610df5565b9190916001600160401b0380809416911601918211610e2357565b15610e4a57565b60405162461bcd60e51b8152602060048201526014602482015273125b99195e081b1a5b5a5d08195e18d95959195960621b6044820152606490fd5b9081602091031261010c575161064081610296565b15610ea257565b60405162461bcd60e51b815260206004820152601960248201527f53656c662d666565646261636b206e6f7420616c6c6f776564000000000000006044820152606490fd5b15610eee57565b60405162461bcd60e51b815260206004820152600e60248201526d24b73b30b634b21039b4b3b732b960911b6044820152606490fd5b15610f2b57565b60405162461bcd60e51b815260206004820152601860248201527f496e76616c696420617574682064617461206c656e67746800000000000000006044820152606490fd5b60405190610f7d82610154565b604182526060366020840137565b6040519061010082018281106001600160401b0382111761016f5760405260e0808352366020840137565b60051115610fc057565b634e487b7160e01b600052602160045260246000fd5b15610fdd57565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b8054906801000000000000000082101561016f5760018201808255821015611078576000526020600020019060018060a01b03166bffffffffffffffffffffffff60a01b825416179055565b611016565b908060209392818452848401376000828201840152601f01601f1916010190565b9594936110c49260ff60609593168852602088015260806040880152608087019161107d565b930152565b9496939291969590956110e2606460ff89161115610c08565b60405163de99f15760e01b8152600481018790526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa80156115df57611140916000916116a5575b50610c6b565b611149856120cf565b9461115687875114610cae565b6020860151611178906001600160a01b03166001600160a01b03163314610ced565b61118760808701514614610d32565b60a08601516111c5906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169181161614610d71565b6111d460608701514210610db1565b6111f76111f26104aa3361076b8b6000526001602052604060002090565b610e0b565b9061122f61121e61121260408a01516001600160401b031690565b6001600160401b031690565b6001600160401b0384161115610e43565b6040516331a9108f60e11b8152600481018990526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115df57600091611686575b506001600160a01b03811661129b33821415610e9b565b60c08901516001600160a01b03169081149182156115e4575b505080156114f4575b9161136a7f54b3254e4cc01969e70376d20390cc2f6af90539d9adaa78a53ebcda17f7815498611418936112f5611450989796610ee7565b61130461012182511015610f24565b6101008101519061014061012082015191015160001a611323846121a9565b9161132c610f70565b9360208501526040840152606083015361134682826116c4565b61135281959295610fb6565b1593846114be575b50831561149f575b505050610fd6565b6113e56113756101b1565b60ff8c1681528c6020820152876040820152600060608201526113ab8361032e8d61076b33916000526000602052604060002090565b60ff606060038285511693831994858254161781556020860151600182015560408601516002820155019301511515918354169116179055565b6113fd3361076b8b6000526001602052604060002090565b906001600160401b03166001600160401b0319825416179055565b61143d61074a6114363361076b8b6000526003602052604060002090565b5460ff1690565b611455575b60405194859433998661109e565b0390a4565b6114723361146d896000526002602052604060002090565b61102c565b61149a61148d3361076b8a6000526003602052604060002090565b805460ff19166001179055565b611442565b60c001516114b693506001600160a01b031661177e565b388080611362565b60c0820151919450906114e1906001600160a01b03165b6001600160a01b031690565b6001600160a01b0390911614923861135a565b5060405163020604bf60e21b815260048101899052909392916020826024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa80156115df577f54b3254e4cc01969e70376d20390cc2f6af90539d9adaa78a53ebcda17f7815498611450966112f561136a93611418966000916115b0575b5060c0840151611596906001600160a01b03166114d5565b9060018060a01b03161494969798505093509850506112bd565b6115d2915060203d6020116115d8575b6115ca8183610190565b810190610e86565b3861157e565b503d6115c0565b610c5f565b60405163e985e9c560e01b81526001600160a01b039182166004820152911660248201529050602081806044810103817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156115df57600091611657575b5038806112b4565b611679915060203d60201161167f575b6116718183610190565b810190610c4a565b3861164f565b503d611667565b61169f915060203d6020116115d8576115ca8183610190565b38611284565b6116be915060203d60201161167f576116718183610190565b3861113a565b9060418151146000146116f2576116ee916020820151906060604084015193015160001a906116fc565b9091565b5050600090600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116117725791608094939160ff602094604051948552168484015260408301526060820152600093849182805260015afa156115df5781516001600160a01b0381161561176c579190565b50600190565b50505050600090600390565b919061178a82826116c4565b6005819592951015610fc057159384611879575b5083156117ac575b50505090565b909192506040519260209384810191630b135d3f60e11b948584526024830152604060448301528051908160648401528660005b83811061186457505050918161181660848286600083819a9899829a010152601f80199101168101036064810184520182610190565b51915afa9161182361188f565b9083611857575b8361183b575b5050503880806117a6565b61184e92935080825183010191016118bf565b14388080611830565b809350815110159261182a565b818184010151608482870101520187906117e0565b6001600160a01b0382811691161493503861179e565b3d156118ba573d906118a0826101c0565b916118ae6040519384610190565b82523d6000602084013e565b606090565b9081602091031261010c575190565b156118d557565b60405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c840d2dcc8caf609b1b6044820152606490fd5b1561191157565b60405162461bcd60e51b815260206004820152600f60248201526e105b1c9958591e481c995d9bdad959608a1b6044820152606490fd5b1561194f57565b60405162461bcd60e51b8152602060048201526009602482015268456d7074792055524960b81b6044820152606490fd5b6001600160401b03809116908114610e235760010190565b949392916040926001600160401b036110c49316875260606020880152606087019161107d565b9060405191828154918282526020928383019160005283600020936000905b8282106119f4575050506101be92500383610190565b85546001600160a01b0316845260019586019588955093810193909101906119de565b6001600160401b03811161016f5760051b60200190565b9291611a3982611a17565b91611a476040519384610190565b829481845260208094019160051b810192831161010c57905b828210611a6d5750505050565b8380918335611a7b81610296565b815201910190611a60565b6000198114610e235760010190565b80518210156110785760209160051b010190565b91908201809211610e2357565b8115611ac0570490565b634e487b7160e01b600052601260045260246000fd5b90611ae082611a17565b611aed6040519182610190565b8281528092611afe601f1991611a17565b0190602036910137565b939192909260009360005b8151811015611c0857611b366104aa61047a896000526001602052604060002090565b6001906001600160401b038091165b808284161115611b6057505050611b5b90611a86565b611b13565b909197611b8b8961032e611b7e8d6000526000602052604060002090565b610494610487898b611a95565b881580611bfa575b611be15786151580611bec575b611be157878015159182611bd2575b5050611bc857610551611bc191611a86565b9190611b45565b97611bc190611980565b60020154141590508738611baf565b5097611bc190611980565b508660018201541415611ba0565b50600381015460ff16611b93565b5050505050919050565b9894919695939092600097885b8551811015611dab57611c4f6104aa611c428e6000526001602052604060002090565b610494610487858b611a95565b6001906001600160401b039081168d8f5b828486161115611c7d575050505050611c7890611a86565b611c1f565b849e92939461032e8c6104946104878a611ca4611caa976000526000602052604060002090565b93611a95565b90881580611d9d575b611d755786151580611d8f575b611d755787151580611d81575b611d755791611d618d8f838f8f8f97611d6c99611d5a95611d328f95611d1b611d669d611d0c8a611d06610487611d069c60039c611a95565b92611a95565b6001600160a01b039091169052565b611d2a88611d06875460ff1690565b9060ff169052565b611d4186600185015492611a95565b52611d5185600284015492611a95565b52015460ff1690565b9015159052565b611a86565b9c611980565b91908d8f611c60565b50509b611d6c90611980565b508760028301541415611ccd565b508660018301541415611cc0565b50600382015460ff16611cb3565b505050505050505050505050565b91908110156110785760051b0190565b3561064081610296565b929094936000926000968215611fda576001600160a01b038116611ed4575050611e0a6105e6856000526002602052604060002090565b865b8151811015611ecb57611e2f6104aa61047a886000526001602052604060002090565b6001600160401b0390811660015b818382161115611e5857505050611e5390611a86565b611e0c565b8a5b89878210611e72575050611e6d90611980565b611e3d565b8199611eba6104aa8b610494611eb5611ec6978e611eaf8f8f8d9161049461048761032e93611ec09f611ca4906000526004602052604060002090565b94611db9565b611dc9565b90610e28565b98611a86565b611e5a565b50505050925050565b93966001600160401b03959491868116611f775750611f046104aa8361076b886000526001602052604060002090565b861660015b818882161115611f1d575050505050505050565b825b858110611f355750611f3090611980565b611f09565b99611f6c611f7291611eba6104aa8a8f611eb5610494918f611eaf8f918f61032e9061076b8f936000526004602052604060002090565b9a611a86565b611f1f565b919550945b828610611f8b57505050505050565b909192939496611fca611fd091611eba6104aa8b610494611eb5611fc28a61032e8f61076b8d916000526004602052604060002090565b928b8d611db9565b97611a86565b9493929190611f7c565b505050505050565b6040519060e082018281106001600160401b0382111761016f576040528160c06000918281528260208201528260408201528260608201528260808201528260a08201520152565b1561203157565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c69642061757468206461746160781b6044820152606490fd5b908151811015611078570160200190565b908160e091031261010c57805191602082015161209781610296565b9160408101516120a6816102a7565b9160608201519160808101519160c060a08301516120c381610296565b92015161064081610296565b906120d8611fe2565b916120e86101218251101561202a565b6120f0610f8b565b906000805b60e081106121735750505061216f60c061214461211e8460208061215e9751830101910161207b565b6001600160a01b03908116988d0198909852871660a08c01529397929492939192909190565b60808a015260608901526001600160401b03166040880152565b166001600160a01b03166020850152565b8252565b806121926121846121a4938661206a565b516001600160f81b03191690565b831a61219e828761206a565b53611a86565b6120f5565b80519060018060a01b0380602083015116916001600160401b0360408201511691606082015160808301519160c08160a0860151169401511693604051956020870197885260408701526060860152608085015260a084015260c083015260e082015260e0815261221981610174565b5190206040516122698161225b602082019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b03601f198101835282610190565b5190209056fea2646970667358221220cc1a7dff5f6dbd693329fa02fdc5938a2ea94f1b9db5a2e4e94da30d5908d83364736f6c63430008130033",
}

// ReputationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationRegistryMetaData.ABI instead.
var ReputationRegistryABI = ReputationRegistryMetaData.ABI

// ReputationRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReputationRegistryMetaData.Bin instead.
var ReputationRegistryBin = ReputationRegistryMetaData.Bin

// DeployReputationRegistry deploys a new Ethereum contract, binding an instance of ReputationRegistry to it.
func DeployReputationRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _identityRegistry common.Address) (common.Address, *types.Transaction, *ReputationRegistry, error) {
	parsed, err := ReputationRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReputationRegistryBin), backend, _identityRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReputationRegistry{ReputationRegistryCaller: ReputationRegistryCaller{contract: contract}, ReputationRegistryTransactor: ReputationRegistryTransactor{contract: contract}, ReputationRegistryFilterer: ReputationRegistryFilterer{contract: contract}}, nil
}

// ReputationRegistry is an auto generated Go binding around an Ethereum contract.
type ReputationRegistry struct {
	ReputationRegistryCaller     // Read-only binding to the contract
	ReputationRegistryTransactor // Write-only binding to the contract
	ReputationRegistryFilterer   // Log filterer for contract events
}

// ReputationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationRegistrySession struct {
	Contract     *ReputationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReputationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationRegistryCallerSession struct {
	Contract *ReputationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ReputationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationRegistryTransactorSession struct {
	Contract     *ReputationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ReputationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRegistryRaw struct {
	Contract *ReputationRegistry // Generic contract binding to access the raw methods on
}

// ReputationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationRegistryCallerRaw struct {
	Contract *ReputationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationRegistryTransactorRaw struct {
	Contract *ReputationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationRegistry creates a new instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistry(address common.Address, backend bind.ContractBackend) (*ReputationRegistry, error) {
	contract, err := bindReputationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistry{ReputationRegistryCaller: ReputationRegistryCaller{contract: contract}, ReputationRegistryTransactor: ReputationRegistryTransactor{contract: contract}, ReputationRegistryFilterer: ReputationRegistryFilterer{contract: contract}}, nil
}

// NewReputationRegistryCaller creates a new read-only instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryCaller(address common.Address, caller bind.ContractCaller) (*ReputationRegistryCaller, error) {
	contract, err := bindReputationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryCaller{contract: contract}, nil
}

// NewReputationRegistryTransactor creates a new write-only instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationRegistryTransactor, error) {
	contract, err := bindReputationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryTransactor{contract: contract}, nil
}

// NewReputationRegistryFilterer creates a new log filterer instance of ReputationRegistry, bound to a specific deployed contract.
func NewReputationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationRegistryFilterer, error) {
	contract, err := bindReputationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryFilterer{contract: contract}, nil
}

// bindReputationRegistry binds a generic wrapper to an already deployed contract.
func bindReputationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationRegistry *ReputationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationRegistry.Contract.ReputationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationRegistry *ReputationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.ReputationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationRegistry *ReputationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.ReputationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationRegistry *ReputationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationRegistry *ReputationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationRegistry *ReputationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[] clientList)
func (_ReputationRegistry *ReputationRegistryCaller) GetClients(opts *bind.CallOpts, agentId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getClients", agentId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[] clientList)
func (_ReputationRegistry *ReputationRegistrySession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _ReputationRegistry.Contract.GetClients(&_ReputationRegistry.CallOpts, agentId)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[] clientList)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _ReputationRegistry.Contract.GetClients(&_ReputationRegistry.CallOpts, agentId)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ReputationRegistry *ReputationRegistryCaller) GetIdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getIdentityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ReputationRegistry *ReputationRegistrySession) GetIdentityRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.GetIdentityRegistry(&_ReputationRegistry.CallOpts)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address registry)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetIdentityRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.GetIdentityRegistry(&_ReputationRegistry.CallOpts)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64 lastIndex)
func (_ReputationRegistry *ReputationRegistryCaller) GetLastIndex(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address) (uint64, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getLastIndex", agentId, clientAddress)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64 lastIndex)
func (_ReputationRegistry *ReputationRegistrySession) GetLastIndex(agentId *big.Int, clientAddress common.Address) (uint64, error) {
	return _ReputationRegistry.Contract.GetLastIndex(&_ReputationRegistry.CallOpts, agentId, clientAddress)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64 lastIndex)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetLastIndex(agentId *big.Int, clientAddress common.Address) (uint64, error) {
	return _ReputationRegistry.Contract.GetLastIndex(&_ReputationRegistry.CallOpts, agentId, clientAddress)
}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_ReputationRegistry *ReputationRegistryCaller) GetResponseCount(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getResponseCount", agentId, clientAddress, feedbackIndex, responders)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_ReputationRegistry *ReputationRegistrySession) GetResponseCount(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	return _ReputationRegistry.Contract.GetResponseCount(&_ReputationRegistry.CallOpts, agentId, clientAddress, feedbackIndex, responders)
}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetResponseCount(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	return _ReputationRegistry.Contract.GetResponseCount(&_ReputationRegistry.CallOpts, agentId, clientAddress, feedbackIndex, responders)
}

// GetSummary is a free data retrieval call binding the contract method 0x31259cff.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2) view returns(uint64 count, uint8 averageScore)
func (_ReputationRegistry *ReputationRegistryCaller) GetSummary(opts *bind.CallOpts, agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte) (struct {
	Count        uint64
	AverageScore uint8
}, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "getSummary", agentId, clientAddresses, tag1, tag2)

	outstruct := new(struct {
		Count        uint64
		AverageScore uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AverageScore = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSummary is a free data retrieval call binding the contract method 0x31259cff.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2) view returns(uint64 count, uint8 averageScore)
func (_ReputationRegistry *ReputationRegistrySession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte) (struct {
	Count        uint64
	AverageScore uint8
}, error) {
	return _ReputationRegistry.Contract.GetSummary(&_ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// GetSummary is a free data retrieval call binding the contract method 0x31259cff.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2) view returns(uint64 count, uint8 averageScore)
func (_ReputationRegistry *ReputationRegistryCallerSession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte) (struct {
	Count        uint64
	AverageScore uint8
}, error) {
	return _ReputationRegistry.Contract.GetSummary(&_ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistryCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistrySession) IdentityRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.IdentityRegistry(&_ReputationRegistry.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_ReputationRegistry *ReputationRegistryCallerSession) IdentityRegistry() (common.Address, error) {
	return _ReputationRegistry.Contract.IdentityRegistry(&_ReputationRegistry.CallOpts)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x80b87594.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2, bool includeRevoked) view returns(address[] clients, uint8[] scores, bytes32[] tag1s, bytes32[] tag2s, bool[] revokedStatuses)
func (_ReputationRegistry *ReputationRegistryCaller) ReadAllFeedback(opts *bind.CallOpts, agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte, includeRevoked bool) (struct {
	Clients         []common.Address
	Scores          []uint8
	Tag1s           [][32]byte
	Tag2s           [][32]byte
	RevokedStatuses []bool
}, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "readAllFeedback", agentId, clientAddresses, tag1, tag2, includeRevoked)

	outstruct := new(struct {
		Clients         []common.Address
		Scores          []uint8
		Tag1s           [][32]byte
		Tag2s           [][32]byte
		RevokedStatuses []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Clients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Scores = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)
	outstruct.Tag1s = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	outstruct.Tag2s = *abi.ConvertType(out[3], new([][32]byte)).(*[][32]byte)
	outstruct.RevokedStatuses = *abi.ConvertType(out[4], new([]bool)).(*[]bool)

	return *outstruct, err

}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x80b87594.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2, bool includeRevoked) view returns(address[] clients, uint8[] scores, bytes32[] tag1s, bytes32[] tag2s, bool[] revokedStatuses)
func (_ReputationRegistry *ReputationRegistrySession) ReadAllFeedback(agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte, includeRevoked bool) (struct {
	Clients         []common.Address
	Scores          []uint8
	Tag1s           [][32]byte
	Tag2s           [][32]byte
	RevokedStatuses []bool
}, error) {
	return _ReputationRegistry.Contract.ReadAllFeedback(&_ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2, includeRevoked)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0x80b87594.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, bytes32 tag1, bytes32 tag2, bool includeRevoked) view returns(address[] clients, uint8[] scores, bytes32[] tag1s, bytes32[] tag2s, bool[] revokedStatuses)
func (_ReputationRegistry *ReputationRegistryCallerSession) ReadAllFeedback(agentId *big.Int, clientAddresses []common.Address, tag1 [32]byte, tag2 [32]byte, includeRevoked bool) (struct {
	Clients         []common.Address
	Scores          []uint8
	Tag1s           [][32]byte
	Tag2s           [][32]byte
	RevokedStatuses []bool
}, error) {
	return _ReputationRegistry.Contract.ReadAllFeedback(&_ReputationRegistry.CallOpts, agentId, clientAddresses, tag1, tag2, includeRevoked)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 index) view returns(uint8 score, bytes32 tag1, bytes32 tag2, bool isRevoked)
func (_ReputationRegistry *ReputationRegistryCaller) ReadFeedback(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address, index uint64) (struct {
	Score     uint8
	Tag1      [32]byte
	Tag2      [32]byte
	IsRevoked bool
}, error) {
	var out []interface{}
	err := _ReputationRegistry.contract.Call(opts, &out, "readFeedback", agentId, clientAddress, index)

	outstruct := new(struct {
		Score     uint8
		Tag1      [32]byte
		Tag2      [32]byte
		IsRevoked bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Tag1 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Tag2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IsRevoked = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 index) view returns(uint8 score, bytes32 tag1, bytes32 tag2, bool isRevoked)
func (_ReputationRegistry *ReputationRegistrySession) ReadFeedback(agentId *big.Int, clientAddress common.Address, index uint64) (struct {
	Score     uint8
	Tag1      [32]byte
	Tag2      [32]byte
	IsRevoked bool
}, error) {
	return _ReputationRegistry.Contract.ReadFeedback(&_ReputationRegistry.CallOpts, agentId, clientAddress, index)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 index) view returns(uint8 score, bytes32 tag1, bytes32 tag2, bool isRevoked)
func (_ReputationRegistry *ReputationRegistryCallerSession) ReadFeedback(agentId *big.Int, clientAddress common.Address, index uint64) (struct {
	Score     uint8
	Tag1      [32]byte
	Tag2      [32]byte
	IsRevoked bool
}, error) {
	return _ReputationRegistry.Contract.ReadFeedback(&_ReputationRegistry.CallOpts, agentId, clientAddress, index)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseUri, bytes32 responseHash) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) AppendResponse(opts *bind.TransactOpts, agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseUri string, responseHash [32]byte) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "appendResponse", agentId, clientAddress, feedbackIndex, responseUri, responseHash)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseUri, bytes32 responseHash) returns()
func (_ReputationRegistry *ReputationRegistrySession) AppendResponse(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseUri string, responseHash [32]byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.AppendResponse(&_ReputationRegistry.TransactOpts, agentId, clientAddress, feedbackIndex, responseUri, responseHash)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseUri, bytes32 responseHash) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) AppendResponse(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseUri string, responseHash [32]byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.AppendResponse(&_ReputationRegistry.TransactOpts, agentId, clientAddress, feedbackIndex, responseUri, responseHash)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x155e5bbd.
//
// Solidity: function giveFeedback(uint256 agentId, uint8 score, bytes32 tag1, bytes32 tag2, string fileuri, bytes32 filehash, bytes feedbackAuth) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) GiveFeedback(opts *bind.TransactOpts, agentId *big.Int, score uint8, tag1 [32]byte, tag2 [32]byte, fileuri string, filehash [32]byte, feedbackAuth []byte) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "giveFeedback", agentId, score, tag1, tag2, fileuri, filehash, feedbackAuth)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x155e5bbd.
//
// Solidity: function giveFeedback(uint256 agentId, uint8 score, bytes32 tag1, bytes32 tag2, string fileuri, bytes32 filehash, bytes feedbackAuth) returns()
func (_ReputationRegistry *ReputationRegistrySession) GiveFeedback(agentId *big.Int, score uint8, tag1 [32]byte, tag2 [32]byte, fileuri string, filehash [32]byte, feedbackAuth []byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.GiveFeedback(&_ReputationRegistry.TransactOpts, agentId, score, tag1, tag2, fileuri, filehash, feedbackAuth)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x155e5bbd.
//
// Solidity: function giveFeedback(uint256 agentId, uint8 score, bytes32 tag1, bytes32 tag2, string fileuri, bytes32 filehash, bytes feedbackAuth) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) GiveFeedback(agentId *big.Int, score uint8, tag1 [32]byte, tag2 [32]byte, fileuri string, filehash [32]byte, feedbackAuth []byte) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.GiveFeedback(&_ReputationRegistry.TransactOpts, agentId, score, tag1, tag2, fileuri, filehash, feedbackAuth)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ReputationRegistry *ReputationRegistryTransactor) RevokeFeedback(opts *bind.TransactOpts, agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ReputationRegistry.contract.Transact(opts, "revokeFeedback", agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ReputationRegistry *ReputationRegistrySession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RevokeFeedback(&_ReputationRegistry.TransactOpts, agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_ReputationRegistry *ReputationRegistryTransactorSession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _ReputationRegistry.Contract.RevokeFeedback(&_ReputationRegistry.TransactOpts, agentId, feedbackIndex)
}

// ReputationRegistryFeedbackRevokedIterator is returned from FilterFeedbackRevoked and is used to iterate over the raw logs and unpacked data for FeedbackRevoked events raised by the ReputationRegistry contract.
type ReputationRegistryFeedbackRevokedIterator struct {
	Event *ReputationRegistryFeedbackRevoked // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryFeedbackRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryFeedbackRevoked)
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
		it.Event = new(ReputationRegistryFeedbackRevoked)
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
func (it *ReputationRegistryFeedbackRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryFeedbackRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryFeedbackRevoked represents a FeedbackRevoked event raised by the ReputationRegistry contract.
type ReputationRegistryFeedbackRevoked struct {
	AgentId       *big.Int
	ClientAddress common.Address
	FeedbackIndex uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeedbackRevoked is a free log retrieval operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterFeedbackRevoked(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, feedbackIndex []uint64) (*ReputationRegistryFeedbackRevokedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var feedbackIndexRule []interface{}
	for _, feedbackIndexItem := range feedbackIndex {
		feedbackIndexRule = append(feedbackIndexRule, feedbackIndexItem)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "FeedbackRevoked", agentIdRule, clientAddressRule, feedbackIndexRule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryFeedbackRevokedIterator{contract: _ReputationRegistry.contract, event: "FeedbackRevoked", logs: logs, sub: sub}, nil
}

// WatchFeedbackRevoked is a free log subscription operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchFeedbackRevoked(opts *bind.WatchOpts, sink chan<- *ReputationRegistryFeedbackRevoked, agentId []*big.Int, clientAddress []common.Address, feedbackIndex []uint64) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var feedbackIndexRule []interface{}
	for _, feedbackIndexItem := range feedbackIndex {
		feedbackIndexRule = append(feedbackIndexRule, feedbackIndexItem)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "FeedbackRevoked", agentIdRule, clientAddressRule, feedbackIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryFeedbackRevoked)
				if err := _ReputationRegistry.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
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

// ParseFeedbackRevoked is a log parse operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_ReputationRegistry *ReputationRegistryFilterer) ParseFeedbackRevoked(log types.Log) (*ReputationRegistryFeedbackRevoked, error) {
	event := new(ReputationRegistryFeedbackRevoked)
	if err := _ReputationRegistry.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryNewFeedbackIterator is returned from FilterNewFeedback and is used to iterate over the raw logs and unpacked data for NewFeedback events raised by the ReputationRegistry contract.
type ReputationRegistryNewFeedbackIterator struct {
	Event *ReputationRegistryNewFeedback // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryNewFeedbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryNewFeedback)
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
		it.Event = new(ReputationRegistryNewFeedback)
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
func (it *ReputationRegistryNewFeedbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryNewFeedbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryNewFeedback represents a NewFeedback event raised by the ReputationRegistry contract.
type ReputationRegistryNewFeedback struct {
	AgentId       *big.Int
	ClientAddress common.Address
	Score         uint8
	Tag1          [32]byte
	Tag2          [32]byte
	Fileuri       string
	Filehash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewFeedback is a free log retrieval operation binding the contract event 0x54b3254e4cc01969e70376d20390cc2f6af90539d9adaa78a53ebcda17f78154.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint8 score, bytes32 indexed tag1, bytes32 tag2, string fileuri, bytes32 filehash)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterNewFeedback(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, tag1 [][32]byte) (*ReputationRegistryNewFeedbackIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var tag1Rule []interface{}
	for _, tag1Item := range tag1 {
		tag1Rule = append(tag1Rule, tag1Item)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "NewFeedback", agentIdRule, clientAddressRule, tag1Rule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryNewFeedbackIterator{contract: _ReputationRegistry.contract, event: "NewFeedback", logs: logs, sub: sub}, nil
}

// WatchNewFeedback is a free log subscription operation binding the contract event 0x54b3254e4cc01969e70376d20390cc2f6af90539d9adaa78a53ebcda17f78154.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint8 score, bytes32 indexed tag1, bytes32 tag2, string fileuri, bytes32 filehash)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchNewFeedback(opts *bind.WatchOpts, sink chan<- *ReputationRegistryNewFeedback, agentId []*big.Int, clientAddress []common.Address, tag1 [][32]byte) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var tag1Rule []interface{}
	for _, tag1Item := range tag1 {
		tag1Rule = append(tag1Rule, tag1Item)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "NewFeedback", agentIdRule, clientAddressRule, tag1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryNewFeedback)
				if err := _ReputationRegistry.contract.UnpackLog(event, "NewFeedback", log); err != nil {
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

// ParseNewFeedback is a log parse operation binding the contract event 0x54b3254e4cc01969e70376d20390cc2f6af90539d9adaa78a53ebcda17f78154.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint8 score, bytes32 indexed tag1, bytes32 tag2, string fileuri, bytes32 filehash)
func (_ReputationRegistry *ReputationRegistryFilterer) ParseNewFeedback(log types.Log) (*ReputationRegistryNewFeedback, error) {
	event := new(ReputationRegistryNewFeedback)
	if err := _ReputationRegistry.contract.UnpackLog(event, "NewFeedback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationRegistryResponseAppendedIterator is returned from FilterResponseAppended and is used to iterate over the raw logs and unpacked data for ResponseAppended events raised by the ReputationRegistry contract.
type ReputationRegistryResponseAppendedIterator struct {
	Event *ReputationRegistryResponseAppended // Event containing the contract specifics and raw log

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
func (it *ReputationRegistryResponseAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationRegistryResponseAppended)
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
		it.Event = new(ReputationRegistryResponseAppended)
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
func (it *ReputationRegistryResponseAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationRegistryResponseAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationRegistryResponseAppended represents a ResponseAppended event raised by the ReputationRegistry contract.
type ReputationRegistryResponseAppended struct {
	AgentId       *big.Int
	ClientAddress common.Address
	FeedbackIndex uint64
	Responder     common.Address
	ResponseUri   string
	ResponseHash  [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterResponseAppended is a free log retrieval operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseUri, bytes32 responseHash)
func (_ReputationRegistry *ReputationRegistryFilterer) FilterResponseAppended(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, responder []common.Address) (*ReputationRegistryResponseAppendedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var responderRule []interface{}
	for _, responderItem := range responder {
		responderRule = append(responderRule, responderItem)
	}

	logs, sub, err := _ReputationRegistry.contract.FilterLogs(opts, "ResponseAppended", agentIdRule, clientAddressRule, responderRule)
	if err != nil {
		return nil, err
	}
	return &ReputationRegistryResponseAppendedIterator{contract: _ReputationRegistry.contract, event: "ResponseAppended", logs: logs, sub: sub}, nil
}

// WatchResponseAppended is a free log subscription operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseUri, bytes32 responseHash)
func (_ReputationRegistry *ReputationRegistryFilterer) WatchResponseAppended(opts *bind.WatchOpts, sink chan<- *ReputationRegistryResponseAppended, agentId []*big.Int, clientAddress []common.Address, responder []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var responderRule []interface{}
	for _, responderItem := range responder {
		responderRule = append(responderRule, responderItem)
	}

	logs, sub, err := _ReputationRegistry.contract.WatchLogs(opts, "ResponseAppended", agentIdRule, clientAddressRule, responderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationRegistryResponseAppended)
				if err := _ReputationRegistry.contract.UnpackLog(event, "ResponseAppended", log); err != nil {
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

// ParseResponseAppended is a log parse operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseUri, bytes32 responseHash)
func (_ReputationRegistry *ReputationRegistryFilterer) ParseResponseAppended(log types.Log) (*ReputationRegistryResponseAppended, error) {
	event := new(ReputationRegistryResponseAppended)
	if err := _ReputationRegistry.contract.UnpackLog(event, "ResponseAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
