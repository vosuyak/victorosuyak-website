import { Component, OnInit, Input } from '@angular/core';
import { DataService } from 'src/app/core/data.service';
import { Education } from '../models/education';

@Component({
  selector: 'education-snippet',
  templateUrl: './education-snippet.component.html',
  styleUrls: ['./education-snippet.component.scss']
})
export class EducationSnippetComponent implements OnInit {
  page:string = 'education';
  educations:Array<any>;
  books: any[];
  onlines: any[];
  schools: any  [];
  education:Education;

  @Input() Education:Education;
  @Input() id:any;
  @Input() school:any;
  @Input() degree:any;
  @Input() start_date:any;
  @Input() end_date:any;
  @Input() city:any;
  @Input() state:any;
  @Input() online:any;
  @Input() website:any;
  @Input() price:any;
  @Input() type:any;
  @Input() courses:any;


  constructor(private data:DataService) { }

  ngOnInit(): void {
    this.getEducation();
    this.getBooks();
    this.getOnline();
    this.getSchools();
  }

  
  getEducation(){
    this.data.getEducation().subscribe(data =>{
      this.educations = data['response']['data']
    })
  }
  getBooks(){
    this.data.getEducation().subscribe(data =>{
      this.educations = data['response']['data']
      this.books = this.educations.filter(edu => edu.type == 'book')
    })
  }
  getOnline(){
    this.data.getEducation().subscribe(data =>{
      this.educations = data['response']['data']
      this.onlines = this.educations.filter(online => online.type == 'online')
    })
  }
  getSchools(){
    this.data.getEducation().subscribe(data =>{
      this.educations = data['response']['data']
      this.schools = this.educations.filter(school => school.type == 'school')
    })
  }
}
