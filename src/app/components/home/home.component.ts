import { Component } from "@angular/core";

export interface User {
  name: string;
  position: number;
  score: number;
}

const ELEMENT_DATA: User[] = [
  { position: 1, name: "Hydrogen", score: 1.0079 },
  { position: 2, name: "Helium", score: 4.0026 },
  { position: 3, name: "Lithium", score: 6.941 },
  { position: 4, name: "Beryllium", score: 9.0122 },
  { position: 5, name: "Boron", score: 10.811 },
  { position: 6, name: "Carbon", score: 12.0107 },
  { position: 7, name: "Nitrogen", score: 14.0067 },
  { position: 8, name: "Oxygen", score: 15.9994 },
  { position: 9, name: "Fluorine", score: 18.9984 },
  { position: 10, name: "Neon", score: 20.1797 },
];

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
  styleUrls: ["./home.component.css"],
})
export class HomeComponent {
  displayedColumns: string[] = ["position", "name", "score"];
  dataSource = ELEMENT_DATA;
}
