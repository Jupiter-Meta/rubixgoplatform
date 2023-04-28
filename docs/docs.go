// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "murali.c@ensurity.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/commit-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "batchID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/create-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User/Entity ID",
                        "name": "UserID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User/Entity Info",
                        "name": "UserInfo",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Committer DID",
                        "name": "CommitterDID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "BacthID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "File Info is json string {",
                        "name": "FileInfo",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "File to be committed",
                        "name": "FileContent",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-account-info": {
            "get": {
                "description": "For a mentioned DID, check the account balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Check account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-comment": {
            "get": {
                "description": "Retrieves the details of a transaction based on its comment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation Comment",
                "operationId": "get-by-comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment to identify the transaction",
                        "name": "Comment",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-did": {
            "get": {
                "description": "Retrieves the details of a transaction based on dID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by dID",
                "operationId": "get-by-did",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID of sender/receiver",
                        "name": "DID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter by role as sender or receiver",
                        "name": "Role",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-data-token": {
            "get": {
                "description": "This API will get all data token belong to the did",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Get Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-txn-details-by-id": {
            "get": {
                "description": "Retrieves the details of a transaction based on its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation ID",
                "operationId": "get-txn-details-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the transaction to retrieve",
                        "name": "txnID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/initiate-rbt-transfer": {
            "post": {
                "description": "This API will initiate RBT transfer to the specified dID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Initiate RBT Transfer",
                "operationId": "initiate-rbt-transfer",
                "parameters": [
                    {
                        "description": "The decentralized identifier of the receiver",
                        "name": "receiver",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The decentralized identifier of the sender",
                        "name": "sender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The number of RBT tokens to transfer",
                        "name": "tokenCount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "A comment for the transfer",
                        "name": "comment",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The type of transfer (1 for direct transfer, 2 for group transfer)",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/start": {
            "get": {
                "description": "It will setup the core if not done before",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Start Core",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "SessionToken": {
            "type": "apiKey",
            "name": "Session-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.9",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Rubix Core",
	Description:      "Rubix core API to control & manage the node.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
