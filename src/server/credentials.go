package main

const DB_USERNAME = "cen3031"
const DB_PASSWORD = "bestprojectever_123"
const DB_NAME = "user_database"
const DB_HOST = "cen3031-server.mysql.database.azure.com"
const DB_PORT = "3306"

const DSN = DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
