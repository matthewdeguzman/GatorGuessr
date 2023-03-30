# Sprint 3

## [Sprint 3 Video](https://youtu.be/XHNM0PuyEt8)

## Work Completed

### Front-end

### Back-end

- We changed the user struct to include a score field which allows us to order users based on scores. The changes were migrated using the `db.AutoMigrate()` and automatically integrated into the MySql database.
- A new end point at the route `/api/login/` was introduced to provide user validation with username and password. This allows the front-end team to easily identify if the provided username and password exist in the database.
- Added new endpoint for sorting users by the `Score` attribute at the `api/users/{limit}/` where limit denotes how many users will be returned. This allows the front-end to display a leaderboard of users who have the highest score field.
- We restructured the server directory to allow for more organization and readadbility of our repository and code.
- We use environment variables to store database password rather than reading from `credentials.txt`. This is a more secure way and standard way of reading secure data and can easily be implemented if deployed on a server.

## Tests

### Front-end

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