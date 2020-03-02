import { Component, OnInit, Input } from '@angular/core';
import { Location } from '@angular/common';
@Component({
  selector: 'shared-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.scss']
})
export class FooterComponent implements OnInit {
  @Input() more: string;
  @Input() prev: string;
  @Input() next: string;

  constructor(private location: Location) { }

  ngOnInit(): void {
  }

  checkIfHome(){
    let path = this.location.path();
    return  path == "/home"? true:false;
  }
}
