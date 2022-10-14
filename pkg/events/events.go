package events

import (
	"github.com/uinb/chain-custom-types/pkg/keystore"
)

type Events struct {
	ChainBridge_FungibleTransfer        []EventChainBridgeFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventChainBridgeNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventChainBridgeGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainBridgeChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventChainBridgeRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventChainBridgeRelayerThresholdChanged //nolint:stylecheck,golint

	Claims_Claimed        []EventClaimsClaimed        //nolint:stylecheck,golint
	Claims_RootHashStored []EventClaimsRootHashStored //nolint:stylecheck,golint

	CrowdloanClaim_ClaimPalletInitialized        []EventCrowdloanClaimClaimPalletInitialized    //nolint:stylecheck,golint
	CrowdloanClaim_RewardClaimed                 []EventCrowdloanClaimRewardClaimed             //nolint:stylecheck,golint
	CrowdloanClaim_ClaimLockedAtUpdated          []EventCrowdloanClaimLockedAtUpdated           //nolint:stylecheck,golint
	CrowdloanClaim_ClaimContributionsRootUpdated []EventCrowdloanClaimContributionsRootUpdated  //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanTrieIndexUpdated     []EventCrowdloanClaimCrowdloanTrieIndexUpdated //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeaseStartUpdated    []EventCrowdloanClaimLeaseStartUpdated         //nolint:stylecheck,golint
	CrowdloanClaim_CrowdloanLeasePeriodUpdated   []EventCrowdloanClaimLeasePeriodUpdated        //nolint:stylecheck,golint

	CrowdloanReward_RewardClaimed            []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_RewardPalletInitialized  []EventCrowdloanRewardRewardClaimed            //nolint:stylecheck,golint
	CrowdloanReward_DirectPayoutRatioUpdated []EventCrowdloanRewardDirectPayoutRatioUpdated //nolint:stylecheck,golint
	CrowdloanReward_VestingPeriodUpdated     []EventCrowdloanRewardVestingPeriodUpdated     //nolint:stylecheck,golint
	CrowdloanReward_VestingStartUpdated      []EventCrowdloanRewardVestingStartUpdated      //nolint:stylecheck,golint

	CollatorAllowlist_CollatorAdded   []EventCollatorAllowlistCollatorAdded   //nolint:stylecheck,golint
	CollatorAllowlist_CollatorRemoved []EventCollatorAllowlistCollatorRemoved //nolint:stylecheck,golint

	Fees_FeeChanged    []EventFeesFeeChanged    //nolint:stylecheck,golint
	Fees_FeeToAuthor   []EventFeesFeeToAuthor   //nolint:stylecheck,golint
	Fees_FeeToBurn     []EventFeesFeeToBurn     //nolint:stylecheck,golint
	Fees_FeeToTreasury []EventFeesFeeToTreasury //nolint:stylecheck,golint

	Keystore_KeyAdded   []keystore.EventKeystoreKeyAdded   //nolint:stylecheck,golint
	Keystore_KeyRevoked []keystore.EventKeystoreKeyRevoked //nolint:stylecheck,golint
	Keystore_DepositSet []keystore.EventKeystoreDepositSet //nolint:stylecheck,golint

	Nfts_DepositAsset []EventNftsDepositAsset //nolint:stylecheck,golint
	Nft_Transferred   []EventNftTransferred   //nolint:stylecheck,golint

	Registry_RegistryCreated []EventRegistryRegistryCreated //nolint:stylecheck,golint
	Registry_Mint            []EventRegistryNftMint         //nolint:stylecheck,golint
}
