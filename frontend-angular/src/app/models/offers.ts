import {Shipping} from "./shipping";
import {User} from "./user";

export class Offers {

    id?: number;
    shippingCost?: number;
    message?: string;
    duration?: number;
    shippingId?: number;
    shipping?: Shipping
    distributorId?: number;
    distributor?: User;

    constructor() {
    }
}