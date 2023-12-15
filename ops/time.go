package ops

import (
	"fmt"
	"time"
)

func TimeIt(f func()) {
	start := time.Now()
	f()
	fmt.Println(time.Since(start))
}
