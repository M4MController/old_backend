package error_codes

const (
	OK                      = iota
	NOT_LOGGED
	INVALID_AUTH_DATA
	INVALID_OWNER
	SESSION_ERROR
	DATABASE_ERROR
	DATABASE_INVALID_SENSOR
	ENCODING_ERROR
	INVALID_ARGUMENT
	INVALID_QUERY
	INVALID_BODY_READ
	INVALID_BODY_CLOSE
	UNMARSHAL_ERROR
	VALIDATION_FAILED
	LOGIN_FAILED
)