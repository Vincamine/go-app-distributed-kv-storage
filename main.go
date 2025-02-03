package main

import (
	"flag"
	"github.com/Vincamine/go-app-distributed-kv-storage/config"
)

var (
)

// Server side
func main() {
	flag.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	config.GetConfig()
	config.DefaultConfig.Wal.Store.Path = *storeFilePath
	config.DefaultConfig.Wal.Restore.Path = *restoreFilePath
	options.DefaultOption()
	dbServer := server.NewServer()
	go route.RouterStart(dbServer)
	raft.RaftServer(dbServer)
}