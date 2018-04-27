package service

import (
	"github.com/bukalapak/kingsman/app/entity"
	"github.com/jmoiron/sqlx"
)

func Get(db *sqlx.DB, ID int64, table string) (entity.Agent, entity.CustomError) {
	var agent entity.Agent
	db.QueryRowx("SELECT * FROM "+table+" WHERE id=?", ID).StructScan(&agent)

	if agent.ID == 0 {
		return agent, entity.AgentNotFoundError
	}

	return agent, entity.CustomError{}
}
