import { Component, OnInit, Input, Output, EventEmitter } from "@angular/core";
import { IssueService } from "./issue.service";

type User = {
  ID: number;
  username: string;
  password: string;
  created: string;
  updated: string;
};

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent {
  Users: User[] = [];
  title = "GatorGuessr";
  constructor(private IssueService: IssueService) {}
}
