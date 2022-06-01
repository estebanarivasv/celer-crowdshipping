import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {DistributorComponent} from "./components/views/distributor/distributor.component";
import {SenderComponent} from "./components/views/sender/sender.component";
import {DashboardComponent} from "./components/views/dashboard/dashboard.component";

const routes: Routes = [
    {path: '', redirectTo: 'dashboard', pathMatch: 'full'},
    {path: 'dashboard', component: DashboardComponent},
    {path: 'sender', component: SenderComponent},
    {path: 'distributor', component: DistributorComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
