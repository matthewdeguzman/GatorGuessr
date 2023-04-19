import { Component } from "@angular/core";
import { MatDialogRef } from "@angular/material/dialog";
import { IssueService } from "src/app/services/issue.service";
import { FormGroup, FormControl } from "@angular/forms";

@Component({
  selector: "app-account",
  templateUrl: "./account.component.html",
  styleUrls: ["./account.component.scss"],
})
export class AccountComponent {
  currentPassword: string;
  newPassword: string;
  newUsername: string;
  public showPassword: boolean = false;

  updateForm = new FormGroup({
    newUsername: new FormControl(""),
    newPassword: new FormControl(""),
  });

  constructor(
    private IssueService: IssueService,
    private dialogRef: MatDialogRef<AccountComponent>
  ) {}
  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }
  closeDialog() {
    this.dialogRef.close();
  }
  updateUser(newUsername: string, newPassword: string) {
    this.IssueService.updateUser(newUsername, newPassword).subscribe((res) => {
      console.log(res);
    });

    this.closeDialog();
  }
}
