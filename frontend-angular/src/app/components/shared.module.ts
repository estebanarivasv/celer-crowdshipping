import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FooterModule, SidebarModule, TableModule} from '@coreui/angular';

@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        FooterModule,
        SidebarModule,
        TableModule
    ],
    exports: [
        FooterModule,
        SidebarModule,
        TableModule
    ]
})
export class SharedModule {
}
