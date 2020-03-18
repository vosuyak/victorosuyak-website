import { Component, OnInit, OnChanges, Input, ViewChild, ElementRef, AfterViewInit } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'shared-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, AfterViewInit, OnChanges {
  @Input() page:string
  @ViewChild('menu') menu:ElementRef;
  isMenuOpen:boolean;

  constructor(private location: Location) {
    this.isMenuOpen = false;
   }

  ngOnChanges() {
    this.checkIfHome();
  }

  ngOnInit(): void {

  }
  ngAfterViewInit() {
    
  }
  checkIfHome(){
    let path = this.location.path();
    return  path == "/home"? false:true;
  }

  back() {
    this.location.back();
  }

  closeMenu(){
    
  }
  openMenu(){
    
    if(this.isMenuOpen == false){
      this.isMenuOpen = true;
      this.menu.nativeElement.style.display = "block"

    }else{
      this.isMenuOpen = false;
      this.menu.nativeElement.style.display = "none"
    }
    
  }
  removeShowMoreBTN(){
    if ("/"+this.page == this.location.path()){
      return false
    } else if(this.page == undefined){
      return false  
    }else{
      return true
    }
  }
}
