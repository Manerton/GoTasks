package functions

import (
	"context"
	"math/rand"
	"time"
)

func GenerateRandomNumber(ctx context.Context, random chan int) {
	defer close(random)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		select {
		case <-ctx.Done():
			return
		case random <- r.Int():
		}
	}
}
