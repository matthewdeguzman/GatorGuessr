import { Loader } from "@googlemaps/js-api-loader";
import { Component } from "@angular/core";
//import { GoogleMapsModule } from "@angular/google-maps";

@Component({
  selector: "app-landing-page",
  templateUrl: "./landing-page.component.html",
  styleUrls: ["./landing-page.component.css"],
})
export class LandingPageComponent {
  title = "GatorMap";

  ngOnInit(): void {
    let loader = new Loader({
      apiKey: "",
      version: "weekly",
    });
    loader.load().then(() => {
      new google.maps.Map(document.getElementById("Gmap") as HTMLElement, {
        center: { lat: 29.643946, lng: -82.355659 },
        zoom: 6,
      });
    });
  }
}
