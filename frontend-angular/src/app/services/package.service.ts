import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {ApiResponse} from "../models/response";
import {Package} from "../models/package";
import {environment} from "../../environments/environment";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class PackageService {

  constructor(private http: HttpClient) { }

  addPackage(body: Package): Observable<ApiResponse<Package>> {

    const URL = `${environment.apiBaseUrl}/packages`
    return this.http.post<ApiResponse<Package>>(URL, body);
  }

}
