-- Section for Products 
CREATE PROC sp_delete_by_id_products
@Id nvarchar(255)
		AS
		BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				IF NOT EXISTS(SELECT 1 FROM Products WHERE Id=@Id) 
					BEGIN
						SET @StatusCode = 404
						SET @Message = 'Product Not found'
						SELECT @StatusCode as code, @Message as message
						RETURN
					END
				ELSE
					BEGIN
					  SET @StatusCode = 200
					  SET @Message = 'Succesfully deleted'
			   
					  DELETE FROM Products WHERE Id=@Id
			   
					  SELECT @StatusCode as code, @Message as message
					  RETURN
				END
		END
		GO
		-- Section for Branches 
CREATE PROC sp_delete_by_id_branches
@ID nvarchar(255)
		AS
		BEGIN
				SET NOCOUNT ON;
				DECLARE @StatusCode int
				DECLARE @Message nvarchar(100)

				IF NOT EXISTS(SELECT 1 FROM Branches WHERE ID=@ID) 
					BEGIN
						SET @StatusCode = 404
						SET @Message = 'Product Not found'
						SELECT @StatusCode as code, @Message as message
						RETURN
					END
				ELSE
					BEGIN
					  SET @StatusCode = 200
					  SET @Message = 'Succesfully deleted'
			   
					  DELETE FROM Branches WHERE ID=@ID
			   
					  SELECT @StatusCode as code, @Message as message
					  RETURN
				END
		END
		GO
		