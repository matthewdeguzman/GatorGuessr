import { Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
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
    return this.http
      .get(`${this.uri}/cookies/set/${username}/`, {
        observe: "response",
        responseType: "text",
      })
      .subscribe((res) => {
        console.log(res.body);
      });
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
