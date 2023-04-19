import { Injectable } from "@angular/core";
import { CanActivate } from "@angular/router";
import { Router } from "@angular/router";

@Injectable({
  providedIn: "root",
})
export class LoginGuard implements CanActivate {
  constructor(private Router: Router) {}
  canActivate(): boolean {
    if (localStorage.getItem("username") != null) {
      console.log("login false");
      this.Router.navigate(["/landing-page"]);
      return false;
    } else {
      console.log("login true");
      return true;
    }
  }
}
