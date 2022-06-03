import {Component, OnInit} from '@angular/core';
import {Shipping} from "../../../../../models/shipping";
import {DistributorService} from "../../../../../services/distributor.service";

@Component({
    selector: 'app-view-shipping-requests',
    templateUrl: './view-shipping-requests.component.html'
})
export class ViewShippingRequestsComponent implements OnInit {

    pendingRequests?: Shipping[]

    constructor(private distributorService: DistributorService) {
    }

    ngOnInit(): void {
        this.distributorService.searchPendingRequests().subscribe(
            (res) => {
                this.pendingRequests = res.data
            })
    }
}
