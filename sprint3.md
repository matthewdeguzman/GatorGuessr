# Sprint 3

## Work Completed

### Front-end
- We successfully implemented Google Maps API in our landing page component. We display one map using the default navigation projection seen in Google Maps on the right of the landing page focused on the Gainesville city area. We defined the latitude and longitude bounds to make a rectangle that encapsulates the campus and the surrounding regions from Celebration Pointe in the Southwest to Satchels Pizza in the Northeast. The implementation begins zoomed out but allows users to zoom into any location within defined bounds and restricts users from navigating outside.
- On the left of the landing page is a street view implementation of the Google Maps API showcasing a random street in Gainesville. We randomize this function using (Math.random() * (max-min)+min) to calculate random latitude and longitude coordinates which we would then pass into a validator function to ensure that there is a street view available at that location. If not, we would randomize it continuously until it returns a valid location. These functions provide that users consistently have a new and unique location. This street view map also allows users to browse the surrounding areas.

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

