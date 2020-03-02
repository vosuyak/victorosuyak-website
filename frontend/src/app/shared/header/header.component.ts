import { Component, OnInit, OnChanges } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'shared-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, OnChanges {


  constructor(private location: Location) { }

  ngOnChanges() {
    this.checkIfHome();
  }

  ngOnInit(): void {
  }

  checkIfHome(){
    let path = this.location.path();
    return  path == "/home"? false:true;
  }

  back() {
    this.location.back();
  }
}
