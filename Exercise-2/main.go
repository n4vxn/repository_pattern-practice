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
	return nil, fmt.Errorf("user not found")
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func main() {
	repo := &InMemoryUserRepo{}
	service := &UserService{repo: repo}

	repo.Create(&User{
		ID: 1, Username: "user1",
		Email: "user1@email.com",
	})

	user, err := service.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found user: %v\n", user)
	}
}

/*

What we built:

    We refactored the code from Exercise 1 to introduce a UserService layer that acts as an intermediary between the business logic and the repository.
    The UserService uses the repository to fetch or store users, but it could add business logic in the future if needed.
    This exercise demonstrated how to separate concerns and ensure that the repository is only responsible for data access, and the service layer handles business operations.

	*/