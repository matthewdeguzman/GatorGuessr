import { IssueService } from './../../issue.service';
import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';



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

  constructor(private IssueService: IssueService, private router: Router) { }
  

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

submitLogin(username:string,password:string){
  this.IssueService.getUsersWithUsername(username).subscribe((res) => {
    console.log(res);
    if (res.Username != username){
      console.log("Username does not exist");
      this.invalidName = true;
    }
    else{
      this.invalidName = false;
      if(res.Password != password){
        console.log("Incorrect password");
        this.invalidUser = true;
      }
      else{
        console.log("Login successful");
        this.invalidUser = false;
        this.router.navigate(['/home']);
        
    }
    } 
    
  });
}
}

