
// Author -- Ritankar Saha <ritankar.saha786@gmail.com>
// Github -- https://github.com/ritankarsaha

package main

import (
	"database/sql"
	"log"

	"github.com/ritankarsaha/HelloDB/api"
	"github.com/ritankarsaha/HelloDB/chat"
	conv "github.com/ritankarsaha/HelloDB/converter"
	"github.com/ritankarsaha/HelloDB/cron"
	db "github.com/ritankarsaha/HelloDB/internal/database"

	"github.com/ritankarsaha/HelloDB/rag"
	"github.com/ritankarsaha/HelloDB/util"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)

	converter := conv.NewSQLConverter(rag.LLMOpts{
		ApiKey:    config.ApiKey,
		OrgId:     config.OrgId,
		ProjectId: config.ProjectId,
		Model:     config.Model,
		Temp:      config.Temp,
	})

	dbcron := cron.NewDBCron(store, cron.CronConfig{
		BatchSize: config.CronBatchSize,
		LogPath:   config.LogPath,
	})

	err = dbcron.InitCron()
	if err != nil {
		log.Fatal("error initializing database cron", err)
	}

	runGinServer(config, store, converter)
}

func runGinServer(config util.Config, store db.Store, converter conv.Converter) {
	websocketSrv, err := chat.NewWebSocketServer(config, converter)
	if err != nil {
		log.Fatal("couldn't initialize the chat-server:", err)
	}

	server, err := api.NewServer(config, store, websocketSrv)
	if err != nil {
		log.Fatal("couldn't initialize the server:", err)
	}

	err = server.Start(config.Port)
	if err != nil {
		log.Fatal("couldn't start up server:", err)
	}
}
