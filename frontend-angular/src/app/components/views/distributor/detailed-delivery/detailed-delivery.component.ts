import {Component, OnInit} from '@angular/core';
import {MessageService} from "primeng/api";
import {State} from "../../../../models/state";
import {Shipping} from "../../../../models/shipping";
import {Offer} from "../../../../models/offer";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {SenderService} from "../../../../services/sender.service";
import {DistributorService} from "../../../../services/distributor.service";

@Component({
    selector: 'app-detailed-delivery',
    templateUrl: './detailed-delivery.component.html',
    providers: [MessageService]
})
export class DetailedDeliveryComponent implements OnInit {

    reload: boolean = false;
    requestId?: number;
    delivery?: Shipping;
    selectedOffer?: Offer;
    deliveryState?: State;
    deliveryCompleted?: boolean;

    constructor(private route: ActivatedRoute,
                private messageService: MessageService,
                private distributorService: DistributorService,
                private router: Router) {
    }

    reloadComponent() {
        this.reload = !this.reload;
        this.loadParamFromUrl()
        this.getDeliveryInformation();
        this.reload = !this.reload;
    }

    ngOnInit(): void {
        this.loadParamFromUrl()
        this.getDeliveryInformation();
    }

    loadParamFromUrl() {
        this.route.paramMap.subscribe((params: ParamMap) => this.requestId = parseInt(<string>params.get("id")));
    }

    getDeliveryInformation() {
        if (this.requestId != null) {
            this.distributorService.getDeliveryDetails(this.requestId).subscribe(
                (res) => {
                    this.delivery = res.data;
                    if (res.data.selectedOfferId !== null) {
                        this.distributorService.getSelectedOfferByShippingId(res.data.id!).subscribe(
                            (res) => {
                                this.selectedOffer = res.data
                            }
                        )
                    }
                }
            )

            this.distributorService.getShippingStateById(this.requestId).subscribe(
                (res) => {
                    if (res.data !== null) {
                        this.deliveryState = res.data;
                    } else {
                        this.deliveryCompleted = true;
                    }
                }
            )

        }
    }

    onClickPkgInTransit(id: number) {
        this.distributorService.sendMsgToWorflowEngine(id, 1).subscribe(
            ((res) => {
                if (res.message === "" && res.success) {
                    this.messageService.add({
                        severity: 'success',
                        summary: 'Sent',
                        detail: 'Message sent to workflow engine and returned with no errors'
                    });

                    this.reloadComponent()
                }
            }), error => {
                this.messageService.add({
                    severity: 'error',
                    summary: 'Bad request',
                    detail: error.error.message!
                })
            });
    }

    onClickPkgDelivered(id: number) {
        this.distributorService.sendMsgToWorflowEngine(id, 2).subscribe(
            ((res) => {
                if (res.message === "" && res.success) {
                    this.messageService.add({
                        severity: 'success',
                        summary: 'Sent',
                        detail: 'Message sent to workflow engine and returned with no errors'
                    });

                    this.reloadComponent()
                }
            }), error => {
                this.messageService.add({
                    severity: 'error',
                    summary: 'Bad request',
                    detail: error.error.message!
                })
            });
    }
}
