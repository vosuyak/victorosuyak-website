import { Component, OnInit, Input } from '@angular/core';
import { DataService } from './../core/data.service';

@Component({
  selector: 'page-experience',
  templateUrl: './experience.component.html',
  styleUrls: ['./experience.component.scss']
})
export class ExperienceComponent implements OnInit {
  page:string = 'experience';
  experiences:Array<any>;
  jobs: any[];
  founders: any[];
  projects: any[];

  @Input() employer:any;

  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getExperience();
  }

  getExperience(){
    this.data.getExperience().subscribe(data =>{
      let array = data['response']['data']
      this.experiences = array.filter(exp => exp.type == 'job')

    })
  }
  // getJobs(){
  //   this.data.getExperience().subscribe(data =>{
  //     this.experiences = data['response']['data']
  //     this.jobs = this.experiences.filter(exp => exp.type == 'job')
  //   })
  // }
  // getFounder(){
  //   this.data.getExperience().subscribe(data =>{
  //     this.experiences = data['response']['data']
  //     this.founders = this.experiences.filter(exp => exp.type == 'founder')
      
  //   })
  // }
  // getProject(){
  //   this.data.getExperience().subscribe(data =>{
  //     this.experiences = data['response']['data']
  //     this.projects = this.experiences.filter(exp => exp.type == 'project')
  //   })
  // }

}
