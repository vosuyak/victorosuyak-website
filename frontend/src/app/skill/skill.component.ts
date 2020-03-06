import { Component, OnInit } from '@angular/core';
import { DataService } from '../core/data.service';

@Component({
  selector: 'page-skill',
  templateUrl: './skill.component.html',
  styleUrls: ['./skill.component.scss']
})
export class SkillComponent implements OnInit {
  page:string = 'skill';
  skills: any;

  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getSkill();
  }

  getSkill(){
    this.data.getSkill().subscribe(data =>{
      this.skills = data['response']['data']
      console.log('this.skills: ', this.skills);
    })
  }
}
