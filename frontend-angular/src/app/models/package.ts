

class Package {

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

/*
* 	ID          int `gorm:"primarykey"`
	Name        string
	ImageURL    string
	Description string
	Dimensions  string
	Value       float64 `json:"value"` // Monetary value
}
* */