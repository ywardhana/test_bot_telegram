package service

import (
  "github.com/bukalapak/kingsman/app/entity"
  "github.com/jmoiron/sqlx"
)

func GetInternalTransaction(db *sqlx.DB, agentID int64) (entity.Agent, entity.CustomError) {
  var agent entity.Agent
  db.QueryRowx("SELECT * FROM agents WHERE id=?", agentID).StructScan(&agent)

  if agent.ID == 0 {
    return agent, entity.AgentNotFoundError
  }

  return agent, entity.CustomError{}
}
