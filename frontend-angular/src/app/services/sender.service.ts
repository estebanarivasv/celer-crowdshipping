import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Shipping} from "../models/shipping";
import {map, Observable} from "rxjs";
import {ApiResponse} from "../models/response";
import {Package} from "../models/package";
import {environment} from "../../environments/environment";
import {Offers} from "../models/offers";


/*
* HttpClient makes requests asynchronously and  returns an Observable that emits the
* requested data when the response is received.
* */
@Injectable({
    providedIn: 'root'
})
export class SenderService {

    constructor(private http: HttpClient) {
    }

    getAllShippings(): Observable<ApiResponse<Shipping[]>> {
        const URL = `${environment.apiBaseUrl}/sender/shippings`
        return this.http.get<ApiResponse<Shipping[]>>(URL);
    }

    addShipping(body: Shipping): Observable<ApiResponse<Shipping>> {
        const URL = `${environment.apiBaseUrl}/sender/shippings`
        return this.http.post<ApiResponse<Shipping>>(URL, body);
    }

    getShippingDetails(id: number): Observable<ApiResponse<Shipping>> {
        const URL = `${environment.apiBaseUrl}/sender/shippings/${id}`
        return this.http.get<ApiResponse<Shipping>>(URL);
    }

    getShippingOffers(id: number): Observable<ApiResponse<Offers[]>> {
        const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers`
        return this.http.get<ApiResponse<Offers[]>>(URL);
    }

}

