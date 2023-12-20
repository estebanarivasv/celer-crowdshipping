import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Shipping } from "../models/shipping";
import { Observable } from "rxjs";
import { ApiResponse } from "../models/response";
import { environment } from "../../environments/environment";
import { Offer } from "../models/offer";
import { State } from "../models/state";
import { AuthService } from "./auth.service";

/*
 * HttpClient makes requests asynchronously and  returns an Observable that emits the
 * requested data when the response is received.
 * */
@Injectable({
  providedIn: "root",
})
export class SenderService {
  constructor(
    private http: HttpClient,
    private authService: AuthService,
  ) {}

  getAllShippings(): Observable<ApiResponse<Shipping[]>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings`;
    const token = this.authService.getToken();
    return this.http.get<ApiResponse<Shipping[]>>(URL, {
      headers: { Authorization: `${token}` },
    });
  }

  addShipping(body: Shipping): Observable<ApiResponse<Shipping>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings`;
    const token = this.authService.getToken();
    return this.http.post<ApiResponse<Shipping>>(URL, body, {
      headers: { Authorization: `${token}` },
    });
  }

  getShippingDetails(id: number): Observable<ApiResponse<Shipping>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings/${id}`;
    const token = this.authService.getToken();
    return this.http.get<ApiResponse<Shipping>>(URL, {
      headers: { Authorization: `${token}` },
    });
  }

  getShippingOffers(id: number): Observable<ApiResponse<Offer[]>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers`;
    const token = this.authService.getToken();
    return this.http.get<ApiResponse<Offer[]>>(URL, {
      headers: { Authorization: `${token}` },
    });
  }

  getSelectedOfferByShippingId(id: number): Observable<ApiResponse<Offer>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers/selected`;
    const token = this.authService.getToken();
    return this.http.get<ApiResponse<Offer>>(URL, {
      headers: { Authorization: `${token}` },
    });
  }

  getShippingStateById(id: number): Observable<ApiResponse<State>> {
    const URL = `${environment.apiBaseUrl}/shippings/${id}/state`;
    const token = this.authService.getToken();
    return this.http.get<ApiResponse<State>>(URL, {
      headers: { Authorization: `${token}` },
    });
  }

  setSelectedOfferById(
    id: number,
    offerId: number,
  ): Observable<ApiResponse<Shipping>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings/${id}/offers/selected`;
    const token = this.authService.getToken();
    return this.http.patch<ApiResponse<Shipping>>(
      URL,
      {
        selectedOfferId: offerId,
      },
      { headers: { Authorization: `${token}` } },
    );
  }

  sendMsgToWorflowEngine(id: number): Observable<ApiResponse<State>> {
    const URL = `${environment.apiBaseUrl}/sender/shippings/${id}`;
    const token = this.authService.getToken();

    let body = {
      messageName: "PackageOk",
    };
    return this.http.patch<ApiResponse<State>>(URL, body, {
      headers: { Authorization: `${token}` },
    });
  }
}
