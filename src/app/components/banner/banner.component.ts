import { AppComponent } from "./../../app.component";
import { Component, HostBinding, OnInit } from "@angular/core";

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.scss"],
})
export class BannerComponent implements OnInit {
  selectedValue = "lightMode";
  selectedTheme = "light_mode";

  constructor(private AppComponent: AppComponent) {}

  ngOnInit() {
    this.selectedTheme = localStorage.getItem("selectedTheme") || "light_mode";
    this.selectedValue =
      this.selectedTheme === "light_mode" ? "lightMode" : "darkMode";
  }
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
