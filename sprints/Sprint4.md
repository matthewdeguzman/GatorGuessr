# Sprint 4

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

