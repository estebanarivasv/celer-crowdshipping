import { Component, OnInit } from "@angular/core";
import { AuthService } from "../../../../services/auth.service";
import { ToastrService } from "ngx-toastr";

@Component({
  selector: "app-register",
  templateUrl: "./register.component.html",
  styleUrls: ["./register.component.scss"],
})
export class RegisterComponent implements OnInit {
  firstName: string = "";
  lastName: string = "";
  phone: string = "";
  username: string = "";
  email: string = "";
  password: string = "";

  constructor(
    private authService: AuthService,
    private toastr: ToastrService,
  ) {}

  ngOnInit(): void {}

  onSubmit(
    firstName: string,
    lastName: string,
    phone: string,
    username: string,
    email: string,
    password: string,
  ) {
    this.authService
      .register(firstName, lastName, phone, username, email, password)
      .subscribe({
        next: (response: any) => {
          this.toastr.success("Account created", "Success", {
            closeButton: true,
          });
        },
        error: (error) => {
          this.toastr.error(error.error.message, "", {
            closeButton: true,
          });
        },
      });
  }
}
