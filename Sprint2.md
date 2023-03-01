# Sprint 2

## Work Completed

### Front-end

- TBD

### Back-end

- Reworked the API to take usernames instead of userIDs
- Created unit tests to test the API function calls
- Fixed bug to allow cross-origin requests from different ports to call the API

## Front-end Unit Tests

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

Retrieves the username from r, then searches for the user with the corresponding username and deletes the user from the database. The fields of the user are reset to the default values and encoded as a JSON object into w.
