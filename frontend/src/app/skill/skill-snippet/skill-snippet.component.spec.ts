import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { SkillSnippetComponent } from './skill-snippet.component';

describe('SkillSnippetComponent', () => {
  let component: SkillSnippetComponent;
  let fixture: ComponentFixture<SkillSnippetComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ SkillSnippetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SkillSnippetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
