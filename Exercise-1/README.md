Exercise 1: Simple Repository for In-memory Storage
Objective:

Understand the basic structure of the Repository Pattern by implementing it with in-memory storage (no database).
Instructions:

    Create a User struct with fields like ID, Username, and Email.
    Define a UserRepository interface with basic methods like Create and FindByID.
    Implement the UserRepository interface with a struct that stores users in memory (a simple slice).
    Write a function to add a user and retrieve a user by ID.