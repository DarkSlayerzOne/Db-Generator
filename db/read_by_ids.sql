-- Section for Products 
CREATE PROC sp_read_by_id_products
@Id nvarchar(255)
		AS
		BEGIN
				SET NOCOUNT ON;

				IF EXISTS(SELECT 1 FROM Products WHERE Id=@Id) 
				BEGIN
					SET @StatusCode = 200
					SET @json=(	SELECT Id 'id', ProductName 'productName', Price 'price', isActive 'isActive', CreatedDate 'createdDate' FROM Products WHERE Id=@Id)
					
					SET @StatusCode=200
					SELECT @StatusCode as code, @json as json
					RETURN
				END
			ELSE
				BEGIN
					SET @StatusCode = 404
					SELECT @StatusCode as code, '' as json
					RETURN
				END
		END
		GO
		-- Section for Branches 
CREATE PROC sp_read_by_id_branches
@ID nvarchar(255)
		AS
		BEGIN
				SET NOCOUNT ON;

				IF EXISTS(SELECT 1 FROM Branches WHERE ID=@ID) 
				BEGIN
					SET @StatusCode = 200
					SET @json=(	SELECT ID 'iD', BranchName 'branchName' FROM Branches WHERE ID=@ID)
					
					SET @StatusCode=200
					SELECT @StatusCode as code, @json as json
					RETURN
				END
			ELSE
				BEGIN
					SET @StatusCode = 404
					SELECT @StatusCode as code, '' as json
					RETURN
				END
		END
		GO
		