// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
		db "cloudrack-lambda-core/db"
		model "cloudrack-lambda-core/config/model"
		usecase "cloudrack-lambda-fn/usecase"
		"os"
        "context"
        "log"
        "github.com/aws/aws-lambda-go/lambda"
        "errors"
)



var FN_CONFIG_LIST_HOTEL string = "listHotels"
var FN_CONFIG_SAVE_HOTEL string = "saveHotel"
var FN_CONFIG_GET_HOTEL string = "getHotel"
var FN_CONFIG_DELETE_HOTEL string = "deleteHotel"
var FN_CONFIG_ADD_PICTURE string = "addPicture"
var FN_CONFIG_DELETE_PICTURE string = "deletePicture"
var FN_CONFIG_SAVE_ROOMS string = "saveRooms"
var FN_CONFIG_ADD_SELLABLE string = "addSellable"
var FN_CONFIG_ADD_TOOM_TYPE string = "addRoomType"
var FN_CONFIG_ADD_TAG string = "addTag"
var FN_CONFIG_DELETE_SELLABLE string = "deleteSellable"
var FN_CONFIG_DELETE_ROOM_TYPE string = "deleteRoomType"
var FN_CONFIG_DELETE_TAG string = "deleteTag"
var FN_CONFIG_ADD_SELLABLE_PICTURE string = "addSellablePicture"
var FN_CONFIG_ADD_ROOM_TYPE_PICTURE string = "addRoomTypePicture"
var FN_CONFIG_DELETE_SELLABLE_PICTURE string = "deleteSellablePicture"
var FN_CONFIG_DELETE_ROOM_TYPE_PICTURE string = "deleteRoomTypePicture"
var FN_CONFIG_PUBLISH_CHANGES string = "publishChanges"

var LAMBDA_ENV = os.Getenv("LAMBDA_ENV")
var DB_TABLE_CONFIG_NAME string = "aws-crud-demo-config-"+LAMBDA_ENV
var DB_TABLE_CONFIG_PK string = "code"
var DB_TABLE_CONFIG_SK string = "itemType"

//It is a best practice to instanciate the dynamoDB client outside
//of the lambda function handler.
//https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.Lambda.BestPracticesWithDynamoDB.html
var configDb =  db.Init(DB_TABLE_CONFIG_NAME,DB_TABLE_CONFIG_PK, DB_TABLE_CONFIG_SK)

func HandleRequest(ctx context.Context, wrapper model.RqWrapper) (model.ResWrapper, error) {
	log.Printf("Received Request %+v", wrapper)
	//Validate request body with rues (may need to be a different function per use case)
	err := ValidateRequest(wrapper.Request)
	res := model.ResWrapper{}
	if(err == nil) {
		if(wrapper.SubFunction == FN_CONFIG_LIST_HOTEL) {
			log.Printf("Selected Use Case %v",FN_CONFIG_LIST_HOTEL)
			res, err = usecase.ListHotel(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_SAVE_HOTEL) {
			log.Printf("Selected Use Case %v",FN_CONFIG_SAVE_HOTEL)
			res, err = usecase.SaveHotel(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_GET_HOTEL) {
			log.Printf("Selected Use Case %v",FN_CONFIG_GET_HOTEL)
			res, err = usecase.GetHotel(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_PICTURE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_PICTURE)
			res, err = usecase.AddPictures(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_PICTURE) {
				log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_PICTURE)
				res, err = usecase.DeletePicture(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_SAVE_ROOMS) {
			log.Printf("Selected Use Case %v",FN_CONFIG_SAVE_ROOMS)
			res, err = usecase.SaveRooms(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_SELLABLE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_SELLABLE)
			res, err = usecase.AddSellable(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_TOOM_TYPE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_TOOM_TYPE)
			res, err = usecase.AddRoomType(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_TAG) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_TAG)
			res, err = usecase.AddTag(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_SELLABLE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_SELLABLE)
			res, err = usecase.DeleteSellable(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_ROOM_TYPE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_ROOM_TYPE)
			res, err = usecase.DeleteRoomType(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_TAG) {
			log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_TAG)
			res, err = usecase.DeleteTag(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_SELLABLE_PICTURE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_SELLABLE_PICTURE)
			res, err = usecase.AddSellablePicture(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_ADD_ROOM_TYPE_PICTURE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_ADD_ROOM_TYPE_PICTURE)
			res, err = usecase.AddRoomTypePicture(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_SELLABLE_PICTURE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_SELLABLE_PICTURE)
			res, err = usecase.DeleteSellablePicture(wrapper,configDb);
		} else if(wrapper.SubFunction == FN_CONFIG_DELETE_ROOM_TYPE_PICTURE) {
			log.Printf("Selected Use Case %v",FN_CONFIG_DELETE_ROOM_TYPE_PICTURE)
			res, err = usecase.DeleteRoomTypePicture(wrapper,configDb)
		} else if(wrapper.SubFunction == FN_CONFIG_PUBLISH_CHANGES) {
			log.Printf("Selected Use Case %v",FN_CONFIG_PUBLISH_CHANGES)
			res, err = usecase.PublishChanges(wrapper,configDb)
		} else {
			log.Printf("No Use Case Found for %v",wrapper.SubFunction)
			err = errors.New("No Use Case Found for sub function " + wrapper.SubFunction)
		}
	}
	log.Printf("Lambda function execusion completed with res: %+v and error: %+v", res,err)
	return res, err
}

//TODO: Write the validation logic
func ValidateRequest(hotel model.Hotel) error{
		return nil
}



func main() {
        lambda.Start(HandleRequest)
}