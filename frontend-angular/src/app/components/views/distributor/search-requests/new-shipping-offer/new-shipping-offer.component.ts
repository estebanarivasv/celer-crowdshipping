import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap} from "@angular/router";
import {MessageService} from "primeng/api";
import {Package} from "../../../../../models/package";
import {Offer} from "../../../../../models/offer";
import {ApiResponse} from "../../../../../models/response";
import {DistributorService} from "../../../../../services/distributor.service";
import {Shipping} from "../../../../../models/shipping";

@Component({
    selector: 'app-new-shipping-offer',
    templateUrl: './new-shipping-offer.component.html',
    providers: [MessageService]
})
export class NewShippingOfferComponent implements OnInit {

    offer = new Offer();
    requestId?: number;

    constructor(private route: ActivatedRoute,
                private messageService: MessageService,
                private distributorService: DistributorService) {
    }

    loadParamFromUrl() {
        this.route.paramMap.subscribe((params: ParamMap) => this.requestId = parseInt(<string>params.get("id")));
    }

    ngOnInit(): void {
        this.loadParamFromUrl();
    }

    onSubmit(offer: Offer ) {
        this.distributorService.makeOfferToRequest(this.requestId!, offer).subscribe({
            next: (res: ApiResponse<Package>) => {
                this.messageService.add({
                    severity: 'success',
                    summary: 'Created',
                    detail: 'Offer successfully made'
                });
            },
            error: (error) => {
                this.messageService.add({
                    severity: 'danger',
                    summary: error.summary,
                    detail: error.detail
                });
            },
        });
    }
}
