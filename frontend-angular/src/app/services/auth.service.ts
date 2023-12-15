import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { environment } from "../../environments/environment";
import { BehaviorSubject, Observable } from "rxjs";
import { ApiResponse } from "../models/response";
import { Shipping } from "../models/shipping";
import { Login } from "../models/auth";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  constructor(private http: HttpClient) {
    const storedToken = localStorage.getItem("token"); // Initialize token value from local storage
    this.tokenSubject.next(storedToken);
  }

  private tokenSubject = new BehaviorSubject<string | null>(null);

  getToken(): string | null {
    return localStorage.getItem("token");
  }

  setToken(token: string | null) {
    // Store token in local storage
    localStorage.setItem("token", token || ""); // Store an empty string if token is null
    // Updates the BehaviorSubject with the new token
    this.tokenSubject.next(token);
  }

  login(email: string, password: string): Observable<ApiResponse<Login>> {
    const endpoint = `${environment.apiBaseUrl}/auth/login`;
    const requestBody = {
      username: email,
      password: password,
    };

    return this.http.post<ApiResponse<Login>>(endpoint, requestBody);
  }

  register(
    firstName: string,
    lastName: string,
    phone: string,
    username: string,
    email: string,
    password: string,
  ): Observable<ApiResponse<Login>> {
    const endpoint = `${environment.apiBaseUrl}/auth/register`;
    const requestBody = {
      firstName: firstName,
      lastName: lastName,
      phone: phone,
      username: username,
      email: email,
      password: password,
    };

    return this.http.post<ApiResponse<Login>>(endpoint, requestBody);
  }
}
