import { Component } from "@angular/core";
import { MatDialogRef } from "@angular/material/dialog";
import { IssueService } from "src/app/services/issue.service";

@Component({
  selector: "app-account",
  templateUrl: "./account.component.html",
  styleUrls: ["./account.component.scss"],
})
export class AccountComponent {
  constructor(
    private IssueService: IssueService,
    private dialogRef: MatDialogRef<AccountComponent>
  ) {}

  deleteAccount() {
    this.IssueService.deleteUser("test");
    this.closeDialog();
  }

  closeDialog() {
    this.dialogRef.close();
  }
}
