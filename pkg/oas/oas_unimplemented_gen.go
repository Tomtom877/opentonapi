// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DnsBackResolve implements dnsBackResolve operation.
//
// Get domains for wallet account.
//
// GET /v2/accounts/{account_id}/dns/backresolve
func (UnimplementedHandler) DnsBackResolve(ctx context.Context, params DnsBackResolveParams) (r *DomainNames, _ error) {
	return r, ht.ErrNotImplemented
}

// DnsInfo implements dnsInfo operation.
//
// Get full information about domain name.
//
// GET /v2/dns/{domain_name}
func (UnimplementedHandler) DnsInfo(ctx context.Context, params DnsInfoParams) (r *DomainInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// DnsResolve implements dnsResolve operation.
//
// DNS resolve for domain name.
//
// GET /v2/dns/{domain_name}/resolve
func (UnimplementedHandler) DnsResolve(ctx context.Context, params DnsResolveParams) (r *DnsRecord, _ error) {
	return r, ht.ErrNotImplemented
}

// EmulateMessageToAccountEvent implements emulateMessageToAccountEvent operation.
//
// Emulate sending message to blockchain.
//
// POST /v2/accounts/{account_id}/events/emulate
func (UnimplementedHandler) EmulateMessageToAccountEvent(ctx context.Context, req *EmulateMessageToAccountEventReq, params EmulateMessageToAccountEventParams) (r *AccountEvent, _ error) {
	return r, ht.ErrNotImplemented
}

// EmulateMessageToEvent implements emulateMessageToEvent operation.
//
// Emulate sending message to blockchain.
//
// POST /v2/events/emulate
func (UnimplementedHandler) EmulateMessageToEvent(ctx context.Context, req *EmulateMessageToEventReq, params EmulateMessageToEventParams) (r *Event, _ error) {
	return r, ht.ErrNotImplemented
}

// EmulateMessageToTrace implements emulateMessageToTrace operation.
//
// Emulate sending message to blockchain.
//
// POST /v2/traces/emulate
func (UnimplementedHandler) EmulateMessageToTrace(ctx context.Context, req *EmulateMessageToTraceReq) (r *Trace, _ error) {
	return r, ht.ErrNotImplemented
}

// EmulateWalletMessage implements emulateWalletMessage operation.
//
// Emulate sending message to blockchain.
//
// POST /v2/wallet/emulate
func (UnimplementedHandler) EmulateWalletMessage(ctx context.Context, req *EmulateWalletMessageReq, params EmulateWalletMessageParams) (r *MessageConsequences, _ error) {
	return r, ht.ErrNotImplemented
}

// ExecGetMethod implements execGetMethod operation.
//
// Execute get method for account.
//
// GET /v2/blockchain/accounts/{account_id}/methods/{method_name}
func (UnimplementedHandler) ExecGetMethod(ctx context.Context, params ExecGetMethodParams) (r *MethodExecutionResult, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccount implements getAccount operation.
//
// Get human-friendly information about an account without low-level details.
//
// GET /v2/accounts/{account_id}
func (UnimplementedHandler) GetAccount(ctx context.Context, params GetAccountParams) (r *Account, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountInfoByStateInit implements getAccountInfoByStateInit operation.
//
// Get account info by state init.
//
// POST /v2/tonconnect/stateinit
func (UnimplementedHandler) GetAccountInfoByStateInit(ctx context.Context, req *GetAccountInfoByStateInitReq) (r *AccountInfoByStateInit, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountSeqno implements getAccountSeqno operation.
//
// Get account seqno.
//
// GET /v2/wallet/{account_id}/seqno
func (UnimplementedHandler) GetAccountSeqno(ctx context.Context, params GetAccountSeqnoParams) (r *Seqno, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountStateLiteServer implements getAccountStateLiteServer operation.
//
// Get account state.
//
// GET /v2/liteserver/get_account_state/{account_id}
func (UnimplementedHandler) GetAccountStateLiteServer(ctx context.Context, params GetAccountStateLiteServerParams) (r *GetAccountStateLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountTransactions implements getAccountTransactions operation.
//
// Get account transactions.
//
// GET /v2/blockchain/accounts/{account_id}/transactions
func (UnimplementedHandler) GetAccountTransactions(ctx context.Context, params GetAccountTransactionsParams) (r *Transactions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccounts implements getAccounts operation.
//
// Get human-friendly information about several accounts without low-level details.
//
// POST /v2/accounts/_bulk
func (UnimplementedHandler) GetAccounts(ctx context.Context, req OptGetAccountsReq) (r *Accounts, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllAuctions implements getAllAuctions operation.
//
// Get all auctions.
//
// GET /v2/dns/auctions
func (UnimplementedHandler) GetAllAuctions(ctx context.Context, params GetAllAuctionsParams) (r *Auctions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllShardsInfoLiteServer implements getAllShardsInfoLiteServer operation.
//
// Get all shards info.
//
// GET /v2/liteserver/get_all_shards_info/{block_id}
func (UnimplementedHandler) GetAllShardsInfoLiteServer(ctx context.Context, params GetAllShardsInfoLiteServerParams) (r *GetAllShardsInfoLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlock implements getBlock operation.
//
// Get block data.
//
// GET /v2/blockchain/blocks/{block_id}
func (UnimplementedHandler) GetBlock(ctx context.Context, params GetBlockParams) (r *Block, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockHeaderLiteServer implements getBlockHeaderLiteServer operation.
//
// Get block header.
//
// GET /v2/liteserver/get_block_header/{block_id}
func (UnimplementedHandler) GetBlockHeaderLiteServer(ctx context.Context, params GetBlockHeaderLiteServerParams) (r *GetBlockHeaderLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockLiteServer implements getBlockLiteServer operation.
//
// Get block.
//
// GET /v2/liteserver/get_block/{block_id}
func (UnimplementedHandler) GetBlockLiteServer(ctx context.Context, params GetBlockLiteServerParams) (r *GetBlockLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockProofLiteServer implements getBlockProofLiteServer operation.
//
// Get block proof.
//
// GET /v2/liteserver/get_block_proof
func (UnimplementedHandler) GetBlockProofLiteServer(ctx context.Context, params GetBlockProofLiteServerParams) (r *GetBlockProofLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockTransactions implements getBlockTransactions operation.
//
// Get transactions from block.
//
// GET /v2/blockchain/blocks/{block_id}/transactions
func (UnimplementedHandler) GetBlockTransactions(ctx context.Context, params GetBlockTransactionsParams) (r *Transactions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChartsRates implements getChartsRates operation.
//
// Get charts by token.
//
// GET /v2/rates/charts
func (UnimplementedHandler) GetChartsRates(ctx context.Context, params GetChartsRatesParams) (r *GetChartsRatesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetConfig implements getConfig operation.
//
// Get blockchain config.
//
// GET /v2/blockchain/config
func (UnimplementedHandler) GetConfig(ctx context.Context) (r *Config, _ error) {
	return r, ht.ErrNotImplemented
}

// GetConfigAllLiteServer implements getConfigAllLiteServer operation.
//
// Get config all.
//
// GET /v2/liteserver/get_config_all/{block_id}
func (UnimplementedHandler) GetConfigAllLiteServer(ctx context.Context, params GetConfigAllLiteServerParams) (r *GetConfigAllLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDnsExpiring implements getDnsExpiring operation.
//
// Get expiring .ton dns.
//
// GET /v2/accounts/{account_id}/dns/expiring
func (UnimplementedHandler) GetDnsExpiring(ctx context.Context, params GetDnsExpiringParams) (r *DnsExpiring, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDomainBids implements getDomainBids operation.
//
// Get domain bids.
//
// GET /v2/dns/{domain_name}/bids
func (UnimplementedHandler) GetDomainBids(ctx context.Context, params GetDomainBidsParams) (r *DomainBids, _ error) {
	return r, ht.ErrNotImplemented
}

// GetEvent implements getEvent operation.
//
// Get an event either by event ID or a hash of any transaction in a trace. An event is built on top
// of a trace which is a series of transactions caused by one inbound message. TonAPI looks for known
// patterns inside the trace and splits the trace into actions, where a single action represents a
// meaningful high-level operation like a Jetton Transfer or an NFT Purchase. Actions are expected to
// be shown to users. It is advised not to build any logic on top of actions because actions can be
// changed at any time.
//
// GET /v2/events/{event_id}
func (UnimplementedHandler) GetEvent(ctx context.Context, params GetEventParams) (r *Event, _ error) {
	return r, ht.ErrNotImplemented
}

// GetEventsByAccount implements getEventsByAccount operation.
//
// Get events for an account. Each event is built on top of a trace which is a series of transactions
// caused by one inbound message. TonAPI looks for known patterns inside the trace and splits the
// trace into actions, where a single action represents a meaningful high-level operation like a
// Jetton Transfer or an NFT Purchase. Actions are expected to be shown to users. It is advised not
// to build any logic on top of actions because actions can be changed at any time.
//
// GET /v2/accounts/{account_id}/events
func (UnimplementedHandler) GetEventsByAccount(ctx context.Context, params GetEventsByAccountParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetItemsFromCollection implements getItemsFromCollection operation.
//
// Get NFT items from collection by collection address.
//
// GET /v2/nfts/collections/{account_id}/items
func (UnimplementedHandler) GetItemsFromCollection(ctx context.Context, params GetItemsFromCollectionParams) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonInfo implements getJettonInfo operation.
//
// Get jetton metadata by jetton master address.
//
// GET /v2/jettons/{account_id}
func (UnimplementedHandler) GetJettonInfo(ctx context.Context, params GetJettonInfoParams) (r *JettonInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettons implements getJettons operation.
//
// Get a list of all indexed jetton masters in the blockchain.
//
// GET /v2/jettons
func (UnimplementedHandler) GetJettons(ctx context.Context, params GetJettonsParams) (r *Jettons, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonsBalances implements getJettonsBalances operation.
//
// Get all Jettons balances by owner address.
//
// GET /v2/accounts/{account_id}/jettons
func (UnimplementedHandler) GetJettonsBalances(ctx context.Context, params GetJettonsBalancesParams) (r *JettonsBalances, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonsHistory implements getJettonsHistory operation.
//
// Get the transfer jettons history for account_id.
//
// GET /v2/accounts/{account_id}/jettons/history
func (UnimplementedHandler) GetJettonsHistory(ctx context.Context, params GetJettonsHistoryParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonsHistoryByID implements getJettonsHistoryByID operation.
//
// Get the transfer jetton history for account_id and jetton_id.
//
// GET /v2/accounts/{account_id}/jettons/{jetton_id}/history
func (UnimplementedHandler) GetJettonsHistoryByID(ctx context.Context, params GetJettonsHistoryByIDParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetListBlockTransactionsLiteServer implements getListBlockTransactionsLiteServer operation.
//
// Get list block transactions.
//
// GET /v2/liteserver/list_block_transactions/{block_id}
func (UnimplementedHandler) GetListBlockTransactionsLiteServer(ctx context.Context, params GetListBlockTransactionsLiteServerParams) (r *GetListBlockTransactionsLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMasterchainHead implements getMasterchainHead operation.
//
// Get last known masterchain block.
//
// GET /v2/blockchain/masterchain-head
func (UnimplementedHandler) GetMasterchainHead(ctx context.Context) (r *Block, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMasterchainInfoExtLiteServer implements getMasterchainInfoExtLiteServer operation.
//
// Get masterchain info ext.
//
// GET /v2/liteserver/get_masterchain_info_ext
func (UnimplementedHandler) GetMasterchainInfoExtLiteServer(ctx context.Context, params GetMasterchainInfoExtLiteServerParams) (r *GetMasterchainInfoExtLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMasterchainInfoLiteServer implements getMasterchainInfoLiteServer operation.
//
// Get masterchain info.
//
// GET /v2/liteserver/get_masterchain_info
func (UnimplementedHandler) GetMasterchainInfoLiteServer(ctx context.Context) (r *GetMasterchainInfoLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftCollection implements getNftCollection operation.
//
// Get NFT collection by collection address.
//
// GET /v2/nfts/collections/{account_id}
func (UnimplementedHandler) GetNftCollection(ctx context.Context, params GetNftCollectionParams) (r *NftCollection, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftCollections implements getNftCollections operation.
//
// Get NFT collections.
//
// GET /v2/nfts/collections
func (UnimplementedHandler) GetNftCollections(ctx context.Context, params GetNftCollectionsParams) (r *NftCollections, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftItemByAddress implements getNftItemByAddress operation.
//
// Get NFT item by its address.
//
// GET /v2/nfts/{account_id}
func (UnimplementedHandler) GetNftItemByAddress(ctx context.Context, params GetNftItemByAddressParams) (r *NftItem, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftItemsByAddresses implements getNftItemsByAddresses operation.
//
// Get NFT items by their addresses.
//
// POST /v2/nfts/_bulk
func (UnimplementedHandler) GetNftItemsByAddresses(ctx context.Context, req OptGetNftItemsByAddressesReq) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftItemsByOwner implements getNftItemsByOwner operation.
//
// Get all NFT items by owner address.
//
// GET /v2/accounts/{account_id}/nfts
func (UnimplementedHandler) GetNftItemsByOwner(ctx context.Context, params GetNftItemsByOwnerParams) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPublicKeyByAccountID implements getPublicKeyByAccountID operation.
//
// Get public key by account id.
//
// GET /v2/accounts/{account_id}/publickey
func (UnimplementedHandler) GetPublicKeyByAccountID(ctx context.Context, params GetPublicKeyByAccountIDParams) (r *GetPublicKeyByAccountIDOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRates implements getRates operation.
//
// Get the token price to the currency.
//
// GET /v2/rates
func (UnimplementedHandler) GetRates(ctx context.Context, params GetRatesParams) (r *GetRatesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawAccount implements getRawAccount operation.
//
// Get low-level information about an account taken directly from the blockchain.
//
// GET /v2/blockchain/accounts/{account_id}
func (UnimplementedHandler) GetRawAccount(ctx context.Context, params GetRawAccountParams) (r *RawAccount, _ error) {
	return r, ht.ErrNotImplemented
}

// GetSearchAccounts implements getSearchAccounts operation.
//
// Search for accounts by name. You can find the account by the first characters of the domain.
//
// GET /v2/accounts/search
func (UnimplementedHandler) GetSearchAccounts(ctx context.Context, params GetSearchAccountsParams) (r *FoundAccounts, _ error) {
	return r, ht.ErrNotImplemented
}

// GetShardBlockProofLiteServer implements getShardBlockProofLiteServer operation.
//
// Get shard block proof.
//
// GET /v2/liteserver/get_shard_block_proof/{block_id}
func (UnimplementedHandler) GetShardBlockProofLiteServer(ctx context.Context, params GetShardBlockProofLiteServerParams) (r *GetShardBlockProofLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetShardInfoLiteServer implements getShardInfoLiteServer operation.
//
// Get shard info.
//
// GET /v2/liteserver/get_shard_info/{block_id}
func (UnimplementedHandler) GetShardInfoLiteServer(ctx context.Context, params GetShardInfoLiteServerParams) (r *GetShardInfoLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStateLiteServer implements getStateLiteServer operation.
//
// Get block state.
//
// GET /v2/liteserver/get_state/{block_id}
func (UnimplementedHandler) GetStateLiteServer(ctx context.Context, params GetStateLiteServerParams) (r *GetStateLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStorageProviders implements getStorageProviders operation.
//
// Get TON storage providers deployed to the blockchain.
//
// GET /v2/storage/providers
func (UnimplementedHandler) GetStorageProviders(ctx context.Context) (r *GetStorageProvidersOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetSubscriptionsByAccount implements getSubscriptionsByAccount operation.
//
// Get all subscriptions by wallet address.
//
// GET /v2/accounts/{account_id}/subscriptions
func (UnimplementedHandler) GetSubscriptionsByAccount(ctx context.Context, params GetSubscriptionsByAccountParams) (r *Subscriptions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTimeLiteServer implements getTimeLiteServer operation.
//
// Get time.
//
// GET /v2/liteserver/get_time
func (UnimplementedHandler) GetTimeLiteServer(ctx context.Context) (r *GetTimeLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTonConnectPayload implements getTonConnectPayload operation.
//
// Get a payload for further token receipt.
//
// GET /v2/tonconnect/payload
func (UnimplementedHandler) GetTonConnectPayload(ctx context.Context) (r *GetTonConnectPayloadOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTrace implements getTrace operation.
//
// Get the trace by trace ID or hash of any transaction in trace.
//
// GET /v2/traces/{trace_id}
func (UnimplementedHandler) GetTrace(ctx context.Context, params GetTraceParams) (r *Trace, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTracesByAccount implements getTracesByAccount operation.
//
// Get traces for account.
//
// GET /v2/accounts/{account_id}/traces
func (UnimplementedHandler) GetTracesByAccount(ctx context.Context, params GetTracesByAccountParams) (r *TraceIds, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTransaction implements getTransaction operation.
//
// Get transaction data.
//
// GET /v2/blockchain/transactions/{transaction_id}
func (UnimplementedHandler) GetTransaction(ctx context.Context, params GetTransactionParams) (r *Transaction, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTransactionByMessageHash implements getTransactionByMessageHash operation.
//
// Get transaction data by message hash.
//
// GET /v2/blockchain/messages/{msg_id}/transaction
func (UnimplementedHandler) GetTransactionByMessageHash(ctx context.Context, params GetTransactionByMessageHashParams) (r *Transaction, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTransactionsLiteServer implements getTransactionsLiteServer operation.
//
// Get transactions.
//
// GET /v2/liteserver/get_transactions/{account_id}
func (UnimplementedHandler) GetTransactionsLiteServer(ctx context.Context, params GetTransactionsLiteServerParams) (r *GetTransactionsLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetValidators implements getValidators operation.
//
// Get validators.
//
// GET /v2/blockchain/validators
func (UnimplementedHandler) GetValidators(ctx context.Context) (r *Validators, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWalletBackup implements getWalletBackup operation.
//
// Get backup info.
//
// GET /v2/wallet/backup
func (UnimplementedHandler) GetWalletBackup(ctx context.Context, params GetWalletBackupParams) (r *GetWalletBackupOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWalletsByPublicKey implements getWalletsByPublicKey operation.
//
// Get wallets by public key.
//
// GET /v2/pubkeys/{public_key}/wallets
func (UnimplementedHandler) GetWalletsByPublicKey(ctx context.Context, params GetWalletsByPublicKeyParams) (r *Accounts, _ error) {
	return r, ht.ErrNotImplemented
}

// PoolsByNominators implements poolsByNominators operation.
//
// All pools where account participates.
//
// GET /v2/staking/nominator/{account_id}/pools
func (UnimplementedHandler) PoolsByNominators(ctx context.Context, params PoolsByNominatorsParams) (r *AccountStaking, _ error) {
	return r, ht.ErrNotImplemented
}

// ReindexAccount implements reindexAccount operation.
//
// Update internal cache for a particular account.
//
// POST /v2/accounts/{account_id}/reindex
func (UnimplementedHandler) ReindexAccount(ctx context.Context, params ReindexAccountParams) error {
	return ht.ErrNotImplemented
}

// SendMessage implements sendMessage operation.
//
// Send message to blockchain.
//
// POST /v2/blockchain/message
func (UnimplementedHandler) SendMessage(ctx context.Context, req *SendMessageReq) error {
	return ht.ErrNotImplemented
}

// SendMessageLiteServer implements sendMessageLiteServer operation.
//
// Send message.
//
// POST /v2/liteserver/send_message
func (UnimplementedHandler) SendMessageLiteServer(ctx context.Context, req *SendMessageLiteServerReq) (r *SendMessageLiteServerOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SetWalletBackup implements setWalletBackup operation.
//
// Set backup info.
//
// PUT /v2/wallet/backup
func (UnimplementedHandler) SetWalletBackup(ctx context.Context, req SetWalletBackupReq, params SetWalletBackupParams) error {
	return ht.ErrNotImplemented
}

// StakingPoolHistory implements stakingPoolHistory operation.
//
// Pool info.
//
// GET /v2/staking/pool/{account_id}/history
func (UnimplementedHandler) StakingPoolHistory(ctx context.Context, params StakingPoolHistoryParams) (r *StakingPoolHistoryOK, _ error) {
	return r, ht.ErrNotImplemented
}

// StakingPoolInfo implements stakingPoolInfo operation.
//
// Pool info.
//
// GET /v2/staking/pool/{account_id}
func (UnimplementedHandler) StakingPoolInfo(ctx context.Context, params StakingPoolInfoParams) (r *StakingPoolInfoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// StakingPools implements stakingPools operation.
//
// All pools available in network.
//
// GET /v2/staking/pools
func (UnimplementedHandler) StakingPools(ctx context.Context, params StakingPoolsParams) (r *StakingPoolsOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TonConnectProof implements tonConnectProof operation.
//
// Account verification and token issuance.
//
// POST /v2/wallet/auth/proof
func (UnimplementedHandler) TonConnectProof(ctx context.Context, req *TonConnectProofReq) (r *TonConnectProofOK, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
