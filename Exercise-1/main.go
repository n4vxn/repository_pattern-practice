package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
}

type InMemoryUserRepo struct {
	users []User
}

func (repo *InMemoryUserRepo) Create(user *User) error {
	repo.users = append(repo.users, *user)
	return nil
}

func (repo *InMemoryUserRepo) FindByID(id int) (*User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("Not found")
}

func main() {
	repo := &InMemoryUserRepo{}

	repo.Create(&User{
		ID:       1,
		Username: "user1",
		Email:    "user1@mail.com",
	})

	user, err := repo.FindByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found user: %v\n", user)
	}
}

/*

What we built:

    A simple in-memory repository that stores users in a slice.
    Implemented the UserRepository interface with methods like Create (to add users) and FindByID (to retrieve users by their ID).
    We kept all the data in memory, so it is lost when the program stops.
    This exercise focused on understanding how to organize the data and interact with it using an interface.

*/