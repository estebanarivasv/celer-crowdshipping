import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {Offer} from "../../../../models/offer";
import {SenderService} from "../../../../services/sender.service";
import {Shipping} from "../../../../models/shipping";
import {State} from "../../../../models/state";
import {MessageService} from "primeng/api";

@Component({
    selector: 'app-detailed-request',
    templateUrl: './detailed-request.component.html',
    providers: [MessageService]
})
export class DetailedRequestComponent implements OnInit {

    requestId?: number;
    shipping?: Shipping;
    offers?: Offer[];
    selectedOffer?: Offer;
    shippingState?: State;
    shippingCompleted?: boolean;

    constructor(private route: ActivatedRoute,
                private messageService: MessageService,
                private senderService: SenderService) {
    }

    loadParamFromUrl() {
        this.route.paramMap.subscribe((params: ParamMap) => this.requestId = parseInt(<string>params.get("id")));
    }

    getShippingInformation() {
        if (this.requestId != null) {
            this.senderService.getShippingDetails(this.requestId).subscribe(
                (res) => {
                    this.shipping = res.data;
                    if (res.data.selectedOfferId !== null) {
                        this.senderService.getSelectedOfferByShippingId(res.data.id!).subscribe(
                            (res) => {
                                this.selectedOffer = res.data
                            }
                        )
                    }
                }
            )
            this.senderService.getShippingOffers(this.requestId).subscribe(
                (res) => this.offers = res.data
            )

            this.senderService.getShippingStateById(this.requestId).subscribe(
                (res) => {
                    if (res.data !== null) {
                        this.shippingState = res.data;
                    } else {
                        this.shippingCompleted = true;
                    }
                }
            )

        }
    }

    markOfferAsSelected(shippingId: number, offerId: number) {
        console.log(shippingId, offerId)
        this.senderService.setSelectedOfferById(shippingId, offerId).subscribe(res => {
            this.messageService.add({
                severity: 'success',
                summary: 'Selected',
                detail: 'Offer set as selected'
            });
        });

    }

    ngOnInit(): void {
        this.loadParamFromUrl();
        this.getShippingInformation();
    }

    onClickPkgOk(id: number) {
        this.senderService.sendMsgToWorflowEngine(id).subscribe(
            ((res) => {
                if (res.message === "" && res.success) {
                    this.messageService.add({
                        severity: 'success',
                        summary: 'Sent',
                        detail: 'Message sent to workflow engine and returned with no errors'
                    });
                }
            }), error => {
                this.messageService.add({
                    severity: 'error',
                    summary: 'Bad request',
                    detail: error.error.message!
                })
            });
    }

    protected readonly JSON = JSON;
}
