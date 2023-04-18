import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { map } from "rxjs/operators";

export type User = {
  Username: string;
  Password: string;
  Score: number;
};

@Injectable({
  providedIn: "root",
})
export class UserService {
  private uri = "http://localhost:9000";
  constructor(private http: HttpClient) {}

  // Builds a user
  buildUser(username: string, password: string): User {
    const user: User = {
      Username: username,
      Password: password,
      Score: 0,
    };
    return user;
  }
  // Checks if login is valid
  validateUser(username: string, password: string) {
    const user = this.buildUser(username, password);
    return this.http
      .post(`${this.uri}/api/login/`, user, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Creates a user
  createUser(username: string, password: string) {
    const user = this.buildUser(username, password);
    return this.http
      .post(`${this.uri}/api/users/`, user, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Gets a user
  getUser(username: string) {
    return this.http
      .get(`${this.uri}/api/users/${username}/`, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Gets api key
  getApiKey() {
    return this.http
      .get(`${this.uri}/apikey/get/`, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Deletes a user
  deleteUser(username: string) {
    return this.http
      .delete(`${this.uri}/api/users/${username}/`, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Gets a user's score
  getScore(username: string) {
    return this.http
      .get(`${this.uri}/api/users/${username}/`, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }
}

/*
import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { map } from "rxjs/operators";

export type User = {
  ID: number;
  Username: string;
  Password: string;
};

@Injectable({
  providedIn: "root",
})
export class IssueService {
  uri = "http://localhost:9000";
  constructor(private http: HttpClient) {}

  // Checks if login is valid
  validateUser(username: string, password: string) {
    const user = {
      username: username,
      password: password,
    };
    return this.http
      .post(`${this.uri}/api/login/`, user, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Makes a user
  createUser(username: string, password: string) {
    const user = {
      username: username,
      password: password,
    };
    return this.http
      .post(`${this.uri}/api/users/`, user, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  getUser(username: string) {
    return this.http
      .get(`${this.uri}/api/users/${username}/`, {
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }
}
*/
