package handlers

import (
	"Db-Generator/src/core/domain/parameters"
	"Db-Generator/src/core/domain/response"
	"Db-Generator/src/core/ports/service"
	"Db-Generator/src/pkg"
	"Db-Generator/src/pkg/validations"
	"encoding/json"
	"net/http"
)

type HandlerInterface interface {
	GenerateDb(r http.ResponseWriter, request *http.Request)
}

var (
	_genDBService service.GenerateMSSQLScriptsServiceInterface
)

type handler struct{}

func HandlerInterfaceImple(genDBService service.GenerateMSSQLScriptsServiceInterface) HandlerInterface {
	_genDBService = genDBService
	return &handler{}
}

func (*handler) GenerateDb(r http.ResponseWriter, request *http.Request) {

	var dbGenParams parameters.DbGenParameters

	err := json.NewDecoder(request.Body).Decode(&dbGenParams)

	if err != nil {
		r.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(r).Encode(response.APIResponse{StatusCode: pkg.Unprocessable, Message: "Error unmarshalling data"})
		return
	}

	validations, errs := validations.ValidateFields(dbGenParams)

	if errs != nil {
		r.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(r).Encode(response.APIResponse{StatusCode: pkg.BadRequest, Message: errs.Error(), Data: validations})
		return
	}

	_genDBService.Generate(dbGenParams)

	resultResponse := response.APIResponse{
		StatusCode: pkg.Created,
		Message:    "Successfully created.",
		Data:       nil,
	}

	r.WriteHeader(http.StatusCreated)
	json.NewEncoder(r).Encode(resultResponse)
}
