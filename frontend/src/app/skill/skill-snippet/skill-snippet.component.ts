import { Component, OnInit, Input } from '@angular/core';
import { DataService } from 'src/app/core/data.service';

@Component({
  selector: 'skill-snippet',
  templateUrl: './skill-snippet.component.html',
  styleUrls: ['./skill-snippet.component.scss']
})
export class SkillSnippetComponent implements OnInit {
  page:string = 'skill';
  skills: any;

  @Input() language:any;
  @Input() years_experience:any;
  @Input() img:any;
  @Input() background_color:any;
  @Input() end:any;
  @Input() type_two:any;
  
  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getSkill();
  }

  getSkill(){
    this.data.getSkill().subscribe(data =>{
      this.skills = data['response']['data']
      
    })
  }

}
