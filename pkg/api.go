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
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/ethdb"
)

// APIName is the namespace used for the state diffing service API
const APIName = "leveldb"

// APIVersion is the version of the state diffing service API
const APIVersion = "0.0.1"

var (
	errWriteNotAllowed = errors.New("write endpoints are not enabled")
)

type PublicLevelDBAPI struct {
	b *LevelDBBackend
}

func NewPublicLevelDBAPI(b *LevelDBBackend) *PublicLevelDBAPI {
	return &PublicLevelDBAPI{b: b}
}

func (s *PublicLevelDBAPI) Has(ctx context.Context, key []byte) (bool, error) {
	return s.b.Has(key)
}

func (s *PublicLevelDBAPI) Get(ctx context.Context, key []byte) ([]byte, error) {
	return s.b.Get(key)
}

func (s *PublicLevelDBAPI) HasAncient(ctx context.Context, kind string, number uint64) (bool, error) {
	return s.b.HasAncient(kind, number)
}

func (s *PublicLevelDBAPI) Ancient(ctx context.Context, kind string, number uint64) ([]byte, error) {
	return s.b.Ancient(kind, number)
}

func (s *PublicLevelDBAPI) AncientRange(ctx context.Context, kind string, start, count, maxBytes uint64) ([][]byte, error) {
	return s.b.AncientRange(kind, start, count, maxBytes)
}

func (s *PublicLevelDBAPI) ReadAncients(fn func(ethdb.AncientReader) error) error {
	return s.b.ReadAncients(fn)
}

func (s *PublicLevelDBAPI) Ancients(ctx context.Context) (uint64, error) {
	return s.b.Ancients()
}

func (s *PublicLevelDBAPI) AncientSize(ctx context.Context, kind string) (uint64, error) {
	return s.b.AncientSize(kind)
}

func (s *PublicLevelDBAPI) Stat(ctx context.Context, property string) (string, error) {
	return s.b.Stat(property)
}
