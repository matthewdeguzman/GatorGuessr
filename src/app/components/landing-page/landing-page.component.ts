import { Loader } from "@googlemaps/js-api-loader";
import { Component } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";

//import { GoogleMapsModule } from "@angular/google-maps";

@Component({
  selector: "app-landing-page",
  templateUrl: "./landing-page.component.html",
  styleUrls: ["./landing-page.component.css"],
})
export class LandingPageComponent {
  title = "GatorMap";
  string = "";
  // content: string;
  // private _jsonURL = "src/assets/env.json";
  // constructor(private http: HttpClient) {
  //   this.getJSON().subscribe((data) => {
  //     console.log(data);
  //   });
  // }
  // public getJSON(): Observable<any> {
  //   return this.http.get(this._jsonURL);
  // }

  // getTextFile() {
  //   this.http
  //     .get("/env.txt", {
  //       responseType: "text",
  //     })
  //     .subscribe((data) => {
  //       this.content = data;
  //       //console.log(this.content);
  //     });
  // }

  // center: LatLngLiteral = { lat: 29.6516, lng: -82.3248 };

  ngOnInit(): void {
    //this.getTextFile();
    let loader = new Loader({
      apiKey: this.string,
      version: "weekly",
    });

    loader.load().then(() => {
      new google.maps.StreetViewPanorama(
        document.getElementById("Smap") as HTMLElement,
        {
          position: { lat: 29.652031, lng: -82.342953 },
          pov: {
            heading: 34,
            pitch: 10,
          },
        }
      );
      new google.maps.Map(document.getElementById("Gmap") as HTMLElement, {
        center: { lat: 29.653288, lng: -82.338712 },
        zoom: 5,
        fullscreenControl: false,
        restriction: {
          latLngBounds: {
            //North: nw 53rd ave
            east: -82.263414,
            north: 29.769872,
            south: 29.602758,
            west: -82.420207,
          },
          strictBounds: false,
        },
      });
    });
  }
}
