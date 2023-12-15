import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {DistributorComponent} from "./components/views/distributor/distributor.component";
import {SenderComponent} from "./components/views/sender/sender.component";
import {DashboardComponent} from "./components/views/dashboard/dashboard.component";
import {NewRequestComponent} from "./components/views/sender/new-request/new-request.component";
import {ViewShippingsComponent} from "./components/views/sender/view-shippings/view-shippings.component";
import {DetailedRequestComponent} from "./components/views/sender/detailed-request/detailed-request.component";
import {SearchRequestsComponent} from "./components/views/distributor/search-requests/search-requests.component";
import {ViewDeliveriesComponent} from "./components/views/distributor/view-deliveries/view-deliveries.component";
import {
    NewShippingOfferComponent
} from "./components/views/distributor/search-requests/new-shipping-offer/new-shipping-offer.component";
import {DetailedDeliveryComponent} from "./components/views/distributor/detailed-delivery/detailed-delivery.component";
import {
    ViewShippingRequestsComponent
} from "./components/views/distributor/search-requests/view-shipping-requests/view-shipping-requests.component";
import {LoginComponent} from "./components/views/auth/login/login.component";
import {RegisterComponent} from "./components/views/auth/register/register.component";

const routes: Routes = [
    {path: '', redirectTo: 'dashboard', pathMatch: 'full'},
    {path: 'dashboard', component: DashboardComponent},
    {
        path: 'auth', children: [
            {
                path: '',
                redirectTo: 'login',
                pathMatch: 'full'
            },
            {
                path: 'login',
                component: LoginComponent,
            },
            {
                path: 'register',
                component: RegisterComponent,
            }
        ]
    },
    {
        path: 'sender', component: SenderComponent, children: [
            {
                path: '',
                component: ViewShippingsComponent,
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
    {
        path: 'distributor', component: DistributorComponent, children: [
            {
                path: '',
                component: ViewDeliveriesComponent,
            },
            {
                path: 'deliveries/:id',
                component: DetailedDeliveryComponent,
            },
            {
                path: 'requests',
                component: SearchRequestsComponent, children: [
                    {
                        path: '',
                        component: ViewShippingRequestsComponent,
                    },
                    {
                        path: ':id/offers/new',
                        component: NewShippingOfferComponent,
                    }]
            },
        ]
    }
];

@NgModule({
    imports: [RouterModule.forRoot(routes, {onSameUrlNavigation: 'reload'})],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
