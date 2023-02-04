# Db-Generator
<br/>

### Description
This my script to generate db scripts for MSSQL. 

It can create
- Database 
- Tables
- Basic CRUD Store procedures

<br/>

### Run the project 

Clone the project
```
git clone https://github.com/DarkSlayerzOne/Db-Generator.git
```

Go to directory

```
cd Db-Generator
```


Run
```
go run main.go
```

Endpoint : localhost:8100/gen/mssql

Sample Prameters

```
{
    "dbName" : "DogToyStore",
     "tables" : [
         {
             "tableName" : "Products",
             "fields" : [
                 {
                     "fieldName" : "Id",
                     "dataType" : "nvarchar",
                     "isPrimaryKey" : true,
                     "isNullable" : false,
                     "length" : 255
                 },
                 {
                     "fieldName" : "ProductName",
                     "dataType" : "nvarchar",
                     "isPrimaryKey" : false,
                     "isNullable" : true,
                     "length" : 255
                 },
                 {
                     "fieldName" : "Price",
                     "dataType" : "decimal",
                     "isPrimaryKey" : false,
                     "isNullable" : true,
                     "length" : 0
                 },
                 {
                     "fieldName" : "isActive",
                     "dataType" : "bit",
                     "isPrimaryKey" : false,
                     "isNullable" : false,
                     "length" : 0
                 },
                 {
                     "fieldName" : "CreatedDate",
                     "dataType" : "datetime",
                     "isPrimaryKey" : false,
                     "isNullable" : true,
                     "length" : 0
                 }
             ]
         },
          {
             "tableName" : "Branches",
             "fields" : [
                 {
                     "fieldName" : "ID",
                     "dataType" : "nvarchar",
                     "isPrimaryKey" : true,
                     "isNullable" : false,
                     "length" : 255
                 },
                 {
                     "fieldName" : "BranchName",
                     "dataType" : "nvarchar",
                     "isPrimaryKey" : false,
                     "isNullable" : true,
                     "length" : 255
                 }
             ]
         }
     ]
}
```

<br/>

TODO: 
 - Add more database support Postgres, MySQL
 - More validations
 - Database Relationship
 - UI