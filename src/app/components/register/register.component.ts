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
  usernameRequirements = new FormControl('',[Validators.required, Validators.minLength(4)]);
  passwordRequirements = new FormControl('',[
    Validators.required, 
    Validators.minLength(8),
    Validators.pattern(/^(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{8,}$/)
  ]);
  constructor(private IssueService: IssueService) { }
  // @Input() error: string | null;
  @Output() submitEM = new EventEmitter();
  
  public showPassword: boolean = false;
  public showSubmit: boolean = false;
  public showConfirmPassword: boolean = false;

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
  public toggleConfirmPassWordVisivility(): void {
    this.showConfirmPassword = !this.showConfirmPassword;
  }
  submitRegistration(username:string,password:string)
  {
    this.IssueService.getUsersWithUsername(username).subscribe((res: User) => {
      if(res.ID!=0){
        console.log("User already exists");
      }
      else{
        this.IssueService.createUser(username,password).subscribe((res) => {
          console.log(res);
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
