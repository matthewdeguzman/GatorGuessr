import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { IssueService } from '../../issue.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit{
  registerForm = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  });
  constructor(private IssueService: IssueService) { }
  ngOnInit(){
    this.IssueService.getUsers().subscribe((issue) => {
      console.log(issue);
    })
  }
submitRegister() {
  
  }
  // @Input() error: string | null;
  @Output() submitEM = new EventEmitter();
  

  public showPassword: boolean = false;

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
}
