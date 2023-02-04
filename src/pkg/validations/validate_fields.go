package validations

import (
	"Db-Generator/src/core/domain/parameters"
	"Db-Generator/src/pkg/constants"
	customErrors "Db-Generator/src/pkg/errors"

	"fmt"
)

func ValidateFields(params parameters.DbGenParameters) ([]string, error) {

	var errorList []string

	for _, table := range params.Table {

		if table.TableName == "" {
			errorList = append(errorList, "Table name is required!")
		}

		for _, field := range table.FieldCollections {

			if field.FieldName == "" {
				errorList = append(errorList, "Field name is required!")
			}

			isFound := false
			for _, dbTypes := range constants.ListOfMSSQLDataType() {
				if field.DataType == dbTypes {
					isFound = true
				}
			}

			if !isFound {
				errorList = append(errorList, fmt.Sprintf("Invalid data type %s for %s", field.DataType, field.FieldName))
			}
		}

	}

	if len(errorList) > 0 {
		errs := customErrors.InvalidParamsErrorBiddingErrorWrapper()
		return errorList, errs
	}

	return nil, nil
}
