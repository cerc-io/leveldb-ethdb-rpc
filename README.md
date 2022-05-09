# leveldb-ethdb-rpc
Thin RPC wrapper around LevelDB to expose data remotely

## Setup

Run the following

```bash
make build
```

Create a `config.toml` file from [example.toml](./environments/example.toml) file.

Update the config for path to geth leveldb and geth ancient database

```toml
[leveldb]
    path = "/path/to/eth/data/geth/chaindata" # $LEVELDB_PATH
    ancient = "/path/to/eth/data/geth/chaindata/ancient" # $LEVELDB_ANCIENT_PATH
```

## Usage

After building the binary, run as

`./leveldb-ethdb-rpc serve --config ./environments/config.toml`
