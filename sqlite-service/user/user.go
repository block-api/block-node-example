package user

import (
	"github.com/block-api/block-node/block"
	"github.com/block-api/block-node/log"
	"github.com/block-api/block-node/transporter"
)

type ServiceUserBlock struct {
	block.Block
}

func NewUserBlock(bn *block.BlockNode) ServiceUserBlock {
	serviceBlock := ServiceUserBlock{
		block.NewBlock(bn, "user"),
	}

	serviceBlock.AddAction("add", serviceBlock.ActionAddUser)
	serviceBlock.AddAction("get", serviceBlock.ActionGetUser)
	serviceBlock.AddAction("delete", serviceBlock.ActionDeleteUser)

	return serviceBlock
}

func (ab *ServiceUserBlock) ActionAddUser(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	var response = new(transporter.PayloadMessage)

	s3db := ab.BlockNode().Database().GetSQLite("default")

	_, err := s3db.Db.Exec("INSERT INTO user (name) VALUES (?)", payload.Data)
	if err != nil {
		log.Warning(err.Error())
		response.Data = false
		return response, nil
	}

	response.Data = true

	return response, nil
}

func (ab *ServiceUserBlock) ActionGetUser(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	var response = new(transporter.PayloadMessage)

	s3db := ab.BlockNode().Database().GetSQLite("default")

	row := s3db.Db.QueryRow("SELECT id FROM user WHERE name = ?", payload.Data)
	if row.Err() != nil {
		log.Warning(row.Err().Error())
		response.Data = false
		return response, nil
	}

	var id uint64
	_ = row.Scan(&id)

	if id < 1 {
		response.Data = false
		return response, nil
	}

	response.Data = id
	return response, nil
}

func (ab *ServiceUserBlock) ActionDeleteUser(payload transporter.PayloadMessage) (*transporter.PayloadMessage, error) {
	var response = new(transporter.PayloadMessage)

	s3db := ab.BlockNode().Database().GetSQLite("default")

	_, err := s3db.Db.Exec("DELETE FROM user WHERE name = ?", payload.Data)
	if err != nil {
		log.Warning(err.Error())
		response.Data = false
		return response, nil
	}

	response.Data = true

	return response, nil
}
