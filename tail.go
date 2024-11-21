package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	sleepInterval = 1 * time.Second
	filePath      = "./error.log"
	timeout       = 10 * time.Minute
)

func getLastPosition(file *os.File) (int64, error) {
	return file.Seek(0, io.SeekEnd)
}

func watchFile(ctx context.Context, filePath string, interval time.Duration) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	position, err := getLastPosition(file)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(interval):
			nextPosition, err := getLastPosition(file)
			if err != nil {
				return err
			}

			if nextPosition > position {
				if _, err := file.Seek(position, io.SeekStart); err != nil {
					return err
				}

				buffer := make([]byte, nextPosition-position)
				if _, err = file.Read(buffer); err != nil {
					return err
				}

				fmt.Print(string(buffer))
				position = nextPosition
			}
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := watchFile(ctx, filePath, sleepInterval); err != nil {
		log.Fatal(err)
	}
}
