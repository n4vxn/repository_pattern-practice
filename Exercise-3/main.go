package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
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

func (repo *InMemoryUserRepo) Update(user *User) error {
	for i, u := range repo.users {
		if u.ID == user.ID {
			repo.users[i] = *user
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (repo *InMemoryUserRepo) Delete(id int) error {
	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(user *User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func main() {
	repo := &InMemoryUserRepo{}

	service := UserService{repo: repo}

	service.repo.Create(&User{ID: 1, Username: "john_doe", Email: "john@example.com"})
	service.repo.Create(&User{ID: 2, Username: "jane_doe", Email: "jane@example.com"})

	user, err := service.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Updated user: %v\n", user)
	}

	service.DeleteUser(2)

	_, err = service.GetUserByID(2)
	if err != nil {
		fmt.Println(err)
	}
}

/*

What we built:

    We extended the UserRepository interface to support more operations: Update (to modify user details) and Delete (to remove a user).
    We implemented these methods in the InMemoryUserRepo, allowing us to create, find, update, and delete users.
    This exercise helped us practice working with more advanced CRUD operations in a repository pattern, while still using in-memory storage.

*/