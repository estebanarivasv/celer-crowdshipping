import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {ApiResponse} from "../models/response";
import {Shipping} from "../models/shipping";

@Injectable({
  providedIn: 'root'
})
export class DistributorService {

  constructor(private http: HttpClient) { }

  getAllRequests(): Observable<ApiResponse<Shipping[]>> {
    // TODO: how to env variables
    const URL = 'http://localhost:5000/api/v1/distributor/shippings'

    // `${environment.baseURL}student.json`

    return this.http.get<ApiResponse<Shipping[]>>(URL);
  }
}
