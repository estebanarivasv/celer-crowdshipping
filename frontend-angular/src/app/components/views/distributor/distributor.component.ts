import {Component, OnInit} from '@angular/core';
import {Shipping} from "../../../models/shipping";
import {SenderService} from "../../../services/sender.service";
import {DistributorService} from "../../../services/distributor.service";

@Component({
    selector: 'app-distributor',
    templateUrl: './distributor.component.html'
})
export class DistributorComponent implements OnInit {

    constructor() {
    }

    ngOnInit(): void {
    }

}
