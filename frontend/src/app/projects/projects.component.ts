import { Component, OnInit, Input } from '@angular/core';
import { DataService } from '../core/data.service';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.scss']
})
export class ProjectsComponent implements OnInit {
  page:string = 'projects';
  projects: any;


  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getProject();
  }

  getProject(){
    this.data.getExperience().subscribe(data =>{
      let array = data['response']['data']
      this.projects = array.filter(i => i.type == 'project')
    })
  }
}
