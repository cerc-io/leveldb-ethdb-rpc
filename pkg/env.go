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

const (
	LOGRUS_LEVEL = "LOGRUS_LEVEL"
	LOGRUS_FILE  = "LOGRUS_FILE"

	IPC_ENABLED   = "IPC_ENABLED"
	IPC_ENDPOINT  = "IPC_PATH"
	HTTP_ENABLED  = "HTTP_ENABLED"
	HTTP_ENDPOINT = "HTTP_PATH"

	LEVELDB_PATH         = "LEVELDB_PATH"
	LEVELDB_CACHE_SIZE   = "LEVELDB_CACHE_SIZE"
	LEVELDB_ANCIENT_PATH = "LEVELDB_ANCIENT_PATH"
	LEVELDB_NAMESPACE    = "LEVELDB_NAMESPACE"

	TOML_LOGRUS_LEVEL = "log.level"
	TOML_LOGRUS_FILE  = "log.file"

	TOML_IPC_ENABLED   = "leveldb.ipcEnabled"
	TOML_IPC_ENDPOINT  = "leveldb.ipcPath"
	TOML_HTTP_ENABLED  = "leveldb.httpEnabled"
	TOML_HTTP_ENDPOINT = "leveldb.httpPath"

	TOML_LEVELDB_PATH         = "leveldb.path"
	TOML_LEVELDB_CACHE_SIZE   = "leveldb.cacheSize"
	TOML_LEVELDB_ANCIENT_PATH = "leveldb.ancient"
	TOML_LEVELDB_NAMESPACE    = "leveldb.namespace"
)
