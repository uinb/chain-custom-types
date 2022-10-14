package events

import (
	"github.com/uinb/go-substrate-rpc-client/v4/types"
)

// EventCollatorAllowlistCollatorAdded is emitted when a collator has been added to the allowlist
type EventCollatorAllowlistCollatorAdded struct {
	Phase    types.Phase
	Collator types.AccountID
	Topics   []types.Hash
}

// EventCollatorAllowlistCollatorRemoved is emitted when a collator has been removed from the allowlist
type EventCollatorAllowlistCollatorRemoved struct {
	Phase    types.Phase
	Collator types.AccountID
	Topics   []types.Hash
}

// EventCrowdloanClaimClaimPalletInitialized is emitted when the crowdloan claim pallet is properly configured.
type EventCrowdloanClaimClaimPalletInitialized struct {
	Phase  types.Phase
	Topics []types.Hash
}

// EventCrowdloanClaimRewardClaimed is emitted when a reward has been claimed successfully.
type EventCrowdloanClaimRewardClaimed struct {
	Phase            types.Phase
	RelayAccount     types.AccountID
	ParachainAccount types.AccountID
	Amount           types.U128
	Topics           []types.Hash
}

// EventCrowdloanClaimLockedAtUpdated is emitted when, the block number,
// where we lock the contributions has been updated
type EventCrowdloanClaimLockedAtUpdated struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

// EventCrowdloanClaimContributionsRootUpdated is emitted when the relay-chain Root hash changed
// which allows to verify contributions
type EventCrowdloanClaimContributionsRootUpdated struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventCrowdloanClaimLeaseStartUpdated is emitted when the lease start of the parachain slot.
// Used to define when we can initialize the next time
type EventCrowdloanClaimLeaseStartUpdated struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

// EventCrowdloanClaimLeasePeriodUpdated is emitted when the lease start of the parachain slot.
// Used to define when we can initialize the next time
type EventCrowdloanClaimLeasePeriodUpdated struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

// EventCrowdloanClaimCrowdloanTrieIndexUpdated is emitted when the trie index of the crowdloan
// inside the relay-chains crowdloan child storage has been updated
type EventCrowdloanClaimCrowdloanTrieIndexUpdated struct {
	Phase     types.Phase
	TrieIndex types.U32
	Topics    []types.Hash
}

// EventCrowdloanRewardRewardClaimed is emitted when a reward claim was processed successfully.
// [who, direct_reward, vested_reward]
type EventCrowdloanRewardRewardClaimed struct {
	Phase        types.Phase
	Who          types.AccountID
	DirectReward types.U128
	VestedReward types.U128
	Topics       []types.Hash
}

// EventCrowdloanRewardRewardPalletInitialized is emitted when the reward module is ready to reward contributors
// [vesting_start, vesting_period, direct_payout_ratio]
type EventCrowdloanRewardRewardPalletInitialized struct {
	Phase             types.Phase
	VestingStart      types.BlockNumber
	VestingPeriod     types.BlockNumber
	DirectPayoutRatio types.U32
	Topics            []types.Hash
}

// EventCrowdloanRewardDirectPayoutRatioUpdated is emitted when direct payout ratio for contributors has been updated
// [payout_ratio]
type EventCrowdloanRewardDirectPayoutRatioUpdated struct {
	Phase             types.Phase
	DirectPayoutRatio types.U32
	Topics            []types.Hash
}

// EventCrowdloanRewardVestingPeriodUpdated is emitted when the vesting period has been updated
type EventCrowdloanRewardVestingPeriodUpdated struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

// EventCrowdloanRewardVestingStartUpdated is emitted when the vesting period has been updated
type EventCrowdloanRewardVestingStartUpdated struct {
	Phase       types.Phase
	BlockNumber types.BlockNumber
	Topics      []types.Hash
}

// EventNftsDepositAsset is emitted when NFT is ready to be deposited to other chain.
type EventNftsDepositAsset struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeesFeeChanged is emitted when a new fee has been set for a key
type EventFeesFeeChanged struct {
	Phase  types.Phase
	Key    types.U8
	Fee    types.U128
	Topics []types.Hash
}

// EventFeesFeeToAuthor is emitted when a fee has been sent to the block author
type EventFeesFeeToAuthor struct {
	Phase   types.Phase
	From    types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

// EventFeesFeeToBurn is emitted when a fee has been burnt
type EventFeesFeeToBurn struct {
	Phase   types.Phase
	From    types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

// EventFeesFeeToTreasury is emitted when a fee has been sent to the treasury
type EventFeesFeeToTreasury struct {
	Phase   types.Phase
	From    types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

// EventClaimsClaimed is emitted when CFG Tokens have been claimed
type EventClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventClaimsRootHashStored is emitted when RootHash has been stored for the correspondent CFG Claims batch
type EventClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.BlockNumber
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventChainBridgeFungibleTransfer is emitted when a bridge fungible token transfer is executed
type EventChainBridgeFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceID   types.Bytes32
	Amount       types.U32
	Recipient    types.Bytes
	Topics       []types.Hash
}

// EventChainBridgeNonFungibleTransfer is emitted when a bridge non fungible token transfer is executed
type EventChainBridgeNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceID   types.Bytes32
	TokenID      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

// EventChainBridgeGenericTransfer is emitted when a bridge generic transfer is executed
type EventChainBridgeGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceID   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

// EventChainBridgeChainWhitelisted is emitted when a new chain has been whitelisted to interact with the bridge
type EventChainBridgeChainWhitelisted struct {
	Phase       types.Phase
	Destination types.U8
	Topics      []types.Hash
}

// EventChainBridgeRelayerAdded is emitted when a new bridge relayer has been whitelisted
type EventChainBridgeRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

// EventChainBridgeRelayerThresholdChanged is emitted when the relayer threshold is changed
type EventChainBridgeRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

// EventRegistryNftMint is emitted when an NFT with tokenID is minted in a given registryID
type EventRegistryNftMint struct {
	Phase      types.Phase
	RegistryID types.H160
	TokenID    types.U256
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when a new registry is created
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryID types.H160
	Topics     []types.Hash
}

// AssetID is a combination of RegistryID and TokenID
type AssetID struct {
	RegistryID types.H160
	TokenID    types.U256
}

// EventNftTransferred is emitted when an NFT is transferred to a new owner
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryID types.H160
	AssetID    AssetID
	AccountID  types.AccountID
	Topics     []types.Hash
}
