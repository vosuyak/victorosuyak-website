import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ExperienceSnippetComponent } from './experience-snippet.component';

describe('ExperienceSnippetComponent', () => {
  let component: ExperienceSnippetComponent;
  let fixture: ComponentFixture<ExperienceSnippetComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ExperienceSnippetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ExperienceSnippetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
