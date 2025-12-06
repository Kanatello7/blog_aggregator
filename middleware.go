package main

import (
	"context"

	"github.com/Kanatello7/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		current_user, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			return err
		}
		return handler(s, cmd, current_user)
	}
}
