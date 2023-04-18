import { Component } from "@angular/core";
import { MatDialogRef } from "@angular/material/dialog";

@Component({
  selector: "app-account",
  templateUrl: "./account.component.html",
  styleUrls: ["./account.component.scss"],
})
export class AccountComponent {
  constructor(private dialogRef: MatDialogRef<AccountComponent>) {}

  ngOnInit() {}

  close() {
    this.dialogRef.close();
  }
}
