# Sprint 3

## Video:
''

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
- A new end point at the route `/api/login/` was introduced to provide user validation.
- Added new endpoint for getting users with top score
- restructured server directory
- use environment variables to store database password

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

#### More Visiting Pages Tests:
1. Ensures all buttons on the banner take the user to the correct page no matter what page they are on
2. Ensures that when a user puts in a url that doesn't exist, they get taken to the page-not-found component
3. When the user purposely goes to the page-not-found component, it works and the button to go back takes them to their previous page

### Back-end

## Documentation



