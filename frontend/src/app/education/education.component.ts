import { Component, OnInit } from '@angular/core';
import { DataService } from './../core/data.service';
@Component({
  selector: 'page-education',
  templateUrl: './education.component.html',
  styleUrls: ['./education.component.scss']
})
export class EducationComponent implements OnInit {
  page:string = 'education';

  constructor(private data:DataService) { }

  ngOnInit(): void {
    // this.getEducation();
  }

  getEducation(){
    this.data.getEducation().subscribe(data =>{
      data
      console.log('data: ', data);
    })
  }

}
