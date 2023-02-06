package service

import (
	"Db-Generator/src/core/domain/parameters"
	"Db-Generator/src/core/ports/service"
	"Db-Generator/src/pkg/constants"
	"fmt"
	"os"
	"strings"
)

type mssql_service struct{}

func GenerateMSSQLScriptsServiceInterfaceImpl() service.GenerateMSSQLScriptsServiceInterface {
	return &mssql_service{}
}

func (*mssql_service) Generate(params parameters.DbGenParameters) string {
	generateDbAndTables(params)
	generateInsert(params)
	generateBulkInserts(params)
	generateReadADataByID(params)
	generateReadWithPaging(params)
	generateDelete(params)
	generateUpdate(params)
	return ""
}

func generateDbAndTables(params parameters.DbGenParameters) string {

	var sb strings.Builder
	sb.WriteString(fmt.Sprintln("CREATE DATABASE", params.DbName, ";"))
	sb.WriteString("GO \n")
	sb.WriteString(fmt.Sprintln("USE", params.DbName, ";"))
	sb.WriteString("GO \n")
	sb.WriteString("")

	var tableStringBuilder strings.Builder

	for _, table := range params.Table {

		tableStringBuilder.WriteString(fmt.Sprintln("CREATE TABLE", table.TableName, "("))

		var countFields int = 0
		fields := len(table.FieldCollections)

		for _, field := range table.FieldCollections {

			tableStringBuilder.WriteString(createField(field, countFields, fields, true))
			countFields++

		}

		tableStringBuilder.WriteString(fmt.Sprintln(")"))
		tableStringBuilder.WriteString(fmt.Sprintln("GO"))

	}

	sb.WriteString(tableStringBuilder.String())

	filePath := "db/db_and_tables.sql"

	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println("File deleted successfully.")
	}

	file, err := os.Create("db/db_and_tables.sql")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	file.WriteString(sb.String())

	return sb.String()
}

func generateInsert(params parameters.DbGenParameters) {
	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))
		sb.WriteString(fmt.Sprintf("CREATE PROC sp_add_%s\n", strings.ToLower(table.TableName)))

		var countFields int = 0
		fields := len(table.FieldCollections)

		for _, field := range table.FieldCollections {

			sb.WriteString(createField(field, countFields, fields, false))
			countFields++

		}
		sb.WriteString(fmt.Sprintf(`AS
			BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				INSERT INTO %s	
				(`, table.TableName))

		var countInsert int = 0
		for _, field := range table.FieldCollections {

			if countInsert == fields-1 {
				sb.WriteString(fmt.Sprintf("\t%s \n", field.FieldName))
			} else {
				sb.WriteString(fmt.Sprintf("\t%s, \n", field.FieldName))
			}
			countInsert++

		}
		sb.WriteString(fmt.Sprintln(`
							)
							VALUES
						   (`))

		var countParams int = 0
		for _, field := range table.FieldCollections {

			if countParams == fields-1 {
				sb.WriteString(fmt.Sprintf("\t@%s \n", field.FieldName))
			} else {
				sb.WriteString(fmt.Sprintf("\t@%s, \n", field.FieldName))
			}
			countParams++

		}
		sb.WriteString(fmt.Sprintf(`
							)
			 SET @StatusCode = 201
			 SET @Message = 'Sucesfully created %s'

			 SELECT @StatusCode, @Message

		END
		GO
		`, table.TableName))
	}

	filePath := "db/inserts.sql"
	generateFile(&sb, filePath)
}

func generateBulkInserts(params parameters.DbGenParameters) {

	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))
		sb.WriteString(fmt.Sprintf("CREATE PROC sp_bulk_insert_%s\n", strings.ToLower(table.TableName)))
		sb.WriteString(fmt.Sprintf("@%sJson nvarchar(MAX)\n", table.TableName))

		var fieldsForInsert string = ""

		fields := len(table.FieldCollections)
		var countInsert int = 0
		for _, field := range table.FieldCollections {

			if countInsert == fields-1 {
				fieldsForInsert += fmt.Sprintf("\t%s \n", field.FieldName)
			} else {
				fieldsForInsert += fmt.Sprintf("\t%s, \n", field.FieldName)
			}
			countInsert++

		}

		var jsonifyFields string = ""

		var countFields int = 0

		for _, e := range table.FieldCollections {

			transformToCamelCase := fmt.Sprintf("%s%s", strings.ToLower(string(e.FieldName[0])), e.FieldName[1:])

			if countFields == fields-1 {

				if e.DataType == constants.MSSQL_decimal || e.DataType == constants.MSSQL_numeric {
					jsonifyFields += fmt.Sprintf("%s %s '$.%s'\n", e.FieldName, "(16,8)", transformToCamelCase)

				} else {
					jsonifyFields += fmt.Sprintf("%s %s '$.%s'\n", e.FieldName, e.DataType, transformToCamelCase)
				}

			} else {

				if e.DataType == constants.MSSQL_decimal || e.DataType == constants.MSSQL_numeric {
					jsonifyFields += fmt.Sprintf("%s %s '$.%s',\n ", e.FieldName, "(16,8)", transformToCamelCase)
				} else {
					jsonifyFields += fmt.Sprintf("%s %s '$.%s',\n ", e.FieldName, e.DataType, transformToCamelCase)

				}

			}

			countFields++
		}

		sb.WriteString(fmt.Sprintf(`
			AS
			BEGIN

				SET NOCOUNT ON;

				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				INSERT INTO %s
				SELECT %s
				FROM OPENJSON(@%sJson)
				WITH(
					%s
				)
				SELECT @StatusCode, @Message
			END
		`, table.TableName, fieldsForInsert, table.TableName, jsonifyFields))

	}

	filePath := "db/bulk_inserts.sql"
	generateFile(&sb, filePath)

}

func generateReadADataByID(params parameters.DbGenParameters) {
	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))

		sb.WriteString(fmt.Sprintf("CREATE PROC sp_read_by_id_%s\n", strings.ToLower(table.TableName)))

		var primaryKey string = ""

		var jsonifyFields string = ""

		var countFields int = 0
		sumOfFields := len(table.FieldCollections)

		for _, e := range table.FieldCollections {

			if e.IsPrimayKey {
				sb.WriteString(fmt.Sprintf("@%s %s(%d)", e.FieldName, e.DataType, e.Length))
				primaryKey = e.FieldName
			}

			transformToCamelCase := fmt.Sprintf("%s%s", strings.ToLower(string(e.FieldName[0])), e.FieldName[1:])

			if countFields == sumOfFields-1 {
				jsonifyFields += fmt.Sprintf("%s '%s'", e.FieldName, transformToCamelCase)
			} else {
				jsonifyFields += fmt.Sprintf("%s '%s', ", e.FieldName, transformToCamelCase)
			}

			countFields++
		}

		sb.WriteString(fmt.Sprintf(`
		AS
		BEGIN
				SET NOCOUNT ON;

				IF EXISTS(SELECT 1 FROM %s WHERE %s=@%s) 
				BEGIN
					SET @StatusCode = 200
					SET @json=(	SELECT %s FROM %s WHERE %s=@%s)
					
					SET @StatusCode=200
					SELECT @StatusCode as code, @json as json
					RETURN
				END
			ELSE
				BEGIN
					SET @StatusCode = 404
					SELECT @StatusCode as code, '' as json
					RETURN
				END
		END
		GO
		`, table.TableName, primaryKey, primaryKey, jsonifyFields, table.TableName, primaryKey, primaryKey))

	}

	filePath := "db/read_by_ids.sql"
	generateFile(&sb, filePath)

}

func generateReadWithPaging(params parameters.DbGenParameters) {

	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))
		sb.WriteString(fmt.Sprintf("CREATE PROC sp_read_%s\n", strings.ToLower(table.TableName)))

		var countFields int = 0
		sumOfFields := len(table.FieldCollections)
		var jsonifyFields string = ""

		var orderByField string = "''"

		for _, e := range table.FieldCollections {

			transformToCamelCase := fmt.Sprintf("%s%s", strings.ToLower(string(e.FieldName[0])), e.FieldName[1:])

			if countFields == sumOfFields-1 {
				jsonifyFields += fmt.Sprintf("%s '%s'", e.FieldName, transformToCamelCase)
			} else {
				jsonifyFields += fmt.Sprintf("%s '%s', ", e.FieldName, transformToCamelCase)
			}

			if e.DataType == constants.MSSQL_datetime || e.DataType == constants.MSSQL_datetime2 {
				orderByField = e.FieldName
			}

			countFields++
		}

		sb.WriteString(fmt.Sprintf(`
		@Offset int,
		@Row int,
		@Total int output
		AS
		BEGIN
			SET NOCOUNT ON;

			DECLARE @StatusCode int
			DECLARE @json nvarchar(MAX)

			SET @json =(
				
			  	SELECT 
				%s
				-- Please change the order by to an available column in the table
				FROM %s ORDER BY %s DESC
				OFFSET @Row * (@Offset -1) ROWS
				FETCH NEXT @Row ROWS ONLY 
				FOR JSON PATH )

			 SET @StatusCode = 200
			 SELECT @StatusCode as code, @json as json

			SELECT @Total=COUNT(*) FROM %s
		END
		GO
		`, jsonifyFields, table.TableName, orderByField, table.TableName))

	}

	filePath := "db/read.sql"
	generateFile(&sb, filePath)

}

func generateDelete(params parameters.DbGenParameters) {
	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))

		sb.WriteString(fmt.Sprintf("CREATE PROC sp_delete_by_id_%s\n", strings.ToLower(table.TableName)))

		var primaryKey string = ""

		for _, e := range table.FieldCollections {

			if e.IsPrimayKey {
				sb.WriteString(fmt.Sprintf("@%s %s(%d)", e.FieldName, e.DataType, e.Length))
				primaryKey = e.FieldName
			}
		}

		sb.WriteString(fmt.Sprintf(`
		AS
		BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				IF NOT EXISTS(SELECT 1 FROM %s WHERE %s=@%s) 
					BEGIN
						SET @StatusCode = 404
						SET @Message = 'Product Not found'
						SELECT @StatusCode as code, @Message as message
						RETURN
					END
				ELSE
					BEGIN
					  SET @StatusCode = 200
					  SET @Message = 'Succesfully deleted'
			   
					  DELETE FROM %s WHERE %s=@%s
			   
					  SELECT @StatusCode as code, @Message as message
					  RETURN
				END
		END
		GO
		`, table.TableName, primaryKey, primaryKey, table.TableName, primaryKey, primaryKey))

	}

	filePath := "db/delete.sql"
	generateFile(&sb, filePath)

}

func generateUpdate(params parameters.DbGenParameters) {
	var sb strings.Builder

	for _, table := range params.Table {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("-- Section for %s \n", table.TableName))
		sb.WriteString(fmt.Sprintf("CREATE PROC sp_update_%s\n", strings.ToLower(table.TableName)))

		var countFields int = 0
		fields := len(table.FieldCollections)

		var primaryKey string = ""

		for _, field := range table.FieldCollections {

			sb.WriteString(createField(field, countFields, fields, false))
			countFields++

			if field.IsPrimayKey {
				primaryKey = field.FieldName
			}
		}
		sb.WriteString(fmt.Sprintf(`AS
		BEGIN
			SET NOCOUNT ON;

			DECLARE @StatusCode int
			DECLARE @Message nvarchar(100)
		
			IF NOT EXISTS(SELECT 1 FROM %s WHERE %s=@%s) 
				BEGIN
					SET @StatusCode = 404
					SET @Message = 'Notfound'
		
					SELECT @StatusCode as code, @Message as message
				RETURN
				END
			ELSE`, table.TableName, primaryKey, primaryKey))

		var countUpdate int = 0
		var updateFields string = ""
		for _, field := range table.FieldCollections {

			if countUpdate == fields-1 {
				updateFields = fmt.Sprintf("\t%s=@%s \n", field.FieldName, field.FieldName)
			} else {
				updateFields = fmt.Sprintf("\t%s=@%s, \n", field.FieldName, field.FieldName)
			}
			countUpdate++
		}
		sb.WriteString(fmt.Sprintf(`
		
			BEGIN
				UPDATE %s SET 
						%s
						WHERE %s=@%s

						SET @StatusCode = 200
						SET @Message = 'Succesflly update ' + @%s
			
						SELECT @StatusCode as code, @Message as message
					  RETURN
					END
			END
			GO
		
		`, table.TableName, updateFields, primaryKey, primaryKey, primaryKey))
	}

	filePath := "db/update.sql"
	generateFile(&sb, filePath)
}

func createField(field parameters.Fields, countFields, fields int, isTable bool) string {

	var columns strings.Builder

	var allowNullable string = ""
	var hasPrimaryKey string = ""
	var hasLength string = ""

	if field.IsNullable && isTable {
		allowNullable = "null"
	} else {
		allowNullable = "not null"
	}

	if field.IsPrimayKey && isTable {
		hasPrimaryKey = "primary key"
	}

	if field.Length > 0 {
		hasLength = fmt.Sprintf("(%d)", field.Length)
	}

	// TODO : Too much complexity
	f := fields - 1
	if isTable {
		fieldsNonParameters(&columns, countFields, f, &field, &hasLength, &allowNullable, &hasPrimaryKey)
	} else {

		fieldsWithParameters(&columns, countFields, f, &field, &hasLength, &allowNullable, &hasPrimaryKey)
	}

	return columns.String()
}

func fieldsWithParameters(columns *strings.Builder, counFields, fieldsTotal int, field *parameters.Fields, hasLength, allowNullable, hasPrimaryKey *string) {
	if counFields == fieldsTotal {

		if field.DataType == constants.MSSQL_decimal || field.DataType == constants.MSSQL_numeric {
			columns.WriteString(fmt.Sprintf("\t@%v %v%v \n", field.FieldName, field.DataType, "(16,8)"))
		} else {
			columns.WriteString(fmt.Sprintf("\t@%v %v%v \n", field.FieldName, field.DataType, *hasLength))
		}

	} else {

		if field.DataType == constants.MSSQL_decimal || field.DataType == constants.MSSQL_numeric {
			columns.WriteString(fmt.Sprintf("\t@%v %v%v \n", field.FieldName, field.DataType, "(16,8),"))
		} else {
			columns.WriteString(fmt.Sprintf("\t@%v %v%v, \n", field.FieldName, field.DataType, *hasLength))
		}

	}
}

func fieldsNonParameters(columns *strings.Builder, counFields, fieldsTotal int, field *parameters.Fields, hasLength, allowNullable, hasPrimaryKey *string) {
	if counFields == fieldsTotal {

		if field.DataType == constants.MSSQL_decimal || field.DataType == constants.MSSQL_numeric {
			columns.WriteString(fmt.Sprintf("\t%v %v%v \n", field.FieldName, field.DataType, "(16,8)"))
		} else {
			columns.WriteString(fmt.Sprintf("\t%v %v%v %v %v \n", field.FieldName, field.DataType, *hasLength, *allowNullable, *hasPrimaryKey))
		}

	} else {

		if field.DataType == constants.MSSQL_decimal || field.DataType == constants.MSSQL_numeric {
			columns.WriteString(fmt.Sprintf("\t%v %v%v \n", field.FieldName, field.DataType, "(16,8),"))
		} else {
			columns.WriteString(fmt.Sprintf("\t%v %v%v %v %v, \n", field.FieldName, field.DataType, *hasLength, *allowNullable, *hasPrimaryKey))
		}

	}
}

func generateFile(sb *strings.Builder, fileName string) {
	filePath := fileName

	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("File deleted successfully.")
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(sb.String())

}
