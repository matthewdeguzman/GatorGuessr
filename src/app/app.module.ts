import { LoginComponent } from "./components/login/login.component";
import { RegisterComponent } from "./components/register/register.component";
import { PageNotFoundComponent } from "./components/page-not-found/page-not-found.component";
import { LandingPageComponent } from "./components/landing-page/landing-page.component";

import { NgModule, Component } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";

import { RouterModule, Routes } from "@angular/router";
import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { MaterialModule } from "./material.module";
import { HomeComponent } from "./components/home/home.component";
import { BannerComponent } from "./components/banner/banner.component";

import { IssueService } from "./services/issue.service";
import { HttpClientModule } from "@angular/common/http";
import { CookieService } from "ngx-cookie-service";

//Don't know if we will still need this
/*const routes: Routes = [
  {
    path: "login",
    title: "Login",
    //canActivate: [AppGuard],
    component: LoginComponent,
  },
  {
    path: "landing-page",
    title: "GatorGuessr",
    component: LandingPageComponent,
  },
  { path: "register", title: "Register", component: RegisterComponent },
  {
    path: "page-not-found",
    title: "404 Error",
    component: PageNotFoundComponent,
  },

  //redirect
  { path: "", redirectTo: "home", pathMatch: "full" },
  { path: "**", redirectTo: "page-not-found" },
  // Add additional routes here
];
*/

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    BannerComponent,
    RegisterComponent,
    HomeComponent,
    LandingPageComponent,
    PageNotFoundComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MaterialModule,
    //RouterModule.forRoot(routes),
  ],
  providers: [IssueService, CookieService],
  bootstrap: [AppComponent],
  exports: [RouterModule],
})
export class AppModule {}
