import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'page-experience',
  templateUrl: './experience.component.html',
  styleUrls: ['./experience.component.scss']
})
export class ExperienceComponent implements OnInit {
  page:string = 'experience';

  constructor() { }

  ngOnInit(): void {
  }

}
