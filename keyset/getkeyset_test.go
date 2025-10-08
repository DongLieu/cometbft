package keyset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFilePVsFromPubkeysJSON(t *testing.T) {
	privVals, err := LoadFilePVsFromPubkeysJSON("/Users/donglieu/925/cometbft/keyset/keys/publickeys.json")
	fmt.Println(err)
	require.NoError(t, err)
	fmt.Println(privVals[0].Key.PubKey)
	require.Equal(t, len(privVals), 4)
}
