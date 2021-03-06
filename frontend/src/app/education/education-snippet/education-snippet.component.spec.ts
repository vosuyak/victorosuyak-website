import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EducationSnippetComponent } from './education-snippet.component';

describe('EducationSnippetComponent', () => {
  let component: EducationSnippetComponent;
  let fixture: ComponentFixture<EducationSnippetComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EducationSnippetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EducationSnippetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
