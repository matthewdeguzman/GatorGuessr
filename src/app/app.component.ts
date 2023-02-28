import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { IssueService } from './issue.service';
//import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'GatorGuessr';
  constructor(private IssueService: IssueService) { }
  ngOnInit(){
  }
}