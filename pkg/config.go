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

package leveldb_ethdb_rpc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/fdlimit"
	"github.com/spf13/viper"
)

// Config struct holds the configuration parameters for the levelDB RPC service
type Config struct {
	IPCEnabled   bool
	IPCEndpoint  string
	HTTPEnabled  bool
	HTTPEndpoint string

	FilePath    string
	Cache       int
	Handles     int
	FreezerPath string
	Namespace   string
}

// NewConfig returns a new Config from viper parameters
func NewConfig() (*Config, error) {
	viper.BindEnv(TOML_IPC_ENABLED, IPC_ENABLED)
	viper.BindEnv(TOML_IPC_ENDPOINT, IPC_ENDPOINT)
	viper.BindEnv(TOML_HTTP_ENABLED, HTTP_ENABLED)
	viper.BindEnv(TOML_HTTP_ENDPOINT, HTTP_ENDPOINT)

	viper.BindEnv(TOML_LEVELDB_PATH, LEVELDB_PATH)
	viper.BindEnv(TOML_LEVELDB_CACHE_SIZE, LEVELDB_CACHE_SIZE)
	viper.BindEnv(TOML_LEVELDB_ANCIENT_PATH, LEVELDB_ANCIENT_PATH)
	viper.BindEnv(TOML_LEVELDB_NAMESPACE, LEVELDB_NAMESPACE)

	numHandles, err := MakeDatabaseHandles()
	if err != nil {
		return nil, err
	}
	return &Config{
		IPCEnabled:   viper.GetBool(TOML_IPC_ENABLED),
		IPCEndpoint:  viper.GetString(TOML_IPC_ENDPOINT),
		HTTPEnabled:  viper.GetBool(TOML_HTTP_ENABLED),
		HTTPEndpoint: viper.GetString(TOML_HTTP_ENDPOINT),
		FilePath:     viper.GetString(TOML_LEVELDB_PATH),
		Cache:        viper.GetInt(TOML_LEVELDB_CACHE_SIZE),
		Handles:      numHandles,
		FreezerPath:  viper.GetString(TOML_LEVELDB_ANCIENT_PATH),
		Namespace:    viper.GetString(TOML_LEVELDB_NAMESPACE),
	}, nil
}

// MakeDatabaseHandles raises out the number of allowed file handles per process
// for Geth and returns half of the allowance to assign to the database.
func MakeDatabaseHandles() (int, error) {
	limit, err := fdlimit.Maximum()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve file descriptor allowance: %v", err)
	}
	raised, err := fdlimit.Raise(uint64(limit))
	if err != nil {
		return 0, fmt.Errorf("failed to raise file descriptor allowance: %v", err)
	}
	return int(raised / 2), nil // Leave half for networking and other stuff
}
