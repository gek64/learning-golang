package main

import (
	"context"
	"fmt"
	"time"
)

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Printf("\n我要闭嘴了\n")
			return
		default:
			fmt.Printf("bala")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
