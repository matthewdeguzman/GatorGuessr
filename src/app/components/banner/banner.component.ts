import { AppComponent } from "./../../app.component";
import { Component, HostBinding, OnInit } from "@angular/core";
import { IssueService } from "src/app/services/issue.service";
import { MatDialog, MatDialogConfig } from "@angular/material/dialog";
import { AccountComponent } from "../account/account.component";
import { DeleteComponent } from "../delete/delete.component";
import { Router } from "@angular/router";

interface User {
  ID: number;
  Username: string;
  Password: string;
  Score: number;
  CreatedAt: string;
  UpdatedAt: string;
}

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.scss"],
})
export class BannerComponent implements OnInit {
  selectedValue = "lightMode";
  selectedTheme = "light_mode";
  username = "Guest";
  score: number;
  highscore: number;

  constructor(
    private AppComponent: AppComponent,
    private IssueService: IssueService,
    private dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit() {
    this.selectedTheme = localStorage.getItem("selectedTheme") || "light_mode";
    this.selectedValue =
      this.selectedTheme === "light_mode" ? "lightMode" : "darkMode";
    this.updateBanner();
  }

  openDeleteDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.autoFocus = true;

    this.dialog.open(DeleteComponent, dialogConfig);
  }

  openAccountDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.autoFocus = true;

    this.dialog.open(AccountComponent, dialogConfig);
  }
  updateScore(score: number) {
    this.score = score;
    if (this.score > this.highscore) {
      this.updateHighScore();
    }
  }
  updateHighScore() {
    this.highscore = this.score;
    const temp = localStorage.getItem("username");
    if (temp != null) {
      this.IssueService.updateScore(temp, this.score).subscribe((res) => {
        console.log("Score updated");
        console.log(res);
      });
    }
  }

  // Login methods
  updateBanner() {
    this.username = localStorage.getItem("username") || "null";
    if (this.username !== "null") {
      this.IssueService.getUserScore(this.username).subscribe((data) => {
        this.highscore = (data as User).Score;
      });
    }
  }

  // Theme methods
  getTheme() {
    return this.selectedValue;
  }
  signOut() {
    localStorage.removeItem("username");
    this.router.navigate(["/home"]);
  }
  loggedIn() {
    return localStorage.getItem("username") != null;
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
    this.AppComponent.theme = this.selectedValue;
  }
}
