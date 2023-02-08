import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  readonly ROOT_URL='cen3031-server.mysql.database.azure.com'
  constructor(private http: HttpClient) {}
  title = 'GatorGuessr';
  getData() {
    this.http.get('cen3031-server.mysql.database.azure.com').subscribe(data => {
      console.log(data);
    });
  }
  sendData() {
    const User = { username: 'president', password: 'joebidenisreal'};
    this.http.post('https://cen3031-server.mysql.database.azure.com:3306', User).subscribe(data => {
      console.log(data);
    })
  }
}
