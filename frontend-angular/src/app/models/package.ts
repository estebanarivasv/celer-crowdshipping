export class Package {

    id: number;
    name: string;
    imageUrl: string;
    description: string;
    dimensions: string;
    value: number;

    constructor(id: number,
                name: string,
                imageUrl: string,
                description: string,
                dimensions: string,
                value: number) {
        this.id = id;
        this.name = name;
        this.imageUrl = imageUrl;
        this.description = description;
        this.dimensions = dimensions;
        this.value = value;
    }
}