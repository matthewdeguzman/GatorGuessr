import { LeaderboardService } from "./leaderboard/leaderboard.service";
import { Injectable } from "@angular/core";
import { UserService } from "./user/user.service";
import { HttpResponse } from "@angular/common/http";
import { Observable } from "rxjs";
import { CookiesService } from "./cookies/cookies.service";

@Injectable({
  providedIn: "root",
})
export class IssueService {
  constructor(
    private userService: UserService,
    private leaderboardService: LeaderboardService,
    private cookiesService: CookiesService
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
  getUserScore(username: string) {
    return this.userService.getUserScore(username);
  }
  deleteUser(username: string) {
    return this.userService.deleteUser(username);
  }
  updateUser(username: string, password: string) {
    return this.userService.updateUser(username, password);
  }

  // Leaderboard methods
  getLeaderboard() {
    return this.leaderboardService.getLeaderboard();
  }

  // Cookie methods
  setCookie(username: string) {
    return this.cookiesService.getCookie(username);
  }
  verifyCookie() {
    return this.cookiesService.validateCookie();
  }

  // Api Key methods
  getApiKey() {
    return this.userService.getApiKey();
  }
}
