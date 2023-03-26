import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { map, catchError } from "rxjs/operators";
import { throwError } from "rxjs";
import { error } from "cypress/types/jquery";

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
