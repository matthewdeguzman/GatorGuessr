import { Component, OnInit } from "@angular/core";

@Component({
  selector: "app-banner",
  templateUrl: "./banner.component.html",
  styleUrls: ["./banner.component.css"],
})
export class BannerComponent implements OnInit {
  selectedValue = "Light";

  ngOnInit() {
    this.selectedValue = localStorage.getItem("selectedTheme") || "Light";
  }

  onToggleChange(event: any) {
    localStorage.setItem("selectedTheme", event.value);
  }
}
