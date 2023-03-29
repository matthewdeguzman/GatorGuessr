import { Component, OnInit } from "@angular/core";

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.scss"],
})
export class BannerComponent implements OnInit {
  selectedValue = "Light";
  selectedTheme = "light_mode";

  ngOnInit() {
    this.selectedTheme = localStorage.getItem("selectedTheme") || "light_mode";
    this.selectedValue = this.selectedTheme === "light_mode" ? "Light" : "Dark";
  }

  toggleTheme() {
    if (this.selectedValue === "Light") {
      this.selectedValue = "Dark";
      this.selectedTheme = "dark_mode";
    } else {
      this.selectedValue = "Light";
      this.selectedTheme = "light_mode";
    }
    localStorage.setItem("selectedTheme", this.selectedTheme);
  }
}
