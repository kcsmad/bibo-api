package db

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	Client  *sql.DB
	Context *context.Context
}

var repositoryLogger = log.WithFields(log.Fields{
	"[domain]": "[Repository]",
})

func (r *Repository) Connect(dbURL string) {
	if r.Client != nil {
		return
	}

	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		repositoryLogger.WithFields(log.Fields{
			"method": "[Connect]",
			"action": "[Open SQL Connection]",
		}).Fatal(err.Error())
	}

	r.Client = db
}

func (r *Repository) Disconnect() {
	if err := r.Client.Close(); err != nil {
		repositoryLogger.WithFields(log.Fields{
			"method": "[Disconnect]",
			"action": "[Client Close Connection]",
		}).Fatal(err.Error())
	}
}
