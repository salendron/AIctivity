package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func newBool(b bool) *bool {
	return &b
}

func newInt(i int) *int {
	return &i
}

type API struct {
	Storage StorageInterface
}

func (api *API) SetStorage(storage StorageInterface) {
	api.Storage = storage
}

func (api *API) SaveData(ctx echo.Context) error {
	d := new(Data)
	if err := ctx.Bind(d); err != nil {
		return RaiseError(ctx, fmt.Sprintf("Failed to parse request body. %v", err.Error()), http.StatusBadRequest, ErrorCodeInvalidRequestBody)
	}

	if err := api.Storage.SaveData(d); err != nil {
		log.Errorf("Failed to save data. %v", err.Error())
		return err
	}

	return ctx.String(http.StatusCreated, "ok")
}
