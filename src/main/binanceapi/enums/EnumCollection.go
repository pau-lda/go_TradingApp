package enums

type RequestType string

const (
	GET    RequestType = "get"
	POST   RequestType = "post"
	PUT    RequestType = "put"
	DELETE RequestType = "delete"
	PATCH  RequestType = "patch"
)
