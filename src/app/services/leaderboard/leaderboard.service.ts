import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { map } from "rxjs/operators";

@Injectable({
  providedIn: "root",
})
export class LeaderboardService {
  private uri = "http://localhost:9000";
  constructor(private http: HttpClient) {}

  //Gets the leaderboard
  getLeaderboard() {
    return this.http.get(`${this.uri}/api/leaderboard/10/`).pipe(
      map((response) => {
        return response;
      })
    );
  }
}
