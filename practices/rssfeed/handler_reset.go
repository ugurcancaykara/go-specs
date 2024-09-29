package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("Failed to reset the database : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database reset successful.")
	return nil

}
