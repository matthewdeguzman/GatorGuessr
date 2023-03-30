# Sprint 3

## Video:
''

## Work Completed

### Front-end

#### Dark Mode to Light Mode
By utilizing Material UI's style templates, we devised a method for users to switch the website's theme between light mode and dark mode, providing them with the flexibility to tailor the website's design to their individual preferences. Moreover, this approach involved converting all the CSS files to SCSS, an efficient design practice that streamlines website development and eases our work in the future.
#### Recieving HTTP Status Code
Rather than returning the user itself, logging in and registering users yield HTTP status codes, a security measure that safeguards passwords and data on the API, while also facilitating communication with the backend. To enhance efficiency, we developed a user service that streamlines HTTP requests and enhances their organization, particularly since we need to make numerous backend calls.

### Back-end

- We changed the user struct to include a score field which allows us to order users based on scores. The changes were migrated using the `db.AutoMigrate()` and automatically integrated into the MySql database.
- A new end point at the route `/api/login/` was introduced to provide user validation.
- Added new endpoint for getting users with top score
- restructured server directory
- use environment variables to store database password

## Tests

### Front-end

### Back-end

## Documentation

# Sprint 3


