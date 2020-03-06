import { Component, OnInit } from '@angular/core';
import { DataService } from '../../core/data.service';

@Component({
  selector: 'experience-snippet',
  templateUrl: './experience-snippet.component.html',
  styleUrls: ['./experience-snippet.component.scss']
})
export class ExperienceSnippetComponent implements OnInit {
  page:string = 'experience';
  experiences:Array<any>;
  jobs: any[];
  founder: any[];
  projects: any[];

  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getExperience();
    this.getJobs();
    this.getFounder();
    this.getProject();
  }

  getExperience(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
    })
  }
  getJobs(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
      this.jobs = this.experiences.filter(edu => edu.type == 'book')
    })
  }
  getFounder(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
      this.founder = this.experiences.filter(online => online.type == 'online')
    })
  }
  getProject(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
      this.projects = this.experiences.filter(school => school.type == 'school')
    })
  }

}
