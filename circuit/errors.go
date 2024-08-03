package circuit

import (
	"errors"
)

var (
	ErrInvalidThresholdValue = errors.New("invalid threshold value")
	ErrThresholdExceeded     = errors.New("threshold exceeded")
)
