import {Package} from "./package";
import {User} from "./user";

export class Shipping {

    id?: number;
    details?: string;
    pickedUpAt?: Date;
    deliveredAt?: Date;
    originAddress?: string;
    destinationAddress?: string;
    packageId?: number;
    selectedOfferId?: number;
    senderId?: number;
    recipientId?: number;
    processId?: string;
    package: Package = new Package();
    sender?: User;
    recipient?: User;

    constructor() {
    }
}


/*

type Shipping struct {
    gorm.Model
    ID                 int `gorm:"primarykey"`
    Details            string
    PackageID          int
    SelectedOfferID    *int `sql:"default:null"` // Selected Offer
    SenderID           int
    RecipientID        int
    ProcessID          string
    PickedUpAt         time.Time
    DeliveredAt        time.Time
    OriginAddress      string
    DestinationAddress string
    Offers             []Offer       `gorm:"preload:true"` // Has many
    Route              []RouteDetail `gorm:"preload:true"` // Has many
    SelectedOffer      Offer         `gorm:"preload:true"` // Has one
    Sender             User          `gorm:"preload:true"` // Has one
    Recipient          User          `gorm:"preload:true"` // Has one
    Package            Package       `gorm:"preload:true"` // Has one
}
*/