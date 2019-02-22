package graphql_poc

import (
	"context"
	"errors"
)

type Resolver struct {
	Messages map[int]*Message
	Users    map[int]*User
}

func Data() (map[int]*User, map[int]*Message) {
	users := make(map[int]*User)
	users[1] = &User{
		ID:       "1",
		Username: "Robin Wieruch",
	}
	users[2] = &User{
		ID:       "2",
		Username: "Dave Davids",
	}

	messages := make(map[int]*Message)
	messages[1] = &Message{
		ID:   "1",
		Text: "Hello World",
		User: "1",
	}
	messages[2] = &Message{
		ID:   "2",
		Text: "Bye World",
		User: "2",
	}

	return users, messages
}

func (r *Resolver) Message() MessageResolver {
	users, messages := Data()

	r.Users = users
	r.Messages = messages
	return &messageResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	users, messages := Data()

	r.Users = users
	r.Messages = messages
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	users, messages := Data()

	r.Users = users
	r.Messages = messages
	return &userResolver{r}
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) User(ctx context.Context, obj *Message) (User, error) {
	for _, user := range r.Resolver.Users {
		if user.ID == obj.User {
			return *user, nil
		}
	}
	return User{}, nil
}

func (r *messageResolver) Message(ctx context.Context, obj *Message, id string) (Message, error) {
	for _, message := range r.Messages {
		if message.ID == id {
			return *message, nil
		}
	}

	return Message{}, errors.New("Message not found")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	users := make([]User, 0)
	for _, user := range r.Resolver.Users {
		users = append(users, *user)
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	users, err := r.Users(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	user := r.Resolver.Users[1]

	messages := make([]*Message, 0)

	for _, message := range r.Resolver.Messages {
		if message.User == user.ID {
			messages = append(messages, message)
		}
	}
	user.Messages = messages
	return user, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]Message, error) {
	messages := make([]Message, 0)

	for _, message := range r.Resolver.Messages {
		messages = append(messages, *message)
	}

	return messages, nil
}

func (r *queryResolver) Message(ctx context.Context, id string) (Message, error) {
	messages, err := r.Messages(ctx)
	if err != nil {
		return Message{}, err
	}

	for _, message := range messages {
		if message.ID == id {
			return message, nil
		}
	}

	return Message{}, errors.New("Message not found")
}

type userResolver struct{ *Resolver }

func (r *userResolver) User(ctx context.Context, obj *User, id string) (*User, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("User not found")
}

func (r *userResolver) Me(ctx context.Context, obj *User) (*User, error) {
	user := r.Resolver.Users[1]
	return user, nil
}
