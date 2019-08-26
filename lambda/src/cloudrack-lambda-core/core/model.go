// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package core

var DYNAMO_STREAM_EVENT_NAME_CREATE string = "INSERT"
var DYNAMO_STREAM_EVENT_NAME_UPDATE string = "MODIFY"
var DYNAMO_STREAM_EVENT_NAME_DELETE string = "REMOVE"

type User struct {
	Sub   string `json:"sub"`
        Email string `json:"email"`
        Username string `json:"username"`
}

type ResError struct{
        Code string `json:"code"`
        Msg string `json:"msg"`
}

type DynamoRecord struct{
        Code string `json:"code"`
        ItemType string `json:"itemType"`
}
