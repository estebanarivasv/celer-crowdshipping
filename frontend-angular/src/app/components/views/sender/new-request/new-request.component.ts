import {Component, OnInit} from '@angular/core';
import {Package} from "../../../../models/package";
import {MessageService} from "primeng/api";
import {PackageService} from "../../../../services/package.service";
import {ApiResponse} from "../../../../models/response";
import {Shipping} from "../../../../models/shipping";
import {SenderService} from "../../../../services/sender.service";
import {Router} from "@angular/router";

@Component({
    selector: 'app-new-request',
    templateUrl: './new-request.component.html',
    providers: [MessageService]
})
export class NewRequestComponent implements OnInit {

    constructor(
        private messageService: MessageService,
        private packageService: PackageService,
        private senderService: SenderService,
        private router: Router,
    ) {
    }

    packageModel = new Package();
    packageFromApi = new Package();
    savePkgClicked = false;
    saveShipClicked = false;
    dimensionsOpts: string[] = [
        "Bag - <2 Kg",
        "XS Box - 2 Kg",
        "S Box - 3 Kg",
        "M Box - 5 Kg",
        "L Box - 10 Kg",
        "XL Box - 20 Kg"
    ];

    shippingModel = new Shipping();

    ngOnInit(): void {
    }

    onClickAddPackage() {
        this.savePkgClicked = true;
        this.packageService.addPackage(this.packageModel)
            .subscribe(
            (res: ApiResponse<Package>) => {
                this.messageService.add({
                    severity: 'success',
                    summary: 'Created',
                    detail: 'Package successfully created'
                });
                this.packageFromApi = res.data;
            }, (error) => {
                this.messageService.add({
                    severity: 'danger',
                    summary: error.summary,
                    detail: error.detail
                });
            });
    }


    onClickAddShipping() {
        this.saveShipClicked = true;
        this.shippingModel.packageId = this.packageFromApi.id;
        this.senderService.addShipping(this.shippingModel)
            .subscribe(
                (res: ApiResponse<Shipping>) => {
                    this.messageService.add({
                        severity: 'success',
                        summary: 'Created',
                        detail: 'Shipping successfully created'
                    });
                    this.shippingModel = res.data;
                    this.router.navigate([`sender/requests/${this.shippingModel.id}`]);
                }, (error) => {
                    this.messageService.add({
                        severity: 'danger',
                        summary: error.summary,
                        detail: error.detail
                    });
                });
    }
}
