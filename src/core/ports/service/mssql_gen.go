package service

import "Db-Generator/src/core/domain/parameters"

type GenerateMSSQLScriptsServiceInterface interface {
	Generate(params parameters.DbGenParameters) string
}
