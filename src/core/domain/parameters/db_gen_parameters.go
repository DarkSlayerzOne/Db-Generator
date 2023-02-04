package parameters

type DbGenParameters struct {
	DbName string  `json:"dbName"`
	Table  []Table `json:"tables"`
	DbApp  string  `json:"dbApp"`
}

type Table struct {
	TableName        string   `json:"tableName"`
	FieldCollections []Fields `json:"fields"`
}

type Fields struct {
	FieldName   string     `json:"fieldName"`
	DataType    string     `json:"dataType"`
	IsPrimayKey bool       `json:"isPrimaryKey"`
	IsNullable  bool       `json:"isNullable"`
	Length      int        `json:"length"`
	Precisions  Precisions `json:"precisions"`
}

type Precisions struct {
	FirstPresc  int `json:"firstPresc"`
	SecondPresc int `json:"secondPresc"`
}
