// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This sample API demonstrates how to compose several authentication schemes \nand configure complex security requirements for your operations.\n\nThis API simulates a very simple market place with customers and resellers  \nof items.\n\nPersonas:\n  - as a first time user, I want to see all items on sales\n  - as a registered customer, I want to post orders for items and \n    consult my past orders\n  - as a registered reseller, I want to see all pending orders on the items \n    I am selling on the market place\n  - as a reseller managing my own inventories, I want to post replenishment orders for the items I provide\n  - as a register user, I want to consult my personal account infos\n\nThe situation we defined on the authentication side is as follows:\n  - every known user is authenticated using a basic token\n  - resellers are authenticated using API keys - we let the option to authenticate using a header or query param\n  - any registered user (customer or reseller) will add a signed JWT to access more API endpoints\n\nObviously, there are several ways to achieve the same result. We just wanted to demonstrate here how\nsecurity requirements may compose several schemes.\n\nNote that we used the \"OAuth2\" declaration here but don't implement a real  \nOAuth2 workflow: our intend here is just to be able to extract scopes from a passed JWT token (the \nonly way to manipulate scoped authorizers with Swagger 2.0 is to declare them with type \"oauth2\").\n",
    "title": "Composing authorizations",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/account": {
      "get": {
        "security": [
          {
            "isRegistered": []
          }
        ],
        "description": "Every registered user should be able to access this operation\n",
        "summary": "registered user account",
        "operationId": "GetAccount",
        "responses": {
          "200": {
            "description": "registered user personal account infos",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "401": {
            "$ref": "#/responses/unauthorized"
          },
          "default": {
            "$ref": "#/responses/otherError"
          }
        }
      }
    },
    "/items": {
      "get": {
        "security": [],
        "description": "Everybody should be able to access this operation\n",
        "summary": "items on sale",
        "operationId": "GetItems",
        "responses": {
          "200": {
            "$ref": "#/responses/multipleItems"
          },
          "default": {
            "$ref": "#/responses/otherError"
          }
        }
      }
    },
    "/order/add": {
      "post": {
        "security": [
          {
            "hasRole": [
              "customer"
            ],
            "isRegistered": []
          },
          {
            "hasRole": [
              "inventoryManager"
            ],
            "isReseller": []
          },
          {
            "hasRole": [
              "inventoryManager"
            ],
            "isResellerQuery": []
          }
        ],
        "description": "Registered customers should be able to add purchase orders.\nRegistered inventory managers should be able to add replenishment orders.\n",
        "summary": "post a new order",
        "operationId": "AddOrder",
        "parameters": [
          {
            "name": "order",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "empty response"
          },
          "401": {
            "$ref": "#/responses/unauthorized"
          },
          "403": {
            "$ref": "#/responses/forbidden"
          },
          "default": {
            "$ref": "#/responses/otherError"
          }
        }
      }
    },
    "/order/{orderID}": {
      "get": {
        "security": [
          {
            "hasRole": [
              "customer"
            ],
            "isRegistered": []
          }
        ],
        "description": "Only registered customers should be able to retrieve orders\n",
        "summary": "retrieves an order",
        "operationId": "GetOrder",
        "parameters": [
          {
            "type": "string",
            "name": "orderID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/singleOrder"
          },
          "401": {
            "$ref": "#/responses/unauthorized"
          },
          "403": {
            "$ref": "#/responses/forbidden"
          },
          "default": {
            "$ref": "#/responses/otherError"
          }
        }
      }
    },
    "/orders/{itemID}": {
      "get": {
        "security": [
          {
            "isReseller": []
          },
          {
            "isResellerQuery": []
          }
        ],
        "description": "Only registered resellers should be able to search orders for an item\n",
        "summary": "retrieves all orders for an item",
        "operationId": "GetOrdersForItem",
        "parameters": [
          {
            "type": "string",
            "name": "itemID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/multipleOrders"
          },
          "401": {
            "$ref": "#/responses/unauthorized"
          },
          "403": {
            "$ref": "#/responses/forbidden"
          },
          "default": {
            "$ref": "#/responses/otherError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Item": {
      "type": "string"
    },
    "Order": {
      "type": "object",
      "required": [
        "orderID"
      ],
      "properties": {
        "orderID": {
          "type": "string"
        },
        "orderLines": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "quantity",
              "purchasedItem"
            ],
            "properties": {
              "purchasedItem": {
                "$ref": "#/definitions/Item"
              },
              "quantity": {
                "type": "string",
                "format": "uint32",
                "minimum": 1
              }
            },
            "x-go-name": "orderLine"
          }
        }
      }
    },
    "principal": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "responses": {
    "forbidden": {
      "description": "forbidden access for a lack of sufficient privileges"
    },
    "multipleItems": {
      "description": "multiple items",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Item"
        }
      }
    },
    "multipleOrders": {
      "description": "multiple orders",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Order"
        }
      }
    },
    "otherError": {
      "description": "other error response",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "singleItem": {
      "description": "single item",
      "schema": {
        "type": "string"
      }
    },
    "singleOrder": {
      "description": "content of an order",
      "schema": {
        "$ref": "#/definitions/Order"
      }
    },
    "unauthorized": {
      "description": "unauthorized access for a lack of authentication"
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "customer": "scope of registered customers",
        "inventoryManager": "scope of resellers acting as inventory managers"
      }
    },
    "isRegistered": {
      "type": "basic"
    },
    "isReseller": {
      "type": "apiKey",
      "name": "X-Custom-Key",
      "in": "header"
    },
    "isResellerQuery": {
      "type": "apiKey",
      "name": "CustomKeyAsQuery",
      "in": "query"
    }
  },
  "security": [
    {
      "isRegistered": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This sample API demonstrates how to compose several authentication schemes \nand configure complex security requirements for your operations.\n\nThis API simulates a very simple market place with customers and resellers  \nof items.\n\nPersonas:\n  - as a first time user, I want to see all items on sales\n  - as a registered customer, I want to post orders for items and \n    consult my past orders\n  - as a registered reseller, I want to see all pending orders on the items \n    I am selling on the market place\n  - as a reseller managing my own inventories, I want to post replenishment orders for the items I provide\n  - as a register user, I want to consult my personal account infos\n\nThe situation we defined on the authentication side is as follows:\n  - every known user is authenticated using a basic token\n  - resellers are authenticated using API keys - we let the option to authenticate using a header or query param\n  - any registered user (customer or reseller) will add a signed JWT to access more API endpoints\n\nObviously, there are several ways to achieve the same result. We just wanted to demonstrate here how\nsecurity requirements may compose several schemes.\n\nNote that we used the \"OAuth2\" declaration here but don't implement a real  \nOAuth2 workflow: our intend here is just to be able to extract scopes from a passed JWT token (the \nonly way to manipulate scoped authorizers with Swagger 2.0 is to declare them with type \"oauth2\").\n",
    "title": "Composing authorizations",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/account": {
      "get": {
        "security": [
          {
            "isRegistered": []
          }
        ],
        "description": "Every registered user should be able to access this operation\n",
        "summary": "registered user account",
        "operationId": "GetAccount",
        "responses": {
          "200": {
            "description": "registered user personal account infos",
            "schema": {
              "type": "object",
              "additionalProperties": true
            }
          },
          "401": {
            "description": "unauthorized access for a lack of authentication"
          },
          "default": {
            "description": "other error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/items": {
      "get": {
        "security": [],
        "description": "Everybody should be able to access this operation\n",
        "summary": "items on sale",
        "operationId": "GetItems",
        "responses": {
          "200": {
            "description": "multiple items",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Item"
              }
            }
          },
          "default": {
            "description": "other error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/order/add": {
      "post": {
        "security": [
          {
            "hasRole": [
              "customer"
            ],
            "isRegistered": []
          },
          {
            "hasRole": [
              "inventoryManager"
            ],
            "isReseller": []
          },
          {
            "hasRole": [
              "inventoryManager"
            ],
            "isResellerQuery": []
          }
        ],
        "description": "Registered customers should be able to add purchase orders.\nRegistered inventory managers should be able to add replenishment orders.\n",
        "summary": "post a new order",
        "operationId": "AddOrder",
        "parameters": [
          {
            "name": "order",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "empty response"
          },
          "401": {
            "description": "unauthorized access for a lack of authentication"
          },
          "403": {
            "description": "forbidden access for a lack of sufficient privileges"
          },
          "default": {
            "description": "other error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/order/{orderID}": {
      "get": {
        "security": [
          {
            "hasRole": [
              "customer"
            ],
            "isRegistered": []
          }
        ],
        "description": "Only registered customers should be able to retrieve orders\n",
        "summary": "retrieves an order",
        "operationId": "GetOrder",
        "parameters": [
          {
            "type": "string",
            "name": "orderID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "content of an order",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "401": {
            "description": "unauthorized access for a lack of authentication"
          },
          "403": {
            "description": "forbidden access for a lack of sufficient privileges"
          },
          "default": {
            "description": "other error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/orders/{itemID}": {
      "get": {
        "security": [
          {
            "isReseller": []
          },
          {
            "isResellerQuery": []
          }
        ],
        "description": "Only registered resellers should be able to search orders for an item\n",
        "summary": "retrieves all orders for an item",
        "operationId": "GetOrdersForItem",
        "parameters": [
          {
            "type": "string",
            "name": "itemID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "multiple orders",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Order"
              }
            }
          },
          "401": {
            "description": "unauthorized access for a lack of authentication"
          },
          "403": {
            "description": "forbidden access for a lack of sufficient privileges"
          },
          "default": {
            "description": "other error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Item": {
      "type": "string"
    },
    "Order": {
      "type": "object",
      "required": [
        "orderID"
      ],
      "properties": {
        "orderID": {
          "type": "string"
        },
        "orderLines": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/orderOrderLinesItems"
          }
        }
      }
    },
    "orderOrderLinesItems": {
      "type": "object",
      "required": [
        "quantity",
        "purchasedItem"
      ],
      "properties": {
        "purchasedItem": {
          "$ref": "#/definitions/Item"
        },
        "quantity": {
          "type": "string",
          "format": "uint32",
          "minimum": 1
        }
      },
      "x-go-gen-location": "models",
      "x-go-name": "orderLine"
    },
    "principal": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  },
  "responses": {
    "forbidden": {
      "description": "forbidden access for a lack of sufficient privileges"
    },
    "multipleItems": {
      "description": "multiple items",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Item"
        }
      }
    },
    "multipleOrders": {
      "description": "multiple orders",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Order"
        }
      }
    },
    "otherError": {
      "description": "other error response",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "singleItem": {
      "description": "single item",
      "schema": {
        "type": "string"
      }
    },
    "singleOrder": {
      "description": "content of an order",
      "schema": {
        "$ref": "#/definitions/Order"
      }
    },
    "unauthorized": {
      "description": "unauthorized access for a lack of authentication"
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "customer": "scope of registered customers",
        "inventoryManager": "scope of resellers acting as inventory managers"
      }
    },
    "isRegistered": {
      "type": "basic"
    },
    "isReseller": {
      "type": "apiKey",
      "name": "X-Custom-Key",
      "in": "header"
    },
    "isResellerQuery": {
      "type": "apiKey",
      "name": "CustomKeyAsQuery",
      "in": "query"
    }
  },
  "security": [
    {
      "isRegistered": []
    }
  ]
}`))
}
