import { Component, OnInit } from "@angular/core";
import { Package } from "../../../../models/package";
import { MessageService } from "primeng/api";
import { PackageService } from "../../../../services/package.service";
import { ApiResponse } from "../../../../models/response";
import { Shipping } from "../../../../models/shipping";
import { SenderService } from "../../../../services/sender.service";
import { Router } from "@angular/router";

@Component({
  selector: "app-new-request",
  templateUrl: "./new-request.component.html",
  providers: [MessageService],
})
export class NewRequestComponent implements OnInit {
  package = new Package();
  shipping = new Shipping();

  constructor(
    private messageService: MessageService,
    private packageService: PackageService,
    private senderService: SenderService,
    private router: Router,
  ) {}

  dimensionsOpts: string[] = [
    "Bag - <2 Kg",
    "XS Box - 2 Kg",
    "S Box - 3 Kg",
    "M Box - 5 Kg",
    "L Box - 10 Kg",
    "XL Box - 20 Kg",
  ];



  ngOnInit(): void {}

  onSubmit(packageData: Package, shippingData: Shipping) {
    console.log({ package: packageData, shipping: shippingData });
    this.packageService.addPackage(packageData).subscribe({
      next: (res: ApiResponse<Package>) => {
        const packageFromApi = res.data;
        shippingData.packageId = packageFromApi.id

        this.senderService.addShipping(shippingData).subscribe({
          next: (_res: ApiResponse<Shipping>) => {
            this.messageService.add({
              severity: 'success',
              summary: 'Created',
              detail: 'Shipping successfully created'
            });
          },
          error: (error) => {
            this.messageService.add({
              severity: "danger",
              summary: error.summary,
              detail: error.detail,
            });
          },
        })
      },
      error: (error) => {
        this.messageService.add({
          severity: "danger",
          summary: error.summary,
          detail: error.detail,
        });
      },
    });

    /*
    * this.authService.login(username, password).subscribe({
      next: (response: any) => {
        const jwtToken = response.data.token;

        this.authService.setToken(jwtToken);
        this.router.navigate(["/dashboard"]);
      },
      error: (error) => {
        this.toastr.error(error.error.message, "", {
          closeButton: true,
        });
      },
    });
    * */
  }

  /*
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

         */
}
