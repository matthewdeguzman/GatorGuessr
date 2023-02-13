import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { IssueService } from '../../issue.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit{
  loginForm = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  });
  constructor(private IssueService: IssueService) { }
  ngOnInit(){
    this.IssueService.getUsers().subscribe((issue) => {
      console.log(issue);
    })
  }
submitLogin() {
  
}
// @Input() error: string | null;
@Output() submitEM = new EventEmitter();
  
public showPassword: boolean = false;

public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
}

