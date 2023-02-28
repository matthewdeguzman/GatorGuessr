import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { IssueService, User } from '../../issue.service';

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

public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
  submitLogin(username:string,password:string){
    this.IssueService.getUsersWithUsername(username).subscribe((res: User) => {
      if(res.ID==0){
        console.log("User does not exist");
      }
      else if(res.Password==password){
        console.log("Login successful");
      }
      else{
        console.log("Incorrect password");
      }
    });
  }
}

