import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListComponent } from './list/list.component';
import { ViewComponent } from './view/view.component';
import { ErrorComponent } from './error/error.component';

const routes: Routes = [
  {
    path: 'list/:tags',
    component: ListComponent,
  },
  {
    path: 'view/:imgId',
    component: ViewComponent,
  },
  {
    path: '',
    redirectTo: '/list/',
    pathMatch: 'full'
  },
  { path: 'error',
    component: ErrorComponent
  },
  {
    path: '**',
    redirectTo: '/list/',
    pathMatch: 'full'
  },

]

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
