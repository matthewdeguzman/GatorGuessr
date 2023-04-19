import { PageNotFoundComponent } from "./components/page-not-found/page-not-found.component";
import { LandingPageComponent } from "./components/landing-page/landing-page.component";
import { LoginComponent } from "./components/login/login.component";
import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { RegisterComponent } from "./components/register/register.component";
import { HomeComponent } from "./components/home/home.component";
import { AuthGuard } from "./guards/auth/auth.guard";
import { LoginGuard } from "./guards/login/login.guard";

const routes: Routes = [
  {
    canActivate: [LoginGuard],
    path: "home",
    title: "Home",
    component: HomeComponent,
  },
  {
    canActivate: [LoginGuard],
    path: "login",
    title: "Login",
    component: LoginComponent,
  },
  {
    path: "landing-page",
    title: "GatorGuessr",
    canActivate: [AuthGuard],
    component: LandingPageComponent,
  },
  {
    canActivate: [LoginGuard],
    path: "register",
    title: "Register",
    component: RegisterComponent,
  },
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
