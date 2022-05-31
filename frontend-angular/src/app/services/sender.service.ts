import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Shipping} from "../models/shipping";
import {map, Observable} from "rxjs";
import {ApiResponse} from "../models/response";


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
        // TODO: how to env variables
        const URL = 'http://localhost:5000/api/v1/sender/shippings'

        // `${environment.baseURL}student.json`
        return this.http.get<ApiResponse<Shipping[]>>(URL);
    }
}

