import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProjectsComponent } from './projects.component';
import { SocialLoginComponent } from './social-login/social-login.component';


const routes: Routes = [
  {
    path: '',
    pathMatch:"project"
  },  
  {
    path: 'project',
    component: ProjectsComponent,
    children :[
      {
        path: 'social-login',
        component: SocialLoginComponent
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectsRoutingModule { }
