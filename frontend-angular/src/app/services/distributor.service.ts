import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {ApiResponse} from "../models/response";
import {Shipping} from "../models/shipping";
import {environment} from "../../environments/environment";
import {Offer} from "../models/offer";
import {State} from "../models/state";

@Injectable({
    providedIn: 'root'
})
export class DistributorService {

    constructor(private http: HttpClient) {
    }

    getAllDeliveries(): Observable<ApiResponse<Shipping[]>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings`
        return this.http.get<ApiResponse<Shipping[]>>(URL);
    }

    searchPendingRequests(): Observable<ApiResponse<Shipping[]>> {
        const URL = `${environment.apiBaseUrl}/distributor/requests`
        return this.http.get<ApiResponse<Shipping[]>>(URL);
    }

    makeOfferToRequest(id: number, body: Offer): Observable<ApiResponse<Offer>> {
        const URL = `${environment.apiBaseUrl}/distributor/requests/${id}/offers`
        return this.http.post<ApiResponse<Offer>>(URL, body);
    }

    getDeliveryDetails(id: number): Observable<ApiResponse<Shipping>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings/${id}`
        return this.http.get<ApiResponse<Shipping>>(URL);
    }

    getSelectedOfferByShippingId(id: number): Observable<ApiResponse<Offer>> {
        // TODO: Business model failing
        const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers/selected`
        return this.http.get<ApiResponse<Offer>>(URL);
    }

    getShippingStateById(id: number): Observable<ApiResponse<State>> {
        const URL = `${environment.apiBaseUrl}/shippings/${id}/state`
        return this.http.get<ApiResponse<State>>(URL);
    }

    sendMsgToWorflowEngine(id: number, messageNum: number): Observable<ApiResponse<State>> {
        const URL = `${environment.apiBaseUrl}/distributor/shippings/${id}`

        switch(messageNum) {
            case 1: {
                let body = {
                    "messageName": "PackageInTransit"
                }
                return this.http.patch<ApiResponse<State>>(URL, body);
            }
            case 2: {
                let body = {
                    "messageName": "DeliveredToRecipient"
                }
                return this.http.patch<ApiResponse<State>>(URL, body);
            }
            default: {
                return this.getShippingStateById(id);
            }
        }

    }
}
