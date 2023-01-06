package controllers

import (
	"erply/entity"
	erply "github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"net/http"
)

func validateUser(con *Controller) (*erply.Client, error) {

	if con.sessionKey == nil || con.clientCode == nil {
		return nil, entity.ErrLoginInfoMissing
	}
	httpCli := http.Client{}
	client, err := erply.NewClient(*con.sessionKey, *con.clientCode, &httpCli)
	if err != nil {
		return nil, entity.ErrFailedEstablishErplyClient
	}
	return client, nil
}

func handleCustomerError(err error) error {
	if erplyError, ok := err.(*common.ErplyError); ok {
		switch erplyError.Code {
		case 1011:
			return entity.ErrCustomerNotFound
		default:
			return err
		}
	} else {
		return err
	}
}
