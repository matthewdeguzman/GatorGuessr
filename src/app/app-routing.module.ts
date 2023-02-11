import { BannerComponent } from './components/banner/banner.component';
import { LoginComponent } from './components/login/login.component';
import { AppComponent } from './app.component';
import { NgModule, Component } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './components/register/register.component';

const routes: Routes = [
  { path: 'app', title:'GatorGuessr',component:AppComponent},
  { path: 'banner', title:'Banner',component: BannerComponent},
  { path: 'login', title:'Login', component: LoginComponent },
  { path: 'register', title:'Register', component: RegisterComponent},

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
