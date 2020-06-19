import { Component, OnInit, Input } from '@angular/core';
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
  founders: any[];
  projects: any[];

  @Input() position:any;
  @Input() employer:any;
  @Input() image:any;
  @Input() start_date:any;
  @Input() end_date:any;
  @Input() description:any;
  @Input() shoutOuts:any;

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
      this.jobs = this.experiences.filter(exp => exp.type == 'job')
    })
  }
  getFounder(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
      this.founders = this.experiences.filter(exp => exp.type == 'founder')
      
    })
  }
  getProject(){
    this.data.getExperience().subscribe(data =>{
      this.experiences = data['response']['data']
      this.projects = this.experiences.filter(exp => exp.type == 'project')
    })
  }

}
