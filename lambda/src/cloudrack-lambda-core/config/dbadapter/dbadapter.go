// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package dbadapter

import (
	core "cloudrack-lambda-core/core"
	model "cloudrack-lambda-core/config/model"
	"strconv"
	"strings"
)

var ITEM_TYPE_CONFIG_PREFIX string = "cfg"
var ITEM_TYPE_CONFIG_GENERAL string = ITEM_TYPE_CONFIG_PREFIX + "-general"
var ITEM_TYPE_CONFIG_HISTORY string = ITEM_TYPE_CONFIG_PREFIX + "-history"
var ITEM_TYPE_CONFIG_INVENTORY_ROOM string = ITEM_TYPE_CONFIG_PREFIX + "-room"
var ITEM_TYPE_CONFIG_INVENTORY_ROOM_TYPE string = ITEM_TYPE_CONFIG_INVENTORY_ROOM + "-type"


func IdToDynamo(id string) model.DynamoHotelRec{
	return model.DynamoHotelRec{ Code : id, ItemType: ITEM_TYPE_CONFIG_GENERAL }
}

func DynamoToBom(dbr model.DynamoHotelRec) model.Hotel{
	bom := model.Hotel{ Code : dbr.Code, 
		Name : dbr.Name,
		Description : dbr.Description,
		CurrencyCode : dbr.CurrencyCode,
		DefaultMarketting : model.HotelMarketing{
			DefaultTagLine : dbr.DefaultTagLine,
			DefaultDescription : dbr.DefaultDescription},
		Options: model.HotelOptions{} }

		bom.Options.Bookable,_ = strconv.ParseBool(dbr.Bookable)
		bom.Options.Shoppable,_ = strconv.ParseBool(dbr.Shoppable)

		return bom
}

func DynamoListToBom(dbrList []model.DynamoHotelRec) model.Hotel{
	var bom model.Hotel
	var hotelConfigChanges []model.HotelConfigChange = make([]model.HotelConfigChange,0,0)

	for _, dbr := range dbrList {
			if(strings.HasPrefix(dbr.ItemType,ITEM_TYPE_CONFIG_INVENTORY_ROOM_TYPE)) {
				bom.RoomTypes  = append(bom.RoomTypes,DynamoToBomRoomType(dbr))
			} else if(dbr.ItemType == ITEM_TYPE_CONFIG_GENERAL) {
				bom = DynamoToBom(dbr)
			} else if(strings.HasPrefix(dbr.ItemType,ITEM_TYPE_CONFIG_HISTORY)) {
				hotelConfigChanges = append(hotelConfigChanges,DynamoToBomConfigChange(dbr))
			}
	}		
	bom.PendingChanges = hotelConfigChanges
	return bom
}


func BomToDynamo(bom model.Hotel, user core.User) model.DynamoHotelRec{
	return model.DynamoHotelRec{ Code : bom.Code, 
		Name : bom.Name,
		Description : bom.Description,
		CurrencyCode : bom.CurrencyCode,
		DefaultTagLine : bom.DefaultMarketting.DefaultTagLine,
		DefaultDescription : bom.DefaultMarketting.DefaultDescription,
		Bookable : strconv.FormatBool(bom.Options.Bookable),
		Shoppable : strconv.FormatBool(bom.Options.Shoppable),
		User : user.Username,
		LastUpdatedBy : user.Username,
		ItemType: ITEM_TYPE_CONFIG_GENERAL }
}

func BomRoomTypeToDynamo(bom model.Hotel, user core.User) model.DynamoRoomTypeRec{
	return model.DynamoRoomTypeRec{ 
		Code : bom.Code, 
		Name : bom.RoomTypes[0].Name,
		Description : bom.RoomTypes[0].Description,
		LowPrice : bom.RoomTypes[0].LowPrice,
    	MedPrice : bom.RoomTypes[0].MedPrice,
    	HighPrice : bom.RoomTypes[0].HighPrice,
		ItemType: ITEM_TYPE_CONFIG_INVENTORY_ROOM_TYPE + "-" +  bom.RoomTypes[0].Code,
		LastUpdatedBy : user.Username,
		}
}

//Bom to generic Dynamo reccord to be used for delete
func BomRoomTypeToDynamoRecord(bom model.Hotel) core.DynamoRecord{
	return core.DynamoRecord{ 
		Code : bom.Code, 
		ItemType: ITEM_TYPE_CONFIG_INVENTORY_ROOM_TYPE + "-" +  bom.RoomTypes[0].Code}
}


func DynamoToBomRoomType(dbr model.DynamoHotelRec) model.HotelRoomType{
	return model.HotelRoomType{ 
		Code : getDbrCode(dbr.ItemType), 
		Name : dbr.Name,
		LowPrice : dbr.LowPrice,
    	MedPrice : dbr.MedPrice,
    	HighPrice : dbr.HighPrice,
		Description : dbr.Description}
}

func DynamoToBomConfigChange(histRec model.DynamoHotelRec) model.HotelConfigChange{
	return model.HotelConfigChange{
    TimeStamp : histRec.TimeStamp,
    EventName  : histRec.EventName,
    ObjectName  : histRec.ObjectName ,
    FieldName  : histRec.FieldName,
    OldValue  : histRec.OldValue,
    NewValue  : histRec.NewValue,
    HumanReadableComment : histRec.HumanReadableComment,
	}
}

//returns code form bynamo dbitem by splitinting itemType field and getting teh last element following our datamodel convention
func getDbrCode(itemType string) string  {
	return strings.Split(itemType,"-")[len(strings.Split(itemType,"-"))-1]
}
