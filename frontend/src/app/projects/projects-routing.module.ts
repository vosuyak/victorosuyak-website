import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProjectsComponent } from './projects.component';
import { SocialLoginComponent } from './social-login/social-login.component';
import { TestComponent } from './test/test.component';


const routes: Routes = [
  {
    path: '',
    pathMatch:"full",
    redirectTo:'/project'
  },  
  {
    path: '',
    component: SocialLoginComponent,
    children :[
      {
        path: 'social-login',
        component: TestComponent 
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProjectsRoutingModule { }
