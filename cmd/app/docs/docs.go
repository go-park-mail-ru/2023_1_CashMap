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
            "name": "CashMap team"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/logout": {
            "post": {
                "description": "Delete user session and invalidate session cookie",
                "tags": [
                    "logout"
                ],
                "summary": "Log out",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Authorize client with credentials (login and password).",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "signin"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "User login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Register client with credentials and other user info.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "signup"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "User email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User first name",
                        "name": "first_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User last name",
                        "name": "last_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/feed": {
            "get": {
                "description": "Get users's new feed part by last post id and batch size.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "Get feed part",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Posts amount",
                        "name": "batch_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date and time of last post given. If not specified the newest posts will be sent",
                        "name": "last_post_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Comment": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "postID": {
                    "type": "integer"
                },
                "replyTo": {
                    "description": "id коммента в посте, к которому сделан коммент. null, если коммент верхнего уровня",
                    "type": "integer"
                },
                "sender": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "entities.Post": {
            "description": "All post information",
            "type": "object",
            "properties": {
                "attachments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Comment"
                    }
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "likes": {
                    "type": "integer"
                },
                "sender_name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Depeche API",
	Description:      "Api for Depeche social network. VK Education project. Spring, 2023",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
