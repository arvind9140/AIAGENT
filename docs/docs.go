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
        "/initializ/v1/ai/allusers": {
            "get": {
                "description": "Get all Data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserData Apis"
                ],
                "summary": "Get User Data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/casestudy": {
            "get": {
                "description": "Get Case Study Api",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Case Study Apis"
                ],
                "summary": "Get Case Study",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Save Case Study",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Case Study Apis"
                ],
                "summary": "Save Case Study",
                "parameters": [
                    {
                        "description": "Case Study",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Casestudy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/generatewithAI": {
            "post": {
                "description": "Generate with AI",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIAgent Apis"
                ],
                "summary": "Generate with AI",
                "parameters": [
                    {
                        "description": "Generate Body Response",
                        "name": "GenerateAI",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GenerateAIBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/painpoints": {
            "get": {
                "description": "Get all Pain Points and Value Proposition",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pain Points Apis"
                ],
                "summary": "Get Pain Points and Value Proposition",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Save Pain Points and Value Proposition",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pain Points Apis"
                ],
                "summary": "Save Pain Points and Value Proposition",
                "parameters": [
                    {
                        "description": "Pain Points and Value Proposition",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PainPointRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/prompt/{promptId}": {
            "get": {
                "description": "Get AI Prompts by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompt Apis"
                ],
                "summary": "Get Prompt by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "promptId",
                        "name": "promptId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/prompts": {
            "get": {
                "description": "Get all AI Prompts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompt Apis"
                ],
                "summary": "Get Prompts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/saveprompt": {
            "post": {
                "description": "Save Prompt",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompt Apis"
                ],
                "summary": "Save Prompt",
                "parameters": [
                    {
                        "description": "Upload the prompt in the Db",
                        "name": "UploadExcel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Prompts"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/updateprompt/{promptId}": {
            "put": {
                "description": "Update Prompt In Db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompt Apis"
                ],
                "summary": "Update Prompt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "promptId",
                        "name": "promptId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update the prompt in the Db",
                        "name": "UploadExcel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Prompts"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        },
        "/initializ/v1/ai/upload": {
            "post": {
                "description": "Upload Excel File",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserData Apis"
                ],
                "summary": "Upload Excel File",
                "parameters": [
                    {
                        "description": "File Data in base64 encoded",
                        "name": "UploadExcel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UploadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ApplicationResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Casestudy": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "models.GenerateAIBody": {
            "type": "object",
            "properties": {
                "company_url": {
                    "type": "string"
                },
                "linkedin_url": {
                    "type": "string"
                },
                "stream": {
                    "type": "boolean"
                },
                "system_prompt": {
                    "type": "string"
                },
                "task": {
                    "type": "string"
                },
                "to_do_research": {
                    "type": "boolean"
                }
            }
        },
        "models.PainPointRole": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "string"
                }
            }
        },
        "models.Prompts": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "prompt": {
                    "type": "string"
                },
                "prompt_rule": {
                    "type": "string"
                },
                "purpose": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "models.UploadRequest": {
            "type": "object",
            "properties": {
                "file_data": {
                    "type": "string"
                }
            }
        },
        "responses.ApplicationResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Init App aiagent",
	Description:      "Init App aiagent Open Api Spec",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
