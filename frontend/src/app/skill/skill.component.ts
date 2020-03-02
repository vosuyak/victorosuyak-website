import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'page-skill',
  templateUrl: './skill.component.html',
  styleUrls: ['./skill.component.scss']
})
export class SkillComponent implements OnInit {
  fullPage:string = 'skill';

  constructor() { }

  ngOnInit(): void {
  }

}
