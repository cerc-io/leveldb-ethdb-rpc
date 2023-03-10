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

package cmd

import (
	"os"
	"os/signal"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	leveldb_ethdb_rpc "github.com/cerc-io/leveldb-ethdb-rpc/pkg"
	srpc "github.com/cerc-io/leveldb-ethdb-rpc/pkg/rpc"
	"github.com/cerc-io/leveldb-ethdb-rpc/version"
)

var (
	subCommand     string
	logWithCommand log.Entry
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "RPC Server for LevelDB eth.Database",
	Long:  `This service exposes a remote RPC server interface for a local levelDB backed ethdb.Database`,
	Run: func(cmd *cobra.Command, args []string) {
		subCommand = cmd.CalledAs()
		logWithCommand = *log.WithField("SubCommand", subCommand)
		serve()
	},
}

func serve() {
	logWithCommand.Infof("running ipld-eth-server version: %s", version.VersionWithMeta)

	wg := new(sync.WaitGroup)
	logWithCommand.Debug("loading server configuration variables")
	serverConfig, err := leveldb_ethdb_rpc.NewConfig()
	if err != nil {
		logWithCommand.Fatal(err)
	}
	logWithCommand.Infof("server config: %+v", serverConfig)
	logWithCommand.Debug("initializing new server service")
	server, err := leveldb_ethdb_rpc.NewServer(serverConfig)
	if err != nil {
		logWithCommand.Fatal(err)
	}

	logWithCommand.Info("starting up servers")
	server.Serve(wg)
	if err := startServers(server, serverConfig); err != nil {
		logWithCommand.Fatal(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	server.Stop()
	wg.Wait()
}

func startServers(server leveldb_ethdb_rpc.Server, settings *leveldb_ethdb_rpc.Config) error {
	if settings.IPCEnabled {
		logWithCommand.Info("starting up IPC server")
		_, _, err := srpc.StartIPCEndpoint(settings.IPCEndpoint, server.APIs())
		if err != nil {
			return err
		}
	} else {
		logWithCommand.Info("IPC server is disabled")
	}

	if settings.HTTPEnabled {
		logWithCommand.Info("starting up HTTP server")
		_, err := srpc.StartHTTPEndpoint(settings.HTTPEndpoint, server.APIs(), []string{"leveldb"}, nil, []string{"*"}, rpc.HTTPTimeouts{})
		if err != nil {
			return err
		}
	} else {
		logWithCommand.Info("HTTP server is disabled")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// CLI flags
	serveCmd.PersistentFlags().Bool("ipc-enabled", false, "turn on ipc server")
	serveCmd.PersistentFlags().String("ipc-path", "", "ipc server endpoint")
	serveCmd.PersistentFlags().Bool("http-enabled", true, "turn on http server; on by default")
	serveCmd.PersistentFlags().String("http-path", "127.0.0.1:8500", "http server endpoint; default = 127.0.0.1:8545")

	serveCmd.PersistentFlags().String("leveldb-path", "", "leveldb filesystem path")
	serveCmd.PersistentFlags().Int("leveldb-cache-size", 0, "leveldb cache size")
	serveCmd.PersistentFlags().String("leveldb-ancient-path", "", "filesystem path to freezer")
	serveCmd.PersistentFlags().String("leveldb-namespace", "eth/db/chaindata/", "leveldb namespace")

	// toml bindings
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_IPC_ENABLED, serveCmd.PersistentFlags().Lookup("ipc-enabled"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_IPC_ENDPOINT, serveCmd.PersistentFlags().Lookup("ipc-path"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_HTTP_ENABLED, serveCmd.PersistentFlags().Lookup("http-enabled"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_HTTP_ENDPOINT, serveCmd.PersistentFlags().Lookup("http-path"))

	viper.BindPFlag(leveldb_ethdb_rpc.TOML_LEVELDB_PATH, serveCmd.PersistentFlags().Lookup("leveldb-path"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_LEVELDB_CACHE_SIZE, serveCmd.PersistentFlags().Lookup("leveldb-cache-size"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_LEVELDB_ANCIENT_PATH, serveCmd.PersistentFlags().Lookup("leveldb-ancient-path"))
	viper.BindPFlag(leveldb_ethdb_rpc.TOML_LEVELDB_NAMESPACE, serveCmd.PersistentFlags().Lookup("leveldb-namespace"))
}
