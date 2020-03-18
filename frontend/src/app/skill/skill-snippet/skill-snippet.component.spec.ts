import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SkillSnippetComponent } from './skill-snippet.component';

describe('SkillSnippetComponent', () => {
  let component: SkillSnippetComponent;
  let fixture: ComponentFixture<SkillSnippetComponent>;

  beforeEach(async(() => {
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
