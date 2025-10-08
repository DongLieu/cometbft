package tooling_keyset

import (
	"fmt"
	"os"

	"github.com/cometbft/cometbft/privval"
)

var (
	pubKeys              []*privval.FilePV
	keyProposal          *privval.FilePV
	path_file_publickeys = "/Users/donglieu/script/onomy/tooling/publickeys.json"
)

func init() {
	var err error
	pubKeys, err = LoadFilePVsFromPubkeysJSON(path_file_publickeys)
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

func LoadFilePVsFromPubkeysJSON(filePath string) ([]*privval.FilePV, error) {
	var filePVs []*privval.FilePV

	dir_save_keys, err := os.MkdirTemp("", "tmp_keys")
	if err != nil {
		panic(err)
	}

	_, paths, err := getValidatorsInfo(filePath, dir_save_keys)
	if err != nil {
		panic(err)
	}

	defer func() {
		for _, p := range paths {
			if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
				fmt.Printf("⚠️ Failed to remove file %s: %v\n", p, err)
			}
		}
	}()

	for _, path := range paths {
		tmFilePV := privval.LoadFilePVEmptyState(path, "")
		filePVs = append(filePVs, tmFilePV)
	}
	return filePVs, nil
}
