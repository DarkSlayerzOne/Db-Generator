CREATE DATABASE Inventory ;
GO 
USE Inventory ;
GO 
CREATE TABLE Products (
	Id nvarchar(255) not null primary key, 
	ProductName nvarchar(255) null , 
	Price decimal(16,8), 
	isActive bit not null , 
	CreatedDate datetime null , 
	CreatedBy nvarchar null  
)
GO
CREATE TABLE Branches (
	ID nvarchar(255) not null primary key, 
	BranchName nvarchar(255) null , 
	Manager nvarchar(555) null , 
	Manager nvarchar(555) null  
)
GO
