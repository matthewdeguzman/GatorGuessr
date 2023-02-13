import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
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
  usernameRequirements = new FormControl('',[Validators.required, Validators.minLength(4)]);
  passwordRequirements = new FormControl('',[
    Validators.required, 
    Validators.minLength(8),
    Validators.pattern(/^(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{8,}$/)
  ]);
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
  public showConfirmPassword: boolean = false;

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
  public toggleConfirmPassWordVisivility(): void {
    this.showConfirmPassword = !this.showConfirmPassword;
  }
}
