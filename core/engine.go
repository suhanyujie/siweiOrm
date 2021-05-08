package core

import (
	"database/sql"
	"github.com/suhanyujie/siweiOrm/log"
	"github.com/suhanyujie/siweiOrm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver string, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// 发送 ping 使数据库连接保持活跃
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{
		db: db,
	}
	log.Info("Connection database success.")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Errorf("Failed to close database: %s", err)
		return
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
