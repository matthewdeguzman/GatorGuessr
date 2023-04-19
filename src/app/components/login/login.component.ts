import { IssueService } from "../../services/issue.service";
import { Component, OnInit, Input, Output, EventEmitter } from "@angular/core";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import { CookieService } from "ngx-cookie-service";

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
  styleUrls: ["./login.component.scss"],
})
export class LoginComponent {
  loginForm = new FormGroup({
    username: new FormControl(""),
    password: new FormControl(""),
  });

  constructor(
    private IssueService: IssueService,
    private router: Router,
    private CookieService: CookieService
  ) {}

  cookie: string;

  // @Input() error: string | null;
  @Output() submitEM = new EventEmitter();

  public showPassword: boolean = false;

  public invalidName: boolean = false;

  public invalidUser: boolean = false;

  public showUserError(): boolean {
    return this.invalidName;
  }

  public showPassError(): boolean {
    return this.invalidUser;
  }

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }

  async timer() {
    await new Promise((resolve) => setTimeout(resolve, 500));
  }

  async submitLogin(username: string, password: string) {
    this.IssueService.validateUser(username, password).subscribe(
      (res) => {
        if (res.status == 200) {
          const cookie = res.body as string;
          localStorage.setItem("username", username);
          localStorage.setItem("token", cookie);
          this.router.navigate(["/landing-page"]);
          console.log(this.cookie);
        }
      },
      (error) => {
        if (error.status == 500) {
          this.invalidName = true;
          this.invalidUser = false;
        } else if (error.status == 404) {
          this.invalidUser = true;
          this.invalidName = false;
        }
      }
    );
    console.log(localStorage.getItem("token"));
  }
}
