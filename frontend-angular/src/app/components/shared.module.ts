import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TableModule} from "primeng/table";
import {MenubarModule} from "primeng/menubar";
import {InputTextModule} from "primeng/inputtext";
import {ButtonModule} from "primeng/button";
import {AvatarModule} from "primeng/avatar";
import {SidebarModule} from "primeng/sidebar";
import {ToggleButtonModule} from "primeng/togglebutton";
import {CardModule} from "primeng/card";
import {RippleModule} from "primeng/ripple";
import {ImageModule} from "primeng/image";

@NgModule({
    imports: [
        CommonModule
    ],
    exports: [
        TableModule,
        MenubarModule,
        InputTextModule,
        ButtonModule,
        AvatarModule,
        SidebarModule,
        ToggleButtonModule,
        CardModule,
        RippleModule,
        ImageModule,
    ], providers: []
})
export class SharedModule {
}
