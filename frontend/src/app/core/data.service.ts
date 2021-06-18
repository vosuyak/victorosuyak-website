import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

var educationUrl;

@Injectable({
  providedIn: 'root'
})

export class DataService {
  educationUrl = '';

  constructor(private http:HttpClient) { 
    this.http.get('/assets/pages/endpoints.json').subscribe(
      data =>{
            educationUrl = data['education']   
      }
    )
  }

  getEducation(){
    let url = '/assets/pages/education.json';
    return this.http.get(url)
  }

  getExperience(){
    let url = '/assets/pages/experience.json';
    return this.http.get(url)
  }

  getSkill(){
    let url = '/assets/pages/skill.json';
    return this.http.get(url)
  }
}
