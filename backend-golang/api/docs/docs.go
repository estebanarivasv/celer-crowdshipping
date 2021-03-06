// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/distributor/offers": {
            "get": {
                "description": "Get offers associated to a distributor from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get offers made by a distributor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Distributor ID",
                        "name": "distributorId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/offers/{id}": {
            "delete": {
                "description": "Delete offer from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete offer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/requests": {
            "get": {
                "description": "Get all requests which are stored in the database and expect offers",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all pending shipping requests",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/requests/{id}/offers": {
            "get": {
                "description": "Get offers associated to a request from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get request offers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping request ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new shipping request offer instance and stores it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates a new offer for a shipping request",
                "parameters": [
                    {
                        "description": "Fill the body to create a new shipping request offer",
                        "name": "Offer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.NewOfferInDTO"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Shipping request ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/shippings": {
            "get": {
                "description": "Get shippings that started their track into the recipient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get shippings in progress",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/shippings/{id}": {
            "get": {
                "description": "Get Shipping stored in the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Shipping",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Change shipping state by sending a message to a camunda process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Shipping State",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Fill the body to change shipping state",
                        "name": "Shipping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.MessageToProcessInDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/distributor/shippings/{id}/route": {
            "post": {
                "description": "Add new coordinates to a shipping and stores it in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new coordinates",
                "parameters": [
                    {
                        "description": "Fill the body to update coordinates",
                        "name": "Offer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.NewRouteDetailInDTO"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/packages": {
            "post": {
                "description": "Creates a new package in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates a new package",
                "parameters": [
                    {
                        "description": "Fill the body to create a new package",
                        "name": "Package",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.PackageInDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/packages/{id}": {
            "delete": {
                "description": "Delete a Package by ID which is stored in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Package",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Package ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/sender/shippings": {
            "get": {
                "description": "Get all shippings stored in the database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all shippings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new shipping instance in camunda and in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates a new shipping",
                "parameters": [
                    {
                        "description": "Fill the body to create a new shipping",
                        "name": "Shipping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.ShippingInDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/sender/shippings/{id}": {
            "get": {
                "description": "Get Shipping stored in the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Shipping",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Shipping by ID which is stored in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Shipping",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Change shipping state by sending a message to a camunda process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Shipping State",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Fill the body to change shipping state",
                        "name": "Shipping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.MessageToProcessInDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/sender/shippings/{id}/offers": {
            "get": {
                "description": "Get offers associated to a shipping from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get shipping offers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/sender/shippings/{id}/offers/selected": {
            "get": {
                "description": "Get a shipping selected route from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Shipping selected offer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Add selected offer and send a message to a camunda process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Shipping's selected offer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add selected offer ID",
                        "name": "Shipping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.ShippingInPatchDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/sender/shippings/{id}/route": {
            "get": {
                "description": "Get a shipping route from the database by passing an ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Shipping route",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        },
        "/shippings/{id}/state": {
            "get": {
                "description": "Get Shipping state making reference to Camunda instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Shipping state",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Shipping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.MessageToProcessInDTO": {
            "type": "object",
            "properties": {
                "messageName": {
                    "type": "string"
                }
            }
        },
        "dtos.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "entities.NewOfferInDTO": {
            "type": "object",
            "properties": {
                "distributorId": {
                    "type": "integer"
                },
                "duration": {
                    "type": "number"
                },
                "message": {
                    "type": "string"
                },
                "shippingCost": {
                    "type": "number"
                }
            }
        },
        "entities.NewRouteDetailInDTO": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "string"
                }
            }
        },
        "entities.PackageInDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dimensions": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entities.ShippingInDTO": {
            "type": "object",
            "properties": {
                "destinationAddress": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "originAddress": {
                    "type": "string"
                },
                "packageId": {
                    "type": "integer"
                },
                "recipientId": {
                    "type": "integer"
                },
                "senderId": {
                    "type": "integer"
                }
            }
        },
        "entities.ShippingInPatchDTO": {
            "type": "object",
            "properties": {
                "selectedOfferId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
