package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
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

type PostgresUserRepo struct {
	db *sql.DB
}

func (repo *PostgresUserRepo) Create(user *User) error {
	query := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id`
	return repo.db.QueryRow(query, user.Username, user.Email).Scan(&user.ID)
}

func (repo *PostgresUserRepo) FindByID(id int) (*User, error) {
	query := `SELECT id, username, email FROM users WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return &user, err
}

func (repo *PostgresUserRepo) Update(user *User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`
	result, err := repo.db.Exec(query, user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", user.ID)
	}
	return nil
}

func (repo *PostgresUserRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}
	return nil
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, email string) (*User, error) {
	if username == "" || email == "" {
		return nil, fmt.Errorf("username and email are required")
	}
	user := &User{Username: username, Email: email}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	addr := os.Getenv("DB_ADDR")
	if addr == "" {
		log.Fatalf("DB_ADDR not set in environment")
	}

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to database successfully")
	return db, nil
}

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}
	defer db.Close()

	repo := &PostgresUserRepo{db: db}
	userService := NewUserService(repo)

	newUser, err := userService.CreateUser("jane_doe", "jane@example.com")
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("Created user: %+v\n", newUser)

	user, err := userService.GetUserByID(newUser.ID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	fmt.Printf("Fetched user: %+v\n", user)

	user.Username = "jane_updated"
	err = userService.UpdateUser(user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Printf("Updated user: %+v\n", user)

	err = userService.DeleteUser(user.ID)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	fmt.Printf("Deleted user with ID: %d\n", user.ID)
}
