import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'experience-snippet',
  templateUrl: './experience-snippet.component.html',
  styleUrls: ['./experience-snippet.component.scss']
})
export class ExperienceSnippetComponent implements OnInit {
  page:string = 'experience';

  constructor() { }

  ngOnInit(): void {
  }

}
