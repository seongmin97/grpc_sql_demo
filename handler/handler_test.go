package handler

import (
	"exercise/gRPC_test/model"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Test model recover")
		}
	}()
	fmt.Println("Test model init...")

	// init sql
	model.InitDatabase()
	// init log
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	code := m.Run()
	fmt.Println("Test model close")
	os.Exit(code)
}
