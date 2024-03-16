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
        "license": {
            "name": "MIT license"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/queue": {
            "get": {
                "description": "Returns a list of queue names",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Queue list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.QueueList"
                        }
                    }
                }
            },
            "put": {
                "description": "Creates a new queue with a name and size",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Creates a new queue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Queue name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Queue size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.QueueCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a queue by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Deletes a queue",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Queue name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.QueueDeleteResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/queue/{qname}": {
            "get": {
                "description": "Returns information about the queue by its name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queue"
                ],
                "summary": "Information about queue",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.QueueInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "info": {
                    "type": "string"
                }
            }
        },
        "response.QueueCreateResponse": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "response.QueueDeleteResponse": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "response.QueueInfo": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "head": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "tail": {
                    "type": "integer"
                }
            }
        },
        "response.QueueList": {
            "type": "object",
            "properties": {
                "queue_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7920",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Limero",
	Description:      "This is a message broker",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}