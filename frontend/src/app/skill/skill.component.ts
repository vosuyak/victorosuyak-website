import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'page-skill',
  templateUrl: './skill.component.html',
  styleUrls: ['./skill.component.scss']
})
export class SkillComponent implements OnInit {
  page:string = 'skill';

  constructor() { }

  ngOnInit(): void {
  }

}
