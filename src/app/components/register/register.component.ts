import { Router } from "@angular/router";
import { Component, OnInit, Input, Output, EventEmitter } from "@angular/core";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { IssueService, User } from "../../issue.service";

@Component({
  selector: "app-register",
  templateUrl: "./register.component.html",
  styleUrls: ["./register.component.css"],
})
export class RegisterComponent {
  registerForm = new FormGroup({
    username: new FormControl(""),
    password: new FormControl(""),
  });
  usernameRequirements = new FormControl("", [
    Validators.required,
    Validators.minLength(4),
    Validators.maxLength(20),
  ]);
  passwordRequirements = new FormControl("", [
    Validators.required,
    Validators.minLength(8),
    Validators.maxLength(25),
    Validators.pattern(
      "^(?=[^A-Z]*[A-Z])(?=[^a-z]*[a-z])(?=\\D*\\d)[A-Za-z\\d!$%@#£€*?&]{8,}$"
    ),
  ]);
  constructor(private IssueService: IssueService, private router: Router) {}
  // @Input() error: string | null;
  @Output() submitEM = new EventEmitter();

  public showPassword: boolean = false;
  public showSubmit: boolean = false;

  public invalidUser: boolean = false;

  public showUserError(): boolean {
    return this.invalidUser;
  }

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }

  submitRegistration(username: string, password: string) {
    this.IssueService.getUser(username).subscribe(
      (res) => {
        if (res == 200) {
          this.invalidUser = true;
        } else {
          console.log("Something weird happened");
        }
      },
      (error) => {
        if (error.status == 404) {
          this.IssueService.createUser(username, password).subscribe((res) => {
            this.invalidUser = false;
            this.router.navigate(["/login"]);
          });
        }
      }
    );
  }
  verifySubmit(): boolean {
    if (this.usernameRequirements.valid && this.passwordRequirements.valid) {
      return false; //means it will display
    } else return true;
  }
}
