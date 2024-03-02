package logs

import (
	"fmt"
	"os"
	"time"
)

func Examples_3_14_a() {
	fmt.Fprintf(os.Stderr, "%s: %s", time.Now(), "I'm here!")
}
