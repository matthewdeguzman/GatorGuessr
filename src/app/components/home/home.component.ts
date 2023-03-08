import { Component } from '@angular/core';

export interface Tile {
  cols: number;
  title: string;
  info: string;
  color: string;
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  tiles: Tile[] = [
    {title: 'About', cols: 2, color: 'lightblue', info: 'This is a test'},
    {title: 'Leaderboard', cols: 1, color: 'lightgreen', info: 'This is a test'},
  ];
}

