import { AppComponent } from "./../../app.component";
import { Component, HostBinding, OnInit } from "@angular/core";
import { IssueService } from "src/app/services/issue.service";
import { MatDialog, MatDialogConfig } from "@angular/material/dialog";
import { AccountComponent } from "../account/account.component";

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.scss"],
})
export class BannerComponent implements OnInit {
  selectedValue = "lightMode";
  selectedTheme = "light_mode";
  username = "Guest";
  score = 0;

  constructor(
    private AppComponent: AppComponent,
    private IssueService: IssueService,
    private dialog: MatDialog
  ) {}

  ngOnInit() {
    this.selectedTheme = localStorage.getItem("selectedTheme") || "light_mode";
    this.selectedValue =
      this.selectedTheme === "light_mode" ? "lightMode" : "darkMode";
    this.updateBanner();
  }

  openDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.autoFocus = true;

    this.dialog.open(AccountComponent, dialogConfig);
  }

  // Login methods
  updateBanner() {
    this.username = localStorage.getItem("username") || "Guest";
    if (this.username !== "null") {
      this.IssueService.getUserScore(this.username).subscribe((data) => {
        this.score = data;
      });
    }
  }

  // Theme methods
  getTheme() {
    return this.selectedValue;
  }

  toggleTheme() {
    if (this.selectedValue === "lightMode") {
      this.selectedValue = "darkMode";
      this.selectedTheme = "dark_mode";
    } else {
      this.selectedValue = "lightMode";
      this.selectedTheme = "light_mode";
    }
    localStorage.setItem("selectedTheme", this.selectedTheme);
    localStorage.setItem("selectedValue", this.selectedValue);
    this.AppComponent.theme = this.selectedValue;
  }
}
