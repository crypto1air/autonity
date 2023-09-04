// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package autonity

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/autonity/autonity"
	"github.com/autonity/autonity/accounts/abi"
	"github.com/autonity/autonity/accounts/abi/bind"
	"github.com/autonity/autonity/common"
	"github.com/autonity/autonity/core/types"
	"github.com/autonity/autonity/event"
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

// AccountabilityConfig is an auto generated low-level Go binding around an user-defined struct.
type AccountabilityConfig struct {
	InnocenceProofSubmissionWindow  *big.Int
	LatestAccountabilityEventsRange *big.Int
	BaseSlashingRateLow             *big.Int
	BaseSlashingRateMid             *big.Int
	CollusionFactor                 *big.Int
	HistoryFactor                   *big.Int
	JailFactor                      *big.Int
	SlashingRatePrecision           *big.Int
}

// AccountabilityEvent is an auto generated low-level Go binding around an user-defined struct.
type AccountabilityEvent struct {
	Chunks         uint8
	ChunkId        uint8
	EventType      uint8
	Rule           uint8
	Reporter       common.Address
	Offender       common.Address
	RawProof       []byte
	Block          *big.Int
	Epoch          *big.Int
	ReportingBlock *big.Int
	MessageHash    *big.Int
}

// AutonityCommitteeMember is an auto generated low-level Go binding around an user-defined struct.
type AutonityCommitteeMember struct {
	Addr        common.Address
	VotingPower *big.Int
}

// AutonityConfig is an auto generated low-level Go binding around an user-defined struct.
type AutonityConfig struct {
	Policy          AutonityPolicy
	Contracts       AutonityContracts
	Protocol        AutonityProtocol
	ContractVersion *big.Int
}

// AutonityContracts is an auto generated low-level Go binding around an user-defined struct.
type AutonityContracts struct {
	AccountabilityContract common.Address
	OracleContract         common.Address
	AcuContract            common.Address
	SupplyControlContract  common.Address
	StabilizationContract  common.Address
}

// AutonityPolicy is an auto generated low-level Go binding around an user-defined struct.
type AutonityPolicy struct {
	TreasuryFee     *big.Int
	MinBaseFee      *big.Int
	DelegationRate  *big.Int
	UnbondingPeriod *big.Int
	TreasuryAccount common.Address
}

// AutonityProtocol is an auto generated low-level Go binding around an user-defined struct.
type AutonityProtocol struct {
	OperatorAccount common.Address
	EpochPeriod     *big.Int
	BlockPeriod     *big.Int
	CommitteeSize   *big.Int
}

// AutonityValidator is an auto generated low-level Go binding around an user-defined struct.
type AutonityValidator struct {
	Treasury            common.Address
	NodeAddress         common.Address
	OracleAddress       common.Address
	Enode               string
	CommissionRate      *big.Int
	BondedStake         *big.Int
	UnbondingStake      *big.Int
	UnbondingShares     *big.Int
	SelfBondedStake     *big.Int
	SelfUnbondingStake  *big.Int
	SelfUnbondingShares *big.Int
	LiquidContract      common.Address
	LiquidSupply        *big.Int
	RegistrationBlock   *big.Int
	TotalSlashed        *big.Int
	JailReleaseBlock    *big.Int
	ProvableFaultCount  *big.Int
	State               uint8
}

// IOracleRoundData is an auto generated low-level Go binding around an user-defined struct.
type IOracleRoundData struct {
	Round     *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Status    *big.Int
}

// StabilizationConfig is an auto generated low-level Go binding around an user-defined struct.
type StabilizationConfig struct {
	BorrowInterestRate        *big.Int
	LiquidationRatio          *big.Int
	MinCollateralizationRatio *big.Int
	MinDebtRequirement        *big.Int
	TargetPrice               *big.Int
}

// ACUMetaData contains all meta data concerning the ACU contract.
var ACUMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"symbols_\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"scale_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"autonity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidBasket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoACUValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"symbols\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"BasketModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"symbols_\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"scale_\",\"type\":\"uint256\"}],\"name\":\"modifyBasket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quantities\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaleFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbols\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"44b4708a": "modifyBasket(string[],uint256[],uint256)",
		"d54d2799": "quantities()",
		"146ca531": "round()",
		"f51e181a": "scale()",
		"683dd191": "scaleFactor()",
		"b3ab15fb": "setOperator(address)",
		"7adbf973": "setOracle(address)",
		"07039ff9": "symbols()",
		"a2e62045": "update()",
		"3fa4f245": "value()",
	},
	Bin: "0x608060405234801562000010575f80fd5b50604051620018433803806200184383398101604081905262000033916200035b565b858580518251146200005857604051634ff799c560e01b815260040160405180910390fd5b5f5b8151811015620000be576001600160ff1b038282815181106200008157620000816200050d565b60200260200101511115620000a957604051634ff799c560e01b815260040160405180910390fd5b80620000b58162000535565b9150506200005a565b508751620000d49060039060208b019062000149565b508651620000ea9060049060208a0190620001a4565b506001869055620000fd86600a6200064b565b6002555050600680546001600160a01b039485166001600160a01b03199182161790915560078054938516938216939093179092556008805491909316911617905550620007b3915050565b828054828255905f5260205f2090810192821562000192579160200282015b82811115620001925782518290620001819082620006eb565b509160200191906001019062000168565b50620001a0929150620001ee565b5090565b828054828255905f5260205f20908101928215620001e0579160200282015b82811115620001e0578251825591602001919060010190620001c3565b50620001a09291506200020e565b80821115620001a0575f62000204828262000224565b50600101620001ee565b5b80821115620001a0575f81556001016200020f565b50805462000232906200065f565b5f825580601f1062000242575050565b601f0160209004905f5260205f20908101906200026091906200020e565b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715620002a257620002a262000263565b604052919050565b5f6001600160401b03821115620002c557620002c562000263565b5060051b60200190565b5f82601f830112620002df575f80fd5b81516020620002f8620002f283620002aa565b62000277565b82815260059290921b8401810191818101908684111562000317575f80fd5b8286015b848110156200033457805183529183019183016200031b565b509695505050505050565b80516001600160a01b038116811462000356575f80fd5b919050565b5f805f805f8060c0878903121562000371575f80fd5b86516001600160401b0381111562000387575f80fd5b8701601f8101891362000398575f80fd5b8051620003a9620002f282620002aa565b808282526020820191508b60208460051b8601011115620003c8575f80fd5b602084015b60208460051b8601018110156200049a5780516001600160401b03811115620003f4575f80fd5b8d603f828801011262000405575f80fd5b858101602001516001600160401b0381111562000426576200042662000263565b6200043b601f8201601f191660200162000277565b8181528f604083858b010101111562000452575f80fd5b5f5b828110156200047957604081858b010101516020828401015260208101905062000454565b505f60208383010152808652505050602083019250602081019050620003cd565b5060208b0151909950925050506001600160401b03811115620004bb575f80fd5b620004c989828a01620002cf565b95505060408701519350620004e1606088016200033f565b9250620004f1608088016200033f565b91506200050160a088016200033f565b90509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f6001820162000549576200054962000521565b5060010190565b600181815b808511156200059057815f190482111562000574576200057462000521565b808516156200058257918102915b93841c939080029062000555565b509250929050565b5f82620005a85750600162000645565b81620005b657505f62000645565b8160018114620005cf5760028114620005da57620005fa565b600191505062000645565b60ff841115620005ee57620005ee62000521565b50506001821b62000645565b5060208310610133831016604e8410600b84101617156200061f575081810a62000645565b6200062b838362000550565b805f190482111562000641576200064162000521565b0290505b92915050565b5f62000658838362000598565b9392505050565b600181811c908216806200067457607f821691505b6020821081036200069357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620006e6575f81815260208120601f850160051c81016020861015620006c15750805b601f850160051c820191505b81811015620006e257828155600101620006cd565b5050505b505050565b81516001600160401b0381111562000707576200070762000263565b6200071f816200071884546200065f565b8462000699565b602080601f83116001811462000755575f84156200073d5750858301515b5f19600386901b1c1916600185901b178555620006e2565b5f85815260208120601f198616915b82811015620007855788860151825594840194600190910190840162000764565b5085821015620007a357878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b61108280620007c15f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c80637adbf973116100635780637adbf973146100f9578063a2e620451461010c578063b3ab15fb14610124578063d54d279914610137578063f51e181a1461014c575f80fd5b806307039ff91461009f578063146ca531146100bd5780633fa4f245146100d357806344b4708a146100db578063683dd191146100f0575b5f80fd5b6100a7610155565b6040516100b4919061090c565b60405180910390f35b6100c55f5481565b6040519081526020016100b4565b6100c5610229565b6100ee6100e93660046109f5565b610252565b005b6100c560025481565b6100ee610107366004610b1c565b61037d565b6101146103c9565b60405190151581526020016100b4565b6100ee610132366004610b1c565b6106e1565b61013f61072d565b6040516100b49190610b7b565b6100c560015481565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015610220578382905f5260205f2001805461019590610b8d565b80601f01602080910402602001604051908101604052809291908181526020018280546101c190610b8d565b801561020c5780601f106101e35761010080835404028352916020019161020c565b820191905f5260205f20905b8154815290600101906020018083116101ef57829003601f168201915b505050505081526020019060010190610178565b50505050905090565b5f80545f0361024b57604051631d3e00bb60e11b815260040160405180910390fd5b5060055490565b8282805182511461027657604051634ff799c560e01b815260040160405180910390fd5b5f5b81518110156102d4576001600160ff1b0382828151811061029b5761029b610bc5565b602002602001015111156102c257604051634ff799c560e01b815260040160405180910390fd5b806102cc81610bed565b915050610278565b506007546001600160a01b031633146102ff576040516282b42960e81b815260040160405180910390fd5b8451610312906003906020880190610783565b5083516103269060049060208701906107d7565b50600183905561033783600a610ce7565b6002556040517fdbdcd10543a20811a4a332247f28d03b22686d3281043de35824a06075c06c099061036e90879087908790610cf2565b60405180910390a15050505050565b6006546001600160a01b031633146103a7576040516282b42960e81b815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6006545f906001600160a01b031633146103f5576040516282b42960e81b815260040160405180910390fd5b5f600160085f9054906101000a90046001600160a01b03166001600160a01b0316639f8743f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610448573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046c9190610d27565b6104769190610d3e565b9050805f5410610487575f91505090565b5f8060085f9054906101000a90046001600160a01b03166001600160a01b0316639670c0bc6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fd9190610d27565b90505f5b60035481101561067d57604051661554d10bd554d160ca1b60208201525f90602701604051602081830303815290604052805190602001206003838154811061054c5761054c610bc5565b905f5260205f20016040516020016105649190610d51565b6040516020818303038152906040528051906020012003610586575081610636565b600854600380545f926001600160a01b031691633c8510fd91899190879081106105b2576105b2610bc5565b905f5260205f20016040518363ffffffff1660e01b81526004016105d7929190610dc3565b608060405180830381865afa1580156105f2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106169190610e50565b905080606001515f1461062f575f965050505050505090565b6020015190505b6004828154811061064957610649610bc5565b905f5260205f2001548161065d9190610eb4565b6106679085610ee3565b935050808061067590610bed565b915050610501565b506106888183610f0a565b60058190555f849055604080514381524260208201528082018690526060810192909252517f23f161ca67071b3e902d4fa7afade82672c6160677e89d373a830145bdda6d269181900360800190a16001935050505090565b6006546001600160a01b0316331461070b576040516282b42960e81b815260040160405180910390fd5b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6060600480548060200260200160405190810160405280929190818152602001828054801561077957602002820191905f5260205f20905b815481526020019060010190808311610765575b5050505050905090565b828054828255905f5260205f209081019282156107c7579160200282015b828111156107c757825182906107b79082610f90565b50916020019190600101906107a1565b506107d392915061081c565b5090565b828054828255905f5260205f20908101928215610810579160200282015b828111156108105782518255916020019190600101906107f5565b506107d3929150610838565b808211156107d3575f61082f828261084c565b5060010161081c565b5b808211156107d3575f8155600101610839565b50805461085890610b8d565b5f825580601f10610867575050565b601f0160209004905f5260205f20908101906108839190610838565b50565b5f82825180855260208086019550808260051b8401018186015f805b858110156108fe57601f1980888603018b5283518051808752845b818110156108d8578281018901518882018a015288016108bd565b5086810188018590529b87019b601f0190911690940185019350918401916001016108a2565b509198975050505050505050565b602081525f61091e6020830184610886565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561096257610962610925565b604052919050565b5f67ffffffffffffffff82111561098357610983610925565b5060051b60200190565b5f82601f83011261099c575f80fd5b813560206109b16109ac8361096a565b610939565b82815260059290921b840181019181810190868411156109cf575f80fd5b8286015b848110156109ea57803583529183019183016109d3565b509695505050505050565b5f805f60608486031215610a07575f80fd5b833567ffffffffffffffff80821115610a1e575f80fd5b818601915086601f830112610a31575f80fd5b81356020610a416109ac8361096a565b82815260059290921b8401810191818101908a841115610a5f575f80fd5b8286015b84811015610ae857803586811115610a7a575f8081fd5b8701603f81018d13610a8b575f8081fd5b84810135604088821115610aa157610aa1610925565b610ab3601f8301601f19168801610939565b8281528f82848601011115610ac7575f8081fd5b82828501898301375f92810188019290925250845250918301918301610a63565b5097505087013592505080821115610afe575f80fd5b50610b0b8682870161098d565b925050604084013590509250925092565b5f60208284031215610b2c575f80fd5b81356001600160a01b038116811461091e575f80fd5b5f8151808452602080850194508084015f5b83811015610b7057815187529582019590820190600101610b54565b509495945050505050565b602081525f61091e6020830184610b42565b600181811c90821680610ba157607f821691505b602082108103610bbf57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201610bfe57610bfe610bd9565b5060010190565b600181815b80851115610c3f57815f1904821115610c2557610c25610bd9565b80851615610c3257918102915b93841c9390800290610c0a565b509250929050565b5f82610c5557506001610ce1565b81610c6157505f610ce1565b8160018114610c775760028114610c8157610c9d565b6001915050610ce1565b60ff841115610c9257610c92610bd9565b50506001821b610ce1565b5060208310610133831016604e8410600b8410161715610cc0575081810a610ce1565b610cca8383610c05565b805f1904821115610cdd57610cdd610bd9565b0290505b92915050565b5f61091e8383610c47565b606081525f610d046060830186610886565b8281036020840152610d168186610b42565b915050826040830152949350505050565b5f60208284031215610d37575f80fd5b5051919050565b81810381811115610ce157610ce1610bd9565b5f808354610d5e81610b8d565b60018281168015610d765760018114610d8b57610db7565b60ff1984168752821515830287019450610db7565b875f526020805f205f5b85811015610dae5781548a820152908401908201610d95565b50505082870194505b50929695505050505050565b8281525f60206040818401525f8454610ddb81610b8d565b806040870152606060018084165f8114610dfc5760018114610e1657610e41565b60ff1985168984015283151560051b890183019550610e41565b895f52865f205f5b85811015610e395781548b8201860152908301908801610e1e565b8a0184019650505b50939998505050505050505050565b5f60808284031215610e60575f80fd5b6040516080810181811067ffffffffffffffff82111715610e8357610e83610925565b8060405250825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b8082025f8212600160ff1b84141615610ecf57610ecf610bd9565b8181058314821517610ce157610ce1610bd9565b8082018281125f831280158216821582161715610f0257610f02610bd9565b505092915050565b5f82610f2457634e487b7160e01b5f52601260045260245ffd5b600160ff1b82145f1984141615610f3d57610f3d610bd9565b500590565b601f821115610f8b575f81815260208120601f850160051c81016020861015610f685750805b601f850160051c820191505b81811015610f8757828155600101610f74565b5050505b505050565b815167ffffffffffffffff811115610faa57610faa610925565b610fbe81610fb88454610b8d565b84610f42565b602080601f831160018114610ff1575f8415610fda5750858301515b5f19600386901b1c1916600185901b178555610f87565b5f85815260208120601f198616915b8281101561101f57888601518255948401946001909101908401611000565b508582101561103c57878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220200551c3027de9b78d276b7cbb2c0a2663963dbea17fa35e97946ee3c1956aa464736f6c63430008150033",
}

// ACUABI is the input ABI used to generate the binding from.
// Deprecated: Use ACUMetaData.ABI instead.
var ACUABI = ACUMetaData.ABI

// Deprecated: Use ACUMetaData.Sigs instead.
// ACUFuncSigs maps the 4-byte function signature to its string representation.
var ACUFuncSigs = ACUMetaData.Sigs

// ACUBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ACUMetaData.Bin instead.
var ACUBin = ACUMetaData.Bin

// DeployACU deploys a new Ethereum contract, binding an instance of ACU to it.
func DeployACU(auth *bind.TransactOpts, backend bind.ContractBackend, symbols_ []string, quantities_ []*big.Int, scale_ *big.Int, autonity common.Address, operator common.Address, oracle common.Address) (common.Address, *types.Transaction, *ACU, error) {
	parsed, err := ACUMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ACUBin), backend, symbols_, quantities_, scale_, autonity, operator, oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ACU{ACUCaller: ACUCaller{contract: contract}, ACUTransactor: ACUTransactor{contract: contract}, ACUFilterer: ACUFilterer{contract: contract}}, nil
}

// ACU is an auto generated Go binding around an Ethereum contract.
type ACU struct {
	ACUCaller     // Read-only binding to the contract
	ACUTransactor // Write-only binding to the contract
	ACUFilterer   // Log filterer for contract events
}

// ACUCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACUCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACUTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACUTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACUFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACUFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACUSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACUSession struct {
	Contract     *ACU              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACUCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACUCallerSession struct {
	Contract *ACUCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ACUTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACUTransactorSession struct {
	Contract     *ACUTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACURaw is an auto generated low-level Go binding around an Ethereum contract.
type ACURaw struct {
	Contract *ACU // Generic contract binding to access the raw methods on
}

// ACUCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACUCallerRaw struct {
	Contract *ACUCaller // Generic read-only contract binding to access the raw methods on
}

// ACUTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACUTransactorRaw struct {
	Contract *ACUTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACU creates a new instance of ACU, bound to a specific deployed contract.
func NewACU(address common.Address, backend bind.ContractBackend) (*ACU, error) {
	contract, err := bindACU(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACU{ACUCaller: ACUCaller{contract: contract}, ACUTransactor: ACUTransactor{contract: contract}, ACUFilterer: ACUFilterer{contract: contract}}, nil
}

// NewACUCaller creates a new read-only instance of ACU, bound to a specific deployed contract.
func NewACUCaller(address common.Address, caller bind.ContractCaller) (*ACUCaller, error) {
	contract, err := bindACU(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACUCaller{contract: contract}, nil
}

// NewACUTransactor creates a new write-only instance of ACU, bound to a specific deployed contract.
func NewACUTransactor(address common.Address, transactor bind.ContractTransactor) (*ACUTransactor, error) {
	contract, err := bindACU(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACUTransactor{contract: contract}, nil
}

// NewACUFilterer creates a new log filterer instance of ACU, bound to a specific deployed contract.
func NewACUFilterer(address common.Address, filterer bind.ContractFilterer) (*ACUFilterer, error) {
	contract, err := bindACU(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACUFilterer{contract: contract}, nil
}

// bindACU binds a generic wrapper to an already deployed contract.
func bindACU(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ACUABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACU *ACURaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACU.Contract.ACUCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACU *ACURaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACU.Contract.ACUTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACU *ACURaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACU.Contract.ACUTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACU *ACUCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACU.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACU *ACUTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACU.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACU *ACUTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACU.Contract.contract.Transact(opts, method, params...)
}

// Quantities is a free data retrieval call binding the contract method 0xd54d2799.
//
// Solidity: function quantities() view returns(uint256[])
func (_ACU *ACUCaller) Quantities(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "quantities")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Quantities is a free data retrieval call binding the contract method 0xd54d2799.
//
// Solidity: function quantities() view returns(uint256[])
func (_ACU *ACUSession) Quantities() ([]*big.Int, error) {
	return _ACU.Contract.Quantities(&_ACU.CallOpts)
}

// Quantities is a free data retrieval call binding the contract method 0xd54d2799.
//
// Solidity: function quantities() view returns(uint256[])
func (_ACU *ACUCallerSession) Quantities() ([]*big.Int, error) {
	return _ACU.Contract.Quantities(&_ACU.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_ACU *ACUCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_ACU *ACUSession) Round() (*big.Int, error) {
	return _ACU.Contract.Round(&_ACU.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_ACU *ACUCallerSession) Round() (*big.Int, error) {
	return _ACU.Contract.Round(&_ACU.CallOpts)
}

// Scale is a free data retrieval call binding the contract method 0xf51e181a.
//
// Solidity: function scale() view returns(uint256)
func (_ACU *ACUCaller) Scale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scale is a free data retrieval call binding the contract method 0xf51e181a.
//
// Solidity: function scale() view returns(uint256)
func (_ACU *ACUSession) Scale() (*big.Int, error) {
	return _ACU.Contract.Scale(&_ACU.CallOpts)
}

// Scale is a free data retrieval call binding the contract method 0xf51e181a.
//
// Solidity: function scale() view returns(uint256)
func (_ACU *ACUCallerSession) Scale() (*big.Int, error) {
	return _ACU.Contract.Scale(&_ACU.CallOpts)
}

// ScaleFactor is a free data retrieval call binding the contract method 0x683dd191.
//
// Solidity: function scaleFactor() view returns(uint256)
func (_ACU *ACUCaller) ScaleFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "scaleFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaleFactor is a free data retrieval call binding the contract method 0x683dd191.
//
// Solidity: function scaleFactor() view returns(uint256)
func (_ACU *ACUSession) ScaleFactor() (*big.Int, error) {
	return _ACU.Contract.ScaleFactor(&_ACU.CallOpts)
}

// ScaleFactor is a free data retrieval call binding the contract method 0x683dd191.
//
// Solidity: function scaleFactor() view returns(uint256)
func (_ACU *ACUCallerSession) ScaleFactor() (*big.Int, error) {
	return _ACU.Contract.ScaleFactor(&_ACU.CallOpts)
}

// Symbols is a free data retrieval call binding the contract method 0x07039ff9.
//
// Solidity: function symbols() view returns(string[])
func (_ACU *ACUCaller) Symbols(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "symbols")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Symbols is a free data retrieval call binding the contract method 0x07039ff9.
//
// Solidity: function symbols() view returns(string[])
func (_ACU *ACUSession) Symbols() ([]string, error) {
	return _ACU.Contract.Symbols(&_ACU.CallOpts)
}

// Symbols is a free data retrieval call binding the contract method 0x07039ff9.
//
// Solidity: function symbols() view returns(string[])
func (_ACU *ACUCallerSession) Symbols() ([]string, error) {
	return _ACU.Contract.Symbols(&_ACU.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(int256)
func (_ACU *ACUCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ACU.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(int256)
func (_ACU *ACUSession) Value() (*big.Int, error) {
	return _ACU.Contract.Value(&_ACU.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(int256)
func (_ACU *ACUCallerSession) Value() (*big.Int, error) {
	return _ACU.Contract.Value(&_ACU.CallOpts)
}

// ModifyBasket is a paid mutator transaction binding the contract method 0x44b4708a.
//
// Solidity: function modifyBasket(string[] symbols_, uint256[] quantities_, uint256 scale_) returns()
func (_ACU *ACUTransactor) ModifyBasket(opts *bind.TransactOpts, symbols_ []string, quantities_ []*big.Int, scale_ *big.Int) (*types.Transaction, error) {
	return _ACU.contract.Transact(opts, "modifyBasket", symbols_, quantities_, scale_)
}

// ModifyBasket is a paid mutator transaction binding the contract method 0x44b4708a.
//
// Solidity: function modifyBasket(string[] symbols_, uint256[] quantities_, uint256 scale_) returns()
func (_ACU *ACUSession) ModifyBasket(symbols_ []string, quantities_ []*big.Int, scale_ *big.Int) (*types.Transaction, error) {
	return _ACU.Contract.ModifyBasket(&_ACU.TransactOpts, symbols_, quantities_, scale_)
}

// ModifyBasket is a paid mutator transaction binding the contract method 0x44b4708a.
//
// Solidity: function modifyBasket(string[] symbols_, uint256[] quantities_, uint256 scale_) returns()
func (_ACU *ACUTransactorSession) ModifyBasket(symbols_ []string, quantities_ []*big.Int, scale_ *big.Int) (*types.Transaction, error) {
	return _ACU.Contract.ModifyBasket(&_ACU.TransactOpts, symbols_, quantities_, scale_)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ACU *ACUTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ACU.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ACU *ACUSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ACU.Contract.SetOperator(&_ACU.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ACU *ACUTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ACU.Contract.SetOperator(&_ACU.TransactOpts, operator)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_ACU *ACUTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _ACU.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_ACU *ACUSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _ACU.Contract.SetOracle(&_ACU.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_ACU *ACUTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _ACU.Contract.SetOracle(&_ACU.TransactOpts, oracle)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_ACU *ACUTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACU.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_ACU *ACUSession) Update() (*types.Transaction, error) {
	return _ACU.Contract.Update(&_ACU.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_ACU *ACUTransactorSession) Update() (*types.Transaction, error) {
	return _ACU.Contract.Update(&_ACU.TransactOpts)
}

// ACUBasketModifiedIterator is returned from FilterBasketModified and is used to iterate over the raw logs and unpacked data for BasketModified events raised by the ACU contract.
type ACUBasketModifiedIterator struct {
	Event *ACUBasketModified // Event containing the contract specifics and raw log

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
func (it *ACUBasketModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACUBasketModified)
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
		it.Event = new(ACUBasketModified)
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
func (it *ACUBasketModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACUBasketModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACUBasketModified represents a BasketModified event raised by the ACU contract.
type ACUBasketModified struct {
	Symbols    []string
	Quantities []*big.Int
	Scale      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBasketModified is a free log retrieval operation binding the contract event 0xdbdcd10543a20811a4a332247f28d03b22686d3281043de35824a06075c06c09.
//
// Solidity: event BasketModified(string[] symbols, uint256[] quantities, uint256 scale)
func (_ACU *ACUFilterer) FilterBasketModified(opts *bind.FilterOpts) (*ACUBasketModifiedIterator, error) {

	logs, sub, err := _ACU.contract.FilterLogs(opts, "BasketModified")
	if err != nil {
		return nil, err
	}
	return &ACUBasketModifiedIterator{contract: _ACU.contract, event: "BasketModified", logs: logs, sub: sub}, nil
}

// WatchBasketModified is a free log subscription operation binding the contract event 0xdbdcd10543a20811a4a332247f28d03b22686d3281043de35824a06075c06c09.
//
// Solidity: event BasketModified(string[] symbols, uint256[] quantities, uint256 scale)
func (_ACU *ACUFilterer) WatchBasketModified(opts *bind.WatchOpts, sink chan<- *ACUBasketModified) (event.Subscription, error) {

	logs, sub, err := _ACU.contract.WatchLogs(opts, "BasketModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACUBasketModified)
				if err := _ACU.contract.UnpackLog(event, "BasketModified", log); err != nil {
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

// ParseBasketModified is a log parse operation binding the contract event 0xdbdcd10543a20811a4a332247f28d03b22686d3281043de35824a06075c06c09.
//
// Solidity: event BasketModified(string[] symbols, uint256[] quantities, uint256 scale)
func (_ACU *ACUFilterer) ParseBasketModified(log types.Log) (*ACUBasketModified, error) {
	event := new(ACUBasketModified)
	if err := _ACU.contract.UnpackLog(event, "BasketModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACUUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the ACU contract.
type ACUUpdatedIterator struct {
	Event *ACUUpdated // Event containing the contract specifics and raw log

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
func (it *ACUUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACUUpdated)
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
		it.Event = new(ACUUpdated)
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
func (it *ACUUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACUUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACUUpdated represents a Updated event raised by the ACU contract.
type ACUUpdated struct {
	Height    *big.Int
	Timestamp *big.Int
	Round     *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0x23f161ca67071b3e902d4fa7afade82672c6160677e89d373a830145bdda6d26.
//
// Solidity: event Updated(uint256 height, uint256 timestamp, uint256 round, int256 value)
func (_ACU *ACUFilterer) FilterUpdated(opts *bind.FilterOpts) (*ACUUpdatedIterator, error) {

	logs, sub, err := _ACU.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &ACUUpdatedIterator{contract: _ACU.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0x23f161ca67071b3e902d4fa7afade82672c6160677e89d373a830145bdda6d26.
//
// Solidity: event Updated(uint256 height, uint256 timestamp, uint256 round, int256 value)
func (_ACU *ACUFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *ACUUpdated) (event.Subscription, error) {

	logs, sub, err := _ACU.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACUUpdated)
				if err := _ACU.contract.UnpackLog(event, "Updated", log); err != nil {
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

// ParseUpdated is a log parse operation binding the contract event 0x23f161ca67071b3e902d4fa7afade82672c6160677e89d373a830145bdda6d26.
//
// Solidity: event Updated(uint256 height, uint256 timestamp, uint256 round, int256 value)
func (_ACU *ACUFilterer) ParseUpdated(log types.Log) (*ACUUpdated, error) {
	event := new(ACUUpdated)
	if err := _ACU.contract.UnpackLog(event, "Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountabilityMetaData contains all meta data concerning the Accountability contract.
var AccountabilityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_autonity\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"innocenceProofSubmissionWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestAccountabilityEventsRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseSlashingRateLow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseSlashingRateMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collusionFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"historyFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRatePrecision\",\"type\":\"uint256\"}],\"internalType\":\"structAccountability.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"InnocenceProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_severity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewAccusation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_severity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewFaultProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseBlock\",\"type\":\"uint256\"}],\"name\":\"SlashingEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"beneficiaries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"_rule\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"canAccuse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"_rule\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"canSlash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"innocenceProofSubmissionWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestAccountabilityEventsRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseSlashingRateLow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseSlashingRateMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collusionFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"historyFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashingRatePrecision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"chunks\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"chunkId\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.EventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"rule\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageHash\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_epochEnd\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"getValidatorAccusation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"chunks\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"chunkId\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.EventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"rule\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageHash\",\"type\":\"uint256\"}],\"internalType\":\"structAccountability.Event\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"getValidatorFaults\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"chunks\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"chunkId\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.EventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"rule\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageHash\",\"type\":\"uint256\"}],\"internalType\":\"structAccountability.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"chunks\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"chunkId\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.EventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"enumAccountability.Rule\",\"name\":\"rule\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageHash\",\"type\":\"uint256\"}],\"internalType\":\"structAccountability.Event\",\"name\":\"_event\",\"type\":\"tuple\"}],\"name\":\"handleEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01567739": "beneficiaries(address)",
		"7ccecadd": "canAccuse(address,uint8,uint256)",
		"4108a95a": "canSlash(address,uint8,uint256)",
		"79502c55": "config()",
		"1de9d9b6": "distributeRewards(address)",
		"b5b7a184": "epochPeriod()",
		"0b791430": "events(uint256)",
		"6c9789b0": "finalize(bool)",
		"9cb22b06": "getValidatorAccusation(address)",
		"bebaa8fc": "getValidatorFaults(address)",
		"069f6863": "handleEvent((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256))",
		"6b5f444c": "setEpochPeriod(uint256)",
		"e7bb0b52": "slashingHistory(address,uint256)",
	},
	Bin: "0x60806040525f60125534801562000014575f80fd5b5060405162003942380380620039428339810160408190526200003791620000fa565b600180546001600160a01b0319166001600160a01b03841690811790915560408051636fd8d26960e11b8152905163dfb1a4d2916004808201926020929091908290030181865afa1580156200008f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001c8565b5f558051600355602081015160045560408101516005556060810151600655608081015160075560a081015160085560c081015160095560e00151600a5550620001e0565b5f808284036101208112156200010e575f80fd5b83516001600160a01b038116811462000125575f80fd5b9250610100601f1982018113156200013b575f80fd5b60405191508082016001600160401b03811183821017156200016b57634e487b7160e01b5f52604160045260245ffd5b80604052506020850151825260408501516020830152606085015160408301526080850151606083015260a0850151608083015260c085015160a083015260e085015160c08301528085015160e083015250809150509250929050565b5f60208284031215620001d9575f80fd5b5051919050565b61375480620001ee5f395ff3fe6080604052600436106100bf575f3560e01c80636c9789b01161007c5780639cb22b06116100575780639cb22b061461028e578063b5b7a184146102ba578063bebaa8fc146102dc578063e7bb0b5214610308575f80fd5b80636c9789b0146101cc57806379502c55146101eb5780637ccecadd14610258575f80fd5b806301567739146100c3578063069f6863146101145780630b791430146101355780631de9d9b61461016b5780634108a95a1461017e5780636b5f444c146101ad575b5f80fd5b3480156100ce575f80fd5b506100f76100dd366004612bba565b600b6020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561011f575f80fd5b5061013361012e366004612d1d565b61033e565b005b348015610140575f80fd5b5061015461014f366004612e19565b61069b565b60405161010b9b9a99989796959493929190612ebe565b610133610179366004612bba565b6107a4565b348015610189575f80fd5b5061019d610198366004612f45565b61099b565b604051901515815260200161010b565b3480156101b8575f80fd5b506101336101c7366004612e19565b610a46565b3480156101d7575f80fd5b506101336101e6366004612f80565b610a74565b3480156101f6575f80fd5b50600354600454600554600654600754600854600954600a5461021d979695949392919088565b604080519889526020890197909752958701949094526060860192909252608085015260a084015260c083015260e08201526101000161010b565b348015610263575f80fd5b50610277610272366004612f45565b610ab7565b60408051921515835260208301919091520161010b565b348015610299575f80fd5b506102ad6102a8366004612bba565b610bfd565b60405161010b9190613069565b3480156102c5575f80fd5b506102ce5f5481565b60405190815260200161010b565b3480156102e7575f80fd5b506102fb6102f6366004612bba565b610dc1565b60405161010b919061307b565b348015610313575f80fd5b506102ce6103223660046130db565b600f60209081525f928352604080842090915290825290205481565b600154604051630c825d9760e11b81523360048201525f916001600160a01b031690631904bb2e906024015f60405180830381865afa158015610383573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103aa9190810190613165565b60208101519091506001600160a01b031633146104245760405162461bcd60e51b815260206004820152602d60248201527f66756e6374696f6e207265737472696374656420746f2061207265676973746560448201526c3932b2103b30b634b230ba37b960991b60648201526084015b60405180910390fd5b60808201516001600160a01b031633146104805760405162461bcd60e51b815260206004820152601d60248201527f6576656e74207265706f72746572206d7573742062652063616c6c6572000000604482015260640161041b565b6001825f015160ff161115610626575f6104998361104d565b9050806104a557505050565b335f908152600e6020908152604091829020825161016081018452815460ff808216835261010082048116948301949094529093919291840191620100009091041660028111156104f8576104f8612e30565b600281111561050957610509612e30565b815281546020909101906301000000900460ff16600d81111561052e5761052e612e30565b600d81111561053f5761053f612e30565b815281546001600160a01b03600160201b90910481166020830152600183015416604082015260028201805460609092019161057a906132ad565b80601f01602080910402602001604051908101604052809291908181526020018280546105a6906132ad565b80156105f15780601f106105c8576101008083540402835291602001916105f1565b820191905f5260205f20905b8154815290600101906020018083116105d457829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815250509250505b5f8260400151600281111561063d5761063d612e30565b0361064f5761064b82611273565b5050565b60018260400151600281111561066757610667612e30565b036106755761064b82611669565b60028260400151600281111561068d5761068d612e30565b0361064b5761064b82611a7d565b600281815481106106aa575f80fd5b5f91825260209091206007909102018054600182015460028301805460ff8085169650610100850481169562010000860482169563010000008104909216946001600160a01b03600160201b90930483169492169290919061070b906132ad565b80601f0160208091040260200160405190810160405280929190818152602001828054610737906132ad565b80156107825780601f1061075957610100808354040283529160200191610782565b820191905f5260205f20905b81548152906001019060200180831161076557829003601f168201915b505050505090806003015490806004015490806005015490806006015490508b565b6001546001600160a01b031633146107ce5760405162461bcd60e51b815260040161041b906132e5565b6001546001600160a01b038281165f908152600b6020526040808220549051630c825d9760e11b8152908316600482015290929190911690631904bb2e906024015f60405180830381865afa158015610829573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108509190810190613165565b516040519091505f906001600160a01b038316906108fc90349084818181858888f193505050503d805f81146108a1576040519150601f19603f3d011682016040523d82523d5f602084013e6108a6565b606091505b50509050806109735760015f9054906101000a90046001600160a01b03166001600160a01b031663f7866ee36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109239190613329565b6001600160a01b0316346040515f6040518083038185875af1925050503d805f811461096a576040519150601f19603f3d011682016040523d82523d5f602084013e61096f565b606091505b5050505b50506001600160a01b03165f908152600b6020526040902080546001600160a01b0319169055565b5f806109a684611e81565b6001546040516396b477cb60e01b8152600481018690529192505f916001600160a01b03909116906396b477cb90602401602060405180830381865afa1580156109f2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a169190613344565b6001600160a01b0387165f908152600f602090815260408083209383529290522054919091109150509392505050565b6001546001600160a01b03163314610a705760405162461bcd60e51b815260040161041b906132e5565b5f55565b6001546001600160a01b03163314610a9e5760405162461bcd60e51b815260040161041b906132e5565b610aa6611eea565b8015610ab457610ab461222c565b50565b5f805f610ac385611e81565b6001546040516396b477cb60e01b8152600481018790529192505f916001600160a01b03909116906396b477cb90602401602060405180830381865afa158015610b0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b339190613344565b6001600160a01b0388165f908152600f602090815260408083208484529091529020549091508211610b6a575f93505f9250610bf3565b6001600160a01b0387165f908152600d602052604090205415610beb576001600160a01b0387165f908152600d6020526040812054600290610bae9060019061336f565b81548110610bbe57610bbe613382565b905f5260205f20906007020190505f945060035f01548160030154610be39190613396565b935050610bf3565b600193505f92505b5050935093915050565b610c05612ae1565b6001600160a01b0382165f908152600d6020526040902054600290610c2c9060019061336f565b81548110610c3c57610c3c613382565b5f91825260209182902060408051610160810182526007909302909101805460ff808216855261010082048116958501959095529293909291840191620100009004166002811115610c9057610c90612e30565b6002811115610ca157610ca1612e30565b815281546020909101906301000000900460ff16600d811115610cc657610cc6612e30565b600d811115610cd757610cd7612e30565b815281546001600160a01b03600160201b909104811660208301526001830154166040820152600282018054606090920191610d12906132ad565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3e906132ad565b8015610d895780601f10610d6057610100808354040283529160200191610d89565b820191905f5260205f20905b815481529060010190602001808311610d6c57829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815250509050919050565b6001600160a01b0381165f908152600c60205260408120546060919067ffffffffffffffff811115610df557610df5612bdc565b604051908082528060200260200182016040528015610e2e57816020015b610e1b612ae1565b815260200190600190039081610e135790505b5090505f5b6001600160a01b0384165f908152600c6020526040902054811015611046576001600160a01b0384165f908152600c6020526040902080546002919083908110610e7f57610e7f613382565b905f5260205f20015481548110610e9857610e98613382565b5f91825260209182902060408051610160810182526007909302909101805460ff808216855261010082048116958501959095529293909291840191620100009004166002811115610eec57610eec612e30565b6002811115610efd57610efd612e30565b815281546020909101906301000000900460ff16600d811115610f2257610f22612e30565b600d811115610f3357610f33612e30565b815281546001600160a01b03600160201b909104811660208301526001830154166040820152600282018054606090920191610f6e906132ad565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9a906132ad565b8015610fe55780601f10610fbc57610100808354040283529160200191610fe5565b820191905f5260205f20905b815481529060010190602001808311610fc857829003601f168201915b5050505050815260200160038201548152602001600482015481526020016005820154815260200160068201548152505082828151811061102857611028613382565b6020026020010181905250808061103e906133a9565b915050610e33565b5092915050565b5f816020015160ff165f0361117357335f908152600e6020908152604091829020845181549286015160ff9081166101000261ffff1990941691161791909117808255918401518492829062ff00001916620100008360028111156110b4576110b4612e30565b021790555060608201518154829063ff0000001916630100000083600d8111156110e0576110e0612e30565b021790555060808201518154640100000000600160c01b031916600160201b6001600160a01b039283160217825560a08301516001830180546001600160a01b0319169190921617905560c0820151600282019061113e908261340f565b5060e082015160038201556101008201516004820155610120820151600582015561014090910151600690910155505f919050565b602080830151335f908152600e90925260409091205460ff918216916111a09161010090041660016134cb565b60ff16146111f05760405162461bcd60e51b815260206004820152601960248201527f6368756e6b73206d75737420626520636f6e746967756f757300000000000000604482015260640161041b565b335f908152600e6020526040902060c083015161121091600201906124f4565b335f908152600e602052604090208054600191908290611239908290610100900460ff166134cb565b92506101000a81548160ff021916908360ff160217905550815f015160ff168260200151600161126991906134cb565b60ff161492915050565b5f805f805f61128760fe8760c0015161263b565b94509450945094509450846112de5760405162461bcd60e51b815260206004820152601960248201527f6661696c65642070726f6f6620766572696669636174696f6e00000000000000604482015260640161041b565b8560a001516001600160a01b0316846001600160a01b0316146113135760405162461bcd60e51b815260040161041b906134e4565b8560600151600d81111561132957611329612e30565b83146113475760405162461bcd60e51b815260040161041b9061350f565b4382106113665760405162461bcd60e51b815260040161041b90613539565b815f0361137b5761137860014361336f565b91505b5f6113898760600151611e81565b6001546040516396b477cb60e01b8152600481018690529192505f916001600160a01b03909116906396b477cb90602401602060405180830381865afa1580156113d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113f99190613344565b6001600160a01b0387165f908152600f60209081526040808320848452909152902054909150821161143d5760405162461bcd60e51b815260040161041b90613569565b60e08801849052610100808901829052436101208a01526101408901849052600280546001810182555f8290528a5160079091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805460208d015160ff90811690950261ffff1990911694909216939093171780835560408b01518b93929091839162ff00001990911690620100009084908111156114e0576114e0612e30565b021790555060608201518154829063ff0000001916630100000083600d81111561150c5761150c612e30565b021790555060808201518154640100000000600160c01b031916600160201b6001600160a01b039283160217825560a08301516001830180546001600160a01b0319169190921617905560c0820151600282019061156a908261340f565b5060e0820151600382015561010082015160048201556101208201516005820155610140909101516006909101556002545f906115a99060019061336f565b60a08a01516001600160a01b039081165f908152600c60209081526040808320805460018082018355918552838520018690556010805491820190557f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67201859055928b16808352600f82528383208784528252918390208790558251878152908101849052929350917f6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f91015b60405180910390a2505050505050505050565b5f805f805f61167d60fc8760c0015161263b565b94509450945094509450846116d45760405162461bcd60e51b815260206004820152601e60248201527f6661696c65642061636375736174696f6e20766572696669636174696f6e0000604482015260640161041b565b8560a001516001600160a01b0316846001600160a01b0316146117095760405162461bcd60e51b815260040161041b906134e4565b8560600151600d81111561171f5761171f612e30565b831461173d5760405162461bcd60e51b815260040161041b9061350f565b43821061175c5760405162461bcd60e51b815260040161041b90613539565b5f61176a8760600151611e81565b6001546040516396b477cb60e01b8152600481018690529192505f916001600160a01b03909116906396b477cb90602401602060405180830381865afa1580156117b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117da9190613344565b6001600160a01b0387165f908152600f60209081526040808320848452909152902054909150821161181e5760405162461bcd60e51b815260040161041b90613569565b60a08801516001600160a01b03165f908152600d6020526040902054156118875760405162461bcd60e51b815260206004820181905260248201527f616c72656164792070726f63657373696e6720616e2061636375736174696f6e604482015260640161041b565b60e08801849052610100808901829052436101208a01526101408901849052600280546001810182555f8290528a5160079091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805460208d015160ff90811690950261ffff1990911694909216939093171780835560408b01518b93929091839162ff000019909116906201000090849081111561192a5761192a612e30565b021790555060608201518154829063ff0000001916630100000083600d81111561195657611956612e30565b021790555060808201518154640100000000600160c01b031916600160201b6001600160a01b039283160217825560a08301516001830180546001600160a01b0319169190921617905560c082015160028201906119b4908261340f565b5060e0820151600382015561010082015160048201556101208201516005820155610140909101516006909101556002545f906119f39060019061336f565b9050611a00816001613396565b60a08a01516001600160a01b03165f908152600d60205260409020556011611a29826001613396565b81546001810183555f9283526020928390200155604080518581529182018390526001600160a01b038916917f2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba403519101611656565b5f805f805f611a9160fd8760c0015161263b565b9450945094509450945084611ae85760405162461bcd60e51b815260206004820152601d60248201527f6661696c656420696e6e6f63656e636520766572696669636174696f6e000000604482015260640161041b565b8560a001516001600160a01b0316846001600160a01b031614611b1d5760405162461bcd60e51b815260040161041b906134e4565b8560600151600d811115611b3357611b33612e30565b8314611b515760405162461bcd60e51b815260040161041b9061350f565b438210611b705760405162461bcd60e51b815260040161041b90613539565b60a08601516001600160a01b03165f908152600d602052604081205490819003611bdc5760405162461bcd60e51b815260206004820152601860248201527f6e6f206173736f6369617465642061636375736174696f6e0000000000000000604482015260640161041b565b83600d811115611bee57611bee612e30565b600d811115611bff57611bff612e30565b6002611c0c60018461336f565b81548110611c1c57611c1c613382565b5f9182526020909120600790910201546301000000900460ff16600d811115611c4757611c47612e30565b14611ca45760405162461bcd60e51b815260206004820152602760248201527f756e6d61746368696e672070726f6f6620616e642061636375736174696f6e206044820152661c9d5b19481a5960ca1b606482015260840161041b565b826002611cb260018461336f565b81548110611cc257611cc2613382565b905f5260205f2090600702016003015414611d2d5760405162461bcd60e51b815260206004820152602560248201527f756e6d61746368696e672070726f6f6620616e642061636375736174696f6e20604482015264626c6f636b60d81b606482015260840161041b565b816002611d3b60018461336f565b81548110611d4b57611d4b613382565b905f5260205f2090600702016006015414611db45760405162461bcd60e51b8152602060048201526024808201527f756e6d61746368696e672070726f6f6620616e642061636375736174696f6e206044820152630d0c2e6d60e31b606482015260840161041b565b6012545b601154811015611e1a578160118281548110611dd657611dd6613382565b905f5260205f20015403611e08575f60118281548110611df857611df8613382565b5f91825260209091200155611e1a565b80611e12816133a9565b915050611db8565b506001600160a01b038086165f908152600d602052604080822082905560a08a015190519216917f1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f91611e709190815260200190565b60405180910390a250505050505050565b5f600a82600d811115611e9657611e96612e30565b03611ea45760025b92915050565b5f82600d811115611eb757611eb7612e30565b03611ec3576002611e9e565b600182600d811115611ed757611ed7612e30565b03611ee3576002611e9e565b6002611e9e565b6012545b601154811015612227575f60118281548110611f0c57611f0c613382565b905f5260205f2001549050805f03611f245750612215565b611f2f60018261336f565b90505f60028281548110611f4557611f45613382565b5f91825260209182902060408051610160810182526007909302909101805460ff808216855261010082048116958501959095529293909291840191620100009004166002811115611f9957611f99612e30565b6002811115611faa57611faa612e30565b815281546020909101906301000000900460ff16600d811115611fcf57611fcf612e30565b600d811115611fe057611fe0612e30565b815281546001600160a01b03600160201b90910481166020830152600183015416604082015260028201805460609092019161201b906132ad565b80601f0160208091040260200160405190810160405280929190818152602001828054612047906132ad565b80156120925780601f1061206957610100808354040283529160200191612092565b820191905f5260205f20905b81548152906001019060200180831161207557829003601f168201915b5050505050815260200160038201548152602001600482015481526020016005820154815260200160068201548152505090504360035f01548261012001516120db9190613396565b11156120e8575050601255565b60a08101516001600160a01b03165f908152600d60205260408120819055606082015161211490611e81565b60a08301516001600160a01b03165f908152600f602090815260408083206101008701518452909152902054909150811161215157505050612215565b60a0820180516001600160a01b039081165f908152600f6020908152604080832061010088015184528252808320869055845184168352600c8252808320805460018082018355918552838520018990556010805491820181559093527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67290920187905592518151858152938401879052909116917f6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f910160405180910390a25050505b8061221f816133a9565b915050611eee565b601255565b5f8060015f9054906101000a90046001600160a01b03166001600160a01b031663c9d97af46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561227e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122a29190613344565b90505f5b60105481101561231457816002601083815481106122c6576122c6613382565b905f5260205f200154815481106122df576122df613382565b905f5260205f2090600702016004015403612302576122ff600184613396565b92505b8061230c816133a9565b9150506122a6565b505f5b6010548110156124e8576124d660026010838154811061233957612339613382565b905f5260205f2001548154811061235257612352613382565b5f91825260209182902060408051610160810182526007909302909101805460ff8082168552610100820481169585019590955292939092918401916201000090041660028111156123a6576123a6612e30565b60028111156123b7576123b7612e30565b815281546020909101906301000000900460ff16600d8111156123dc576123dc612e30565b600d8111156123ed576123ed612e30565b815281546001600160a01b03600160201b909104811660208301526001830154166040820152600282018054606090920191612428906132ad565b80601f0160208091040260200160405190810160405280929190818152602001828054612454906132ad565b801561249f5780601f106124765761010080835404028352916020019161249f565b820191905f5260205f20905b81548152906001019060200180831161248257829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815250508461269e565b806124e0816133a9565b915050612317565b5061064b60105f612b45565b81546002600180831615610100020382160482518082016020811060208410016002811461259c57600181146125c157865f526020840460205f2001600160028402018855602085068060200390508088018589016001836101000a0392508282511684540184556001840193506020820191505b808210156125865781518455600184019350602082019150612569565b815191036101000a908190040290915550612632565b60028302826020036101000a846020036101000a602089015104020185018755612632565b865f526020840460205f2001600160028402018855846020038088018589016001836101000a0392508282511660ff198a160184556020820191506001840193505b808210156126205781518455600184019350602082019150612603565b815191036101000a9081900402909155505b50505050505050565b5f805f805f808651602061264f9190613396565b9050612659612b60565b60a081838a8c5afa612669575f80fd5b805160010361267757600196505b602081015160408201516060830151608090930151989b919a509850909695509350505050565b60015460a0830151604051630c825d9760e11b81526001600160a01b0391821660048201525f929190911690631904bb2e906024015f60405180830381865afa1580156126ed573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526127149190810190613165565b608084015160a08501516001600160a01b039081165f908152600b6020526040812080546001600160a01b03191692909316919091179091556060850151919250906127689061276390611e81565b612a89565b610200830151600854919250905f9061278190836135ad565b60075461278e90876135ad565b6127989085613396565b6127a29190613396565b600a549091508111156127b45750600a545b5f8461012001518560c001518660a001516127cf9190613396565b6127d99190613396565b600a549091505f906127eb83856135ad565b6127f591906135c4565b90505f81905080876101200151106128255780876101200181815161281a919061336f565b9052505f905061283f565b610120870151612835908261336f565b5f61012089015290505b80156128ba57808761010001511061288657808761010001818151612864919061336f565b90525060a08701805182919061287b90839061336f565b9052505f90506128ba565b610100870151612896908261336f565b90508661010001518760a0018181516128af919061336f565b9052505f6101008801525b5f811180156128eb57505f8761010001518860a001516128da919061336f565b8860c001516128e99190613396565b115b15612968575f8761010001518860a00151612906919061336f565b8860c001516129159190613396565b60c089015161292490846135ad565b61292e91906135c4565b9050808860c001818151612942919061336f565b90525061294f818361336f565b9150818860a001818151612963919061336f565b905250505b81876101c00181815161297b9190613396565b9052506102008701805160019190612994908390613396565b9052505f546102008801516009546129ac91906135ad565b6129b691906135ad565b6129c09043613396565b6101e08801526002610220880152600154604051631b9d315560e21b81526001600160a01b0390911690636e74c554906129fe908a906004016135e3565b5f604051808303815f87803b158015612a15575f80fd5b505af1158015612a27573d5f803e3d5ffd5b505050506020878101516101e0890151604080516001600160a01b0390931683529282018590528183015290517f416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc9181900360600190a1505050505050505050565b5f81612a9757505060065490565b60018203612aa757505060065490565b60028203612ab757505060065490565b60038203612ac757505060065490565b60048203612ad85750612710919050565b50612710919050565b60408051610160810182525f80825260208201819052909182019081526020015f81526020015f6001600160a01b031681526020015f6001600160a01b03168152602001606081526020015f81526020015f81526020015f81526020015f81525090565b5080545f8255905f5260205f2090810190610ab49190612b7e565b6040518060a001604052806005906020820280368337509192915050565b5b80821115612b92575f8155600101612b7f565b5090565b6001600160a01b0381168114610ab4575f80fd5b8035612bb581612b96565b919050565b5f60208284031215612bca575f80fd5b8135612bd581612b96565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051610160810167ffffffffffffffff81118282101715612c1457612c14612bdc565b60405290565b604051610240810167ffffffffffffffff81118282101715612c1457612c14612bdc565b604051601f8201601f1916810167ffffffffffffffff81118282101715612c6757612c67612bdc565b604052919050565b803560ff81168114612bb5575f80fd5b60038110610ab4575f80fd5b8035612bb581612c7f565b8035600e8110612bb5575f80fd5b5f67ffffffffffffffff821115612cbd57612cbd612bdc565b50601f01601f191660200190565b5f82601f830112612cda575f80fd5b8135612ced612ce882612ca4565b612c3e565b818152846020838601011115612d01575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215612d2d575f80fd5b813567ffffffffffffffff80821115612d44575f80fd5b908301906101608286031215612d58575f80fd5b612d60612bf0565b612d6983612c6f565b8152612d7760208401612c6f565b6020820152612d8860408401612c8b565b6040820152612d9960608401612c96565b6060820152612daa60808401612baa565b6080820152612dbb60a08401612baa565b60a082015260c083013582811115612dd1575f80fd5b612ddd87828601612ccb565b60c08301525060e08381013590820152610100808401359082015261012080840135908201526101409283013592810192909252509392505050565b5f60208284031215612e29575f80fd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b60038110610ab457610ab4612e30565b612e5d81612e44565b9052565b600e8110612e5d57612e5d612e30565b5f5b83811015612e8b578181015183820152602001612e73565b50505f910152565b5f8151808452612eaa816020860160208601612e71565b601f01601f19169290920160200192915050565b5f61016060ff8e16835260ff8d166020840152612eda8c612e44565b8b6040840152612eed606084018c612e61565b6001600160a01b038a81166080850152891660a084015260c08301819052612f1781840189612e93565b9150508560e08301528461010083015283610120830152826101408301529c9b505050505050505050505050565b5f805f60608486031215612f57575f80fd5b8335612f6281612b96565b9250612f7060208501612c96565b9150604084013590509250925092565b5f60208284031215612f90575f80fd5b81358015158114612bd5575f80fd5b805160ff1682525f6101606020830151612fbe602086018260ff169052565b506040830151612fd16040860182612e54565b506060830151612fe46060860182612e61565b506080830151612fff60808601826001600160a01b03169052565b5060a083015161301a60a08601826001600160a01b03169052565b5060c08301518160c086015261303282860182612e93565b60e0858101519087015261010080860151908701526101208086015190870152610140948501519490950193909352509192915050565b602081525f612bd56020830184612f9f565b5f602080830181845280855180835260408601915060408160051b87010192508387015f5b828110156130ce57603f198886030184526130bc858351612f9f565b945092850192908501906001016130a0565b5092979650505050505050565b5f80604083850312156130ec575f80fd5b82356130f781612b96565b946020939093013593505050565b8051612bb581612b96565b5f82601f83011261311f575f80fd5b815161312d612ce882612ca4565b818152846020838601011115613141575f80fd5b613152826020830160208701612e71565b949350505050565b8051612bb581612c7f565b5f60208284031215613175575f80fd5b815167ffffffffffffffff8082111561318c575f80fd5b9083019061024082860312156131a0575f80fd5b6131a8612c1a565b6131b183613105565b81526131bf60208401613105565b60208201526131d060408401613105565b60408201526060830151828111156131e6575f80fd5b6131f287828601613110565b6060830152506080838101519082015260a0808401519082015260c0808401519082015260e080840151908201526101008084015190820152610120808401519082015261014080840151908201526101609150613251828401613105565b9181019190915261018082810151908201526101a080830151908201526101c080830151908201526101e0808301519082015261020080830151908201526102209061329e82840161315a565b91810191909152949350505050565b600181811c908216806132c157607f821691505b6020821081036132df57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526024908201527f66756e6374696f6e207265737472696374656420746f207468652076616c696460408201526330ba37b960e11b606082015260800190565b5f60208284031215613339575f80fd5b8151612bd581612b96565b5f60208284031215613354575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115611e9e57611e9e61335b565b634e487b7160e01b5f52603260045260245ffd5b80820180821115611e9e57611e9e61335b565b5f600182016133ba576133ba61335b565b5060010190565b601f82111561340a575f81815260208120601f850160051c810160208610156133e75750805b601f850160051c820191505b81811015613406578281556001016133f3565b5050505b505050565b815167ffffffffffffffff81111561342957613429612bdc565b61343d8161343784546132ad565b846133c1565b602080601f831160018114613470575f84156134595750858301515b5f19600386901b1c1916600185901b178555613406565b5f85815260208120601f198616915b8281101561349e5788860151825594840194600190910190840161347f565b50858210156134bb57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b60ff8181168382160190811115611e9e57611e9e61335b565b6020808252601190820152700decccccadcc8cae440dad2e6dac2e8c6d607b1b604082015260600190565b60208082526010908201526f0e4ead8ca40d2c840dad2e6dac2e8c6d60831b604082015260600190565b60208082526016908201527563616e277420626520696e207468652066757475726560501b604082015260600190565b60208082526024908201527f616c726561647920736c6173686564206174207468652070726f6f66277320656040820152630e0dec6d60e31b606082015260800190565b8082028115828204841417611e9e57611e9e61335b565b5f826135de57634e487b7160e01b5f52601260045260245ffd5b500490565b602081526135fd6020820183516001600160a01b03169052565b5f602083015161361860408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015161024080608085015261364a610260850183612e93565b9150608085015160a085015260a085015160c085015260c085015160e085015260e08501516101008181870152808701519150506101208181870152808701519150506101408181870152808701519150506101608181870152808701519150506101806136c2818701836001600160a01b03169052565b8601516101a0868101919091528601516101c0808701919091528601516101e0808701919091528601516102008087019190915286015161022080870191909152860151905061371482860182612e54565b509094935050505056fea26469706673582212204ee2dbacf54476e41b4cc6fd62d6e1d8298fea0d088cd797bf26749233b7ef4864736f6c63430008150033",
}

// AccountabilityABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountabilityMetaData.ABI instead.
var AccountabilityABI = AccountabilityMetaData.ABI

// Deprecated: Use AccountabilityMetaData.Sigs instead.
// AccountabilityFuncSigs maps the 4-byte function signature to its string representation.
var AccountabilityFuncSigs = AccountabilityMetaData.Sigs

// AccountabilityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountabilityMetaData.Bin instead.
var AccountabilityBin = AccountabilityMetaData.Bin

// DeployAccountability deploys a new Ethereum contract, binding an instance of Accountability to it.
func DeployAccountability(auth *bind.TransactOpts, backend bind.ContractBackend, _autonity common.Address, _config AccountabilityConfig) (common.Address, *types.Transaction, *Accountability, error) {
	parsed, err := AccountabilityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountabilityBin), backend, _autonity, _config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accountability{AccountabilityCaller: AccountabilityCaller{contract: contract}, AccountabilityTransactor: AccountabilityTransactor{contract: contract}, AccountabilityFilterer: AccountabilityFilterer{contract: contract}}, nil
}

// Accountability is an auto generated Go binding around an Ethereum contract.
type Accountability struct {
	AccountabilityCaller     // Read-only binding to the contract
	AccountabilityTransactor // Write-only binding to the contract
	AccountabilityFilterer   // Log filterer for contract events
}

// AccountabilityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountabilityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountabilityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountabilitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountabilitySession struct {
	Contract     *Accountability   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountabilityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountabilityCallerSession struct {
	Contract *AccountabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountabilityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountabilityTransactorSession struct {
	Contract     *AccountabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountabilityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountabilityRaw struct {
	Contract *Accountability // Generic contract binding to access the raw methods on
}

// AccountabilityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountabilityCallerRaw struct {
	Contract *AccountabilityCaller // Generic read-only contract binding to access the raw methods on
}

// AccountabilityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountabilityTransactorRaw struct {
	Contract *AccountabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountability creates a new instance of Accountability, bound to a specific deployed contract.
func NewAccountability(address common.Address, backend bind.ContractBackend) (*Accountability, error) {
	contract, err := bindAccountability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accountability{AccountabilityCaller: AccountabilityCaller{contract: contract}, AccountabilityTransactor: AccountabilityTransactor{contract: contract}, AccountabilityFilterer: AccountabilityFilterer{contract: contract}}, nil
}

// NewAccountabilityCaller creates a new read-only instance of Accountability, bound to a specific deployed contract.
func NewAccountabilityCaller(address common.Address, caller bind.ContractCaller) (*AccountabilityCaller, error) {
	contract, err := bindAccountability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountabilityCaller{contract: contract}, nil
}

// NewAccountabilityTransactor creates a new write-only instance of Accountability, bound to a specific deployed contract.
func NewAccountabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountabilityTransactor, error) {
	contract, err := bindAccountability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountabilityTransactor{contract: contract}, nil
}

// NewAccountabilityFilterer creates a new log filterer instance of Accountability, bound to a specific deployed contract.
func NewAccountabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountabilityFilterer, error) {
	contract, err := bindAccountability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountabilityFilterer{contract: contract}, nil
}

// bindAccountability binds a generic wrapper to an already deployed contract.
func bindAccountability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountability *AccountabilityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountability.Contract.AccountabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountability *AccountabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountability.Contract.AccountabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountability *AccountabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountability.Contract.AccountabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountability *AccountabilityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountability *AccountabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountability *AccountabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountability.Contract.contract.Transact(opts, method, params...)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x01567739.
//
// Solidity: function beneficiaries(address ) view returns(address)
func (_Accountability *AccountabilityCaller) Beneficiaries(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "beneficiaries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiaries is a free data retrieval call binding the contract method 0x01567739.
//
// Solidity: function beneficiaries(address ) view returns(address)
func (_Accountability *AccountabilitySession) Beneficiaries(arg0 common.Address) (common.Address, error) {
	return _Accountability.Contract.Beneficiaries(&_Accountability.CallOpts, arg0)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x01567739.
//
// Solidity: function beneficiaries(address ) view returns(address)
func (_Accountability *AccountabilityCallerSession) Beneficiaries(arg0 common.Address) (common.Address, error) {
	return _Accountability.Contract.Beneficiaries(&_Accountability.CallOpts, arg0)
}

// CanAccuse is a free data retrieval call binding the contract method 0x7ccecadd.
//
// Solidity: function canAccuse(address _offender, uint8 _rule, uint256 _block) view returns(bool _result, uint256 _deadline)
func (_Accountability *AccountabilityCaller) CanAccuse(opts *bind.CallOpts, _offender common.Address, _rule uint8, _block *big.Int) (struct {
	Result   bool
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "canAccuse", _offender, _rule, _block)

	outstruct := new(struct {
		Result   bool
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanAccuse is a free data retrieval call binding the contract method 0x7ccecadd.
//
// Solidity: function canAccuse(address _offender, uint8 _rule, uint256 _block) view returns(bool _result, uint256 _deadline)
func (_Accountability *AccountabilitySession) CanAccuse(_offender common.Address, _rule uint8, _block *big.Int) (struct {
	Result   bool
	Deadline *big.Int
}, error) {
	return _Accountability.Contract.CanAccuse(&_Accountability.CallOpts, _offender, _rule, _block)
}

// CanAccuse is a free data retrieval call binding the contract method 0x7ccecadd.
//
// Solidity: function canAccuse(address _offender, uint8 _rule, uint256 _block) view returns(bool _result, uint256 _deadline)
func (_Accountability *AccountabilityCallerSession) CanAccuse(_offender common.Address, _rule uint8, _block *big.Int) (struct {
	Result   bool
	Deadline *big.Int
}, error) {
	return _Accountability.Contract.CanAccuse(&_Accountability.CallOpts, _offender, _rule, _block)
}

// CanSlash is a free data retrieval call binding the contract method 0x4108a95a.
//
// Solidity: function canSlash(address _offender, uint8 _rule, uint256 _block) view returns(bool)
func (_Accountability *AccountabilityCaller) CanSlash(opts *bind.CallOpts, _offender common.Address, _rule uint8, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "canSlash", _offender, _rule, _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSlash is a free data retrieval call binding the contract method 0x4108a95a.
//
// Solidity: function canSlash(address _offender, uint8 _rule, uint256 _block) view returns(bool)
func (_Accountability *AccountabilitySession) CanSlash(_offender common.Address, _rule uint8, _block *big.Int) (bool, error) {
	return _Accountability.Contract.CanSlash(&_Accountability.CallOpts, _offender, _rule, _block)
}

// CanSlash is a free data retrieval call binding the contract method 0x4108a95a.
//
// Solidity: function canSlash(address _offender, uint8 _rule, uint256 _block) view returns(bool)
func (_Accountability *AccountabilityCallerSession) CanSlash(_offender common.Address, _rule uint8, _block *big.Int) (bool, error) {
	return _Accountability.Contract.CanSlash(&_Accountability.CallOpts, _offender, _rule, _block)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 innocenceProofSubmissionWindow, uint256 latestAccountabilityEventsRange, uint256 baseSlashingRateLow, uint256 baseSlashingRateMid, uint256 collusionFactor, uint256 historyFactor, uint256 jailFactor, uint256 slashingRatePrecision)
func (_Accountability *AccountabilityCaller) Config(opts *bind.CallOpts) (struct {
	InnocenceProofSubmissionWindow  *big.Int
	LatestAccountabilityEventsRange *big.Int
	BaseSlashingRateLow             *big.Int
	BaseSlashingRateMid             *big.Int
	CollusionFactor                 *big.Int
	HistoryFactor                   *big.Int
	JailFactor                      *big.Int
	SlashingRatePrecision           *big.Int
}, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		InnocenceProofSubmissionWindow  *big.Int
		LatestAccountabilityEventsRange *big.Int
		BaseSlashingRateLow             *big.Int
		BaseSlashingRateMid             *big.Int
		CollusionFactor                 *big.Int
		HistoryFactor                   *big.Int
		JailFactor                      *big.Int
		SlashingRatePrecision           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnocenceProofSubmissionWindow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestAccountabilityEventsRange = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BaseSlashingRateLow = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BaseSlashingRateMid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CollusionFactor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HistoryFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.JailFactor = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.SlashingRatePrecision = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 innocenceProofSubmissionWindow, uint256 latestAccountabilityEventsRange, uint256 baseSlashingRateLow, uint256 baseSlashingRateMid, uint256 collusionFactor, uint256 historyFactor, uint256 jailFactor, uint256 slashingRatePrecision)
func (_Accountability *AccountabilitySession) Config() (struct {
	InnocenceProofSubmissionWindow  *big.Int
	LatestAccountabilityEventsRange *big.Int
	BaseSlashingRateLow             *big.Int
	BaseSlashingRateMid             *big.Int
	CollusionFactor                 *big.Int
	HistoryFactor                   *big.Int
	JailFactor                      *big.Int
	SlashingRatePrecision           *big.Int
}, error) {
	return _Accountability.Contract.Config(&_Accountability.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 innocenceProofSubmissionWindow, uint256 latestAccountabilityEventsRange, uint256 baseSlashingRateLow, uint256 baseSlashingRateMid, uint256 collusionFactor, uint256 historyFactor, uint256 jailFactor, uint256 slashingRatePrecision)
func (_Accountability *AccountabilityCallerSession) Config() (struct {
	InnocenceProofSubmissionWindow  *big.Int
	LatestAccountabilityEventsRange *big.Int
	BaseSlashingRateLow             *big.Int
	BaseSlashingRateMid             *big.Int
	CollusionFactor                 *big.Int
	HistoryFactor                   *big.Int
	JailFactor                      *big.Int
	SlashingRatePrecision           *big.Int
}, error) {
	return _Accountability.Contract.Config(&_Accountability.CallOpts)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_Accountability *AccountabilityCaller) EpochPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "epochPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_Accountability *AccountabilitySession) EpochPeriod() (*big.Int, error) {
	return _Accountability.Contract.EpochPeriod(&_Accountability.CallOpts)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_Accountability *AccountabilityCallerSession) EpochPeriod() (*big.Int, error) {
	return _Accountability.Contract.EpochPeriod(&_Accountability.CallOpts)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint8 chunks, uint8 chunkId, uint8 eventType, uint8 rule, address reporter, address offender, bytes rawProof, uint256 block, uint256 epoch, uint256 reportingBlock, uint256 messageHash)
func (_Accountability *AccountabilityCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Chunks         uint8
	ChunkId        uint8
	EventType      uint8
	Rule           uint8
	Reporter       common.Address
	Offender       common.Address
	RawProof       []byte
	Block          *big.Int
	Epoch          *big.Int
	ReportingBlock *big.Int
	MessageHash    *big.Int
}, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "events", arg0)

	outstruct := new(struct {
		Chunks         uint8
		ChunkId        uint8
		EventType      uint8
		Rule           uint8
		Reporter       common.Address
		Offender       common.Address
		RawProof       []byte
		Block          *big.Int
		Epoch          *big.Int
		ReportingBlock *big.Int
		MessageHash    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Chunks = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ChunkId = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.EventType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Rule = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Reporter = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Offender = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.RawProof = *abi.ConvertType(out[6], new([]byte)).(*[]byte)
	outstruct.Block = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Epoch = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ReportingBlock = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MessageHash = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint8 chunks, uint8 chunkId, uint8 eventType, uint8 rule, address reporter, address offender, bytes rawProof, uint256 block, uint256 epoch, uint256 reportingBlock, uint256 messageHash)
func (_Accountability *AccountabilitySession) Events(arg0 *big.Int) (struct {
	Chunks         uint8
	ChunkId        uint8
	EventType      uint8
	Rule           uint8
	Reporter       common.Address
	Offender       common.Address
	RawProof       []byte
	Block          *big.Int
	Epoch          *big.Int
	ReportingBlock *big.Int
	MessageHash    *big.Int
}, error) {
	return _Accountability.Contract.Events(&_Accountability.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint8 chunks, uint8 chunkId, uint8 eventType, uint8 rule, address reporter, address offender, bytes rawProof, uint256 block, uint256 epoch, uint256 reportingBlock, uint256 messageHash)
func (_Accountability *AccountabilityCallerSession) Events(arg0 *big.Int) (struct {
	Chunks         uint8
	ChunkId        uint8
	EventType      uint8
	Rule           uint8
	Reporter       common.Address
	Offender       common.Address
	RawProof       []byte
	Block          *big.Int
	Epoch          *big.Int
	ReportingBlock *big.Int
	MessageHash    *big.Int
}, error) {
	return _Accountability.Contract.Events(&_Accountability.CallOpts, arg0)
}

// GetValidatorAccusation is a free data retrieval call binding the contract method 0x9cb22b06.
//
// Solidity: function getValidatorAccusation(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256))
func (_Accountability *AccountabilityCaller) GetValidatorAccusation(opts *bind.CallOpts, _val common.Address) (AccountabilityEvent, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "getValidatorAccusation", _val)

	if err != nil {
		return *new(AccountabilityEvent), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountabilityEvent)).(*AccountabilityEvent)

	return out0, err

}

// GetValidatorAccusation is a free data retrieval call binding the contract method 0x9cb22b06.
//
// Solidity: function getValidatorAccusation(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256))
func (_Accountability *AccountabilitySession) GetValidatorAccusation(_val common.Address) (AccountabilityEvent, error) {
	return _Accountability.Contract.GetValidatorAccusation(&_Accountability.CallOpts, _val)
}

// GetValidatorAccusation is a free data retrieval call binding the contract method 0x9cb22b06.
//
// Solidity: function getValidatorAccusation(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256))
func (_Accountability *AccountabilityCallerSession) GetValidatorAccusation(_val common.Address) (AccountabilityEvent, error) {
	return _Accountability.Contract.GetValidatorAccusation(&_Accountability.CallOpts, _val)
}

// GetValidatorFaults is a free data retrieval call binding the contract method 0xbebaa8fc.
//
// Solidity: function getValidatorFaults(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256)[])
func (_Accountability *AccountabilityCaller) GetValidatorFaults(opts *bind.CallOpts, _val common.Address) ([]AccountabilityEvent, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "getValidatorFaults", _val)

	if err != nil {
		return *new([]AccountabilityEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccountabilityEvent)).(*[]AccountabilityEvent)

	return out0, err

}

// GetValidatorFaults is a free data retrieval call binding the contract method 0xbebaa8fc.
//
// Solidity: function getValidatorFaults(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256)[])
func (_Accountability *AccountabilitySession) GetValidatorFaults(_val common.Address) ([]AccountabilityEvent, error) {
	return _Accountability.Contract.GetValidatorFaults(&_Accountability.CallOpts, _val)
}

// GetValidatorFaults is a free data retrieval call binding the contract method 0xbebaa8fc.
//
// Solidity: function getValidatorFaults(address _val) view returns((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256)[])
func (_Accountability *AccountabilityCallerSession) GetValidatorFaults(_val common.Address) ([]AccountabilityEvent, error) {
	return _Accountability.Contract.GetValidatorFaults(&_Accountability.CallOpts, _val)
}

// SlashingHistory is a free data retrieval call binding the contract method 0xe7bb0b52.
//
// Solidity: function slashingHistory(address , uint256 ) view returns(uint256)
func (_Accountability *AccountabilityCaller) SlashingHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Accountability.contract.Call(opts, &out, "slashingHistory", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingHistory is a free data retrieval call binding the contract method 0xe7bb0b52.
//
// Solidity: function slashingHistory(address , uint256 ) view returns(uint256)
func (_Accountability *AccountabilitySession) SlashingHistory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Accountability.Contract.SlashingHistory(&_Accountability.CallOpts, arg0, arg1)
}

// SlashingHistory is a free data retrieval call binding the contract method 0xe7bb0b52.
//
// Solidity: function slashingHistory(address , uint256 ) view returns(uint256)
func (_Accountability *AccountabilityCallerSession) SlashingHistory(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Accountability.Contract.SlashingHistory(&_Accountability.CallOpts, arg0, arg1)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_Accountability *AccountabilityTransactor) DistributeRewards(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Accountability.contract.Transact(opts, "distributeRewards", _validator)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_Accountability *AccountabilitySession) DistributeRewards(_validator common.Address) (*types.Transaction, error) {
	return _Accountability.Contract.DistributeRewards(&_Accountability.TransactOpts, _validator)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_Accountability *AccountabilityTransactorSession) DistributeRewards(_validator common.Address) (*types.Transaction, error) {
	return _Accountability.Contract.DistributeRewards(&_Accountability.TransactOpts, _validator)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_Accountability *AccountabilityTransactor) Finalize(opts *bind.TransactOpts, _epochEnd bool) (*types.Transaction, error) {
	return _Accountability.contract.Transact(opts, "finalize", _epochEnd)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_Accountability *AccountabilitySession) Finalize(_epochEnd bool) (*types.Transaction, error) {
	return _Accountability.Contract.Finalize(&_Accountability.TransactOpts, _epochEnd)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_Accountability *AccountabilityTransactorSession) Finalize(_epochEnd bool) (*types.Transaction, error) {
	return _Accountability.Contract.Finalize(&_Accountability.TransactOpts, _epochEnd)
}

// HandleEvent is a paid mutator transaction binding the contract method 0x069f6863.
//
// Solidity: function handleEvent((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256) _event) returns()
func (_Accountability *AccountabilityTransactor) HandleEvent(opts *bind.TransactOpts, _event AccountabilityEvent) (*types.Transaction, error) {
	return _Accountability.contract.Transact(opts, "handleEvent", _event)
}

// HandleEvent is a paid mutator transaction binding the contract method 0x069f6863.
//
// Solidity: function handleEvent((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256) _event) returns()
func (_Accountability *AccountabilitySession) HandleEvent(_event AccountabilityEvent) (*types.Transaction, error) {
	return _Accountability.Contract.HandleEvent(&_Accountability.TransactOpts, _event)
}

// HandleEvent is a paid mutator transaction binding the contract method 0x069f6863.
//
// Solidity: function handleEvent((uint8,uint8,uint8,uint8,address,address,bytes,uint256,uint256,uint256,uint256) _event) returns()
func (_Accountability *AccountabilityTransactorSession) HandleEvent(_event AccountabilityEvent) (*types.Transaction, error) {
	return _Accountability.Contract.HandleEvent(&_Accountability.TransactOpts, _event)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_Accountability *AccountabilityTransactor) SetEpochPeriod(opts *bind.TransactOpts, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Accountability.contract.Transact(opts, "setEpochPeriod", _newPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_Accountability *AccountabilitySession) SetEpochPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Accountability.Contract.SetEpochPeriod(&_Accountability.TransactOpts, _newPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_Accountability *AccountabilityTransactorSession) SetEpochPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Accountability.Contract.SetEpochPeriod(&_Accountability.TransactOpts, _newPeriod)
}

// AccountabilityInnocenceProvenIterator is returned from FilterInnocenceProven and is used to iterate over the raw logs and unpacked data for InnocenceProven events raised by the Accountability contract.
type AccountabilityInnocenceProvenIterator struct {
	Event *AccountabilityInnocenceProven // Event containing the contract specifics and raw log

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
func (it *AccountabilityInnocenceProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountabilityInnocenceProven)
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
		it.Event = new(AccountabilityInnocenceProven)
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
func (it *AccountabilityInnocenceProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountabilityInnocenceProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountabilityInnocenceProven represents a InnocenceProven event raised by the Accountability contract.
type AccountabilityInnocenceProven struct {
	Offender common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInnocenceProven is a free log retrieval operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_Accountability *AccountabilityFilterer) FilterInnocenceProven(opts *bind.FilterOpts, _offender []common.Address) (*AccountabilityInnocenceProvenIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.FilterLogs(opts, "InnocenceProven", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &AccountabilityInnocenceProvenIterator{contract: _Accountability.contract, event: "InnocenceProven", logs: logs, sub: sub}, nil
}

// WatchInnocenceProven is a free log subscription operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_Accountability *AccountabilityFilterer) WatchInnocenceProven(opts *bind.WatchOpts, sink chan<- *AccountabilityInnocenceProven, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.WatchLogs(opts, "InnocenceProven", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountabilityInnocenceProven)
				if err := _Accountability.contract.UnpackLog(event, "InnocenceProven", log); err != nil {
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

// ParseInnocenceProven is a log parse operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_Accountability *AccountabilityFilterer) ParseInnocenceProven(log types.Log) (*AccountabilityInnocenceProven, error) {
	event := new(AccountabilityInnocenceProven)
	if err := _Accountability.contract.UnpackLog(event, "InnocenceProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountabilityNewAccusationIterator is returned from FilterNewAccusation and is used to iterate over the raw logs and unpacked data for NewAccusation events raised by the Accountability contract.
type AccountabilityNewAccusationIterator struct {
	Event *AccountabilityNewAccusation // Event containing the contract specifics and raw log

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
func (it *AccountabilityNewAccusationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountabilityNewAccusation)
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
		it.Event = new(AccountabilityNewAccusation)
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
func (it *AccountabilityNewAccusationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountabilityNewAccusationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountabilityNewAccusation represents a NewAccusation event raised by the Accountability contract.
type AccountabilityNewAccusation struct {
	Offender common.Address
	Severity *big.Int
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAccusation is a free log retrieval operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) FilterNewAccusation(opts *bind.FilterOpts, _offender []common.Address) (*AccountabilityNewAccusationIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.FilterLogs(opts, "NewAccusation", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &AccountabilityNewAccusationIterator{contract: _Accountability.contract, event: "NewAccusation", logs: logs, sub: sub}, nil
}

// WatchNewAccusation is a free log subscription operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) WatchNewAccusation(opts *bind.WatchOpts, sink chan<- *AccountabilityNewAccusation, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.WatchLogs(opts, "NewAccusation", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountabilityNewAccusation)
				if err := _Accountability.contract.UnpackLog(event, "NewAccusation", log); err != nil {
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

// ParseNewAccusation is a log parse operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) ParseNewAccusation(log types.Log) (*AccountabilityNewAccusation, error) {
	event := new(AccountabilityNewAccusation)
	if err := _Accountability.contract.UnpackLog(event, "NewAccusation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountabilityNewFaultProofIterator is returned from FilterNewFaultProof and is used to iterate over the raw logs and unpacked data for NewFaultProof events raised by the Accountability contract.
type AccountabilityNewFaultProofIterator struct {
	Event *AccountabilityNewFaultProof // Event containing the contract specifics and raw log

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
func (it *AccountabilityNewFaultProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountabilityNewFaultProof)
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
		it.Event = new(AccountabilityNewFaultProof)
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
func (it *AccountabilityNewFaultProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountabilityNewFaultProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountabilityNewFaultProof represents a NewFaultProof event raised by the Accountability contract.
type AccountabilityNewFaultProof struct {
	Offender common.Address
	Severity *big.Int
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFaultProof is a free log retrieval operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) FilterNewFaultProof(opts *bind.FilterOpts, _offender []common.Address) (*AccountabilityNewFaultProofIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.FilterLogs(opts, "NewFaultProof", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &AccountabilityNewFaultProofIterator{contract: _Accountability.contract, event: "NewFaultProof", logs: logs, sub: sub}, nil
}

// WatchNewFaultProof is a free log subscription operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) WatchNewFaultProof(opts *bind.WatchOpts, sink chan<- *AccountabilityNewFaultProof, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _Accountability.contract.WatchLogs(opts, "NewFaultProof", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountabilityNewFaultProof)
				if err := _Accountability.contract.UnpackLog(event, "NewFaultProof", log); err != nil {
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

// ParseNewFaultProof is a log parse operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_Accountability *AccountabilityFilterer) ParseNewFaultProof(log types.Log) (*AccountabilityNewFaultProof, error) {
	event := new(AccountabilityNewFaultProof)
	if err := _Accountability.contract.UnpackLog(event, "NewFaultProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountabilitySlashingEventIterator is returned from FilterSlashingEvent and is used to iterate over the raw logs and unpacked data for SlashingEvent events raised by the Accountability contract.
type AccountabilitySlashingEventIterator struct {
	Event *AccountabilitySlashingEvent // Event containing the contract specifics and raw log

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
func (it *AccountabilitySlashingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountabilitySlashingEvent)
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
		it.Event = new(AccountabilitySlashingEvent)
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
func (it *AccountabilitySlashingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountabilitySlashingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountabilitySlashingEvent represents a SlashingEvent event raised by the Accountability contract.
type AccountabilitySlashingEvent struct {
	Validator    common.Address
	Amount       *big.Int
	ReleaseBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSlashingEvent is a free log retrieval operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_Accountability *AccountabilityFilterer) FilterSlashingEvent(opts *bind.FilterOpts) (*AccountabilitySlashingEventIterator, error) {

	logs, sub, err := _Accountability.contract.FilterLogs(opts, "SlashingEvent")
	if err != nil {
		return nil, err
	}
	return &AccountabilitySlashingEventIterator{contract: _Accountability.contract, event: "SlashingEvent", logs: logs, sub: sub}, nil
}

// WatchSlashingEvent is a free log subscription operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_Accountability *AccountabilityFilterer) WatchSlashingEvent(opts *bind.WatchOpts, sink chan<- *AccountabilitySlashingEvent) (event.Subscription, error) {

	logs, sub, err := _Accountability.contract.WatchLogs(opts, "SlashingEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountabilitySlashingEvent)
				if err := _Accountability.contract.UnpackLog(event, "SlashingEvent", log); err != nil {
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

// ParseSlashingEvent is a log parse operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_Accountability *AccountabilityFilterer) ParseSlashingEvent(log types.Log) (*AccountabilitySlashingEvent, error) {
	event := new(AccountabilitySlashingEvent)
	if err := _Accountability.contract.UnpackLog(event, "SlashingEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityMetaData contains all meta data concerning the Autonity contract.
var AutonityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfBondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"contractLiquid\",\"name\":\"liquidContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSlashed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailReleaseBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provableFaultCount\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structAutonity.Validator[]\",\"name\":\"_validators\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAccount\",\"type\":\"address\"}],\"internalType\":\"structAutonity.Policy\",\"name\":\"policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIAccountability\",\"name\":\"accountabilityContract\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"oracleContract\",\"type\":\"address\"},{\"internalType\":\"contractIACU\",\"name\":\"acuContract\",\"type\":\"address\"},{\"internalType\":\"contractISupplyControl\",\"name\":\"supplyControlContract\",\"type\":\"address\"},{\"internalType\":\"contractIStabilization\",\"name\":\"stabilizationContract\",\"type\":\"address\"}],\"internalType\":\"structAutonity.Contracts\",\"name\":\"contracts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeSize\",\"type\":\"uint256\"}],\"internalType\":\"structAutonity.Protocol\",\"name\":\"protocol\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"internalType\":\"structAutonity.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlock\",\"type\":\"uint256\"}],\"name\":\"ActivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"EpochPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"MinimumBaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"selfBonded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewBondingRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"selfBonded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewUnbondingRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlock\",\"type\":\"uint256\"}],\"name\":\"PausedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidContract\",\"type\":\"address\"}],\"name\":\"RegisteredValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Rewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"COMMISSION_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"changeCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"computeCommittee\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAccount\",\"type\":\"address\"}],\"internalType\":\"structAutonity.Policy\",\"name\":\"policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIAccountability\",\"name\":\"accountabilityContract\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"oracleContract\",\"type\":\"address\"},{\"internalType\":\"contractIACU\",\"name\":\"acuContract\",\"type\":\"address\"},{\"internalType\":\"contractISupplyControl\",\"name\":\"supplyControlContract\",\"type\":\"address\"},{\"internalType\":\"contractIStabilization\",\"name\":\"stabilizationContract\",\"type\":\"address\"}],\"internalType\":\"structAutonity.Contracts\",\"name\":\"contracts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committeeSize\",\"type\":\"uint256\"}],\"internalType\":\"structAutonity.Protocol\",\"name\":\"protocol\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochTotalBondedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"structAutonity.CommitteeMember[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeInitialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPower\",\"type\":\"uint256\"}],\"internalType\":\"structAutonity.CommitteeMember[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitteeEnodes\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"getEpochFromBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastEpochBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxCommitteeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewContract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfBondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"contractLiquid\",\"name\":\"liquidContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSlashed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailReleaseBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provableFaultCount\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structAutonity.Validator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastEpochBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"pauseValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_enode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_multisig\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAccountability\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAccountabilityContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIACU\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAcuContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setCommitteeSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setMinimumBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"setOperatorAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setOracleContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStabilization\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setStabilizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISupplyControl\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSupplyControlContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"setTreasuryAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setUnbondingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRedistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfBondedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfUnbondingShares\",\"type\":\"uint256\"},{\"internalType\":\"contractLiquid\",\"name\":\"liquidContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"registrationBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSlashed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailReleaseBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provableFaultCount\",\"type\":\"uint256\"},{\"internalType\":\"enumValidatorState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structAutonity.Validator\",\"name\":\"_val\",\"type\":\"tuple\"}],\"name\":\"updateValidatorAndTransferSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytecode\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"2f2c3f2e": "COMMISSION_RATE_PRECISION()",
		"b46e5520": "activateValidator(address)",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"a515366a": "bond(address,uint256)",
		"9dc29fac": "burn(address,uint256)",
		"852c4849": "changeCommissionRate(address,uint256)",
		"872cf059": "completeContractUpgrade()",
		"ae1f5fa0": "computeCommittee()",
		"79502c55": "config()",
		"313ce567": "decimals()",
		"d5f39488": "deployer()",
		"c9d97af4": "epochID()",
		"1604e416": "epochReward()",
		"9c98e471": "epochTotalBondedStake()",
		"4bb278f3": "finalize()",
		"d861b0e8": "finalizeInitialization()",
		"43645969": "getBlockPeriod()",
		"ab8f6ffe": "getCommittee()",
		"a8b2216e": "getCommitteeEnodes()",
		"96b477cb": "getEpochFromBlock(uint256)",
		"dfb1a4d2": "getEpochPeriod()",
		"731b3a03": "getLastEpochBlock()",
		"819b6463": "getMaxCommitteeSize()",
		"11220633": "getMinimumBaseFee()",
		"b66b3e79": "getNewContract()",
		"e7f43c68": "getOperator()",
		"833b1fce": "getOracle()",
		"5f7d3949": "getProposer(uint256,uint256)",
		"f7866ee3": "getTreasuryAccount()",
		"29070c6d": "getTreasuryFee()",
		"6fd2c80b": "getUnbondingPeriod()",
		"1904bb2e": "getValidator(address)",
		"b7ab4db5": "getValidators()",
		"0d8e6e2c": "getVersion()",
		"c2362dd5": "lastEpochBlock()",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"0ae65e7a": "pauseValidator(address)",
		"ad722d4d": "registerValidator(string,address,bytes)",
		"cf9c5719": "resetContractUpgrade()",
		"1250a28d": "setAccountabilityContract(address)",
		"d372c07e": "setAcuContract(address)",
		"8bac7dad": "setCommitteeSize(uint256)",
		"6b5f444c": "setEpochPeriod(uint256)",
		"cb696f54": "setMinimumBaseFee(uint256)",
		"520fdbbc": "setOperatorAccount(address)",
		"496ccd9b": "setOracleContract(address)",
		"cfd19fb9": "setStabilizationContract(address)",
		"b3ecbadd": "setSupplyControlContract(address)",
		"d886f8a2": "setTreasuryAccount(address)",
		"77e741c7": "setTreasuryFee(uint256)",
		"114eaf55": "setUnbondingPeriod(uint256)",
		"95d89b41": "symbol()",
		"9bb851c0": "totalRedistributed()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"a5d059ca": "unbond(address,uint256)",
		"6e74c554": "updateValidatorAndTransferSlashedFunds((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))",
		"b2ea9adb": "upgradeContract(bytes,string)",
	},
	Bin: "0x60806040525f600b555f600c5534801562000018575f80fd5b5060405162009ae938038062009ae98339810160408190526200003b9162000d7b565b601b545f036200006357602980546001600160a01b031916331790556200006382826200006b565b5050620011ec565b80518051600d55602080820151600e55604080830151600f55606080840151601055608093840151601180546001600160a01b03199081166001600160a01b039384161790915584870151805160128054841691851691909117905580860151601380548416918516919091179055808501516014805484169185169190911790558084015160158054841691851691909117905590950151601680548716918316919091179055828601518051601780549097169216919091179094559183015160185582015160195590810151601a55810151601b555f5b8251811015620003a5575f83828151811062000165576200016562000f7b565b602002602001015160a0015190505f84838151811062000189576200018962000f7b565b60200260200101516101800181815250505f848381518110620001b057620001b062000f7b565b602002602001015161016001906001600160a01b031690816001600160a01b0316815250505f848381518110620001eb57620001eb62000f7b565b602002602001015160a00181815250505f84838151811062000211576200021162000f7b565b60209081029190910101516101a00152600f5484518590849081106200023b576200023b62000f7b565b602002602001015160800181815250505f84838151811062000261576200026162000f7b565b60200260200101516102200190600281111562000282576200028262000f8f565b9081600281111562000298576200029862000f8f565b81525050620002c9848381518110620002b557620002b562000f7b565b6020026020010151620003aa60201b60201c565b8060265f868581518110620002e257620002e262000f7b565b60200260200101515f01516001600160a01b03166001600160a01b031681526020019081526020015f205f8282546200031c919062000fb7565b925050819055508060285f82825462000336919062000fb7565b925050819055506200038f84838151811062000356576200035662000f7b565b6020026020010151602001518286858151811062000378576200037862000f7b565b60200260200101515f0151620003c360201b60201c565b50806200039c8162000fd3565b91505062000145565b505050565b620003b581620005ad565b620003c081620006e2565b50565b5f8211620004245760405162461bcd60e51b815260206004820152602360248201527f616d6f756e74206e65656420746f206265207374726963746c7920706f73697460448201526269766560e81b60648201526084015b60405180910390fd5b6001600160a01b0381165f908152602660205260409020548211156200048d5760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e74204e6577746f6e2062616c616e6365000000000060448201526064016200041b565b6001600160a01b0381165f9081526026602052604081208054849290620004b690849062000fee565b9091555050604080516080810182526001600160a01b03808416825285811660208084019182528385018781524360608601908152600580545f908152600394859052978820875181549088166001600160a01b0319918216178255955160018201805491909816961695909517909555905160028401555191015580549192620005418362000fd3565b90915550506001600160a01b038481165f81815260276020908152604091829020548251908516948716948514808252918101889052909392917fc46aaee12f38035617ad448c04a7956119f7c7ed395ecc347b898817451ddb8d910160405180910390a35050505050565b5f620005c38260600151620008ea60201b60201c565b6001600160a01b03909116602084015290508015620006135760405162461bcd60e51b815260206004820152600b60248201526a32b737b2329032b93937b960a91b60448201526064016200041b565b6020808301516001600160a01b039081165f90815260279092526040909120600101541615620006865760405162461bcd60e51b815260206004820152601c60248201527f76616c696461746f7220616c726561647920726567697374657265640000000060448201526064016200041b565b61271082608001511115620006de5760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420636f6d6d697373696f6e207261746500000000000000000060448201526064016200041b565b5050565b6101608101516001600160a01b03166200076057601c545f90620007069062000930565b90508160200151825f0151836080015183604051620007259062000a4c565b62000734949392919062001004565b604051809103905ff0801580156200074e573d5f803e3d5ffd5b506001600160a01b0316610160830152505b60208181018051601c8054600180820183555f9283527f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a21190910180546001600160a01b03199081166001600160a01b0395861617909155845184168352602790955260409182902086518154871690851617815593519084018054861691841691909117905584015160028301805490941691161790915560608201518291906003820190620008119082620010e4565b506080820151600482015560a0820151600582015560c0820151600682015560e0820151600782015561010082015160088201556101208201516009820155610140820151600a820155610160820151600b820180546001600160a01b0319166001600160a01b03909216919091179055610180820151600c8201556101a0820151600d8201556101c0820151600e8201556101e0820151600f820155610200820151601082015561022082015160118201805460ff19166001836002811115620008e057620008e062000f8f565b0217905550505050565b5f80620008f662000a5a565b5f60408286516020880160ff5afa6200090d575f80fd5b5080516020909101516c0100000000000000000000000090910494909350915050565b6060815f03620009575750506040805180820190915260018152600360fc1b602082015290565b815f5b81156200098657806200096d8162000fd3565b91506200097e9050600a83620011c0565b91506200095a565b5f816001600160401b03811115620009a257620009a262000a78565b6040519080825280601f01601f191660200182016040528015620009cd576020820181803683370190505b5090505b841562000a4457620009e560018362000fee565b9150620009f4600a86620011d6565b62000a0190603062000fb7565b60f81b81838151811062000a195762000a1962000f7b565b60200101906001600160f81b03191690815f1a90535062000a3c600a86620011c0565b9450620009d1565b949350505050565b6114b4806200863583390190565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b60405160a081016001600160401b038111828210171562000ab15762000ab162000a78565b60405290565b604051608081016001600160401b038111828210171562000ab15762000ab162000a78565b60405161024081016001600160401b038111828210171562000ab15762000ab162000a78565b604051601f8201601f191681016001600160401b038111828210171562000b2d5762000b2d62000a78565b604052919050565b6001600160a01b0381168114620003c0575f80fd5b805162000b578162000b35565b919050565b5f5b8381101562000b7857818101518382015260200162000b5e565b50505f910152565b5f82601f83011262000b90575f80fd5b81516001600160401b0381111562000bac5762000bac62000a78565b62000bc1601f8201601f191660200162000b02565b81815284602083860101111562000bd6575f80fd5b62000a4482602083016020870162000b5c565b80516003811062000b57575f80fd5b5f60a0828403121562000c09575f80fd5b62000c1362000a8c565b9050815162000c228162000b35565b8152602082015162000c348162000b35565b6020820152604082015162000c498162000b35565b6040820152606082015162000c5e8162000b35565b6060820152608082015162000c738162000b35565b608082015292915050565b5f6080828403121562000c8f575f80fd5b62000c9962000ab7565b9050815162000ca88162000b35565b8082525060208201516020820152604082015160408201526060820151606082015292915050565b5f8183036101e081121562000ce3575f80fd5b62000ced62000ab7565b915060a081121562000cfd575f80fd5b5062000d0862000a8c565b82518152602083015160208201526040830151604082015260608301516060820152608083015162000d3a8162000b35565b6080820152815262000d508360a0840162000bf8565b602082015262000d6583610140840162000c7e565b60408201526101c0820151606082015292915050565b5f8061020080848603121562000d8f575f80fd5b83516001600160401b038082111562000da6575f80fd5b818601915086601f83011262000dba575f80fd5b815160208282111562000dd15762000dd162000a78565b8160051b62000de282820162000b02565b928352848101820192828101908b85111562000dfc575f80fd5b83870192505b8483101562000f595782518681111562000e1a575f80fd5b8701610240818e03601f1901121562000e31575f80fd5b62000e3b62000adc565b62000e4886830162000b4a565b815262000e586040830162000b4a565b8682015262000e6a6060830162000b4a565b604082015260808201518881111562000e81575f80fd5b62000e918f888386010162000b80565b60608301525060a0820151608082015260c082015160a082015260e082015160c08201526101008083015160e0830152610120808401518284015261014091508184015181840152506101608084015182840152610180915062000ef782850162000b4a565b81840152506101a080840151828401526101c091508184015181840152506101e080840151828401528b840151818401525050610220808301518b83015262000f44610240840162000be9565b90820152835250918301919083019062000e02565b80995050505062000f6d89828a0162000cd0565b955050505050509250929050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111562000fcd5762000fcd62000fa3565b92915050565b5f6001820162000fe75762000fe762000fa3565b5060010190565b8181038181111562000fcd5762000fcd62000fa3565b5f60018060a01b038087168352808616602084015250836040830152608060608301528251806080840152620010428160a085016020870162000b5c565b601f01601f19169190910160a00195945050505050565b600181811c908216806200106e57607f821691505b6020821081036200108d57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620003a5575f81815260208120601f850160051c81016020861015620010bb5750805b601f850160051c820191505b81811015620010dc57828155600101620010c7565b505050505050565b81516001600160401b0381111562001100576200110062000a78565b620011188162001111845462001059565b8462001093565b602080601f8311600181146200114e575f8415620011365750858301515b5f19600386901b1c1916600185901b178555620010dc565b5f85815260208120601f198616915b828110156200117e578886015182559484019460019091019084016200115d565b50858210156200119c57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601260045260245ffd5b5f82620011d157620011d1620011ac565b500490565b5f82620011e757620011e7620011ac565b500690565b61743b80620011fa5f395ff3fe608060405260043610620003ce575f3560e01c80638bac7dad11620001f6578063b46e55201162000116578063d372c07e11620000a6578063dd62ed3e1162000074578063dd62ed3e1462000cd7578063dfb1a4d21462000d1f578063e7f43c681462000d35578063f7866ee31462000d5457005b8063d372c07e1462000c57578063d5f394881462000c7b578063d861b0e81462000c9c578063d886f8a21462000cb357005b8063c9d97af411620000e4578063c9d97af41462000be1578063cb696f541462000bf8578063cf9c57191462000c1c578063cfd19fb91462000c3357005b8063b46e55201462000b68578063b66b3e791462000b8c578063b7ab4db51462000bb3578063c2362dd51462000bca57005b8063a5d059ca1162000192578063ad722d4d1162000160578063ad722d4d1462000ad6578063ae1f5fa01462000afa578063b2ea9adb1462000b20578063b3ecbadd1462000b4457005b8063a5d059ca1462000a42578063a8b2216e1462000a66578063a9059cbb1462000a8c578063ab8f6ffe1462000ab057005b80639bb851c011620001d05780639bb851c014620009cc5780639c98e47114620009e35780639dc29fac14620009fa578063a515366a1462000a1e57005b80638bac7dad146200094c57806395d89b41146200097057806396b477cb146200099d57005b80634364596911620002ee57806370a08231116200027e578063819b6463116200024c578063819b646314620008dc578063833b1fce14620008f2578063852c48491462000911578063872cf059146200093557005b806370a082311462000704578063731b3a03146200073c57806377e741c7146200075257806379502c55146200077657005b80635f7d394911620002bc5780635f7d394914620006695780636b5f444c14620006a65780636e74c55414620006ca5780636fd2c80b14620006ee57005b80634364596914620005e4578063496ccd9b14620005fa5780634bb278f3146200061e578063520fdbbc146200064557005b80631604e416116200036a57806329070c6d116200033857806329070c6d14620005765780632f2c3f2e146200058c578063313ce56714620005a357806340c10f1914620005c057005b80631604e41614620004f257806318160ddd14620005095780631904bb2e146200051f57806323b872dd146200055257005b80630d8e6e2c11620003a85780630d8e6e2c1462000474578063112206331462000494578063114eaf5514620004aa5780631250a28d14620004ce57005b806306fdde0314620003d8578063095ea7b3146200041b5780630ae65e7a146200045057005b36620003d657005b005b348015620003e4575f80fd5b506040805180820190915260068152652732bbba37b760d11b60208201525b604051620004129190620051ad565b60405180910390f35b34801562000427575f80fd5b506200043f62000439366004620051dd565b62000d73565b604051901515815260200162000412565b3480156200045c575f80fd5b50620003d66200046e3660046200520a565b62000d8b565b34801562000480575f80fd5b50601b545b60405190815260200162000412565b348015620004a0575f80fd5b50600e5462000485565b348015620004b6575f80fd5b50620003d6620004c836600462005228565b62000e1c565b348015620004da575f80fd5b50620003d6620004ec3660046200520a565b62000e4e565b348015620004fe575f80fd5b506200048560235481565b34801562000515575f80fd5b5060285462000485565b3480156200052b575f80fd5b50620005436200053d3660046200520a565b62000e9d565b60405162000412919062005275565b3480156200055e575f80fd5b506200043f62000570366004620053b7565b6200109e565b34801562000582575f80fd5b50600d5462000485565b34801562000598575f80fd5b506200048561271081565b348015620005af575f80fd5b506040516012815260200162000412565b348015620005cc575f80fd5b50620003d6620005de366004620051dd565b620010f6565b348015620005f0575f80fd5b5060195462000485565b34801562000606575f80fd5b50620003d6620006183660046200520a565b620011ae565b3480156200062a575f80fd5b5062000635620012ae565b6040516200041292919062005449565b34801562000651575f80fd5b50620003d6620006633660046200520a565b620015bb565b34801562000675575f80fd5b506200068d6200068736600462005465565b62001746565b6040516001600160a01b03909116815260200162000412565b348015620006b2575f80fd5b50620003d6620006c436600462005228565b62001952565b348015620006d6575f80fd5b50620003d6620006e836600462005486565b62001ad3565b348015620006fa575f80fd5b5060105462000485565b34801562000710575f80fd5b5062000485620007223660046200520a565b6001600160a01b03165f9081526026602052604090205490565b34801562000748575f80fd5b50601f5462000485565b3480156200075e575f80fd5b50620003d66200077036600462005228565b62001ca6565b34801562000782575f80fd5b506040805160a08082018352600d548252600e54602080840191909152600f54838501526010546060808501919091526011546001600160a01b0390811660808087019190915286519485018752601254821685526013548216858501526014548216858801526015548216858401526016548216858201528651908101875260175490911681526018549281019290925260195494820194909452601a5493810193909352601b5462000834939084565b6040805185518152602080870151818301528683015182840152606080880151818401526080978801516001600160a01b03908116898501528751811660a085015282880151811660c085015287850151811660e0850152818801518116610100850152979096015187166101208301528451909616610140820152948301516101608601528201516101808501529101516101a08301526101c08201526101e00162000412565b348015620008e8575f80fd5b50601a5462000485565b348015620008fe575f80fd5b506013546001600160a01b03166200068d565b3480156200091d575f80fd5b50620003d66200092f366004620051dd565b62001cd8565b34801562000941575f80fd5b50620003d662001e66565b34801562000958575f80fd5b50620003d66200096a36600462005228565b62001ea2565b3480156200097c575f80fd5b50604080518082019091526003815262272a2760e91b602082015262000403565b348015620009a9575f80fd5b5062000485620009bb36600462005228565b5f908152601e602052604090205490565b348015620009d8575f80fd5b506200048560225481565b348015620009ef575f80fd5b506200048560205481565b34801562000a06575f80fd5b50620003d662000a18366004620051dd565b62001f25565b34801562000a2a575f80fd5b50620003d662000a3c366004620051dd565b62002038565b34801562000a4e575f80fd5b50620003d662000a60366004620051dd565b62002109565b34801562000a72575f80fd5b5062000a7d62002155565b604051620004129190620054c1565b34801562000a98575f80fd5b506200043f62000aaa366004620051dd565b62002233565b34801562000abc575f80fd5b5062000ac762002241565b60405162000412919062005525565b34801562000ae2575f80fd5b50620003d662000af4366004620055de565b620022ad565b34801562000b06575f80fd5b5062000b11620023bf565b6040516200041291906200565a565b34801562000b2c575f80fd5b50620003d662000b3e366004620056a8565b62002b9a565b34801562000b50575f80fd5b50620003d662000b623660046200520a565b62002be0565b34801562000b74575f80fd5b50620003d662000b863660046200520a565b62002c2f565b34801562000b98575f80fd5b5062000ba362002e42565b604051620004129291906200570f565b34801562000bbf575f80fd5b5062000b1162002f74565b34801562000bd6575f80fd5b5062000485601f5481565b34801562000bed575f80fd5b5062000485601d5481565b34801562000c04575f80fd5b50620003d662000c1636600462005228565b62002fd6565b34801562000c28575f80fd5b50620003d662003039565b34801562000c3f575f80fd5b50620003d662000c513660046200520a565b6200308b565b34801562000c63575f80fd5b50620003d662000c753660046200520a565b620030da565b34801562000c87575f80fd5b506029546200068d906001600160a01b031681565b34801562000ca8575f80fd5b50620003d662003129565b34801562000cbf575f80fd5b50620003d662000cd13660046200520a565b6200316a565b34801562000ce3575f80fd5b506200048562000cf536600462005740565b6001600160a01b039182165f90815260256020908152604080832093909416825291909152205490565b34801562000d2b575f80fd5b5060185462000485565b34801562000d41575f80fd5b506017546001600160a01b03166200068d565b34801562000d60575f80fd5b506011546001600160a01b03166200068d565b5f62000d81338484620031b9565b5060015b92915050565b6001600160a01b038082165f818152602760205260409020600101549091161462000dd35760405162461bcd60e51b815260040162000dca906200577c565b60405180910390fd5b6001600160a01b038181165f9081526027602052604090205416331462000e0e5760405162461bcd60e51b815260040162000dca90620057b3565b62000e1981620032e1565b50565b6017546001600160a01b0316331462000e495760405162461bcd60e51b815260040162000dca90620057ff565b601055565b6017546001600160a01b0316331462000e7b5760405162461bcd60e51b815260040162000dca90620057ff565b601280546001600160a01b0319166001600160a01b0392909216919091179055565b62000ea762004fa5565b6001600160a01b038083165f818152602760205260409020600101549091161462000ee65760405162461bcd60e51b815260040162000dca9062005836565b6001600160a01b038083165f90815260276020908152604091829020825161024081018452815485168152600182015485169281019290925260028101549093169181019190915260038201805491929160608401919062000f48906200586d565b80601f016020809104026020016040519081016040528092919081815260200182805462000f76906200586d565b801562000fc55780601f1062000f9b5761010080835404028352916020019162000fc5565b820191905f5260205f20905b81548152906001019060200180831162000fa757829003601f168201915b505050918352505060048201546020820152600582015460408201526006820154606082015260078201546080820152600882015460a0820152600982015460c0820152600a82015460e0820152600b8201546001600160a01b0316610100820152600c820154610120820152600d820154610140820152600e820154610160820152600f82015461018082015260108201546101a082015260118201546101c09091019060ff16600281111562001081576200108162005240565b600281111562001095576200109562005240565b90525092915050565b5f620010ac848484620033b7565b6001600160a01b0384165f908152602560209081526040808320338452909152812054620010dc908490620058bb565b9050620010eb853383620031b9565b506001949350505050565b6017546001600160a01b03163314620011235760405162461bcd60e51b815260040162000dca90620057ff565b6001600160a01b0382165f90815260266020526040812080548392906200114c908490620058d1565b925050819055508060285f828254620011669190620058d1565b90915550506040518181526001600160a01b038316907f48490b4407bb949b708ec5f514b4167f08f4969baaf78d53b05028adf369bfcf906020015b60405180910390a25050565b6017546001600160a01b03163314620011db5760405162461bcd60e51b815260040162000dca90620057ff565b601380546001600160a01b0319166001600160a01b03838116918217909255601454604051637adbf97360e01b8152600481019290925290911690637adbf973906024015f604051808303815f87803b15801562001237575f80fd5b505af11580156200124a573d5f803e3d5ffd5b5050601654604051637adbf97360e01b81526001600160a01b0385811660048301529091169250637adbf97391506024015b5f604051808303815f87803b15801562001294575f80fd5b505af1158015620012a7573d5f803e3d5ffd5b5050505050565b6029545f906060906001600160a01b03163314620012e05760405162461bcd60e51b815260040162000dca90620058e7565b601d54435f818152601e6020526040812092909255601854601f54620013079190620058d1565b6012546040516306c9789b60e41b8152929091146004830181905292506001600160a01b031690636c9789b0906024015f604051808303815f87803b1580156200134f575f80fd5b505af115801562001362573d5f803e3d5ffd5b505050508015620014505762001377620034bd565b6200138162003879565b6200138b62003967565b5f62001396620023bf565b60135460405163422811f960e11b81529192506001600160a01b03169063845023f290620013c99084906004016200565a565b5f604051808303815f87803b158015620013e1575f80fd5b505af1158015620013f4573d5f803e3d5ffd5b5050505043601f819055506001601d5f828254620014139190620058d1565b9091555050601d546040519081527febad8099c467528a56c98b63c8d476d251cf1ffb4c75db94b4d23fa2b6a1e3359060200160405180910390a1505b60135460408051634bb278f360e01b815290515f926001600160a01b031691634bb278f3916004808301926020929190829003018187875af115801562001499573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620014bf91906200592a565b905080156200153a576014546040805163a2e6204560e01b815290516001600160a01b039092169163a2e620459160048082019260209290919082900301815f875af192505050801562001532575060408051601f3d908101601f191682019092526200152f918101906200592a565b60015b156200153a57505b600254602180546040805160208084028201810190925282815260ff909416939183915f9084015b82821015620015ab575f848152602090819020604080518082019091526002850290910180546001600160a01b0316825260019081015482840152908352909201910162001562565b5050505090509350935050509091565b6017546001600160a01b03163314620015e85760405162461bcd60e51b815260040162000dca90620057ff565b601780546001600160a01b0319166001600160a01b0383811691821790925560135460405163b3ab15fb60e01b815260048101929092529091169063b3ab15fb906024015f604051808303815f87803b15801562001644575f80fd5b505af115801562001657573d5f803e3d5ffd5b505060145460405163b3ab15fb60e01b81526001600160a01b038581166004830152909116925063b3ab15fb91506024015f604051808303815f87803b158015620016a0575f80fd5b505af1158015620016b3573d5f803e3d5ffd5b505060155460405163b3ab15fb60e01b81526001600160a01b038581166004830152909116925063b3ab15fb91506024015f604051808303815f87803b158015620016fc575f80fd5b505af11580156200170f573d5f803e3d5ffd5b505060165460405163b3ab15fb60e01b81526001600160a01b038581166004830152909116925063b3ab15fb91506024016200127c565b5f80805b6021548110156200179f57602181815481106200176b576200176b6200594b565b905f5260205f2090600202016001015482620017889190620058d1565b91508062001796816200595f565b9150506200174a565b50805f03620017f15760405162461bcd60e51b815260206004820152601c60248201527f54686520636f6d6d6974746565206973206e6f74207374616b696e6700000000604482015260640162000dca565b5f83620018006063876200597a565b6200180c9190620058d1565b90505f816040516020016200182391815260200190565b60408051601f19818403018152919052805160209091012090505f6200184a8483620059a8565b90505f805b602154811015620018f657602181815481106200187057620018706200594b565b905f5260205f20906002020160010154826200188d9190620058d1565b91506200189c600183620058bb565b8311620018e15760218181548110620018b957620018b96200594b565b5f9182526020909120600290910201546001600160a01b0316965062000d8595505050505050565b80620018ed816200595f565b9150506200184f565b5060405162461bcd60e51b815260206004820152602960248201527f5468657265206973206e6f2076616c696461746f72206c65667420696e20746860448201526865206e6574776f726b60b81b606482015260840162000dca565b6017546001600160a01b031633146200197f5760405162461bcd60e51b815260040162000dca90620057ff565b60185481101562001a365780601f546200199a9190620058d1565b431062001a365760405162461bcd60e51b815260206004820152605760248201527f63757272656e7420636861696e2068656164206578636565642074686520776960448201527f6e646f773a206c617374426c6f636b45706f6368202b205f6e6577506572696f60648201527f642c2074727920616761696e206c6174746572206f6e2e000000000000000000608482015260a40162000dca565b6018819055601254604051631ad7d11360e21b8152600481018390526001600160a01b0390911690636b5f444c906024015f604051808303815f87803b15801562001a7f575f80fd5b505af115801562001a92573d5f803e3d5ffd5b505050507fd7f1279ded354dbf22a69fcc2fd661763a6e2956a5d2891af9410af880fa5f818160405162001ac891815260200190565b60405180910390a150565b6012546001600160a01b0316331462001b3b5760405162461bcd60e51b815260206004820152602360248201527f63616c6c6572206973206e6f742074686520736c617368696e6720636f6e74726044820152621858dd60ea1b606482015260840162000dca565b5f61012082013560278262001b5760408601602087016200520a565b6001600160a01b03166001600160a01b031681526020019081526020015f206009015462001b869190620058bb565b60c083013560275f62001ba060408701602088016200520a565b6001600160a01b03166001600160a01b031681526020019081526020015f206006015462001bcf9190620058bb565b60a084013560275f62001be960408801602089016200520a565b6001600160a01b03166001600160a01b031681526020019081526020015f206005015462001c189190620058bb565b62001c249190620058d1565b62001c309190620058d1565b6011546001600160a01b03165f9081526026602052604081208054929350839290919062001c60908490620058d1565b9091555082905060275f62001c7c60408401602085016200520a565b6001600160a01b0316815260208101919091526040015f2062001ca0828262005b8a565b50505050565b6017546001600160a01b0316331462001cd35760405162461bcd60e51b815260040162000dca90620057ff565b600d55565b6001600160a01b038083165f818152602760205260409020600101549091161462001d175760405162461bcd60e51b815260040162000dca906200577c565b6001600160a01b038281165f9081526027602052604090205416331462001d525760405162461bcd60e51b815260040162000dca90620057b3565b61271081111562001da65760405162461bcd60e51b815260206004820152601f60248201527f7265717569726520636f727265637420636f6d6d697373696f6e207261746500604482015260640162000dca565b604080516060810182526001600160a01b038481168252436020808401918252838501868152600c80545f908152600a909352958220855181546001600160a01b03191695169490941784559151600180850191909155915160029093019290925583549293909290919062001e1e908490620058d1565b90915550506040518281526001600160a01b038416907f4fba51c92fa3d6ad8374d394f6cd5766857552e153d7384a8f23aa4ce9a8a7cf9060200160405180910390a2505050565b6017546001600160a01b0316331462001e935760405162461bcd60e51b815260040162000dca90620057ff565b6002805460ff19166001179055565b6017546001600160a01b0316331462001ecf5760405162461bcd60e51b815260040162000dca90620057ff565b5f811162001f205760405162461bcd60e51b815260206004820152601960248201527f636f6d6d69747465652073697a652063616e2774206265203000000000000000604482015260640162000dca565b601a55565b6017546001600160a01b0316331462001f525760405162461bcd60e51b815260040162000dca90620057ff565b6001600160a01b0382165f9081526026602052604090205481111562001fb45760405162461bcd60e51b8152602060048201526016602482015275416d6f756e7420657863656564732062616c616e636560501b604482015260640162000dca565b6001600160a01b0382165f908152602660205260408120805483929062001fdd908490620058bb565b925050819055508060285f82825462001ff79190620058bb565b90915550506040518181526001600160a01b038316907f5024dbeedf0c06664c9bd7be836915730c955e936972c020683dadf11d5488a390602001620011a2565b6001600160a01b038083165f8181526027602052604090206001015490911614620020775760405162461bcd60e51b815260040162000dca9062005836565b6001600160a01b0382165f9081526027602052604081206011015460ff166002811115620020a957620020a962005240565b14620020f85760405162461bcd60e51b815260206004820152601b60248201527f76616c696461746f72206e65656420746f206265206163746976650000000000604482015260640162000dca565b6200210582823362003a7e565b5050565b6001600160a01b038083165f8181526027602052604090206001015490911614620021485760405162461bcd60e51b815260040162000dca9062005836565b6200210582823362003c65565b60606024805480602002602001604051908101604052809291908181526020015f905b828210156200222a578382905f5260205f2001805462002198906200586d565b80601f0160208091040260200160405190810160405280929190818152602001828054620021c6906200586d565b8015620022155780601f10620021eb5761010080835404028352916020019162002215565b820191905f5260205f20905b815481529060010190602001808311620021f757829003601f168201915b50505050508152602001906001019062002178565b50505050905090565b5f62000d81338484620033b7565b60606021805480602002602001604051908101604052809291908181526020015f905b828210156200222a575f848152602090819020604080518082019091526002850290910180546001600160a01b0316825260019081015482840152908352909201910162002264565b5f604051806102400160405280336001600160a01b031681526020015f6001600160a01b03168152602001846001600160a01b03168152602001858152602001600d5f016002015481526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f6001600160a01b031681526020015f81526020014381526020015f81526020015f81526020015f81526020015f600281111562002360576200236062005240565b9052905062002370818362003f61565b60208101516101608201516040517f8ad8bd2eb6950e5f332fd3a6dca48cb358ecfe3057848902b98cbdfe455c915c92620023b192339288918a9162005cb3565b60405180910390a150505050565b6029546060906001600160a01b03163314620023ef5760405162461bcd60e51b815260040162000dca90620058e7565b601c54620024405760405162461bcd60e51b815260206004820152601860248201527f5468657265206d7573742062652076616c696461746f72730000000000000000604482015260640162000dca565b5f805b601c548110156200251d575f60275f601c84815481106200246857620024686200594b565b5f9182526020808320909101546001600160a01b0316835282019290925260400190206011015460ff166002811115620024a657620024a662005240565b148015620024f257505f60275f601c8481548110620024c957620024c96200594b565b5f9182526020808320909101546001600160a01b03168352820192909252604001902060050154115b1562002508578162002504816200595f565b9250505b8062002514816200595f565b91505062002443565b50601a548181106200252c5750805b5f8267ffffffffffffffff81111562002549576200254962005539565b6040519080825280602002602001820160405280156200258657816020015b6200257262004fa5565b815260200190600190039081620025685790505b5090505f8267ffffffffffffffff811115620025a657620025a662005539565b604051908082528060200260200182016040528015620025e357816020015b620025cf62004fa5565b815260200190600190039081620025c55790505b5090505f8367ffffffffffffffff81111562002603576200260362005539565b6040519080825280602002602001820160405280156200262d578160200160208202803683370190505b5090505f805b601c548110156200290c575f60275f601c84815481106200265857620026586200594b565b5f9182526020808320909101546001600160a01b0316835282019290925260400190206011015460ff16600281111562002696576200269662005240565b148015620026e257505f60275f601c8481548110620026b957620026b96200594b565b5f9182526020808320909101546001600160a01b03168352820192909252604001902060050154115b15620028f7575f60275f601c84815481106200270257620027026200594b565b5f9182526020808320909101546001600160a01b039081168452838201949094526040928301909120825161024081018452815485168152600182015485169281019290925260028101549093169181019190915260038201805491929160608401919062002771906200586d565b80601f01602080910402602001604051908101604052809291908181526020018280546200279f906200586d565b8015620027ee5780601f10620027c457610100808354040283529160200191620027ee565b820191905f5260205f20905b815481529060010190602001808311620027d057829003601f168201915b505050918352505060048201546020820152600582015460408201526006820154606082015260078201546080820152600882015460a0820152600982015460c0820152600a82015460e0820152600b8201546001600160a01b0316610100820152600c820154610120820152600d820154610140820152600e820154610160820152600f82015461018082015260108201546101a082015260118201546101c09091019060ff166002811115620028aa57620028aa62005240565b6002811115620028be57620028be62005240565b81525050905080868481518110620028da57620028da6200594b565b60200260200101819052508280620028f2906200595f565b935050505b8062002903816200595f565b91505062002633565b50601a54845111156200298b57620029248462004297565b5f5b601a5481101562002984578481815181106200294657620029466200594b565b60200260200101518482815181106200296357620029636200594b565b602002602001018190525080806200297b906200595f565b91505062002926565b506200298f565b8392505b6200299c60215f62005056565b620029a960245f62005076565b5f60208190555b8581101562002b8e575f6040518060400160405280868481518110620029da57620029da6200594b565b6020026020010151602001516001600160a01b0316815260200186848151811062002a095762002a096200594b565b60209081029190910181015160a00151909152602180546001810182555f9190915282517f3a6357012c1a3ae0a17d304c9920310382d968ebcc4b1771f41c6b304205b570600290920291820180546001600160a01b0319166001600160a01b03909216919091179055908201517f3a6357012c1a3ae0a17d304c9920310382d968ebcc4b1771f41c6b304205b57190910155855190915060249086908490811062002ab95762002ab96200594b565b6020908102919091018101516060015182546001810184555f93845291909220019062002ae7908262005cfb565b5084828151811062002afd5762002afd6200594b565b60200260200101516040015184838151811062002b1e5762002b1e6200594b565b60200260200101906001600160a01b031690816001600160a01b03168152505084828151811062002b535762002b536200594b565b602002602001015160a0015160205f82825462002b719190620058d1565b9091555082915062002b859050816200595f565b915050620029b0565b50909550505050505090565b6017546001600160a01b0316331462002bc75760405162461bcd60e51b815260040162000dca90620057ff565b62002bd35f83620042b3565b62002105600182620042b3565b6017546001600160a01b0316331462002c0d5760405162461bcd60e51b815260040162000dca90620057ff565b601580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038082165f818152602760205260409020600101549091161462002c6e5760405162461bcd60e51b815260040162000dca906200577c565b6001600160a01b038082165f9081526027602052604090208054909116331462002cf35760405162461bcd60e51b815260206004820152602f60248201527f726571756972652063616c6c657220746f2062652076616c696461746f72207460448201526e1c99585cdd5c9e481858d8dbdd5b9d608a1b606482015260840162000dca565b5f601182015460ff16600281111562002d105762002d1062005240565b0362002d5f5760405162461bcd60e51b815260206004820152601860248201527f76616c696461746f7220616c7265616479206163746976650000000000000000604482015260640162000dca565b6002601182015460ff16600281111562002d7d5762002d7d62005240565b14801562002d8e57504381600f0154115b1562002ddd5760405162461bcd60e51b815260206004820152601760248201527f76616c696461746f72207374696c6c20696e206a61696c000000000000000000604482015260640162000dca565b60118101805460ff191690558054601854601f546001600160a01b038581169316917f60fcbf2d07dc712a93e59fb28f1edb626d7c2497c57ba71a8c0b3999ecb9a3b59162002e2d9190620058d1565b60405190815260200160405180910390a35050565b6060805f600181805462002e56906200586d565b80601f016020809104026020016040519081016040528092919081815260200182805462002e84906200586d565b801562002ed35780601f1062002ea95761010080835404028352916020019162002ed3565b820191905f5260205f20905b81548152906001019060200180831162002eb557829003601f168201915b5050505050915080805462002ee8906200586d565b80601f016020809104026020016040519081016040528092919081815260200182805462002f16906200586d565b801562002f655780601f1062002f3b5761010080835404028352916020019162002f65565b820191905f5260205f20905b81548152906001019060200180831162002f4757829003601f168201915b50505050509050915091509091565b6060601c80548060200260200160405190810160405280929190818152602001828054801562002fcc57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831162002fad575b5050505050905090565b6017546001600160a01b03163314620030035760405162461bcd60e51b815260040162000dca90620057ff565b600e8190556040518181527f1f4d2fc7529047a5bd96d3229bfea127fd18b7748f13586e097c69fccd3891289060200162001ac8565b6017546001600160a01b03163314620030665760405162461bcd60e51b815260040162000dca90620057ff565b620030725f8062005093565b6200307f60015f62005093565b6002805460ff19169055565b6017546001600160a01b03163314620030b85760405162461bcd60e51b815260040162000dca90620057ff565b601680546001600160a01b0319166001600160a01b0392909216919091179055565b6017546001600160a01b03163314620031075760405162461bcd60e51b815260040162000dca90620057ff565b601480546001600160a01b0319166001600160a01b0392909216919091179055565b6029546001600160a01b03163314620031565760405162461bcd60e51b815260040162000dca90620058e7565b6200316062003879565b62000e19620023bf565b6017546001600160a01b03163314620031975760405162461bcd60e51b815260040162000dca90620057ff565b601180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0383166200321d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840162000dca565b6001600160a01b038216620032805760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840162000dca565b6001600160a01b038381165f8181526025602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0381165f90815260276020526040812090601182015460ff16600281111562003315576200331562005240565b14620033645760405162461bcd60e51b815260206004820152601860248201527f76616c696461746f72206d757374206265206163746976650000000000000000604482015260640162000dca565b60118101805460ff191660011790558054601854601f546001600160a01b038581169316917f75bdcdbe540758778e669d108fbcb7ede734f27f46e4e5525eeb8ecf91849a9c9162002e2d9190620058d1565b6001600160a01b0383165f90815260266020526040902054811115620034195760405162461bcd60e51b8152602060048201526016602482015275616d6f756e7420657863656564732062616c616e636560501b604482015260640162000dca565b6001600160a01b0383165f908152602660205260408120805483929062003442908490620058bb565b90915550506001600160a01b0382165f908152602660205260408120805483929062003470908490620058d1565b92505081905550816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620032d491815260200190565b475f03620034c757565b600d5447905f90670de0b6b3a764000090620034e59084906200597a565b620034f1919062005dc4565b905080156200356c576011546040515f916001600160a01b03169083908381818185875af1925050503d805f811462003546576040519150601f19603f3d011682016040523d82523d5f602084013e6200354b565b606091505b50909150508015156001036200356a57620035678284620058bb565b92505b505b8160225f8282546200357f9190620058d1565b909155505f90505b60215481101562003874575f60275f60218481548110620035ac57620035ac6200594b565b905f5260205f2090600202015f015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f2090505f60205485602185815481106200360957620036096200594b565b905f5260205f209060020201600101546200362591906200597a565b62003631919062005dc4565b905080156200385c576002601183015460ff16600281111562003658576200365862005240565b03620036fb57601254602180546001600160a01b0390921691631de9d9b6918491879081106200368c576200368c6200594b565b5f91825260209091206002909102015460405160e084901b6001600160e01b03191681526001600160a01b0390911660048201526024015f604051808303818588803b158015620036db575f80fd5b505af1158015620036ee573d5f803e3d5ffd5b505050505050506200385f565b5f82600501548284600801546200371391906200597a565b6200371f919062005dc4565b90505f6200372e8284620058bb565b90508115620037905783546040516001600160a01b03909116906108fc9084905f818181858888f193505050503d805f811462003787576040519150601f19603f3d011682016040523d82523d5f602084013e6200378c565b606091505b5050505b8015620038155783600b015f9054906101000a90046001600160a01b03166001600160a01b031663fb489a7b826040518263ffffffff1660e01b815260040160206040518083038185885af1158015620037ec573d5f803e3d5ffd5b50505050506040513d601f19601f8201168201806040525081019062003813919062005dda565b505b60018401546040518481526001600160a01b03909116907fb3b7a071186534c03b40695710096f289fd4ed6c1a374aff0bb648955e4fe5639060200160405180910390a250505b50505b806200386b816200595f565b91505062003587565b505050565b6004545b600554811015620038a657620038a08162003898816200595f565b925062004402565b6200387d565b5060055460045560085460075403620038bb57565b6009545b600854811015620038e857620038e281620038da816200595f565b92506200455c565b620038bf565b50600854600955600754805b60085481101562003961576010545f8281526006602052604090206004015443916200392091620058d1565b1162003946576200393181620047f2565b6200393e600183620058d1565b91506200394c565b62003961565b8062003958816200595f565b915050620038f4565b50600755565b600c54600b54101562003a7c57600b545f908152600a60205260409020601054600182015443916200399991620058d1565b1115620039a35750565b600281015481546001600160a01b039081165f90815260276020526040808220600490810185905585548416835291819020600b015490516319fac8fd60e01b81529216926319fac8fd92620039fd920190815260200190565b5f604051808303815f87803b15801562003a15575f80fd5b505af115801562003a28573d5f803e3d5ffd5b5050600b80545f908152600a6020526040812080546001600160a01b03191681556001808201839055600290910182905582549094509192509062003a6f908490620058d1565b9091555062003967915050565b565b5f821162003adb5760405162461bcd60e51b815260206004820152602360248201527f616d6f756e74206e65656420746f206265207374726963746c7920706f73697460448201526269766560e81b606482015260840162000dca565b6001600160a01b0381165f9081526026602052604090205482111562003b445760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e74204e6577746f6e2062616c616e63650000000000604482015260640162000dca565b6001600160a01b0381165f908152602660205260408120805484929062003b6d908490620058bb565b9091555050604080516080810182526001600160a01b03808416825285811660208084019182528385018781524360608601908152600580545f908152600394859052978820875181549088166001600160a01b031991821617825595516001820180549190981696169590951790955590516002840155519101558054919262003bf8836200595f565b90915550506001600160a01b038481165f81815260276020908152604091829020548251908516948716948514808252918101889052909392917fc46aaee12f38035617ad448c04a7956119f7c7ed395ecc347b898817451ddb8d91015b60405180910390a35050505050565b6001600160a01b038084165f908152602760205260409020805490918381169116148062003dd257600b820154604051631092ab9160e31b81526001600160a01b0385811660048301525f9216906384955c8890602401602060405180830381865afa15801562003cd8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003cfe919062005dda565b90508481101562003d665760405162461bcd60e51b815260206004820152602b60248201527f696e73756666696369656e7420756e6c6f636b6564204c6971756964204e657760448201526a746f6e2062616c616e636560a81b606482015260840162000dca565b600b83015460405163282d3fdf60e01b81526001600160a01b038681166004830152602482018890529091169063282d3fdf906044015f604051808303815f87803b15801562003db4575f80fd5b505af115801562003dc7573d5f803e3d5ffd5b505050505062003e38565b838260080154101562003e385760405162461bcd60e51b815260206004820152602760248201527f696e73756666696369656e742073656c6620626f6e646564206e6577746f6e2060448201526662616c616e636560c81b606482015260840162000dca565b6040805160e0810182526001600160a01b03808616825287811660208084019182528385018981525f60608601818152436080880190815260a088018381528a151560c08a019081526008805486526006909752998420985189549089166001600160a01b0319918216178a55965160018a01805491909916971696909617909655915160028701559051600386015592516004850155905160059093018054945115156101000261ff00199415159490941661ffff19909516949094179290921790925580549162003f0b836200595f565b9190505550826001600160a01b0316856001600160a01b03167f63f8870909f7c59c9c4932bf98dbd491647c8d2e89ca0a032aacdd943a13e2fc838760405162003c569291909115158252602082015260400190565b805160821462003fab5760405162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b604482015260640162000dca565b62003fb68262004925565b604080518082018252601a81527f19457468657265756d205369676e6564204d6573736167653a0a0000000000006020808301919091528451925191925f9262004018920160609190911b6bffffffffffffffffffffffff1916815260140190565b60405160208183030381529060405290505f8262004037835162004a50565b836040516020016200404c9392919062005df2565b60408051601f19818403018152828252805160209182012060028085526060850184529094505f93929091830190803683370190505090505f808060205b88518110156200415357620040a0898262004b6d565b604080515f8152602081018083528b905260ff8316918101919091526060810184905260808101839052929650909450925060019060a0016020604051602081039080840390855afa158015620040f9573d5f803e3d5ffd5b5050604051601f1901519050856200411360418462005dc4565b815181106200412657620041266200594b565b6001600160a01b03909216602092830291909101909101526200414b604182620058d1565b90506200408a565b5088602001516001600160a01b0316845f815181106200417757620041776200594b565b60200260200101516001600160a01b031614620041e95760405162461bcd60e51b815260206004820152602960248201527f496e76616c6964206e6f6465206b6579206f776e6572736869702070726f6f66604482015268081c1c9bdd9a59195960ba1b606482015260840162000dca565b88604001516001600160a01b0316846001815181106200420d576200420d6200594b565b60200260200101516001600160a01b031614620042815760405162461bcd60e51b815260206004820152602b60248201527f496e76616c6964206f7261636c65206b6579206f776e6572736869702070726f60448201526a1bd9881c1c9bdd9a59195960aa1b606482015260840162000dca565b6200428c8962004ba3565b505050505050505050565b62000e19815f60018451620042ad9190620058bb565b62004dab565b8154600260018083161561010002038216048251808201602081106020841001600281146200436057600181146200438657865f526020840460205f2001600160028402018855602085068060200390508088018589016001836101000a0392508282511684540184556001840193506020820191505b808210156200434957815184556001840193506020820191506200432a565b815191036101000a908190040290915550620043f9565b60028302826020036101000a846020036101000a602089015104020185018755620043f9565b865f526020840460205f2001600160028402018855846020038088018589016001836101000a0392508282511660ff198a160184556020820191506001840193505b80821015620043e75781518455600184019350602082019150620043c8565b815191036101000a9081900402909155505b50505050505050565b5f81815260036020908152604080832060018101546001600160a01b039081168552602790935292208054835491929182169116146200451a575f8082600801548360050154620044549190620058bb565b9050805f036200446b578360020154915062004491565b80846002015484600c01546200448291906200597a565b6200448e919062005dc4565b91505b600b83015484546040516340c10f1960e01b81526001600160a01b039182166004820152602481018590529116906340c10f19906044015f604051808303815f87803b158015620044e0575f80fd5b505af1158015620044f3573d5f803e3d5ffd5b505050508183600c015f8282546200450c9190620058d1565b909155506200453992505050565b8160020154816008015f828254620045339190620058d1565b90915550505b8160020154816005015f828254620045529190620058d1565b9091555050505050565b5f81815260066020908152604080832060018101546001600160a01b0316845260279092528220600582015491929091610100900460ff1662004740576002830154600b8301548454604051637eee288d60e01b81526001600160a01b03918216600482015260248101849052911690637eee288d906044015f604051808303815f87803b158015620045ed575f80fd5b505af115801562004600573d5f803e3d5ffd5b50505050600b8301548454604051632770a7eb60e21b81526001600160a01b03918216600482015260248101849052911690639dc29fac906044015f604051808303815f87803b15801562004653575f80fd5b505af115801562004666573d5f803e3d5ffd5b505050505f83600801548460050154620046819190620058bb565b600c8501549091506200469582846200597a565b620046a1919062005dc4565b92508184600c015f828254620046b89190620058bb565b909155505060068401545f03620046d65760038501839055620046fe565b60068401546007850154620046ec90856200597a565b620046f8919062005dc4565b60038601555b82846006015f828254620047139190620058d1565b909155505060038501546007850180545f9062004732908490620058d1565b90915550620047c192505050565b50600282015460098201545f036200475f576003830181905562004787565b6009820154600a8301546200477590836200597a565b62004781919062005dc4565b60038401555b80826009015f8282546200479c9190620058d1565b90915550506003830154600a830180545f90620047bb908490620058d1565b90915550505b6005808401805460ff19166001179055820180548291905f90620047e7908490620058bb565b909155505050505050565b5f81815260066020908152604080832060018101546001600160a01b0316845260279092528220600582015491929091610100900460ff1662004898578160070154826006015484600301546200484a91906200597a565b62004856919062005dc4565b905080826006015f8282546200486d9190620058bb565b909155505060038301546007830180545f906200488c908490620058bb565b90915550620048fb9050565b81600a015482600901548460030154620048b391906200597a565b620048bf919062005dc4565b905080826009015f828254620048d69190620058bb565b90915550506003830154600a830180545f90620048f5908490620058bb565b90915550505b82546001600160a01b03165f9081526026602052604081208054839290620047e7908490620058d1565b5f62004935826060015162004f68565b6001600160a01b03909116602084015290508015620049855760405162461bcd60e51b815260206004820152600b60248201526a32b737b2329032b93937b960a91b604482015260640162000dca565b6020808301516001600160a01b039081165f90815260279092526040909120600101541615620049f85760405162461bcd60e51b815260206004820152601c60248201527f76616c696461746f7220616c7265616479207265676973746572656400000000604482015260640162000dca565b61271082608001511115620021055760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420636f6d6d697373696f6e2072617465000000000000000000604482015260640162000dca565b6060815f0362004a775750506040805180820190915260018152600360fc1b602082015290565b815f5b811562004aa6578062004a8d816200595f565b915062004a9e9050600a8362005dc4565b915062004a7a565b5f8167ffffffffffffffff81111562004ac35762004ac362005539565b6040519080825280601f01601f19166020018201604052801562004aee576020820181803683370190505b5090505b841562004b655762004b06600183620058bb565b915062004b15600a86620059a8565b62004b22906030620058d1565b60f81b81838151811062004b3a5762004b3a6200594b565b60200101906001600160f81b03191690815f1a90535062004b5d600a8662005dc4565b945062004af2565b949350505050565b818101805160208201516040909201519091905f1a601b81101562004b9c5762004b99601b8262005e3a565b90505b9250925092565b6101608101516001600160a01b031662004c2157601c545f9062004bc79062004a50565b90508160200151825f015183608001518360405162004be690620050cf565b62004bf5949392919062005e56565b604051809103905ff08015801562004c0f573d5f803e3d5ffd5b506001600160a01b0316610160830152505b60208181018051601c8054600180820183555f9283527f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a21190910180546001600160a01b03199081166001600160a01b039586161790915584518416835260279095526040918290208651815487169085161781559351908401805486169184169190911790558401516002830180549094169116179091556060820151829190600382019062004cd2908262005cfb565b506080820151600482015560a0820151600582015560c0820151600682015560e0820151600782015561010082015160088201556101208201516009820155610140820151600a820155610160820151600b820180546001600160a01b0319166001600160a01b03909216919091179055610180820151600c8201556101a0820151600d8201556101c0820151600e8201556101e0820151600f820155610200820151601082015561022082015160118201805460ff1916600183600281111562004da15762004da162005240565b0217905550505050565b818180820362004dbc575050505050565b5f85600262004dcc878762005e94565b62004dd8919062005ebd565b62004de4908762005eef565b8151811062004df75762004df76200594b565b602002602001015160a0015190505b81831362004f34575b8086848151811062004e255762004e256200594b565b602002602001015160a00151111562004e4d578262004e448162005f19565b93505062004e0f565b85828151811062004e625762004e626200594b565b602002602001015160a0015181111562004e8b578162004e828162005f33565b92505062004e4d565b81831362004f2e5785828151811062004ea85762004ea86200594b565b602002602001015186848151811062004ec55762004ec56200594b565b602002602001015187858151811062004ee25762004ee26200594b565b6020026020010188858151811062004efe5762004efe6200594b565b602002602001018290528290525050828062004f1a9062005f19565b935050818062004f2a9062005f33565b9250505b62004e06565b8185121562004f4a5762004f4a86868462004dab565b8383121562004f605762004f6086848662004dab565b505050505050565b5f8062004f74620050dd565b5f60408286516020880160ff5afa62004f8b575f80fd5b508051602090910151600160601b90910494909350915050565b6040518061024001604052805f6001600160a01b031681526020015f6001600160a01b031681526020015f6001600160a01b03168152602001606081526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f600281111562005051576200505162005240565b905290565b5080545f8255600202905f5260205f209081019062000e199190620050fb565b5080545f8255905f5260205f209081019062000e19919062005126565b508054620050a1906200586d565b5f825580601f10620050b1575050565b601f0160209004905f5260205f209081019062000e19919062005146565b6114b48062005f5283390190565b60405180604001604052806002906020820280368337509192915050565b5b80821115620051225780546001600160a01b03191681555f6001820155600201620050fc565b5090565b8082111562005122575f6200513c828262005093565b5060010162005126565b5b8082111562005122575f815560010162005147565b5f5b83811015620051785781810151838201526020016200515e565b50505f910152565b5f8151808452620051998160208601602086016200515c565b601f01601f19169290920160200192915050565b602081525f620051c1602083018462005180565b9392505050565b6001600160a01b038116811462000e19575f80fd5b5f8060408385031215620051ef575f80fd5b8235620051fc81620051c8565b946020939093013593505050565b5f602082840312156200521b575f80fd5b8135620051c181620051c8565b5f6020828403121562005239575f80fd5b5035919050565b634e487b7160e01b5f52602160045260245ffd5b600381106200527157634e487b7160e01b5f52602160045260245ffd5b9052565b60208152620052906020820183516001600160a01b03169052565b5f6020830151620052ac60408401826001600160a01b03169052565b5060408301516001600160a01b0381166060840152506060830151610240806080850152620052e061026085018362005180565b9150608085015160a085015260a085015160c085015260c085015160e085015260e085015161010081818701528087015191505061012081818701528087015191505061014081818701528087015191505061016081818701528087015191505061018062005359818701836001600160a01b03169052565b8601516101a0868101919091528601516101c0808701919091528601516101e08087019190915286015161020080870191909152860151610220808701919091528601519050620053ad8286018262005254565b5090949350505050565b5f805f60608486031215620053ca575f80fd5b8335620053d781620051c8565b92506020840135620053e981620051c8565b929592945050506040919091013590565b5f8151808452602080850194508084015f5b838110156200543e57815180516001600160a01b0316885283015183880152604090960195908201906001016200540c565b509495945050505050565b8215158152604060208201525f62004b656040830184620053fa565b5f806040838503121562005477575f80fd5b50508035926020909101359150565b5f6020828403121562005497575f80fd5b813567ffffffffffffffff811115620054ae575f80fd5b82016102408185031215620051c1575f80fd5b5f602080830181845280855180835260408601915060408160051b87010192508387015f5b828110156200551857603f198886030184526200550585835162005180565b94509285019290850190600101620054e6565b5092979650505050505050565b602081525f620051c16020830184620053fa565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126200555d575f80fd5b813567ffffffffffffffff808211156200557b576200557b62005539565b604051601f8301601f19908116603f01168101908282118183101715620055a657620055a662005539565b81604052838152866020858801011115620055bf575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f60608486031215620055f1575f80fd5b833567ffffffffffffffff8082111562005609575f80fd5b62005617878388016200554d565b9450602086013591506200562b82620051c8565b9092506040850135908082111562005641575f80fd5b5062005650868287016200554d565b9150509250925092565b602080825282518282018190525f9190848201906040850190845b818110156200569c5783516001600160a01b03168352928401929184019160010162005675565b50909695505050505050565b5f8060408385031215620056ba575f80fd5b823567ffffffffffffffff80821115620056d2575f80fd5b620056e0868387016200554d565b93506020850135915080821115620056f6575f80fd5b5062005705858286016200554d565b9150509250929050565b604081525f62005723604083018562005180565b828103602084015262005737818562005180565b95945050505050565b5f806040838503121562005752575f80fd5b82356200575f81620051c8565b915060208301356200577181620051c8565b809150509250929050565b6020808252601c908201527f76616c696461746f72206d757374206265207265676973746572656400000000604082015260600190565b6020808252602c908201527f726571756972652063616c6c657220746f2062652076616c696461746f72206160408201526b191b5a5b881858d8dbdd5b9d60a21b606082015260800190565b6020808252601a908201527f63616c6c6572206973206e6f7420746865206f70657261746f72000000000000604082015260600190565b60208082526018908201527f76616c696461746f72206e6f7420726567697374657265640000000000000000604082015260600190565b600181811c908216806200588257607f821691505b602082108103620058a157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111562000d855762000d85620058a7565b8082018082111562000d855762000d85620058a7565b60208082526023908201527f66756e6374696f6e207265737472696374656420746f207468652070726f746f60408201526218dbdb60ea1b606082015260800190565b5f602082840312156200593b575f80fd5b81518015158114620051c1575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f60018201620059735762005973620058a7565b5060010190565b808202811582820484141762000d855762000d85620058a7565b634e487b7160e01b5f52601260045260245ffd5b5f82620059b957620059b962005994565b500690565b5f813562000d8581620051c8565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f808335601e1984360301811262005a02575f80fd5b83018035915067ffffffffffffffff82111562005a1d575f80fd5b60200191503681900382131562005a32575f80fd5b9250929050565b601f82111562003874575f81815260208120601f850160051c8101602086101562005a615750805b601f850160051c820191505b8181101562004f605782815560010162005a6d565b67ffffffffffffffff83111562005a9d5762005a9d62005539565b62005ab58362005aae83546200586d565b8362005a39565b5f601f84116001811462005ae9575f851562005ad15750838201355b5f19600387901b1c1916600186901b178355620012a7565b5f83815260209020601f19861690835b8281101562005b1b578685013582556020948501946001909201910162005af9565b508682101562005b38575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f81356003811062000d85575f80fd5b6003821062005b7757634e487b7160e01b5f52602160045260245ffd5b60ff1981541660ff831681178255505050565b62005ba062005b9983620059be565b82620059cc565b62005bbc62005bb260208401620059be565b60018301620059cc565b62005bd862005bce60408401620059be565b60028301620059cc565b62005be76060830183620059ec565b62005bf781836003860162005a82565b50506080820135600482015560a0820135600582015560c0820135600682015560e0820135600782015561010082013560088201556101208201356009820155610140820135600a82015562005c5f62005c556101608401620059be565b600b8301620059cc565b610180820135600c8201556101a0820135600d8201556101c0820135600e8201556101e0820135600f82015561020082013560108201556200210562005ca9610220840162005b4a565b6011830162005b5a565b5f60018060a01b0380881683528087166020840152808616604084015260a0606084015262005ce660a084018662005180565b91508084166080840152509695505050505050565b815167ffffffffffffffff81111562005d185762005d1862005539565b62005d308162005d2984546200586d565b8462005a39565b602080601f83116001811462005d66575f841562005d4e5750858301515b5f19600386901b1c1916600185901b17855562004f60565b5f85815260208120601f198616915b8281101562005d965788860151825594840194600190910190840162005d75565b508582101562005db457878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8262005dd55762005dd562005994565b500490565b5f6020828403121562005deb575f80fd5b5051919050565b5f845162005e058184602089016200515c565b84519083019062005e1b8183602089016200515c565b845191019062005e308183602088016200515c565b0195945050505050565b60ff818116838216019081111562000d855762000d85620058a7565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f9062005e8a9083018462005180565b9695505050505050565b8181035f83128015838313168383128216171562005eb65762005eb6620058a7565b5092915050565b5f8262005ece5762005ece62005994565b600160ff1b82145f198414161562005eea5762005eea620058a7565b500590565b8082018281125f83128015821682158216171562005f115762005f11620058a7565b505092915050565b5f6001600160ff1b018201620059735762005973620058a7565b5f600160ff1b820162005f4a5762005f4a620058a7565b505f19019056fe608060405234801562000010575f80fd5b50604051620014b4380380620014b4833981016040819052620000339162000149565b61271082111562000042575f80fd5b600a80546001600160a01b038087166001600160a01b031992831617909255600b805492861692909116919091179055600c8290556040516200008a90829060200162000230565b60405160208183030381529060405260089081620000a99190620002ea565b5080604051602001620000bd919062000230565b60405160208183030381529060405260099081620000dc9190620002ea565b50505f80546001600160a01b0319163317905550620003b2915050565b6001600160a01b03811681146200010e575f80fd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200014157818101518382015260200162000127565b50505f910152565b5f805f80608085870312156200015d575f80fd5b84516200016a81620000f9565b60208601519094506200017d81620000f9565b6040860151606087015191945092506001600160401b0380821115620001a1575f80fd5b818701915087601f830112620001b5575f80fd5b815181811115620001ca57620001ca62000111565b604051601f8201601f19908116603f01168101908382118183101715620001f557620001f562000111565b816040528281528a60208487010111156200020e575f80fd5b6200022183602083016020880162000125565b979a9699509497505050505050565b644c4e544e2d60d81b81525f82516200025181600585016020870162000125565b9190910160050192915050565b600181811c908216806200027357607f821691505b6020821081036200029257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002e5575f81815260208120601f850160051c81016020861015620002c05750805b601f850160051c820191505b81811015620002e157828155600101620002cc565b5050505b505050565b81516001600160401b0381111562000306576200030662000111565b6200031e816200031784546200025e565b8462000298565b602080601f83116001811462000354575f84156200033c5750858301515b5f19600386901b1c1916600185901b178555620002e1565b5f85815260208120601f198616915b82811015620003845788860151825594840194600190910190840162000363565b5085821015620003a257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6110f480620003c05f395ff3fe608060405260043610610147575f3560e01c806359355736116100b3578063949813b81161006d578063949813b8146103ac57806395d89b41146103cb5780639dc29fac146103df578063a9059cbb146103fe578063dd62ed3e1461041d578063fb489a7b14610461575f80fd5b806359355736146102d25780635ea1d6f81461030657806361d027b31461031b57806370a082311461033a5780637eee288d1461036e57806384955c881461038d575f80fd5b8063282d3fdf11610104578063282d3fdf146102195780632f2c3f2e14610238578063313ce5671461024d578063372500ab146102685780633a5381b51461027c57806340c10f19146102b3575f80fd5b806306fdde031461014b578063095ea7b31461017557806318160ddd146101a4578063187cf4d7146101c257806319fac8fd146101d957806323b872dd146101fa575b5f80fd5b348015610156575f80fd5b5061015f610469565b60405161016c9190610e7f565b60405180910390f35b348015610180575f80fd5b5061019461018f366004610ee5565b6104f9565b604051901515815260200161016c565b3480156101af575f80fd5b506004545b60405190815260200161016c565b3480156101cd575f80fd5b506101b4633b9aca0081565b3480156101e4575f80fd5b506101f86101f3366004610f0d565b61050f565b005b348015610205575f80fd5b50610194610214366004610f24565b610546565b348015610224575f80fd5b506101f8610233366004610ee5565b610637565b348015610243575f80fd5b506101b461271081565b348015610258575f80fd5b506040516012815260200161016c565b348015610273575f80fd5b506101f8610719565b348015610287575f80fd5b50600a5461029b906001600160a01b031681565b6040516001600160a01b03909116815260200161016c565b3480156102be575f80fd5b506101f86102cd366004610ee5565b6107c3565b3480156102dd575f80fd5b506101b46102ec366004610f5d565b6001600160a01b03165f9081526002602052604090205490565b348015610311575f80fd5b506101b4600c5481565b348015610326575f80fd5b50600b5461029b906001600160a01b031681565b348015610345575f80fd5b506101b4610354366004610f5d565b6001600160a01b03165f9081526001602052604090205490565b348015610379575f80fd5b506101f8610388366004610ee5565b610828565b348015610398575f80fd5b506101b46103a7366004610f5d565b6108eb565b3480156103b7575f80fd5b506101b46103c6366004610f5d565b610918565b3480156103d6575f80fd5b5061015f610944565b3480156103ea575f80fd5b506101f86103f9366004610ee5565b610953565b348015610409575f80fd5b50610194610418366004610ee5565b6109b0565b348015610428575f80fd5b506101b4610437366004610f7d565b6001600160a01b039182165f90815260036020908152604080832093909416825291909152205490565b6101b46109fb565b60606008805461047890610fae565b80601f01602080910402602001604051908101604052809291908181526020018280546104a490610fae565b80156104ef5780601f106104c6576101008083540402835291602001916104ef565b820191905f5260205f20905b8154815290600101906020018083116104d257829003601f168201915b5050505050905090565b5f610505338484610b5c565b5060015b92915050565b5f546001600160a01b031633146105415760405162461bcd60e51b815260040161053890610fe6565b60405180910390fd5b600c55565b6001600160a01b0383165f908152600360209081526040808320338452909152812054828110156105ca5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152608401610538565b6105de85336105d98685611042565b610b5c565b6105e88584610c7f565b6105f28484610d63565b836001600160a01b0316856001600160a01b03165f8051602061109f8339815191528560405161062491815260200190565b60405180910390a3506001949350505050565b5f546001600160a01b031633146106605760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f90815260026020908152604080832054600190925290912054829161068f91611042565b10156106e95760405162461bcd60e51b8152602060048201526024808201527f63616e2774206c6f636b206d6f72652066756e6473207468616e20617661696c60448201526361626c6560e01b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611055565b90915550505050565b5f61072333610dac565b335f81815260056020526040808220829055519293509183908381818185875af1925050503d805f8114610772576040519150601f19603f3d011682016040523d82523d5f602084013e610777565b606091505b50509050806107bf5760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b6044820152606401610538565b5050565b5f546001600160a01b031633146107ec5760405162461bcd60e51b815260040161053890610fe6565b6107f68282610d63565b6040518181526001600160a01b038316905f905f8051602061109f833981519152906020015b60405180910390a35050565b5f546001600160a01b031633146108515760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f908152600260205260409020548111156108c45760405162461bcd60e51b815260206004820152602360248201527f63616e277420756e6c6f636b206d6f72652066756e6473207468616e206c6f636044820152621ad95960ea1b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611042565b6001600160a01b0381165f9081526002602090815260408083205460019092528220546105099190611042565b5f61092282610e0e565b6001600160a01b0383165f908152600560205260409020546105099190611055565b60606009805461047890610fae565b5f546001600160a01b0316331461097c5760405162461bcd60e51b815260040161053890610fe6565b6109868282610c7f565b6040518181525f906001600160a01b038416905f8051602061109f8339815191529060200161081c565b5f6109bb3383610c7f565b6109c58383610d63565b6040518281526001600160a01b0384169033905f8051602061109f8339815191529060200160405180910390a350600192915050565b5f80546001600160a01b03163314610a255760405162461bcd60e51b815260040161053890610fe6565b600c5434905f9061271090610a3a9084611068565b610a44919061107f565b905081811115610a965760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642076616c696461746f722072657761726400000000000000006044820152606401610538565b610aa08183611042565b600b546040519193506001600160a01b0316906108fc9083905f818181858888f193505050503d805f8114610af0576040519150601f19603f3d011682016040523d82523d5f602084013e610af5565b606091505b50506004545f9150610b0b633b9aca0085611068565b610b15919061107f565b905080600754610b259190611055565b6007556004545f90633b9aca0090610b3d9084611068565b610b47919061107f565b9050610b538184611055565b94505050505090565b6001600160a01b038316610bbe5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610538565b6001600160a01b038216610c1f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610538565b6001600160a01b038381165f8181526003602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b610c8882610dac565b506001600160a01b0382165f90815260016020908152604080832054600290925290912054610cb79082611042565b821115610d065760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e7420756e6c6f636b65642066756e647300000000006044820152606401610538565b610d108282611042565b6001600160a01b0384165f90815260016020526040902055808203610d48576001600160a01b0383165f908152600660205260408120555b8160045f828254610d599190611042565b9091555050505050565b610d6c82610dac565b506001600160a01b0382165f9081526001602052604081208054839290610d94908490611055565b925050819055508060045f8282546107109190611055565b5f80610db783610e0e565b6001600160a01b0384165f90815260056020526040902054909150610ddd908290611055565b6001600160a01b039093165f9081526005602090815260408083208690556007546006909252909120555090919050565b6001600160a01b0381165f90815260016020526040812054808203610e3557505f92915050565b6001600160a01b0383165f90815260066020526040812054600754610e5a9190611042565b90505f633b9aca00610e6c8484611068565b610e76919061107f565b95945050505050565b5f6020808352835180828501525f5b81811015610eaa57858101830151858201604001528201610e8e565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610ee0575f80fd5b919050565b5f8060408385031215610ef6575f80fd5b610eff83610eca565b946020939093013593505050565b5f60208284031215610f1d575f80fd5b5035919050565b5f805f60608486031215610f36575f80fd5b610f3f84610eca565b9250610f4d60208501610eca565b9150604084013590509250925092565b5f60208284031215610f6d575f80fd5b610f7682610eca565b9392505050565b5f8060408385031215610f8e575f80fd5b610f9783610eca565b9150610fa560208401610eca565b90509250929050565b600181811c90821680610fc257607f821691505b602082108103610fe057634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526028908201527f43616c6c207265737472696374656420746f20746865204175746f6e6974792060408201526710dbdb9d1c9858dd60c21b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156105095761050961102e565b808201808211156105095761050961102e565b80820281158282048414176105095761050961102e565b5f8261109957634e487b7160e01b5f52601260045260245ffd5b50049056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212208b971a8921a90e53cf2766d536cfa590f79864d3328a6958e1150f86b9db10c264736f6c63430008150033a26469706673582212206bd9a25d9274e60f60615496173a1e9c49109341e58f40f1e4af1249c92227d864736f6c63430008150033608060405234801562000010575f80fd5b50604051620014b4380380620014b4833981016040819052620000339162000149565b61271082111562000042575f80fd5b600a80546001600160a01b038087166001600160a01b031992831617909255600b805492861692909116919091179055600c8290556040516200008a90829060200162000230565b60405160208183030381529060405260089081620000a99190620002ea565b5080604051602001620000bd919062000230565b60405160208183030381529060405260099081620000dc9190620002ea565b50505f80546001600160a01b0319163317905550620003b2915050565b6001600160a01b03811681146200010e575f80fd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200014157818101518382015260200162000127565b50505f910152565b5f805f80608085870312156200015d575f80fd5b84516200016a81620000f9565b60208601519094506200017d81620000f9565b6040860151606087015191945092506001600160401b0380821115620001a1575f80fd5b818701915087601f830112620001b5575f80fd5b815181811115620001ca57620001ca62000111565b604051601f8201601f19908116603f01168101908382118183101715620001f557620001f562000111565b816040528281528a60208487010111156200020e575f80fd5b6200022183602083016020880162000125565b979a9699509497505050505050565b644c4e544e2d60d81b81525f82516200025181600585016020870162000125565b9190910160050192915050565b600181811c908216806200027357607f821691505b6020821081036200029257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002e5575f81815260208120601f850160051c81016020861015620002c05750805b601f850160051c820191505b81811015620002e157828155600101620002cc565b5050505b505050565b81516001600160401b0381111562000306576200030662000111565b6200031e816200031784546200025e565b8462000298565b602080601f83116001811462000354575f84156200033c5750858301515b5f19600386901b1c1916600185901b178555620002e1565b5f85815260208120601f198616915b82811015620003845788860151825594840194600190910190840162000363565b5085821015620003a257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6110f480620003c05f395ff3fe608060405260043610610147575f3560e01c806359355736116100b3578063949813b81161006d578063949813b8146103ac57806395d89b41146103cb5780639dc29fac146103df578063a9059cbb146103fe578063dd62ed3e1461041d578063fb489a7b14610461575f80fd5b806359355736146102d25780635ea1d6f81461030657806361d027b31461031b57806370a082311461033a5780637eee288d1461036e57806384955c881461038d575f80fd5b8063282d3fdf11610104578063282d3fdf146102195780632f2c3f2e14610238578063313ce5671461024d578063372500ab146102685780633a5381b51461027c57806340c10f19146102b3575f80fd5b806306fdde031461014b578063095ea7b31461017557806318160ddd146101a4578063187cf4d7146101c257806319fac8fd146101d957806323b872dd146101fa575b5f80fd5b348015610156575f80fd5b5061015f610469565b60405161016c9190610e7f565b60405180910390f35b348015610180575f80fd5b5061019461018f366004610ee5565b6104f9565b604051901515815260200161016c565b3480156101af575f80fd5b506004545b60405190815260200161016c565b3480156101cd575f80fd5b506101b4633b9aca0081565b3480156101e4575f80fd5b506101f86101f3366004610f0d565b61050f565b005b348015610205575f80fd5b50610194610214366004610f24565b610546565b348015610224575f80fd5b506101f8610233366004610ee5565b610637565b348015610243575f80fd5b506101b461271081565b348015610258575f80fd5b506040516012815260200161016c565b348015610273575f80fd5b506101f8610719565b348015610287575f80fd5b50600a5461029b906001600160a01b031681565b6040516001600160a01b03909116815260200161016c565b3480156102be575f80fd5b506101f86102cd366004610ee5565b6107c3565b3480156102dd575f80fd5b506101b46102ec366004610f5d565b6001600160a01b03165f9081526002602052604090205490565b348015610311575f80fd5b506101b4600c5481565b348015610326575f80fd5b50600b5461029b906001600160a01b031681565b348015610345575f80fd5b506101b4610354366004610f5d565b6001600160a01b03165f9081526001602052604090205490565b348015610379575f80fd5b506101f8610388366004610ee5565b610828565b348015610398575f80fd5b506101b46103a7366004610f5d565b6108eb565b3480156103b7575f80fd5b506101b46103c6366004610f5d565b610918565b3480156103d6575f80fd5b5061015f610944565b3480156103ea575f80fd5b506101f86103f9366004610ee5565b610953565b348015610409575f80fd5b50610194610418366004610ee5565b6109b0565b348015610428575f80fd5b506101b4610437366004610f7d565b6001600160a01b039182165f90815260036020908152604080832093909416825291909152205490565b6101b46109fb565b60606008805461047890610fae565b80601f01602080910402602001604051908101604052809291908181526020018280546104a490610fae565b80156104ef5780601f106104c6576101008083540402835291602001916104ef565b820191905f5260205f20905b8154815290600101906020018083116104d257829003601f168201915b5050505050905090565b5f610505338484610b5c565b5060015b92915050565b5f546001600160a01b031633146105415760405162461bcd60e51b815260040161053890610fe6565b60405180910390fd5b600c55565b6001600160a01b0383165f908152600360209081526040808320338452909152812054828110156105ca5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152608401610538565b6105de85336105d98685611042565b610b5c565b6105e88584610c7f565b6105f28484610d63565b836001600160a01b0316856001600160a01b03165f8051602061109f8339815191528560405161062491815260200190565b60405180910390a3506001949350505050565b5f546001600160a01b031633146106605760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f90815260026020908152604080832054600190925290912054829161068f91611042565b10156106e95760405162461bcd60e51b8152602060048201526024808201527f63616e2774206c6f636b206d6f72652066756e6473207468616e20617661696c60448201526361626c6560e01b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611055565b90915550505050565b5f61072333610dac565b335f81815260056020526040808220829055519293509183908381818185875af1925050503d805f8114610772576040519150601f19603f3d011682016040523d82523d5f602084013e610777565b606091505b50509050806107bf5760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b6044820152606401610538565b5050565b5f546001600160a01b031633146107ec5760405162461bcd60e51b815260040161053890610fe6565b6107f68282610d63565b6040518181526001600160a01b038316905f905f8051602061109f833981519152906020015b60405180910390a35050565b5f546001600160a01b031633146108515760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f908152600260205260409020548111156108c45760405162461bcd60e51b815260206004820152602360248201527f63616e277420756e6c6f636b206d6f72652066756e6473207468616e206c6f636044820152621ad95960ea1b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611042565b6001600160a01b0381165f9081526002602090815260408083205460019092528220546105099190611042565b5f61092282610e0e565b6001600160a01b0383165f908152600560205260409020546105099190611055565b60606009805461047890610fae565b5f546001600160a01b0316331461097c5760405162461bcd60e51b815260040161053890610fe6565b6109868282610c7f565b6040518181525f906001600160a01b038416905f8051602061109f8339815191529060200161081c565b5f6109bb3383610c7f565b6109c58383610d63565b6040518281526001600160a01b0384169033905f8051602061109f8339815191529060200160405180910390a350600192915050565b5f80546001600160a01b03163314610a255760405162461bcd60e51b815260040161053890610fe6565b600c5434905f9061271090610a3a9084611068565b610a44919061107f565b905081811115610a965760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642076616c696461746f722072657761726400000000000000006044820152606401610538565b610aa08183611042565b600b546040519193506001600160a01b0316906108fc9083905f818181858888f193505050503d805f8114610af0576040519150601f19603f3d011682016040523d82523d5f602084013e610af5565b606091505b50506004545f9150610b0b633b9aca0085611068565b610b15919061107f565b905080600754610b259190611055565b6007556004545f90633b9aca0090610b3d9084611068565b610b47919061107f565b9050610b538184611055565b94505050505090565b6001600160a01b038316610bbe5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610538565b6001600160a01b038216610c1f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610538565b6001600160a01b038381165f8181526003602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b610c8882610dac565b506001600160a01b0382165f90815260016020908152604080832054600290925290912054610cb79082611042565b821115610d065760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e7420756e6c6f636b65642066756e647300000000006044820152606401610538565b610d108282611042565b6001600160a01b0384165f90815260016020526040902055808203610d48576001600160a01b0383165f908152600660205260408120555b8160045f828254610d599190611042565b9091555050505050565b610d6c82610dac565b506001600160a01b0382165f9081526001602052604081208054839290610d94908490611055565b925050819055508060045f8282546107109190611055565b5f80610db783610e0e565b6001600160a01b0384165f90815260056020526040902054909150610ddd908290611055565b6001600160a01b039093165f9081526005602090815260408083208690556007546006909252909120555090919050565b6001600160a01b0381165f90815260016020526040812054808203610e3557505f92915050565b6001600160a01b0383165f90815260066020526040812054600754610e5a9190611042565b90505f633b9aca00610e6c8484611068565b610e76919061107f565b95945050505050565b5f6020808352835180828501525f5b81811015610eaa57858101830151858201604001528201610e8e565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610ee0575f80fd5b919050565b5f8060408385031215610ef6575f80fd5b610eff83610eca565b946020939093013593505050565b5f60208284031215610f1d575f80fd5b5035919050565b5f805f60608486031215610f36575f80fd5b610f3f84610eca565b9250610f4d60208501610eca565b9150604084013590509250925092565b5f60208284031215610f6d575f80fd5b610f7682610eca565b9392505050565b5f8060408385031215610f8e575f80fd5b610f9783610eca565b9150610fa560208401610eca565b90509250929050565b600181811c90821680610fc257607f821691505b602082108103610fe057634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526028908201527f43616c6c207265737472696374656420746f20746865204175746f6e6974792060408201526710dbdb9d1c9858dd60c21b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156105095761050961102e565b808201808211156105095761050961102e565b80820281158282048414176105095761050961102e565b5f8261109957634e487b7160e01b5f52601260045260245ffd5b50049056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212208b971a8921a90e53cf2766d536cfa590f79864d3328a6958e1150f86b9db10c264736f6c63430008150033",
}

// AutonityABI is the input ABI used to generate the binding from.
// Deprecated: Use AutonityMetaData.ABI instead.
var AutonityABI = AutonityMetaData.ABI

// Deprecated: Use AutonityMetaData.Sigs instead.
// AutonityFuncSigs maps the 4-byte function signature to its string representation.
var AutonityFuncSigs = AutonityMetaData.Sigs

// AutonityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AutonityMetaData.Bin instead.
var AutonityBin = AutonityMetaData.Bin

// DeployAutonity deploys a new Ethereum contract, binding an instance of Autonity to it.
func DeployAutonity(auth *bind.TransactOpts, backend bind.ContractBackend, _validators []AutonityValidator, _config AutonityConfig) (common.Address, *types.Transaction, *Autonity, error) {
	parsed, err := AutonityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AutonityBin), backend, _validators, _config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Autonity{AutonityCaller: AutonityCaller{contract: contract}, AutonityTransactor: AutonityTransactor{contract: contract}, AutonityFilterer: AutonityFilterer{contract: contract}}, nil
}

// Autonity is an auto generated Go binding around an Ethereum contract.
type Autonity struct {
	AutonityCaller     // Read-only binding to the contract
	AutonityTransactor // Write-only binding to the contract
	AutonityFilterer   // Log filterer for contract events
}

// AutonityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AutonityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutonityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AutonityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutonityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutonityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutonitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutonitySession struct {
	Contract     *Autonity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AutonityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutonityCallerSession struct {
	Contract *AutonityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AutonityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutonityTransactorSession struct {
	Contract     *AutonityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AutonityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AutonityRaw struct {
	Contract *Autonity // Generic contract binding to access the raw methods on
}

// AutonityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutonityCallerRaw struct {
	Contract *AutonityCaller // Generic read-only contract binding to access the raw methods on
}

// AutonityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutonityTransactorRaw struct {
	Contract *AutonityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAutonity creates a new instance of Autonity, bound to a specific deployed contract.
func NewAutonity(address common.Address, backend bind.ContractBackend) (*Autonity, error) {
	contract, err := bindAutonity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Autonity{AutonityCaller: AutonityCaller{contract: contract}, AutonityTransactor: AutonityTransactor{contract: contract}, AutonityFilterer: AutonityFilterer{contract: contract}}, nil
}

// NewAutonityCaller creates a new read-only instance of Autonity, bound to a specific deployed contract.
func NewAutonityCaller(address common.Address, caller bind.ContractCaller) (*AutonityCaller, error) {
	contract, err := bindAutonity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutonityCaller{contract: contract}, nil
}

// NewAutonityTransactor creates a new write-only instance of Autonity, bound to a specific deployed contract.
func NewAutonityTransactor(address common.Address, transactor bind.ContractTransactor) (*AutonityTransactor, error) {
	contract, err := bindAutonity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutonityTransactor{contract: contract}, nil
}

// NewAutonityFilterer creates a new log filterer instance of Autonity, bound to a specific deployed contract.
func NewAutonityFilterer(address common.Address, filterer bind.ContractFilterer) (*AutonityFilterer, error) {
	contract, err := bindAutonity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutonityFilterer{contract: contract}, nil
}

// bindAutonity binds a generic wrapper to an already deployed contract.
func bindAutonity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AutonityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Autonity *AutonityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Autonity.Contract.AutonityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Autonity *AutonityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.Contract.AutonityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Autonity *AutonityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Autonity.Contract.AutonityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Autonity *AutonityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Autonity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Autonity *AutonityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Autonity *AutonityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Autonity.Contract.contract.Transact(opts, method, params...)
}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Autonity *AutonityCaller) COMMISSIONRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "COMMISSION_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Autonity *AutonitySession) COMMISSIONRATEPRECISION() (*big.Int, error) {
	return _Autonity.Contract.COMMISSIONRATEPRECISION(&_Autonity.CallOpts)
}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Autonity *AutonityCallerSession) COMMISSIONRATEPRECISION() (*big.Int, error) {
	return _Autonity.Contract.COMMISSIONRATEPRECISION(&_Autonity.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Autonity *AutonityCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Autonity *AutonitySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Autonity.Contract.Allowance(&_Autonity.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Autonity *AutonityCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Autonity.Contract.Allowance(&_Autonity.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_Autonity *AutonityCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_Autonity *AutonitySession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _Autonity.Contract.BalanceOf(&_Autonity.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_Autonity *AutonityCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _Autonity.Contract.BalanceOf(&_Autonity.CallOpts, _addr)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns((uint256,uint256,uint256,uint256,address) policy, (address,address,address,address,address) contracts, (address,uint256,uint256,uint256) protocol, uint256 contractVersion)
func (_Autonity *AutonityCaller) Config(opts *bind.CallOpts) (struct {
	Policy          AutonityPolicy
	Contracts       AutonityContracts
	Protocol        AutonityProtocol
	ContractVersion *big.Int
}, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Policy          AutonityPolicy
		Contracts       AutonityContracts
		Protocol        AutonityProtocol
		ContractVersion *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Policy = *abi.ConvertType(out[0], new(AutonityPolicy)).(*AutonityPolicy)
	outstruct.Contracts = *abi.ConvertType(out[1], new(AutonityContracts)).(*AutonityContracts)
	outstruct.Protocol = *abi.ConvertType(out[2], new(AutonityProtocol)).(*AutonityProtocol)
	outstruct.ContractVersion = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns((uint256,uint256,uint256,uint256,address) policy, (address,address,address,address,address) contracts, (address,uint256,uint256,uint256) protocol, uint256 contractVersion)
func (_Autonity *AutonitySession) Config() (struct {
	Policy          AutonityPolicy
	Contracts       AutonityContracts
	Protocol        AutonityProtocol
	ContractVersion *big.Int
}, error) {
	return _Autonity.Contract.Config(&_Autonity.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns((uint256,uint256,uint256,uint256,address) policy, (address,address,address,address,address) contracts, (address,uint256,uint256,uint256) protocol, uint256 contractVersion)
func (_Autonity *AutonityCallerSession) Config() (struct {
	Policy          AutonityPolicy
	Contracts       AutonityContracts
	Protocol        AutonityProtocol
	ContractVersion *big.Int
}, error) {
	return _Autonity.Contract.Config(&_Autonity.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Autonity *AutonityCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Autonity *AutonitySession) Decimals() (uint8, error) {
	return _Autonity.Contract.Decimals(&_Autonity.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Autonity *AutonityCallerSession) Decimals() (uint8, error) {
	return _Autonity.Contract.Decimals(&_Autonity.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Autonity *AutonityCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Autonity *AutonitySession) Deployer() (common.Address, error) {
	return _Autonity.Contract.Deployer(&_Autonity.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Autonity *AutonityCallerSession) Deployer() (common.Address, error) {
	return _Autonity.Contract.Deployer(&_Autonity.CallOpts)
}

// EpochID is a free data retrieval call binding the contract method 0xc9d97af4.
//
// Solidity: function epochID() view returns(uint256)
func (_Autonity *AutonityCaller) EpochID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "epochID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochID is a free data retrieval call binding the contract method 0xc9d97af4.
//
// Solidity: function epochID() view returns(uint256)
func (_Autonity *AutonitySession) EpochID() (*big.Int, error) {
	return _Autonity.Contract.EpochID(&_Autonity.CallOpts)
}

// EpochID is a free data retrieval call binding the contract method 0xc9d97af4.
//
// Solidity: function epochID() view returns(uint256)
func (_Autonity *AutonityCallerSession) EpochID() (*big.Int, error) {
	return _Autonity.Contract.EpochID(&_Autonity.CallOpts)
}

// EpochReward is a free data retrieval call binding the contract method 0x1604e416.
//
// Solidity: function epochReward() view returns(uint256)
func (_Autonity *AutonityCaller) EpochReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "epochReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochReward is a free data retrieval call binding the contract method 0x1604e416.
//
// Solidity: function epochReward() view returns(uint256)
func (_Autonity *AutonitySession) EpochReward() (*big.Int, error) {
	return _Autonity.Contract.EpochReward(&_Autonity.CallOpts)
}

// EpochReward is a free data retrieval call binding the contract method 0x1604e416.
//
// Solidity: function epochReward() view returns(uint256)
func (_Autonity *AutonityCallerSession) EpochReward() (*big.Int, error) {
	return _Autonity.Contract.EpochReward(&_Autonity.CallOpts)
}

// EpochTotalBondedStake is a free data retrieval call binding the contract method 0x9c98e471.
//
// Solidity: function epochTotalBondedStake() view returns(uint256)
func (_Autonity *AutonityCaller) EpochTotalBondedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "epochTotalBondedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTotalBondedStake is a free data retrieval call binding the contract method 0x9c98e471.
//
// Solidity: function epochTotalBondedStake() view returns(uint256)
func (_Autonity *AutonitySession) EpochTotalBondedStake() (*big.Int, error) {
	return _Autonity.Contract.EpochTotalBondedStake(&_Autonity.CallOpts)
}

// EpochTotalBondedStake is a free data retrieval call binding the contract method 0x9c98e471.
//
// Solidity: function epochTotalBondedStake() view returns(uint256)
func (_Autonity *AutonityCallerSession) EpochTotalBondedStake() (*big.Int, error) {
	return _Autonity.Contract.EpochTotalBondedStake(&_Autonity.CallOpts)
}

// GetBlockPeriod is a free data retrieval call binding the contract method 0x43645969.
//
// Solidity: function getBlockPeriod() view returns(uint256)
func (_Autonity *AutonityCaller) GetBlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getBlockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockPeriod is a free data retrieval call binding the contract method 0x43645969.
//
// Solidity: function getBlockPeriod() view returns(uint256)
func (_Autonity *AutonitySession) GetBlockPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetBlockPeriod(&_Autonity.CallOpts)
}

// GetBlockPeriod is a free data retrieval call binding the contract method 0x43645969.
//
// Solidity: function getBlockPeriod() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetBlockPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetBlockPeriod(&_Autonity.CallOpts)
}

// GetCommittee is a free data retrieval call binding the contract method 0xab8f6ffe.
//
// Solidity: function getCommittee() view returns((address,uint256)[])
func (_Autonity *AutonityCaller) GetCommittee(opts *bind.CallOpts) ([]AutonityCommitteeMember, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getCommittee")

	if err != nil {
		return *new([]AutonityCommitteeMember), err
	}

	out0 := *abi.ConvertType(out[0], new([]AutonityCommitteeMember)).(*[]AutonityCommitteeMember)

	return out0, err

}

// GetCommittee is a free data retrieval call binding the contract method 0xab8f6ffe.
//
// Solidity: function getCommittee() view returns((address,uint256)[])
func (_Autonity *AutonitySession) GetCommittee() ([]AutonityCommitteeMember, error) {
	return _Autonity.Contract.GetCommittee(&_Autonity.CallOpts)
}

// GetCommittee is a free data retrieval call binding the contract method 0xab8f6ffe.
//
// Solidity: function getCommittee() view returns((address,uint256)[])
func (_Autonity *AutonityCallerSession) GetCommittee() ([]AutonityCommitteeMember, error) {
	return _Autonity.Contract.GetCommittee(&_Autonity.CallOpts)
}

// GetCommitteeEnodes is a free data retrieval call binding the contract method 0xa8b2216e.
//
// Solidity: function getCommitteeEnodes() view returns(string[])
func (_Autonity *AutonityCaller) GetCommitteeEnodes(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getCommitteeEnodes")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCommitteeEnodes is a free data retrieval call binding the contract method 0xa8b2216e.
//
// Solidity: function getCommitteeEnodes() view returns(string[])
func (_Autonity *AutonitySession) GetCommitteeEnodes() ([]string, error) {
	return _Autonity.Contract.GetCommitteeEnodes(&_Autonity.CallOpts)
}

// GetCommitteeEnodes is a free data retrieval call binding the contract method 0xa8b2216e.
//
// Solidity: function getCommitteeEnodes() view returns(string[])
func (_Autonity *AutonityCallerSession) GetCommitteeEnodes() ([]string, error) {
	return _Autonity.Contract.GetCommitteeEnodes(&_Autonity.CallOpts)
}

// GetEpochFromBlock is a free data retrieval call binding the contract method 0x96b477cb.
//
// Solidity: function getEpochFromBlock(uint256 _block) view returns(uint256)
func (_Autonity *AutonityCaller) GetEpochFromBlock(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getEpochFromBlock", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochFromBlock is a free data retrieval call binding the contract method 0x96b477cb.
//
// Solidity: function getEpochFromBlock(uint256 _block) view returns(uint256)
func (_Autonity *AutonitySession) GetEpochFromBlock(_block *big.Int) (*big.Int, error) {
	return _Autonity.Contract.GetEpochFromBlock(&_Autonity.CallOpts, _block)
}

// GetEpochFromBlock is a free data retrieval call binding the contract method 0x96b477cb.
//
// Solidity: function getEpochFromBlock(uint256 _block) view returns(uint256)
func (_Autonity *AutonityCallerSession) GetEpochFromBlock(_block *big.Int) (*big.Int, error) {
	return _Autonity.Contract.GetEpochFromBlock(&_Autonity.CallOpts, _block)
}

// GetEpochPeriod is a free data retrieval call binding the contract method 0xdfb1a4d2.
//
// Solidity: function getEpochPeriod() view returns(uint256)
func (_Autonity *AutonityCaller) GetEpochPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getEpochPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochPeriod is a free data retrieval call binding the contract method 0xdfb1a4d2.
//
// Solidity: function getEpochPeriod() view returns(uint256)
func (_Autonity *AutonitySession) GetEpochPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetEpochPeriod(&_Autonity.CallOpts)
}

// GetEpochPeriod is a free data retrieval call binding the contract method 0xdfb1a4d2.
//
// Solidity: function getEpochPeriod() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetEpochPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetEpochPeriod(&_Autonity.CallOpts)
}

// GetLastEpochBlock is a free data retrieval call binding the contract method 0x731b3a03.
//
// Solidity: function getLastEpochBlock() view returns(uint256)
func (_Autonity *AutonityCaller) GetLastEpochBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getLastEpochBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEpochBlock is a free data retrieval call binding the contract method 0x731b3a03.
//
// Solidity: function getLastEpochBlock() view returns(uint256)
func (_Autonity *AutonitySession) GetLastEpochBlock() (*big.Int, error) {
	return _Autonity.Contract.GetLastEpochBlock(&_Autonity.CallOpts)
}

// GetLastEpochBlock is a free data retrieval call binding the contract method 0x731b3a03.
//
// Solidity: function getLastEpochBlock() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetLastEpochBlock() (*big.Int, error) {
	return _Autonity.Contract.GetLastEpochBlock(&_Autonity.CallOpts)
}

// GetMaxCommitteeSize is a free data retrieval call binding the contract method 0x819b6463.
//
// Solidity: function getMaxCommitteeSize() view returns(uint256)
func (_Autonity *AutonityCaller) GetMaxCommitteeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getMaxCommitteeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxCommitteeSize is a free data retrieval call binding the contract method 0x819b6463.
//
// Solidity: function getMaxCommitteeSize() view returns(uint256)
func (_Autonity *AutonitySession) GetMaxCommitteeSize() (*big.Int, error) {
	return _Autonity.Contract.GetMaxCommitteeSize(&_Autonity.CallOpts)
}

// GetMaxCommitteeSize is a free data retrieval call binding the contract method 0x819b6463.
//
// Solidity: function getMaxCommitteeSize() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetMaxCommitteeSize() (*big.Int, error) {
	return _Autonity.Contract.GetMaxCommitteeSize(&_Autonity.CallOpts)
}

// GetMinimumBaseFee is a free data retrieval call binding the contract method 0x11220633.
//
// Solidity: function getMinimumBaseFee() view returns(uint256)
func (_Autonity *AutonityCaller) GetMinimumBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getMinimumBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBaseFee is a free data retrieval call binding the contract method 0x11220633.
//
// Solidity: function getMinimumBaseFee() view returns(uint256)
func (_Autonity *AutonitySession) GetMinimumBaseFee() (*big.Int, error) {
	return _Autonity.Contract.GetMinimumBaseFee(&_Autonity.CallOpts)
}

// GetMinimumBaseFee is a free data retrieval call binding the contract method 0x11220633.
//
// Solidity: function getMinimumBaseFee() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetMinimumBaseFee() (*big.Int, error) {
	return _Autonity.Contract.GetMinimumBaseFee(&_Autonity.CallOpts)
}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Autonity *AutonityCaller) GetNewContract(opts *bind.CallOpts) ([]byte, string, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getNewContract")

	if err != nil {
		return *new([]byte), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Autonity *AutonitySession) GetNewContract() ([]byte, string, error) {
	return _Autonity.Contract.GetNewContract(&_Autonity.CallOpts)
}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Autonity *AutonityCallerSession) GetNewContract() ([]byte, string, error) {
	return _Autonity.Contract.GetNewContract(&_Autonity.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Autonity *AutonityCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Autonity *AutonitySession) GetOperator() (common.Address, error) {
	return _Autonity.Contract.GetOperator(&_Autonity.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Autonity *AutonityCallerSession) GetOperator() (common.Address, error) {
	return _Autonity.Contract.GetOperator(&_Autonity.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Autonity *AutonityCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Autonity *AutonitySession) GetOracle() (common.Address, error) {
	return _Autonity.Contract.GetOracle(&_Autonity.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Autonity *AutonityCallerSession) GetOracle() (common.Address, error) {
	return _Autonity.Contract.GetOracle(&_Autonity.CallOpts)
}

// GetProposer is a free data retrieval call binding the contract method 0x5f7d3949.
//
// Solidity: function getProposer(uint256 height, uint256 round) view returns(address)
func (_Autonity *AutonityCaller) GetProposer(opts *bind.CallOpts, height *big.Int, round *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getProposer", height, round)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProposer is a free data retrieval call binding the contract method 0x5f7d3949.
//
// Solidity: function getProposer(uint256 height, uint256 round) view returns(address)
func (_Autonity *AutonitySession) GetProposer(height *big.Int, round *big.Int) (common.Address, error) {
	return _Autonity.Contract.GetProposer(&_Autonity.CallOpts, height, round)
}

// GetProposer is a free data retrieval call binding the contract method 0x5f7d3949.
//
// Solidity: function getProposer(uint256 height, uint256 round) view returns(address)
func (_Autonity *AutonityCallerSession) GetProposer(height *big.Int, round *big.Int) (common.Address, error) {
	return _Autonity.Contract.GetProposer(&_Autonity.CallOpts, height, round)
}

// GetTreasuryAccount is a free data retrieval call binding the contract method 0xf7866ee3.
//
// Solidity: function getTreasuryAccount() view returns(address)
func (_Autonity *AutonityCaller) GetTreasuryAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getTreasuryAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasuryAccount is a free data retrieval call binding the contract method 0xf7866ee3.
//
// Solidity: function getTreasuryAccount() view returns(address)
func (_Autonity *AutonitySession) GetTreasuryAccount() (common.Address, error) {
	return _Autonity.Contract.GetTreasuryAccount(&_Autonity.CallOpts)
}

// GetTreasuryAccount is a free data retrieval call binding the contract method 0xf7866ee3.
//
// Solidity: function getTreasuryAccount() view returns(address)
func (_Autonity *AutonityCallerSession) GetTreasuryAccount() (common.Address, error) {
	return _Autonity.Contract.GetTreasuryAccount(&_Autonity.CallOpts)
}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint256)
func (_Autonity *AutonityCaller) GetTreasuryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getTreasuryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint256)
func (_Autonity *AutonitySession) GetTreasuryFee() (*big.Int, error) {
	return _Autonity.Contract.GetTreasuryFee(&_Autonity.CallOpts)
}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetTreasuryFee() (*big.Int, error) {
	return _Autonity.Contract.GetTreasuryFee(&_Autonity.CallOpts)
}

// GetUnbondingPeriod is a free data retrieval call binding the contract method 0x6fd2c80b.
//
// Solidity: function getUnbondingPeriod() view returns(uint256)
func (_Autonity *AutonityCaller) GetUnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getUnbondingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondingPeriod is a free data retrieval call binding the contract method 0x6fd2c80b.
//
// Solidity: function getUnbondingPeriod() view returns(uint256)
func (_Autonity *AutonitySession) GetUnbondingPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetUnbondingPeriod(&_Autonity.CallOpts)
}

// GetUnbondingPeriod is a free data retrieval call binding the contract method 0x6fd2c80b.
//
// Solidity: function getUnbondingPeriod() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetUnbondingPeriod() (*big.Int, error) {
	return _Autonity.Contract.GetUnbondingPeriod(&_Autonity.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _addr) view returns((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Autonity *AutonityCaller) GetValidator(opts *bind.CallOpts, _addr common.Address) (AutonityValidator, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getValidator", _addr)

	if err != nil {
		return *new(AutonityValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(AutonityValidator)).(*AutonityValidator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _addr) view returns((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Autonity *AutonitySession) GetValidator(_addr common.Address) (AutonityValidator, error) {
	return _Autonity.Contract.GetValidator(&_Autonity.CallOpts, _addr)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _addr) view returns((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8))
func (_Autonity *AutonityCallerSession) GetValidator(_addr common.Address) (AutonityValidator, error) {
	return _Autonity.Contract.GetValidator(&_Autonity.CallOpts, _addr)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Autonity *AutonityCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Autonity *AutonitySession) GetValidators() ([]common.Address, error) {
	return _Autonity.Contract.GetValidators(&_Autonity.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Autonity *AutonityCallerSession) GetValidators() ([]common.Address, error) {
	return _Autonity.Contract.GetValidators(&_Autonity.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(uint256)
func (_Autonity *AutonityCaller) GetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(uint256)
func (_Autonity *AutonitySession) GetVersion() (*big.Int, error) {
	return _Autonity.Contract.GetVersion(&_Autonity.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(uint256)
func (_Autonity *AutonityCallerSession) GetVersion() (*big.Int, error) {
	return _Autonity.Contract.GetVersion(&_Autonity.CallOpts)
}

// LastEpochBlock is a free data retrieval call binding the contract method 0xc2362dd5.
//
// Solidity: function lastEpochBlock() view returns(uint256)
func (_Autonity *AutonityCaller) LastEpochBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "lastEpochBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEpochBlock is a free data retrieval call binding the contract method 0xc2362dd5.
//
// Solidity: function lastEpochBlock() view returns(uint256)
func (_Autonity *AutonitySession) LastEpochBlock() (*big.Int, error) {
	return _Autonity.Contract.LastEpochBlock(&_Autonity.CallOpts)
}

// LastEpochBlock is a free data retrieval call binding the contract method 0xc2362dd5.
//
// Solidity: function lastEpochBlock() view returns(uint256)
func (_Autonity *AutonityCallerSession) LastEpochBlock() (*big.Int, error) {
	return _Autonity.Contract.LastEpochBlock(&_Autonity.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Autonity *AutonityCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Autonity *AutonitySession) Name() (string, error) {
	return _Autonity.Contract.Name(&_Autonity.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Autonity *AutonityCallerSession) Name() (string, error) {
	return _Autonity.Contract.Name(&_Autonity.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Autonity *AutonityCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Autonity *AutonitySession) Symbol() (string, error) {
	return _Autonity.Contract.Symbol(&_Autonity.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Autonity *AutonityCallerSession) Symbol() (string, error) {
	return _Autonity.Contract.Symbol(&_Autonity.CallOpts)
}

// TotalRedistributed is a free data retrieval call binding the contract method 0x9bb851c0.
//
// Solidity: function totalRedistributed() view returns(uint256)
func (_Autonity *AutonityCaller) TotalRedistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "totalRedistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRedistributed is a free data retrieval call binding the contract method 0x9bb851c0.
//
// Solidity: function totalRedistributed() view returns(uint256)
func (_Autonity *AutonitySession) TotalRedistributed() (*big.Int, error) {
	return _Autonity.Contract.TotalRedistributed(&_Autonity.CallOpts)
}

// TotalRedistributed is a free data retrieval call binding the contract method 0x9bb851c0.
//
// Solidity: function totalRedistributed() view returns(uint256)
func (_Autonity *AutonityCallerSession) TotalRedistributed() (*big.Int, error) {
	return _Autonity.Contract.TotalRedistributed(&_Autonity.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Autonity *AutonityCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Autonity.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Autonity *AutonitySession) TotalSupply() (*big.Int, error) {
	return _Autonity.Contract.TotalSupply(&_Autonity.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Autonity *AutonityCallerSession) TotalSupply() (*big.Int, error) {
	return _Autonity.Contract.TotalSupply(&_Autonity.CallOpts)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address _address) returns()
func (_Autonity *AutonityTransactor) ActivateValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "activateValidator", _address)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address _address) returns()
func (_Autonity *AutonitySession) ActivateValidator(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.ActivateValidator(&_Autonity.TransactOpts, _address)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address _address) returns()
func (_Autonity *AutonityTransactorSession) ActivateValidator(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.ActivateValidator(&_Autonity.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Autonity *AutonityTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Autonity *AutonitySession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Approve(&_Autonity.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Autonity *AutonityTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Approve(&_Autonity.TransactOpts, spender, amount)
}

// Bond is a paid mutator transaction binding the contract method 0xa515366a.
//
// Solidity: function bond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonityTransactor) Bond(opts *bind.TransactOpts, _validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "bond", _validator, _amount)
}

// Bond is a paid mutator transaction binding the contract method 0xa515366a.
//
// Solidity: function bond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonitySession) Bond(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Bond(&_Autonity.TransactOpts, _validator, _amount)
}

// Bond is a paid mutator transaction binding the contract method 0xa515366a.
//
// Solidity: function bond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonityTransactorSession) Bond(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Bond(&_Autonity.TransactOpts, _validator, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _addr, uint256 _amount) returns()
func (_Autonity *AutonityTransactor) Burn(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "burn", _addr, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _addr, uint256 _amount) returns()
func (_Autonity *AutonitySession) Burn(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Burn(&_Autonity.TransactOpts, _addr, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _addr, uint256 _amount) returns()
func (_Autonity *AutonityTransactorSession) Burn(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Burn(&_Autonity.TransactOpts, _addr, _amount)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x852c4849.
//
// Solidity: function changeCommissionRate(address _validator, uint256 _rate) returns()
func (_Autonity *AutonityTransactor) ChangeCommissionRate(opts *bind.TransactOpts, _validator common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "changeCommissionRate", _validator, _rate)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x852c4849.
//
// Solidity: function changeCommissionRate(address _validator, uint256 _rate) returns()
func (_Autonity *AutonitySession) ChangeCommissionRate(_validator common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.ChangeCommissionRate(&_Autonity.TransactOpts, _validator, _rate)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x852c4849.
//
// Solidity: function changeCommissionRate(address _validator, uint256 _rate) returns()
func (_Autonity *AutonityTransactorSession) ChangeCommissionRate(_validator common.Address, _rate *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.ChangeCommissionRate(&_Autonity.TransactOpts, _validator, _rate)
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Autonity *AutonityTransactor) CompleteContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "completeContractUpgrade")
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Autonity *AutonitySession) CompleteContractUpgrade() (*types.Transaction, error) {
	return _Autonity.Contract.CompleteContractUpgrade(&_Autonity.TransactOpts)
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Autonity *AutonityTransactorSession) CompleteContractUpgrade() (*types.Transaction, error) {
	return _Autonity.Contract.CompleteContractUpgrade(&_Autonity.TransactOpts)
}

// ComputeCommittee is a paid mutator transaction binding the contract method 0xae1f5fa0.
//
// Solidity: function computeCommittee() returns(address[])
func (_Autonity *AutonityTransactor) ComputeCommittee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "computeCommittee")
}

// ComputeCommittee is a paid mutator transaction binding the contract method 0xae1f5fa0.
//
// Solidity: function computeCommittee() returns(address[])
func (_Autonity *AutonitySession) ComputeCommittee() (*types.Transaction, error) {
	return _Autonity.Contract.ComputeCommittee(&_Autonity.TransactOpts)
}

// ComputeCommittee is a paid mutator transaction binding the contract method 0xae1f5fa0.
//
// Solidity: function computeCommittee() returns(address[])
func (_Autonity *AutonityTransactorSession) ComputeCommittee() (*types.Transaction, error) {
	return _Autonity.Contract.ComputeCommittee(&_Autonity.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool, (address,uint256)[])
func (_Autonity *AutonityTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool, (address,uint256)[])
func (_Autonity *AutonitySession) Finalize() (*types.Transaction, error) {
	return _Autonity.Contract.Finalize(&_Autonity.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool, (address,uint256)[])
func (_Autonity *AutonityTransactorSession) Finalize() (*types.Transaction, error) {
	return _Autonity.Contract.Finalize(&_Autonity.TransactOpts)
}

// FinalizeInitialization is a paid mutator transaction binding the contract method 0xd861b0e8.
//
// Solidity: function finalizeInitialization() returns()
func (_Autonity *AutonityTransactor) FinalizeInitialization(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "finalizeInitialization")
}

// FinalizeInitialization is a paid mutator transaction binding the contract method 0xd861b0e8.
//
// Solidity: function finalizeInitialization() returns()
func (_Autonity *AutonitySession) FinalizeInitialization() (*types.Transaction, error) {
	return _Autonity.Contract.FinalizeInitialization(&_Autonity.TransactOpts)
}

// FinalizeInitialization is a paid mutator transaction binding the contract method 0xd861b0e8.
//
// Solidity: function finalizeInitialization() returns()
func (_Autonity *AutonityTransactorSession) FinalizeInitialization() (*types.Transaction, error) {
	return _Autonity.Contract.FinalizeInitialization(&_Autonity.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _addr, uint256 _amount) returns()
func (_Autonity *AutonityTransactor) Mint(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "mint", _addr, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _addr, uint256 _amount) returns()
func (_Autonity *AutonitySession) Mint(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Mint(&_Autonity.TransactOpts, _addr, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _addr, uint256 _amount) returns()
func (_Autonity *AutonityTransactorSession) Mint(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Mint(&_Autonity.TransactOpts, _addr, _amount)
}

// PauseValidator is a paid mutator transaction binding the contract method 0x0ae65e7a.
//
// Solidity: function pauseValidator(address _address) returns()
func (_Autonity *AutonityTransactor) PauseValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "pauseValidator", _address)
}

// PauseValidator is a paid mutator transaction binding the contract method 0x0ae65e7a.
//
// Solidity: function pauseValidator(address _address) returns()
func (_Autonity *AutonitySession) PauseValidator(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.PauseValidator(&_Autonity.TransactOpts, _address)
}

// PauseValidator is a paid mutator transaction binding the contract method 0x0ae65e7a.
//
// Solidity: function pauseValidator(address _address) returns()
func (_Autonity *AutonityTransactorSession) PauseValidator(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.PauseValidator(&_Autonity.TransactOpts, _address)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xad722d4d.
//
// Solidity: function registerValidator(string _enode, address _oracleAddress, bytes _multisig) returns()
func (_Autonity *AutonityTransactor) RegisterValidator(opts *bind.TransactOpts, _enode string, _oracleAddress common.Address, _multisig []byte) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "registerValidator", _enode, _oracleAddress, _multisig)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xad722d4d.
//
// Solidity: function registerValidator(string _enode, address _oracleAddress, bytes _multisig) returns()
func (_Autonity *AutonitySession) RegisterValidator(_enode string, _oracleAddress common.Address, _multisig []byte) (*types.Transaction, error) {
	return _Autonity.Contract.RegisterValidator(&_Autonity.TransactOpts, _enode, _oracleAddress, _multisig)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xad722d4d.
//
// Solidity: function registerValidator(string _enode, address _oracleAddress, bytes _multisig) returns()
func (_Autonity *AutonityTransactorSession) RegisterValidator(_enode string, _oracleAddress common.Address, _multisig []byte) (*types.Transaction, error) {
	return _Autonity.Contract.RegisterValidator(&_Autonity.TransactOpts, _enode, _oracleAddress, _multisig)
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Autonity *AutonityTransactor) ResetContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "resetContractUpgrade")
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Autonity *AutonitySession) ResetContractUpgrade() (*types.Transaction, error) {
	return _Autonity.Contract.ResetContractUpgrade(&_Autonity.TransactOpts)
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Autonity *AutonityTransactorSession) ResetContractUpgrade() (*types.Transaction, error) {
	return _Autonity.Contract.ResetContractUpgrade(&_Autonity.TransactOpts)
}

// SetAccountabilityContract is a paid mutator transaction binding the contract method 0x1250a28d.
//
// Solidity: function setAccountabilityContract(address _address) returns()
func (_Autonity *AutonityTransactor) SetAccountabilityContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setAccountabilityContract", _address)
}

// SetAccountabilityContract is a paid mutator transaction binding the contract method 0x1250a28d.
//
// Solidity: function setAccountabilityContract(address _address) returns()
func (_Autonity *AutonitySession) SetAccountabilityContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetAccountabilityContract(&_Autonity.TransactOpts, _address)
}

// SetAccountabilityContract is a paid mutator transaction binding the contract method 0x1250a28d.
//
// Solidity: function setAccountabilityContract(address _address) returns()
func (_Autonity *AutonityTransactorSession) SetAccountabilityContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetAccountabilityContract(&_Autonity.TransactOpts, _address)
}

// SetAcuContract is a paid mutator transaction binding the contract method 0xd372c07e.
//
// Solidity: function setAcuContract(address _address) returns()
func (_Autonity *AutonityTransactor) SetAcuContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setAcuContract", _address)
}

// SetAcuContract is a paid mutator transaction binding the contract method 0xd372c07e.
//
// Solidity: function setAcuContract(address _address) returns()
func (_Autonity *AutonitySession) SetAcuContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetAcuContract(&_Autonity.TransactOpts, _address)
}

// SetAcuContract is a paid mutator transaction binding the contract method 0xd372c07e.
//
// Solidity: function setAcuContract(address _address) returns()
func (_Autonity *AutonityTransactorSession) SetAcuContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetAcuContract(&_Autonity.TransactOpts, _address)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x8bac7dad.
//
// Solidity: function setCommitteeSize(uint256 _size) returns()
func (_Autonity *AutonityTransactor) SetCommitteeSize(opts *bind.TransactOpts, _size *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setCommitteeSize", _size)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x8bac7dad.
//
// Solidity: function setCommitteeSize(uint256 _size) returns()
func (_Autonity *AutonitySession) SetCommitteeSize(_size *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetCommitteeSize(&_Autonity.TransactOpts, _size)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x8bac7dad.
//
// Solidity: function setCommitteeSize(uint256 _size) returns()
func (_Autonity *AutonityTransactorSession) SetCommitteeSize(_size *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetCommitteeSize(&_Autonity.TransactOpts, _size)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _period) returns()
func (_Autonity *AutonityTransactor) SetEpochPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setEpochPeriod", _period)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _period) returns()
func (_Autonity *AutonitySession) SetEpochPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetEpochPeriod(&_Autonity.TransactOpts, _period)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _period) returns()
func (_Autonity *AutonityTransactorSession) SetEpochPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetEpochPeriod(&_Autonity.TransactOpts, _period)
}

// SetMinimumBaseFee is a paid mutator transaction binding the contract method 0xcb696f54.
//
// Solidity: function setMinimumBaseFee(uint256 _price) returns()
func (_Autonity *AutonityTransactor) SetMinimumBaseFee(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setMinimumBaseFee", _price)
}

// SetMinimumBaseFee is a paid mutator transaction binding the contract method 0xcb696f54.
//
// Solidity: function setMinimumBaseFee(uint256 _price) returns()
func (_Autonity *AutonitySession) SetMinimumBaseFee(_price *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetMinimumBaseFee(&_Autonity.TransactOpts, _price)
}

// SetMinimumBaseFee is a paid mutator transaction binding the contract method 0xcb696f54.
//
// Solidity: function setMinimumBaseFee(uint256 _price) returns()
func (_Autonity *AutonityTransactorSession) SetMinimumBaseFee(_price *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetMinimumBaseFee(&_Autonity.TransactOpts, _price)
}

// SetOperatorAccount is a paid mutator transaction binding the contract method 0x520fdbbc.
//
// Solidity: function setOperatorAccount(address _account) returns()
func (_Autonity *AutonityTransactor) SetOperatorAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setOperatorAccount", _account)
}

// SetOperatorAccount is a paid mutator transaction binding the contract method 0x520fdbbc.
//
// Solidity: function setOperatorAccount(address _account) returns()
func (_Autonity *AutonitySession) SetOperatorAccount(_account common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetOperatorAccount(&_Autonity.TransactOpts, _account)
}

// SetOperatorAccount is a paid mutator transaction binding the contract method 0x520fdbbc.
//
// Solidity: function setOperatorAccount(address _account) returns()
func (_Autonity *AutonityTransactorSession) SetOperatorAccount(_account common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetOperatorAccount(&_Autonity.TransactOpts, _account)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(address _address) returns()
func (_Autonity *AutonityTransactor) SetOracleContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setOracleContract", _address)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(address _address) returns()
func (_Autonity *AutonitySession) SetOracleContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetOracleContract(&_Autonity.TransactOpts, _address)
}

// SetOracleContract is a paid mutator transaction binding the contract method 0x496ccd9b.
//
// Solidity: function setOracleContract(address _address) returns()
func (_Autonity *AutonityTransactorSession) SetOracleContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetOracleContract(&_Autonity.TransactOpts, _address)
}

// SetStabilizationContract is a paid mutator transaction binding the contract method 0xcfd19fb9.
//
// Solidity: function setStabilizationContract(address _address) returns()
func (_Autonity *AutonityTransactor) SetStabilizationContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setStabilizationContract", _address)
}

// SetStabilizationContract is a paid mutator transaction binding the contract method 0xcfd19fb9.
//
// Solidity: function setStabilizationContract(address _address) returns()
func (_Autonity *AutonitySession) SetStabilizationContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetStabilizationContract(&_Autonity.TransactOpts, _address)
}

// SetStabilizationContract is a paid mutator transaction binding the contract method 0xcfd19fb9.
//
// Solidity: function setStabilizationContract(address _address) returns()
func (_Autonity *AutonityTransactorSession) SetStabilizationContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetStabilizationContract(&_Autonity.TransactOpts, _address)
}

// SetSupplyControlContract is a paid mutator transaction binding the contract method 0xb3ecbadd.
//
// Solidity: function setSupplyControlContract(address _address) returns()
func (_Autonity *AutonityTransactor) SetSupplyControlContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setSupplyControlContract", _address)
}

// SetSupplyControlContract is a paid mutator transaction binding the contract method 0xb3ecbadd.
//
// Solidity: function setSupplyControlContract(address _address) returns()
func (_Autonity *AutonitySession) SetSupplyControlContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetSupplyControlContract(&_Autonity.TransactOpts, _address)
}

// SetSupplyControlContract is a paid mutator transaction binding the contract method 0xb3ecbadd.
//
// Solidity: function setSupplyControlContract(address _address) returns()
func (_Autonity *AutonityTransactorSession) SetSupplyControlContract(_address common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetSupplyControlContract(&_Autonity.TransactOpts, _address)
}

// SetTreasuryAccount is a paid mutator transaction binding the contract method 0xd886f8a2.
//
// Solidity: function setTreasuryAccount(address _account) returns()
func (_Autonity *AutonityTransactor) SetTreasuryAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setTreasuryAccount", _account)
}

// SetTreasuryAccount is a paid mutator transaction binding the contract method 0xd886f8a2.
//
// Solidity: function setTreasuryAccount(address _account) returns()
func (_Autonity *AutonitySession) SetTreasuryAccount(_account common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetTreasuryAccount(&_Autonity.TransactOpts, _account)
}

// SetTreasuryAccount is a paid mutator transaction binding the contract method 0xd886f8a2.
//
// Solidity: function setTreasuryAccount(address _account) returns()
func (_Autonity *AutonityTransactorSession) SetTreasuryAccount(_account common.Address) (*types.Transaction, error) {
	return _Autonity.Contract.SetTreasuryAccount(&_Autonity.TransactOpts, _account)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Autonity *AutonityTransactor) SetTreasuryFee(opts *bind.TransactOpts, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setTreasuryFee", _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Autonity *AutonitySession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetTreasuryFee(&_Autonity.TransactOpts, _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Autonity *AutonityTransactorSession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetTreasuryFee(&_Autonity.TransactOpts, _treasuryFee)
}

// SetUnbondingPeriod is a paid mutator transaction binding the contract method 0x114eaf55.
//
// Solidity: function setUnbondingPeriod(uint256 _period) returns()
func (_Autonity *AutonityTransactor) SetUnbondingPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "setUnbondingPeriod", _period)
}

// SetUnbondingPeriod is a paid mutator transaction binding the contract method 0x114eaf55.
//
// Solidity: function setUnbondingPeriod(uint256 _period) returns()
func (_Autonity *AutonitySession) SetUnbondingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetUnbondingPeriod(&_Autonity.TransactOpts, _period)
}

// SetUnbondingPeriod is a paid mutator transaction binding the contract method 0x114eaf55.
//
// Solidity: function setUnbondingPeriod(uint256 _period) returns()
func (_Autonity *AutonityTransactorSession) SetUnbondingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.SetUnbondingPeriod(&_Autonity.TransactOpts, _period)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Autonity *AutonityTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Autonity *AutonitySession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Transfer(&_Autonity.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Autonity *AutonityTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Transfer(&_Autonity.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Autonity *AutonityTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Autonity *AutonitySession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.TransferFrom(&_Autonity.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Autonity *AutonityTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.TransferFrom(&_Autonity.TransactOpts, sender, recipient, amount)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonityTransactor) Unbond(opts *bind.TransactOpts, _validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "unbond", _validator, _amount)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonitySession) Unbond(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Unbond(&_Autonity.TransactOpts, _validator, _amount)
}

// Unbond is a paid mutator transaction binding the contract method 0xa5d059ca.
//
// Solidity: function unbond(address _validator, uint256 _amount) returns()
func (_Autonity *AutonityTransactorSession) Unbond(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Autonity.Contract.Unbond(&_Autonity.TransactOpts, _validator, _amount)
}

// UpdateValidatorAndTransferSlashedFunds is a paid mutator transaction binding the contract method 0x6e74c554.
//
// Solidity: function updateValidatorAndTransferSlashedFunds((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8) _val) returns()
func (_Autonity *AutonityTransactor) UpdateValidatorAndTransferSlashedFunds(opts *bind.TransactOpts, _val AutonityValidator) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "updateValidatorAndTransferSlashedFunds", _val)
}

// UpdateValidatorAndTransferSlashedFunds is a paid mutator transaction binding the contract method 0x6e74c554.
//
// Solidity: function updateValidatorAndTransferSlashedFunds((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8) _val) returns()
func (_Autonity *AutonitySession) UpdateValidatorAndTransferSlashedFunds(_val AutonityValidator) (*types.Transaction, error) {
	return _Autonity.Contract.UpdateValidatorAndTransferSlashedFunds(&_Autonity.TransactOpts, _val)
}

// UpdateValidatorAndTransferSlashedFunds is a paid mutator transaction binding the contract method 0x6e74c554.
//
// Solidity: function updateValidatorAndTransferSlashedFunds((address,address,address,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint8) _val) returns()
func (_Autonity *AutonityTransactorSession) UpdateValidatorAndTransferSlashedFunds(_val AutonityValidator) (*types.Transaction, error) {
	return _Autonity.Contract.UpdateValidatorAndTransferSlashedFunds(&_Autonity.TransactOpts, _val)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Autonity *AutonityTransactor) UpgradeContract(opts *bind.TransactOpts, _bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Autonity.contract.Transact(opts, "upgradeContract", _bytecode, _abi)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Autonity *AutonitySession) UpgradeContract(_bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Autonity.Contract.UpgradeContract(&_Autonity.TransactOpts, _bytecode, _abi)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Autonity *AutonityTransactorSession) UpgradeContract(_bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Autonity.Contract.UpgradeContract(&_Autonity.TransactOpts, _bytecode, _abi)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Autonity *AutonityTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Autonity.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Autonity *AutonitySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Autonity.Contract.Fallback(&_Autonity.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Autonity *AutonityTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Autonity.Contract.Fallback(&_Autonity.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Autonity *AutonityTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Autonity.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Autonity *AutonitySession) Receive() (*types.Transaction, error) {
	return _Autonity.Contract.Receive(&_Autonity.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Autonity *AutonityTransactorSession) Receive() (*types.Transaction, error) {
	return _Autonity.Contract.Receive(&_Autonity.TransactOpts)
}

// AutonityActivatedValidatorIterator is returned from FilterActivatedValidator and is used to iterate over the raw logs and unpacked data for ActivatedValidator events raised by the Autonity contract.
type AutonityActivatedValidatorIterator struct {
	Event *AutonityActivatedValidator // Event containing the contract specifics and raw log

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
func (it *AutonityActivatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityActivatedValidator)
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
		it.Event = new(AutonityActivatedValidator)
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
func (it *AutonityActivatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityActivatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityActivatedValidator represents a ActivatedValidator event raised by the Autonity contract.
type AutonityActivatedValidator struct {
	Treasury       common.Address
	Addr           common.Address
	EffectiveBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterActivatedValidator is a free log retrieval operation binding the contract event 0x60fcbf2d07dc712a93e59fb28f1edb626d7c2497c57ba71a8c0b3999ecb9a3b5.
//
// Solidity: event ActivatedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) FilterActivatedValidator(opts *bind.FilterOpts, treasury []common.Address, addr []common.Address) (*AutonityActivatedValidatorIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "ActivatedValidator", treasuryRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &AutonityActivatedValidatorIterator{contract: _Autonity.contract, event: "ActivatedValidator", logs: logs, sub: sub}, nil
}

// WatchActivatedValidator is a free log subscription operation binding the contract event 0x60fcbf2d07dc712a93e59fb28f1edb626d7c2497c57ba71a8c0b3999ecb9a3b5.
//
// Solidity: event ActivatedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) WatchActivatedValidator(opts *bind.WatchOpts, sink chan<- *AutonityActivatedValidator, treasury []common.Address, addr []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "ActivatedValidator", treasuryRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityActivatedValidator)
				if err := _Autonity.contract.UnpackLog(event, "ActivatedValidator", log); err != nil {
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

// ParseActivatedValidator is a log parse operation binding the contract event 0x60fcbf2d07dc712a93e59fb28f1edb626d7c2497c57ba71a8c0b3999ecb9a3b5.
//
// Solidity: event ActivatedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) ParseActivatedValidator(log types.Log) (*AutonityActivatedValidator, error) {
	event := new(AutonityActivatedValidator)
	if err := _Autonity.contract.UnpackLog(event, "ActivatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Autonity contract.
type AutonityApprovalIterator struct {
	Event *AutonityApproval // Event containing the contract specifics and raw log

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
func (it *AutonityApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityApproval)
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
		it.Event = new(AutonityApproval)
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
func (it *AutonityApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityApproval represents a Approval event raised by the Autonity contract.
type AutonityApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Autonity *AutonityFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AutonityApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AutonityApprovalIterator{contract: _Autonity.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Autonity *AutonityFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AutonityApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityApproval)
				if err := _Autonity.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Autonity *AutonityFilterer) ParseApproval(log types.Log) (*AutonityApproval, error) {
	event := new(AutonityApproval)
	if err := _Autonity.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityBurnedStakeIterator is returned from FilterBurnedStake and is used to iterate over the raw logs and unpacked data for BurnedStake events raised by the Autonity contract.
type AutonityBurnedStakeIterator struct {
	Event *AutonityBurnedStake // Event containing the contract specifics and raw log

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
func (it *AutonityBurnedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityBurnedStake)
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
		it.Event = new(AutonityBurnedStake)
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
func (it *AutonityBurnedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityBurnedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityBurnedStake represents a BurnedStake event raised by the Autonity contract.
type AutonityBurnedStake struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnedStake is a free log retrieval operation binding the contract event 0x5024dbeedf0c06664c9bd7be836915730c955e936972c020683dadf11d5488a3.
//
// Solidity: event BurnedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) FilterBurnedStake(opts *bind.FilterOpts, addr []common.Address) (*AutonityBurnedStakeIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "BurnedStake", addrRule)
	if err != nil {
		return nil, err
	}
	return &AutonityBurnedStakeIterator{contract: _Autonity.contract, event: "BurnedStake", logs: logs, sub: sub}, nil
}

// WatchBurnedStake is a free log subscription operation binding the contract event 0x5024dbeedf0c06664c9bd7be836915730c955e936972c020683dadf11d5488a3.
//
// Solidity: event BurnedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) WatchBurnedStake(opts *bind.WatchOpts, sink chan<- *AutonityBurnedStake, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "BurnedStake", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityBurnedStake)
				if err := _Autonity.contract.UnpackLog(event, "BurnedStake", log); err != nil {
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

// ParseBurnedStake is a log parse operation binding the contract event 0x5024dbeedf0c06664c9bd7be836915730c955e936972c020683dadf11d5488a3.
//
// Solidity: event BurnedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) ParseBurnedStake(log types.Log) (*AutonityBurnedStake, error) {
	event := new(AutonityBurnedStake)
	if err := _Autonity.contract.UnpackLog(event, "BurnedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityCommissionRateChangeIterator is returned from FilterCommissionRateChange and is used to iterate over the raw logs and unpacked data for CommissionRateChange events raised by the Autonity contract.
type AutonityCommissionRateChangeIterator struct {
	Event *AutonityCommissionRateChange // Event containing the contract specifics and raw log

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
func (it *AutonityCommissionRateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityCommissionRateChange)
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
		it.Event = new(AutonityCommissionRateChange)
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
func (it *AutonityCommissionRateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityCommissionRateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityCommissionRateChange represents a CommissionRateChange event raised by the Autonity contract.
type AutonityCommissionRateChange struct {
	Validator common.Address
	Rate      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateChange is a free log retrieval operation binding the contract event 0x4fba51c92fa3d6ad8374d394f6cd5766857552e153d7384a8f23aa4ce9a8a7cf.
//
// Solidity: event CommissionRateChange(address indexed validator, uint256 rate)
func (_Autonity *AutonityFilterer) FilterCommissionRateChange(opts *bind.FilterOpts, validator []common.Address) (*AutonityCommissionRateChangeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "CommissionRateChange", validatorRule)
	if err != nil {
		return nil, err
	}
	return &AutonityCommissionRateChangeIterator{contract: _Autonity.contract, event: "CommissionRateChange", logs: logs, sub: sub}, nil
}

// WatchCommissionRateChange is a free log subscription operation binding the contract event 0x4fba51c92fa3d6ad8374d394f6cd5766857552e153d7384a8f23aa4ce9a8a7cf.
//
// Solidity: event CommissionRateChange(address indexed validator, uint256 rate)
func (_Autonity *AutonityFilterer) WatchCommissionRateChange(opts *bind.WatchOpts, sink chan<- *AutonityCommissionRateChange, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "CommissionRateChange", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityCommissionRateChange)
				if err := _Autonity.contract.UnpackLog(event, "CommissionRateChange", log); err != nil {
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

// ParseCommissionRateChange is a log parse operation binding the contract event 0x4fba51c92fa3d6ad8374d394f6cd5766857552e153d7384a8f23aa4ce9a8a7cf.
//
// Solidity: event CommissionRateChange(address indexed validator, uint256 rate)
func (_Autonity *AutonityFilterer) ParseCommissionRateChange(log types.Log) (*AutonityCommissionRateChange, error) {
	event := new(AutonityCommissionRateChange)
	if err := _Autonity.contract.UnpackLog(event, "CommissionRateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityEpochPeriodUpdatedIterator is returned from FilterEpochPeriodUpdated and is used to iterate over the raw logs and unpacked data for EpochPeriodUpdated events raised by the Autonity contract.
type AutonityEpochPeriodUpdatedIterator struct {
	Event *AutonityEpochPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *AutonityEpochPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityEpochPeriodUpdated)
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
		it.Event = new(AutonityEpochPeriodUpdated)
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
func (it *AutonityEpochPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityEpochPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityEpochPeriodUpdated represents a EpochPeriodUpdated event raised by the Autonity contract.
type AutonityEpochPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochPeriodUpdated is a free log retrieval operation binding the contract event 0xd7f1279ded354dbf22a69fcc2fd661763a6e2956a5d2891af9410af880fa5f81.
//
// Solidity: event EpochPeriodUpdated(uint256 period)
func (_Autonity *AutonityFilterer) FilterEpochPeriodUpdated(opts *bind.FilterOpts) (*AutonityEpochPeriodUpdatedIterator, error) {

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "EpochPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &AutonityEpochPeriodUpdatedIterator{contract: _Autonity.contract, event: "EpochPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochPeriodUpdated is a free log subscription operation binding the contract event 0xd7f1279ded354dbf22a69fcc2fd661763a6e2956a5d2891af9410af880fa5f81.
//
// Solidity: event EpochPeriodUpdated(uint256 period)
func (_Autonity *AutonityFilterer) WatchEpochPeriodUpdated(opts *bind.WatchOpts, sink chan<- *AutonityEpochPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "EpochPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityEpochPeriodUpdated)
				if err := _Autonity.contract.UnpackLog(event, "EpochPeriodUpdated", log); err != nil {
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

// ParseEpochPeriodUpdated is a log parse operation binding the contract event 0xd7f1279ded354dbf22a69fcc2fd661763a6e2956a5d2891af9410af880fa5f81.
//
// Solidity: event EpochPeriodUpdated(uint256 period)
func (_Autonity *AutonityFilterer) ParseEpochPeriodUpdated(log types.Log) (*AutonityEpochPeriodUpdated, error) {
	event := new(AutonityEpochPeriodUpdated)
	if err := _Autonity.contract.UnpackLog(event, "EpochPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityMinimumBaseFeeUpdatedIterator is returned from FilterMinimumBaseFeeUpdated and is used to iterate over the raw logs and unpacked data for MinimumBaseFeeUpdated events raised by the Autonity contract.
type AutonityMinimumBaseFeeUpdatedIterator struct {
	Event *AutonityMinimumBaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AutonityMinimumBaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityMinimumBaseFeeUpdated)
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
		it.Event = new(AutonityMinimumBaseFeeUpdated)
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
func (it *AutonityMinimumBaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityMinimumBaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityMinimumBaseFeeUpdated represents a MinimumBaseFeeUpdated event raised by the Autonity contract.
type AutonityMinimumBaseFeeUpdated struct {
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinimumBaseFeeUpdated is a free log retrieval operation binding the contract event 0x1f4d2fc7529047a5bd96d3229bfea127fd18b7748f13586e097c69fccd389128.
//
// Solidity: event MinimumBaseFeeUpdated(uint256 gasPrice)
func (_Autonity *AutonityFilterer) FilterMinimumBaseFeeUpdated(opts *bind.FilterOpts) (*AutonityMinimumBaseFeeUpdatedIterator, error) {

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "MinimumBaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AutonityMinimumBaseFeeUpdatedIterator{contract: _Autonity.contract, event: "MinimumBaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumBaseFeeUpdated is a free log subscription operation binding the contract event 0x1f4d2fc7529047a5bd96d3229bfea127fd18b7748f13586e097c69fccd389128.
//
// Solidity: event MinimumBaseFeeUpdated(uint256 gasPrice)
func (_Autonity *AutonityFilterer) WatchMinimumBaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *AutonityMinimumBaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "MinimumBaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityMinimumBaseFeeUpdated)
				if err := _Autonity.contract.UnpackLog(event, "MinimumBaseFeeUpdated", log); err != nil {
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

// ParseMinimumBaseFeeUpdated is a log parse operation binding the contract event 0x1f4d2fc7529047a5bd96d3229bfea127fd18b7748f13586e097c69fccd389128.
//
// Solidity: event MinimumBaseFeeUpdated(uint256 gasPrice)
func (_Autonity *AutonityFilterer) ParseMinimumBaseFeeUpdated(log types.Log) (*AutonityMinimumBaseFeeUpdated, error) {
	event := new(AutonityMinimumBaseFeeUpdated)
	if err := _Autonity.contract.UnpackLog(event, "MinimumBaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityMintedStakeIterator is returned from FilterMintedStake and is used to iterate over the raw logs and unpacked data for MintedStake events raised by the Autonity contract.
type AutonityMintedStakeIterator struct {
	Event *AutonityMintedStake // Event containing the contract specifics and raw log

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
func (it *AutonityMintedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityMintedStake)
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
		it.Event = new(AutonityMintedStake)
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
func (it *AutonityMintedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityMintedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityMintedStake represents a MintedStake event raised by the Autonity contract.
type AutonityMintedStake struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintedStake is a free log retrieval operation binding the contract event 0x48490b4407bb949b708ec5f514b4167f08f4969baaf78d53b05028adf369bfcf.
//
// Solidity: event MintedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) FilterMintedStake(opts *bind.FilterOpts, addr []common.Address) (*AutonityMintedStakeIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "MintedStake", addrRule)
	if err != nil {
		return nil, err
	}
	return &AutonityMintedStakeIterator{contract: _Autonity.contract, event: "MintedStake", logs: logs, sub: sub}, nil
}

// WatchMintedStake is a free log subscription operation binding the contract event 0x48490b4407bb949b708ec5f514b4167f08f4969baaf78d53b05028adf369bfcf.
//
// Solidity: event MintedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) WatchMintedStake(opts *bind.WatchOpts, sink chan<- *AutonityMintedStake, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "MintedStake", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityMintedStake)
				if err := _Autonity.contract.UnpackLog(event, "MintedStake", log); err != nil {
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

// ParseMintedStake is a log parse operation binding the contract event 0x48490b4407bb949b708ec5f514b4167f08f4969baaf78d53b05028adf369bfcf.
//
// Solidity: event MintedStake(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) ParseMintedStake(log types.Log) (*AutonityMintedStake, error) {
	event := new(AutonityMintedStake)
	if err := _Autonity.contract.UnpackLog(event, "MintedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityNewBondingRequestIterator is returned from FilterNewBondingRequest and is used to iterate over the raw logs and unpacked data for NewBondingRequest events raised by the Autonity contract.
type AutonityNewBondingRequestIterator struct {
	Event *AutonityNewBondingRequest // Event containing the contract specifics and raw log

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
func (it *AutonityNewBondingRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityNewBondingRequest)
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
		it.Event = new(AutonityNewBondingRequest)
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
func (it *AutonityNewBondingRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityNewBondingRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityNewBondingRequest represents a NewBondingRequest event raised by the Autonity contract.
type AutonityNewBondingRequest struct {
	Validator  common.Address
	Delegator  common.Address
	SelfBonded bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBondingRequest is a free log retrieval operation binding the contract event 0xc46aaee12f38035617ad448c04a7956119f7c7ed395ecc347b898817451ddb8d.
//
// Solidity: event NewBondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) FilterNewBondingRequest(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*AutonityNewBondingRequestIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "NewBondingRequest", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &AutonityNewBondingRequestIterator{contract: _Autonity.contract, event: "NewBondingRequest", logs: logs, sub: sub}, nil
}

// WatchNewBondingRequest is a free log subscription operation binding the contract event 0xc46aaee12f38035617ad448c04a7956119f7c7ed395ecc347b898817451ddb8d.
//
// Solidity: event NewBondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) WatchNewBondingRequest(opts *bind.WatchOpts, sink chan<- *AutonityNewBondingRequest, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "NewBondingRequest", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityNewBondingRequest)
				if err := _Autonity.contract.UnpackLog(event, "NewBondingRequest", log); err != nil {
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

// ParseNewBondingRequest is a log parse operation binding the contract event 0xc46aaee12f38035617ad448c04a7956119f7c7ed395ecc347b898817451ddb8d.
//
// Solidity: event NewBondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) ParseNewBondingRequest(log types.Log) (*AutonityNewBondingRequest, error) {
	event := new(AutonityNewBondingRequest)
	if err := _Autonity.contract.UnpackLog(event, "NewBondingRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityNewEpochIterator is returned from FilterNewEpoch and is used to iterate over the raw logs and unpacked data for NewEpoch events raised by the Autonity contract.
type AutonityNewEpochIterator struct {
	Event *AutonityNewEpoch // Event containing the contract specifics and raw log

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
func (it *AutonityNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityNewEpoch)
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
		it.Event = new(AutonityNewEpoch)
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
func (it *AutonityNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityNewEpoch represents a NewEpoch event raised by the Autonity contract.
type AutonityNewEpoch struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewEpoch is a free log retrieval operation binding the contract event 0xebad8099c467528a56c98b63c8d476d251cf1ffb4c75db94b4d23fa2b6a1e335.
//
// Solidity: event NewEpoch(uint256 epoch)
func (_Autonity *AutonityFilterer) FilterNewEpoch(opts *bind.FilterOpts) (*AutonityNewEpochIterator, error) {

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return &AutonityNewEpochIterator{contract: _Autonity.contract, event: "NewEpoch", logs: logs, sub: sub}, nil
}

// WatchNewEpoch is a free log subscription operation binding the contract event 0xebad8099c467528a56c98b63c8d476d251cf1ffb4c75db94b4d23fa2b6a1e335.
//
// Solidity: event NewEpoch(uint256 epoch)
func (_Autonity *AutonityFilterer) WatchNewEpoch(opts *bind.WatchOpts, sink chan<- *AutonityNewEpoch) (event.Subscription, error) {

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityNewEpoch)
				if err := _Autonity.contract.UnpackLog(event, "NewEpoch", log); err != nil {
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

// ParseNewEpoch is a log parse operation binding the contract event 0xebad8099c467528a56c98b63c8d476d251cf1ffb4c75db94b4d23fa2b6a1e335.
//
// Solidity: event NewEpoch(uint256 epoch)
func (_Autonity *AutonityFilterer) ParseNewEpoch(log types.Log) (*AutonityNewEpoch, error) {
	event := new(AutonityNewEpoch)
	if err := _Autonity.contract.UnpackLog(event, "NewEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityNewUnbondingRequestIterator is returned from FilterNewUnbondingRequest and is used to iterate over the raw logs and unpacked data for NewUnbondingRequest events raised by the Autonity contract.
type AutonityNewUnbondingRequestIterator struct {
	Event *AutonityNewUnbondingRequest // Event containing the contract specifics and raw log

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
func (it *AutonityNewUnbondingRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityNewUnbondingRequest)
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
		it.Event = new(AutonityNewUnbondingRequest)
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
func (it *AutonityNewUnbondingRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityNewUnbondingRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityNewUnbondingRequest represents a NewUnbondingRequest event raised by the Autonity contract.
type AutonityNewUnbondingRequest struct {
	Validator  common.Address
	Delegator  common.Address
	SelfBonded bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUnbondingRequest is a free log retrieval operation binding the contract event 0x63f8870909f7c59c9c4932bf98dbd491647c8d2e89ca0a032aacdd943a13e2fc.
//
// Solidity: event NewUnbondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) FilterNewUnbondingRequest(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*AutonityNewUnbondingRequestIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "NewUnbondingRequest", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &AutonityNewUnbondingRequestIterator{contract: _Autonity.contract, event: "NewUnbondingRequest", logs: logs, sub: sub}, nil
}

// WatchNewUnbondingRequest is a free log subscription operation binding the contract event 0x63f8870909f7c59c9c4932bf98dbd491647c8d2e89ca0a032aacdd943a13e2fc.
//
// Solidity: event NewUnbondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) WatchNewUnbondingRequest(opts *bind.WatchOpts, sink chan<- *AutonityNewUnbondingRequest, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "NewUnbondingRequest", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityNewUnbondingRequest)
				if err := _Autonity.contract.UnpackLog(event, "NewUnbondingRequest", log); err != nil {
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

// ParseNewUnbondingRequest is a log parse operation binding the contract event 0x63f8870909f7c59c9c4932bf98dbd491647c8d2e89ca0a032aacdd943a13e2fc.
//
// Solidity: event NewUnbondingRequest(address indexed validator, address indexed delegator, bool selfBonded, uint256 amount)
func (_Autonity *AutonityFilterer) ParseNewUnbondingRequest(log types.Log) (*AutonityNewUnbondingRequest, error) {
	event := new(AutonityNewUnbondingRequest)
	if err := _Autonity.contract.UnpackLog(event, "NewUnbondingRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityPausedValidatorIterator is returned from FilterPausedValidator and is used to iterate over the raw logs and unpacked data for PausedValidator events raised by the Autonity contract.
type AutonityPausedValidatorIterator struct {
	Event *AutonityPausedValidator // Event containing the contract specifics and raw log

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
func (it *AutonityPausedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityPausedValidator)
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
		it.Event = new(AutonityPausedValidator)
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
func (it *AutonityPausedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityPausedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityPausedValidator represents a PausedValidator event raised by the Autonity contract.
type AutonityPausedValidator struct {
	Treasury       common.Address
	Addr           common.Address
	EffectiveBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPausedValidator is a free log retrieval operation binding the contract event 0x75bdcdbe540758778e669d108fbcb7ede734f27f46e4e5525eeb8ecf91849a9c.
//
// Solidity: event PausedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) FilterPausedValidator(opts *bind.FilterOpts, treasury []common.Address, addr []common.Address) (*AutonityPausedValidatorIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "PausedValidator", treasuryRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &AutonityPausedValidatorIterator{contract: _Autonity.contract, event: "PausedValidator", logs: logs, sub: sub}, nil
}

// WatchPausedValidator is a free log subscription operation binding the contract event 0x75bdcdbe540758778e669d108fbcb7ede734f27f46e4e5525eeb8ecf91849a9c.
//
// Solidity: event PausedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) WatchPausedValidator(opts *bind.WatchOpts, sink chan<- *AutonityPausedValidator, treasury []common.Address, addr []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "PausedValidator", treasuryRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityPausedValidator)
				if err := _Autonity.contract.UnpackLog(event, "PausedValidator", log); err != nil {
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

// ParsePausedValidator is a log parse operation binding the contract event 0x75bdcdbe540758778e669d108fbcb7ede734f27f46e4e5525eeb8ecf91849a9c.
//
// Solidity: event PausedValidator(address indexed treasury, address indexed addr, uint256 effectiveBlock)
func (_Autonity *AutonityFilterer) ParsePausedValidator(log types.Log) (*AutonityPausedValidator, error) {
	event := new(AutonityPausedValidator)
	if err := _Autonity.contract.UnpackLog(event, "PausedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityRegisteredValidatorIterator is returned from FilterRegisteredValidator and is used to iterate over the raw logs and unpacked data for RegisteredValidator events raised by the Autonity contract.
type AutonityRegisteredValidatorIterator struct {
	Event *AutonityRegisteredValidator // Event containing the contract specifics and raw log

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
func (it *AutonityRegisteredValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityRegisteredValidator)
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
		it.Event = new(AutonityRegisteredValidator)
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
func (it *AutonityRegisteredValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityRegisteredValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityRegisteredValidator represents a RegisteredValidator event raised by the Autonity contract.
type AutonityRegisteredValidator struct {
	Treasury       common.Address
	Addr           common.Address
	OracleAddress  common.Address
	Enode          string
	LiquidContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegisteredValidator is a free log retrieval operation binding the contract event 0x8ad8bd2eb6950e5f332fd3a6dca48cb358ecfe3057848902b98cbdfe455c915c.
//
// Solidity: event RegisteredValidator(address treasury, address addr, address oracleAddress, string enode, address liquidContract)
func (_Autonity *AutonityFilterer) FilterRegisteredValidator(opts *bind.FilterOpts) (*AutonityRegisteredValidatorIterator, error) {

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "RegisteredValidator")
	if err != nil {
		return nil, err
	}
	return &AutonityRegisteredValidatorIterator{contract: _Autonity.contract, event: "RegisteredValidator", logs: logs, sub: sub}, nil
}

// WatchRegisteredValidator is a free log subscription operation binding the contract event 0x8ad8bd2eb6950e5f332fd3a6dca48cb358ecfe3057848902b98cbdfe455c915c.
//
// Solidity: event RegisteredValidator(address treasury, address addr, address oracleAddress, string enode, address liquidContract)
func (_Autonity *AutonityFilterer) WatchRegisteredValidator(opts *bind.WatchOpts, sink chan<- *AutonityRegisteredValidator) (event.Subscription, error) {

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "RegisteredValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityRegisteredValidator)
				if err := _Autonity.contract.UnpackLog(event, "RegisteredValidator", log); err != nil {
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

// ParseRegisteredValidator is a log parse operation binding the contract event 0x8ad8bd2eb6950e5f332fd3a6dca48cb358ecfe3057848902b98cbdfe455c915c.
//
// Solidity: event RegisteredValidator(address treasury, address addr, address oracleAddress, string enode, address liquidContract)
func (_Autonity *AutonityFilterer) ParseRegisteredValidator(log types.Log) (*AutonityRegisteredValidator, error) {
	event := new(AutonityRegisteredValidator)
	if err := _Autonity.contract.UnpackLog(event, "RegisteredValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityRewardedIterator is returned from FilterRewarded and is used to iterate over the raw logs and unpacked data for Rewarded events raised by the Autonity contract.
type AutonityRewardedIterator struct {
	Event *AutonityRewarded // Event containing the contract specifics and raw log

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
func (it *AutonityRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityRewarded)
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
		it.Event = new(AutonityRewarded)
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
func (it *AutonityRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityRewarded represents a Rewarded event raised by the Autonity contract.
type AutonityRewarded struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewarded is a free log retrieval operation binding the contract event 0xb3b7a071186534c03b40695710096f289fd4ed6c1a374aff0bb648955e4fe563.
//
// Solidity: event Rewarded(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) FilterRewarded(opts *bind.FilterOpts, addr []common.Address) (*AutonityRewardedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "Rewarded", addrRule)
	if err != nil {
		return nil, err
	}
	return &AutonityRewardedIterator{contract: _Autonity.contract, event: "Rewarded", logs: logs, sub: sub}, nil
}

// WatchRewarded is a free log subscription operation binding the contract event 0xb3b7a071186534c03b40695710096f289fd4ed6c1a374aff0bb648955e4fe563.
//
// Solidity: event Rewarded(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) WatchRewarded(opts *bind.WatchOpts, sink chan<- *AutonityRewarded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "Rewarded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityRewarded)
				if err := _Autonity.contract.UnpackLog(event, "Rewarded", log); err != nil {
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

// ParseRewarded is a log parse operation binding the contract event 0xb3b7a071186534c03b40695710096f289fd4ed6c1a374aff0bb648955e4fe563.
//
// Solidity: event Rewarded(address indexed addr, uint256 amount)
func (_Autonity *AutonityFilterer) ParseRewarded(log types.Log) (*AutonityRewarded, error) {
	event := new(AutonityRewarded)
	if err := _Autonity.contract.UnpackLog(event, "Rewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutonityTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Autonity contract.
type AutonityTransferIterator struct {
	Event *AutonityTransfer // Event containing the contract specifics and raw log

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
func (it *AutonityTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutonityTransfer)
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
		it.Event = new(AutonityTransfer)
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
func (it *AutonityTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutonityTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutonityTransfer represents a Transfer event raised by the Autonity contract.
type AutonityTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Autonity *AutonityFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutonityTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Autonity.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutonityTransferIterator{contract: _Autonity.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Autonity *AutonityFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AutonityTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Autonity.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutonityTransfer)
				if err := _Autonity.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Autonity *AutonityFilterer) ParseTransfer(log types.Log) (*AutonityTransfer, error) {
	event := new(AutonityTransfer)
	if err := _Autonity.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BytesLibMetaData contains all meta data concerning the BytesLib contract.
var BytesLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220a382116c91b2d99ddbd669d16236b0462d4240ef2dacde934e5c8103df52ee0a64736f6c63430008150033",
}

// BytesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesLibMetaData.ABI instead.
var BytesLibABI = BytesLibMetaData.ABI

// BytesLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesLibMetaData.Bin instead.
var BytesLibBin = BytesLibMetaData.Bin

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// HelpersMetaData contains all meta data concerning the Helpers contract.
var HelpersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122056f83c93698d67e38eaec1853be4c450ebd5bab90617892b34666041bdfa261a64736f6c63430008150033",
}

// HelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use HelpersMetaData.ABI instead.
var HelpersABI = HelpersMetaData.ABI

// HelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelpersMetaData.Bin instead.
var HelpersBin = HelpersMetaData.Bin

// DeployHelpers deploys a new Ethereum contract, binding an instance of Helpers to it.
func DeployHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helpers, error) {
	parsed, err := HelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// Helpers is an auto generated Go binding around an Ethereum contract.
type Helpers struct {
	HelpersCaller     // Read-only binding to the contract
	HelpersTransactor // Write-only binding to the contract
	HelpersFilterer   // Log filterer for contract events
}

// HelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelpersSession struct {
	Contract     *Helpers          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelpersCallerSession struct {
	Contract *HelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelpersTransactorSession struct {
	Contract     *HelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelpersRaw struct {
	Contract *Helpers // Generic contract binding to access the raw methods on
}

// HelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelpersCallerRaw struct {
	Contract *HelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelpersTransactorRaw struct {
	Contract *HelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelpers creates a new instance of Helpers, bound to a specific deployed contract.
func NewHelpers(address common.Address, backend bind.ContractBackend) (*Helpers, error) {
	contract, err := bindHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// NewHelpersCaller creates a new read-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersCaller(address common.Address, caller bind.ContractCaller) (*HelpersCaller, error) {
	contract, err := bindHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersCaller{contract: contract}, nil
}

// NewHelpersTransactor creates a new write-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HelpersTransactor, error) {
	contract, err := bindHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersTransactor{contract: contract}, nil
}

// NewHelpersFilterer creates a new log filterer instance of Helpers, bound to a specific deployed contract.
func NewHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HelpersFilterer, error) {
	contract, err := bindHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelpersFilterer{contract: contract}, nil
}

// bindHelpers binds a generic wrapper to an already deployed contract.
func bindHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.HelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transact(opts, method, params...)
}

// IACUMetaData contains all meta data concerning the IACU contract.
var IACUMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b3ab15fb": "setOperator(address)",
		"7adbf973": "setOracle(address)",
		"a2e62045": "update()",
	},
}

// IACUABI is the input ABI used to generate the binding from.
// Deprecated: Use IACUMetaData.ABI instead.
var IACUABI = IACUMetaData.ABI

// Deprecated: Use IACUMetaData.Sigs instead.
// IACUFuncSigs maps the 4-byte function signature to its string representation.
var IACUFuncSigs = IACUMetaData.Sigs

// IACU is an auto generated Go binding around an Ethereum contract.
type IACU struct {
	IACUCaller     // Read-only binding to the contract
	IACUTransactor // Write-only binding to the contract
	IACUFilterer   // Log filterer for contract events
}

// IACUCaller is an auto generated read-only Go binding around an Ethereum contract.
type IACUCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACUTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IACUTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACUFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IACUFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACUSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IACUSession struct {
	Contract     *IACU             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IACUCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IACUCallerSession struct {
	Contract *IACUCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IACUTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IACUTransactorSession struct {
	Contract     *IACUTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IACURaw is an auto generated low-level Go binding around an Ethereum contract.
type IACURaw struct {
	Contract *IACU // Generic contract binding to access the raw methods on
}

// IACUCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IACUCallerRaw struct {
	Contract *IACUCaller // Generic read-only contract binding to access the raw methods on
}

// IACUTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IACUTransactorRaw struct {
	Contract *IACUTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIACU creates a new instance of IACU, bound to a specific deployed contract.
func NewIACU(address common.Address, backend bind.ContractBackend) (*IACU, error) {
	contract, err := bindIACU(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IACU{IACUCaller: IACUCaller{contract: contract}, IACUTransactor: IACUTransactor{contract: contract}, IACUFilterer: IACUFilterer{contract: contract}}, nil
}

// NewIACUCaller creates a new read-only instance of IACU, bound to a specific deployed contract.
func NewIACUCaller(address common.Address, caller bind.ContractCaller) (*IACUCaller, error) {
	contract, err := bindIACU(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IACUCaller{contract: contract}, nil
}

// NewIACUTransactor creates a new write-only instance of IACU, bound to a specific deployed contract.
func NewIACUTransactor(address common.Address, transactor bind.ContractTransactor) (*IACUTransactor, error) {
	contract, err := bindIACU(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IACUTransactor{contract: contract}, nil
}

// NewIACUFilterer creates a new log filterer instance of IACU, bound to a specific deployed contract.
func NewIACUFilterer(address common.Address, filterer bind.ContractFilterer) (*IACUFilterer, error) {
	contract, err := bindIACU(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IACUFilterer{contract: contract}, nil
}

// bindIACU binds a generic wrapper to an already deployed contract.
func bindIACU(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IACUABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACU *IACURaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACU.Contract.IACUCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACU *IACURaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACU.Contract.IACUTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACU *IACURaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACU.Contract.IACUTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACU *IACUCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACU.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACU *IACUTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACU.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACU *IACUTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACU.Contract.contract.Transact(opts, method, params...)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IACU *IACUTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IACU.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IACU *IACUSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _IACU.Contract.SetOperator(&_IACU.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IACU *IACUTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _IACU.Contract.SetOperator(&_IACU.TransactOpts, operator)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IACU *IACUTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _IACU.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IACU *IACUSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _IACU.Contract.SetOracle(&_IACU.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IACU *IACUTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _IACU.Contract.SetOracle(&_IACU.TransactOpts, oracle)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_IACU *IACUTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACU.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_IACU *IACUSession) Update() (*types.Transaction, error) {
	return _IACU.Contract.Update(&_IACU.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(bool status)
func (_IACU *IACUTransactorSession) Update() (*types.Transaction, error) {
	return _IACU.Contract.Update(&_IACU.TransactOpts)
}

// IAccountabilityMetaData contains all meta data concerning the IAccountability contract.
var IAccountabilityMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"InnocenceProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_severity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewAccusation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_offender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_severity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"NewFaultProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseBlock\",\"type\":\"uint256\"}],\"name\":\"SlashingEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_epochEnd\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1de9d9b6": "distributeRewards(address)",
		"6c9789b0": "finalize(bool)",
		"6b5f444c": "setEpochPeriod(uint256)",
	},
}

// IAccountabilityABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountabilityMetaData.ABI instead.
var IAccountabilityABI = IAccountabilityMetaData.ABI

// Deprecated: Use IAccountabilityMetaData.Sigs instead.
// IAccountabilityFuncSigs maps the 4-byte function signature to its string representation.
var IAccountabilityFuncSigs = IAccountabilityMetaData.Sigs

// IAccountability is an auto generated Go binding around an Ethereum contract.
type IAccountability struct {
	IAccountabilityCaller     // Read-only binding to the contract
	IAccountabilityTransactor // Write-only binding to the contract
	IAccountabilityFilterer   // Log filterer for contract events
}

// IAccountabilityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountabilityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountabilityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountabilitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountabilitySession struct {
	Contract     *IAccountability  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountabilityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountabilityCallerSession struct {
	Contract *IAccountabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountabilityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountabilityTransactorSession struct {
	Contract     *IAccountabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountabilityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountabilityRaw struct {
	Contract *IAccountability // Generic contract binding to access the raw methods on
}

// IAccountabilityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountabilityCallerRaw struct {
	Contract *IAccountabilityCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountabilityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountabilityTransactorRaw struct {
	Contract *IAccountabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountability creates a new instance of IAccountability, bound to a specific deployed contract.
func NewIAccountability(address common.Address, backend bind.ContractBackend) (*IAccountability, error) {
	contract, err := bindIAccountability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountability{IAccountabilityCaller: IAccountabilityCaller{contract: contract}, IAccountabilityTransactor: IAccountabilityTransactor{contract: contract}, IAccountabilityFilterer: IAccountabilityFilterer{contract: contract}}, nil
}

// NewIAccountabilityCaller creates a new read-only instance of IAccountability, bound to a specific deployed contract.
func NewIAccountabilityCaller(address common.Address, caller bind.ContractCaller) (*IAccountabilityCaller, error) {
	contract, err := bindIAccountability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityCaller{contract: contract}, nil
}

// NewIAccountabilityTransactor creates a new write-only instance of IAccountability, bound to a specific deployed contract.
func NewIAccountabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountabilityTransactor, error) {
	contract, err := bindIAccountability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityTransactor{contract: contract}, nil
}

// NewIAccountabilityFilterer creates a new log filterer instance of IAccountability, bound to a specific deployed contract.
func NewIAccountabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountabilityFilterer, error) {
	contract, err := bindIAccountability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityFilterer{contract: contract}, nil
}

// bindIAccountability binds a generic wrapper to an already deployed contract.
func bindIAccountability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccountabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountability *IAccountabilityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountability.Contract.IAccountabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountability *IAccountabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountability.Contract.IAccountabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountability *IAccountabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountability.Contract.IAccountabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountability *IAccountabilityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountability *IAccountabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountability *IAccountabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountability.Contract.contract.Transact(opts, method, params...)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_IAccountability *IAccountabilityTransactor) DistributeRewards(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _IAccountability.contract.Transact(opts, "distributeRewards", _validator)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_IAccountability *IAccountabilitySession) DistributeRewards(_validator common.Address) (*types.Transaction, error) {
	return _IAccountability.Contract.DistributeRewards(&_IAccountability.TransactOpts, _validator)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address _validator) payable returns()
func (_IAccountability *IAccountabilityTransactorSession) DistributeRewards(_validator common.Address) (*types.Transaction, error) {
	return _IAccountability.Contract.DistributeRewards(&_IAccountability.TransactOpts, _validator)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_IAccountability *IAccountabilityTransactor) Finalize(opts *bind.TransactOpts, _epochEnd bool) (*types.Transaction, error) {
	return _IAccountability.contract.Transact(opts, "finalize", _epochEnd)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_IAccountability *IAccountabilitySession) Finalize(_epochEnd bool) (*types.Transaction, error) {
	return _IAccountability.Contract.Finalize(&_IAccountability.TransactOpts, _epochEnd)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool _epochEnd) returns()
func (_IAccountability *IAccountabilityTransactorSession) Finalize(_epochEnd bool) (*types.Transaction, error) {
	return _IAccountability.Contract.Finalize(&_IAccountability.TransactOpts, _epochEnd)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_IAccountability *IAccountabilityTransactor) SetEpochPeriod(opts *bind.TransactOpts, _newPeriod *big.Int) (*types.Transaction, error) {
	return _IAccountability.contract.Transact(opts, "setEpochPeriod", _newPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_IAccountability *IAccountabilitySession) SetEpochPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _IAccountability.Contract.SetEpochPeriod(&_IAccountability.TransactOpts, _newPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _newPeriod) returns()
func (_IAccountability *IAccountabilityTransactorSession) SetEpochPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _IAccountability.Contract.SetEpochPeriod(&_IAccountability.TransactOpts, _newPeriod)
}

// IAccountabilityInnocenceProvenIterator is returned from FilterInnocenceProven and is used to iterate over the raw logs and unpacked data for InnocenceProven events raised by the IAccountability contract.
type IAccountabilityInnocenceProvenIterator struct {
	Event *IAccountabilityInnocenceProven // Event containing the contract specifics and raw log

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
func (it *IAccountabilityInnocenceProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountabilityInnocenceProven)
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
		it.Event = new(IAccountabilityInnocenceProven)
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
func (it *IAccountabilityInnocenceProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountabilityInnocenceProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountabilityInnocenceProven represents a InnocenceProven event raised by the IAccountability contract.
type IAccountabilityInnocenceProven struct {
	Offender common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInnocenceProven is a free log retrieval operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) FilterInnocenceProven(opts *bind.FilterOpts, _offender []common.Address) (*IAccountabilityInnocenceProvenIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.FilterLogs(opts, "InnocenceProven", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityInnocenceProvenIterator{contract: _IAccountability.contract, event: "InnocenceProven", logs: logs, sub: sub}, nil
}

// WatchInnocenceProven is a free log subscription operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) WatchInnocenceProven(opts *bind.WatchOpts, sink chan<- *IAccountabilityInnocenceProven, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.WatchLogs(opts, "InnocenceProven", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountabilityInnocenceProven)
				if err := _IAccountability.contract.UnpackLog(event, "InnocenceProven", log); err != nil {
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

// ParseInnocenceProven is a log parse operation binding the contract event 0x1fa96beb8dddcb7d4484dd00c4059e872439f7a474a2ecf49c430fc6e86c9e1f.
//
// Solidity: event InnocenceProven(address indexed _offender, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) ParseInnocenceProven(log types.Log) (*IAccountabilityInnocenceProven, error) {
	event := new(IAccountabilityInnocenceProven)
	if err := _IAccountability.contract.UnpackLog(event, "InnocenceProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountabilityNewAccusationIterator is returned from FilterNewAccusation and is used to iterate over the raw logs and unpacked data for NewAccusation events raised by the IAccountability contract.
type IAccountabilityNewAccusationIterator struct {
	Event *IAccountabilityNewAccusation // Event containing the contract specifics and raw log

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
func (it *IAccountabilityNewAccusationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountabilityNewAccusation)
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
		it.Event = new(IAccountabilityNewAccusation)
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
func (it *IAccountabilityNewAccusationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountabilityNewAccusationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountabilityNewAccusation represents a NewAccusation event raised by the IAccountability contract.
type IAccountabilityNewAccusation struct {
	Offender common.Address
	Severity *big.Int
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAccusation is a free log retrieval operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) FilterNewAccusation(opts *bind.FilterOpts, _offender []common.Address) (*IAccountabilityNewAccusationIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.FilterLogs(opts, "NewAccusation", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityNewAccusationIterator{contract: _IAccountability.contract, event: "NewAccusation", logs: logs, sub: sub}, nil
}

// WatchNewAccusation is a free log subscription operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) WatchNewAccusation(opts *bind.WatchOpts, sink chan<- *IAccountabilityNewAccusation, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.WatchLogs(opts, "NewAccusation", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountabilityNewAccusation)
				if err := _IAccountability.contract.UnpackLog(event, "NewAccusation", log); err != nil {
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

// ParseNewAccusation is a log parse operation binding the contract event 0x2e8e354b41470731dafa7c3df150e9498a8d5b9c51ff0259fbf77f721ba40351.
//
// Solidity: event NewAccusation(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) ParseNewAccusation(log types.Log) (*IAccountabilityNewAccusation, error) {
	event := new(IAccountabilityNewAccusation)
	if err := _IAccountability.contract.UnpackLog(event, "NewAccusation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountabilityNewFaultProofIterator is returned from FilterNewFaultProof and is used to iterate over the raw logs and unpacked data for NewFaultProof events raised by the IAccountability contract.
type IAccountabilityNewFaultProofIterator struct {
	Event *IAccountabilityNewFaultProof // Event containing the contract specifics and raw log

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
func (it *IAccountabilityNewFaultProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountabilityNewFaultProof)
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
		it.Event = new(IAccountabilityNewFaultProof)
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
func (it *IAccountabilityNewFaultProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountabilityNewFaultProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountabilityNewFaultProof represents a NewFaultProof event raised by the IAccountability contract.
type IAccountabilityNewFaultProof struct {
	Offender common.Address
	Severity *big.Int
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFaultProof is a free log retrieval operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) FilterNewFaultProof(opts *bind.FilterOpts, _offender []common.Address) (*IAccountabilityNewFaultProofIterator, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.FilterLogs(opts, "NewFaultProof", _offenderRule)
	if err != nil {
		return nil, err
	}
	return &IAccountabilityNewFaultProofIterator{contract: _IAccountability.contract, event: "NewFaultProof", logs: logs, sub: sub}, nil
}

// WatchNewFaultProof is a free log subscription operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) WatchNewFaultProof(opts *bind.WatchOpts, sink chan<- *IAccountabilityNewFaultProof, _offender []common.Address) (event.Subscription, error) {

	var _offenderRule []interface{}
	for _, _offenderItem := range _offender {
		_offenderRule = append(_offenderRule, _offenderItem)
	}

	logs, sub, err := _IAccountability.contract.WatchLogs(opts, "NewFaultProof", _offenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountabilityNewFaultProof)
				if err := _IAccountability.contract.UnpackLog(event, "NewFaultProof", log); err != nil {
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

// ParseNewFaultProof is a log parse operation binding the contract event 0x6b7783718ab8e152c193eb08bf76eed1191fcd1677a23a7fe9d338265aad132f.
//
// Solidity: event NewFaultProof(address indexed _offender, uint256 _severity, uint256 _id)
func (_IAccountability *IAccountabilityFilterer) ParseNewFaultProof(log types.Log) (*IAccountabilityNewFaultProof, error) {
	event := new(IAccountabilityNewFaultProof)
	if err := _IAccountability.contract.UnpackLog(event, "NewFaultProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountabilitySlashingEventIterator is returned from FilterSlashingEvent and is used to iterate over the raw logs and unpacked data for SlashingEvent events raised by the IAccountability contract.
type IAccountabilitySlashingEventIterator struct {
	Event *IAccountabilitySlashingEvent // Event containing the contract specifics and raw log

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
func (it *IAccountabilitySlashingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountabilitySlashingEvent)
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
		it.Event = new(IAccountabilitySlashingEvent)
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
func (it *IAccountabilitySlashingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountabilitySlashingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountabilitySlashingEvent represents a SlashingEvent event raised by the IAccountability contract.
type IAccountabilitySlashingEvent struct {
	Validator    common.Address
	Amount       *big.Int
	ReleaseBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSlashingEvent is a free log retrieval operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_IAccountability *IAccountabilityFilterer) FilterSlashingEvent(opts *bind.FilterOpts) (*IAccountabilitySlashingEventIterator, error) {

	logs, sub, err := _IAccountability.contract.FilterLogs(opts, "SlashingEvent")
	if err != nil {
		return nil, err
	}
	return &IAccountabilitySlashingEventIterator{contract: _IAccountability.contract, event: "SlashingEvent", logs: logs, sub: sub}, nil
}

// WatchSlashingEvent is a free log subscription operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_IAccountability *IAccountabilityFilterer) WatchSlashingEvent(opts *bind.WatchOpts, sink chan<- *IAccountabilitySlashingEvent) (event.Subscription, error) {

	logs, sub, err := _IAccountability.contract.WatchLogs(opts, "SlashingEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountabilitySlashingEvent)
				if err := _IAccountability.contract.UnpackLog(event, "SlashingEvent", log); err != nil {
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

// ParseSlashingEvent is a log parse operation binding the contract event 0x416d384429edd812a063948dc473d928a51a43b9763a0d5bde902d3df5dcddcc.
//
// Solidity: event SlashingEvent(address validator, uint256 amount, uint256 releaseBlock)
func (_IAccountability *IAccountabilityFilterer) ParseSlashingEvent(log types.Log) (*IAccountabilitySlashingEvent, error) {
	event := new(IAccountabilitySlashingEvent)
	if err := _IAccountability.contract.UnpackLog(event, "SlashingEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAutonityMetaData contains all meta data concerning the IAutonity contract.
var IAutonityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e7f43c68": "getOperator()",
		"833b1fce": "getOracle()",
	},
}

// IAutonityABI is the input ABI used to generate the binding from.
// Deprecated: Use IAutonityMetaData.ABI instead.
var IAutonityABI = IAutonityMetaData.ABI

// Deprecated: Use IAutonityMetaData.Sigs instead.
// IAutonityFuncSigs maps the 4-byte function signature to its string representation.
var IAutonityFuncSigs = IAutonityMetaData.Sigs

// IAutonity is an auto generated Go binding around an Ethereum contract.
type IAutonity struct {
	IAutonityCaller     // Read-only binding to the contract
	IAutonityTransactor // Write-only binding to the contract
	IAutonityFilterer   // Log filterer for contract events
}

// IAutonityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAutonityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAutonityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAutonityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAutonityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAutonityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAutonitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAutonitySession struct {
	Contract     *IAutonity        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAutonityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAutonityCallerSession struct {
	Contract *IAutonityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IAutonityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAutonityTransactorSession struct {
	Contract     *IAutonityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IAutonityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAutonityRaw struct {
	Contract *IAutonity // Generic contract binding to access the raw methods on
}

// IAutonityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAutonityCallerRaw struct {
	Contract *IAutonityCaller // Generic read-only contract binding to access the raw methods on
}

// IAutonityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAutonityTransactorRaw struct {
	Contract *IAutonityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAutonity creates a new instance of IAutonity, bound to a specific deployed contract.
func NewIAutonity(address common.Address, backend bind.ContractBackend) (*IAutonity, error) {
	contract, err := bindIAutonity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAutonity{IAutonityCaller: IAutonityCaller{contract: contract}, IAutonityTransactor: IAutonityTransactor{contract: contract}, IAutonityFilterer: IAutonityFilterer{contract: contract}}, nil
}

// NewIAutonityCaller creates a new read-only instance of IAutonity, bound to a specific deployed contract.
func NewIAutonityCaller(address common.Address, caller bind.ContractCaller) (*IAutonityCaller, error) {
	contract, err := bindIAutonity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAutonityCaller{contract: contract}, nil
}

// NewIAutonityTransactor creates a new write-only instance of IAutonity, bound to a specific deployed contract.
func NewIAutonityTransactor(address common.Address, transactor bind.ContractTransactor) (*IAutonityTransactor, error) {
	contract, err := bindIAutonity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAutonityTransactor{contract: contract}, nil
}

// NewIAutonityFilterer creates a new log filterer instance of IAutonity, bound to a specific deployed contract.
func NewIAutonityFilterer(address common.Address, filterer bind.ContractFilterer) (*IAutonityFilterer, error) {
	contract, err := bindIAutonity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAutonityFilterer{contract: contract}, nil
}

// bindIAutonity binds a generic wrapper to an already deployed contract.
func bindIAutonity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAutonityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAutonity *IAutonityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAutonity.Contract.IAutonityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAutonity *IAutonityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAutonity.Contract.IAutonityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAutonity *IAutonityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAutonity.Contract.IAutonityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAutonity *IAutonityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAutonity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAutonity *IAutonityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAutonity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAutonity *IAutonityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAutonity.Contract.contract.Transact(opts, method, params...)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IAutonity *IAutonityCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAutonity.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IAutonity *IAutonitySession) GetOperator() (common.Address, error) {
	return _IAutonity.Contract.GetOperator(&_IAutonity.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IAutonity *IAutonityCallerSession) GetOperator() (common.Address, error) {
	return _IAutonity.Contract.GetOperator(&_IAutonity.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IAutonity *IAutonityCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAutonity.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IAutonity *IAutonitySession) GetOracle() (common.Address, error) {
	return _IAutonity.Contract.GetOracle(&_IAutonity.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_IAutonity *IAutonityCallerSession) GetOracle() (common.Address, error) {
	return _IAutonity.Contract.GetOracle(&_IAutonity.CallOpts)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOracleMetaData contains all meta data concerning the IOracle contract.
var IOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_votePeriod\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"NewSymbols\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"_votes\",\"type\":\"int256[]\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getRoundData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structIOracle.RoundData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbols\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"latestRoundData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structIOracle.RoundData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"}],\"name\":\"setSymbols\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newVoters\",\"type\":\"address[]\"}],\"name\":\"setVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commit\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"_reports\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4bb278f3": "finalize()",
		"9670c0bc": "getPrecision()",
		"9f8743f7": "getRound()",
		"3c8510fd": "getRoundData(uint256,string)",
		"df7f710e": "getSymbols()",
		"b78dec52": "getVotePeriod()",
		"cdd72253": "getVoters()",
		"33f98c77": "latestRoundData(string)",
		"b3ab15fb": "setOperator(address)",
		"8d4f75d2": "setSymbols(string[])",
		"845023f2": "setVoters(address[])",
		"307de9b6": "vote(uint256,int256[],uint256)",
	},
}

// IOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IOracleMetaData.ABI instead.
var IOracleABI = IOracleMetaData.ABI

// Deprecated: Use IOracleMetaData.Sigs instead.
// IOracleFuncSigs maps the 4-byte function signature to its string representation.
var IOracleFuncSigs = IOracleMetaData.Sigs

// IOracle is an auto generated Go binding around an Ethereum contract.
type IOracle struct {
	IOracleCaller     // Read-only binding to the contract
	IOracleTransactor // Write-only binding to the contract
	IOracleFilterer   // Log filterer for contract events
}

// IOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOracleSession struct {
	Contract     *IOracle          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOracleCallerSession struct {
	Contract *IOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOracleTransactorSession struct {
	Contract     *IOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOracleRaw struct {
	Contract *IOracle // Generic contract binding to access the raw methods on
}

// IOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOracleCallerRaw struct {
	Contract *IOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOracleTransactorRaw struct {
	Contract *IOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOracle creates a new instance of IOracle, bound to a specific deployed contract.
func NewIOracle(address common.Address, backend bind.ContractBackend) (*IOracle, error) {
	contract, err := bindIOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOracle{IOracleCaller: IOracleCaller{contract: contract}, IOracleTransactor: IOracleTransactor{contract: contract}, IOracleFilterer: IOracleFilterer{contract: contract}}, nil
}

// NewIOracleCaller creates a new read-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleCaller(address common.Address, caller bind.ContractCaller) (*IOracleCaller, error) {
	contract, err := bindIOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleCaller{contract: contract}, nil
}

// NewIOracleTransactor creates a new write-only instance of IOracle, bound to a specific deployed contract.
func NewIOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IOracleTransactor, error) {
	contract, err := bindIOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOracleTransactor{contract: contract}, nil
}

// NewIOracleFilterer creates a new log filterer instance of IOracle, bound to a specific deployed contract.
func NewIOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IOracleFilterer, error) {
	contract, err := bindIOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOracleFilterer{contract: contract}, nil
}

// bindIOracle binds a generic wrapper to an already deployed contract.
func bindIOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.IOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.IOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOracle *IOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOracle *IOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOracle *IOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOracle.Contract.contract.Transact(opts, method, params...)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_IOracle *IOracleCaller) GetPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_IOracle *IOracleSession) GetPrecision() (*big.Int, error) {
	return _IOracle.Contract.GetPrecision(&_IOracle.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_IOracle *IOracleCallerSession) GetPrecision() (*big.Int, error) {
	return _IOracle.Contract.GetPrecision(&_IOracle.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOracle *IOracleCaller) GetRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOracle *IOracleSession) GetRound() (*big.Int, error) {
	return _IOracle.Contract.GetRound(&_IOracle.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOracle *IOracleCallerSession) GetRound() (*big.Int, error) {
	return _IOracle.Contract.GetRound(&_IOracle.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleCaller) GetRoundData(opts *bind.CallOpts, _round *big.Int, _symbol string) (IOracleRoundData, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getRoundData", _round, _symbol)

	if err != nil {
		return *new(IOracleRoundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleRoundData)).(*IOracleRoundData)

	return out0, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleSession) GetRoundData(_round *big.Int, _symbol string) (IOracleRoundData, error) {
	return _IOracle.Contract.GetRoundData(&_IOracle.CallOpts, _round, _symbol)
}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleCallerSession) GetRoundData(_round *big.Int, _symbol string) (IOracleRoundData, error) {
	return _IOracle.Contract.GetRoundData(&_IOracle.CallOpts, _round, _symbol)
}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[] _symbols)
func (_IOracle *IOracleCaller) GetSymbols(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getSymbols")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[] _symbols)
func (_IOracle *IOracleSession) GetSymbols() ([]string, error) {
	return _IOracle.Contract.GetSymbols(&_IOracle.CallOpts)
}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[] _symbols)
func (_IOracle *IOracleCallerSession) GetSymbols() ([]string, error) {
	return _IOracle.Contract.GetSymbols(&_IOracle.CallOpts)
}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_IOracle *IOracleCaller) GetVotePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getVotePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_IOracle *IOracleSession) GetVotePeriod() (*big.Int, error) {
	return _IOracle.Contract.GetVotePeriod(&_IOracle.CallOpts)
}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_IOracle *IOracleCallerSession) GetVotePeriod() (*big.Int, error) {
	return _IOracle.Contract.GetVotePeriod(&_IOracle.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_IOracle *IOracleCaller) GetVoters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "getVoters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_IOracle *IOracleSession) GetVoters() ([]common.Address, error) {
	return _IOracle.Contract.GetVoters(&_IOracle.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_IOracle *IOracleCallerSession) GetVoters() ([]common.Address, error) {
	return _IOracle.Contract.GetVoters(&_IOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleCaller) LatestRoundData(opts *bind.CallOpts, _symbol string) (IOracleRoundData, error) {
	var out []interface{}
	err := _IOracle.contract.Call(opts, &out, "latestRoundData", _symbol)

	if err != nil {
		return *new(IOracleRoundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleRoundData)).(*IOracleRoundData)

	return out0, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleSession) LatestRoundData(_symbol string) (IOracleRoundData, error) {
	return _IOracle.Contract.LatestRoundData(&_IOracle.CallOpts, _symbol)
}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_IOracle *IOracleCallerSession) LatestRoundData(_symbol string) (IOracleRoundData, error) {
	return _IOracle.Contract.LatestRoundData(&_IOracle.CallOpts, _symbol)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_IOracle *IOracleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_IOracle *IOracleSession) Finalize() (*types.Transaction, error) {
	return _IOracle.Contract.Finalize(&_IOracle.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_IOracle *IOracleTransactorSession) Finalize() (*types.Transaction, error) {
	return _IOracle.Contract.Finalize(&_IOracle.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IOracle *IOracleTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IOracle *IOracleSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _IOracle.Contract.SetOperator(&_IOracle.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IOracle *IOracleTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _IOracle.Contract.SetOperator(&_IOracle.TransactOpts, _operator)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_IOracle *IOracleTransactor) SetSymbols(opts *bind.TransactOpts, _symbols []string) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "setSymbols", _symbols)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_IOracle *IOracleSession) SetSymbols(_symbols []string) (*types.Transaction, error) {
	return _IOracle.Contract.SetSymbols(&_IOracle.TransactOpts, _symbols)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_IOracle *IOracleTransactorSession) SetSymbols(_symbols []string) (*types.Transaction, error) {
	return _IOracle.Contract.SetSymbols(&_IOracle.TransactOpts, _symbols)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_IOracle *IOracleTransactor) SetVoters(opts *bind.TransactOpts, _newVoters []common.Address) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "setVoters", _newVoters)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_IOracle *IOracleSession) SetVoters(_newVoters []common.Address) (*types.Transaction, error) {
	return _IOracle.Contract.SetVoters(&_IOracle.TransactOpts, _newVoters)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_IOracle *IOracleTransactorSession) SetVoters(_newVoters []common.Address) (*types.Transaction, error) {
	return _IOracle.Contract.SetVoters(&_IOracle.TransactOpts, _newVoters)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_IOracle *IOracleTransactor) Vote(opts *bind.TransactOpts, _commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _IOracle.contract.Transact(opts, "vote", _commit, _reports, _salt)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_IOracle *IOracleSession) Vote(_commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _IOracle.Contract.Vote(&_IOracle.TransactOpts, _commit, _reports, _salt)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_IOracle *IOracleTransactorSession) Vote(_commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _IOracle.Contract.Vote(&_IOracle.TransactOpts, _commit, _reports, _salt)
}

// IOracleNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the IOracle contract.
type IOracleNewRoundIterator struct {
	Event *IOracleNewRound // Event containing the contract specifics and raw log

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
func (it *IOracleNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOracleNewRound)
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
		it.Event = new(IOracleNewRound)
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
func (it *IOracleNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOracleNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOracleNewRound represents a NewRound event raised by the IOracle contract.
type IOracleNewRound struct {
	Round      *big.Int
	Height     *big.Int
	Timestamp  *big.Int
	VotePeriod *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_IOracle *IOracleFilterer) FilterNewRound(opts *bind.FilterOpts) (*IOracleNewRoundIterator, error) {

	logs, sub, err := _IOracle.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &IOracleNewRoundIterator{contract: _IOracle.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_IOracle *IOracleFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *IOracleNewRound) (event.Subscription, error) {

	logs, sub, err := _IOracle.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOracleNewRound)
				if err := _IOracle.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_IOracle *IOracleFilterer) ParseNewRound(log types.Log) (*IOracleNewRound, error) {
	event := new(IOracleNewRound)
	if err := _IOracle.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOracleNewSymbolsIterator is returned from FilterNewSymbols and is used to iterate over the raw logs and unpacked data for NewSymbols events raised by the IOracle contract.
type IOracleNewSymbolsIterator struct {
	Event *IOracleNewSymbols // Event containing the contract specifics and raw log

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
func (it *IOracleNewSymbolsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOracleNewSymbols)
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
		it.Event = new(IOracleNewSymbols)
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
func (it *IOracleNewSymbolsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOracleNewSymbolsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOracleNewSymbols represents a NewSymbols event raised by the IOracle contract.
type IOracleNewSymbols struct {
	Symbols []string
	Round   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSymbols is a free log retrieval operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_IOracle *IOracleFilterer) FilterNewSymbols(opts *bind.FilterOpts) (*IOracleNewSymbolsIterator, error) {

	logs, sub, err := _IOracle.contract.FilterLogs(opts, "NewSymbols")
	if err != nil {
		return nil, err
	}
	return &IOracleNewSymbolsIterator{contract: _IOracle.contract, event: "NewSymbols", logs: logs, sub: sub}, nil
}

// WatchNewSymbols is a free log subscription operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_IOracle *IOracleFilterer) WatchNewSymbols(opts *bind.WatchOpts, sink chan<- *IOracleNewSymbols) (event.Subscription, error) {

	logs, sub, err := _IOracle.contract.WatchLogs(opts, "NewSymbols")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOracleNewSymbols)
				if err := _IOracle.contract.UnpackLog(event, "NewSymbols", log); err != nil {
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

// ParseNewSymbols is a log parse operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_IOracle *IOracleFilterer) ParseNewSymbols(log types.Log) (*IOracleNewSymbols, error) {
	event := new(IOracleNewSymbols)
	if err := _IOracle.contract.UnpackLog(event, "NewSymbols", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOracleVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the IOracle contract.
type IOracleVotedIterator struct {
	Event *IOracleVoted // Event containing the contract specifics and raw log

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
func (it *IOracleVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOracleVoted)
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
		it.Event = new(IOracleVoted)
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
func (it *IOracleVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOracleVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOracleVoted represents a Voted event raised by the IOracle contract.
type IOracleVoted struct {
	Voter common.Address
	Votes []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_IOracle *IOracleFilterer) FilterVoted(opts *bind.FilterOpts, _voter []common.Address) (*IOracleVotedIterator, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _IOracle.contract.FilterLogs(opts, "Voted", _voterRule)
	if err != nil {
		return nil, err
	}
	return &IOracleVotedIterator{contract: _IOracle.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_IOracle *IOracleFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *IOracleVoted, _voter []common.Address) (event.Subscription, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _IOracle.contract.WatchLogs(opts, "Voted", _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOracleVoted)
				if err := _IOracle.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_IOracle *IOracleFilterer) ParseVoted(log types.Log) (*IOracleVoted, error) {
	event := new(IOracleVoted)
	if err := _IOracle.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStabilizationMetaData contains all meta data concerning the IStabilization contract.
var IStabilizationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b3ab15fb": "setOperator(address)",
		"7adbf973": "setOracle(address)",
	},
}

// IStabilizationABI is the input ABI used to generate the binding from.
// Deprecated: Use IStabilizationMetaData.ABI instead.
var IStabilizationABI = IStabilizationMetaData.ABI

// Deprecated: Use IStabilizationMetaData.Sigs instead.
// IStabilizationFuncSigs maps the 4-byte function signature to its string representation.
var IStabilizationFuncSigs = IStabilizationMetaData.Sigs

// IStabilization is an auto generated Go binding around an Ethereum contract.
type IStabilization struct {
	IStabilizationCaller     // Read-only binding to the contract
	IStabilizationTransactor // Write-only binding to the contract
	IStabilizationFilterer   // Log filterer for contract events
}

// IStabilizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStabilizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStabilizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStabilizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStabilizationSession struct {
	Contract     *IStabilization   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStabilizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStabilizationCallerSession struct {
	Contract *IStabilizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IStabilizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStabilizationTransactorSession struct {
	Contract     *IStabilizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IStabilizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStabilizationRaw struct {
	Contract *IStabilization // Generic contract binding to access the raw methods on
}

// IStabilizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStabilizationCallerRaw struct {
	Contract *IStabilizationCaller // Generic read-only contract binding to access the raw methods on
}

// IStabilizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStabilizationTransactorRaw struct {
	Contract *IStabilizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStabilization creates a new instance of IStabilization, bound to a specific deployed contract.
func NewIStabilization(address common.Address, backend bind.ContractBackend) (*IStabilization, error) {
	contract, err := bindIStabilization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStabilization{IStabilizationCaller: IStabilizationCaller{contract: contract}, IStabilizationTransactor: IStabilizationTransactor{contract: contract}, IStabilizationFilterer: IStabilizationFilterer{contract: contract}}, nil
}

// NewIStabilizationCaller creates a new read-only instance of IStabilization, bound to a specific deployed contract.
func NewIStabilizationCaller(address common.Address, caller bind.ContractCaller) (*IStabilizationCaller, error) {
	contract, err := bindIStabilization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStabilizationCaller{contract: contract}, nil
}

// NewIStabilizationTransactor creates a new write-only instance of IStabilization, bound to a specific deployed contract.
func NewIStabilizationTransactor(address common.Address, transactor bind.ContractTransactor) (*IStabilizationTransactor, error) {
	contract, err := bindIStabilization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStabilizationTransactor{contract: contract}, nil
}

// NewIStabilizationFilterer creates a new log filterer instance of IStabilization, bound to a specific deployed contract.
func NewIStabilizationFilterer(address common.Address, filterer bind.ContractFilterer) (*IStabilizationFilterer, error) {
	contract, err := bindIStabilization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStabilizationFilterer{contract: contract}, nil
}

// bindIStabilization binds a generic wrapper to an already deployed contract.
func bindIStabilization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStabilizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStabilization *IStabilizationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStabilization.Contract.IStabilizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStabilization *IStabilizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStabilization.Contract.IStabilizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStabilization *IStabilizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStabilization.Contract.IStabilizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStabilization *IStabilizationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStabilization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStabilization *IStabilizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStabilization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStabilization *IStabilizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStabilization.Contract.contract.Transact(opts, method, params...)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IStabilization *IStabilizationTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _IStabilization.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IStabilization *IStabilizationSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _IStabilization.Contract.SetOperator(&_IStabilization.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_IStabilization *IStabilizationTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _IStabilization.Contract.SetOperator(&_IStabilization.TransactOpts, operator)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IStabilization *IStabilizationTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _IStabilization.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IStabilization *IStabilizationSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _IStabilization.Contract.SetOracle(&_IStabilization.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_IStabilization *IStabilizationTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _IStabilization.Contract.SetOracle(&_IStabilization.TransactOpts, oracle)
}

// ISupplyControlMetaData contains all meta data concerning the ISupplyControl contract.
var ISupplyControlMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"availableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stabilizer_\",\"type\":\"address\"}],\"name\":\"setStabilizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stabilizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7ecc2b56": "availableSupply()",
		"44df8e70": "burn()",
		"40c10f19": "mint(address,uint256)",
		"b3ab15fb": "setOperator(address)",
		"db7f521a": "setStabilizer(address)",
		"7e47961c": "stabilizer()",
		"18160ddd": "totalSupply()",
	},
}

// ISupplyControlABI is the input ABI used to generate the binding from.
// Deprecated: Use ISupplyControlMetaData.ABI instead.
var ISupplyControlABI = ISupplyControlMetaData.ABI

// Deprecated: Use ISupplyControlMetaData.Sigs instead.
// ISupplyControlFuncSigs maps the 4-byte function signature to its string representation.
var ISupplyControlFuncSigs = ISupplyControlMetaData.Sigs

// ISupplyControl is an auto generated Go binding around an Ethereum contract.
type ISupplyControl struct {
	ISupplyControlCaller     // Read-only binding to the contract
	ISupplyControlTransactor // Write-only binding to the contract
	ISupplyControlFilterer   // Log filterer for contract events
}

// ISupplyControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISupplyControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISupplyControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISupplyControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISupplyControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISupplyControlSession struct {
	Contract     *ISupplyControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISupplyControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISupplyControlCallerSession struct {
	Contract *ISupplyControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISupplyControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISupplyControlTransactorSession struct {
	Contract     *ISupplyControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISupplyControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISupplyControlRaw struct {
	Contract *ISupplyControl // Generic contract binding to access the raw methods on
}

// ISupplyControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISupplyControlCallerRaw struct {
	Contract *ISupplyControlCaller // Generic read-only contract binding to access the raw methods on
}

// ISupplyControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISupplyControlTransactorRaw struct {
	Contract *ISupplyControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISupplyControl creates a new instance of ISupplyControl, bound to a specific deployed contract.
func NewISupplyControl(address common.Address, backend bind.ContractBackend) (*ISupplyControl, error) {
	contract, err := bindISupplyControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISupplyControl{ISupplyControlCaller: ISupplyControlCaller{contract: contract}, ISupplyControlTransactor: ISupplyControlTransactor{contract: contract}, ISupplyControlFilterer: ISupplyControlFilterer{contract: contract}}, nil
}

// NewISupplyControlCaller creates a new read-only instance of ISupplyControl, bound to a specific deployed contract.
func NewISupplyControlCaller(address common.Address, caller bind.ContractCaller) (*ISupplyControlCaller, error) {
	contract, err := bindISupplyControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISupplyControlCaller{contract: contract}, nil
}

// NewISupplyControlTransactor creates a new write-only instance of ISupplyControl, bound to a specific deployed contract.
func NewISupplyControlTransactor(address common.Address, transactor bind.ContractTransactor) (*ISupplyControlTransactor, error) {
	contract, err := bindISupplyControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISupplyControlTransactor{contract: contract}, nil
}

// NewISupplyControlFilterer creates a new log filterer instance of ISupplyControl, bound to a specific deployed contract.
func NewISupplyControlFilterer(address common.Address, filterer bind.ContractFilterer) (*ISupplyControlFilterer, error) {
	contract, err := bindISupplyControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISupplyControlFilterer{contract: contract}, nil
}

// bindISupplyControl binds a generic wrapper to an already deployed contract.
func bindISupplyControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISupplyControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupplyControl *ISupplyControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupplyControl.Contract.ISupplyControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupplyControl *ISupplyControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupplyControl.Contract.ISupplyControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupplyControl *ISupplyControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupplyControl.Contract.ISupplyControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISupplyControl *ISupplyControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISupplyControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISupplyControl *ISupplyControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupplyControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISupplyControl *ISupplyControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISupplyControl.Contract.contract.Transact(opts, method, params...)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x7ecc2b56.
//
// Solidity: function availableSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlCaller) AvailableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISupplyControl.contract.Call(opts, &out, "availableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSupply is a free data retrieval call binding the contract method 0x7ecc2b56.
//
// Solidity: function availableSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlSession) AvailableSupply() (*big.Int, error) {
	return _ISupplyControl.Contract.AvailableSupply(&_ISupplyControl.CallOpts)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x7ecc2b56.
//
// Solidity: function availableSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlCallerSession) AvailableSupply() (*big.Int, error) {
	return _ISupplyControl.Contract.AvailableSupply(&_ISupplyControl.CallOpts)
}

// Stabilizer is a free data retrieval call binding the contract method 0x7e47961c.
//
// Solidity: function stabilizer() view returns(address)
func (_ISupplyControl *ISupplyControlCaller) Stabilizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISupplyControl.contract.Call(opts, &out, "stabilizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stabilizer is a free data retrieval call binding the contract method 0x7e47961c.
//
// Solidity: function stabilizer() view returns(address)
func (_ISupplyControl *ISupplyControlSession) Stabilizer() (common.Address, error) {
	return _ISupplyControl.Contract.Stabilizer(&_ISupplyControl.CallOpts)
}

// Stabilizer is a free data retrieval call binding the contract method 0x7e47961c.
//
// Solidity: function stabilizer() view returns(address)
func (_ISupplyControl *ISupplyControlCallerSession) Stabilizer() (common.Address, error) {
	return _ISupplyControl.Contract.Stabilizer(&_ISupplyControl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISupplyControl.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlSession) TotalSupply() (*big.Int, error) {
	return _ISupplyControl.Contract.TotalSupply(&_ISupplyControl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ISupplyControl *ISupplyControlCallerSession) TotalSupply() (*big.Int, error) {
	return _ISupplyControl.Contract.TotalSupply(&_ISupplyControl.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_ISupplyControl *ISupplyControlTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISupplyControl.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_ISupplyControl *ISupplyControlSession) Burn() (*types.Transaction, error) {
	return _ISupplyControl.Contract.Burn(&_ISupplyControl.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_ISupplyControl *ISupplyControlTransactorSession) Burn() (*types.Transaction, error) {
	return _ISupplyControl.Contract.Burn(&_ISupplyControl.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_ISupplyControl *ISupplyControlTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISupplyControl.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_ISupplyControl *ISupplyControlSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISupplyControl.Contract.Mint(&_ISupplyControl.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_ISupplyControl *ISupplyControlTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISupplyControl.Contract.Mint(&_ISupplyControl.TransactOpts, recipient, amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ISupplyControl *ISupplyControlTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ISupplyControl.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ISupplyControl *ISupplyControlSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ISupplyControl.Contract.SetOperator(&_ISupplyControl.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_ISupplyControl *ISupplyControlTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _ISupplyControl.Contract.SetOperator(&_ISupplyControl.TransactOpts, operator)
}

// SetStabilizer is a paid mutator transaction binding the contract method 0xdb7f521a.
//
// Solidity: function setStabilizer(address stabilizer_) returns()
func (_ISupplyControl *ISupplyControlTransactor) SetStabilizer(opts *bind.TransactOpts, stabilizer_ common.Address) (*types.Transaction, error) {
	return _ISupplyControl.contract.Transact(opts, "setStabilizer", stabilizer_)
}

// SetStabilizer is a paid mutator transaction binding the contract method 0xdb7f521a.
//
// Solidity: function setStabilizer(address stabilizer_) returns()
func (_ISupplyControl *ISupplyControlSession) SetStabilizer(stabilizer_ common.Address) (*types.Transaction, error) {
	return _ISupplyControl.Contract.SetStabilizer(&_ISupplyControl.TransactOpts, stabilizer_)
}

// SetStabilizer is a paid mutator transaction binding the contract method 0xdb7f521a.
//
// Solidity: function setStabilizer(address stabilizer_) returns()
func (_ISupplyControl *ISupplyControlTransactorSession) SetStabilizer(stabilizer_ common.Address) (*types.Transaction, error) {
	return _ISupplyControl.Contract.SetStabilizer(&_ISupplyControl.TransactOpts, stabilizer_)
}

// ISupplyControlBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ISupplyControl contract.
type ISupplyControlBurnIterator struct {
	Event *ISupplyControlBurn // Event containing the contract specifics and raw log

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
func (it *ISupplyControlBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISupplyControlBurn)
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
		it.Event = new(ISupplyControlBurn)
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
func (it *ISupplyControlBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISupplyControlBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISupplyControlBurn represents a Burn event raised by the ISupplyControl contract.
type ISupplyControlBurn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) FilterBurn(opts *bind.FilterOpts) (*ISupplyControlBurnIterator, error) {

	logs, sub, err := _ISupplyControl.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &ISupplyControlBurnIterator{contract: _ISupplyControl.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ISupplyControlBurn) (event.Subscription, error) {

	logs, sub, err := _ISupplyControl.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISupplyControlBurn)
				if err := _ISupplyControl.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) ParseBurn(log types.Log) (*ISupplyControlBurn, error) {
	event := new(ISupplyControlBurn)
	if err := _ISupplyControl.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISupplyControlMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ISupplyControl contract.
type ISupplyControlMintIterator struct {
	Event *ISupplyControlMint // Event containing the contract specifics and raw log

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
func (it *ISupplyControlMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISupplyControlMint)
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
		it.Event = new(ISupplyControlMint)
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
func (it *ISupplyControlMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISupplyControlMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISupplyControlMint represents a Mint event raised by the ISupplyControl contract.
type ISupplyControlMint struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address recipient, uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) FilterMint(opts *bind.FilterOpts) (*ISupplyControlMintIterator, error) {

	logs, sub, err := _ISupplyControl.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &ISupplyControlMintIterator{contract: _ISupplyControl.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address recipient, uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ISupplyControlMint) (event.Subscription, error) {

	logs, sub, err := _ISupplyControl.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISupplyControlMint)
				if err := _ISupplyControl.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address recipient, uint256 amount)
func (_ISupplyControl *ISupplyControlFilterer) ParseMint(log types.Log) (*ISupplyControlMint, error) {
	event := new(ISupplyControlMint)
	if err := _ISupplyControl.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidMetaData contains all meta data concerning the Liquid contract.
var LiquidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_index\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COMMISSION_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_FACTOR_UNIT_RECIP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"lockedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"unlockedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f2c3f2e": "COMMISSION_RATE_PRECISION()",
		"187cf4d7": "FEE_FACTOR_UNIT_RECIP()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"9dc29fac": "burn(address,uint256)",
		"372500ab": "claimRewards()",
		"5ea1d6f8": "commissionRate()",
		"313ce567": "decimals()",
		"282d3fdf": "lock(address,uint256)",
		"59355736": "lockedBalanceOf(address)",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"fb489a7b": "redistribute()",
		"19fac8fd": "setCommissionRate(uint256)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"61d027b3": "treasury()",
		"949813b8": "unclaimedRewards(address)",
		"7eee288d": "unlock(address,uint256)",
		"84955c88": "unlockedBalanceOf(address)",
		"3a5381b5": "validator()",
	},
	Bin: "0x608060405234801562000010575f80fd5b50604051620014b4380380620014b4833981016040819052620000339162000149565b61271082111562000042575f80fd5b600a80546001600160a01b038087166001600160a01b031992831617909255600b805492861692909116919091179055600c8290556040516200008a90829060200162000230565b60405160208183030381529060405260089081620000a99190620002ea565b5080604051602001620000bd919062000230565b60405160208183030381529060405260099081620000dc9190620002ea565b50505f80546001600160a01b0319163317905550620003b2915050565b6001600160a01b03811681146200010e575f80fd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200014157818101518382015260200162000127565b50505f910152565b5f805f80608085870312156200015d575f80fd5b84516200016a81620000f9565b60208601519094506200017d81620000f9565b6040860151606087015191945092506001600160401b0380821115620001a1575f80fd5b818701915087601f830112620001b5575f80fd5b815181811115620001ca57620001ca62000111565b604051601f8201601f19908116603f01168101908382118183101715620001f557620001f562000111565b816040528281528a60208487010111156200020e575f80fd5b6200022183602083016020880162000125565b979a9699509497505050505050565b644c4e544e2d60d81b81525f82516200025181600585016020870162000125565b9190910160050192915050565b600181811c908216806200027357607f821691505b6020821081036200029257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002e5575f81815260208120601f850160051c81016020861015620002c05750805b601f850160051c820191505b81811015620002e157828155600101620002cc565b5050505b505050565b81516001600160401b0381111562000306576200030662000111565b6200031e816200031784546200025e565b8462000298565b602080601f83116001811462000354575f84156200033c5750858301515b5f19600386901b1c1916600185901b178555620002e1565b5f85815260208120601f198616915b82811015620003845788860151825594840194600190910190840162000363565b5085821015620003a257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6110f480620003c05f395ff3fe608060405260043610610147575f3560e01c806359355736116100b3578063949813b81161006d578063949813b8146103ac57806395d89b41146103cb5780639dc29fac146103df578063a9059cbb146103fe578063dd62ed3e1461041d578063fb489a7b14610461575f80fd5b806359355736146102d25780635ea1d6f81461030657806361d027b31461031b57806370a082311461033a5780637eee288d1461036e57806384955c881461038d575f80fd5b8063282d3fdf11610104578063282d3fdf146102195780632f2c3f2e14610238578063313ce5671461024d578063372500ab146102685780633a5381b51461027c57806340c10f19146102b3575f80fd5b806306fdde031461014b578063095ea7b31461017557806318160ddd146101a4578063187cf4d7146101c257806319fac8fd146101d957806323b872dd146101fa575b5f80fd5b348015610156575f80fd5b5061015f610469565b60405161016c9190610e7f565b60405180910390f35b348015610180575f80fd5b5061019461018f366004610ee5565b6104f9565b604051901515815260200161016c565b3480156101af575f80fd5b506004545b60405190815260200161016c565b3480156101cd575f80fd5b506101b4633b9aca0081565b3480156101e4575f80fd5b506101f86101f3366004610f0d565b61050f565b005b348015610205575f80fd5b50610194610214366004610f24565b610546565b348015610224575f80fd5b506101f8610233366004610ee5565b610637565b348015610243575f80fd5b506101b461271081565b348015610258575f80fd5b506040516012815260200161016c565b348015610273575f80fd5b506101f8610719565b348015610287575f80fd5b50600a5461029b906001600160a01b031681565b6040516001600160a01b03909116815260200161016c565b3480156102be575f80fd5b506101f86102cd366004610ee5565b6107c3565b3480156102dd575f80fd5b506101b46102ec366004610f5d565b6001600160a01b03165f9081526002602052604090205490565b348015610311575f80fd5b506101b4600c5481565b348015610326575f80fd5b50600b5461029b906001600160a01b031681565b348015610345575f80fd5b506101b4610354366004610f5d565b6001600160a01b03165f9081526001602052604090205490565b348015610379575f80fd5b506101f8610388366004610ee5565b610828565b348015610398575f80fd5b506101b46103a7366004610f5d565b6108eb565b3480156103b7575f80fd5b506101b46103c6366004610f5d565b610918565b3480156103d6575f80fd5b5061015f610944565b3480156103ea575f80fd5b506101f86103f9366004610ee5565b610953565b348015610409575f80fd5b50610194610418366004610ee5565b6109b0565b348015610428575f80fd5b506101b4610437366004610f7d565b6001600160a01b039182165f90815260036020908152604080832093909416825291909152205490565b6101b46109fb565b60606008805461047890610fae565b80601f01602080910402602001604051908101604052809291908181526020018280546104a490610fae565b80156104ef5780601f106104c6576101008083540402835291602001916104ef565b820191905f5260205f20905b8154815290600101906020018083116104d257829003601f168201915b5050505050905090565b5f610505338484610b5c565b5060015b92915050565b5f546001600160a01b031633146105415760405162461bcd60e51b815260040161053890610fe6565b60405180910390fd5b600c55565b6001600160a01b0383165f908152600360209081526040808320338452909152812054828110156105ca5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152608401610538565b6105de85336105d98685611042565b610b5c565b6105e88584610c7f565b6105f28484610d63565b836001600160a01b0316856001600160a01b03165f8051602061109f8339815191528560405161062491815260200190565b60405180910390a3506001949350505050565b5f546001600160a01b031633146106605760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f90815260026020908152604080832054600190925290912054829161068f91611042565b10156106e95760405162461bcd60e51b8152602060048201526024808201527f63616e2774206c6f636b206d6f72652066756e6473207468616e20617661696c60448201526361626c6560e01b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611055565b90915550505050565b5f61072333610dac565b335f81815260056020526040808220829055519293509183908381818185875af1925050503d805f8114610772576040519150601f19603f3d011682016040523d82523d5f602084013e610777565b606091505b50509050806107bf5760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b6044820152606401610538565b5050565b5f546001600160a01b031633146107ec5760405162461bcd60e51b815260040161053890610fe6565b6107f68282610d63565b6040518181526001600160a01b038316905f905f8051602061109f833981519152906020015b60405180910390a35050565b5f546001600160a01b031633146108515760405162461bcd60e51b815260040161053890610fe6565b6001600160a01b0382165f908152600260205260409020548111156108c45760405162461bcd60e51b815260206004820152602360248201527f63616e277420756e6c6f636b206d6f72652066756e6473207468616e206c6f636044820152621ad95960ea1b6064820152608401610538565b6001600160a01b0382165f9081526002602052604081208054839290610710908490611042565b6001600160a01b0381165f9081526002602090815260408083205460019092528220546105099190611042565b5f61092282610e0e565b6001600160a01b0383165f908152600560205260409020546105099190611055565b60606009805461047890610fae565b5f546001600160a01b0316331461097c5760405162461bcd60e51b815260040161053890610fe6565b6109868282610c7f565b6040518181525f906001600160a01b038416905f8051602061109f8339815191529060200161081c565b5f6109bb3383610c7f565b6109c58383610d63565b6040518281526001600160a01b0384169033905f8051602061109f8339815191529060200160405180910390a350600192915050565b5f80546001600160a01b03163314610a255760405162461bcd60e51b815260040161053890610fe6565b600c5434905f9061271090610a3a9084611068565b610a44919061107f565b905081811115610a965760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642076616c696461746f722072657761726400000000000000006044820152606401610538565b610aa08183611042565b600b546040519193506001600160a01b0316906108fc9083905f818181858888f193505050503d805f8114610af0576040519150601f19603f3d011682016040523d82523d5f602084013e610af5565b606091505b50506004545f9150610b0b633b9aca0085611068565b610b15919061107f565b905080600754610b259190611055565b6007556004545f90633b9aca0090610b3d9084611068565b610b47919061107f565b9050610b538184611055565b94505050505090565b6001600160a01b038316610bbe5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610538565b6001600160a01b038216610c1f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610538565b6001600160a01b038381165f8181526003602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b610c8882610dac565b506001600160a01b0382165f90815260016020908152604080832054600290925290912054610cb79082611042565b821115610d065760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e7420756e6c6f636b65642066756e647300000000006044820152606401610538565b610d108282611042565b6001600160a01b0384165f90815260016020526040902055808203610d48576001600160a01b0383165f908152600660205260408120555b8160045f828254610d599190611042565b9091555050505050565b610d6c82610dac565b506001600160a01b0382165f9081526001602052604081208054839290610d94908490611055565b925050819055508060045f8282546107109190611055565b5f80610db783610e0e565b6001600160a01b0384165f90815260056020526040902054909150610ddd908290611055565b6001600160a01b039093165f9081526005602090815260408083208690556007546006909252909120555090919050565b6001600160a01b0381165f90815260016020526040812054808203610e3557505f92915050565b6001600160a01b0383165f90815260066020526040812054600754610e5a9190611042565b90505f633b9aca00610e6c8484611068565b610e76919061107f565b95945050505050565b5f6020808352835180828501525f5b81811015610eaa57858101830151858201604001528201610e8e565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610ee0575f80fd5b919050565b5f8060408385031215610ef6575f80fd5b610eff83610eca565b946020939093013593505050565b5f60208284031215610f1d575f80fd5b5035919050565b5f805f60608486031215610f36575f80fd5b610f3f84610eca565b9250610f4d60208501610eca565b9150604084013590509250925092565b5f60208284031215610f6d575f80fd5b610f7682610eca565b9392505050565b5f8060408385031215610f8e575f80fd5b610f9783610eca565b9150610fa560208401610eca565b90509250929050565b600181811c90821680610fc257607f821691505b602082108103610fe057634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526028908201527f43616c6c207265737472696374656420746f20746865204175746f6e6974792060408201526710dbdb9d1c9858dd60c21b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156105095761050961102e565b808201808211156105095761050961102e565b80820281158282048414176105095761050961102e565b5f8261109957634e487b7160e01b5f52601260045260245ffd5b50049056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212208b971a8921a90e53cf2766d536cfa590f79864d3328a6958e1150f86b9db10c264736f6c63430008150033",
}

// LiquidABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidMetaData.ABI instead.
var LiquidABI = LiquidMetaData.ABI

// Deprecated: Use LiquidMetaData.Sigs instead.
// LiquidFuncSigs maps the 4-byte function signature to its string representation.
var LiquidFuncSigs = LiquidMetaData.Sigs

// LiquidBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidMetaData.Bin instead.
var LiquidBin = LiquidMetaData.Bin

// DeployLiquid deploys a new Ethereum contract, binding an instance of Liquid to it.
func DeployLiquid(auth *bind.TransactOpts, backend bind.ContractBackend, _validator common.Address, _treasury common.Address, _commissionRate *big.Int, _index string) (common.Address, *types.Transaction, *Liquid, error) {
	parsed, err := LiquidMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidBin), backend, _validator, _treasury, _commissionRate, _index)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Liquid{LiquidCaller: LiquidCaller{contract: contract}, LiquidTransactor: LiquidTransactor{contract: contract}, LiquidFilterer: LiquidFilterer{contract: contract}}, nil
}

// Liquid is an auto generated Go binding around an Ethereum contract.
type Liquid struct {
	LiquidCaller     // Read-only binding to the contract
	LiquidTransactor // Write-only binding to the contract
	LiquidFilterer   // Log filterer for contract events
}

// LiquidCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidSession struct {
	Contract     *Liquid           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidCallerSession struct {
	Contract *LiquidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LiquidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidTransactorSession struct {
	Contract     *LiquidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidRaw struct {
	Contract *Liquid // Generic contract binding to access the raw methods on
}

// LiquidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidCallerRaw struct {
	Contract *LiquidCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidTransactorRaw struct {
	Contract *LiquidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquid creates a new instance of Liquid, bound to a specific deployed contract.
func NewLiquid(address common.Address, backend bind.ContractBackend) (*Liquid, error) {
	contract, err := bindLiquid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquid{LiquidCaller: LiquidCaller{contract: contract}, LiquidTransactor: LiquidTransactor{contract: contract}, LiquidFilterer: LiquidFilterer{contract: contract}}, nil
}

// NewLiquidCaller creates a new read-only instance of Liquid, bound to a specific deployed contract.
func NewLiquidCaller(address common.Address, caller bind.ContractCaller) (*LiquidCaller, error) {
	contract, err := bindLiquid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidCaller{contract: contract}, nil
}

// NewLiquidTransactor creates a new write-only instance of Liquid, bound to a specific deployed contract.
func NewLiquidTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidTransactor, error) {
	contract, err := bindLiquid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidTransactor{contract: contract}, nil
}

// NewLiquidFilterer creates a new log filterer instance of Liquid, bound to a specific deployed contract.
func NewLiquidFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidFilterer, error) {
	contract, err := bindLiquid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidFilterer{contract: contract}, nil
}

// bindLiquid binds a generic wrapper to an already deployed contract.
func bindLiquid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquid *LiquidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquid.Contract.LiquidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquid *LiquidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquid.Contract.LiquidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquid *LiquidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquid.Contract.LiquidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquid *LiquidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquid *LiquidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquid *LiquidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquid.Contract.contract.Transact(opts, method, params...)
}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Liquid *LiquidCaller) COMMISSIONRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "COMMISSION_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Liquid *LiquidSession) COMMISSIONRATEPRECISION() (*big.Int, error) {
	return _Liquid.Contract.COMMISSIONRATEPRECISION(&_Liquid.CallOpts)
}

// COMMISSIONRATEPRECISION is a free data retrieval call binding the contract method 0x2f2c3f2e.
//
// Solidity: function COMMISSION_RATE_PRECISION() view returns(uint256)
func (_Liquid *LiquidCallerSession) COMMISSIONRATEPRECISION() (*big.Int, error) {
	return _Liquid.Contract.COMMISSIONRATEPRECISION(&_Liquid.CallOpts)
}

// FEEFACTORUNITRECIP is a free data retrieval call binding the contract method 0x187cf4d7.
//
// Solidity: function FEE_FACTOR_UNIT_RECIP() view returns(uint256)
func (_Liquid *LiquidCaller) FEEFACTORUNITRECIP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "FEE_FACTOR_UNIT_RECIP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEFACTORUNITRECIP is a free data retrieval call binding the contract method 0x187cf4d7.
//
// Solidity: function FEE_FACTOR_UNIT_RECIP() view returns(uint256)
func (_Liquid *LiquidSession) FEEFACTORUNITRECIP() (*big.Int, error) {
	return _Liquid.Contract.FEEFACTORUNITRECIP(&_Liquid.CallOpts)
}

// FEEFACTORUNITRECIP is a free data retrieval call binding the contract method 0x187cf4d7.
//
// Solidity: function FEE_FACTOR_UNIT_RECIP() view returns(uint256)
func (_Liquid *LiquidCallerSession) FEEFACTORUNITRECIP() (*big.Int, error) {
	return _Liquid.Contract.FEEFACTORUNITRECIP(&_Liquid.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Liquid *LiquidCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Liquid *LiquidSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Liquid.Contract.Allowance(&_Liquid.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Liquid *LiquidCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Liquid.Contract.Allowance(&_Liquid.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCaller) BalanceOf(opts *bind.CallOpts, _delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "balanceOf", _delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidSession) BalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.BalanceOf(&_Liquid.CallOpts, _delegator)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCallerSession) BalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.BalanceOf(&_Liquid.CallOpts, _delegator)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Liquid *LiquidCaller) CommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "commissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Liquid *LiquidSession) CommissionRate() (*big.Int, error) {
	return _Liquid.Contract.CommissionRate(&_Liquid.CallOpts)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Liquid *LiquidCallerSession) CommissionRate() (*big.Int, error) {
	return _Liquid.Contract.CommissionRate(&_Liquid.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Liquid *LiquidCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Liquid *LiquidSession) Decimals() (uint8, error) {
	return _Liquid.Contract.Decimals(&_Liquid.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Liquid *LiquidCallerSession) Decimals() (uint8, error) {
	return _Liquid.Contract.Decimals(&_Liquid.CallOpts)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCaller) LockedBalanceOf(opts *bind.CallOpts, _delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "lockedBalanceOf", _delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidSession) LockedBalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.LockedBalanceOf(&_Liquid.CallOpts, _delegator)
}

// LockedBalanceOf is a free data retrieval call binding the contract method 0x59355736.
//
// Solidity: function lockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCallerSession) LockedBalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.LockedBalanceOf(&_Liquid.CallOpts, _delegator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquid *LiquidCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquid *LiquidSession) Name() (string, error) {
	return _Liquid.Contract.Name(&_Liquid.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquid *LiquidCallerSession) Name() (string, error) {
	return _Liquid.Contract.Name(&_Liquid.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquid *LiquidCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquid *LiquidSession) Symbol() (string, error) {
	return _Liquid.Contract.Symbol(&_Liquid.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquid *LiquidCallerSession) Symbol() (string, error) {
	return _Liquid.Contract.Symbol(&_Liquid.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquid *LiquidCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquid *LiquidSession) TotalSupply() (*big.Int, error) {
	return _Liquid.Contract.TotalSupply(&_Liquid.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquid *LiquidCallerSession) TotalSupply() (*big.Int, error) {
	return _Liquid.Contract.TotalSupply(&_Liquid.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquid *LiquidCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquid *LiquidSession) Treasury() (common.Address, error) {
	return _Liquid.Contract.Treasury(&_Liquid.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Liquid *LiquidCallerSession) Treasury() (common.Address, error) {
	return _Liquid.Contract.Treasury(&_Liquid.CallOpts)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address _account) view returns(uint256)
func (_Liquid *LiquidCaller) UnclaimedRewards(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "unclaimedRewards", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address _account) view returns(uint256)
func (_Liquid *LiquidSession) UnclaimedRewards(_account common.Address) (*big.Int, error) {
	return _Liquid.Contract.UnclaimedRewards(&_Liquid.CallOpts, _account)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address _account) view returns(uint256)
func (_Liquid *LiquidCallerSession) UnclaimedRewards(_account common.Address) (*big.Int, error) {
	return _Liquid.Contract.UnclaimedRewards(&_Liquid.CallOpts, _account)
}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCaller) UnlockedBalanceOf(opts *bind.CallOpts, _delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "unlockedBalanceOf", _delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidSession) UnlockedBalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.UnlockedBalanceOf(&_Liquid.CallOpts, _delegator)
}

// UnlockedBalanceOf is a free data retrieval call binding the contract method 0x84955c88.
//
// Solidity: function unlockedBalanceOf(address _delegator) view returns(uint256)
func (_Liquid *LiquidCallerSession) UnlockedBalanceOf(_delegator common.Address) (*big.Int, error) {
	return _Liquid.Contract.UnlockedBalanceOf(&_Liquid.CallOpts, _delegator)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Liquid *LiquidCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquid.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Liquid *LiquidSession) Validator() (common.Address, error) {
	return _Liquid.Contract.Validator(&_Liquid.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Liquid *LiquidCallerSession) Validator() (common.Address, error) {
	return _Liquid.Contract.Validator(&_Liquid.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Liquid *LiquidTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Liquid *LiquidSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Approve(&_Liquid.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Liquid *LiquidTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Approve(&_Liquid.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_Liquid *LiquidSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Burn(&_Liquid.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Burn(&_Liquid.TransactOpts, _account, _amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Liquid *LiquidTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Liquid *LiquidSession) ClaimRewards() (*types.Transaction, error) {
	return _Liquid.Contract.ClaimRewards(&_Liquid.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Liquid *LiquidTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _Liquid.Contract.ClaimRewards(&_Liquid.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactor) Lock(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "lock", _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Lock(&_Liquid.TransactOpts, _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactorSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Lock(&_Liquid.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Liquid *LiquidSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Mint(&_Liquid.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Mint(&_Liquid.TransactOpts, _account, _amount)
}

// Redistribute is a paid mutator transaction binding the contract method 0xfb489a7b.
//
// Solidity: function redistribute() payable returns(uint256)
func (_Liquid *LiquidTransactor) Redistribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "redistribute")
}

// Redistribute is a paid mutator transaction binding the contract method 0xfb489a7b.
//
// Solidity: function redistribute() payable returns(uint256)
func (_Liquid *LiquidSession) Redistribute() (*types.Transaction, error) {
	return _Liquid.Contract.Redistribute(&_Liquid.TransactOpts)
}

// Redistribute is a paid mutator transaction binding the contract method 0xfb489a7b.
//
// Solidity: function redistribute() payable returns(uint256)
func (_Liquid *LiquidTransactorSession) Redistribute() (*types.Transaction, error) {
	return _Liquid.Contract.Redistribute(&_Liquid.TransactOpts)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 _rate) returns()
func (_Liquid *LiquidTransactor) SetCommissionRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "setCommissionRate", _rate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 _rate) returns()
func (_Liquid *LiquidSession) SetCommissionRate(_rate *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.SetCommissionRate(&_Liquid.TransactOpts, _rate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 _rate) returns()
func (_Liquid *LiquidTransactorSession) SetCommissionRate(_rate *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.SetCommissionRate(&_Liquid.TransactOpts, _rate)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Transfer(&_Liquid.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Transfer(&_Liquid.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.TransferFrom(&_Liquid.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool _success)
func (_Liquid *LiquidTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.TransferFrom(&_Liquid.TransactOpts, _sender, _recipient, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactor) Unlock(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.contract.Transact(opts, "unlock", _account, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidSession) Unlock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Unlock(&_Liquid.TransactOpts, _account, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _account, uint256 _amount) returns()
func (_Liquid *LiquidTransactorSession) Unlock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquid.Contract.Unlock(&_Liquid.TransactOpts, _account, _amount)
}

// LiquidApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Liquid contract.
type LiquidApprovalIterator struct {
	Event *LiquidApproval // Event containing the contract specifics and raw log

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
func (it *LiquidApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidApproval)
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
		it.Event = new(LiquidApproval)
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
func (it *LiquidApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidApproval represents a Approval event raised by the Liquid contract.
type LiquidApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquid *LiquidFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LiquidApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Liquid.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidApprovalIterator{contract: _Liquid.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquid *LiquidFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LiquidApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Liquid.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidApproval)
				if err := _Liquid.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquid *LiquidFilterer) ParseApproval(log types.Log) (*LiquidApproval, error) {
	event := new(LiquidApproval)
	if err := _Liquid.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Liquid contract.
type LiquidTransferIterator struct {
	Event *LiquidTransfer // Event containing the contract specifics and raw log

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
func (it *LiquidTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidTransfer)
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
		it.Event = new(LiquidTransfer)
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
func (it *LiquidTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidTransfer represents a Transfer event raised by the Liquid contract.
type LiquidTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquid *LiquidFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LiquidTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Liquid.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LiquidTransferIterator{contract: _Liquid.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquid *LiquidFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LiquidTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Liquid.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidTransfer)
				if err := _Liquid.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquid *LiquidFilterer) ParseTransfer(log types.Log) (*LiquidTransfer, error) {
	event := new(LiquidTransfer)
	if err := _Liquid.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_autonity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_votePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_votePeriod\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"NewSymbols\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"_votes\",\"type\":\"int256[]\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getRoundData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structIOracle.RoundData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbols\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRoundBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVoterUpdateRound\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"latestRoundData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structIOracle.RoundData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newSymbols\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reports\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"}],\"name\":\"setSymbols\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newVoters\",\"type\":\"address[]\"}],\"name\":\"setVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbolUpdatedRound\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"symbols\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commit\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"_reports\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVoter\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"4bb278f3": "finalize()",
		"9670c0bc": "getPrecision()",
		"9f8743f7": "getRound()",
		"3c8510fd": "getRoundData(uint256,string)",
		"df7f710e": "getSymbols()",
		"b78dec52": "getVotePeriod()",
		"cdd72253": "getVoters()",
		"e6a02a28": "lastRoundBlock()",
		"aa2f89b5": "lastVoterUpdateRound()",
		"33f98c77": "latestRoundData(string)",
		"5281b5c6": "newSymbols(uint256)",
		"4c56ea56": "reports(string,address)",
		"146ca531": "round()",
		"b3ab15fb": "setOperator(address)",
		"8d4f75d2": "setSymbols(string[])",
		"845023f2": "setVoters(address[])",
		"08f21ff5": "symbolUpdatedRound()",
		"ccce413b": "symbols(uint256)",
		"307de9b6": "vote(uint256,int256[],uint256)",
		"a7813587": "votePeriod()",
		"5412b3ae": "votingInfo(address)",
	},
	Bin: "0x6080604052600160ff1b600755600160ff1b60085534801562000020575f80fd5b5060405162002cd338038062002cd383398101604081905262000043916200061d565b600280546001600160a01b038087166001600160a01b0319928316179092556003805492861692909116919091179055815162000087905f90602085019062000358565b5081516200009d90600190602085019062000358565b5080600981905550620000c1855f60018851620000bb919062000728565b6200017b565b8451620000d6906004906020880190620003b3565b508451620000ec906005906020880190620003b3565b5060016006819055600d805490910181555f9081525b85518110156200016f576001600b5f88848151811062000126576200012662000744565b6020908102919091018101516001600160a01b031682528101919091526040015f20600201805460ff191691151591909117905580620001668162000758565b91505062000102565b5050505050506200098a565b8082126200018857505050565b81815f8560026200019a858562000773565b620001a691906200079c565b620001b29087620007d8565b81518110620001c557620001c562000744565b602002602001015190505b81831362000324575b806001600160a01b0316868481518110620001f857620001f862000744565b60200260200101516001600160a01b031610156200022557826200021c8162000802565b935050620001d9565b806001600160a01b031686838151811062000244576200024462000744565b60200260200101516001600160a01b0316111562000271578162000268816200081c565b92505062000225565b8183136200031e578582815181106200028e576200028e62000744565b6020026020010151868481518110620002ab57620002ab62000744565b6020026020010151878581518110620002c857620002c862000744565b60200260200101888581518110620002e457620002e462000744565b6001600160a01b03938416602091820292909201015291169052826200030a8162000802565b93505081806200031a906200081c565b9250505b620001d0565b818512156200033a576200033a8686846200017b565b838312156200035057620003508684866200017b565b505050505050565b828054828255905f5260205f20908101928215620003a1579160200282015b82811115620003a15782518290620003909082620008c2565b509160200191906001019062000377565b50620003af92915062000417565b5090565b828054828255905f5260205f2090810192821562000409579160200282015b828111156200040957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003d2565b50620003af92915062000437565b80821115620003af575f6200042d82826200044d565b5060010162000417565b5b80821115620003af575f815560010162000438565b5080546200045b906200083a565b5f825580601f106200046b575050565b601f0160209004905f5260205f209081019062000489919062000437565b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715620004cb57620004cb6200048c565b604052919050565b5f6001600160401b03821115620004ee57620004ee6200048c565b5060051b60200190565b80516001600160a01b03811681146200050f575f80fd5b919050565b5f601f838184011262000525575f80fd5b825160206200053e6200053883620004d3565b620004a0565b82815260059290921b850181019181810190878411156200055d575f80fd5b8287015b84811015620006115780516001600160401b038082111562000582575f8081fd5b818a0191508a603f83011262000597575f8081fd5b8582015181811115620005ae57620005ae6200048c565b620005c1818a01601f19168801620004a0565b915080825260408c81838601011115620005da575f8081fd5b5f5b82811015620005f9578481018201518482018a01528801620005dc565b50505f90820187015284525091830191830162000561565b50979650505050505050565b5f805f805f60a0868803121562000632575f80fd5b85516001600160401b038082111562000649575f80fd5b818801915088601f8301126200065d575f80fd5b81516020620006706200053883620004d3565b82815260059290921b8401810191818101908c8411156200068f575f80fd5b948201945b83861015620006b857620006a886620004f8565b8252948201949082019062000694565b9950620006c990508a8201620004f8565b97505050620006db60408901620004f8565b94506060880151915080821115620006f1575f80fd5b50620007008882890162000514565b925050608086015190509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b818103818111156200073e576200073e62000714565b92915050565b634e487b7160e01b5f52603260045260245ffd5b5f600182016200076c576200076c62000714565b5060010190565b8181035f83128015838313168383128216171562000795576200079562000714565b5092915050565b5f82620007b757634e487b7160e01b5f52601260045260245ffd5b600160ff1b82145f1984141615620007d357620007d362000714565b500590565b8082018281125f831280158216821582161715620007fa57620007fa62000714565b505092915050565b5f6001600160ff1b0182016200076c576200076c62000714565b5f600160ff1b820162000833576200083362000714565b505f190190565b600181811c908216806200084f57607f821691505b6020821081036200086e57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620008bd575f81815260208120601f850160051c810160208610156200089c5750805b601f850160051c820191505b818110156200035057828155600101620008a8565b505050565b81516001600160401b03811115620008de57620008de6200048c565b620008f681620008ef84546200083a565b8462000874565b602080601f8311600181146200092c575f8415620009145750858301515b5f19600386901b1c1916600185901b17855562000350565b5f85815260208120601f198616915b828110156200095c578886015182559484019460019091019084016200093b565b50858210156200097a57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b61233b80620009985f395ff3fe608060405260043610610121575f3560e01c80638d4f75d2116100a4578063b3ab15fb1161006b578063b3ab15fb14610378578063b78dec5214610397578063ccce413b146103ab578063cdd72253146103ca578063df7f710e146103eb578063e6a02a281461040c57005b80638d4f75d2146103065780639670c0bc146103255780639f8743f71461033a578063a78135871461034e578063aa2f89b51461036357005b80634bb278f3116100e85780634bb278f3146101f75780634c56ea561461021b5780635281b5c6146102615780635412b3ae1461028d578063845023f2146102e757005b806308f21ff51461012a578063146ca53114610152578063307de9b61461016757806333f98c77146101865780633c8510fd146101d857005b3661012857005b005b348015610135575f80fd5b5061013f60085481565b6040519081526020015b60405180910390f35b34801561015d575f80fd5b5061013f60065481565b348015610172575f80fd5b506101286101813660046119a5565b610421565b348015610191575f80fd5b506101a56101a0366004611ad4565b610658565b60405161014991908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b3480156101e3575f80fd5b506101a56101f2366004611b06565b610770565b348015610202575f80fd5b5061020b61086f565b6040519015158152602001610149565b348015610226575f80fd5b5061013f610235366004611b65565b8151602081840181018051600c825292820194820194909420919093529091525f908152604090205481565b34801561026c575f80fd5b5061028061027b366004611bb0565b610a09565b6040516101499190611c14565b348015610298575f80fd5b506102ca6102a7366004611c2d565b600b6020525f908152604090208054600182015460029092015490919060ff1683565b604080519384526020840192909252151590820152606001610149565b3480156102f2575f80fd5b50610128610301366004611c69565b610aaf565b348015610311575f80fd5b50610128610320366004611d01565b610b56565b348015610330575f80fd5b506298968061013f565b348015610345575f80fd5b5060065461013f565b348015610359575f80fd5b5061013f60095481565b34801561036e575f80fd5b5061013f60075481565b348015610383575f80fd5b50610128610392366004611c2d565b610cc5565b3480156103a2575f80fd5b5060095461013f565b3480156103b6575f80fd5b506102806103c5366004611bb0565b610d11565b3480156103d5575f80fd5b506103de610d1f565b6040516101499190611dac565b3480156103f6575f80fd5b506103ff610d7f565b6040516101499190611e4b565b348015610417575f80fd5b5061013f600a5481565b335f908152600b602052604090206002015460ff166104875760405162461bcd60e51b815260206004820152601960248201527f7265737472696374656420746f206f6e6c7920766f746572730000000000000060448201526064015b60405180910390fd5b600654335f908152600b6020526040902054036104d65760405162461bcd60e51b815260206004820152600d60248201526c185b1c9958591e481d9bdd1959609a1b604482015260640161047e565b335f908152600b6020526040812060018101805490879055815460065490925591819003610505575050610652565b5f548414610514575050610652565b60016006546105239190611e71565b8114158061055f5750848484336040516020016105439493929190611e84565b604051602081830303815290604052805190602001205f1c8214155b156105d5575f5b5f548110156105cd576001600160ff1b03600c5f838154811061058b5761058b611ece565b905f5260205f20016040516105a09190611f1a565b9081526040805160209281900383019020335f9081529252902055806105c581611f8c565b915050610566565b505050610652565b5f5b8481101561064e578585828181106105f1576105f1611ece565b90506020020135600c5f838154811061060c5761060c611ece565b905f5260205f20016040516106219190611f1a565b9081526040805160209281900383019020335f90815292529020558061064681611f8c565b9150506105d7565b5050505b50505050565b61067f60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5f600d60016006546106919190611e71565b815481106106a1576106a1611ece565b905f5260205f2001836040516106b79190611fa4565b90815260200160405180910390206040518060600160405290815f820154815260200160018201548152602001600282015f9054906101000a900460ff16600181111561070657610706611fbf565b600181111561071757610717611fbf565b8152505090505f604051806080016040528060016006546107389190611e71565b8152602001835f01518152602001836020015181526020018360400151600181111561076657610766611fbf565b9052949350505050565b61079760405180608001604052805f81526020015f81526020015f81526020015f81525090565b5f600d84815481106107ab576107ab611ece565b905f5260205f2001836040516107c19190611fa4565b90815260200160405180910390206040518060600160405290815f820154815260200160018201548152602001600282015f9054906101000a900460ff16600181111561081057610810611fbf565b600181111561082157610821611fbf565b8152505090505f6040518060800160405280868152602001835f01518152602001836020015181526020018360400151600181111561086257610862611fbf565b9052925050505b92915050565b6002545f906001600160a01b0316331461089b5760405162461bcd60e51b815260040161047e90611fd3565b600954600a546108ab9190612016565b4310610a04575f5b5f548110156108d7576108c581610f32565b6108d0600182612016565b90506108b3565b5060065460075403610951575f5b60055481101561094f576001600b5f6005848154811061090757610907611ece565b5f918252602080832091909101546001600160a01b031683528201929092526040019020600201805460ff19169115159190911790558061094781611f8c565b9150506108e5565b505b600654600754610962906001612029565b0361096f5761096f611252565b43600a81905550600160065f8282546109889190612016565b909155505060085461099b906002612029565b600654036109b457600180546109b2915f91611808565b505b60065460095460408051928352436020840152429083015260608201527fb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e59060800160405180910390a150600190565b505f90565b60018181548110610a18575f80fd5b905f5260205f20015f915090508054610a3090611ee2565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5c90611ee2565b8015610aa75780601f10610a7e57610100808354040283529160200191610aa7565b820191905f5260205f20905b815481529060010190602001808311610a8a57829003601f168201915b505050505081565b6002546001600160a01b03163314610ad95760405162461bcd60e51b815260040161047e90611fd3565b80515f03610b215760405162461bcd60e51b8152602060048201526015602482015274566f746572732063616e277420626520656d70747960581b604482015260640161047e565b610b39815f60018451610b349190611e71565b611426565b8051610b4c90600590602084019061185c565b5050600654600755565b6003546001600160a01b03163314610ba95760405162461bcd60e51b81526020600482015260166024820152753932b9ba3934b1ba32b2103a379037b832b930ba37b960511b604482015260640161047e565b80515f03610bf25760405162461bcd60e51b815260206004820152601660248201527573796d626f6c732063616e277420626520656d70747960501b604482015260640161047e565b600654600854610c03906001612029565b14158015610c15575060065460085414155b610c615760405162461bcd60e51b815260206004820152601e60248201527f63616e2774206265207570646174656420696e207468697320726f756e640000604482015260640161047e565b8051610c749060019060208401906118bb565b5060065460088190557faa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d908290610cac906001612016565b604051610cba929190612050565b60405180910390a150565b6002546001600160a01b03163314610cef5760405162461bcd60e51b815260040161047e90611fd3565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b5f8181548110610a18575f80fd5b60606005805480602002602001604051908101604052809291908181526020018280548015610d7557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610d57575b5050505050905090565b60606006546008546001610d939190612029565b03610e6a576001805480602002602001604051908101604052809291908181526020015f905b82821015610e61578382905f5260205f20018054610dd690611ee2565b80601f0160208091040260200160405190810160405280929190818152602001828054610e0290611ee2565b8015610e4d5780601f10610e2457610100808354040283529160200191610e4d565b820191905f5260205f20905b815481529060010190602001808311610e3057829003601f168201915b505050505081526020019060010190610db9565b50505050905090565b5f805480602002602001604051908101604052809291908181526020015f905b82821015610e61578382905f5260205f20018054610ea790611ee2565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed390611ee2565b8015610f1e5780601f10610ef557610100808354040283529160200191610f1e565b820191905f5260205f20905b815481529060010190602001808311610f0157829003601f168201915b505050505081526020019060010190610e8a565b5f808281548110610f4557610f45611ece565b905f5260205f20018054610f5890611ee2565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8490611ee2565b8015610fcf5780601f10610fa657610100808354040283529160200191610fcf565b820191905f5260205f20905b815481529060010190602001808311610fb257829003601f168201915b505050505090505f60048054905067ffffffffffffffff811115610ff557610ff5611a23565b60405190808252806020026020018201604052801561101e578160200160208202803683370190505b5090505f805b600454811015611130575f6004828154811061104257611042611ece565b5f9182526020808320909101546006546001600160a01b03909116808452600b9092526040909220549092501415806110b457506001600160ff1b03600c8660405161108e9190611fa4565b90815260408051602092819003830190206001600160a01b0385165f9081529252902054145b156110bf575061111e565b600c856040516110cf9190611fa4565b90815260408051602092819003830190206001600160a01b0384165f908152925290205484846110fe81611f8c565b95508151811061111057611110611ece565b602002602001018181525050505b8061112881611f8c565b915050611024565b505f600d60016006546111439190611e71565b8154811061115357611153611ece565b905f5260205f2001846040516111699190611fa4565b908152604051908190036020019020549050600182156111935761118d84846115d2565b91505f90505b600d8054600190810182555f9190915260408051606081018252848152426020820152919082019083908111156111cc576111cc611fbf565b815250600d600654815481106111e4576111e4611ece565b905f5260205f2001866040516111fa9190611fa4565b90815260200160405180910390205f820151815f0155602082015181600101556040820151816002015f6101000a81548160ff0219169083600181111561124357611243611fbf565b02179055505050505050505050565b5f805b60045482108015611267575060055481105b156113a6576005818154811061127f5761127f611ece565b5f91825260209091200154600480546001600160a01b0390921691849081106112aa576112aa611ece565b5f918252602090912001546001600160a01b0316036112e357816112cd81611f8c565b92505080806112db90611f8c565b915050611255565b600581815481106112f6576112f6611ece565b5f91825260209091200154600480546001600160a01b03909216918490811061132157611321611ece565b5f918252602090912001546001600160a01b0316101561139c57600b5f6004848154811061135157611351611ece565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208181556001810191909155600201805460ff191690558161139481611f8c565b925050611255565b806112db81611f8c565b60045482101561141157600b5f600484815481106113c6576113c6611ece565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208181556001810191909155600201805460ff191690558161140981611f8c565b9250506113a6565b60058054611421916004916118ff565b505050565b80821261143257505050565b81815f8560026114428585612071565b61144c91906120ab565b6114569087612029565b8151811061146657611466611ece565b602002602001015190505b8183136115a4575b806001600160a01b031686848151811061149557611495611ece565b60200260200101516001600160a01b031610156114be57826114b6816120d7565b935050611479565b806001600160a01b03168683815181106114da576114da611ece565b60200260200101516001600160a01b0316111561150357816114fb816120ee565b9250506114be565b81831361159f5785828151811061151c5761151c611ece565b602002602001015186848151811061153657611536611ece565b602002602001015187858151811061155057611550611ece565b6020026020010188858151811061156957611569611ece565b6001600160a01b039384166020918202929092010152911690528261158d816120d7565b935050818061159b906120ee565b9250505b611471565b818512156115b7576115b7868684611426565b838312156115ca576115ca868486611426565b505050505050565b5f815f036115e157505f610869565b6115f6835f6115f1600186611e71565b611691565b5f611602600284612109565b905061160f60028461211c565b156116335783818151811061162657611626611ece565b6020026020010151611689565b600284828151811061164757611647611ece565b60200260200101518560018461165d9190611e71565b8151811061166d5761166d611ece565b602002602001015161167f9190612029565b61168991906120ab565b949350505050565b81818082036116a1575050505050565b5f8560026116af8787612071565b6116b991906120ab565b6116c39087612029565b815181106116d3576116d3611ece565b602002602001015190505b8183136117e2575b808684815181106116f9576116f9611ece565b602002602001015112156117195782611711816120d7565b9350506116e6565b85828151811061172b5761172b611ece565b602002602001015181121561174c5781611744816120ee565b925050611719565b8183136117dd5785828151811061176557611765611ece565b602002602001015186848151811061177f5761177f611ece565b602002602001015187858151811061179957611799611ece565b602002602001018885815181106117b2576117b2611ece565b602090810291909101019190915252816117cb816120ee565b92505082806117d9906120d7565b9350505b6116de565b818512156117f5576117f5868684611691565b838312156115ca576115ca868486611691565b828054828255905f5260205f2090810192821561184c575f5260205f209182015b8281111561184c578161183c8482612174565b5091600101919060010190611829565b5061185892915061193b565b5090565b828054828255905f5260205f209081019282156118af579160200282015b828111156118af57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061187a565b50611858929150611957565b828054828255905f5260205f2090810192821561184c579160200282015b8281111561184c57825182906118ef908261224f565b50916020019190600101906118d9565b828054828255905f5260205f209081019282156118af575f5260205f209182015b828111156118af578254825591600101919060010190611920565b80821115611858575f61194e828261196b565b5060010161193b565b5b80821115611858575f8155600101611958565b50805461197790611ee2565b5f825580601f10611986575050565b601f0160209004905f5260205f20908101906119a29190611957565b50565b5f805f80606085870312156119b8575f80fd5b84359350602085013567ffffffffffffffff808211156119d6575f80fd5b818701915087601f8301126119e9575f80fd5b8135818111156119f7575f80fd5b8860208260051b8501011115611a0b575f80fd5b95986020929092019750949560400135945092505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611a6057611a60611a23565b604052919050565b5f82601f830112611a77575f80fd5b813567ffffffffffffffff811115611a9157611a91611a23565b611aa4601f8201601f1916602001611a37565b818152846020838601011115611ab8575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611ae4575f80fd5b813567ffffffffffffffff811115611afa575f80fd5b61168984828501611a68565b5f8060408385031215611b17575f80fd5b82359150602083013567ffffffffffffffff811115611b34575f80fd5b611b4085828601611a68565b9150509250929050565b80356001600160a01b0381168114611b60575f80fd5b919050565b5f8060408385031215611b76575f80fd5b823567ffffffffffffffff811115611b8c575f80fd5b611b9885828601611a68565b925050611ba760208401611b4a565b90509250929050565b5f60208284031215611bc0575f80fd5b5035919050565b5f5b83811015611be1578181015183820152602001611bc9565b50505f910152565b5f8151808452611c00816020860160208601611bc7565b601f01601f19169290920160200192915050565b602081525f611c266020830184611be9565b9392505050565b5f60208284031215611c3d575f80fd5b611c2682611b4a565b5f67ffffffffffffffff821115611c5f57611c5f611a23565b5060051b60200190565b5f6020808385031215611c7a575f80fd5b823567ffffffffffffffff811115611c90575f80fd5b8301601f81018513611ca0575f80fd5b8035611cb3611cae82611c46565b611a37565b81815260059190911b82018301908381019087831115611cd1575f80fd5b928401925b82841015611cf657611ce784611b4a565b82529284019290840190611cd6565b979650505050505050565b5f6020808385031215611d12575f80fd5b823567ffffffffffffffff80821115611d29575f80fd5b818501915085601f830112611d3c575f80fd5b8135611d4a611cae82611c46565b81815260059190911b83018401908481019088831115611d68575f80fd5b8585015b83811015611d9f57803585811115611d83575f8081fd5b611d918b89838a0101611a68565b845250918601918601611d6c565b5098975050505050505050565b602080825282518282018190525f9190848201906040850190845b81811015611dec5783516001600160a01b031683529284019291840191600101611dc7565b50909695505050505050565b5f81518084526020808501808196508360051b810191508286015f5b85811015611e3e578284038952611e2c848351611be9565b98850198935090840190600101611e14565b5091979650505050505050565b602081525f611c266020830184611df8565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561086957610869611e5d565b5f8186825b87811015611ea7578135835260209283019290910190600101611e89565b5050938452505060601b6bffffffffffffffffffffffff1916602082015260340192915050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680611ef657607f821691505b602082108103611f1457634e487b7160e01b5f52602260045260245ffd5b50919050565b5f808354611f2781611ee2565b60018281168015611f3f5760018114611f5457611f80565b60ff1984168752821515830287019450611f80565b875f526020805f205f5b85811015611f775781548a820152908401908201611f5e565b50505082870194505b50929695505050505050565b5f60018201611f9d57611f9d611e5d565b5060010190565b5f8251611fb5818460208701611bc7565b9190910192915050565b634e487b7160e01b5f52602160045260245ffd5b60208082526023908201527f7265737472696374656420746f20746865206175746f6e69747920636f6e74726040820152621858dd60ea1b606082015260800190565b8082018082111561086957610869611e5d565b8082018281125f83128015821682158216171561204857612048611e5d565b505092915050565b604081525f6120626040830185611df8565b90508260208301529392505050565b8181035f83128015838313168383128216171561209057612090611e5d565b5092915050565b634e487b7160e01b5f52601260045260245ffd5b5f826120b9576120b9612097565b600160ff1b82145f19841416156120d2576120d2611e5d565b500590565b5f6001600160ff1b018201611f9d57611f9d611e5d565b5f600160ff1b820161210257612102611e5d565b505f190190565b5f8261211757612117612097565b500490565b5f8261212a5761212a612097565b500690565b601f821115611421575f81815260208120601f850160051c810160208610156121555750805b601f850160051c820191505b818110156115ca57828155600101612161565b81810361217f575050565b6121898254611ee2565b67ffffffffffffffff8111156121a1576121a1611a23565b6121b5816121af8454611ee2565b8461212f565b5f601f8211600181146121e6575f83156121cf5750848201545b5f19600385901b1c1916600184901b178455612248565b5f85815260209020601f198416905f86815260209020845b8381101561221e57828601548255600195860195909101906020016121fe565b508583101561223b57818501545f19600388901b60f8161c191681555b50505060018360011b0184555b5050505050565b815167ffffffffffffffff81111561226957612269611a23565b612277816121af8454611ee2565b602080601f8311600181146122aa575f84156122935750858301515b5f19600386901b1c1916600185901b1785556115ca565b5f85815260208120601f198616915b828110156122d8578886015182559484019460019091019084016122b9565b50858210156122f557878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220caad931a2f49d9dde758f3b2cee8b3282d200c55c9b9a36a120cd60782f8da3964736f6c63430008150033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Deprecated: Use OracleMetaData.Sigs instead.
// OracleFuncSigs maps the 4-byte function signature to its string representation.
var OracleFuncSigs = OracleMetaData.Sigs

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _voters []common.Address, _autonity common.Address, _operator common.Address, _symbols []string, _votePeriod *big.Int) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend, _voters, _autonity, _operator, _symbols, _votePeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Oracle *OracleCaller) GetPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Oracle *OracleSession) GetPrecision() (*big.Int, error) {
	return _Oracle.Contract.GetPrecision(&_Oracle.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Oracle *OracleCallerSession) GetPrecision() (*big.Int, error) {
	return _Oracle.Contract.GetPrecision(&_Oracle.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_Oracle *OracleCaller) GetRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_Oracle *OracleSession) GetRound() (*big.Int, error) {
	return _Oracle.Contract.GetRound(&_Oracle.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_Oracle *OracleCallerSession) GetRound() (*big.Int, error) {
	return _Oracle.Contract.GetRound(&_Oracle.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleCaller) GetRoundData(opts *bind.CallOpts, _round *big.Int, _symbol string) (IOracleRoundData, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRoundData", _round, _symbol)

	if err != nil {
		return *new(IOracleRoundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleRoundData)).(*IOracleRoundData)

	return out0, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleSession) GetRoundData(_round *big.Int, _symbol string) (IOracleRoundData, error) {
	return _Oracle.Contract.GetRoundData(&_Oracle.CallOpts, _round, _symbol)
}

// GetRoundData is a free data retrieval call binding the contract method 0x3c8510fd.
//
// Solidity: function getRoundData(uint256 _round, string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleCallerSession) GetRoundData(_round *big.Int, _symbol string) (IOracleRoundData, error) {
	return _Oracle.Contract.GetRoundData(&_Oracle.CallOpts, _round, _symbol)
}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[])
func (_Oracle *OracleCaller) GetSymbols(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getSymbols")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[])
func (_Oracle *OracleSession) GetSymbols() ([]string, error) {
	return _Oracle.Contract.GetSymbols(&_Oracle.CallOpts)
}

// GetSymbols is a free data retrieval call binding the contract method 0xdf7f710e.
//
// Solidity: function getSymbols() view returns(string[])
func (_Oracle *OracleCallerSession) GetSymbols() ([]string, error) {
	return _Oracle.Contract.GetSymbols(&_Oracle.CallOpts)
}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_Oracle *OracleCaller) GetVotePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getVotePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_Oracle *OracleSession) GetVotePeriod() (*big.Int, error) {
	return _Oracle.Contract.GetVotePeriod(&_Oracle.CallOpts)
}

// GetVotePeriod is a free data retrieval call binding the contract method 0xb78dec52.
//
// Solidity: function getVotePeriod() view returns(uint256)
func (_Oracle *OracleCallerSession) GetVotePeriod() (*big.Int, error) {
	return _Oracle.Contract.GetVotePeriod(&_Oracle.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Oracle *OracleCaller) GetVoters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getVoters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Oracle *OracleSession) GetVoters() ([]common.Address, error) {
	return _Oracle.Contract.GetVoters(&_Oracle.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0xcdd72253.
//
// Solidity: function getVoters() view returns(address[])
func (_Oracle *OracleCallerSession) GetVoters() ([]common.Address, error) {
	return _Oracle.Contract.GetVoters(&_Oracle.CallOpts)
}

// LastRoundBlock is a free data retrieval call binding the contract method 0xe6a02a28.
//
// Solidity: function lastRoundBlock() view returns(uint256)
func (_Oracle *OracleCaller) LastRoundBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "lastRoundBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRoundBlock is a free data retrieval call binding the contract method 0xe6a02a28.
//
// Solidity: function lastRoundBlock() view returns(uint256)
func (_Oracle *OracleSession) LastRoundBlock() (*big.Int, error) {
	return _Oracle.Contract.LastRoundBlock(&_Oracle.CallOpts)
}

// LastRoundBlock is a free data retrieval call binding the contract method 0xe6a02a28.
//
// Solidity: function lastRoundBlock() view returns(uint256)
func (_Oracle *OracleCallerSession) LastRoundBlock() (*big.Int, error) {
	return _Oracle.Contract.LastRoundBlock(&_Oracle.CallOpts)
}

// LastVoterUpdateRound is a free data retrieval call binding the contract method 0xaa2f89b5.
//
// Solidity: function lastVoterUpdateRound() view returns(int256)
func (_Oracle *OracleCaller) LastVoterUpdateRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "lastVoterUpdateRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVoterUpdateRound is a free data retrieval call binding the contract method 0xaa2f89b5.
//
// Solidity: function lastVoterUpdateRound() view returns(int256)
func (_Oracle *OracleSession) LastVoterUpdateRound() (*big.Int, error) {
	return _Oracle.Contract.LastVoterUpdateRound(&_Oracle.CallOpts)
}

// LastVoterUpdateRound is a free data retrieval call binding the contract method 0xaa2f89b5.
//
// Solidity: function lastVoterUpdateRound() view returns(int256)
func (_Oracle *OracleCallerSession) LastVoterUpdateRound() (*big.Int, error) {
	return _Oracle.Contract.LastVoterUpdateRound(&_Oracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleCaller) LatestRoundData(opts *bind.CallOpts, _symbol string) (IOracleRoundData, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestRoundData", _symbol)

	if err != nil {
		return *new(IOracleRoundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleRoundData)).(*IOracleRoundData)

	return out0, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleSession) LatestRoundData(_symbol string) (IOracleRoundData, error) {
	return _Oracle.Contract.LatestRoundData(&_Oracle.CallOpts, _symbol)
}

// LatestRoundData is a free data retrieval call binding the contract method 0x33f98c77.
//
// Solidity: function latestRoundData(string _symbol) view returns((uint256,int256,uint256,uint256) data)
func (_Oracle *OracleCallerSession) LatestRoundData(_symbol string) (IOracleRoundData, error) {
	return _Oracle.Contract.LatestRoundData(&_Oracle.CallOpts, _symbol)
}

// NewSymbols is a free data retrieval call binding the contract method 0x5281b5c6.
//
// Solidity: function newSymbols(uint256 ) view returns(string)
func (_Oracle *OracleCaller) NewSymbols(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "newSymbols", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NewSymbols is a free data retrieval call binding the contract method 0x5281b5c6.
//
// Solidity: function newSymbols(uint256 ) view returns(string)
func (_Oracle *OracleSession) NewSymbols(arg0 *big.Int) (string, error) {
	return _Oracle.Contract.NewSymbols(&_Oracle.CallOpts, arg0)
}

// NewSymbols is a free data retrieval call binding the contract method 0x5281b5c6.
//
// Solidity: function newSymbols(uint256 ) view returns(string)
func (_Oracle *OracleCallerSession) NewSymbols(arg0 *big.Int) (string, error) {
	return _Oracle.Contract.NewSymbols(&_Oracle.CallOpts, arg0)
}

// Reports is a free data retrieval call binding the contract method 0x4c56ea56.
//
// Solidity: function reports(string , address ) view returns(int256)
func (_Oracle *OracleCaller) Reports(opts *bind.CallOpts, arg0 string, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "reports", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reports is a free data retrieval call binding the contract method 0x4c56ea56.
//
// Solidity: function reports(string , address ) view returns(int256)
func (_Oracle *OracleSession) Reports(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Reports(&_Oracle.CallOpts, arg0, arg1)
}

// Reports is a free data retrieval call binding the contract method 0x4c56ea56.
//
// Solidity: function reports(string , address ) view returns(int256)
func (_Oracle *OracleCallerSession) Reports(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Oracle.Contract.Reports(&_Oracle.CallOpts, arg0, arg1)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Oracle *OracleCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Oracle *OracleSession) Round() (*big.Int, error) {
	return _Oracle.Contract.Round(&_Oracle.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Oracle *OracleCallerSession) Round() (*big.Int, error) {
	return _Oracle.Contract.Round(&_Oracle.CallOpts)
}

// SymbolUpdatedRound is a free data retrieval call binding the contract method 0x08f21ff5.
//
// Solidity: function symbolUpdatedRound() view returns(int256)
func (_Oracle *OracleCaller) SymbolUpdatedRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "symbolUpdatedRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SymbolUpdatedRound is a free data retrieval call binding the contract method 0x08f21ff5.
//
// Solidity: function symbolUpdatedRound() view returns(int256)
func (_Oracle *OracleSession) SymbolUpdatedRound() (*big.Int, error) {
	return _Oracle.Contract.SymbolUpdatedRound(&_Oracle.CallOpts)
}

// SymbolUpdatedRound is a free data retrieval call binding the contract method 0x08f21ff5.
//
// Solidity: function symbolUpdatedRound() view returns(int256)
func (_Oracle *OracleCallerSession) SymbolUpdatedRound() (*big.Int, error) {
	return _Oracle.Contract.SymbolUpdatedRound(&_Oracle.CallOpts)
}

// Symbols is a free data retrieval call binding the contract method 0xccce413b.
//
// Solidity: function symbols(uint256 ) view returns(string)
func (_Oracle *OracleCaller) Symbols(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "symbols", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbols is a free data retrieval call binding the contract method 0xccce413b.
//
// Solidity: function symbols(uint256 ) view returns(string)
func (_Oracle *OracleSession) Symbols(arg0 *big.Int) (string, error) {
	return _Oracle.Contract.Symbols(&_Oracle.CallOpts, arg0)
}

// Symbols is a free data retrieval call binding the contract method 0xccce413b.
//
// Solidity: function symbols(uint256 ) view returns(string)
func (_Oracle *OracleCallerSession) Symbols(arg0 *big.Int) (string, error) {
	return _Oracle.Contract.Symbols(&_Oracle.CallOpts, arg0)
}

// VotePeriod is a free data retrieval call binding the contract method 0xa7813587.
//
// Solidity: function votePeriod() view returns(uint256)
func (_Oracle *OracleCaller) VotePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "votePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePeriod is a free data retrieval call binding the contract method 0xa7813587.
//
// Solidity: function votePeriod() view returns(uint256)
func (_Oracle *OracleSession) VotePeriod() (*big.Int, error) {
	return _Oracle.Contract.VotePeriod(&_Oracle.CallOpts)
}

// VotePeriod is a free data retrieval call binding the contract method 0xa7813587.
//
// Solidity: function votePeriod() view returns(uint256)
func (_Oracle *OracleCallerSession) VotePeriod() (*big.Int, error) {
	return _Oracle.Contract.VotePeriod(&_Oracle.CallOpts)
}

// VotingInfo is a free data retrieval call binding the contract method 0x5412b3ae.
//
// Solidity: function votingInfo(address ) view returns(uint256 round, uint256 commit, bool isVoter)
func (_Oracle *OracleCaller) VotingInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Round   *big.Int
	Commit  *big.Int
	IsVoter bool
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "votingInfo", arg0)

	outstruct := new(struct {
		Round   *big.Int
		Commit  *big.Int
		IsVoter bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Round = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsVoter = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// VotingInfo is a free data retrieval call binding the contract method 0x5412b3ae.
//
// Solidity: function votingInfo(address ) view returns(uint256 round, uint256 commit, bool isVoter)
func (_Oracle *OracleSession) VotingInfo(arg0 common.Address) (struct {
	Round   *big.Int
	Commit  *big.Int
	IsVoter bool
}, error) {
	return _Oracle.Contract.VotingInfo(&_Oracle.CallOpts, arg0)
}

// VotingInfo is a free data retrieval call binding the contract method 0x5412b3ae.
//
// Solidity: function votingInfo(address ) view returns(uint256 round, uint256 commit, bool isVoter)
func (_Oracle *OracleCallerSession) VotingInfo(arg0 common.Address) (struct {
	Round   *big.Int
	Commit  *big.Int
	IsVoter bool
}, error) {
	return _Oracle.Contract.VotingInfo(&_Oracle.CallOpts, arg0)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Oracle *OracleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Oracle *OracleSession) Finalize() (*types.Transaction, error) {
	return _Oracle.Contract.Finalize(&_Oracle.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_Oracle *OracleTransactorSession) Finalize() (*types.Transaction, error) {
	return _Oracle.Contract.Finalize(&_Oracle.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Oracle *OracleTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Oracle *OracleSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetOperator(&_Oracle.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Oracle *OracleTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetOperator(&_Oracle.TransactOpts, _operator)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_Oracle *OracleTransactor) SetSymbols(opts *bind.TransactOpts, _symbols []string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setSymbols", _symbols)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_Oracle *OracleSession) SetSymbols(_symbols []string) (*types.Transaction, error) {
	return _Oracle.Contract.SetSymbols(&_Oracle.TransactOpts, _symbols)
}

// SetSymbols is a paid mutator transaction binding the contract method 0x8d4f75d2.
//
// Solidity: function setSymbols(string[] _symbols) returns()
func (_Oracle *OracleTransactorSession) SetSymbols(_symbols []string) (*types.Transaction, error) {
	return _Oracle.Contract.SetSymbols(&_Oracle.TransactOpts, _symbols)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_Oracle *OracleTransactor) SetVoters(opts *bind.TransactOpts, _newVoters []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setVoters", _newVoters)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_Oracle *OracleSession) SetVoters(_newVoters []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetVoters(&_Oracle.TransactOpts, _newVoters)
}

// SetVoters is a paid mutator transaction binding the contract method 0x845023f2.
//
// Solidity: function setVoters(address[] _newVoters) returns()
func (_Oracle *OracleTransactorSession) SetVoters(_newVoters []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetVoters(&_Oracle.TransactOpts, _newVoters)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_Oracle *OracleTransactor) Vote(opts *bind.TransactOpts, _commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "vote", _commit, _reports, _salt)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_Oracle *OracleSession) Vote(_commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Vote(&_Oracle.TransactOpts, _commit, _reports, _salt)
}

// Vote is a paid mutator transaction binding the contract method 0x307de9b6.
//
// Solidity: function vote(uint256 _commit, int256[] _reports, uint256 _salt) returns()
func (_Oracle *OracleTransactorSession) Vote(_commit *big.Int, _reports []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Vote(&_Oracle.TransactOpts, _commit, _reports, _salt)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oracle *OracleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Oracle.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oracle *OracleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Fallback(&_Oracle.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Oracle *OracleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Oracle.Contract.Fallback(&_Oracle.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactorSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

// OracleNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Oracle contract.
type OracleNewRoundIterator struct {
	Event *OracleNewRound // Event containing the contract specifics and raw log

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
func (it *OracleNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewRound)
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
		it.Event = new(OracleNewRound)
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
func (it *OracleNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewRound represents a NewRound event raised by the Oracle contract.
type OracleNewRound struct {
	Round      *big.Int
	Height     *big.Int
	Timestamp  *big.Int
	VotePeriod *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_Oracle *OracleFilterer) FilterNewRound(opts *bind.FilterOpts) (*OracleNewRoundIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &OracleNewRoundIterator{contract: _Oracle.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_Oracle *OracleFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *OracleNewRound) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewRound)
				if err := _Oracle.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xb5d8636ab45e6cac7a4a61cb7c77f77f61a454d73aa2e6139ff8dcaf463537e5.
//
// Solidity: event NewRound(uint256 _round, uint256 _height, uint256 _timestamp, uint256 _votePeriod)
func (_Oracle *OracleFilterer) ParseNewRound(log types.Log) (*OracleNewRound, error) {
	event := new(OracleNewRound)
	if err := _Oracle.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleNewSymbolsIterator is returned from FilterNewSymbols and is used to iterate over the raw logs and unpacked data for NewSymbols events raised by the Oracle contract.
type OracleNewSymbolsIterator struct {
	Event *OracleNewSymbols // Event containing the contract specifics and raw log

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
func (it *OracleNewSymbolsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewSymbols)
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
		it.Event = new(OracleNewSymbols)
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
func (it *OracleNewSymbolsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewSymbolsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewSymbols represents a NewSymbols event raised by the Oracle contract.
type OracleNewSymbols struct {
	Symbols []string
	Round   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSymbols is a free log retrieval operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_Oracle *OracleFilterer) FilterNewSymbols(opts *bind.FilterOpts) (*OracleNewSymbolsIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewSymbols")
	if err != nil {
		return nil, err
	}
	return &OracleNewSymbolsIterator{contract: _Oracle.contract, event: "NewSymbols", logs: logs, sub: sub}, nil
}

// WatchNewSymbols is a free log subscription operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_Oracle *OracleFilterer) WatchNewSymbols(opts *bind.WatchOpts, sink chan<- *OracleNewSymbols) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewSymbols")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewSymbols)
				if err := _Oracle.contract.UnpackLog(event, "NewSymbols", log); err != nil {
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

// ParseNewSymbols is a log parse operation binding the contract event 0xaa278e424da680ce5dad66510415760e78e0bd87d45c786c6e88bdde82f9342d.
//
// Solidity: event NewSymbols(string[] _symbols, uint256 _round)
func (_Oracle *OracleFilterer) ParseNewSymbols(log types.Log) (*OracleNewSymbols, error) {
	event := new(OracleNewSymbols)
	if err := _Oracle.contract.UnpackLog(event, "NewSymbols", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Oracle contract.
type OracleVotedIterator struct {
	Event *OracleVoted // Event containing the contract specifics and raw log

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
func (it *OracleVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleVoted)
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
		it.Event = new(OracleVoted)
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
func (it *OracleVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleVoted represents a Voted event raised by the Oracle contract.
type OracleVoted struct {
	Voter common.Address
	Votes []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_Oracle *OracleFilterer) FilterVoted(opts *bind.FilterOpts, _voter []common.Address) (*OracleVotedIterator, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Voted", _voterRule)
	if err != nil {
		return nil, err
	}
	return &OracleVotedIterator{contract: _Oracle.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_Oracle *OracleFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *OracleVoted, _voter []common.Address) (event.Subscription, error) {

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Voted", _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleVoted)
				if err := _Oracle.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0xd0d8560f1076ac6b216b1091a2571d6f9bc3e0889f4dbdbe1c7d1be7136714d3.
//
// Solidity: event Voted(address indexed _voter, int256[] _votes)
func (_Oracle *OracleFilterer) ParseVoted(log types.Log) (*OracleVoted, error) {
	event := new(OracleVoted)
	if err := _Oracle.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrecompiledMetaData contains all meta data concerning the Precompiled contract.
var PrecompiledMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ACCUSATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INNOCENCE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MISBEHAVIOUR_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4dc925d3": "ACCUSATION_CONTRACT()",
		"8e153dc3": "INNOCENCE_CONTRACT()",
		"925c5492": "MISBEHAVIOUR_CONTRACT()",
	},
	Bin: "0x60b1610034600b8282823980515f1a60731461002857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106046575f3560e01c80634dc925d314604a5780638e153dc314606d578063925c5492146074575b5f80fd5b605160fc81565b6040516001600160a01b03909116815260200160405180910390f35b605160fd81565b605160fe8156fea2646970667358221220d86ca6dca7f5d9dd255ce7a0fb74fc56bb32863ed973a046150ff0697d465cc564736f6c63430008150033",
}

// PrecompiledABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompiledMetaData.ABI instead.
var PrecompiledABI = PrecompiledMetaData.ABI

// Deprecated: Use PrecompiledMetaData.Sigs instead.
// PrecompiledFuncSigs maps the 4-byte function signature to its string representation.
var PrecompiledFuncSigs = PrecompiledMetaData.Sigs

// PrecompiledBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompiledMetaData.Bin instead.
var PrecompiledBin = PrecompiledMetaData.Bin

// DeployPrecompiled deploys a new Ethereum contract, binding an instance of Precompiled to it.
func DeployPrecompiled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Precompiled, error) {
	parsed, err := PrecompiledMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompiledBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Precompiled{PrecompiledCaller: PrecompiledCaller{contract: contract}, PrecompiledTransactor: PrecompiledTransactor{contract: contract}, PrecompiledFilterer: PrecompiledFilterer{contract: contract}}, nil
}

// Precompiled is an auto generated Go binding around an Ethereum contract.
type Precompiled struct {
	PrecompiledCaller     // Read-only binding to the contract
	PrecompiledTransactor // Write-only binding to the contract
	PrecompiledFilterer   // Log filterer for contract events
}

// PrecompiledCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompiledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompiledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompiledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompiledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompiledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompiledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompiledSession struct {
	Contract     *Precompiled      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrecompiledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompiledCallerSession struct {
	Contract *PrecompiledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PrecompiledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompiledTransactorSession struct {
	Contract     *PrecompiledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrecompiledRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompiledRaw struct {
	Contract *Precompiled // Generic contract binding to access the raw methods on
}

// PrecompiledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompiledCallerRaw struct {
	Contract *PrecompiledCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompiledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompiledTransactorRaw struct {
	Contract *PrecompiledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompiled creates a new instance of Precompiled, bound to a specific deployed contract.
func NewPrecompiled(address common.Address, backend bind.ContractBackend) (*Precompiled, error) {
	contract, err := bindPrecompiled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Precompiled{PrecompiledCaller: PrecompiledCaller{contract: contract}, PrecompiledTransactor: PrecompiledTransactor{contract: contract}, PrecompiledFilterer: PrecompiledFilterer{contract: contract}}, nil
}

// NewPrecompiledCaller creates a new read-only instance of Precompiled, bound to a specific deployed contract.
func NewPrecompiledCaller(address common.Address, caller bind.ContractCaller) (*PrecompiledCaller, error) {
	contract, err := bindPrecompiled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompiledCaller{contract: contract}, nil
}

// NewPrecompiledTransactor creates a new write-only instance of Precompiled, bound to a specific deployed contract.
func NewPrecompiledTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompiledTransactor, error) {
	contract, err := bindPrecompiled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompiledTransactor{contract: contract}, nil
}

// NewPrecompiledFilterer creates a new log filterer instance of Precompiled, bound to a specific deployed contract.
func NewPrecompiledFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompiledFilterer, error) {
	contract, err := bindPrecompiled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompiledFilterer{contract: contract}, nil
}

// bindPrecompiled binds a generic wrapper to an already deployed contract.
func bindPrecompiled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompiledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompiled *PrecompiledRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompiled.Contract.PrecompiledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompiled *PrecompiledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompiled.Contract.PrecompiledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompiled *PrecompiledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompiled.Contract.PrecompiledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompiled *PrecompiledCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompiled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompiled *PrecompiledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompiled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompiled *PrecompiledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompiled.Contract.contract.Transact(opts, method, params...)
}

// ACCUSATIONCONTRACT is a free data retrieval call binding the contract method 0x4dc925d3.
//
// Solidity: function ACCUSATION_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCaller) ACCUSATIONCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Precompiled.contract.Call(opts, &out, "ACCUSATION_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCUSATIONCONTRACT is a free data retrieval call binding the contract method 0x4dc925d3.
//
// Solidity: function ACCUSATION_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledSession) ACCUSATIONCONTRACT() (common.Address, error) {
	return _Precompiled.Contract.ACCUSATIONCONTRACT(&_Precompiled.CallOpts)
}

// ACCUSATIONCONTRACT is a free data retrieval call binding the contract method 0x4dc925d3.
//
// Solidity: function ACCUSATION_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCallerSession) ACCUSATIONCONTRACT() (common.Address, error) {
	return _Precompiled.Contract.ACCUSATIONCONTRACT(&_Precompiled.CallOpts)
}

// INNOCENCECONTRACT is a free data retrieval call binding the contract method 0x8e153dc3.
//
// Solidity: function INNOCENCE_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCaller) INNOCENCECONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Precompiled.contract.Call(opts, &out, "INNOCENCE_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INNOCENCECONTRACT is a free data retrieval call binding the contract method 0x8e153dc3.
//
// Solidity: function INNOCENCE_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledSession) INNOCENCECONTRACT() (common.Address, error) {
	return _Precompiled.Contract.INNOCENCECONTRACT(&_Precompiled.CallOpts)
}

// INNOCENCECONTRACT is a free data retrieval call binding the contract method 0x8e153dc3.
//
// Solidity: function INNOCENCE_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCallerSession) INNOCENCECONTRACT() (common.Address, error) {
	return _Precompiled.Contract.INNOCENCECONTRACT(&_Precompiled.CallOpts)
}

// MISBEHAVIOURCONTRACT is a free data retrieval call binding the contract method 0x925c5492.
//
// Solidity: function MISBEHAVIOUR_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCaller) MISBEHAVIOURCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Precompiled.contract.Call(opts, &out, "MISBEHAVIOUR_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MISBEHAVIOURCONTRACT is a free data retrieval call binding the contract method 0x925c5492.
//
// Solidity: function MISBEHAVIOUR_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledSession) MISBEHAVIOURCONTRACT() (common.Address, error) {
	return _Precompiled.Contract.MISBEHAVIOURCONTRACT(&_Precompiled.CallOpts)
}

// MISBEHAVIOURCONTRACT is a free data retrieval call binding the contract method 0x925c5492.
//
// Solidity: function MISBEHAVIOUR_CONTRACT() view returns(address)
func (_Precompiled *PrecompiledCallerSession) MISBEHAVIOURCONTRACT() (common.Address, error) {
	return _Precompiled.Contract.MISBEHAVIOURCONTRACT(&_Precompiled.CallOpts)
}

// StabilizationMetaData contains all meta data concerning the Stabilization contract.
var StabilizationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"borrowInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minCollateralizationRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDebtRequirement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetPrice\",\"type\":\"uint256\"}],\"internalType\":\"structStabilization.Config\",\"name\":\"config_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"autonity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"supplyControl\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDebtPosition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Liquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDebtPosition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv18_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UD60x18\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"PRBMath_UD60x18_Exp2_InputTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UD60x18\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"PRBMath_UD60x18_Exp_InputTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCALE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mcr\",\"type\":\"uint256\"}],\"name\":\"borrowLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cdps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minCollateralizationRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDebtRequirement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"debtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeDue\",\"type\":\"uint256\"}],\"name\":\"interestDue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mcr\",\"type\":\"uint256\"}],\"name\":\"minimumCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setLiquidationRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setMinCollateralizationRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinDebtRequirement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"supplyControl\",\"type\":\"address\"}],\"name\":\"setSupplyControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationRatio\",\"type\":\"uint256\"}],\"name\":\"underCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"eced5526": "SCALE()",
		"ce4b5bbe": "SCALE_FACTOR()",
		"5dcc9391": "SECONDS_IN_YEAR()",
		"68cd03f6": "accounts()",
		"c5ebeaec": "borrow(uint256)",
		"83baa174": "borrowLimit(uint256,uint256,uint256,uint256)",
		"840c7e24": "cdps(address)",
		"5891de72": "collateralPrice()",
		"79502c55": "config()",
		"54a9f42c": "debtAmount(address)",
		"50bf06bf": "debtAmount(address,uint256)",
		"b6b55f25": "deposit(uint256)",
		"15184245": "interestDue(uint256,uint256,uint256,uint256)",
		"042e02cf": "isLiquidatable(address)",
		"3fdfb2aa": "isLiquidatable(address,uint256)",
		"2f865568": "liquidate(address)",
		"08796696": "minimumCollateral(uint256,uint256,uint256)",
		"402d8883": "repay()",
		"946ce8cd": "setLiquidationRatio(uint256)",
		"7b44646a": "setMinCollateralizationRatio(uint256)",
		"53afe81d": "setMinDebtRequirement(uint256)",
		"b3ab15fb": "setOperator(address)",
		"7adbf973": "setOracle(address)",
		"52e5a050": "setSupplyControl(address)",
		"fbbe6991": "underCollateralized(uint256,uint256,uint256,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x608060405234801562000010575f80fd5b50604051620027263803806200272683398101604081905262000033916200012a565b8560400151805f036200005957604051630309cb8760e51b815260040160405180910390fd5b866020015187604001518082106200008457604051630309cb8760e51b815260040160405180910390fd5b505086515f5550602086015160015560408601516002556060860151600355608090950151600455600780546001600160a01b03199081166001600160a01b039687161790915560088054821694861694909417909355600a8054841692851692909217909155600b80548316918416919091179055600980549091169190921617905562000211565b80516001600160a01b038116811462000125575f80fd5b919050565b5f805f805f8086880361014081121562000142575f80fd5b60a081121562000150575f80fd5b5060405160a081016001600160401b03811182821017156200018057634e487b7160e01b5f52604160045260245ffd5b8060405250875181526020880151602082015260408801516040820152606088015160608201526080880151608082015280965050620001c360a088016200010e565b9450620001d360c088016200010e565b9350620001e360e088016200010e565b9250620001f461010088016200010e565b91506200020561012088016200010e565b90509295509295509295565b612507806200021f5f395ff3fe60806040526004361061017b575f3560e01c806368cd03f6116100cd578063946ce8cd11610087578063c5ebeaec11610062578063c5ebeaec14610489578063ce4b5bbe146104a8578063eced5526146104bc578063fbbe6991146104d0575f80fd5b8063946ce8cd1461042c578063b3ab15fb1461044b578063b6b55f251461046a575f80fd5b806368cd03f61461030157806379502c55146103225780637adbf9731461036f5780637b44646a1461038e57806383baa174146103ad578063840c7e24146103cc575f80fd5b8063402d88831161013857806353afe81d1161011357806353afe81d1461029857806354a9f42c146102b75780635891de72146102d65780635dcc9391146102ea575f80fd5b8063402d88831461025257806350bf06bf1461025a57806352e5a05014610279575f80fd5b8063042e02cf1461017f57806308796696146101b357806315184245146101e05780632e1a7d4d146101ff5780632f865568146102205780633fdfb2aa14610233575b5f80fd5b34801561018a575f80fd5b5061019e610199366004612175565b6104ef565b60405190151581526020015b60405180910390f35b3480156101be575f80fd5b506101d26101cd36600461218e565b610563565b6040519081526020016101aa565b3480156101eb575f80fd5b506101d26101fa3660046121b7565b6105cb565b34801561020a575f80fd5b5061021e6102193660046121e6565b610657565b005b61021e61022e366004612175565b61081b565b34801561023e575f80fd5b5061019e61024d3660046121fd565b610a4d565b61021e610ad5565b348015610265575f80fd5b506101d26102743660046121fd565b610c9e565b348015610284575f80fd5b5061021e610293366004612175565b610d08565b3480156102a3575f80fd5b5061021e6102b23660046121e6565b610d54565b3480156102c2575f80fd5b506101d26102d1366004612175565b610d83565b3480156102e1575f80fd5b506101d2610df1565b3480156102f5575f80fd5b506101d26301e1338081565b34801561030c575f80fd5b50610315611085565b6040516101aa9190612225565b34801561032d575f80fd5b505f54600154600254600354600454610347949392919085565b604080519586526020860194909452928401919091526060830152608082015260a0016101aa565b34801561037a575f80fd5b5061021e610389366004612175565b6110e5565b348015610399575f80fd5b5061021e6103a83660046121e6565b611131565b3480156103b8575f80fd5b506101d26103c73660046121b7565b6111a8565b3480156103d7575f80fd5b5061040c6103e6366004612175565b60056020525f908152604090208054600182015460028301546003909301549192909184565b6040805194855260208501939093529183015260608201526080016101aa565b348015610437575f80fd5b5061021e6104463660046121e6565b6111fd565b348015610456575f80fd5b5061021e610465366004612175565b611253565b348015610475575f80fd5b5061021e6104843660046121e6565b61129f565b348015610494575f80fd5b5061021e6104a33660046121e6565b611495565b3480156104b3575f80fd5b506101d2611653565b3480156104c7575f80fd5b506101d2601281565b3480156104db575f80fd5b5061019e6104ea3660046121b7565b611662565b604051631fefd95560e11b81526001600160a01b03821660048201524260248201525f903090633fdfb2aa90604401602060405180830381865afa158015610539573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055d9190612271565b92915050565b5f82805f036105845760405162bfc92160e01b815260040160405180910390fd5b83158061058f575082155b156105ad57604051630309cb8760e51b815260040160405180910390fd5b836105b884876122a4565b6105c291906122cf565b95945050505050565b5f818311156105ed57604051630309cb8760e51b815260040160405180910390fd5b84845f61060e6301e1338061060861060589896122ee565b90565b906116b5565b90505f61062361061e84846116d3565b6116e1565b90505f61064961064261063b6106056012600a6123e1565b849061173b565b86906116d3565b9a9950505050505050505050565b80805f036106785760405163162908e360e11b815260040160405180910390fd5b335f90815260056020526040902060018101548311156106ab5760405163162908e360e11b815260040160405180910390fd5b5f6106b68242611749565b5090505f6106c2610df1565b90506106d8836001015482845f60010154611662565b156106f657604051636229415360e01b815260040160405180910390fd5b6107098360020154825f60020154610563565b85846001015461071991906122ee565b101561073857604051633a23d82560e01b815260040160405180910390fd5b84836001015f82825461074b91906122ee565b909155505060095460405163a9059cbb60e01b8152336004820152602481018790526001600160a01b039091169063a9059cbb906044016020604051808303815f875af115801561079e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107c29190612271565b6107df576040516312171d8360e31b815260040160405180910390fd5b60405185815233907f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243649060200160405180910390a25050505050565b345f0361083b57604051637c946ed760e01b815260040160405180910390fd5b6001600160a01b0381165f908152600560205260408120600281015490910361087757604051638aa5baf360e01b815260040160405180910390fd5b5f806108838342611749565b915091506108a18360010154610897610df1565b6001548590611662565b6108be57604051636ef5bcdd60e11b815260040160405180910390fd5b5f6108c983346122ee565b90506001840180544286555f9182905560028601829055600386019190915560095460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044016020604051808303815f875af1158015610936573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095a9190612271565b610977576040516312171d8360e31b815260040160405180910390fd5b600b546001600160a01b03166344df8e7061099285876122ee565b6040518263ffffffff1660e01b81526004015f604051808303818588803b1580156109bb575f80fd5b505af11580156109cd573d5f803e3d5ffd5b50505050505f821115610a0657604051339083156108fc029084905f818181858888f19350505050158015610a04573d5f803e3d5ffd5b505b6040513381526001600160a01b038716907fc3d81b2125598b9a2b024afe09e33981f0aa5b7bcbe3e30c4303a4dec209ddb4906020015b60405180910390a2505050505050565b6001600160a01b0382165f908152600560205260408120805484918491821015610a8a57604051630309cb8760e51b815260040160405180910390fd5b6001600160a01b0386165f90815260056020526040812090610aac8288611749565b509050610ac98260010154610abf610df1565b6001548490611662565b98975050505050505050565b345f03610af557604051637c946ed760e01b815260040160405180910390fd5b335f9081526005602052604081206002810154909103610b2857604051638aa5baf360e01b815260040160405180910390fd5b5f80610b348342611749565b915091508134108015610b515750600354610b4f34846122ee565b105b15610b6f5760405163e6bd447960e01b815260040160405180910390fd5b80836003015f828254610b8291906123ec565b90915550504283555f8080610b9786346117ba565b92509250925081866002015f828254610bb091906122ee565b9250508190555082866003015f828254610bca91906122ee565b90915550508115610c3a57600b5f9054906101000a90046001600160a01b03166001600160a01b03166344df8e70836040518263ffffffff1660e01b81526004015f604051808303818588803b158015610c22575f80fd5b505af1158015610c34573d5f803e3d5ffd5b50505050505b8015610c6c57604051339082156108fc029083905f818181858888f19350505050158015610c6a573d5f803e3d5ffd5b505b60405134815233907f5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a59423190602001610a3d565b6001600160a01b0382165f908152600560205260408120805484918491821015610cdb57604051630309cb8760e51b815260040160405180910390fd5b6001600160a01b0386165f908152600560205260409020610cfc8187611749565b50979650505050505050565b6008546001600160a01b03163314610d32576040516282b42960e81b815260040160405180910390fd5b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6008546001600160a01b03163314610d7e576040516282b42960e81b815260040160405180910390fd5b600355565b6040516350bf06bf60e01b81526001600160a01b03821660048201524260248201525f9030906350bf06bf90604401602060405180830381865afa158015610dcd573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055d91906123ff565b600a546040805180820182526007815266272a2717a0aa2760c91b602082015290516333f98c7760e01b81525f9283926001600160a01b03909116916333f98c7791610e3f91600401612416565b608060405180830381865afa158015610e5a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e7e9190612461565b905080606001515f14610ea45760405163cb08be8160e01b815260040160405180910390fd5b5f816020015113610ec75760405162bfc92160e01b815260040160405180910390fd5b600a5f9054906101000a90046001600160a01b03166001600160a01b0316639670c0bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f17573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f3b91906123ff565b610f476012600a6123e1565b1115610fec57600a5f9054906101000a90046001600160a01b03166001600160a01b0316639670c0bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f9d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc191906123ff565b610fcd6012600a6123e1565b610fd791906122cf565b8160200151610fe691906122a4565b91505090565b610ff86012600a6123e1565b600a5f9054906101000a90046001600160a01b03166001600160a01b0316639670c0bc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611048573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106c91906123ff565b61107691906122cf565b8160200151610fe691906122cf565b606060068054806020026020016040519081016040528092919081815260200182805480156110db57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116110bd575b5050505050905090565b6007546001600160a01b0316331461110f576040516282b42960e81b815260040160405180910390fd5b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b80805f0361115257604051630309cb8760e51b815260040160405180910390fd5b6001548280821061117657604051630309cb8760e51b815260040160405180910390fd5b6008546001600160a01b031633146111a0576040516282b42960e81b815260040160405180910390fd5b505050600255565b5f8315806111b4575081155b156111d257604051630309cb8760e51b815260040160405180910390fd5b6111de6012600a6123e1565b6111e890836122a4565b836111f386886122a4565b6105b891906122a4565b600254819080821061122257604051630309cb8760e51b815260040160405180910390fd5b6008546001600160a01b0316331461124c576040516282b42960e81b815260040160405180910390fd5b5050600155565b6007546001600160a01b0316331461127d576040516282b42960e81b815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b80805f036112c05760405163162908e360e11b815260040160405180910390fd5b600954604051636eb1769f60e11b815233600482015230602482015283916001600160a01b03169063dd62ed3e90604401602060405180830381865afa15801561130c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061133091906123ff565b101561134f576040516313be252b60e01b815260040160405180910390fd5b335f90815260056020526040812080549091036113a857600680546001810182555f919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b031916331790555b4281556001810180548491905f906113c19084906123ec565b90915550506009546040516323b872dd60e01b8152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064016020604051808303815f875af115801561141a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061143e9190612271565b61145b576040516312171d8360e31b815260040160405180910390fd5b60405183815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2505050565b80805f036114b65760405163162908e360e11b815260040160405180910390fd5b335f90815260056020526040812090806114d08342611749565b90925090506114df85836123ec565b6003549092508210156115055760405163e6bd447960e01b815260040160405180910390fd5b5f61150e610df1565b9050611524846001015482855f60010154611662565b1561154257604051636229415360e01b815260040160405180910390fd5b5f61155b8560010154835f600401545f600201546111a8565b90508084111561157e57604051633a23d82560e01b815260040160405180910390fd5b4285556002850180548891905f906115979084906123ec565b9250508190555082856003015f8282546115b191906123ec565b9091555050600b546040516340c10f1960e01b8152336004820152602481018990526001600160a01b03909116906340c10f19906044015f604051808303815f87803b1580156115ff575f80fd5b505af1158015611611573d5f803e3d5ffd5b50506040518981523392507fcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750915060200160405180910390a250505050505050565b61165f6012600a6123e1565b81565b5f83805f036116835760405162bfc92160e01b815260040160405180910390fd5b835f03611692575f91506116ac565b828461169e87896122a4565b6116a891906122cf565b1091505b50949350505050565b5f6116cc61060584670de0b6b3a76400008561182b565b9392505050565b5f6116cc61060584846118f9565b5f81680736ea4425c11ac63081111561171557604051630d7b1d6560e11b8152600481018490526024015b60405180910390fd5b6714057b7ef767814f8102611733670de0b6b3a764000082046119ab565b949350505050565b5f6116cc61060583856122ee565b5f80825f0361176b57604051630309cb8760e51b815260040160405180910390fd5b5f8460030154856002015461178091906123ec565b85549091508403611793575f91506117a6565b5f5485546117a3918391876105cb565b91505b6117b082826123ec565b9250509250929050565b5f805f80856003015486600201546117d291906123ec565b9050856003015485106117e95785600301546117eb565b845b93508085106117fe578560020154611808565b61180884866122ee565b9250808511611817575f611821565b61182181866122ee565b9150509250925092565b5f80805f19858709858702925082811083820303915050805f0361186257838281611858576118586122bb565b04925050506116cc565b83811061189357604051630c740aef60e31b815260048101879052602481018690526044810185905260640161170c565b5f84868809600260036001881981018916988990049182028318808302840302808302840302808302840302808302840302808302840302918202909203025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f80805f19848609848602925082811083820303915050805f0361192a5750670de0b6b3a76400009004905061055d565b670de0b6b3a7640000811061195c57604051635173648d60e01b8152600481018690526024810185905260440161170c565b5f670de0b6b3a764000085870962040000818503049310909103600160ee1b02919091177faccb18165bd6fe31ae1cf318dc5b51eee0e1ba569b88cd74c1773b91fac106690291505092915050565b5f81680a688906bd8affffff8111156119da5760405163b3b6ba1f60e01b81526004810184905260240161170c565b5f6119f1670de0b6b3a7640000604084901b6122cf565b905061173361060582600160bf1b67ff00000000000000821615611b0757678000000000000000821615611a2e5768016a09e667f3bcc9090260401c5b674000000000000000821615611a4d576801306fe0a31b7152df0260401c5b672000000000000000821615611a6c576801172b83c7d517adce0260401c5b671000000000000000821615611a8b5768010b5586cf9890f62a0260401c5b670800000000000000821615611aaa576801059b0d31585743ae0260401c5b670400000000000000821615611ac957680102c9a3e778060ee70260401c5b670200000000000000821615611ae85768010163da9fb33356d80260401c5b670100000000000000821615611b0757680100b1afa5abcbed610260401c5b66ff000000000000821615611c06576680000000000000821615611b345768010058c86da1c09ea20260401c5b6640000000000000821615611b52576801002c605e2e8cec500260401c5b6620000000000000821615611b7057680100162f3904051fa10260401c5b6610000000000000821615611b8e576801000b175effdc76ba0260401c5b6608000000000000821615611bac57680100058ba01fb9f96d0260401c5b6604000000000000821615611bca5768010002c5cc37da94920260401c5b6602000000000000821615611be8576801000162e525ee05470260401c5b6601000000000000821615611c065768010000b17255775c040260401c5b65ff0000000000821615611cfc5765800000000000821615611c31576801000058b91b5bc9ae0260401c5b65400000000000821615611c4e57680100002c5c89d5ec6d0260401c5b65200000000000821615611c6b5768010000162e43f4f8310260401c5b65100000000000821615611c8857680100000b1721bcfc9a0260401c5b65080000000000821615611ca55768010000058b90cf1e6e0260401c5b65040000000000821615611cc2576801000002c5c863b73f0260401c5b65020000000000821615611cdf57680100000162e430e5a20260401c5b65010000000000821615611cfc576801000000b1721835510260401c5b64ff00000000821615611de957648000000000821615611d2557680100000058b90c0b490260401c5b644000000000821615611d415768010000002c5c8601cc0260401c5b642000000000821615611d5d576801000000162e42fff00260401c5b641000000000821615611d795768010000000b17217fbb0260401c5b640800000000821615611d95576801000000058b90bfce0260401c5b640400000000821615611db157680100000002c5c85fe30260401c5b640200000000821615611dcd5768010000000162e42ff10260401c5b640100000000821615611de957680100000000b17217f80260401c5b63ff000000821615611ecd576380000000821615611e105768010000000058b90bfc0260401c5b6340000000821615611e2b576801000000002c5c85fe0260401c5b6320000000821615611e4657680100000000162e42ff0260401c5b6310000000821615611e61576801000000000b17217f0260401c5b6308000000821615611e7c57680100000000058b90c00260401c5b6304000000821615611e975768010000000002c5c8600260401c5b6302000000821615611eb2576801000000000162e4300260401c5b6301000000821615611ecd5768010000000000b172180260401c5b62ff0000821615611fa85762800000821615611ef2576801000000000058b90c0260401c5b62400000821615611f0c57680100000000002c5c860260401c5b62200000821615611f265768010000000000162e430260401c5b62100000821615611f4057680100000000000b17210260401c5b62080000821615611f5a5768010000000000058b910260401c5b62040000821615611f74576801000000000002c5c80260401c5b62020000821615611f8e57680100000000000162e40260401c5b62010000821615611fa8576801000000000000b1720260401c5b61ff0082161561207a57618000821615611fcb57680100000000000058b90260401c5b614000821615611fe45768010000000000002c5d0260401c5b612000821615611ffd576801000000000000162e0260401c5b6110008216156120165768010000000000000b170260401c5b61080082161561202f576801000000000000058c0260401c5b61040082161561204857680100000000000002c60260401c5b61020082161561206157680100000000000001630260401c5b61010082161561207a57680100000000000000b10260401c5b60ff82161561214357608082161561209b57680100000000000000590260401c5b60408216156120b3576801000000000000002c0260401c5b60208216156120cb57680100000000000000160260401c5b60108216156120e3576801000000000000000b0260401c5b60088216156120fb57680100000000000000060260401c5b600482161561211357680100000000000000030260401c5b600282161561212b57680100000000000000010260401c5b600182161561214357680100000000000000010260401c5b670de0b6b3a76400000260409190911c60bf031c90565b80356001600160a01b0381168114612170575f80fd5b919050565b5f60208284031215612185575f80fd5b6116cc8261215a565b5f805f606084860312156121a0575f80fd5b505081359360208301359350604090920135919050565b5f805f80608085870312156121ca575f80fd5b5050823594602084013594506040840135936060013592509050565b5f602082840312156121f6575f80fd5b5035919050565b5f806040838503121561220e575f80fd5b6122178361215a565b946020939093013593505050565b602080825282518282018190525f9190848201906040850190845b818110156122655783516001600160a01b031683529284019291840191600101612240565b50909695505050505050565b5f60208284031215612281575f80fd5b815180151581146116cc575f80fd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761055d5761055d612290565b634e487b7160e01b5f52601260045260245ffd5b5f826122e957634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561055d5761055d612290565b600181815b8085111561233b57815f190482111561232157612321612290565b8085161561232e57918102915b93841c9390800290612306565b509250929050565b5f826123515750600161055d565b8161235d57505f61055d565b8160018114612373576002811461237d57612399565b600191505061055d565b60ff84111561238e5761238e612290565b50506001821b61055d565b5060208310610133831016604e8410600b84101617156123bc575081810a61055d565b6123c68383612301565b805f19048211156123d9576123d9612290565b029392505050565b5f6116cc8383612343565b8082018082111561055d5761055d612290565b5f6020828403121561240f575f80fd5b5051919050565b5f6020808352835180828501525f5b8181101561244157858101830151858201604001528201612425565b505f604082860101526040601f19601f8301168501019250505092915050565b5f60808284031215612471575f80fd5b6040516080810181811067ffffffffffffffff821117156124a057634e487b7160e01b5f52604160045260245ffd5b806040525082518152602083015160208201526040830151604082015260608301516060820152809150509291505056fea2646970667358221220de93ba02def46531e72eb6c8bd6b2d9012b94eb9a91604c2ef4a942df2acd51f64736f6c63430008150033",
}

// StabilizationABI is the input ABI used to generate the binding from.
// Deprecated: Use StabilizationMetaData.ABI instead.
var StabilizationABI = StabilizationMetaData.ABI

// Deprecated: Use StabilizationMetaData.Sigs instead.
// StabilizationFuncSigs maps the 4-byte function signature to its string representation.
var StabilizationFuncSigs = StabilizationMetaData.Sigs

// StabilizationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StabilizationMetaData.Bin instead.
var StabilizationBin = StabilizationMetaData.Bin

// DeployStabilization deploys a new Ethereum contract, binding an instance of Stabilization to it.
func DeployStabilization(auth *bind.TransactOpts, backend bind.ContractBackend, config_ StabilizationConfig, autonity common.Address, operator common.Address, oracle common.Address, supplyControl common.Address, collateralToken common.Address) (common.Address, *types.Transaction, *Stabilization, error) {
	parsed, err := StabilizationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StabilizationBin), backend, config_, autonity, operator, oracle, supplyControl, collateralToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stabilization{StabilizationCaller: StabilizationCaller{contract: contract}, StabilizationTransactor: StabilizationTransactor{contract: contract}, StabilizationFilterer: StabilizationFilterer{contract: contract}}, nil
}

// Stabilization is an auto generated Go binding around an Ethereum contract.
type Stabilization struct {
	StabilizationCaller     // Read-only binding to the contract
	StabilizationTransactor // Write-only binding to the contract
	StabilizationFilterer   // Log filterer for contract events
}

// StabilizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type StabilizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StabilizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StabilizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StabilizationSession struct {
	Contract     *Stabilization    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StabilizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StabilizationCallerSession struct {
	Contract *StabilizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StabilizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StabilizationTransactorSession struct {
	Contract     *StabilizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StabilizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type StabilizationRaw struct {
	Contract *Stabilization // Generic contract binding to access the raw methods on
}

// StabilizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StabilizationCallerRaw struct {
	Contract *StabilizationCaller // Generic read-only contract binding to access the raw methods on
}

// StabilizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StabilizationTransactorRaw struct {
	Contract *StabilizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStabilization creates a new instance of Stabilization, bound to a specific deployed contract.
func NewStabilization(address common.Address, backend bind.ContractBackend) (*Stabilization, error) {
	contract, err := bindStabilization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stabilization{StabilizationCaller: StabilizationCaller{contract: contract}, StabilizationTransactor: StabilizationTransactor{contract: contract}, StabilizationFilterer: StabilizationFilterer{contract: contract}}, nil
}

// NewStabilizationCaller creates a new read-only instance of Stabilization, bound to a specific deployed contract.
func NewStabilizationCaller(address common.Address, caller bind.ContractCaller) (*StabilizationCaller, error) {
	contract, err := bindStabilization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizationCaller{contract: contract}, nil
}

// NewStabilizationTransactor creates a new write-only instance of Stabilization, bound to a specific deployed contract.
func NewStabilizationTransactor(address common.Address, transactor bind.ContractTransactor) (*StabilizationTransactor, error) {
	contract, err := bindStabilization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizationTransactor{contract: contract}, nil
}

// NewStabilizationFilterer creates a new log filterer instance of Stabilization, bound to a specific deployed contract.
func NewStabilizationFilterer(address common.Address, filterer bind.ContractFilterer) (*StabilizationFilterer, error) {
	contract, err := bindStabilization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StabilizationFilterer{contract: contract}, nil
}

// bindStabilization binds a generic wrapper to an already deployed contract.
func bindStabilization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StabilizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stabilization *StabilizationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stabilization.Contract.StabilizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stabilization *StabilizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stabilization.Contract.StabilizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stabilization *StabilizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stabilization.Contract.StabilizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stabilization *StabilizationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stabilization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stabilization *StabilizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stabilization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stabilization *StabilizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stabilization.Contract.contract.Transact(opts, method, params...)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Stabilization *StabilizationCaller) SCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Stabilization *StabilizationSession) SCALE() (*big.Int, error) {
	return _Stabilization.Contract.SCALE(&_Stabilization.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Stabilization *StabilizationCallerSession) SCALE() (*big.Int, error) {
	return _Stabilization.Contract.SCALE(&_Stabilization.CallOpts)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Stabilization *StabilizationCaller) SCALEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "SCALE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Stabilization *StabilizationSession) SCALEFACTOR() (*big.Int, error) {
	return _Stabilization.Contract.SCALEFACTOR(&_Stabilization.CallOpts)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Stabilization *StabilizationCallerSession) SCALEFACTOR() (*big.Int, error) {
	return _Stabilization.Contract.SCALEFACTOR(&_Stabilization.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_Stabilization *StabilizationCaller) SECONDSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "SECONDS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_Stabilization *StabilizationSession) SECONDSINYEAR() (*big.Int, error) {
	return _Stabilization.Contract.SECONDSINYEAR(&_Stabilization.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_Stabilization *StabilizationCallerSession) SECONDSINYEAR() (*big.Int, error) {
	return _Stabilization.Contract.SECONDSINYEAR(&_Stabilization.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address[])
func (_Stabilization *StabilizationCaller) Accounts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "accounts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address[])
func (_Stabilization *StabilizationSession) Accounts() ([]common.Address, error) {
	return _Stabilization.Contract.Accounts(&_Stabilization.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address[])
func (_Stabilization *StabilizationCallerSession) Accounts() ([]common.Address, error) {
	return _Stabilization.Contract.Accounts(&_Stabilization.CallOpts)
}

// BorrowLimit is a free data retrieval call binding the contract method 0x83baa174.
//
// Solidity: function borrowLimit(uint256 collateral, uint256 price, uint256 targetPrice, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationCaller) BorrowLimit(opts *bind.CallOpts, collateral *big.Int, price *big.Int, targetPrice *big.Int, mcr *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "borrowLimit", collateral, price, targetPrice, mcr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowLimit is a free data retrieval call binding the contract method 0x83baa174.
//
// Solidity: function borrowLimit(uint256 collateral, uint256 price, uint256 targetPrice, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationSession) BorrowLimit(collateral *big.Int, price *big.Int, targetPrice *big.Int, mcr *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.BorrowLimit(&_Stabilization.CallOpts, collateral, price, targetPrice, mcr)
}

// BorrowLimit is a free data retrieval call binding the contract method 0x83baa174.
//
// Solidity: function borrowLimit(uint256 collateral, uint256 price, uint256 targetPrice, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationCallerSession) BorrowLimit(collateral *big.Int, price *big.Int, targetPrice *big.Int, mcr *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.BorrowLimit(&_Stabilization.CallOpts, collateral, price, targetPrice, mcr)
}

// Cdps is a free data retrieval call binding the contract method 0x840c7e24.
//
// Solidity: function cdps(address ) view returns(uint256 timestamp, uint256 collateral, uint256 principal, uint256 interest)
func (_Stabilization *StabilizationCaller) Cdps(opts *bind.CallOpts, arg0 common.Address) (struct {
	Timestamp  *big.Int
	Collateral *big.Int
	Principal  *big.Int
	Interest   *big.Int
}, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "cdps", arg0)

	outstruct := new(struct {
		Timestamp  *big.Int
		Collateral *big.Int
		Principal  *big.Int
		Interest   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Principal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Interest = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Cdps is a free data retrieval call binding the contract method 0x840c7e24.
//
// Solidity: function cdps(address ) view returns(uint256 timestamp, uint256 collateral, uint256 principal, uint256 interest)
func (_Stabilization *StabilizationSession) Cdps(arg0 common.Address) (struct {
	Timestamp  *big.Int
	Collateral *big.Int
	Principal  *big.Int
	Interest   *big.Int
}, error) {
	return _Stabilization.Contract.Cdps(&_Stabilization.CallOpts, arg0)
}

// Cdps is a free data retrieval call binding the contract method 0x840c7e24.
//
// Solidity: function cdps(address ) view returns(uint256 timestamp, uint256 collateral, uint256 principal, uint256 interest)
func (_Stabilization *StabilizationCallerSession) Cdps(arg0 common.Address) (struct {
	Timestamp  *big.Int
	Collateral *big.Int
	Principal  *big.Int
	Interest   *big.Int
}, error) {
	return _Stabilization.Contract.Cdps(&_Stabilization.CallOpts, arg0)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() view returns(uint256 price)
func (_Stabilization *StabilizationCaller) CollateralPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "collateralPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() view returns(uint256 price)
func (_Stabilization *StabilizationSession) CollateralPrice() (*big.Int, error) {
	return _Stabilization.Contract.CollateralPrice(&_Stabilization.CallOpts)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x5891de72.
//
// Solidity: function collateralPrice() view returns(uint256 price)
func (_Stabilization *StabilizationCallerSession) CollateralPrice() (*big.Int, error) {
	return _Stabilization.Contract.CollateralPrice(&_Stabilization.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 borrowInterestRate, uint256 liquidationRatio, uint256 minCollateralizationRatio, uint256 minDebtRequirement, uint256 targetPrice)
func (_Stabilization *StabilizationCaller) Config(opts *bind.CallOpts) (struct {
	BorrowInterestRate        *big.Int
	LiquidationRatio          *big.Int
	MinCollateralizationRatio *big.Int
	MinDebtRequirement        *big.Int
	TargetPrice               *big.Int
}, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		BorrowInterestRate        *big.Int
		LiquidationRatio          *big.Int
		MinCollateralizationRatio *big.Int
		MinDebtRequirement        *big.Int
		TargetPrice               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowInterestRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidationRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinCollateralizationRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinDebtRequirement = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TargetPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 borrowInterestRate, uint256 liquidationRatio, uint256 minCollateralizationRatio, uint256 minDebtRequirement, uint256 targetPrice)
func (_Stabilization *StabilizationSession) Config() (struct {
	BorrowInterestRate        *big.Int
	LiquidationRatio          *big.Int
	MinCollateralizationRatio *big.Int
	MinDebtRequirement        *big.Int
	TargetPrice               *big.Int
}, error) {
	return _Stabilization.Contract.Config(&_Stabilization.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 borrowInterestRate, uint256 liquidationRatio, uint256 minCollateralizationRatio, uint256 minDebtRequirement, uint256 targetPrice)
func (_Stabilization *StabilizationCallerSession) Config() (struct {
	BorrowInterestRate        *big.Int
	LiquidationRatio          *big.Int
	MinCollateralizationRatio *big.Int
	MinDebtRequirement        *big.Int
	TargetPrice               *big.Int
}, error) {
	return _Stabilization.Contract.Config(&_Stabilization.CallOpts)
}

// DebtAmount is a free data retrieval call binding the contract method 0x50bf06bf.
//
// Solidity: function debtAmount(address account, uint256 timestamp) view returns(uint256 debt)
func (_Stabilization *StabilizationCaller) DebtAmount(opts *bind.CallOpts, account common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "debtAmount", account, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtAmount is a free data retrieval call binding the contract method 0x50bf06bf.
//
// Solidity: function debtAmount(address account, uint256 timestamp) view returns(uint256 debt)
func (_Stabilization *StabilizationSession) DebtAmount(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.DebtAmount(&_Stabilization.CallOpts, account, timestamp)
}

// DebtAmount is a free data retrieval call binding the contract method 0x50bf06bf.
//
// Solidity: function debtAmount(address account, uint256 timestamp) view returns(uint256 debt)
func (_Stabilization *StabilizationCallerSession) DebtAmount(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.DebtAmount(&_Stabilization.CallOpts, account, timestamp)
}

// DebtAmount0 is a free data retrieval call binding the contract method 0x54a9f42c.
//
// Solidity: function debtAmount(address account) view returns(uint256 debt)
func (_Stabilization *StabilizationCaller) DebtAmount0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "debtAmount0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtAmount0 is a free data retrieval call binding the contract method 0x54a9f42c.
//
// Solidity: function debtAmount(address account) view returns(uint256 debt)
func (_Stabilization *StabilizationSession) DebtAmount0(account common.Address) (*big.Int, error) {
	return _Stabilization.Contract.DebtAmount0(&_Stabilization.CallOpts, account)
}

// DebtAmount0 is a free data retrieval call binding the contract method 0x54a9f42c.
//
// Solidity: function debtAmount(address account) view returns(uint256 debt)
func (_Stabilization *StabilizationCallerSession) DebtAmount0(account common.Address) (*big.Int, error) {
	return _Stabilization.Contract.DebtAmount0(&_Stabilization.CallOpts, account)
}

// InterestDue is a free data retrieval call binding the contract method 0x15184245.
//
// Solidity: function interestDue(uint256 debt, uint256 rate, uint256 timeBorrow, uint256 timeDue) pure returns(uint256)
func (_Stabilization *StabilizationCaller) InterestDue(opts *bind.CallOpts, debt *big.Int, rate *big.Int, timeBorrow *big.Int, timeDue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "interestDue", debt, rate, timeBorrow, timeDue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestDue is a free data retrieval call binding the contract method 0x15184245.
//
// Solidity: function interestDue(uint256 debt, uint256 rate, uint256 timeBorrow, uint256 timeDue) pure returns(uint256)
func (_Stabilization *StabilizationSession) InterestDue(debt *big.Int, rate *big.Int, timeBorrow *big.Int, timeDue *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.InterestDue(&_Stabilization.CallOpts, debt, rate, timeBorrow, timeDue)
}

// InterestDue is a free data retrieval call binding the contract method 0x15184245.
//
// Solidity: function interestDue(uint256 debt, uint256 rate, uint256 timeBorrow, uint256 timeDue) pure returns(uint256)
func (_Stabilization *StabilizationCallerSession) InterestDue(debt *big.Int, rate *big.Int, timeBorrow *big.Int, timeDue *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.InterestDue(&_Stabilization.CallOpts, debt, rate, timeBorrow, timeDue)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Stabilization *StabilizationCaller) IsLiquidatable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "isLiquidatable", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Stabilization *StabilizationSession) IsLiquidatable(account common.Address) (bool, error) {
	return _Stabilization.Contract.IsLiquidatable(&_Stabilization.CallOpts, account)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Stabilization *StabilizationCallerSession) IsLiquidatable(account common.Address) (bool, error) {
	return _Stabilization.Contract.IsLiquidatable(&_Stabilization.CallOpts, account)
}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0x3fdfb2aa.
//
// Solidity: function isLiquidatable(address account, uint256 timestamp) view returns(bool)
func (_Stabilization *StabilizationCaller) IsLiquidatable0(opts *bind.CallOpts, account common.Address, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "isLiquidatable0", account, timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0x3fdfb2aa.
//
// Solidity: function isLiquidatable(address account, uint256 timestamp) view returns(bool)
func (_Stabilization *StabilizationSession) IsLiquidatable0(account common.Address, timestamp *big.Int) (bool, error) {
	return _Stabilization.Contract.IsLiquidatable0(&_Stabilization.CallOpts, account, timestamp)
}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0x3fdfb2aa.
//
// Solidity: function isLiquidatable(address account, uint256 timestamp) view returns(bool)
func (_Stabilization *StabilizationCallerSession) IsLiquidatable0(account common.Address, timestamp *big.Int) (bool, error) {
	return _Stabilization.Contract.IsLiquidatable0(&_Stabilization.CallOpts, account, timestamp)
}

// MinimumCollateral is a free data retrieval call binding the contract method 0x08796696.
//
// Solidity: function minimumCollateral(uint256 principal, uint256 price, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationCaller) MinimumCollateral(opts *bind.CallOpts, principal *big.Int, price *big.Int, mcr *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "minimumCollateral", principal, price, mcr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumCollateral is a free data retrieval call binding the contract method 0x08796696.
//
// Solidity: function minimumCollateral(uint256 principal, uint256 price, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationSession) MinimumCollateral(principal *big.Int, price *big.Int, mcr *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.MinimumCollateral(&_Stabilization.CallOpts, principal, price, mcr)
}

// MinimumCollateral is a free data retrieval call binding the contract method 0x08796696.
//
// Solidity: function minimumCollateral(uint256 principal, uint256 price, uint256 mcr) pure returns(uint256)
func (_Stabilization *StabilizationCallerSession) MinimumCollateral(principal *big.Int, price *big.Int, mcr *big.Int) (*big.Int, error) {
	return _Stabilization.Contract.MinimumCollateral(&_Stabilization.CallOpts, principal, price, mcr)
}

// UnderCollateralized is a free data retrieval call binding the contract method 0xfbbe6991.
//
// Solidity: function underCollateralized(uint256 collateral, uint256 price, uint256 debt, uint256 liquidationRatio) pure returns(bool)
func (_Stabilization *StabilizationCaller) UnderCollateralized(opts *bind.CallOpts, collateral *big.Int, price *big.Int, debt *big.Int, liquidationRatio *big.Int) (bool, error) {
	var out []interface{}
	err := _Stabilization.contract.Call(opts, &out, "underCollateralized", collateral, price, debt, liquidationRatio)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnderCollateralized is a free data retrieval call binding the contract method 0xfbbe6991.
//
// Solidity: function underCollateralized(uint256 collateral, uint256 price, uint256 debt, uint256 liquidationRatio) pure returns(bool)
func (_Stabilization *StabilizationSession) UnderCollateralized(collateral *big.Int, price *big.Int, debt *big.Int, liquidationRatio *big.Int) (bool, error) {
	return _Stabilization.Contract.UnderCollateralized(&_Stabilization.CallOpts, collateral, price, debt, liquidationRatio)
}

// UnderCollateralized is a free data retrieval call binding the contract method 0xfbbe6991.
//
// Solidity: function underCollateralized(uint256 collateral, uint256 price, uint256 debt, uint256 liquidationRatio) pure returns(bool)
func (_Stabilization *StabilizationCallerSession) UnderCollateralized(collateral *big.Int, price *big.Int, debt *big.Int, liquidationRatio *big.Int) (bool, error) {
	return _Stabilization.Contract.UnderCollateralized(&_Stabilization.CallOpts, collateral, price, debt, liquidationRatio)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_Stabilization *StabilizationTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "borrow", amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_Stabilization *StabilizationSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Borrow(&_Stabilization.TransactOpts, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_Stabilization *StabilizationTransactorSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Borrow(&_Stabilization.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stabilization *StabilizationTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stabilization *StabilizationSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Deposit(&_Stabilization.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stabilization *StabilizationTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Deposit(&_Stabilization.TransactOpts, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address account) payable returns()
func (_Stabilization *StabilizationTransactor) Liquidate(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "liquidate", account)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address account) payable returns()
func (_Stabilization *StabilizationSession) Liquidate(account common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.Liquidate(&_Stabilization.TransactOpts, account)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address account) payable returns()
func (_Stabilization *StabilizationTransactorSession) Liquidate(account common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.Liquidate(&_Stabilization.TransactOpts, account)
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_Stabilization *StabilizationTransactor) Repay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "repay")
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_Stabilization *StabilizationSession) Repay() (*types.Transaction, error) {
	return _Stabilization.Contract.Repay(&_Stabilization.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_Stabilization *StabilizationTransactorSession) Repay() (*types.Transaction, error) {
	return _Stabilization.Contract.Repay(&_Stabilization.TransactOpts)
}

// SetLiquidationRatio is a paid mutator transaction binding the contract method 0x946ce8cd.
//
// Solidity: function setLiquidationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationTransactor) SetLiquidationRatio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setLiquidationRatio", ratio)
}

// SetLiquidationRatio is a paid mutator transaction binding the contract method 0x946ce8cd.
//
// Solidity: function setLiquidationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationSession) SetLiquidationRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetLiquidationRatio(&_Stabilization.TransactOpts, ratio)
}

// SetLiquidationRatio is a paid mutator transaction binding the contract method 0x946ce8cd.
//
// Solidity: function setLiquidationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationTransactorSession) SetLiquidationRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetLiquidationRatio(&_Stabilization.TransactOpts, ratio)
}

// SetMinCollateralizationRatio is a paid mutator transaction binding the contract method 0x7b44646a.
//
// Solidity: function setMinCollateralizationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationTransactor) SetMinCollateralizationRatio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setMinCollateralizationRatio", ratio)
}

// SetMinCollateralizationRatio is a paid mutator transaction binding the contract method 0x7b44646a.
//
// Solidity: function setMinCollateralizationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationSession) SetMinCollateralizationRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetMinCollateralizationRatio(&_Stabilization.TransactOpts, ratio)
}

// SetMinCollateralizationRatio is a paid mutator transaction binding the contract method 0x7b44646a.
//
// Solidity: function setMinCollateralizationRatio(uint256 ratio) returns()
func (_Stabilization *StabilizationTransactorSession) SetMinCollateralizationRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetMinCollateralizationRatio(&_Stabilization.TransactOpts, ratio)
}

// SetMinDebtRequirement is a paid mutator transaction binding the contract method 0x53afe81d.
//
// Solidity: function setMinDebtRequirement(uint256 amount) returns()
func (_Stabilization *StabilizationTransactor) SetMinDebtRequirement(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setMinDebtRequirement", amount)
}

// SetMinDebtRequirement is a paid mutator transaction binding the contract method 0x53afe81d.
//
// Solidity: function setMinDebtRequirement(uint256 amount) returns()
func (_Stabilization *StabilizationSession) SetMinDebtRequirement(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetMinDebtRequirement(&_Stabilization.TransactOpts, amount)
}

// SetMinDebtRequirement is a paid mutator transaction binding the contract method 0x53afe81d.
//
// Solidity: function setMinDebtRequirement(uint256 amount) returns()
func (_Stabilization *StabilizationTransactorSession) SetMinDebtRequirement(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.SetMinDebtRequirement(&_Stabilization.TransactOpts, amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Stabilization *StabilizationTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Stabilization *StabilizationSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetOperator(&_Stabilization.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_Stabilization *StabilizationTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetOperator(&_Stabilization.TransactOpts, operator)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Stabilization *StabilizationTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Stabilization *StabilizationSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetOracle(&_Stabilization.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Stabilization *StabilizationTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetOracle(&_Stabilization.TransactOpts, oracle)
}

// SetSupplyControl is a paid mutator transaction binding the contract method 0x52e5a050.
//
// Solidity: function setSupplyControl(address supplyControl) returns()
func (_Stabilization *StabilizationTransactor) SetSupplyControl(opts *bind.TransactOpts, supplyControl common.Address) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "setSupplyControl", supplyControl)
}

// SetSupplyControl is a paid mutator transaction binding the contract method 0x52e5a050.
//
// Solidity: function setSupplyControl(address supplyControl) returns()
func (_Stabilization *StabilizationSession) SetSupplyControl(supplyControl common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetSupplyControl(&_Stabilization.TransactOpts, supplyControl)
}

// SetSupplyControl is a paid mutator transaction binding the contract method 0x52e5a050.
//
// Solidity: function setSupplyControl(address supplyControl) returns()
func (_Stabilization *StabilizationTransactorSession) SetSupplyControl(supplyControl common.Address) (*types.Transaction, error) {
	return _Stabilization.Contract.SetSupplyControl(&_Stabilization.TransactOpts, supplyControl)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stabilization *StabilizationTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stabilization *StabilizationSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Withdraw(&_Stabilization.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stabilization *StabilizationTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stabilization.Contract.Withdraw(&_Stabilization.TransactOpts, amount)
}

// StabilizationBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Stabilization contract.
type StabilizationBorrowIterator struct {
	Event *StabilizationBorrow // Event containing the contract specifics and raw log

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
func (it *StabilizationBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilizationBorrow)
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
		it.Event = new(StabilizationBorrow)
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
func (it *StabilizationBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilizationBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilizationBorrow represents a Borrow event raised by the Stabilization contract.
type StabilizationBorrow struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*StabilizationBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &StabilizationBorrowIterator{contract: _Stabilization.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *StabilizationBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilizationBorrow)
				if err := _Stabilization.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) ParseBorrow(log types.Log) (*StabilizationBorrow, error) {
	event := new(StabilizationBorrow)
	if err := _Stabilization.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StabilizationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Stabilization contract.
type StabilizationDepositIterator struct {
	Event *StabilizationDeposit // Event containing the contract specifics and raw log

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
func (it *StabilizationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilizationDeposit)
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
		it.Event = new(StabilizationDeposit)
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
func (it *StabilizationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilizationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilizationDeposit represents a Deposit event raised by the Stabilization contract.
type StabilizationDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*StabilizationDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &StabilizationDepositIterator{contract: _Stabilization.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StabilizationDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilizationDeposit)
				if err := _Stabilization.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) ParseDeposit(log types.Log) (*StabilizationDeposit, error) {
	event := new(StabilizationDeposit)
	if err := _Stabilization.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StabilizationLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Stabilization contract.
type StabilizationLiquidateIterator struct {
	Event *StabilizationLiquidate // Event containing the contract specifics and raw log

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
func (it *StabilizationLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilizationLiquidate)
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
		it.Event = new(StabilizationLiquidate)
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
func (it *StabilizationLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilizationLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilizationLiquidate represents a Liquidate event raised by the Stabilization contract.
type StabilizationLiquidate struct {
	Account    common.Address
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xc3d81b2125598b9a2b024afe09e33981f0aa5b7bcbe3e30c4303a4dec209ddb4.
//
// Solidity: event Liquidate(address indexed account, address liquidator)
func (_Stabilization *StabilizationFilterer) FilterLiquidate(opts *bind.FilterOpts, account []common.Address) (*StabilizationLiquidateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.FilterLogs(opts, "Liquidate", accountRule)
	if err != nil {
		return nil, err
	}
	return &StabilizationLiquidateIterator{contract: _Stabilization.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xc3d81b2125598b9a2b024afe09e33981f0aa5b7bcbe3e30c4303a4dec209ddb4.
//
// Solidity: event Liquidate(address indexed account, address liquidator)
func (_Stabilization *StabilizationFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *StabilizationLiquidate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.WatchLogs(opts, "Liquidate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilizationLiquidate)
				if err := _Stabilization.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xc3d81b2125598b9a2b024afe09e33981f0aa5b7bcbe3e30c4303a4dec209ddb4.
//
// Solidity: event Liquidate(address indexed account, address liquidator)
func (_Stabilization *StabilizationFilterer) ParseLiquidate(log types.Log) (*StabilizationLiquidate, error) {
	event := new(StabilizationLiquidate)
	if err := _Stabilization.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StabilizationRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Stabilization contract.
type StabilizationRepayIterator struct {
	Event *StabilizationRepay // Event containing the contract specifics and raw log

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
func (it *StabilizationRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilizationRepay)
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
		it.Event = new(StabilizationRepay)
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
func (it *StabilizationRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilizationRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilizationRepay represents a Repay event raised by the Stabilization contract.
type StabilizationRepay struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*StabilizationRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &StabilizationRepayIterator{contract: _Stabilization.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *StabilizationRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilizationRepay)
				if err := _Stabilization.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) ParseRepay(log types.Log) (*StabilizationRepay, error) {
	event := new(StabilizationRepay)
	if err := _Stabilization.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StabilizationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Stabilization contract.
type StabilizationWithdrawIterator struct {
	Event *StabilizationWithdraw // Event containing the contract specifics and raw log

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
func (it *StabilizationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StabilizationWithdraw)
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
		it.Event = new(StabilizationWithdraw)
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
func (it *StabilizationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StabilizationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StabilizationWithdraw represents a Withdraw event raised by the Stabilization contract.
type StabilizationWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*StabilizationWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &StabilizationWithdrawIterator{contract: _Stabilization.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StabilizationWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stabilization.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StabilizationWithdraw)
				if err := _Stabilization.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_Stabilization *StabilizationFilterer) ParseWithdraw(log types.Log) (*StabilizationWithdraw, error) {
	event := new(StabilizationWithdraw)
	if err := _Stabilization.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableMetaData contains all meta data concerning the Upgradeable contract.
var UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"completeContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewContract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytecode\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"872cf059": "completeContractUpgrade()",
		"b66b3e79": "getNewContract()",
		"cf9c5719": "resetContractUpgrade()",
		"b2ea9adb": "upgradeContract(bytes,string)",
	},
}

// UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeableMetaData.ABI instead.
var UpgradeableABI = UpgradeableMetaData.ABI

// Deprecated: Use UpgradeableMetaData.Sigs instead.
// UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var UpgradeableFuncSigs = UpgradeableMetaData.Sigs

// Upgradeable is an auto generated Go binding around an Ethereum contract.
type Upgradeable struct {
	UpgradeableCaller     // Read-only binding to the contract
	UpgradeableTransactor // Write-only binding to the contract
	UpgradeableFilterer   // Log filterer for contract events
}

// UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableSession struct {
	Contract     *Upgradeable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableCallerSession struct {
	Contract *UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableTransactorSession struct {
	Contract     *UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableRaw struct {
	Contract *Upgradeable // Generic contract binding to access the raw methods on
}

// UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableCallerRaw struct {
	Contract *UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableTransactorRaw struct {
	Contract *UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeable creates a new instance of Upgradeable, bound to a specific deployed contract.
func NewUpgradeable(address common.Address, backend bind.ContractBackend) (*Upgradeable, error) {
	contract, err := bindUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Upgradeable{UpgradeableCaller: UpgradeableCaller{contract: contract}, UpgradeableTransactor: UpgradeableTransactor{contract: contract}, UpgradeableFilterer: UpgradeableFilterer{contract: contract}}, nil
}

// NewUpgradeableCaller creates a new read-only instance of Upgradeable, bound to a specific deployed contract.
func NewUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableCaller, error) {
	contract, err := bindUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableCaller{contract: contract}, nil
}

// NewUpgradeableTransactor creates a new write-only instance of Upgradeable, bound to a specific deployed contract.
func NewUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableTransactor, error) {
	contract, err := bindUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTransactor{contract: contract}, nil
}

// NewUpgradeableFilterer creates a new log filterer instance of Upgradeable, bound to a specific deployed contract.
func NewUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableFilterer, error) {
	contract, err := bindUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableFilterer{contract: contract}, nil
}

// bindUpgradeable binds a generic wrapper to an already deployed contract.
func bindUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgradeable *UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgradeable.Contract.UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgradeable *UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgradeable.Contract.UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgradeable *UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgradeable.Contract.UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgradeable *UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgradeable *UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgradeable *UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Upgradeable *UpgradeableCaller) GetNewContract(opts *bind.CallOpts) ([]byte, string, error) {
	var out []interface{}
	err := _Upgradeable.contract.Call(opts, &out, "getNewContract")

	if err != nil {
		return *new([]byte), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Upgradeable *UpgradeableSession) GetNewContract() ([]byte, string, error) {
	return _Upgradeable.Contract.GetNewContract(&_Upgradeable.CallOpts)
}

// GetNewContract is a free data retrieval call binding the contract method 0xb66b3e79.
//
// Solidity: function getNewContract() view returns(bytes, string)
func (_Upgradeable *UpgradeableCallerSession) GetNewContract() ([]byte, string, error) {
	return _Upgradeable.Contract.GetNewContract(&_Upgradeable.CallOpts)
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Upgradeable *UpgradeableTransactor) CompleteContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgradeable.contract.Transact(opts, "completeContractUpgrade")
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Upgradeable *UpgradeableSession) CompleteContractUpgrade() (*types.Transaction, error) {
	return _Upgradeable.Contract.CompleteContractUpgrade(&_Upgradeable.TransactOpts)
}

// CompleteContractUpgrade is a paid mutator transaction binding the contract method 0x872cf059.
//
// Solidity: function completeContractUpgrade() returns()
func (_Upgradeable *UpgradeableTransactorSession) CompleteContractUpgrade() (*types.Transaction, error) {
	return _Upgradeable.Contract.CompleteContractUpgrade(&_Upgradeable.TransactOpts)
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Upgradeable *UpgradeableTransactor) ResetContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgradeable.contract.Transact(opts, "resetContractUpgrade")
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Upgradeable *UpgradeableSession) ResetContractUpgrade() (*types.Transaction, error) {
	return _Upgradeable.Contract.ResetContractUpgrade(&_Upgradeable.TransactOpts)
}

// ResetContractUpgrade is a paid mutator transaction binding the contract method 0xcf9c5719.
//
// Solidity: function resetContractUpgrade() returns()
func (_Upgradeable *UpgradeableTransactorSession) ResetContractUpgrade() (*types.Transaction, error) {
	return _Upgradeable.Contract.ResetContractUpgrade(&_Upgradeable.TransactOpts)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Upgradeable *UpgradeableTransactor) UpgradeContract(opts *bind.TransactOpts, _bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Upgradeable.contract.Transact(opts, "upgradeContract", _bytecode, _abi)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Upgradeable *UpgradeableSession) UpgradeContract(_bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Upgradeable.Contract.UpgradeContract(&_Upgradeable.TransactOpts, _bytecode, _abi)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xb2ea9adb.
//
// Solidity: function upgradeContract(bytes _bytecode, string _abi) returns()
func (_Upgradeable *UpgradeableTransactorSession) UpgradeContract(_bytecode []byte, _abi string) (*types.Transaction, error) {
	return _Upgradeable.Contract.UpgradeContract(&_Upgradeable.TransactOpts, _bytecode, _abi)
}