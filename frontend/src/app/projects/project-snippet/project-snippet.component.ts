import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'project-snippet',
  templateUrl: './project-snippet.component.html',
  styleUrls: ['./project-snippet.component.scss']
})
export class ProjectSnippetComponent implements OnInit {
  @Input() name:string;
  @Input() url:string;
  @Input() position:string;
  @Input() description:string;
  @Input() type_two:string;
  @Input() img:string;
  @Input() bgImg:string;
  @Input() location:string;
  
  constructor() { }

  ngOnInit(): void {
  }

}
