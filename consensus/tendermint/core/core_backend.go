package core

import (
	"context"
	"time"

	"github.com/clearmatics/autonity/common"
	"github.com/clearmatics/autonity/consensus"
	"github.com/clearmatics/autonity/consensus/tendermint/committee"
	"github.com/clearmatics/autonity/core/types"
	"github.com/clearmatics/autonity/event"
)

// Backend provides application specific functions for Istanbul core
type Backend interface {
	Address() common.Address

	AddSeal(block *types.Block) (*types.Block, error)

	AskSync(set committee.Set)

	// Broadcast sends a message to all validators (include self)
	Broadcast(ctx context.Context, valSet committee.Set, payload []byte) error

	// Commit delivers an approved proposal to backend.
	// The delivered proposal will be put into blockchain.
	Commit(proposalBlock *types.Block, round int64, seals [][]byte) error

	GetContractABI() string

	GetContractAddress() common.Address

	// Gossip sends a message to all validators (exclude self)
	Gossip(ctx context.Context, valSet committee.Set, payload []byte)

	HandleUnhandledMsgs(ctx context.Context)

	// LastCommittedProposal retrieves latest committed proposal and the address of proposer
	LastCommittedProposal() (*types.Block, common.Address)

	// GetProposer returns the proposer of the given block height
	GetProposer(number uint64) common.Address

	// GetProposerFromAC returns the proposer of the given block height and round from autontiy contract.
	GetProposerFromAC(height uint64, round int64) common.Address

	// HasBadBlock returns whether the block with the hash is a bad block
	HasBadProposal(hash common.Hash) bool

	Post(ev interface{})

	// Setter for proposed block hash
	SetProposedBlockHash(hash common.Hash)

	// Sign signs input data with the backend's private key
	Sign([]byte) ([]byte, error)

	Subscribe(types ...interface{}) *event.TypeMuxSubscription

	SyncPeer(address common.Address)

	// VerifyProposal verifies the proposal. If a consensus.ErrFutureBlock error is returned,
	// the time difference of the proposal and current time is also returned.
	VerifyProposal(types.Block) (time.Duration, error)

	WhiteList() []string

	BlockChain() consensus.ChainReader
}

type Tendermint interface {
	Start(ctx context.Context)
	Stop()
	GetCurrentHeightMessages() []*Message
}
