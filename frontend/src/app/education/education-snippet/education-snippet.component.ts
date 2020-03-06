import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/core/data.service';

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
  schools: any[];
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
