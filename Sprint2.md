# Sprint 2

## Work Completed

### Front-end

- Created Cypress Tests to test Frontend functionality 

### Back-end

- Reworked the API to take usernames instead of userIDs
- Created unit tests to test the API function calls
- Fixed bug to allow cross-origin requests from different ports to call the API

## Front-end Unit Tests

## Cypress Tests
We made 7 tests in total, each can be split up into their corresponding component
### loginTests
Tests if the proper error message comes up when logging in for 2 different cases:
1. If a user enters a password incorrectly
2. If the enetered user doesn't exist
3. Or nothing comes up If the password and username are correct
### registrationTests
1. Tests that an error message comes up when someone tries to make a user with an already existing username
### visitingPages
Simplest tests, ensures that when a button for that page is clicked it visits the page correctly
1. Home page
2. Login page
3. Register page

## Back-end Unit Tests

### TestGetUsers(t *testing.T)

- Tests the getUsers function by sending a GET request to the server

### TestGetUser1(t* testing.T)

- Tests the getUser function by sending a GET request for a user in the database to the server

### TestGetUser2(t* testing.T)

- Tests the getUser function by sending a GET request with a user not in the database to the server

### TestCreateUser(t *testing.T)

- Tests the createUser function by sending a POST request with a new user to the server

### TestDeleteUser1(t *testing.T)

- Tests the deleteUser function by first creating a user with a GET request then deleting the user with a DELETE request

### TestDeleteUser2(t *testing.T)

- Tests the deleteUser function by sending a DELETE request for a user not in the database

### TestUpdateUser1(t *testing.T)

- Tests the updateUser function by first creating a test user via a GET request then updating the password with a PUT request

### TestUpdateUser2(t *testing.T)

- Tests the updateUser function by sending a PUT request to update a user that does not exist in the database

## API Documentation

### func getUsers(w http.ResponseWriter, r *http.Request)

Retrieves all the users from the database and encodes the data in JSON format into w.

### func getUser(w http.ResponseWriter, r *http.Request)

Retrieves the username parameter from r and encodes the user with the matching username from the database in JSON format into w. If no user is found, then the JSON will have empty values.

### func createUser(w http.ResponseWriter, r *http.Request)

Receives a JSON object from r and creates a new user with the corresponding fields, then encodes the user as a JSON object into w.

### func updateUser(w http.ResponseWriter, r *http.Request)

Retrieves a username from r, searches for the user with the matching username in the database, updates the fields that are different, then encodes the updated user object as a JSON object into w.

### func deleteUser(w http.ResponseWriter, r *http.Request)
takes the id parameter in the request deletes the user with the matching id

## FrontEnd Documentation 

### public showUserError

Returns a boolean on whether the username is already taken. If there is a matching display in an error message on the registration page.

### submitRegistration

Takes in user input of username and password for registration. Checks to see if there is a matching username through HTTP get and, if not, use HTTP post to create a user and navigate to the log-in page. If there is a match, it displays an error message through showUserError implemented in the HTML file. 

### verifySubmit

This checks whether the submit button will be enabled, depending on the user fulfilling username and password requirements.

### showUserError

This boolean displays an error message if the username is not found in the system during login.

### showPassError

This boolean displays an error message if there was a found user, but it was an incorrect password during login.

### submitLogin

Takes in user input of username and password for logging in. This checks whether it is a match for a defined user in the system and, if not, displays error messages utilizing the previous two boolean functions. It scans through an HTTP get. It logs in and navigates the user to the home page if it is a match.


