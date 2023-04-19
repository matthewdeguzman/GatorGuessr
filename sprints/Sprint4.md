# Sprint 4

## Work Completed

## Back-end
- Implemented cookies and proper authorization flow. Previously, there was no way of authorizing the requests which was a big security concern. However, we have worked with the net/http package to create cookies upon certain requests. Certain endpoints now require cookies to be sent. Furthermore, the cookies are verified via an HMAC signature that is prepended to the value upon creation. 

## Tests

- TestGetUsers: Tests if GetUsers endpoint correctly returns the list of users
- TestGetWithoutAuthorization: Tests if an error is returned when calling GetUsers endpoint without proper cookie
- TestGetExistingUser: Tests if GetUser endpoint returns an existing user
- TestGetNonexistantUser: Tests if GetUser responds with a 404 error if a call for a user not in the database is made
- TestCreateNewUser: Tests if CreateUser endpoint correctly creates a new user
- TestCreateUserWithoutPassword: Tests if CreateUser endpoint correctly responds with an error if a user is attempting to be created without a password
- TestCreateUserWithID: Tests if CreateUser endpoint correctly throws an error if a user is attempting to be created with an ID
- TestUpdateNonexistantUser: Tests if UpdateUser endpoint correctly responds with an error if a user that does not exist in the database is being udpated
- TestUpdateUserWithoutAuthorization: Tests if CreateUser endpoint returns an error when called without proper cookie
- TestUpdateExistingUser: Tets if UpdateUser endpoint correctly updates an existing user in the database
- TestUpdateUserID: Tests if UpdateUser endpoint correctly responds with an error if the ID of a user is being updated
- TestUpdateUserScore: Tests if UpdateUser endpoint correctly updates the score field of a user
- TestUpdateUserWithoutAuthorization: Tests if DeleteUser endpoint returns an error when called without proper cookie
- Test DeleteExisingUser: Tests if DeleteUser endpoint correctly deletes an existing user from the database
- TestDeleteNonexistingUser: Tests if DeleteUser endpoint correctly responds with an error if a call to delete a user that does not exist in the database is made
- TestValidateExistingUserCookie: Tests if a cookie is created properly upon a successful call of the ValidateUser endpoint
- TestValidateExisingUser: Tests if ValidateUser endpoint correctly validates an existing user with the correct password in the database
- TestValidateNonexistantUser: Tests if ValidateUser endpoint correctly responds with an error if a nonexistant user is being validated
- TestValidateIncorrectPassword: Tests if ValidateUser endpoint correctly responds with an error if a username is provided with an incorrect password
- TestLeaderboardNegativeInteger: Tests if the GetTopUsers endpoint correctly handles negative integers as a parameter
- TestLeaderboardFloatintPoint: Tests if the GetTopUsers endpoint correctly handles floating point numbers as a parameter
- TestLeaderboardZero: Tests if the GetTopUsers endpoint correctly handles 0 as a parameter
- TestLeaderboardOverMaxLimit: Tests if the GetTopUsers endpoint correctly handles an integer larger than the number of users in the database
- TestLeaderboardSorted: Tests if the GetTopUsers endpoint correctly returns a sorted list of users according to their scores

## Documentation

## Authorization Flow

### Creating signature

The endpoints `CreateUser` and `ValidateUser` both create cookies for their respective users upon successful calls. The value of the cookie is dependent on the ID of the user and is base 64 encoded then made cryptographically secure by creating an HMAC signature and prepending it to the cookie value. The cookie is set with the `Set-Header` header in the response.

### Verification

The endpoints `GetUser`, `UpdateUser`, and `DeleteUser` all required a cookie to be sent in the http request. If no cookie is found, then a `400` response is returned. The value of the cookie is seperated into an HMAC signature and its value based on a size specified by the backend. The HMAC signature is recomputed based on the value and name of the cookie and compared with the HMAC signature sent with the cookie. If the signatures both match, then the request continue, else a `404` response is returned.

## Endpoints

### `GET /api/users/` func GetUsers(w http.ResponseWriter, r *http.Request)

Retrieves users from the MySql database using `db.first` where `db` is a `*gorm.DB` pointer and stores the received data in a slice of users. The slice is encoded in the ResponseWriter.

### `GET /api/users/{username}` func GetUser(w http.ResponseWriter, r *http.Request)

Reads the `username` parameter from the request url and uses retrieves the user in the database with the corresponding username. If no user is found, a 404 error is written in the http response, else the retrieved user is encoded in the http response.

### `POST /api/users/` func CreateUser(w http.ResponseWriter, r *http.Request)

Reads the JSON object from the http request and and encodes it into a user struct. If the user being created has the same username as another user in the database, a 400 error is written in the response. If no password was provided or an ID was provided, a 400 error is written in the http response. Otherwise, the user is inserted in the database.

### `PUT /api/users/{username}/` func UpdateUser(w http.ResponseWriter, r *http.Request)

Reads `username` from the http request and searches for the user in the database with the corresponding username. If no user is found, then a 404 error is written in the response. If the ID is being changed, a 405 error is written in the response. Otherwise, the user in the database has the provided fields updated. Note that if the password is provided, then it is rehashed with argon2.

### `DELETE /api/users/{username}/` func DeleteUser(w http.ResponseWriter, r *http.Request)

Reads `username` from the http request and searches for the user in the database with the corresponding username. If no user is found, then a 404 error is written in the response. Otherwise, the user is removed from the database.

### `POST /api/login/` func ValidateUser(w http.ResponseWriter, r *http.Request)

Encodes the JSON object from the http request into a user struct. The password is decrypted using argon 2. Then, the user in the database with the corresponding username is retrieved. If no user is found, then a 404 error is encoded. Otherwise, the received user has the password decryped with argon2 and is compared to the password provided in the http request. If the passwords do not match, a 404 error is encoded. Otherwise, a 200 OK response is encoded

### `GET /api/leaderboard/{limit}/` func GetTopUsers(w http.ResponseWriter, r *http.Request)

Reads `limit` from URl and returns at most `limit` users from the database such that they are sorted in descending order of score. The users are stored in a slice and encoded into the http response body. If `limit` is not a positive integer, a 400 error is encoded in `w`. 

