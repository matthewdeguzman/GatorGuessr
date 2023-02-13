# Sprint 2
## API Documentation
funtions take http ResponseWriter and Request
### func getUsers(w http.ResponseWriter, r *http.Request)
returns all the users and their information in the database
### func getUser(w http.ResponseWriter, r *http.Request)
takes the id parameter in the request
returns the user whose id matches the input
### func createUser(w http.ResponseWriter, r *http.Request)
takes a json file with a username and password and creates a new user with a new id
### func updateUser(w http.ResponseWriter, r *http.Request)
takes the id parameter in the request and a json file with the new username and password
changes the username and password of the user with the inputed id
### func deleteUser(w http.ResponseWriter, r *http.Request)
takes the id parameter in the request
deletes the user with the matching id
