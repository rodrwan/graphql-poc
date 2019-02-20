package graphql_poc

import (
	"context"
	"errors"
)

type Resolver struct {
	Messages map[int]Message
	Users    map[int]User
}

func (r *Resolver) Query() QueryResolver {
	messages := make(map[int]Message)
	messages[1] = Message{
		ID:   "1",
		Text: "Hello World",
	}
	messages[2] = Message{
		ID:   "2",
		Text: "Bye World",
	}
	r.Messages = messages

	users := make(map[int]User)
	users[1] = User{
		ID:       "1",
		Username: "Robin Wieruch",
	}
	users[2] = User{
		ID:       "2",
		Username: "Dave Davids",
	}
	r.Users = users

	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	users := make([]User, 0)
	for _, user := range r.Resolver.Users {
		users = append(users, user)
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (User, error) {
	for _, user := range r.Resolver.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

func (r *queryResolver) Me(ctx context.Context) (User, error) {
	user := r.Resolver.Users[1]
	return user, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]Message, error) {
	messages := make([]Message, 0)
	for _, message := range r.Resolver.Messages {
		messages = append(messages, message)
	}

	return messages, nil
}

func (r *queryResolver) Message(ctx context.Context, id string) (Message, error) {
	for _, message := range r.Resolver.Messages {
		if message.ID == id {
			return message, nil
		}
	}

	return Message{}, errors.New("Message not found")
}
