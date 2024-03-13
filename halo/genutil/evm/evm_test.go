package evm_test

import (
	"testing"

	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/test/tutil"

	_ "github.com/omni-network/omni/halo/app" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()

	genesis := evm.MakeDevGenesis()
	tutil.RequireGoldenJSON(t, genesis)
}
