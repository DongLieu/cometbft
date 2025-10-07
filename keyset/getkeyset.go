package keyset

import (
	"github.com/cometbft/cometbft/privval"
)

var pubKeys []*privval.FilePV
var keyProposal *privval.FilePV

func init() {
	var err error
	pubKeys, err = privval.LoadFilePVsFromPubkeysJSON("/Users/donglieu/script/onomy/tooling/publickeys.json")
	if err != nil {
		panic(err)
	}
}

func GetPubKeys() []*privval.FilePV {
	return pubKeys
}

func GetkeyProposal() *privval.FilePV {
	return keyProposal
}

func SetkeyProposal(key *privval.FilePV) {
	keyProposal = key
}
