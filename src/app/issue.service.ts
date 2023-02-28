import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class IssueService {
  uri = 'http://localhost:9000';

  constructor(private http: HttpClient) { }

  getUsersWithUsername(username: string) {
    return this.http.get(`${this.uri}/api/users/${username}`)
  }
  createUser(username: string, password: string) {
    const user = {
      username: username,
      password: password,
    };
    return this.http.post(`${this.uri}/api/users`, user);
  }


}
