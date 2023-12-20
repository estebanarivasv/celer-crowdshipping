import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {ApiResponse} from "../models/response";
import {Package} from "../models/package";
import {environment} from "../../environments/environment";
import {HttpClient} from "@angular/common/http";
import {AuthService} from "./auth.service";

@Injectable({
  providedIn: 'root'
})
export class PackageService {

  constructor(private http: HttpClient,
              private authService: AuthService) { }

  addPackage(body: Package): Observable<ApiResponse<Package>> {

    const URL = `${environment.apiBaseUrl}/packages`
    const token = this.authService.getToken();
    return this.http.post<ApiResponse<Package>>(URL, body, {headers: {'Authorization': `${token}`}});
  }

}
