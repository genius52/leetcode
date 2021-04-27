package other

import (
	"context"
	"fmt"
	"time"
)

func Context_timout_test() {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
