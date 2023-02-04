
-- Section for Products 
CREATE PROC sp_add_products
	@Id nvarchar(255), 
	@ProductName nvarchar(255), 
	@Price decimal, 
	@isActive bit, 
	@CreatedDate datetime 
AS
			BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				INSERT INTO Products	
				(	Id, 
	ProductName, 
	Price, 
	isActive, 
	CreatedDate 

							)
							VALUES
						   (
	@Id, 
	@ProductName, 
	@Price, 
	@isActive, 
	@CreatedDate 

							)
			 SET @StatusCode = 201
			 SET @Message = 'Sucesfully created Products'

			 SELECT @StatusCode, @Message

		END
		GO
		
-- Section for Branches 
CREATE PROC sp_add_branches
	@ID nvarchar(255), 
	@BranchName nvarchar(255) 
AS
			BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				INSERT INTO Branches	
				(	ID, 
	BranchName 

							)
							VALUES
						   (
	@ID, 
	@BranchName 

							)
			 SET @StatusCode = 201
			 SET @Message = 'Sucesfully created Branches'

			 SELECT @StatusCode, @Message

		END
		GO
		