# Sprint 4

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

### func SetCookieHandler(w http.ResponseWriter, r *http.Request)

This function handles the path `/cookies/get/`. The function receives a JSON in the http request that formats the [specific fields](https://go.dev/src/net/http/cookie.go) of a cookie. The HMAC signature is encoded with a secret key and based on the name and value of the sent cookie then preprended to the cookie value. The value is then encoded in base 64 and if the string is larger than 4096 characters, an error is thrown. Otherwise, the cookie is written into the response writer. Note that the `COOKIE_SECRET` environment variable must be defined and it is usually a 32-bit cryptographically randomly generated number.

### func GetCookieHandler(w http.ResponseWriter, r *http.Request)

This function handles the path `/cookies/set/{cookie-name}/`. The function receives a cookie in the http request and the name of the cookie as the `cookie-name` routing variable then verifies whether the cookie is valid. If there is no cookie with the given name, a bad request is returned. The name and value of the cookie are base 64 encoded, so they are decoded first. Then, the cookie is verified if the HMAC signature generated with the secret key, cookie name, and cookie value match the expected HMAC signature of the same values. If the signatures do not match, the cookie is not valid and a bad request is returned. Note that the `COOKIE_SECRET` environment variable must be defined and it is usually a 32-bit cryptographically randomly generated number.
