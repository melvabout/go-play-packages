package module

import (
	"fmt"
	"time"
)

func TheTimeIs() string {
	return fmt.Sprintf("The time is %s", time.Now().Format(time.RFC1123))
}
