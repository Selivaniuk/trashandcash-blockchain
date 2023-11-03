package types

const (
	// ModuleName defines the module name
	ModuleName = "trashandcashblockchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trashandcashblockchain"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SpinKey      = "Spin/value/"
	SpinCountKey = "Spin/count/"
)
const (
	SpinReelsEventType    = "spin-reels"
	SpinReelsEventId      = "id"
	SpinReelsEventBet     = "bet"
	SpinReelsEventAddress = "address"
	SpinReelsEventResult  = "result"
	SpinReelsEventWinning = "winning"
)
