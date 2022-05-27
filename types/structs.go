package types

type Newsletter struct {
	EmailAddress string
}

type HTTPResponse struct {
	Success  bool
	Response interface{}
}
