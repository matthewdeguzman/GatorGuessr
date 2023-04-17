import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpHeaders,
  HTTP_INTERCEPTORS,
  HttpParams,
} from "@angular/common/http";
import { map } from "rxjs/operators";
import { CookieService } from "ngx-cookie-service";

@Injectable({
  providedIn: "root",
})
export class CookiesService {
  private uri = "http://localhost:9000";
  constructor(private http: HttpClient, private cookieService: CookieService) {}

  // Gets cookie
  getCookie(username: string) {
    document.cookie = `key=Cookie; path=http://localhost:9000/cookies/set/${username}/`;
    const headers = new HttpHeaders().set(document.cookie, username);
    return this.http
      .get(`${this.uri}/cookies/set/${username}/`, {
        headers: headers,
        observe: "response",
        withCredentials: true,
        responseType: "text",
      })
      .pipe(map((response) => response.status));
  }
  // Checks if cookie is valid
  validateCookie() {
    return this.http
      .get(`${this.uri}/cookies/get/`, {
        observe: "response",
        responseType: "text",
      })
      .subscribe((res) => {
        console.log(res.body);
      });
  }
}
