CREATE DATABASE DogToyStore ;
GO 
USE DogToyStore ;
GO 
CREATE TABLE Products (
	Id nvarchar(255) not null primary key, 
	ProductName nvarchar(255) null , 
	Price decimal null , 
	isActive bit not null , 
	CreatedDate datetime null  
)
GO
CREATE TABLE Branches (
	ID nvarchar(255) not null primary key, 
	BranchName nvarchar(255) null  
)
GO
