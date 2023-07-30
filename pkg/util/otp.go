package util

import (
	"fmt"
	"time"
)

func OtpGenerator() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
