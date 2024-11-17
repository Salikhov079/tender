// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/bid/contractor/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "List all bids placed by a specific contractor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bid"
                ],
                "summary": "List Contractor Bids",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Contractor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Contractor's bids retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllBidsByUserIdRequest"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error while retrieving bids",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bid/list": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "List all bids with optional filtering by price, delivery time, limit, and offset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bid"
                ],
                "summary": "List Bids",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Price filter",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Delivery time filter",
                        "name": "delivery_time",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit the number of bids",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bids retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ListBidsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error while retrieving bids",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bid/submit": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Submit a bid for a tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bid"
                ],
                "summary": "Submit Bid",
                "parameters": [
                    {
                        "description": "Bid details",
                        "name": "bid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.SubmitBidRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bid submitted successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.BidResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error while submitting bid",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bid/tender/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all bids associated with a specific tender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bid"
                ],
                "summary": "Get All Bids by Tender ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tender ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bids retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ListBidsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error while retrieving bids",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenders": {
            "get": {
                "description": "List tenders with optional filtering by title and deadline",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "List tenders",
                "parameters": [
                    {
                        "description": "Filter and pagination parameters",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.ListTendersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of tenders",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ListTendersResponse"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing tender with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "Update a tender",
                "parameters": [
                    {
                        "description": "Updated tender details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.UpdateTenderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender updated successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.TenderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Tender not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new tender with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "Create a new tender",
                "parameters": [
                    {
                        "description": "Tender details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.CreateTenderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender created successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.TenderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenders/award": {
            "post": {
                "description": "Award a tender to a specific bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "Award a tender",
                "parameters": [
                    {
                        "description": "Tender award details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.CreatTenderAwardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender awarded successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.TenderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Tender or bid not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenders/user/{id}": {
            "get": {
                "description": "List all tenders created by a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "List tenders for a specific user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of user tenders",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ListTendersResponse"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tenders/{id}": {
            "delete": {
                "description": "Delete a tender by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "04-Tender"
                ],
                "summary": "Delete a tender",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tender ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tender deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/genprotos.TenderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Tender not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.BidResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "genprotos.CreatTenderAwardRequest": {
            "type": "object",
            "properties": {
                "bid_id": {
                    "type": "string"
                },
                "tender_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.CreateTenderRequest": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "number"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetAllBidResponse": {
            "type": "object",
            "properties": {
                "Tenders": {
                    "$ref": "#/definitions/genprotos.GetTenderResponse"
                },
                "comments": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "delivery_time": {
                    "description": "in days",
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "status": {
                    "description": "\"pending\", \"awarded\"",
                    "type": "string"
                },
                "tender_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetAllBidsByUser": {
            "type": "object",
            "properties": {
                "Tenders": {
                    "$ref": "#/definitions/genprotos.GetTenderResponse"
                },
                "comments": {
                    "type": "string"
                },
                "contractor_id": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "genprotos.GetAllBidsByUserIdRequest": {
            "type": "object",
            "properties": {
                "Binds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetAllBidsByUser"
                    }
                }
            }
        },
        "genprotos.GetTenderResponse": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "number"
                },
                "client_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "genprotos.ListBidsResponse": {
            "type": "object",
            "properties": {
                "bids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetAllBidResponse"
                    }
                }
            }
        },
        "genprotos.ListTendersRequest": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "genprotos.ListTendersResponse": {
            "type": "object",
            "properties": {
                "tenders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.GetTenderResponse"
                    }
                }
            }
        },
        "genprotos.SubmitBidRequest": {
            "type": "object",
            "properties": {
                "comments": {
                    "description": "Optional",
                    "type": "string"
                },
                "contractor_id": {
                    "type": "string"
                },
                "delivery_time": {
                    "description": "in days",
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "tender_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.TenderResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "genprotos.UpdateTenderRequest": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "number"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Gateway",
	Description:      "Dilshod's API Gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
