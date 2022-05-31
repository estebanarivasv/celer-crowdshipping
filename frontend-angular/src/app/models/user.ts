export class User {
    id: number;
    name: string;
    surname: string;
    phone: number;
    email: string;
    bankName: string;
    bankCbu: string;


    constructor(id: number,
                name: string,
                surname: string,
                phone: number,
                email: string,
                bankName: string,
                bankCbu: string) {
        this.id = id;
        this.name = name;
        this.surname = surname;
        this.phone = phone;
        this.email = email;
        this.bankName = bankName;
        this.bankCbu = bankCbu;
    }
}