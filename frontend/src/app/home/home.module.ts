import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import { EducationComponent } from '../education/education.component';
import { SkillComponent } from '../skill/skill.component';
import { ExperienceComponent } from '../experience/experience.component';
import { FooterComponent } from '../shared/footer/footer.component';
import { EducationSnippetComponent } from '../education/education-snippet/education-snippet.component';
import { ExperienceSnippetComponent } from '../experience/experience-snippet/experience-snippet.component';
import { SkillSnippetComponent } from '../skill/skill-snippet/skill-snippet.component';



@NgModule({
  declarations: [
    HomeComponent,
    EducationComponent,
    SkillComponent,
    ExperienceComponent,
    FooterComponent,
    EducationSnippetComponent,
    ExperienceSnippetComponent,
    SkillSnippetComponent,
  ],
  imports: [
    CommonModule
  ],
  exports:[
    HomeComponent
  ]
})
export class HomeModule { }
