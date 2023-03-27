import { PageNotFoundComponent } from "./components/page-not-found/page-not-found.component";
import { LandingPageComponent } from "./components/landing-page/landing-page.component";
import { LoginComponent } from "./components/login/login.component";
import { AppComponent } from "./app.component";
import { NgModule, Component } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { RegisterComponent } from "./components/register/register.component";
import { HomeComponent } from "./components/home/home.component";
import { AuthGuard } from "./guards/auth/auth.guard";

const routes: Routes = [
  { path: "home", title: "Home", component: HomeComponent },
  {
    path: "login",
    title: "Login",
    component: LoginComponent,
    children: [
      {
        canActivate: [AuthGuard],
        path: "landing-page",
        title: "GatorGuessr",
        component: LandingPageComponent,
      },
    ],
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
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
