import { Component, OnInit, AfterContentInit, AfterViewChecked } from '@angular/core';
import { DataService } from './../core/data.service';
@Component({
  selector: 'page-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.scss']
})
export class EducationComponent implements OnInit, AfterViewChecked {
  page:string = 'education';
  educations:Array<any>;
  books: any[];
  onlines: any[];
  schools: any[];
  constructor(private data:DataService) { }

  ngOnInit(){
    this.getEducation();    
  }
  ngAfterViewChecked(): void {
  }

  getEducation(){
    this.data.getEducation().subscribe(data =>{
      this.educations = data['response']['data']
    })
  }
  // getBooks(){
  //   this.data.getEducation().subscribe(data =>{
  //     this.educations = data['response']['data']
  //     this.books = this.educations.filter(edu => edu.type == 'book')
  //   })
  // }
  // getOnline(){
  //   this.data.getEducation().subscribe(data =>{
  //     this.educations = data['response']['data']
  //     this.onlines = this.educations.filter(online => online.type == 'online')
  //   })
  // }
  // getSchools(){
  //   this.data.getEducation().subscribe(data =>{
  //     this.educations = data['response']['data']
  //     this.schools = this.educations.filter(school => school.type == 'school')
  //   })
  // }
}
