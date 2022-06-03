import {Component, OnInit} from '@angular/core';
import {DistributorService} from "../../../../services/distributor.service";
import {Shipping} from "../../../../models/shipping";

@Component({
    selector: 'app-search-requests',
    templateUrl: './search-requests.component.html'
})
export class SearchRequestsComponent implements OnInit {

    constructor() {
    }

    ngOnInit(): void {
    }

}
