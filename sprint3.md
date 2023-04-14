# Sprint 3

## [Sprint 3 Video](https://youtu.be/XHNM0PuyEt8)

## Work Completed

### Front-end

#### Google Maps

We successfully implemented Google Maps API in our landing page component. We display one map using the default navigation projection seen in Google Maps on the right of the landing page focused on the Gainesville city area. We defined the latitude and longitude bounds to make a rectangle that encapsulates the campus and the surrounding regions from Celebration Pointe in the Southwest to Satchels Pizza in the Northeast. The implementation begins zoomed out but allows users to zoom into any location within defined bounds and restricts users from navigating outside.

#### Google Maps Streetview

On the left of the landing page is a street view implementation of the Google Maps API showcasing a random street in Gainesville. We randomize this function using (Math.random() * (max-min)+min) to calculate random latitude and longitude coordinates which we would then pass into a validator function to ensure that there is a street view available at that location. If not, we would randomize it continuously until it returns a valid location. These functions provide that users consistently have a new and unique location. This street view map also allows users to browse the surrounding areas.

#### Leaderboard

We've integrated a leaderboard on the home page that displays the highest-ranking users and their scores, which dynamically updates as more games of GatorGuessr are played. To achieve this, the frontend sends an HTTP Get request to retrieve a JSON file containing an array of the top 10 scores, which are then showcased on the application's front page.

#### Dark Mode to Light Mode

By utilizing Material UI's style templates, we devised a method for users to switch the website's theme between light mode and dark mode, providing them with the flexibility to tailor the website's design to their individual preferences. Moreover, this approach involved converting all the CSS files to SCSS, an efficient design practice that streamlines website development and eases our work in the future.

#### Recieving HTTP Status Code

Rather than returning the user itself, logging in and registering users yield HTTP status codes, a security measure that safeguards passwords and data on the API, while also facilitating communication with the backend. To enhance efficiency, we developed a user service that streamlines HTTP requests and enhances their organization, particularly since we need to make numerous backend calls.

#### Page-not-found Component

When an incorrect URL is entered, users are redirected to a "page not found" component that provides them with the option to return to their previous location, thereby enhancing navigation throughout the application. This feature serves as a safeguard against user confusion, in case they unintentionally navigate to the wrong page or we mistakenly redirect them to an nonexistent location.

### Back-end

- We changed the user struct to include a score field which allows us to order users based on scores. The changes were migrated using the `db.AutoMigrate()` and automatically integrated into the MySql database.
- A new end point at the route `/api/login/` was introduced to provide user validation with username and password. This allows the front-end team to easily identify if the provided username and password exist in the database.
- Added new endpoint for sorting users by the `Score` attribute at the `api/users/{limit}/` where limit denotes how many users will be returned. This allows the front-end to display a leaderboard of users who have the highest score field.
- We restructured the server directory to allow for more organization and readadbility of our repository and code.
- We use environment variables to store database password rather than reading from `credentials.txt`. This is a more secure way and standard way of reading secure data and can easily be implemented if deployed on a server.

## Tests

### Front-end

#### More Registration Tests

Tests that ensure the correct error message comes up depending on the mistake the user makes when registering

1. Username is too short
2. Username is too long
3. No username was given
4. Password is too short
5. Password is too long
6. Password doesn't contain uppercase
7. Password doesn't contain lowercase
8. Password doesn't contain number
9. No password was given
10. Registering correctly takes user to login page

#### More Visiting Pages Tests

1. Ensures all buttons on the banner take the user to the correct page no matter what page they are on
2. Ensures that when a user puts in a url that doesn't exist, they get taken to the page-not-found component
3. When the user purposely goes to the page-not-found component, it works and the button to go back takes them to their previous page

### Back-end

- TestGetUsers: Tests if GetUsers endpoint correctly returns the list of users
- TestGetExistingUser: Tests if GetUser endpoint returns an existing user
- TestGetNonexistantUser: Tests if GetUser responds with a 404 error if a call for a user not in the database is made
- TestCreateNewUser: Tests if CreateUser endpoint correctly creates a new user
- TestCreateUserWithoutPassword: Tests if CreateUser endpoint correctly responds with an error if a user is attempting to be created without a password
- TestCreateUserWithID: Tests if CreateUser endpoint correctly throws an error if a user is attempting to be created with an ID
- TestUpdateNonexistantUser: Tests if UpdateUser endpoint correctly responds with an error if a user that does not exist in the database is being udpated
- TestUpdateExistingUser: Tets if UpdateUser endpoint correctly updates an existing user in the database
- TestUpdateUserID: Tests if UpdateUser endpoint correctly responds with an error if the ID of a user is being updated
- TestUpdateUserScore: Tests if UpdateUser endpoint correctly updates the score field of a user
- Test DeleteExisingUser: Tests if DeleteUser endpoint correctly deletes an existing user from the database
- TestDeleteNonexistingUser: Tests if DeleteUser endpoint correctly responds with an error if a call to delete a user that does not exist in the database is made
- TestValidateExisintUser: Tests if ValidateUser endpoint correctly validates an existing user with the correct password in the database
- TestValidateNonexistantUser: Tests if ValidateUser endpoint correctly responds with an error if a nonexistant user is being validated
- TestValidateIncorrectPassword: Tests if ValidateUser endpoint correctly responds with an error if a username is provided with an incorrect password
- TestLeaderboardNegativeInteger: Tests if the GetTopUsers endpoint correctly handles negative integers as a parameter
- TestLeaderboardFloatintPoint: Tests if the GetTopUsers endpoint correctly handles floating point numbers as a parameter
- TestLeaderboardZero: Tests if the GetTopUsers endpoint correctly handles 0 as a parameter
- TestLeaderboardOverMaxLimit: Tests if the GetTopUsers endpoint correctly handles an integer larger than the number of users in the database
- TestLeaderboardSorted: Tests if the GetTopUsers endpoint correctly returns a sorted list of users according to their scores

## Documentation

### func GetUsers(w http.ResponseWriter, r *http.Request)

Retrieves users from the MySql database using `db.first` and stores them in a slice of users. The slice is encoded in the ResponseWriter and a 200 OK request is encoded as well.

### func GetUser(w http.ResponseWriter, r *http.Request)

Reads the `username` parameter from `http.Request` and uses `gorm.first` to retrieve the user in the database that has the `username`. If no user is found, a 404 error is encoded in the `w`, else the retrieved user is encoded in `w` and a 200 OK request is encoded.

### func CreateUser(w http.ResponseWriter, r *http.Request)

Reads the JSON object from `r` and creates a new user struct with the JSON object. If the user that is being created has the same username as another user in the database, a 400 error is encoded. If the user is being created with an empty password or ID, a 400 error is encoded as both are not allowed. If none of the conditions are satisfied, the user is encoded in `w` and entered in the database using `db.Create`.

### func UpdateUser(w http.ResponseWriter, r *http.Request)

Reads `username` from `r` and uses `db.first` to execute a search query for the first user in the database with `username`. If no user is found, then a 404 error is thrown. If the ID is being changed, a 405 error is encoded in `w` as it is an immutable field and should not be changed. Else, the user has their password encoded with argon2 and is updated in the database using `db.Save`.

### func DeleteUser(w http.ResponseWriter, r *http.Request)

Reads `username` from `r` and uses `db.first` to execute a search query for the first user in the database with `username`. If no user is found, then a 404 error is thrown. Else, the user is removed from the database using `db.Delete`.

### func ValidateUser(w http.ResponseWriter, r *http.Request)

Reads `user` from the JSON object provided in `r`.  `user.password` is decrypted using argon 2. Then, a search query is executed to retrieve the corresponding user in the database with the same username. If no user is found, then a 404 error is encoded. Else, the received user has the password decryped using argon2 and compares if the two passwords match. If the passwords do not match, a 404 error is encoded. Else, the passwords match and a 200 OK is encoded in `w`.

### func GetTopUsers(w http.ResponseWriter, r *http.Request)

Reads `limit` from `r` and executes a search query with `db.Limit` to get `limit` amount of users. The users are sorted in descending order of score using `db.Order` and stored in a slice. If `limit` is not a uint, a 400 error is encoded in `w`. Else, the slice of users is encoded in `w`.
