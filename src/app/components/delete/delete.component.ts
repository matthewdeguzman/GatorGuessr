import { Component } from "@angular/core";
import { MatDialogRef } from "@angular/material/dialog";
import { IssueService } from "src/app/services/issue.service";
import { Router } from "@angular/router";

@Component({
  selector: "app-delete",
  templateUrl: "./delete.component.html",
  styleUrls: ["./delete.component.scss"],
})
export class DeleteComponent {
  constructor(
    private IssueService: IssueService,
    private router: Router,
    private dialogRef: MatDialogRef<DeleteComponent>
  ) {}

  signOut() {
    localStorage.removeItem("username");
    this.router.navigate(["/home"]);
  }

  deleteAccount() {
    const username = localStorage.getItem("username");
    if (username) {
      this.IssueService.deleteUser(username).subscribe((res) => {
        console.log(res);
      });
      this.closeDialog();
      this.signOut();
    }
  }

  closeDialog() {
    this.dialogRef.close();
  }
}
