import { Component, OnInit } from '@angular/core';
import {Shipping} from "../../../models/shipping";
import {SenderService} from "../../../services/sender.service";

@Component({
  selector: 'app-sender',
  templateUrl: './sender.component.html'
})
export class SenderComponent implements OnInit {

  shippingArr: Shipping[] = new Array<Shipping>();

  constructor(private service: SenderService) {
  }

  setShippings() {
    this.service.getAllShippings().subscribe((res) => {
      this.shippingArr = res.data;
    });
  }

  ngOnInit(): void {
    this.setShippings()
  }

}
