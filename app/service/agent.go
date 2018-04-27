package service

import (
	"github.com/bukalapak/kingsman/app/entity"
	"github.com/bukalapak/kingsman/db"
	"github.com/jmoiron/sqlx"
)

func GetAgent(database *sqlx.DB, agentID int64) (entity.Agent, entity.CustomError) {
	var agent entity.Agent
	query, _ := db.Get(database, "agents", agentID)
	query.StructScan(&agent)

	if agent.ID == 0 {
		return agent, entity.AgentNotFoundError
	}

	return agent, entity.CustomError{}
}
