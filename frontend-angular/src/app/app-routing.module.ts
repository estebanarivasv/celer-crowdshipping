import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {DistributorComponent} from "./components/views/distributor/distributor.component";
import {SenderComponent} from "./components/views/sender/sender.component";
import {DashboardComponent} from "./components/views/dashboard/dashboard.component";
import {NewRequestComponent} from "./components/views/sender/new-request/new-request.component";
import {ViewRequestsComponent} from "./components/views/sender/view-requests/view-requests.component";
import {DetailedRequestComponent} from "./components/views/sender/detailed-request/detailed-request.component";

const routes: Routes = [
    {path: '', redirectTo: 'dashboard', pathMatch: 'full'},
    {path: 'dashboard', component: DashboardComponent},
    {
        path: 'sender', component: SenderComponent, children: [
            {
                path: '',
                component: ViewRequestsComponent,
            },
            {
                path: 'new-request',
                component: NewRequestComponent,
            },
            {
                path: 'requests/:id',
                component: DetailedRequestComponent,
            }
        ]
    },
    {path: 'distributor', component: DistributorComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
