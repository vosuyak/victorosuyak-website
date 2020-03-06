import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(private http:HttpClient) { }

  getEducation(){
    let url = './../../assets/pages/education.json';
    return this.http.get(url)
  }

  getSkill(){
    let url = './../../assets/pages/skill.json';
    return this.http.get(url)
  }
}
