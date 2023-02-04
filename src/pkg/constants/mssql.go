package constants

const (
	MSSQL_date           = "date"
	MSSQL_datetime2      = "datetime2"
	MSSQL_datetime       = "datetime"
	MSSQL_datetimeoffset = "datetimeoffset"
	MSSQL_smalldatetime  = "smalldatetime"
	MSSQL_time           = "time"
)

const (
	MSSQL_char    = "char"
	MSSQL_text    = "text"
	MSSQL_varchar = "varchar"
)

const (
	MSSQL_nchar    = "nchar"
	MSSQL_ntext    = "ntext"
	MSSQL_nvarchar = "nvarchar"
)

const (
	MSSQL_binary    = "binary"
	MSSQL_varbinary = "varbinary"
	MSSQL_image     = "image"
)

const (
	MSSQL_bigint     = "bigint"
	MSSQL_bit        = "bit"
	MSSQL_int        = "int"
	MSSQL_decimal    = "decimal"
	MSSQL_money      = "money"
	MSSQL_numeric    = "numeric"
	MSSQL_smallint   = "smallint"
	MSSQL_smallmoney = "smallmoney"
	MSSQL_tinyint    = "tinyint"
)

func ListOfMSSQLDataType() []string {

	var dataTypes []string
	dataTypes = append(dataTypes, MSSQL_date)
	dataTypes = append(dataTypes, MSSQL_datetime2)
	dataTypes = append(dataTypes, MSSQL_datetime)
	dataTypes = append(dataTypes, MSSQL_datetimeoffset)
	dataTypes = append(dataTypes, MSSQL_smalldatetime)
	dataTypes = append(dataTypes, MSSQL_time)

	dataTypes = append(dataTypes, MSSQL_char)
	dataTypes = append(dataTypes, MSSQL_text)
	dataTypes = append(dataTypes, MSSQL_varchar)

	dataTypes = append(dataTypes, MSSQL_nchar)
	dataTypes = append(dataTypes, MSSQL_ntext)
	dataTypes = append(dataTypes, MSSQL_nvarchar)

	dataTypes = append(dataTypes, MSSQL_binary)
	dataTypes = append(dataTypes, MSSQL_varbinary)
	dataTypes = append(dataTypes, MSSQL_image)

	dataTypes = append(dataTypes, MSSQL_bigint)
	dataTypes = append(dataTypes, MSSQL_bit)
	dataTypes = append(dataTypes, MSSQL_int)
	dataTypes = append(dataTypes, MSSQL_decimal)
	dataTypes = append(dataTypes, MSSQL_money)
	dataTypes = append(dataTypes, MSSQL_numeric)
	dataTypes = append(dataTypes, MSSQL_smallint)
	dataTypes = append(dataTypes, MSSQL_smallmoney)
	dataTypes = append(dataTypes, MSSQL_tinyint)

	return dataTypes
}
