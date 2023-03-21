import { LoginComponent } from "./components/login/login.component";
import { AppComponent } from "./app.component";
import { NgModule, Component } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { RegisterComponent } from "./components/register/register.component";
import { HomeComponent } from "./components/home/home.component";

const routes: Routes = [
  { path: "home", title: "Home ", component: HomeComponent },
  { path: "login", title: "Login", component: LoginComponent },
  { path: "register", title: "Register", component: RegisterComponent },

  //redirect
  { path: "**", redirectTo: "home" },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
