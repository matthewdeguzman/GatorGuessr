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

## Front End
### public showUserError - registration
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
