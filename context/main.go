package main

import (
	"context"
	"log"
	"time"
)

func DoStuff(ctx context.CancelFunc) {

	for {
	}

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
}
