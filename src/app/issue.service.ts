import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class IssueService {
  uri = 'http://localhost:9000';

  constructor(private http: HttpClient) { }

  getUsersWithId(id: number) {
    return this.http.get(`${this.uri}/api/users/${id}`)
  }
  getUsers() {
    return this.http.get(`/api/users`)
  }

}
