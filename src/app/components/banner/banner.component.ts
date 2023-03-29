import { Component, OnInit } from "@angular/core";
import { OverlayContainer } from "@angular/cdk/overlay";

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.scss"],
})
export class BannerComponent implements OnInit {
  selectedValue = true;
  selectedTheme = "light_mode";

  constructor(private overlay: OverlayContainer) {}
  ngOnInit() {
    this.selectedTheme = localStorage.getItem("selectedTheme") || "light_mode";
    this.selectedValue = this.selectedTheme === "light_mode" ? true : false;
  }

  toggleTheme() {
    if (this.selectedValue) {
      this.overlay.getContainerElement().classList.add("dark_mode");
      this.selectedValue = false;
      this.selectedTheme = "dark_mode";
    } else {
      this.overlay.getContainerElement().classList.remove("dark_mode");
      this.selectedValue = true;
      this.selectedTheme = "light_mode";
    }
    localStorage.setItem("selectedTheme", this.selectedTheme);
  }
}
