import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { map } from "rxjs/operators";
import { CookieService } from "ngx-cookie-service";

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
  constructor(private http: HttpClient, private CookieService: CookieService) {}

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
    return this.http.post(`${this.uri}/api/login/`, user, {
      withCredentials: true,
      observe: "response",
      responseType: "text",
    });
  }

  // Creates a user
  createUser(username: string, password: string) {
    const user = this.buildUser(username, password);
    return this.http
      .post(`${this.uri}/api/users/`, user, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response));
  }

  // Gets a user
  getUser(username: string) {
    return this.http
      .get(`${this.uri}/api/users/${username}/`, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Gets api key
  getApiKey() {
    return this.http
      .get(`${this.uri}/api/maps-key/`, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response));
  }

  // Deletes a user
  deleteUser(username: string) {
    return this.http
      .delete(`${this.uri}/api/users/${username}/`, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }

  // Gets a user's score
  getUserScore(username: string) {
    console.log(`UserLoginCookie=${this.CookieService.get("UserLoginCookie")}`);
    return this.http
      .get(`${this.uri}/api/users/${username}/`, {
        withCredentials: true,
        observe: "body",
        responseType: "json",
      })
      .pipe(map((response) => response));
  }

  // Updates a users Username and Password
  updateUser(username: string, password: string) {
    console.log(this.CookieService.get("UserLoginCookie"));
    const body = { Username: username, Password: password };
    return this.http
      .put(`${this.uri}/api/users/${username}/`, body, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response));
  }

  // Updates a users score
  updateScore(username: string, score: number) {
    console.log(this.CookieService.get("UserLoginCookie"));
    const body = { Score: score };
    return this.http
      .put(`${this.uri}/api/users/${username}/`, body, {
        withCredentials: true,
        observe: "response",
        responseType: "text",
      })
      .pipe(map((response) => response));
  }
}
