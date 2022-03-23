package types

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type Events struct {
	Claims_Claimed                               []EventClaimsClaimed                           //nolint:stylecheck,golint
	Claims_RootHashStored                        []EventClaimsRootHashStored                    //nolint:stylecheck,golint
	CrowdloanClaim_ClaimPalletInitialized        []EventCrowdloanClaimClaimPalletInitialized    //nolint:stylecheck,golint
	CrowdloanClaim_RewardClaimed                 []EventCrowdloanClaimRewardClaimed             //nolint:stylecheck,golint
	CrowdloanClaim_ClaimLockedAtUpdated          []EventCrowdloanClaimLockedAtUpdated           //nolint:stylecheck,golint
	CrowdloanClaim_ClaimContributionsRootUpdated []EventCrowdloanClaimContributionsRootUpdated  //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanTrieIndexUpdated     []EventCrowdloanClaimCrowdloanTrieIndexUpdated //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeaseStartUpdated    []EventCrowdloanClaimLeaseStartUpdated         //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeasePeriodUpdated   []EventCrowdloanClaimLeasePeriodUpdated        //nolint:stylecheck,golint
	CrowdloanReward_RewardClaimed                []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_RewardPalletInitialized      []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_DirectPayoutRatioUpdated     []EventCrowdloanRewardDirectPayoutRatioUpdated //nolint:stylecheck,golint
	CrowdloanReward_VestingPeriodUpdated         []EventCrowdloanRewardVestingPeriodUpdated     //nolint:stylecheck,golint
	CrowdloanReward_VestingStartUpdated          []EventCrowdloanRewardVestingStartUpdated      //nolint:stylecheck,golint
	CollatorAllowlist_CollatorAdded              []EventCollatorAllowlistCollatorAdded          //nolint:stylecheck,golint
	CollatorAllowlist_CollatorRemoved            []EventCollatorAllowlistCollatorRemoved        //nolint:stylecheck,golint
	Fees_FeeChanged                              []EventFeesFeeChanged                          //nolint:stylecheck,golint
	Nfts_DepositAsset                            []EventNftsDepositAsset                        //nolint:stylecheck,golint
}

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
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
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
