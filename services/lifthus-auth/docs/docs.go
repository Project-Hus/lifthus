// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "lifthus531@gmail.com",
            "email": "lifthus531@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/account": {
            "delete": {
                "tags": [
                    "auth"
                ],
                "summary": "deletes user's lifthus account",
                "responses": {
                    "200": {
                        "description": "Ok, the account is deleted"
                    },
                    "400": {
                        "description": "Bad Request, invalid request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/hus/signin": {
            "patch": {
                "description": "the \"signin_propagation\" token should be included in the request body.",
                "tags": [
                    "auth"
                ],
                "summary": "processes user sign-in propagation from cloudhus.",
                "responses": {
                    "200": {
                        "description": "Ok, session signed"
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
        "/hus/signout": {
            "patch": {
                "description": "the \"signout_propagation\" token should be included in the request body.",
                "tags": [
                    "auth"
                ],
                "summary": "processes user sign-out propagation from cloudhus.",
                "responses": {
                    "200": {
                        "description": "Ok, session signed"
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
        "/relation/follow/{uid}": {
            "post": {
                "tags": [
                    "relation"
                ],
                "summary": "gets uid from path param and makes signed user follow the given user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "new following list"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user following list"
                    }
                }
            }
        },
        "/relation/followers/{uid}": {
            "get": {
                "tags": [
                    "relation"
                ],
                "summary": "gets uid from path param and returns user's follower list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns follower list as list of number"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user follower list"
                    }
                }
            }
        },
        "/relation/following/{uid}": {
            "get": {
                "tags": [
                    "relation"
                ],
                "summary": "gets uid from path param and returns user's following list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns following list as list of number"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user following list"
                    }
                }
            }
        },
        "/relation/unfollow/{uid}": {
            "delete": {
                "tags": [
                    "relation"
                ],
                "summary": "gets uid from path param and makes signed user unfollow the given user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "new following list"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user following list"
                    }
                }
            }
        },
        "/session": {
            "get": {
                "tags": [
                    "auth"
                ],
                "summary": "validates session. publishes new one if it isn't. refreshes expired session.",
                "responses": {
                    "200": {
                        "description": "Ok, session refreshed, session info JSON returned"
                    },
                    "201": {
                        "description": "Created, new session issued, redirect to cloudhus and do connect"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/session/signout": {
            "patch": {
                "tags": [
                    "auth"
                ],
                "summary": "gets sign-out request from the client and propagates it to Cloudhus.",
                "responses": {
                    "200": {
                        "description": "Ok, signed out of the session"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized, the token is expired or the session is not signed"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sid": {
            "get": {
                "tags": [
                    "auth"
                ],
                "summary": "returns client's SID. should be encrypted later.",
                "responses": {
                    "200": {
                        "description": "Ok, session ID"
                    },
                    "401": {
                        "description": "Unauthorized, the token is expired"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "it gets register info and registers user to lifthus",
                "tags": [
                    "user"
                ],
                "summary": "gets user register info and registers user",
                "parameters": [
                    {
                        "description": "user register info",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterInfoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns register info as json"
                    },
                    "400": {
                        "description": "invalid body"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            },
            "patch": {
                "description": "it gets uid from path param and updates user info",
                "tags": [
                    "user"
                ],
                "summary": "gets uid from path param and updates user info",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserInfoDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns user info as json"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to set user info"
                    }
                }
            }
        },
        "/user/{uid}": {
            "get": {
                "description": "if the signed user is the same as the requested user, returns all info while hiding sensitive info if different.",
                "tags": [
                    "user"
                ],
                "summary": "gets uid from path param and returns user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns user info as json"
                    },
                    "400": {
                        "description": "invalid uid"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user info"
                    }
                }
            }
        },
        "/username/{username}": {
            "get": {
                "description": "if the signed user is the same as the requested user, returns all info while hiding sensitive info if different.",
                "tags": [
                    "user"
                ],
                "summary": "gets username from path param and returns user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns user info as json"
                    },
                    "404": {
                        "description": "user not found"
                    },
                    "500": {
                        "description": "failed to get user info"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UpdateUserInfoDto": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.RegisterInfoDto": {
            "type": "object",
            "properties": {
                "benchpress": {
                    "type": "number"
                },
                "bodyWeight": {
                    "type": "number"
                },
                "deadlift": {
                    "type": "number"
                },
                "height": {
                    "type": "number"
                },
                "squat": {
                    "type": "number"
                },
                "trainingType": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "auth.lifthus.com",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Lifthus user server",
	Description:      "This is Project-Hus's subservice Lifthus's user management server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
