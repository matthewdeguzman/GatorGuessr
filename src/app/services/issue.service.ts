import { LeaderboardService } from "./leaderboard/leaderboard.service";
import { Injectable } from "@angular/core";
import { UserService } from "./user/user.service";

@Injectable({
  providedIn: "root",
})
export class IssueService {
  constructor(
    private userService: UserService,
    private LeaderboardService: LeaderboardService
  ) {}

  // User methods
  validateUser(username: string, password: string) {
    return this.userService.validateUser(username, password);
  }
  createUser(username: string, password: string) {
    return this.userService.createUser(username, password);
  }
  getUser(username: string) {
    return this.userService.getUser(username);
  }

  // Leaderboard methods
  getLeaderboard() {
    return this.LeaderboardService.getLeaderboard();
  }

  // Api Key methods
  getApiKey() {
    return this.userService.getApiKey();
  }
}
