-- Section for Products 
CREATE PROC sp_read_products

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
				Id 'id', ProductName 'productName', Price 'price', isActive 'isActive', CreatedDate 'createdDate'
				-- Please change the order by to an available column in the table
				FROM Products ORDER BY CreatedDate DESC
				OFFSET @Row * (@Offset -1) ROWS
				FETCH NEXT @Row ROWS ONLY 
				FOR JSON PATH )

			 SET @StatusCode = 200
			 SELECT @StatusCode as code, @json as json

			SELECT @Total=COUNT(*) FROM Products
		END
		GO
		-- Section for Branches 
CREATE PROC sp_read_branches

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
				ID 'iD', BranchName 'branchName'
				-- Please change the order by to an available column in the table
				FROM Branches ORDER BY '' DESC
				OFFSET @Row * (@Offset -1) ROWS
				FETCH NEXT @Row ROWS ONLY 
				FOR JSON PATH )

			 SET @StatusCode = 200
			 SELECT @StatusCode as code, @json as json

			SELECT @Total=COUNT(*) FROM Branches
		END
		GO
		