package pkg

const (
	POST   string = "POST"
	GET    string = "GET"
	PUT    string = "PUT"
	DELETE string = "DELETE"
	PATCH  string = "PATCH"
)

const (
	OK            int = 200
	BadRequest    int = 400
	UnAuthorized  int = 401
	ServerError   int = 500
	Created       int = 201
	Conflict      int = 409
	NotFound      int = 404
	NoContent     int = 204
	Unprocessable int = 422
)

