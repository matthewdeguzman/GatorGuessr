# Sprint 4

## [Sprint 4 Video](https://youtu.be/J21OCYAhNqc0)

## Work Completed

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

### Landing Page and Map Tests

1. Registers click on map
2. Signs in and signs out
3. Tests delete user by creating a new user, logs in, and deletes user
4. Tests the next button to refresh the map


#### More Visiting Pages Tests:

1. Ensures all buttons on the banner take the user to the correct page no matter what page they are on
2. Ensures that when a user puts in a url that doesn't exist, they get taken to the page-not-found component
3. When the user purposely goes to the page-not-found component, it works and the button to go back takes them to their previous page

#### Google Maps

We updated our layout on the landing page component for our Google Maps API implementation. We display one map, street view panorama, as the base layer and then in a small box on the bottom left of the screen is the second map, the default navigation projection. We defined the latitude and longitude bounds to make a rectangle that encapsulates the campus and the surrounding regions from Celebration Pointe in the Southwest to Satchels Pizza in the Northeast. 

#### Google Maps Streetview

On the left of the landing page is a street view implementation of the Google Maps API showcasing a random street in Gainesville. We randomize this function using (Math.random() * (max-min)+min) to calculate random latitude and longitude coordinates which we would then pass into a validator function to ensure that there is a street view available at that location. If not, we would randomize it continuously until it returns a valid location. These functions provide that users consistently have a new and unique location. This street view map also allows users to browse the surrounding areas.

#### Marker

We created marker functions that allow users to place one dynamically anywhere on the navigation map. We limited it to one at a time by clearing previous markers every time a user clicks a new one.

#### Submit

We created a submit button on top of the google maps instantiation. Adjacent to the submit button is a timer set for 60 seconds, the length of one round. Once the timer hits 0, the game automatically submits and calculates the score as 0 or from wherever the user placed their marker. After the user submits, it also creates a red polyline that connects the user's marker to the randomized location. We also implemented a next button that resets the map, score, and timer.

#### Scoring

We implemented a Euclidean distance formula to calculate the score rather than Haversine since it is a small distance localized in Gainesville. We then calculate it using Math.max(maxPoints - (distance / maxDistance) * maxPoints, minPoints) where the maximum is 1000, min is 0, and maxDistance is 0.04. We then use HTTP put requests and update the user, to update the score on the banner and the user's high score.

#### Leaderboard

We've integrated a leaderboard on the home page that displays the highest-ranking users and their scores, which dynamically updates as more games of GatorGuessr are played. To achieve this, the frontend sends an HTTP Get request to retrieve a JSON file containing an array of the top 10 scores, which are then showcased on the application's front page.

#### Dark Mode to Light Mode

By utilizing Material UI's style templates, we devised a method for users to switch the website's theme between light mode and dark mode, providing them with the flexibility to tailor the website's design to their individual preferences. Moreover, this approach involved converting all the CSS files to SCSS, an efficient design practice that streamlines website development and eases our work in the future.

#### Recieving HTTP Status Code

Rather than returning the user itself, logging in and registering users yield HTTP status codes, a security measure that safeguards passwords and data on the API, while also facilitating communication with the backend. To enhance efficiency, we developed a user service that streamlines HTTP requests and enhances their organization, particularly since we need to make numerous backend calls.

#### Page-not-found Component

When an incorrect URL is entered, users are redirected to a "page not found" component that provides them with the option to return to their previous location, thereby enhancing navigation throughout the application. This feature serves as a safeguard against user confusion, in case they unintentionally navigate to the wrong page or we mistakenly redirect them to an nonexistent location.

#### Deleting User

If a user decides to delete their account, they can do so by sending an `HTTP delete request` to the backend through the app. This action will log the user out of the site and delete any cookies and browsing history associated with their account. Deleting an account gives the user complete control over their data and allows them to remove their account from the app if they wish.

#### App Guards

The App Guards feature ensures that users can only access sites that they are authorized to view. This is achieved by checking if the user has a valid cookie that grants them access to the site. If the cookie is missing or invalid, the App Guards prevent the user from accessing the site. This security measure helps to protect sensitive information and restricts unauthorized access to the app.

##### Cookies

The app utilizes cookies to store user-specific data, including their high score and username. Cookies also play a vital role in user authentication, as they are used to validate the user's identity and grant them access to the rest of the site. By relying on cookies, the app can provide a seamless and personalized experience for each user, while also ensuring the security of their data.

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

