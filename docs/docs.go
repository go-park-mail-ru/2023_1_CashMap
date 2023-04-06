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
            "name": "CashMap team"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/feed": {
            "get": {
                "description": "Get users's new feed part by last post id and batch size.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
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
        },
        "/api/user/friends": {
            "get": {
                "description": "Get friends",
                "tags": [
                    "Profiles"
                ],
                "summary": "Friends",
                "parameters": [
                    {
                        "type": "string",
                        "description": "link to requested profile",
                        "name": "link",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of profiles",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "number of batch",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Profile"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/profile": {
            "get": {
                "description": "Get self profile",
                "tags": [
                    "Profiles"
                ],
                "summary": "Self",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/profile/edit": {
            "post": {
                "description": "Edit profile",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "summary": "EditProfile",
                "parameters": [
                    {
                        "description": "Edited fields",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/profile/link": {
            "get": {
                "description": "Get profile by link",
                "tags": [
                    "Profiles"
                ],
                "summary": "Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "link to requested profile",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/reject": {
            "post": {
                "description": "Reject friend request",
                "tags": [
                    "Subscribes"
                ],
                "summary": "Reject",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/sub": {
            "get": {
                "description": "Get subscribes or subscribers for requested user",
                "tags": [
                    "Profiles"
                ],
                "summary": "Subscribes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "in/out for subscribers/subscribes",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "link to requested profile",
                        "name": "link",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of profiles",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "number of batch",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Profile"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Subscribe to other user",
                "tags": [
                    "Subscribes"
                ],
                "summary": "Subscribe",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/unsub": {
            "post": {
                "description": "Unsubscribe from other user",
                "tags": [
                    "Subscribes"
                ],
                "summary": "Unsubscribe",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Delete user session and invalidate session cookie",
                "tags": [
                    "Auth"
                ],
                "summary": "Log out",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
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
                    "Auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
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
                    "Auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "Required register fields",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUp"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Profile": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": ""
                },
                "bio": {
                    "type": "string",
                    "example": "Текст с информацией о себе."
                },
                "birthday": {
                    "type": "string",
                    "example": "30.04.2002"
                },
                "date_joined": {
                    "type": "string",
                    "example": "10.02.2023"
                },
                "email": {
                    "type": "string",
                    "example": "example@mail.ru"
                },
                "first_name": {
                    "type": "string",
                    "example": "Василий"
                },
                "last_active": {
                    "type": "string",
                    "example": ""
                },
                "last_name": {
                    "type": "string",
                    "example": "Петров"
                },
                "link": {
                    "type": "string",
                    "example": "id100500"
                },
                "private": {
                    "type": "boolean",
                    "example": false
                },
                "sex": {
                    "type": "string",
                    "example": "male"
                },
                "status": {
                    "type": "string",
                    "example": "Текст статуса."
                }
            }
        },
        "dto.SignIn": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SignUp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entities.Comment": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "post_id": {
                    "type": "integer"
                },
                "reply_to": {
                    "type": "integer"
                },
                "sender_name": {
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
        },
        "middleware.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Невалидный запрос."
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Depeche API",
	Description:      "Api for Depeche social network. VK Education project. Spring, 2023",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
