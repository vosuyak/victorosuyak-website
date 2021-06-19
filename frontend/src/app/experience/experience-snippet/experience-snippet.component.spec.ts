import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ExperienceSnippetComponent } from './experience-snippet.component';

describe('ExperienceSnippetComponent', () => {
  let component: ExperienceSnippetComponent;
  let fixture: ComponentFixture<ExperienceSnippetComponent>;

  beforeEach(waitForAsync(() => {
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
