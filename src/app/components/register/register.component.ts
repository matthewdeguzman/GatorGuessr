import { Router } from '@angular/router';
import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { IssueService, User } from '../../issue.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent{
  registerForm = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  });
  usernameRequirements = new FormControl('',[Validators.required, Validators.minLength(4), Validators.maxLength(20)]);
  passwordRequirements = new FormControl('',[
    Validators.required, 
    Validators.minLength(8),
    Validators.maxLength(25),
    Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]$/)
  ]);
  constructor(private IssueService: IssueService, private router: Router) { }
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

  submitRegistration(username:string,password:string)
  {
    this.IssueService.getUsersWithUsername(username).subscribe((res: User) => {
      if(res.ID!=0){
        console.log("User already exists");
        this.invalidUser = true;

      }
      else{
        this.IssueService.createUser(username,password).subscribe((res) => {
          console.log(res);
          this.invalidUser = false;
          this.router.navigate(['/login']);
        });
      }
    });
  }
  verifySubmit(): boolean {
    if (this.usernameRequirements.valid && this.passwordRequirements.valid) {
      return false; //means it will display
    }
    else return true;
    
  }
}
