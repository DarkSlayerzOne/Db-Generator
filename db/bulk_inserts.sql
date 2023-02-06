
-- Section for Products 
CREATE PROC sp_bulk_insert_products
@ProductsJson nvarchar(MAX)
AS
	BEGIN
				SET NOCOUNT ON;

				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				INSERT INTO Products
				SELECT 	Id, 
						ProductName, 
						Price, 
						isActive, 
						CreatedDate, 
						CreatedBy 

				FROM OPENJSON(@ProductsJson)
				WITH(
					Id nvarchar '$.id',
					ProductName nvarchar '$.productName',
					Price (16,8) '$.price',
					isActive bit '$.isActive',
					CreatedDate datetime '$.createdDate',
					CreatedBy nvarchar '$.createdBy'

				)
				SELECT @StatusCode, @Message
	END
		
-- Section for Branches 
CREATE PROC sp_bulk_insert_branches
@BranchesJson nvarchar(MAX)
AS
	BEGIN
		SET NOCOUNT ON;

		DECLARE @StatusCode int
		DECLARE @Message nvarchar(100)

		INSERT INTO Branches
		SELECT 	ID, 
				BranchName, 
				Manager, 
				Manager 

		FROM OPENJSON(@BranchesJson)
		WITH(
			ID nvarchar '$.iD',
			BranchName nvarchar '$.branchName',
			Manager nvarchar '$.manager',
			Manager nvarchar '$.manager'

		)
		SELECT @StatusCode, @Message
	END
		