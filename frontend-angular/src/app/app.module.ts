import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {AppRoutingModule} from './app-routing.module';
import {HttpClientModule} from "@angular/common/http";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {SharedModule} from "./components/shared.module";

import {AppComponent} from './components/app.component';
import {SenderComponent} from './components/views/sender/sender.component';
import {HeaderComponent} from './components/header/header.component';
import {DistributorComponent} from "./components/views/distributor/distributor.component";
import {DashboardComponent} from "./components/views/dashboard/dashboard.component";
import {NewRequestComponent} from "./components/views/sender/new-request/new-request.component";
import {ViewShippingsComponent} from "./components/views/sender/view-shippings/view-shippings.component";
import {FormsModule} from "@angular/forms";
import {DetailedRequestComponent} from "./components/views/sender/detailed-request/detailed-request.component";

@NgModule({
    declarations: [
        AppComponent,
        SenderComponent,
        DistributorComponent,
        HeaderComponent,
        DashboardComponent,
        NewRequestComponent,
        ViewShippingsComponent,
        DetailedRequestComponent
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        HttpClientModule,
        AppRoutingModule,
        FormsModule,
        SharedModule,
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
