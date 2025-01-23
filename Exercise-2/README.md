Exercise 2: Separate Service Layer
Objective:

Learn how to separate the repository logic from business logic by introducing a service layer.
Instructions:

    Refactor the code from Exercise 1 to add a UserService struct.
    The UserService will take a UserRepository as a dependency and provide methods like GetUserByID.
    The service should delegate the actual data fetching to the repository but also handle any additional business logic if needed.