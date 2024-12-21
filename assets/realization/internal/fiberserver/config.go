package fiberserver

type Config struct {
	//StopTimeout                 utils.DurationSec
	Host                        string `validate:"required"`
	ShowUnknownErrorsInResponse bool
	AllowOrigins                string   `validate:"required"`
	AllowHeaders                string   `validate:"required"`
	ExposeHeaders               string   `validate:"required"`
	IpHeader                    string   `validate:"required"`
	SecureReqJsonPaths          []string `validate:"required"`
	SecureResJsonPaths          []string `validate:"required"`
	BodyLimit                   int      `validate:"required"`
	StreamRequestBody           *bool    `validate:"required"`
}
