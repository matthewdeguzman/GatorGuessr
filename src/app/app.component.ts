import { Component, OnInit } from "@angular/core";
import { BannerComponent } from "./components/banner/banner.component";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"],
})
export class AppComponent {
  title = "GatorGuessr";
  theme = localStorage.getItem("selectedValue") || "lightMode";
}
