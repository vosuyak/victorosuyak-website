import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'shared-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(private location: Location) { }

  ngOnInit(): void {
  }

  back() {
    this.location.back();
    let path = this.location.path();
    console.log('path: ', path);
    let isPath = this.location.isCurrentPathEqualTo(path,'/home');
    console.log('isPath: ', isPath);
  }
}
