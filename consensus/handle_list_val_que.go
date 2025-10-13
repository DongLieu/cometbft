package consensus

import (
	"bytes"

	"github.com/cometbft/cometbft/crypto"
	types "github.com/cometbft/cometbft/types"
)

// getProposer()
func (cs *State) getProposer() crypto.PubKey {
	for _, pubkey := range cs.listprivValidatorPubKey {
		address := pubkey.Address()
		if cs.isProposer(address) {
			return pubkey
		}

	}
	return nil
}

// getval()
func (cs *State) getValByAddrees(addr []byte) crypto.PubKey {
	for _, pubkey := range cs.listprivValidatorPubKey {
		address := pubkey.Address()
		if bytes.Equal(address, addr) {
			return pubkey
		}

	}
	return nil
}

// setProposer()
func (cs *State) setProposer(proposer types.Address) crypto.PubKey {
	for _, pubkey := range cs.listprivValidatorPubKey {
		address := pubkey.Address()
		if bytes.Equal(address, proposer) {
			cs.privValidatorPubKey = pubkey
			return pubkey
		}

	}
	return nil
}
