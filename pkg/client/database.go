// Copyright © 2022 Vulcanize, Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"errors"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rpc"
)

var errNotSupported = errors.New("this operation is not supported")

var _ ethdb.Database = &DatabaseClient{}

// Type that satisfies the ethdb.DatabaseClient using a leveldb-ethdb-rpc client
type DatabaseClient struct {
	client *rpc.Client
}

// NewDatabase returns a ethdb.Database interface
func NewDatabaseClient(url string) (ethdb.Database, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	database := DatabaseClient{
		client: rpcClient,
	}

	return &database, nil
}

// Has satisfies the ethdb.KeyValueReader interface
// Has retrieves if a key is present in the key-value data store
func (d *DatabaseClient) Has(key []byte) (bool, error) {
	var resp bool
	err := d.client.Call(&resp, "leveldb_has", key)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Get satisfies the ethdb.KeyValueReader interface
// Get retrieves the given key if it's present in the key-value data store
func (d *DatabaseClient) Get(key []byte) ([]byte, error) {
	var resp []byte
	err := d.client.Call(&resp, "leveldb_get", key)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Put satisfies the ethdb.KeyValueWriter interface
// Put inserts the given value into the key-value data store
// Key is expected to be the keccak256 hash of value
func (d *DatabaseClient) Put(key []byte, value []byte) error {
	return errNotSupported
}

// Delete satisfies the ethdb.KeyValueWriter interface
// Delete removes the key from the key-value data store
func (d *DatabaseClient) Delete(key []byte) error {
	return errNotSupported
}

// Stat satisfies the ethdb.Stater interface
// Stat returns a particular internal stat of the database
func (d *DatabaseClient) Stat(property string) (string, error) {
	var resp string
	err := d.client.Call(&resp, "leveldb_stat", property)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Compact satisfies the ethdb.Compacter interface
// Compact flattens the underlying data store for the given key range
func (d *DatabaseClient) Compact(start []byte, limit []byte) error {
	return errNotSupported
}

// NewBatch satisfies the ethdb.Batcher interface
// NewBatch creates a write-only database that buffers changes to its host db
// until a final write is called
func (d *DatabaseClient) NewBatch() ethdb.Batch {
	return nil
}

// NewBatchWithSize satisfies the ethdb.Batcher interface.
// NewBatchWithSize creates a write-only database batch with pre-allocated buffer.
func (d *DatabaseClient) NewBatchWithSize(size int) ethdb.Batch {
	return nil
}

// NewIterator satisfies the ethdb.Iteratee interface
// it creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
//
// Note: This method assumes that the prefix is NOT part of the start, so there's
// no need for the caller to prepend the prefix to the start
func (d *DatabaseClient) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	return nil
}

// Close satisfies the io.Closer interface
// Close closes the db connection
func (d *DatabaseClient) Close() error {
	return errNotSupported
}

// HasAncient satisfies the ethdb.AncientReader interface
// HasAncient returns an indicator whether the specified data exists in the ancient store
func (d *DatabaseClient) HasAncient(kind string, number uint64) (bool, error) {
	var resp bool
	err := d.client.Call(&resp, "leveldb_hasAncient", kind, number)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Ancient satisfies the ethdb.AncientReader interface
// Ancient retrieves an ancient binary blob from the append-only immutable files
func (d *DatabaseClient) Ancient(kind string, number uint64) ([]byte, error) {
	var resp []byte
	err := d.client.Call(&resp, "leveldb_ancient", kind, number)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Ancients satisfies the ethdb.AncientReader interface
// Ancients returns the ancient item numbers in the ancient store
func (d *DatabaseClient) Ancients() (uint64, error) {
	var resp uint64
	err := d.client.Call(&resp, "leveldb_ancients")
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Tail satisfies the ethdb.AncientReader interface.
// Tail returns the number of first stored item in the freezer.
func (d *DatabaseClient) Tail() (uint64, error) {
	return 0, errNotSupported
}

// AncientSize satisfies the ethdb.AncientReader interface
// AncientSize returns the ancient size of the specified category
func (d *DatabaseClient) AncientSize(kind string) (uint64, error) {
	var resp uint64
	err := d.client.Call(&resp, "leveldb_ancientSize")
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// AncientRange retrieves all the items in a range, starting from the index 'start'.
// It will return
//   - at most 'count' items,
//   - at least 1 item (even if exceeding the maxBytes), but will otherwise
//     return as many items as fit into maxBytes.
func (d *DatabaseClient) AncientRange(kind string, start, count, maxBytes uint64) ([][]byte, error) {
	var resp [][]byte
	err := d.client.Call(&resp, "leveldb_ancientRange", kind, start, count, maxBytes)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ReadAncients applies the provided AncientReader function
func (d *DatabaseClient) ReadAncients(fn func(ethdb.AncientReaderOp) error) (err error) {
	return fn(d)
}

// ModifyAncients satisfies the ethdb.AncientWriter interface.
// ModifyAncients runs a write operation on the ancient store.
func (d *DatabaseClient) ModifyAncients(f func(ethdb.AncientWriteOp) error) (int64, error) {
	return 0, errNotSupported
}

// TruncateHead satisfies the ethdb.AncientWriter interface.
// TruncateHead discards all but the first n ancient data from the ancient store.
func (d *DatabaseClient) TruncateHead(n uint64) (uint64, error) {
	return 0, errNotSupported
}

// TruncateTail satisfies the ethdb.AncientWriter interface.
// TruncateTail discards the first n ancient data from the ancient store.
func (d *DatabaseClient) TruncateTail(n uint64) (uint64, error) {
	return 0, errNotSupported
}

// Sync satisfies the ethdb.AncientWriter interface
// Sync flushes all in-memory ancient store data to disk
func (d *DatabaseClient) Sync() error {
	return errNotSupported
}

// MigrateTable satisfies the ethdb.AncientWriter interface.
// MigrateTable processes and migrates entries of a given table to a new format.
func (d *DatabaseClient) MigrateTable(string, func([]byte) ([]byte, error)) error {
	return errNotSupported
}

// NewSnapshot satisfies the ethdb.Snapshotter interface.
// NewSnapshot creates a database snapshot based on the current state.
func (d *DatabaseClient) NewSnapshot() (ethdb.Snapshot, error) {
	return nil, errNotSupported
}

// AncientDatadir returns an error as we don't have a backing chain freezer.
func (d *DatabaseClient) AncientDatadir() (string, error) {
	return "", errNotSupported
}
