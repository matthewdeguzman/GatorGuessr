import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { IssueService } from '../../issue.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent{
  loginForm = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  });
  constructor(private IssueService: IssueService) { }
// @Input() error: string | null;
@Output() submitEM = new EventEmitter();
  
public showPassword: boolean = false;

public validUser: boolean = false;

public showUserError(): boolean {
 return this.validUser;
  
}

// verifyUser(): boolean {
//   if (this.submitLogin && this.passwordRequirements.valid) {
//     return false; //means it will display
//   }
//   else return true;
  
// }



public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
  submitLogin(username:string,password:string){
    this.IssueService.getUsersWithUsername(username).subscribe((res) => {
      console.log(res);
      if (res.Username != username){
        
        console.log("Username does not exist");
        this.validUser = true;
      }
      else this.validUser = false;
    });
  }
}

