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
        "/": {
            "get": {
                "description": "Responds with an HTML page with the index page",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "html"
                ],
                "summary": "Serve HTML index page",
                "responses": {
                    "200": {
                        "description": "HTML Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pages/totals/rough": {
            "get": {
                "description": "Responds with an HTML page with a rought total graph",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "html"
                ],
                "summary": "Serve HTML rought total page",
                "responses": {
                    "200": {
                        "description": "HTML Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/participants": {
            "get": {
                "description": "Responds with the list of participants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "participants"
                ],
                "summary": "Get Participants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Participant"
                            }
                        }
                    }
                }
            }
        },
        "/votes": {
            "post": {
                "description": "Cast a Vote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "votes"
                ],
                "summary": "Post Vote",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Vote"
                        }
                    }
                }
            }
        },
        "/votes/totals/rough": {
            "get": {
                "description": "Get rough totals",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "totals votes"
                ],
                "summary": "Get Rough Totals",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Participant": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Vote": {
            "type": "object",
            "properties": {
                "participant": {
                    "$ref": "#/definitions/domain.Participant"
                },
                "timestamp": {
                    "type": "string"
                },
                "voteID": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
