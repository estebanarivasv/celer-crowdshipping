import {Component} from '@angular/core';
import {MenuItem} from 'primeng/api';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html'
})
export class HeaderComponent {

    items: MenuItem[] = [
        {
            label: 'My shippings',
            icon: 'pi pi-fw pi-file',
            routerLink: "/sender"
        },
        {
            label: 'My deliveries',
            icon: 'pi pi-fw pi-pencil',
            routerLink: "/distributor"
        }
    ];

}
