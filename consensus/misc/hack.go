// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package misc

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// ErrBadProHackExtra is returned if a header doesn't support the Hack fork on a
	// pro-fork client.
	ErrBadProHackExtra = errors.New("bad Hack pro-fork extra-data")
)

// VerifyHackHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to Hack hard-fork rules.
//
// Hack hard-fork extension to the header validity,
// require block in the specific number to have the unique extra-data set.
func VerifyHackHeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	if config.HackForkBlock == nil {
		return nil
	}
	if header.Number.Cmp(config.HackForkBlock) != 0 {
		return nil
	}
	if len(header.Extra) < len(params.HackForkBlockExtra) {
		return ErrBadProHackExtra
	}
	if !bytes.Equal(header.Extra[:len(params.HackForkBlockExtra)], params.HackForkBlockExtra) {
		return ErrBadProHackExtra
	}
	// All ok, header has the extra-data we expect
	return nil
}

// ApplyHackHardFork modifies the state database according to the Hack hard-fork
// rules, transferring all balances of a set of Hacker accounts to a single refund
// address.
func ApplyHackHardFork(statedb *state.StateDB) {
	// Retrieve the address to refund balances into
	if !statedb.Exist(params.HackRefundAddress) {
		statedb.CreateAccount(params.HackRefundAddress)
	}

	// Move every Hack account and extra-balance account funds into the refund address
	for _, addr := range params.HackDrainList() {
		statedb.AddBalance(params.HackRefundAddress, statedb.GetBalance(addr))
		statedb.SetBalance(addr, new(big.Int))
	}
}
