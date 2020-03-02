import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'page-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.scss']
})
export class EducationComponent implements OnInit {
  page:string = 'education';

  constructor() { }

  ngOnInit(): void {
  }

}
