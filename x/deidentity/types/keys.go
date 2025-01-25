package types

const (
	// ModuleName defines the module name
	ModuleName = "deidentity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_deidentity"
)

var (
	ParamsKey = []byte("p_deidentity")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	IdentityKey      = "Identity/value/"
	IdentityCountKey = "Identity/count/"
)
