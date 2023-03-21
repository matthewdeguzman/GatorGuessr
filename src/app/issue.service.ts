import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

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
  validateUser(username: string, password: string) {
    const user = {
      username: username,
      password: password,
    };
    this.http
      .post(`${this.uri}/api/login/`, user, { observe: "response" })
      .subscribe((response) => {
        response && response.status
          ? console.log("Login successful")
          : console.log("Login failed");
      });
  }
  createUser(username: string, password: string) {
    const user = {
      username: username,
      password: password,
    };
    return this.http
      .post(`${this.uri}/api/users`, user, { observe: "response" })
      .subscribe((response) => {
        console.log(response.status);
      });
  }
  getUser(username: string) {
    return this.http.get(`${this.uri}/api/login/${username}`);
  }
}
