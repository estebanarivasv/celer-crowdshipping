export class Shipping {

    id: number;
    details: string;
    pickedUpAt: Date;
    deliveredAt: Date;
    originAddress: string;
    deliveryAddress: string;
    packageId?: number;
    selectedOfferId?: number;
    senderId?: number;
    recipientId?: number;
    processId?: string;
    package?: Package;
    sender?: User;
    recipient?: User;


    constructor(id: number,
                details: string,
                pickedUpAt: Date,
                deliveredAt: Date,
                originAddress: string,
                deliveryAddress: string,
                packageId?: number,
                selectedOfferId?: number,
                senderId?: number,
                recipientId?: number,
                processId?: string,
                packageDetails?: Package,
                sender?: User,
                recipient?: User) {
        this.id = id;
        this.details = details;
        this.pickedUpAt = pickedUpAt;
        this.deliveredAt = deliveredAt;
        this.originAddress = originAddress;
        this.deliveryAddress = deliveryAddress;
        this.packageId = packageId;
        this.selectedOfferId = selectedOfferId;
        this.senderId = senderId;
        this.recipientId = recipientId;
        this.processId = processId;
        this.package = packageDetails;
        this.sender = sender;
        this.recipient = recipient;
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