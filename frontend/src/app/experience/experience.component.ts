import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'page-experience',
  templateUrl: './experience.component.html',
  styleUrls: ['./experience.component.scss']
})
export class ExperienceComponent implements OnInit {
  fullPage:string = 'experience';

  constructor() { }

  ngOnInit(): void {
  }

}
