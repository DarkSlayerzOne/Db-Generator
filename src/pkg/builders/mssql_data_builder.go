package builders

import (
	"Db-Generator/src/core/domain/parameters"
	"Db-Generator/src/pkg/constants"
	"fmt"
	"strings"
)

var data_type_mssql parameters.Fields

type FieldContruct struct {
	Field string
}

type MSSQLFieldBuilder struct {
	FieldContruct
}

func (mssql *MSSQLFieldBuilder) GetField(field parameters.Fields) *MSSQLFieldBuilder {
	data_type_mssql = field
	return mssql
}

func (mssql *MSSQLFieldBuilder) BuildNVarchar() *MSSQLFieldBuilder {

	if data_type_mssql.DataType == constants.MSSQL_nvarchar {

		var sb strings.Builder

		var allowNullable string = ""
		var hasPrimaryKey string = ""
		var hasLength string = ""

		if data_type_mssql.IsNullable {
			allowNullable = "not null"
		}

		if data_type_mssql.IsPrimayKey {
			hasPrimaryKey = "primary key"
		}

		if data_type_mssql.Length > 0 {
			hasLength = fmt.Sprintf("(%d)", data_type_mssql.Length)
		}

		sb.WriteString(fmt.Sprintf("%s%s%s %s %s, \n", data_type_mssql.FieldName, data_type_mssql.DataType, hasLength, allowNullable, hasPrimaryKey))

		strLength := len(sb.String())

		mssql.Field = sb.String()[0 : strLength-1]

	}

	return mssql
}

func (mssql *MSSQLFieldBuilder) BuildDecimal() *MSSQLFieldBuilder {

	if data_type_mssql.DataType == constants.MSSQL_decimal {

		var sb strings.Builder

		var allowNullable string = ""
		var hasPrimaryKey string = ""
		var hasLength string = ""

		if data_type_mssql.IsNullable {
			allowNullable = "not null"
		}

		if data_type_mssql.IsPrimayKey {
			hasPrimaryKey = "primary key"
		}

		if data_type_mssql.Length > 0 {
			hasLength = fmt.Sprintf("(%d)", data_type_mssql.Length)
		}

		sb.WriteString(fmt.Sprintf("%s%s%s %s %s, \n", data_type_mssql.FieldName, data_type_mssql.DataType, hasLength, allowNullable, hasPrimaryKey))

		strLength := len(sb.String())

		mssql.Field = sb.String()[0 : strLength-1]

	}

	return mssql
}
