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

  ngOnInit(): void {
    //this.getTextFile();
    let loader = new Loader({
      apiKey: this.string,
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
