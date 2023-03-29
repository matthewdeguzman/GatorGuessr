import { Component, OnInit } from "@angular/core";
import { IssueService } from "src/app/services/issue.service";

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
  styleUrls: ["./home.component.css"],
})
export class HomeComponent {
  constructor(private IssueService: IssueService) {}
  displayedColumns: string[] = ["name", "score"];
  leaderboardArray: any;
  ngOnInit() {
    this.IssueService.getLeaderboard().subscribe((data) => {
      this.leaderboardArray = data;
    });
  }
}
