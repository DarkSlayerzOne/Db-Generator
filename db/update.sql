
-- Section for Products 
CREATE PROC sp_update_products
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
		
			IF NOT EXISTS(SELECT 1 FROM Products WHERE Id=@Id) 
				BEGIN
					SET @StatusCode = 404
					SET @Message = 'Notfound'
		
					SELECT @StatusCode as code, @Message as message
				RETURN
				END
			ELSE
		
			BEGIN
				UPDATE Products SET 
							CreatedDate=@CreatedDate 

						WHERE Id=@Id

						SET @StatusCode = 200
						SET @Message = 'Succesflly update ' + @Id
			
						SELECT @StatusCode as code, @Message as message
					  RETURN
					END
			END
			GO
		
		
-- Section for Branches 
CREATE PROC sp_update_branches
	@ID nvarchar(255), 
	@BranchName nvarchar(255) 
AS
		BEGIN
			SET NOCOUNT ON;

			DECLARE @StatusCode int
			DECLARE @Message nvarchar(100)
		
			IF NOT EXISTS(SELECT 1 FROM Branches WHERE ID=@ID) 
				BEGIN
					SET @StatusCode = 404
					SET @Message = 'Notfound'
		
					SELECT @StatusCode as code, @Message as message
				RETURN
				END
			ELSE
		
			BEGIN
				UPDATE Branches SET 
							BranchName=@BranchName 

						WHERE ID=@ID

						SET @StatusCode = 200
						SET @Message = 'Succesflly update ' + @ID
			
						SELECT @StatusCode as code, @Message as message
					  RETURN
					END
			END
			GO
		
		