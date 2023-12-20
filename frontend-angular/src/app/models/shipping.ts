import {Package} from "./package";
import {User} from "./user";

export class Shipping {

    id?: number;
    details?: string;
    pickedUpAt?: Date;
    deliveredAt?: Date;
    createdAt?: Date;
    originAddress?: string;
    destinationAddress?: string;
    packageId?: number;
    selectedOfferId?: number;
    senderId?: number;
    recipientId?: number;
    processId?: string;
    package: Package = new Package();
    sender?: User;
    recipient: User = new User();

    constructor() {
    }
}

