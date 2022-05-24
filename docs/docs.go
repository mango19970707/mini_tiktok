// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "get": {
                "description": "get auth info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "get auth info",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
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
                            "$ref": "#/definitions/biz.QueryInfoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryInfoResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "password",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "password",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/comment/action": {
            "post": {
                "description": "put comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "comment a video",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "video_id",
                        "name": "video_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "1 - comment 2 - delete",
                        "name": "action_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "needed when action_type=1",
                        "name": "comment_text",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "needed when action_type=2",
                        "name": "comment_id",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/biz.CommentActionResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.CommentActionResponse"
                        }
                    }
                }
            }
        },
        "/comment/list": {
            "get": {
                "description": "get comment list by video id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "get comment list",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "video_id",
                        "name": "video_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryCommentListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryCommentListResponse"
                        }
                    }
                }
            }
        },
        "/favorite/action": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "favorite"
                ],
                "summary": "like video action",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "video_id",
                        "name": "video_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "1 - like 2 - unlike",
                        "name": "action_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/biz.FavoriteActionResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.FavoriteActionResponse"
                        }
                    }
                }
            }
        },
        "/favorite/list": {
            "get": {
                "description": "get favorite list by userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "favorite"
                ],
                "summary": "get favorite list",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
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
                            "$ref": "#/definitions/biz.QueryFavoriteListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryFavoriteListResponse"
                        }
                    }
                }
            }
        },
        "/feed": {
            "get": {
                "description": "get feed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "get feed",
                "parameters": [
                    {
                        "description": "the latest time to get",
                        "name": "latest_time",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/biz.FeedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.FeedResponse"
                        }
                    }
                }
            }
        },
        "/publish/action": {
            "post": {
                "description": "publish video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publish"
                ],
                "summary": "publish video",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
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
                            "$ref": "#/definitions/biz.PublishActionResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.PublishActionResponse"
                        }
                    }
                }
            }
        },
        "/publish/list": {
            "get": {
                "description": "get publish video by userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publish"
                ],
                "summary": "get publish list",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "video data",
                        "name": "data",
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
                            "$ref": "#/definitions/biz.QueryPublishListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryPublishListResponse"
                        }
                    }
                }
            }
        },
        "/relation/action": {
            "post": {
                "description": "follow and unfollow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "relation"
                ],
                "summary": "follow and unfollow",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "to_user_id",
                        "name": "to_user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "1 - follow 2 - unfollow",
                        "name": "action_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/biz.RelationActionResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.RelationActionResponse"
                        }
                    }
                }
            }
        },
        "/relation/follow/list": {
            "get": {
                "description": "get follow list by userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "relation"
                ],
                "summary": "get follow list",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
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
                            "$ref": "#/definitions/biz.QueryFollowListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryFollowListResponse"
                        }
                    }
                }
            }
        },
        "/relation/follower/list": {
            "get": {
                "description": "get follower list by userId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "relation"
                ],
                "summary": "get follower list",
                "parameters": [
                    {
                        "description": "user_id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "token",
                        "name": "token",
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
                            "$ref": "#/definitions/biz.QueryFollowerListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/biz.QueryFollowerListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "auth.RegisterResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "biz.Comment": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/biz.User"
                },
                "content": {
                    "type": "string"
                },
                "create_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "biz.CommentActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.FavoriteActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.FeedResponse": {
            "type": "object",
            "properties": {
                "next_time": {
                    "type": "integer"
                },
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "video": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.Video"
                    }
                }
            }
        },
        "biz.PublishActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.QueryCommentListResponse": {
            "type": "object",
            "properties": {
                "comment_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.Comment"
                    }
                },
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.QueryFavoriteListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "video_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.Video"
                    }
                }
            }
        },
        "biz.QueryFollowListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.User"
                    }
                }
            }
        },
        "biz.QueryFollowerListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.User"
                    }
                }
            }
        },
        "biz.QueryInfoResponse": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/biz.User"
                },
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.QueryPublishListResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "video_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/biz.Video"
                    }
                }
            }
        },
        "biz.RelationActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                }
            }
        },
        "biz.User": {
            "type": "object",
            "properties": {
                "follow_count": {
                    "type": "integer"
                },
                "follower_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_follow": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "biz.Video": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/biz.User"
                },
                "comment_count": {
                    "type": "integer"
                },
                "cover_url": {
                    "type": "string"
                },
                "favorite_count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "play_url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/douyin",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}