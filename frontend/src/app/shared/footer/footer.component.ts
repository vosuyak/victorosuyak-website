import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'shared-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.scss']
})
export class FooterComponent implements OnInit {
  @Input() more: string;
  @Input() prev: string;
  @Input() next: string;
  constructor() { }

  ngOnInit(): void {
  }

}
