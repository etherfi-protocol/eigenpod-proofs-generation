// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etherfiNodesManager

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
	_ = abi.ConvertType
)

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// IEtherFiNodesManagerValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEtherFiNodesManagerValidatorInfo struct {
	ValidatorIndex       uint32
	ExitRequestTimestamp uint32
	ExitTimestamp        uint32
	Phase                uint8
}

// EtherfiNodesManagerMetaData contains all meta data concerning the EtherfiNodesManager contract.
var EtherfiNodesManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInstalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwardedCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEtherFiNodeVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidForwardedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPenaltyRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInstalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStakingManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendFail\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"AllowedForwardedEigenpodCallsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"AllowedForwardedExternalCallsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"etherFiNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toOperator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTnft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toBnft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTreasury\",\"type\":\"uint256\"}],\"name\":\"FullWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"NodeEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"NodeExitProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"NodeExitRequestReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"NodeExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"etherFiNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toOperator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTnft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toBnft\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTreasury\",\"type\":\"uint256\"}],\"name\":\"PartialWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\",\"name\":\"_phase\",\"type\":\"uint8\"}],\"name\":\"PhaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"etherFiNode\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\"}],\"name\":\"QueuedRestakingWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPRECATED_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPRECATED_enableNodeRecycling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPRECATED_protocolRevenueManager\",\"outputs\":[{\"internalType\":\"contractIProtocolRevenueManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPRECATED_protocolRevenueManagerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPRECATED_protocolRewardsSplit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"treasury\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodeOperator\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tnft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"bnft\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCALE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enableRestaking\",\"type\":\"bool\"}],\"name\":\"allocateEtherFiNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"withdrawalSafeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"allowedForwardedEigenpodCalls\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedForwardedExternalCalls\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionManager\",\"outputs\":[{\"internalType\":\"contractIAuctionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchFullWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchPartialWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchQueueRestakedWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchSendExitRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnft\",\"outputs\":[{\"internalType\":\"contractBNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beaconBalance\",\"type\":\"uint256\"}],\"name\":\"calculateTVL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toNodeOperator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTnft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBnft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTreasury\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"name\":\"withdrawals\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_receiveAsTokens\",\"type\":\"bool\"}],\"name\":\"completeQueuedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedWithdrawalRouter\",\"outputs\":[{\"internalType\":\"contractIDelayedWithdrawalRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"etherfiNodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_data\",\"type\":\"bytes[]\"}],\"name\":\"forwardEigenpodCall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_etherfiNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_data\",\"type\":\"bytes[]\"}],\"name\":\"forwardEigenpodCall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_etherfiNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_data\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"forwardExternalCall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_data\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"forwardExternalCall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"fullWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"generateWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getEigenPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalPayouts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"toNodeOperator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTnft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toBnft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTreasury\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getNonExitPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonExitPenalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getRewardsPayouts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnusedWithdrawalSafesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitRequestTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\",\"name\":\"phase\",\"type\":\"uint8\"}],\"internalType\":\"structIEtherFiNodesManager.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_count\",\"type\":\"uint64\"}],\"name\":\"incrementNumberOfValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"isExitRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"markBeingSlashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEigenlayerWithdrawals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonExitPenaltyDailyRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonExitPenaltyPrincipal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"numAssociatedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfValidators\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatingAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"partialWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\",\"name\":\"validatorPhase\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_exitTimestamps\",\"type\":\"uint32[]\"}],\"name\":\"processNodeExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enableRestaking\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_withdrawalSafeAddress\",\"type\":\"address\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_max\",\"type\":\"uint8\"}],\"name\":\"setMaxEigenLayerWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonExitPenaltyDailyRate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonExitPenaltyPrincipal\",\"type\":\"uint64\"}],\"name\":\"setNonExitPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newProofSubmitter\",\"type\":\"address\"}],\"name\":\"setProofSubmitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_treasury\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nodeOperator\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_tnft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_bnft\",\"type\":\"uint64\"}],\"name\":\"setStakingRewardsSplit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\",\"name\":\"_phase\",\"type\":\"uint8\"}],\"name\":\"setValidatorPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingManagerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingRewardsSplit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"treasury\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodeOperator\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tnft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"bnft\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_revertIfNoBalance\",\"type\":\"bool\"}],\"name\":\"startCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tnft\",\"outputs\":[{\"internalType\":\"contractTNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"unregisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unusedWithdrawalSafes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedForwardedEigenpodCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedForwardedExternalCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateEigenLayerOperatingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"updateEtherFiNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EtherfiNodesManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiNodesManagerMetaData.ABI instead.
var EtherfiNodesManagerABI = EtherfiNodesManagerMetaData.ABI

// EtherfiNodesManager is an auto generated Go binding around an Ethereum contract.
type EtherfiNodesManager struct {
	EtherfiNodesManagerCaller     // Read-only binding to the contract
	EtherfiNodesManagerTransactor // Write-only binding to the contract
	EtherfiNodesManagerFilterer   // Log filterer for contract events
}

// EtherfiNodesManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiNodesManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiNodesManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiNodesManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiNodesManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiNodesManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiNodesManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiNodesManagerSession struct {
	Contract     *EtherfiNodesManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtherfiNodesManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiNodesManagerCallerSession struct {
	Contract *EtherfiNodesManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EtherfiNodesManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiNodesManagerTransactorSession struct {
	Contract     *EtherfiNodesManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EtherfiNodesManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiNodesManagerRaw struct {
	Contract *EtherfiNodesManager // Generic contract binding to access the raw methods on
}

// EtherfiNodesManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiNodesManagerCallerRaw struct {
	Contract *EtherfiNodesManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiNodesManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiNodesManagerTransactorRaw struct {
	Contract *EtherfiNodesManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiNodesManager creates a new instance of EtherfiNodesManager, bound to a specific deployed contract.
func NewEtherfiNodesManager(address common.Address, backend bind.ContractBackend) (*EtherfiNodesManager, error) {
	contract, err := bindEtherfiNodesManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManager{EtherfiNodesManagerCaller: EtherfiNodesManagerCaller{contract: contract}, EtherfiNodesManagerTransactor: EtherfiNodesManagerTransactor{contract: contract}, EtherfiNodesManagerFilterer: EtherfiNodesManagerFilterer{contract: contract}}, nil
}

// NewEtherfiNodesManagerCaller creates a new read-only instance of EtherfiNodesManager, bound to a specific deployed contract.
func NewEtherfiNodesManagerCaller(address common.Address, caller bind.ContractCaller) (*EtherfiNodesManagerCaller, error) {
	contract, err := bindEtherfiNodesManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerCaller{contract: contract}, nil
}

// NewEtherfiNodesManagerTransactor creates a new write-only instance of EtherfiNodesManager, bound to a specific deployed contract.
func NewEtherfiNodesManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiNodesManagerTransactor, error) {
	contract, err := bindEtherfiNodesManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerTransactor{contract: contract}, nil
}

// NewEtherfiNodesManagerFilterer creates a new log filterer instance of EtherfiNodesManager, bound to a specific deployed contract.
func NewEtherfiNodesManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiNodesManagerFilterer, error) {
	contract, err := bindEtherfiNodesManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerFilterer{contract: contract}, nil
}

// bindEtherfiNodesManager binds a generic wrapper to an already deployed contract.
func bindEtherfiNodesManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherfiNodesManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiNodesManager *EtherfiNodesManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiNodesManager.Contract.EtherfiNodesManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiNodesManager *EtherfiNodesManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.EtherfiNodesManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiNodesManager *EtherfiNodesManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.EtherfiNodesManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiNodesManager *EtherfiNodesManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiNodesManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.contract.Transact(opts, method, params...)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDAdmin(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDAdmin(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DEPRECATEDEnableNodeRecycling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "DEPRECATED_enableNodeRecycling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DEPRECATEDEnableNodeRecycling() (bool, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDEnableNodeRecycling(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DEPRECATEDEnableNodeRecycling() (bool, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDEnableNodeRecycling(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DEPRECATEDProtocolRevenueManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRevenueManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRevenueManager(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRevenueManager(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DEPRECATEDProtocolRevenueManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRevenueManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DEPRECATEDProtocolRevenueManagerContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRevenueManagerContract(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DEPRECATEDProtocolRevenueManagerContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRevenueManagerContract(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DEPRECATEDProtocolRewardsSplit(opts *bind.CallOpts) (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRewardsSplit")

	outstruct := new(struct {
		Treasury     uint64
		NodeOperator uint64
		Tnft         uint64
		Bnft         uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NodeOperator = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Tnft = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Bnft = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DEPRECATEDProtocolRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRewardsSplit(&_EtherfiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DEPRECATEDProtocolRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherfiNodesManager.Contract.DEPRECATEDProtocolRewardsSplit(&_EtherfiNodesManager.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) SCALE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "SCALE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SCALE() (uint64, error) {
	return _EtherfiNodesManager.Contract.SCALE(&_EtherfiNodesManager.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) SCALE() (uint64, error) {
	return _EtherfiNodesManager.Contract.SCALE(&_EtherfiNodesManager.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.Admins(&_EtherfiNodesManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.Admins(&_EtherfiNodesManager.CallOpts, arg0)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) AllowedForwardedEigenpodCalls(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "allowedForwardedEigenpodCalls", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherfiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherfiNodesManager.CallOpts, arg0)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherfiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherfiNodesManager.CallOpts, arg0)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) AllowedForwardedExternalCalls(opts *bind.CallOpts, arg0 [4]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "allowedForwardedExternalCalls", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherfiNodesManager.CallOpts, arg0, arg1)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherfiNodesManager.CallOpts, arg0, arg1)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) AuctionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "auctionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) AuctionManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.AuctionManager(&_EtherfiNodesManager.CallOpts)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) AuctionManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.AuctionManager(&_EtherfiNodesManager.CallOpts)
}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Bnft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "bnft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Bnft() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Bnft(&_EtherfiNodesManager.CallOpts)
}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Bnft() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Bnft(&_EtherfiNodesManager.CallOpts)
}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) CalculateTVL(opts *bind.CallOpts, _validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "calculateTVL", _validatorId, _beaconBalance)

	outstruct := new(struct {
		ToNodeOperator *big.Int
		ToTnft         *big.Int
		ToBnft         *big.Int
		ToTreasury     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToNodeOperator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ToTnft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ToBnft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ToTreasury = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) CalculateTVL(_validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherfiNodesManager.Contract.CalculateTVL(&_EtherfiNodesManager.CallOpts, _validatorId, _beaconBalance)
}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) CalculateTVL(_validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherfiNodesManager.Contract.CalculateTVL(&_EtherfiNodesManager.CallOpts, _validatorId, _beaconBalance)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DelayedWithdrawalRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "delayedWithdrawalRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DelayedWithdrawalRouter(&_EtherfiNodesManager.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DelayedWithdrawalRouter(&_EtherfiNodesManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) DelegationManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DelegationManager(&_EtherfiNodesManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) DelegationManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.DelegationManager(&_EtherfiNodesManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) EigenPodManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.EigenPodManager(&_EtherfiNodesManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _EtherfiNodesManager.Contract.EigenPodManager(&_EtherfiNodesManager.CallOpts)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) EtherfiNodeAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "etherfiNodeAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) EtherfiNodeAddress(arg0 *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.EtherfiNodeAddress(&_EtherfiNodesManager.CallOpts, arg0)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) EtherfiNodeAddress(arg0 *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.EtherfiNodeAddress(&_EtherfiNodesManager.CallOpts, arg0)
}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GenerateWithdrawalCredentials(opts *bind.CallOpts, _address common.Address) ([]byte, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "generateWithdrawalCredentials", _address)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GenerateWithdrawalCredentials(_address common.Address) ([]byte, error) {
	return _EtherfiNodesManager.Contract.GenerateWithdrawalCredentials(&_EtherfiNodesManager.CallOpts, _address)
}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GenerateWithdrawalCredentials(_address common.Address) ([]byte, error) {
	return _EtherfiNodesManager.Contract.GenerateWithdrawalCredentials(&_EtherfiNodesManager.CallOpts, _address)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetEigenPod(opts *bind.CallOpts, _validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getEigenPod", _validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetEigenPod(_validatorId *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetEigenPod(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetEigenPod(_validatorId *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetEigenPod(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetFullWithdrawalPayouts(opts *bind.CallOpts, _validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getFullWithdrawalPayouts", _validatorId)

	outstruct := new(struct {
		ToNodeOperator *big.Int
		ToTnft         *big.Int
		ToBnft         *big.Int
		ToTreasury     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToNodeOperator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ToTnft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ToBnft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ToTreasury = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetFullWithdrawalPayouts(_validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherfiNodesManager.Contract.GetFullWithdrawalPayouts(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetFullWithdrawalPayouts(_validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherfiNodesManager.Contract.GetFullWithdrawalPayouts(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetImplementation() (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetImplementation(&_EtherfiNodesManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetImplementation() (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetImplementation(&_EtherfiNodesManager.CallOpts)
}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetNonExitPenalty(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getNonExitPenalty", _validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetNonExitPenalty(_validatorId *big.Int) (*big.Int, error) {
	return _EtherfiNodesManager.Contract.GetNonExitPenalty(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetNonExitPenalty(_validatorId *big.Int) (*big.Int, error) {
	return _EtherfiNodesManager.Contract.GetNonExitPenalty(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetRewardsPayouts(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getRewardsPayouts", _validatorId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetRewardsPayouts(_validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EtherfiNodesManager.Contract.GetRewardsPayouts(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetRewardsPayouts(_validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EtherfiNodesManager.Contract.GetRewardsPayouts(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetUnusedWithdrawalSafesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getUnusedWithdrawalSafesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetUnusedWithdrawalSafesLength() (*big.Int, error) {
	return _EtherfiNodesManager.Contract.GetUnusedWithdrawalSafesLength(&_EtherfiNodesManager.CallOpts)
}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetUnusedWithdrawalSafesLength() (*big.Int, error) {
	return _EtherfiNodesManager.Contract.GetUnusedWithdrawalSafesLength(&_EtherfiNodesManager.CallOpts)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetValidatorInfo(opts *bind.CallOpts, _validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getValidatorInfo", _validatorId)

	if err != nil {
		return *new(IEtherFiNodesManagerValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEtherFiNodesManagerValidatorInfo)).(*IEtherFiNodesManagerValidatorInfo)

	return out0, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetValidatorInfo(_validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	return _EtherfiNodesManager.Contract.GetValidatorInfo(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetValidatorInfo(_validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	return _EtherfiNodesManager.Contract.GetValidatorInfo(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetWithdrawalCredentials(opts *bind.CallOpts, _validatorId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getWithdrawalCredentials", _validatorId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetWithdrawalCredentials(_validatorId *big.Int) ([]byte, error) {
	return _EtherfiNodesManager.Contract.GetWithdrawalCredentials(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetWithdrawalCredentials(_validatorId *big.Int) ([]byte, error) {
	return _EtherfiNodesManager.Contract.GetWithdrawalCredentials(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) GetWithdrawalSafeAddress(opts *bind.CallOpts, _validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "getWithdrawalSafeAddress", _validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) GetWithdrawalSafeAddress(_validatorId *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetWithdrawalSafeAddress(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) GetWithdrawalSafeAddress(_validatorId *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.GetWithdrawalSafeAddress(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) IsExitRequested(opts *bind.CallOpts, _validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "isExitRequested", _validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) IsExitRequested(_validatorId *big.Int) (bool, error) {
	return _EtherfiNodesManager.Contract.IsExitRequested(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) IsExitRequested(_validatorId *big.Int) (bool, error) {
	return _EtherfiNodesManager.Contract.IsExitRequested(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) MaxEigenlayerWithdrawals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "maxEigenlayerWithdrawals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) MaxEigenlayerWithdrawals() (uint8, error) {
	return _EtherfiNodesManager.Contract.MaxEigenlayerWithdrawals(&_EtherfiNodesManager.CallOpts)
}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) MaxEigenlayerWithdrawals() (uint8, error) {
	return _EtherfiNodesManager.Contract.MaxEigenlayerWithdrawals(&_EtherfiNodesManager.CallOpts)
}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) NonExitPenaltyDailyRate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "nonExitPenaltyDailyRate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) NonExitPenaltyDailyRate() (uint64, error) {
	return _EtherfiNodesManager.Contract.NonExitPenaltyDailyRate(&_EtherfiNodesManager.CallOpts)
}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) NonExitPenaltyDailyRate() (uint64, error) {
	return _EtherfiNodesManager.Contract.NonExitPenaltyDailyRate(&_EtherfiNodesManager.CallOpts)
}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) NonExitPenaltyPrincipal(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "nonExitPenaltyPrincipal")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) NonExitPenaltyPrincipal() (uint64, error) {
	return _EtherfiNodesManager.Contract.NonExitPenaltyPrincipal(&_EtherfiNodesManager.CallOpts)
}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) NonExitPenaltyPrincipal() (uint64, error) {
	return _EtherfiNodesManager.Contract.NonExitPenaltyPrincipal(&_EtherfiNodesManager.CallOpts)
}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) NumAssociatedValidators(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "numAssociatedValidators", _validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) NumAssociatedValidators(_validatorId *big.Int) (*big.Int, error) {
	return _EtherfiNodesManager.Contract.NumAssociatedValidators(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) NumAssociatedValidators(_validatorId *big.Int) (*big.Int, error) {
	return _EtherfiNodesManager.Contract.NumAssociatedValidators(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) NumberOfValidators(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "numberOfValidators")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) NumberOfValidators() (uint64, error) {
	return _EtherfiNodesManager.Contract.NumberOfValidators(&_EtherfiNodesManager.CallOpts)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) NumberOfValidators() (uint64, error) {
	return _EtherfiNodesManager.Contract.NumberOfValidators(&_EtherfiNodesManager.CallOpts)
}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) OperatingAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "operatingAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) OperatingAdmin(arg0 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.OperatingAdmin(&_EtherfiNodesManager.CallOpts, arg0)
}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) OperatingAdmin(arg0 common.Address) (bool, error) {
	return _EtherfiNodesManager.Contract.OperatingAdmin(&_EtherfiNodesManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Owner() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Owner(&_EtherfiNodesManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Owner() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Owner(&_EtherfiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Paused() (bool, error) {
	return _EtherfiNodesManager.Contract.Paused(&_EtherfiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Paused() (bool, error) {
	return _EtherfiNodesManager.Contract.Paused(&_EtherfiNodesManager.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Phase(opts *bind.CallOpts, _validatorId *big.Int) (uint8, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "phase", _validatorId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Phase(_validatorId *big.Int) (uint8, error) {
	return _EtherfiNodesManager.Contract.Phase(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Phase(_validatorId *big.Int) (uint8, error) {
	return _EtherfiNodesManager.Contract.Phase(&_EtherfiNodesManager.CallOpts, _validatorId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiNodesManager.Contract.ProxiableUUID(&_EtherfiNodesManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiNodesManager.Contract.ProxiableUUID(&_EtherfiNodesManager.CallOpts)
}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) StakingManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "stakingManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) StakingManagerContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.StakingManagerContract(&_EtherfiNodesManager.CallOpts)
}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) StakingManagerContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.StakingManagerContract(&_EtherfiNodesManager.CallOpts)
}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) StakingRewardsSplit(opts *bind.CallOpts) (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "stakingRewardsSplit")

	outstruct := new(struct {
		Treasury     uint64
		NodeOperator uint64
		Tnft         uint64
		Bnft         uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NodeOperator = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Tnft = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Bnft = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) StakingRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherfiNodesManager.Contract.StakingRewardsSplit(&_EtherfiNodesManager.CallOpts)
}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) StakingRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherfiNodesManager.Contract.StakingRewardsSplit(&_EtherfiNodesManager.CallOpts)
}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) Tnft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "tnft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Tnft() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Tnft(&_EtherfiNodesManager.CallOpts)
}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) Tnft() (common.Address, error) {
	return _EtherfiNodesManager.Contract.Tnft(&_EtherfiNodesManager.CallOpts)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) TreasuryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "treasuryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) TreasuryContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.TreasuryContract(&_EtherfiNodesManager.CallOpts)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) TreasuryContract() (common.Address, error) {
	return _EtherfiNodesManager.Contract.TreasuryContract(&_EtherfiNodesManager.CallOpts)
}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCaller) UnusedWithdrawalSafes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiNodesManager.contract.Call(opts, &out, "unusedWithdrawalSafes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UnusedWithdrawalSafes(arg0 *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.UnusedWithdrawalSafes(&_EtherfiNodesManager.CallOpts, arg0)
}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherfiNodesManager *EtherfiNodesManagerCallerSession) UnusedWithdrawalSafes(arg0 *big.Int) (common.Address, error) {
	return _EtherfiNodesManager.Contract.UnusedWithdrawalSafes(&_EtherfiNodesManager.CallOpts, arg0)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) AllocateEtherFiNode(opts *bind.TransactOpts, _enableRestaking bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "allocateEtherFiNode", _enableRestaking)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) AllocateEtherFiNode(_enableRestaking bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.AllocateEtherFiNode(&_EtherfiNodesManager.TransactOpts, _enableRestaking)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) AllocateEtherFiNode(_enableRestaking bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.AllocateEtherFiNode(&_EtherfiNodesManager.TransactOpts, _enableRestaking)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) BatchFullWithdraw(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "batchFullWithdraw", _validatorIds)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) BatchFullWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchFullWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) BatchFullWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchFullWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) BatchPartialWithdraw(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "batchPartialWithdraw", _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) BatchPartialWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchPartialWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) BatchPartialWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchPartialWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) BatchQueueRestakedWithdrawal(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "batchQueueRestakedWithdrawal", _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) BatchQueueRestakedWithdrawal(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchQueueRestakedWithdrawal(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) BatchQueueRestakedWithdrawal(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchQueueRestakedWithdrawal(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) BatchSendExitRequest(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "batchSendExitRequest", _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) BatchSendExitRequest(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchSendExitRequest(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) BatchSendExitRequest(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.BatchSendExitRequest(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xde01da5c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[] middlewareTimesIndexes, bool _receiveAsTokens) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, _validatorIds []*big.Int, withdrawals []IDelegationManagerWithdrawal, middlewareTimesIndexes []*big.Int, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "completeQueuedWithdrawals", _validatorIds, withdrawals, middlewareTimesIndexes, _receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xde01da5c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[] middlewareTimesIndexes, bool _receiveAsTokens) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) CompleteQueuedWithdrawals(_validatorIds []*big.Int, withdrawals []IDelegationManagerWithdrawal, middlewareTimesIndexes []*big.Int, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherfiNodesManager.TransactOpts, _validatorIds, withdrawals, middlewareTimesIndexes, _receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xde01da5c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[] middlewareTimesIndexes, bool _receiveAsTokens) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) CompleteQueuedWithdrawals(_validatorIds []*big.Int, withdrawals []IDelegationManagerWithdrawal, middlewareTimesIndexes []*big.Int, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherfiNodesManager.TransactOpts, _validatorIds, withdrawals, middlewareTimesIndexes, _receiveAsTokens)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) ForwardEigenpodCall(opts *bind.TransactOpts, _validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "forwardEigenpodCall", _validatorIds, _data)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ForwardEigenpodCall(_validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardEigenpodCall(&_EtherfiNodesManager.TransactOpts, _validatorIds, _data)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) ForwardEigenpodCall(_validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardEigenpodCall(&_EtherfiNodesManager.TransactOpts, _validatorIds, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) ForwardEigenpodCall0(opts *bind.TransactOpts, _etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "forwardEigenpodCall0", _etherfiNodes, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ForwardEigenpodCall0(_etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardEigenpodCall0(&_EtherfiNodesManager.TransactOpts, _etherfiNodes, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) ForwardEigenpodCall0(_etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardEigenpodCall0(&_EtherfiNodesManager.TransactOpts, _etherfiNodes, _data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) ForwardExternalCall(opts *bind.TransactOpts, _etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "forwardExternalCall", _etherfiNodes, _data, _target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ForwardExternalCall(_etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardExternalCall(&_EtherfiNodesManager.TransactOpts, _etherfiNodes, _data, _target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) ForwardExternalCall(_etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardExternalCall(&_EtherfiNodesManager.TransactOpts, _etherfiNodes, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) ForwardExternalCall0(opts *bind.TransactOpts, _validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "forwardExternalCall0", _validatorIds, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ForwardExternalCall0(_validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardExternalCall0(&_EtherfiNodesManager.TransactOpts, _validatorIds, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) ForwardExternalCall0(_validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ForwardExternalCall0(&_EtherfiNodesManager.TransactOpts, _validatorIds, _data, _target)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) FullWithdraw(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "fullWithdraw", _validatorId)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) FullWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.FullWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) FullWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.FullWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) IncrementNumberOfValidators(opts *bind.TransactOpts, _count uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "incrementNumberOfValidators", _count)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) IncrementNumberOfValidators(_count uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.IncrementNumberOfValidators(&_EtherfiNodesManager.TransactOpts, _count)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) IncrementNumberOfValidators(_count uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.IncrementNumberOfValidators(&_EtherfiNodesManager.TransactOpts, _count)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) MarkBeingSlashed(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "markBeingSlashed", _validatorIds)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) MarkBeingSlashed(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.MarkBeingSlashed(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) MarkBeingSlashed(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.MarkBeingSlashed(&_EtherfiNodesManager.TransactOpts, _validatorIds)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) PartialWithdraw(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "partialWithdraw", _validatorId)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) PartialWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.PartialWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) PartialWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.PartialWithdraw(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) PauseContract() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.PauseContract(&_EtherfiNodesManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.PauseContract(&_EtherfiNodesManager.TransactOpts)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) ProcessNodeExit(opts *bind.TransactOpts, _validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "processNodeExit", _validatorIds, _exitTimestamps)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) ProcessNodeExit(_validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ProcessNodeExit(&_EtherfiNodesManager.TransactOpts, _validatorIds, _exitTimestamps)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) ProcessNodeExit(_validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.ProcessNodeExit(&_EtherfiNodesManager.TransactOpts, _validatorIds, _exitTimestamps)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) RegisterValidator(opts *bind.TransactOpts, _validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "registerValidator", _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) RegisterValidator(_validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.RegisterValidator(&_EtherfiNodesManager.TransactOpts, _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) RegisterValidator(_validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.RegisterValidator(&_EtherfiNodesManager.TransactOpts, _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.RenounceOwnership(&_EtherfiNodesManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.RenounceOwnership(&_EtherfiNodesManager.TransactOpts)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) SetMaxEigenLayerWithdrawals(opts *bind.TransactOpts, _max uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "setMaxEigenLayerWithdrawals", _max)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SetMaxEigenLayerWithdrawals(_max uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetMaxEigenLayerWithdrawals(&_EtherfiNodesManager.TransactOpts, _max)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) SetMaxEigenLayerWithdrawals(_max uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetMaxEigenLayerWithdrawals(&_EtherfiNodesManager.TransactOpts, _max)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) SetNonExitPenalty(opts *bind.TransactOpts, _nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "setNonExitPenalty", _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SetNonExitPenalty(_nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetNonExitPenalty(&_EtherfiNodesManager.TransactOpts, _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) SetNonExitPenalty(_nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetNonExitPenalty(&_EtherfiNodesManager.TransactOpts, _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) SetProofSubmitter(opts *bind.TransactOpts, _validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "setProofSubmitter", _validatorId, _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SetProofSubmitter(_validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetProofSubmitter(&_EtherfiNodesManager.TransactOpts, _validatorId, _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) SetProofSubmitter(_validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetProofSubmitter(&_EtherfiNodesManager.TransactOpts, _validatorId, _newProofSubmitter)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) SetStakingRewardsSplit(opts *bind.TransactOpts, _treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "setStakingRewardsSplit", _treasury, _nodeOperator, _tnft, _bnft)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SetStakingRewardsSplit(_treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetStakingRewardsSplit(&_EtherfiNodesManager.TransactOpts, _treasury, _nodeOperator, _tnft, _bnft)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) SetStakingRewardsSplit(_treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetStakingRewardsSplit(&_EtherfiNodesManager.TransactOpts, _treasury, _nodeOperator, _tnft, _bnft)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) SetValidatorPhase(opts *bind.TransactOpts, _validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "setValidatorPhase", _validatorId, _phase)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) SetValidatorPhase(_validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetValidatorPhase(&_EtherfiNodesManager.TransactOpts, _validatorId, _phase)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) SetValidatorPhase(_validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.SetValidatorPhase(&_EtherfiNodesManager.TransactOpts, _validatorId, _phase)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) StartCheckpoint(opts *bind.TransactOpts, _validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "startCheckpoint", _validatorId, _revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) StartCheckpoint(_validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.StartCheckpoint(&_EtherfiNodesManager.TransactOpts, _validatorId, _revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) StartCheckpoint(_validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.StartCheckpoint(&_EtherfiNodesManager.TransactOpts, _validatorId, _revertIfNoBalance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.TransferOwnership(&_EtherfiNodesManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.TransferOwnership(&_EtherfiNodesManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UnPauseContract(&_EtherfiNodesManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UnPauseContract(&_EtherfiNodesManager.TransactOpts)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UnregisterValidator(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "unregisterValidator", _validatorId)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UnregisterValidator(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UnregisterValidator(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UnregisterValidator(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UnregisterValidator(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAdmin(&_EtherfiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAdmin(&_EtherfiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpdateAllowedForwardedEigenpodCalls(opts *bind.TransactOpts, _selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "updateAllowedForwardedEigenpodCalls", _selector, _allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpdateAllowedForwardedEigenpodCalls(_selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherfiNodesManager.TransactOpts, _selector, _allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpdateAllowedForwardedEigenpodCalls(_selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherfiNodesManager.TransactOpts, _selector, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpdateAllowedForwardedExternalCalls(opts *bind.TransactOpts, _selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "updateAllowedForwardedExternalCalls", _selector, _target, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpdateAllowedForwardedExternalCalls(_selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherfiNodesManager.TransactOpts, _selector, _target, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpdateAllowedForwardedExternalCalls(_selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherfiNodesManager.TransactOpts, _selector, _target, _allowed)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpdateEigenLayerOperatingAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "updateEigenLayerOperatingAdmin", _address, _isAdmin)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpdateEigenLayerOperatingAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateEigenLayerOperatingAdmin(&_EtherfiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpdateEigenLayerOperatingAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateEigenLayerOperatingAdmin(&_EtherfiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpdateEtherFiNode(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "updateEtherFiNode", _validatorId)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpdateEtherFiNode(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateEtherFiNode(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpdateEtherFiNode(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpdateEtherFiNode(&_EtherfiNodesManager.TransactOpts, _validatorId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpgradeTo(&_EtherfiNodesManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpgradeTo(&_EtherfiNodesManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpgradeToAndCall(&_EtherfiNodesManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.UpgradeToAndCall(&_EtherfiNodesManager.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiNodesManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerSession) Receive() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.Receive(&_EtherfiNodesManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiNodesManager *EtherfiNodesManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherfiNodesManager.Contract.Receive(&_EtherfiNodesManager.TransactOpts)
}

// EtherfiNodesManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAdminChangedIterator struct {
	Event *EtherfiNodesManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerAdminChanged)
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
		it.Event = new(EtherfiNodesManagerAdminChanged)
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
func (it *EtherfiNodesManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerAdminChanged represents a AdminChanged event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EtherfiNodesManagerAdminChangedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerAdminChangedIterator{contract: _EtherfiNodesManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerAdminChanged)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseAdminChanged(log types.Log) (*EtherfiNodesManagerAdminChanged, error) {
	event := new(EtherfiNodesManagerAdminChanged)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator is returned from FilterAllowedForwardedEigenpodCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedEigenpodCallsUpdated events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator struct {
	Event *EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
		it.Event = new(EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
func (it *EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated represents a AllowedForwardedEigenpodCallsUpdated event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated struct {
	Selector [4]byte
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedEigenpodCallsUpdated is a free log retrieval operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterAllowedForwardedEigenpodCallsUpdated(opts *bind.FilterOpts, selector [][4]byte) (*EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator{contract: _EtherfiNodesManager.contract, event: "AllowedForwardedEigenpodCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedEigenpodCallsUpdated is a free log subscription operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchAllowedForwardedEigenpodCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated, selector [][4]byte) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedEigenpodCallsUpdated is a log parse operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseAllowedForwardedEigenpodCallsUpdated(log types.Log) (*EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated, error) {
	event := new(EtherfiNodesManagerAllowedForwardedEigenpodCallsUpdated)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator is returned from FilterAllowedForwardedExternalCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedExternalCallsUpdated events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator struct {
	Event *EtherfiNodesManagerAllowedForwardedExternalCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerAllowedForwardedExternalCallsUpdated)
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
		it.Event = new(EtherfiNodesManagerAllowedForwardedExternalCallsUpdated)
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
func (it *EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerAllowedForwardedExternalCallsUpdated represents a AllowedForwardedExternalCallsUpdated event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerAllowedForwardedExternalCallsUpdated struct {
	Selector [4]byte
	Target   common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedExternalCallsUpdated is a free log retrieval operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterAllowedForwardedExternalCallsUpdated(opts *bind.FilterOpts, selector [][4]byte, _target []common.Address) (*EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerAllowedForwardedExternalCallsUpdatedIterator{contract: _EtherfiNodesManager.contract, event: "AllowedForwardedExternalCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedExternalCallsUpdated is a free log subscription operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchAllowedForwardedExternalCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerAllowedForwardedExternalCallsUpdated, selector [][4]byte, _target []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerAllowedForwardedExternalCallsUpdated)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedExternalCallsUpdated is a log parse operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseAllowedForwardedExternalCallsUpdated(log types.Log) (*EtherfiNodesManagerAllowedForwardedExternalCallsUpdated, error) {
	event := new(EtherfiNodesManagerAllowedForwardedExternalCallsUpdated)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerBeaconUpgradedIterator struct {
	Event *EtherfiNodesManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerBeaconUpgraded)
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
		it.Event = new(EtherfiNodesManagerBeaconUpgraded)
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
func (it *EtherfiNodesManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EtherfiNodesManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerBeaconUpgradedIterator{contract: _EtherfiNodesManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerBeaconUpgraded)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EtherfiNodesManagerBeaconUpgraded, error) {
	event := new(EtherfiNodesManagerBeaconUpgraded)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerFullWithdrawalIterator is returned from FilterFullWithdrawal and is used to iterate over the raw logs and unpacked data for FullWithdrawal events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerFullWithdrawalIterator struct {
	Event *EtherfiNodesManagerFullWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerFullWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerFullWithdrawal)
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
		it.Event = new(EtherfiNodesManagerFullWithdrawal)
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
func (it *EtherfiNodesManagerFullWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerFullWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerFullWithdrawal represents a FullWithdrawal event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerFullWithdrawal struct {
	ValidatorId *big.Int
	EtherFiNode common.Address
	ToOperator  *big.Int
	ToTnft      *big.Int
	ToBnft      *big.Int
	ToTreasury  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFullWithdrawal is a free log retrieval operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterFullWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherfiNodesManagerFullWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerFullWithdrawalIterator{contract: _EtherfiNodesManager.contract, event: "FullWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFullWithdrawal is a free log subscription operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchFullWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerFullWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerFullWithdrawal)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
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

// ParseFullWithdrawal is a log parse operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseFullWithdrawal(log types.Log) (*EtherfiNodesManagerFullWithdrawal, error) {
	event := new(EtherfiNodesManagerFullWithdrawal)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerFundsWithdrawnIterator struct {
	Event *EtherfiNodesManagerFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerFundsWithdrawn)
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
		it.Event = new(EtherfiNodesManagerFundsWithdrawn)
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
func (it *EtherfiNodesManagerFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerFundsWithdrawn represents a FundsWithdrawn event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerFundsWithdrawn struct {
	ValidatorId *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xd66662c0ded9e58fd31d5e44944bcfd07ffc15e6927ecc1382e7941cb7bd24c4.
//
// Solidity: event FundsWithdrawn(uint256 indexed _validatorId, uint256 amount)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, _validatorId []*big.Int) (*EtherfiNodesManagerFundsWithdrawnIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "FundsWithdrawn", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerFundsWithdrawnIterator{contract: _EtherfiNodesManager.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xd66662c0ded9e58fd31d5e44944bcfd07ffc15e6927ecc1382e7941cb7bd24c4.
//
// Solidity: event FundsWithdrawn(uint256 indexed _validatorId, uint256 amount)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerFundsWithdrawn, _validatorId []*big.Int) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "FundsWithdrawn", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerFundsWithdrawn)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xd66662c0ded9e58fd31d5e44944bcfd07ffc15e6927ecc1382e7941cb7bd24c4.
//
// Solidity: event FundsWithdrawn(uint256 indexed _validatorId, uint256 amount)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseFundsWithdrawn(log types.Log) (*EtherfiNodesManagerFundsWithdrawn, error) {
	event := new(EtherfiNodesManagerFundsWithdrawn)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerInitializedIterator struct {
	Event *EtherfiNodesManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerInitialized)
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
		it.Event = new(EtherfiNodesManagerInitialized)
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
func (it *EtherfiNodesManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerInitialized represents a Initialized event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtherfiNodesManagerInitializedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerInitializedIterator{contract: _EtherfiNodesManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerInitialized)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseInitialized(log types.Log) (*EtherfiNodesManagerInitialized, error) {
	event := new(EtherfiNodesManagerInitialized)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerNodeEvictedIterator is returned from FilterNodeEvicted and is used to iterate over the raw logs and unpacked data for NodeEvicted events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeEvictedIterator struct {
	Event *EtherfiNodesManagerNodeEvicted // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerNodeEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerNodeEvicted)
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
		it.Event = new(EtherfiNodesManagerNodeEvicted)
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
func (it *EtherfiNodesManagerNodeEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerNodeEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerNodeEvicted represents a NodeEvicted event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeEvicted struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeEvicted is a free log retrieval operation binding the contract event 0x4fbc91e67bf682202f6b274513267e96191f13676abd1e43573d8bacfc2ac136.
//
// Solidity: event NodeEvicted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterNodeEvicted(opts *bind.FilterOpts) (*EtherfiNodesManagerNodeEvictedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "NodeEvicted")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerNodeEvictedIterator{contract: _EtherfiNodesManager.contract, event: "NodeEvicted", logs: logs, sub: sub}, nil
}

// WatchNodeEvicted is a free log subscription operation binding the contract event 0x4fbc91e67bf682202f6b274513267e96191f13676abd1e43573d8bacfc2ac136.
//
// Solidity: event NodeEvicted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchNodeEvicted(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerNodeEvicted) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "NodeEvicted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerNodeEvicted)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeEvicted", log); err != nil {
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

// ParseNodeEvicted is a log parse operation binding the contract event 0x4fbc91e67bf682202f6b274513267e96191f13676abd1e43573d8bacfc2ac136.
//
// Solidity: event NodeEvicted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseNodeEvicted(log types.Log) (*EtherfiNodesManagerNodeEvicted, error) {
	event := new(EtherfiNodesManagerNodeEvicted)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerNodeExitProcessedIterator is returned from FilterNodeExitProcessed and is used to iterate over the raw logs and unpacked data for NodeExitProcessed events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitProcessedIterator struct {
	Event *EtherfiNodesManagerNodeExitProcessed // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerNodeExitProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerNodeExitProcessed)
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
		it.Event = new(EtherfiNodesManagerNodeExitProcessed)
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
func (it *EtherfiNodesManagerNodeExitProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerNodeExitProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerNodeExitProcessed represents a NodeExitProcessed event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitProcessed struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeExitProcessed is a free log retrieval operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterNodeExitProcessed(opts *bind.FilterOpts) (*EtherfiNodesManagerNodeExitProcessedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "NodeExitProcessed")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerNodeExitProcessedIterator{contract: _EtherfiNodesManager.contract, event: "NodeExitProcessed", logs: logs, sub: sub}, nil
}

// WatchNodeExitProcessed is a free log subscription operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchNodeExitProcessed(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerNodeExitProcessed) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "NodeExitProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerNodeExitProcessed)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitProcessed", log); err != nil {
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

// ParseNodeExitProcessed is a log parse operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseNodeExitProcessed(log types.Log) (*EtherfiNodesManagerNodeExitProcessed, error) {
	event := new(EtherfiNodesManagerNodeExitProcessed)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerNodeExitRequestRevertedIterator is returned from FilterNodeExitRequestReverted and is used to iterate over the raw logs and unpacked data for NodeExitRequestReverted events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitRequestRevertedIterator struct {
	Event *EtherfiNodesManagerNodeExitRequestReverted // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerNodeExitRequestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerNodeExitRequestReverted)
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
		it.Event = new(EtherfiNodesManagerNodeExitRequestReverted)
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
func (it *EtherfiNodesManagerNodeExitRequestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerNodeExitRequestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerNodeExitRequestReverted represents a NodeExitRequestReverted event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitRequestReverted struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeExitRequestReverted is a free log retrieval operation binding the contract event 0xaf4cad58a5f970385e5ae6a7e862001a1fa8f35c4dacdfde879f9e9b330b1cdb.
//
// Solidity: event NodeExitRequestReverted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterNodeExitRequestReverted(opts *bind.FilterOpts) (*EtherfiNodesManagerNodeExitRequestRevertedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "NodeExitRequestReverted")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerNodeExitRequestRevertedIterator{contract: _EtherfiNodesManager.contract, event: "NodeExitRequestReverted", logs: logs, sub: sub}, nil
}

// WatchNodeExitRequestReverted is a free log subscription operation binding the contract event 0xaf4cad58a5f970385e5ae6a7e862001a1fa8f35c4dacdfde879f9e9b330b1cdb.
//
// Solidity: event NodeExitRequestReverted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchNodeExitRequestReverted(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerNodeExitRequestReverted) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "NodeExitRequestReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerNodeExitRequestReverted)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitRequestReverted", log); err != nil {
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

// ParseNodeExitRequestReverted is a log parse operation binding the contract event 0xaf4cad58a5f970385e5ae6a7e862001a1fa8f35c4dacdfde879f9e9b330b1cdb.
//
// Solidity: event NodeExitRequestReverted(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseNodeExitRequestReverted(log types.Log) (*EtherfiNodesManagerNodeExitRequestReverted, error) {
	event := new(EtherfiNodesManagerNodeExitRequestReverted)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitRequestReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerNodeExitRequestedIterator is returned from FilterNodeExitRequested and is used to iterate over the raw logs and unpacked data for NodeExitRequested events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitRequestedIterator struct {
	Event *EtherfiNodesManagerNodeExitRequested // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerNodeExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerNodeExitRequested)
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
		it.Event = new(EtherfiNodesManagerNodeExitRequested)
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
func (it *EtherfiNodesManagerNodeExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerNodeExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerNodeExitRequested represents a NodeExitRequested event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerNodeExitRequested struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeExitRequested is a free log retrieval operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterNodeExitRequested(opts *bind.FilterOpts) (*EtherfiNodesManagerNodeExitRequestedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "NodeExitRequested")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerNodeExitRequestedIterator{contract: _EtherfiNodesManager.contract, event: "NodeExitRequested", logs: logs, sub: sub}, nil
}

// WatchNodeExitRequested is a free log subscription operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchNodeExitRequested(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerNodeExitRequested) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "NodeExitRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerNodeExitRequested)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitRequested", log); err != nil {
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

// ParseNodeExitRequested is a log parse operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseNodeExitRequested(log types.Log) (*EtherfiNodesManagerNodeExitRequested, error) {
	event := new(EtherfiNodesManagerNodeExitRequested)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "NodeExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerOwnershipTransferredIterator struct {
	Event *EtherfiNodesManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerOwnershipTransferred)
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
		it.Event = new(EtherfiNodesManagerOwnershipTransferred)
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
func (it *EtherfiNodesManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EtherfiNodesManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerOwnershipTransferredIterator{contract: _EtherfiNodesManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerOwnershipTransferred)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EtherfiNodesManagerOwnershipTransferred, error) {
	event := new(EtherfiNodesManagerOwnershipTransferred)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerPartialWithdrawalIterator is returned from FilterPartialWithdrawal and is used to iterate over the raw logs and unpacked data for PartialWithdrawal events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPartialWithdrawalIterator struct {
	Event *EtherfiNodesManagerPartialWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerPartialWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerPartialWithdrawal)
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
		it.Event = new(EtherfiNodesManagerPartialWithdrawal)
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
func (it *EtherfiNodesManagerPartialWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerPartialWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerPartialWithdrawal represents a PartialWithdrawal event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPartialWithdrawal struct {
	ValidatorId *big.Int
	EtherFiNode common.Address
	ToOperator  *big.Int
	ToTnft      *big.Int
	ToBnft      *big.Int
	ToTreasury  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawal is a free log retrieval operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterPartialWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherfiNodesManagerPartialWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerPartialWithdrawalIterator{contract: _EtherfiNodesManager.contract, event: "PartialWithdrawal", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawal is a free log subscription operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchPartialWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerPartialWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerPartialWithdrawal)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
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

// ParsePartialWithdrawal is a log parse operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParsePartialWithdrawal(log types.Log) (*EtherfiNodesManagerPartialWithdrawal, error) {
	event := new(EtherfiNodesManagerPartialWithdrawal)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPausedIterator struct {
	Event *EtherfiNodesManagerPaused // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerPaused)
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
		it.Event = new(EtherfiNodesManagerPaused)
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
func (it *EtherfiNodesManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerPaused represents a Paused event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherfiNodesManagerPausedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerPausedIterator{contract: _EtherfiNodesManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerPaused)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParsePaused(log types.Log) (*EtherfiNodesManagerPaused, error) {
	event := new(EtherfiNodesManagerPaused)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerPhaseChangedIterator is returned from FilterPhaseChanged and is used to iterate over the raw logs and unpacked data for PhaseChanged events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPhaseChangedIterator struct {
	Event *EtherfiNodesManagerPhaseChanged // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerPhaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerPhaseChanged)
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
		it.Event = new(EtherfiNodesManagerPhaseChanged)
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
func (it *EtherfiNodesManagerPhaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerPhaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerPhaseChanged represents a PhaseChanged event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerPhaseChanged struct {
	ValidatorId *big.Int
	Phase       uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPhaseChanged is a free log retrieval operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterPhaseChanged(opts *bind.FilterOpts, _validatorId []*big.Int) (*EtherfiNodesManagerPhaseChangedIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "PhaseChanged", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerPhaseChangedIterator{contract: _EtherfiNodesManager.contract, event: "PhaseChanged", logs: logs, sub: sub}, nil
}

// WatchPhaseChanged is a free log subscription operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchPhaseChanged(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerPhaseChanged, _validatorId []*big.Int) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "PhaseChanged", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerPhaseChanged)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "PhaseChanged", log); err != nil {
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

// ParsePhaseChanged is a log parse operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParsePhaseChanged(log types.Log) (*EtherfiNodesManagerPhaseChanged, error) {
	event := new(EtherfiNodesManagerPhaseChanged)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "PhaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerQueuedRestakingWithdrawalIterator is returned from FilterQueuedRestakingWithdrawal and is used to iterate over the raw logs and unpacked data for QueuedRestakingWithdrawal events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerQueuedRestakingWithdrawalIterator struct {
	Event *EtherfiNodesManagerQueuedRestakingWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerQueuedRestakingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerQueuedRestakingWithdrawal)
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
		it.Event = new(EtherfiNodesManagerQueuedRestakingWithdrawal)
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
func (it *EtherfiNodesManagerQueuedRestakingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerQueuedRestakingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerQueuedRestakingWithdrawal represents a QueuedRestakingWithdrawal event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerQueuedRestakingWithdrawal struct {
	ValidatorId     *big.Int
	EtherFiNode     common.Address
	WithdrawalRoots [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterQueuedRestakingWithdrawal is a free log retrieval operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterQueuedRestakingWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherfiNodesManagerQueuedRestakingWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerQueuedRestakingWithdrawalIterator{contract: _EtherfiNodesManager.contract, event: "QueuedRestakingWithdrawal", logs: logs, sub: sub}, nil
}

// WatchQueuedRestakingWithdrawal is a free log subscription operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchQueuedRestakingWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerQueuedRestakingWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerQueuedRestakingWithdrawal)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
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

// ParseQueuedRestakingWithdrawal is a log parse operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseQueuedRestakingWithdrawal(log types.Log) (*EtherfiNodesManagerQueuedRestakingWithdrawal, error) {
	event := new(EtherfiNodesManagerQueuedRestakingWithdrawal)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerUnpausedIterator struct {
	Event *EtherfiNodesManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerUnpaused)
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
		it.Event = new(EtherfiNodesManagerUnpaused)
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
func (it *EtherfiNodesManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerUnpaused represents a Unpaused event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherfiNodesManagerUnpausedIterator, error) {

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerUnpausedIterator{contract: _EtherfiNodesManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerUnpaused)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseUnpaused(log types.Log) (*EtherfiNodesManagerUnpaused, error) {
	event := new(EtherfiNodesManagerUnpaused)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiNodesManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerUpgradedIterator struct {
	Event *EtherfiNodesManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiNodesManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiNodesManagerUpgraded)
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
		it.Event = new(EtherfiNodesManagerUpgraded)
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
func (it *EtherfiNodesManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiNodesManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiNodesManagerUpgraded represents a Upgraded event raised by the EtherfiNodesManager contract.
type EtherfiNodesManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EtherfiNodesManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiNodesManagerUpgradedIterator{contract: _EtherfiNodesManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiNodesManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiNodesManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiNodesManagerUpgraded)
				if err := _EtherfiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiNodesManager *EtherfiNodesManagerFilterer) ParseUpgraded(log types.Log) (*EtherfiNodesManagerUpgraded, error) {
	event := new(EtherfiNodesManagerUpgraded)
	if err := _EtherfiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
