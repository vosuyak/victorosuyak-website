import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/core/data.service';

@Component({
  selector: 'skill-snippet',
  templateUrl: './skill-snippet.component.html',
  styleUrls: ['./skill-snippet.component.scss']
})
export class SkillSnippetComponent implements OnInit {
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
