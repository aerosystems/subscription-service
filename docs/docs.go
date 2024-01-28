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
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
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
                        "ApiKeyAuth": []
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
                            "$ref": "#/definitions/handlers.InvoiceRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.InvoiceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "object",
                                                "additionalProperties": {
                                                    "type": "integer"
                                                }
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Subscription"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.InvoiceRequest": {
            "type": "object",
            "required": [
                "durationSubscription",
                "kindSubscription"
            ],
            "properties": {
                "durationSubscription": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.DurationSubscription"
                        }
                    ],
                    "example": "12m"
                },
                "kindSubscription": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.KindSubscription"
                        }
                    ],
                    "example": "business"
                }
            }
        },
        "handlers.InvoiceResponse": {
            "type": "object",
            "properties": {
                "paymentUrl": {
                    "type": "string",
                    "example": "https://api.monobank.ua"
                }
            }
        },
        "handlers.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.DurationSubscription": {
            "type": "string",
            "enum": [
                "1m",
                "12m"
            ],
            "x-enum-varnames": [
                "OneMonthDurationSubscription",
                "TwelveMonthDurationSubscription"
            ]
        },
        "models.KindSubscription": {
            "type": "string",
            "enum": [
                "trial",
                "startup",
                "business"
            ],
            "x-enum-varnames": [
                "TrialSubscription",
                "StartupSubscription",
                "BusinessSubscription"
            ]
        },
        "models.Subscription": {
            "type": "object",
            "properties": {
                "accessTime": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "duration": {
                    "$ref": "#/definitions/models.DurationSubscription"
                },
                "id": {
                    "type": "integer"
                },
                "kind": {
                    "$ref": "#/definitions/models.KindSubscription"
                },
                "updatedAt": {
                    "type": "string"
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
