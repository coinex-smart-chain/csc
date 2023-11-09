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

package params

import (
	"github.com/ethereum/go-ethereum/common"
)

// HackForkBlockExtra is the block header extra-data field to set for the Hack fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("hack-hard-fork").
var HackForkBlockExtra = common.FromHex("0x34df02625f900697fd3cbdf6762ef0ee")

// HackRefundAddress is the address of the refund address to send Hack balances to.
var HackRefundAddress = common.HexToAddress("0xa7017873d7b8d0d817d0e9586de235d680e03ad3")

// HackDrainList is the list of accounts whose full balances will be moved into a
// refund address at the beginning of the Hack-fork block.
func HackDrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("0x6F5A5A1E19687dB59AC9F06284F11D9E66d2E0C2"),
	}
}
