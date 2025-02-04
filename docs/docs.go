// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Artem Kostenko",
            "url": "https://github.com/aerosystems"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/invoices/{payment_method}": {
            "post": {
                "security": [
                    {
                        "ServiceApiKeyAuth": []
                    }
                ],
                "description": "Create invoice by payment method",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "invoices"
                ],
                "summary": "Create invoice",
                "parameters": [
                    {
                        "enum": [
                            "monobank"
                        ],
                        "type": "string",
                        "description": "payment method",
                        "name": "payment_method",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "invoice",
                        "name": "invoice",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.InvoiceRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payment.InvoiceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/prices": {
            "get": {
                "description": "get prices for all available subscriptions, in cents",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prices"
                ],
                "summary": "Get prices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "object",
                                "additionalProperties": {
                                    "type": "integer"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions": {
            "get": {
                "security": [
                    {
                        "ServiceApiKeyAuth": []
                    }
                ],
                "description": "get subscriptions by userUuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Get subscriptions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/subscription.GetSubscriptionResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Create subscription",
                "parameters": [
                    {
                        "description": "Create subscription",
                        "name": "raw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/subscription.CreateSubscriptionRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/subscription.CreateSubscriptionResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions/create-free-trial": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create free trial",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Create free trial",
                "parameters": [
                    {
                        "description": "Create free trial",
                        "name": "raw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/subscription.CreateFreeTrialRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/subscription.CreateSubscriptionResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions/{subscriptionId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Update subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription ID",
                        "name": "subscriptionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Delete subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription ID",
                        "name": "subscriptionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.SubscriptionType": {
            "type": "object"
        },
        "payment.InvoiceRequest": {
            "type": "object"
        },
        "payment.InvoiceResponse": {
            "type": "object",
            "properties": {
                "paymentUrl": {
                    "type": "string",
                    "example": "https://api.monobank.ua"
                }
            }
        },
        "subscription.CreateFreeTrialRequestBody": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "subscription": {
                    "type": "string"
                }
            }
        },
        "subscription.CreateSubscriptionRequestBody": {
            "type": "object",
            "properties": {
                "customerUuid": {
                    "type": "string",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "subscriptionDuration": {
                    "type": "string",
                    "example": ""
                },
                "subscriptionType": {
                    "type": "string",
                    "example": "business"
                }
            }
        },
        "subscription.CreateSubscriptionResponseBody": {
            "type": "object",
            "properties": {
                "accessTime": {
                    "type": "string"
                },
                "customerUuid": {
                    "type": "string"
                },
                "subscriptionDuration": {
                    "type": "string"
                },
                "subscriptionType": {
                    "type": "string"
                }
            }
        },
        "subscription.GetSubscriptionResponse": {
            "type": "object",
            "properties": {
                "accessTime": {
                    "type": "string",
                    "example": "2021-09-01T00:00:00Z"
                },
                "duration": {
                    "type": "string",
                    "example": "12m"
                },
                "name": {
                    "type": "string",
                    "example": "business"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Should contain Access JWT Token, with the Bearer started",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "gw.verifire.dev/subs",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "Subscription Service",
	Description:      "A part of microservice infrastructure, who responsible for user subscriptions",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
