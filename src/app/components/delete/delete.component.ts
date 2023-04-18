import { Component } from "@angular/core";
import { MatDialogRef } from "@angular/material/dialog";
import { IssueService } from "src/app/services/issue.service";

@Component({
  selector: "app-delete",
  templateUrl: "./delete.component.html",
  styleUrls: ["./delete.component.scss"],
})
export class DeleteComponent {
  constructor(
    private IssueService: IssueService,
    private dialogRef: MatDialogRef<DeleteComponent>
  ) {}

  deleteAccount() {
    this.IssueService.deleteUser("test");
    this.closeDialog();
  }

  closeDialog() {
    this.dialogRef.close();
  }
}
