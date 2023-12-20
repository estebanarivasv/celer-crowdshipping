import {Injectable, OnDestroy} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable, Subscription} from "rxjs";
import {ApiResponse} from "../models/response";
import {Shipping} from "../models/shipping";
import {environment} from "../../environments/environment";
import {Offer} from "../models/offer";
import {State} from "../models/state";
import {AuthService} from "./auth.service";

@Injectable({
    providedIn: 'root'
})
export class DistributorService {


    constructor(private http: HttpClient,
                private authService: AuthService) {
    }

    getAllDeliveries(): Observable<ApiResponse<Shipping[]>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings`
        const token = this.authService.getToken();

        return this.http.get<ApiResponse<Shipping[]>>(URL, {headers: {'Authorization': `${token}`}});
    }

    searchPendingRequests(): Observable<ApiResponse<Shipping[]>> {
        const URL = `${environment.apiBaseUrl}/distributor/requests`
        const token = this.authService.getToken();
        return this.http.get<ApiResponse<Shipping[]>>(URL, {headers: {'Authorization': `${token}`}});
    }

    makeOfferToRequest(id: number, body: Offer): Observable<ApiResponse<Offer>> {
        const URL = `${environment.apiBaseUrl}/distributor/requests/${id}/offers`
        const token = this.authService.getToken();
        return this.http.post<ApiResponse<Offer>>(URL, body, {headers: {'Authorization': `${token}`}});
    }

    getDeliveryDetails(id: number): Observable<ApiResponse<Shipping>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings/${id}`
        const token = this.authService.getToken();
        return this.http.get<ApiResponse<Shipping>>(URL, {headers: {'Authorization': `${token}`}});
    }

    getSelectedOfferByShippingId(id: number): Observable<ApiResponse<Offer>> {
        const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers/selected`
        const token = this.authService.getToken();
        return this.http.get<ApiResponse<Offer>>(URL, {headers: {'Authorization': `${token}`}});
    }

    getShippingStateById(id: number): Observable<ApiResponse<State>> {
        const URL = `${environment.apiBaseUrl}/shippings/${id}/state`
        const token = this.authService.getToken();
        return this.http.get<ApiResponse<State>>(URL, {headers: {'Authorization': `${token}`}});
    }

    sendMsgToWorflowEngine(id: number, messageNum: number): Observable<ApiResponse<State>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings/${id}`

        const token = this.authService.getToken();
        switch (messageNum) {
            case 1: {
                let body = {
                    "messageName": "PackageInTransit"
                }
                return this.http.patch<ApiResponse<State>>(URL, body, {headers: {'Authorization': `${token}`}});
            }
            case 2: {
                let body = {
                    "messageName": "DeliveredToRecipient"
                }
                return this.http.patch<ApiResponse<State>>(URL, body, {headers: {'Authorization': `${token}`}});
            }
            default: {
                return this.getShippingStateById(id);
            }
        }

    }
}
