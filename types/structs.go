package types

type Newsletter struct {
	EmailAddress string `json:"email_address"`
}

type HTTPResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}

type DBSettings struct {
	DatabaseUsername string
	DatabasePassword string
	DatabaseNet      string
	DatabaseHost     string
	DatabaseName     string
}
