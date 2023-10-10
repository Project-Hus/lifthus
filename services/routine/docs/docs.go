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
        "/act": {
            "post": {
                "tags": [
                    "act"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "lifthus_st",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create act dto",
                        "name": "creatActDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateActRequestDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/act/upgrade": {
            "post": {
                "tags": [
                    "act"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "lifthus_st",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "upgrade act dto",
                        "name": "upgradeActDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpgradeActRequestDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/act/{code}": {
            "get": {
                "tags": [
                    "act"
                ],
                "summary": "get act by code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "act code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns act info as json"
                    },
                    "400": {
                        "description": "invalid request"
                    },
                    "404": {
                        "description": "act not found"
                    },
                    "500": {
                        "description": "internal server error"
                    }
                }
            }
        },
        "/images/{target}": {
            "post": {
                "tags": [
                    ""
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "images for target",
                        "name": "target",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "images of act",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "lifthus_st",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/program/weekly": {
            "post": {
                "tags": [
                    "program"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "lifthus_st",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create program dto",
                        "name": "creatProgramDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProgramRequestDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/program/{code}": {
            "get": {
                "tags": [
                    "program"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "program code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.CreateActRequestDto": {
            "type": "object",
            "properties": {
                "actType": {
                    "type": "string"
                },
                "author": {
                    "type": "string"
                },
                "imageSrcs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dto.CreateProgramRequestDailyRoutineDto": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "integer"
                },
                "routineActs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateProgramRequestRoutineActDto"
                    }
                }
            }
        },
        "dto.CreateProgramRequestDto": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "dailyRoutines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateProgramRequestDailyRoutineDto"
                    }
                },
                "derivedFrom": {
                    "type": "string"
                },
                "imageSrcs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "programType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.CreateProgramRequestRoutineActDto": {
            "type": "object",
            "properties": {
                "actVersion": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "ratioOrSecs": {
                    "type": "number"
                },
                "repsOrMeters": {
                    "type": "integer"
                },
                "stage": {
                    "type": "string"
                }
            }
        },
        "dto.UpgradeActRequestDto": {
            "type": "object",
            "properties": {
                "actCode": {
                    "type": "string"
                },
                "imageSrcs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Version:          "0.0.0",
	Host:             "localhost:9100",
	BasePath:         "/routine",
	Schemes:          []string{},
	Title:            "Lifthus routine server",
	Description:      "This is Project-Hus's subservice Lifthus's routine management server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
