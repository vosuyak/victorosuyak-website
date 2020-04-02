import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { PagenotfoundComponent } from './pagenotfound/pagenotfound.component';
import { EducationComponent } from './education/education.component';
import { HeaderComponent } from './shared/header/header.component';
import { SkillComponent } from './skill/skill.component';
import { ExperienceComponent } from './experience/experience.component';
import { FooterComponent } from './shared/footer/footer.component';
import { EducationSnippetComponent } from './education/education-snippet/education-snippet.component';
import { ExperienceSnippetComponent } from './experience/experience-snippet/experience-snippet.component';
import { SkillSnippetComponent } from './skill/skill-snippet/skill-snippet.component';
import { HomeModule } from './home/home.module';
import { NgxPageScrollModule } from 'ngx-page-scroll';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment';
import { ProjectsComponent } from './projects/projects.component';
import { ProjectSnippetComponent } from './projects/project-snippet/project-snippet.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    PagenotfoundComponent,
    EducationComponent,
    HeaderComponent,
    SkillComponent,
    ExperienceComponent,
    FooterComponent,
    EducationSnippetComponent,
    ExperienceSnippetComponent,
    SkillSnippetComponent,
    ProjectsComponent,
    ProjectSnippetComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    HomeModule,
    NgxPageScrollModule,
    ServiceWorkerModule.register('ngsw-worker.js', { enabled: environment.production })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
