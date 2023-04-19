import { Location } from "@angular/common";
import { Loader } from "@googlemaps/js-api-loader";
import { Component } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { IssueService } from "src/app/services/issue.service";
import { BannerComponent } from "../banner/banner.component";
import { clear } from "console";

interface User {
  ID: number;
  Username: string;
  Password: string;
  Score: number;
  CreatedAt: string;
  UpdatedAt: string;
}

@Component({
  selector: "app-landing-page",
  templateUrl: "./landing-page.component.html",
  styleUrls: ["./landing-page.component.scss"],
})
export class LandingPageComponent {
  title = "GatorMap";
  string = "";
  invalidLoc: boolean = true;
  time = 60;
  userLat: number = 0;
  userLng: number = 0;
  streetViewLat: number = 0;
  streetViewLng: number = 0;
  navMap: any;
  loader: any;
  timeContinue: boolean = true;
  score: number;
  canClick: boolean = true;
  oldScore: number;


  setStreetView(latLng: google.maps.LatLng) {
    this.streetViewLat = latLng.lat();
    this.streetViewLng = latLng.lng();
  }
  setUserLoc(latLng: google.maps.LatLng) {
    this.userLat = latLng.lat();
    this.userLng = latLng.lng();
  }
  nextButton() {
    this.googleMapsLoad();
    this.timeContinue = true;
    this.time = 60;
    this.countDown();
  }

  submit() {
    this.canClick = false;
    this.timeContinue = false;
    console.log("Submit button clicked");
    const distance = Math.sqrt(
      Math.pow(this.userLat - this.streetViewLat, 2) +
        Math.pow(this.userLng - this.streetViewLng, 2)
    );
    const maxDistance = 0.04; // Maximum allowed distance for maximum points (adjust as needed)
    const maxPoints = 1000; // Maximum points for a perfect guess
    const minPoints = 0; // Minimum points for a guess beyond the maximum allowed distance
    const newScore = Math.round(
      Math.max(maxPoints - (distance / maxDistance) * maxPoints, minPoints)
    );
    this.BannerComponent.updateScore(newScore);
    console.log("Score: " + newScore);
    var orginalLocation = new google.maps.Marker({
      position: { lat: this.streetViewLat, lng: this.streetViewLng },
      map: this.navMap,
      icon: "http://maps.google.com/mapfiles/kml/paddle/purple-stars.png",
    });
    const lineDistance = new google.maps.Polyline({
      path: [
        { lat: this.userLat, lng: this.userLng },
        { lat: this.streetViewLat, lng: this.streetViewLng },
      ],
      strokeColor: "#FF5733",
      strokeOpacity: 1.0,
      strokeWeight: 4,
    });
    lineDistance.setMap(this.navMap);
  }

  countDown() {
    this.time--; // decrements by one second
    if (this.time > 0 && this.timeContinue == true) {
      setTimeout(() => {
        this.countDown();
      }, 1000); //decrement one second
    } else {
      // If remaining time reaches 0, call the submit function
      this.submit();
    }
  }
  async timer() {
    await new Promise((resolve) => setTimeout(resolve, 800));
  }

  googleMapsLoad() {
    this.canClick = true;
    this.loader.load().then(() => {
      const pano = new google.maps.StreetViewPanorama(
        document.getElementById("Smap") as HTMLElement,
        {
          disableDefaultUI: false,
          addressControl: false,
          fullscreenControl: false,
          showRoadLabels: false,
          //position: { lat: this.lat, lng: this.long },

          pov: {
            heading: 34,
            pitch: 10,
          },
        }
      );
      const RandomLoc = (
        callback:
          | ((
              a: google.maps.StreetViewPanoramaData | null,
              b: google.maps.StreetViewStatus
            ) => void)
          | undefined
      ) => {
        var lat = Math.random() * (29.676191 - 29.616823) + 29.616823;
        var long = Math.random() * (-82.295573 - -82.398928) + -82.398928;
        // this.lat = lat;
        // this.long = long;
        var cr = new google.maps.LatLng(lat, long);
        this.setStreetView(cr);
        var sStatus = new google.maps.StreetViewService();
        sStatus.getPanorama(
          {
            location: cr,
            radius: 15,
            source: google.maps.StreetViewSource.OUTDOOR,
          },
          callback
        );
      };
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
      if (pano.getVisible() == false) RandomLoc(HandlePanoramaData);

      const navMap = new google.maps.Map(
        document.getElementById("Gmap") as HTMLElement,
        {
          center: { lat: 29.653288, lng: -82.338712 },
          zoom: 11,
          disableDefaultUI: true,
          mapTypeControl: false,

          // restriction: {
          //   latLngBounds: {
          //     //North: nw 53rd ave
          //     east: -82.295573,
          //     north: 29.676191,
          //     south: 29.616823,
          //     west: -82.398928,
          //   },
          //   strictBounds: false,
          // },
        }
      );
      this.navMap = navMap;
      var marker = new google.maps.Marker({
        position: null,
        map: navMap,
      });
      google.maps.event.addListener(
        navMap,
        "click",
        (e: { latLng: google.maps.LatLng }) => {
          if (this.canClick == true) placeMarker(e.latLng, navMap);
        }
      );
      const placeMarker = (
        Location: google.maps.LatLng,
        Map: google.maps.Map
      ) => {
        marker.setMap(navMap);
        marker.setPosition(Location);
        this.setUserLoc(Location);
      };
    });
  }

  constructor(
    private IssueService: IssueService,
    private BannerComponent: BannerComponent
  ) {}

  async ngOnInit() {
    this.BannerComponent.ngOnInit();
    this.IssueService.getApiKey().subscribe((res) => {
      this.string = res.body as string;
    });
    await this.timer();
    this.string = this.string.substring(1, this.string.length - 2);

    this.loader = new Loader({
      apiKey: this.string,
      version: "weekly",
    });
    this.time = 60;
    this.countDown();
    // setTimeout(() => {
    //   this.time = 60;
    //   this.countDown();
    // }, 1000); //decrements by 1000ms
    this.googleMapsLoad();
  }
}
