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
    let params = new HttpParams()
      .set("Name", JSON.stringify(username))
      .set("MaxAge", 31536000)
      .set("Value", JSON.stringify("miku"));
    return this.http
      .get(`${this.uri}/cookies/set/`, {
        params: params,
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
      .get(`${this.uri}/cookies/verify/`, {
        observe: "response",
        responseType: "text",
      })
      .subscribe((res) => {
        console.log(res.body);
      });
  }
}
