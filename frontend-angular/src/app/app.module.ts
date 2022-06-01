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

@NgModule({
    declarations: [
        AppComponent,
        SenderComponent,
        DistributorComponent,
        HeaderComponent,
        DashboardComponent
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        HttpClientModule,
        AppRoutingModule,
        SharedModule,
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
