import { Component, OnInit } from '@angular/core';
import {Shipping} from "../../../../models/shipping";
import {DistributorService} from "../../../../services/distributor.service";

@Component({
  selector: 'app-view-deliveries',
  templateUrl: './view-deliveries.component.html'
})
export class ViewDeliveriesComponent implements OnInit {

  requestsArr: Shipping[] = new Array<Shipping>();

  constructor(private service: DistributorService) { }

  setRequests() {
    this.service.getAllDeliveries().subscribe((res) => {
      this.requestsArr = res.data;
    });
  }

  ngOnInit(): void {
    this.setRequests()
  }

}
