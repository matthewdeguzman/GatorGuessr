import { Location } from "@angular/common";
import { Loader } from "@googlemaps/js-api-loader";
import { Component } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
// const { StreetViewPreference } = (await google.maps.importLibrary(
//   "streetView"
// )) as google.maps.StreetViewLibrary;

@Component({
  selector: "app-landing-page",
  templateUrl: "./landing-page.component.html",
  styleUrls: ["./landing-page.component.css"],
})
export class LandingPageComponent {
  title = "GatorMap";
  string = "";
  lat = this.randomLat();
  long = this.randomLong();
  invalidLoc: boolean = true;

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
  //    });
  // }

  // center: LatLngLiteral = { lat: 29.6516, lng: -82.3248 };
  randomLat() {
    return Math.random() * (29.769872 - 29.602758) + 29.769872;
  }
  randomLong() {
    return Math.random() * (-82.263414 - -82.420207) + -82.263414;
  }

  ngOnInit(): void {
    let loader = new Loader({
      apiKey: this.string,
      version: "weekly",
    });

    loader.load().then(() => {
      const pano = new google.maps.StreetViewPanorama(
        document.getElementById("Smap") as HTMLElement,
        {
          disableDefaultUI: false,
          //position: { lat: this.lat, lng: this.long },

          pov: {
            heading: 34,
            pitch: 10,
          },
        }
      );
      function RandomLoc(
        callback:
          | ((
              a: google.maps.StreetViewPanoramaData | null,
              b: google.maps.StreetViewStatus
            ) => void)
          | undefined
      ) {
        var lat = Math.random() * (29.676191 - 29.616823) + 29.616823;
        var long = Math.random() * (-82.295573 - -82.398928) + -82.398928;
        var cr = new google.maps.LatLng(lat, long);
        var sStatus = new google.maps.StreetViewService();
        sStatus.getPanorama({ location: cr, radius: 10 }, callback);
      }
      const HandlePanoramaData = (data: any, status: string) => {
        if (status === "OK") {
          console.log("valid panorama");
          pano.setPosition({
            lat: data.location.latLng.lat(),
            lng: data.location.latLng.lng(),
          });
        } else {
          console.log("invalid panorama");
          RandomLoc(HandlePanoramaData);
        }
      };
      RandomLoc(HandlePanoramaData);
      // var cr = new google.maps.LatLng(this.lat, this.long);
      // //streetView.getPanorama({ location: cr, radius: 50 });
      // var sStatus = new google.maps.StreetViewService();

      // sStatus.getPanorama({ location: cr, radius: 1000 }, (data, status) => {
      //   if (status === "OK") {
      //     console.log("valid panorama");
      //   } else {
      //     console.log("invalid panorama");

      //     this.invalidLoc = true;
      //     this.lat = this.randomLat();
      //     this.long = this.randomLong();
      //     cr = new google.maps.LatLng(this.lat, this.long);
      //   }
      // });

      new google.maps.Map(document.getElementById("Gmap") as HTMLElement, {
        center: { lat: 29.653288, lng: -82.338712 },
        zoom: 7,
        disableDefaultUI: true,
        restriction: {
          latLngBounds: {
            //North: nw 53rd ave
            east: -82.295573,
            north: 29.676191,
            south: 29.616823,
            west: -82.398928,
          },
          strictBounds: false,
        },
      });
    });
    //StreetViewService().getPanorama
  }
}
