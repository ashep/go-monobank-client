package monobank

import (
	"fmt"
)

var ErrTooManyRequests = fmt.Errorf("too many requests, try again later")
