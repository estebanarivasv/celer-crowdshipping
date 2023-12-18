import { Component, OnInit } from "@angular/core";
import { AuthService } from "../../../../services/auth.service";
import { ToastrService } from "ngx-toastr";
import { Router } from "@angular/router";

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
  styleUrls: ["./login.component.scss"],
})
export class LoginComponent implements OnInit {
  username: string = "";
  password: string = "";

  constructor(
    private authService: AuthService,
    private toastr: ToastrService,
    private router: Router,
  ) {}

  ngOnInit(): void {}

  onSubmit(username: string, password: string) {
    this.authService.login(username, password).subscribe({
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
  }
}
