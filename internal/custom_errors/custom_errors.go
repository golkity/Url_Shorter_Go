package custom_errors

import "errors"

var (
	Error_InitRepository = errors.New("failed to initialize repository")
	Error_StartServer    = errors.New("failed to start server")
)
