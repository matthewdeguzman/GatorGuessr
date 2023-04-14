# Sprint 1

## User Stories

1. As a competitive player, I want to be able to save my progress and have a personal account.
2. As an amateur user, I want an easily identifiable and clear, user-friendly homepage to get an idea of what is happening.
3. As a social user, I want to see my friends' scores and compare them to mine.
4. As a UF student, I want to familiarize myself with Gainesville, and I want to be able to know what areas I am unfamiliar with, so I can improve.

## Issues to Adress

We planned to implement the login interface for the user to create an account. This requires the front end team to create the UI for the login and the back end team to create a REST API so the front-end can communicate with the server.

## Successful Issues

We successfully implemented an intuitive, clean login interface through angular material. We routed multiple pages, such as login and app. The back-end team created a mysql database and used gorilla mux and gorm to create a REST API to get users, get a specified user, create a user, update a user, and delete a user. Front-end was able to communicate using the Azure server using HTTPS requests.

## Unsuccessful Issues

We don't have a registration page, so the back end must input the login information independently. We currently have two banners on the home page as well.
