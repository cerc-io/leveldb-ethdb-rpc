// VulcanizeDB
// Copyright © 2022 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	c *rpc.Client
}

// NewClient creates a new Client with a connection to leveldb-ethdb-rpc RPC endpoint.
func NewClient(c *rpc.Client) *Client {
	return &Client{
		c: c,
	}
}

func (c *Client) Has(key []byte) (bool, error) {
	var resp bool
	err := c.c.Call(&resp, "leveldb_has", key)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) Get(key []byte) ([]byte, error) {
	var resp []byte
	err := c.c.Call(&resp, "leveldb_get", key)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
