// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

import (
	core "cloudrack-lambda-core/core"
)

var SELLABLE_INVENTORY_TYPE_PER_DAY string = "per_day"
var SELLABLE_INVENTORY_TYPE_PER_NIGHT string = "per_night"
var SELLABLE_INVENTORY_TYPE_PER_UNIT string = "per_unit"
var TAG_PRICE_UPDATE_UNIT_PERCENT string = "percent"
var TAG_PRICE_UPDATE_UNIT_CURRENCY string = "currency"

type Hotel struct {
	Code   string `json:"code"`				//mandatory for Dynamo db config table
        Name string `json:"name"`
        Description string `json:"description"`
        CurrencyCode string `json:"currencyCode"`
        DefaultMarketting HotelMarketing `json:"defaultMarketting"`
        Options HotelOptions `json:"options"`
        Pictures []HotelPicture `json:"pictures"`
        Buildings []HotelBuilding `json:"buildings"`
        RoomTypes []HotelRoomType `json:"roomTypes"`
        Tags []HotelRoomAttribute `json:"tags"`
        Sellables []HotelSellable `json:"sellables"`
        Rules []HotelBusinessRule `json:"businessRules"`
        PendingChanges []HotelConfigChange `json:"pendingChanges"` //changes made n config pending publication
}

type HotelBusinessRule struct {
    Id string `json:"id"`
    AppliesTo []HotelBusinessObjectScope `json:"appliesTo"`
    Effect []HotelBusinessRuleEffect `json:"effect"`
    On []HotelDateRange `json:"on"`
    NotOn []HotelDateRange `json:"notOn"`
}

//Defines the scope of a business rule
type HotelBusinessObjectScope struct {
    Type string `json:"type"`  //Room, RoomType, Attribute, Sellable
    Id string `json:"id"`      //unique identifier of the business object
}

//Effect of a business rule on a hotel product
type HotelBusinessRuleEffect struct {
    EffectType string `json:"effectType"` //availability, pricing, inventory
    Available bool `json:"available"`
    PriceImpact int64 `json:"priceImpact"`
    PriceUpdateUnit string `json:"priceUpdateUnit"`
}

//definition of the time frame a rule applies
type HotelDateRange struct {
    From string `json:"from"` //From date
    To string `json:"from"`   //To date
    Dow []string `json:"dow"`   // Day of weeks
}

type HotelPicture struct {
        Id string `json:"id"`
        RawData string `json:"rawData"`
        Format string `json:"format"`
        Url string `json:"url"`
        Tags []string `json:"tags"`
        Main bool `json:"main"`
        PictureItemCode string `json:"pictureItemCode"`       //code of the business object the picture belongs to
}

type HotelSellable struct{
    Code string `json:"code"`
    Category string `json:"category"`
    Name string `json:"name"`
    Quantity int64 `json:"quantity"`
    InventoryType string `json:"inventoryType"` //per_day / per_unit / per night
    PricePerUnit float64 `json:"pricePerUnit"`
    Description string `json:"description"`
    OptionalTags  []HotelRoomAttribute `json:"optionalTags"`
    Pictures []HotelPicture `json:"pictures"`

}

type HotelRoomType struct{
    Code string `json:"code"`
    Name string `json:"name"`
    LowPrice float64 `json:"lowPrice"`
    MedPrice float64 `json:"medPrice"`
    HighPrice float64 `json:"highPrice"`
    Description string `json:"description"`
    Pictures []HotelPicture `json:"pictures"`
}

type HotelRoomAttribute struct {
    Code      string `json:"code"`
    Name string `json:"name"`
    Description      string `json:"description"`
    PriceImpact float64 `json:"priceImpact"`
    PriceUpdateUnit string `json:"priceUpdateUnit"` //TAG_PRICE_UPDATE_UNIT_<?>
    Category      string `json:"category"`
}

type HotelOptions struct {
        Bookable bool `json:"bookable"`
        Shoppable bool `json:"shoppable"`
}

type HotelMarketing struct {
        DefaultTagLine string `json:"defaultTagLine"`
        DefaultDescription string `json:"defaultDescription"`     
}

type HotelBuilding struct {
        Name string `json:"name"`
        Floors []HotelFloor `json:"floors"`
}

type HotelFloor struct {
        Num int32 `json:"num"`
        Rooms []HotelRoom `json:"rooms"`
}

type HotelRoom struct {
        Number      int64 `json:"number"`
        Name        string `json:"name"`
        Type        string `json:"type"`
        Attributes  []HotelRoomAttribute `json:"attributes"`
}

type HotelConfigChange struct{
    TimeStamp string `json:"timeStamp"`  
    EventName string `json:"eventName"`   
    ObjectName string `json:"objectName"`   
    FieldName string `json:"fieldName"`
    OldValue string `json:"oldValue"`
    NewValue string `json:"newValue"`
    HumanReadableComment string `json:"humanReadableComment"`
}

type DynamoHotelHistory struct{
    Code   string `json:"code"`                      //mandatory for Dynamo db config table
    ItemType string `json:"itemType"`               //mandatory for dynamo db config tabble
    TimeStamp string `json:"timeStamp"`  
    EventName string `json:"eventName"`   
    ObjectName string `json:"objectName"`   
    FieldName string `json:"fieldName"`
    OldValue string `json:"oldValue"`
    NewValue string `json:"newValue"`
    HumanReadableComment string `json:"humanReadableComment"`
}

type DynamoHotelRec struct {
        Code   string `json:"code"`                      //mandatory for Dynamo db config table
        ItemType string `json:"itemType"`               //mandatory for dynamo db config tabble
        Name string `json:"name"`
        Description string `json:"description"`
        CurrencyCode string `json:"currencyCode"`
        DefaultTagLine string `json:"defaultTagLine"`
        DefaultDescription string `json:"defaultDescription"`
        Bookable string `json:"bookable"`
        Shoppable string `json:"shoppable"`
        User string `json:"user",omitempty`
        LastUpdatedBy string `json:"lastUpdatedBy"`
        //PICTURE SPECIFIC ATTRIBUTES
        Url string `json:"url"`
        Tags string `json:"tags"`
        Main bool `json:"main"`
        PictureItemCode string `json:"pictureItemCode"`       //code of the business object the picture belongs to
        //ROOM SPECIFIC ATTRIBUTES
        Number      int64 `json:"number"`
        Type        string `json:"type"`
        Floor       int32 `json:"floor"`
        Building    int32 `json:"building"`
        BuildingName string `json:"buildingName"`
        Attributes  string `json:"attributes"`
        //ROOM TYPE, TAGS AND SELLABLE
        Quantity  int64 `json:"quantity"`
        Category  string `json:"category"`
        LowPrice float64 `json:"lowPrice"`
        MedPrice float64 `json:"medPrice"`
        HighPrice float64 `json:"highPrice"`
        PriceImpact float64 `json:"priceImpact"`
        PriceUpdateUnit string `json:"priceUpdateUnit"`
        InventoryType string `json:"inventoryType"`
        PricePerUnit float64 `json:"pricePerUnit"`
        //HISTORY
        TimeStamp string `json:"timeStamp"`  
        EventName string `json:"eventName"`   
        ObjectName string `json:"objectName"`   
        FieldName string `json:"fieldName"`
        OldValue string `json:"oldValue"`
        NewValue string `json:"newValue"`
        HumanReadableComment string `json:"humanReadableComment"`
}

//Use for write (this avoid having indexed attribute sent in request empty (like user))
type DynamoHotelPictureRec struct {
        Code   string `json:"code"`                      //mandatory for Dynamo db config table
        ItemType string `json:"itemType"`               //mandatory for dynamo db config tabble
        //PICTURE SPECIFIC ATTRIBUTES
        PictureItemCode string `json:"pictureItemCode"`       //code of the business object the picture belongs to
        Url string `json:"url"`
        Tags string `json:"tags"`
        Main bool `json:"main"`
        LastUpdatedBy string `json:"lastUpdatedBy"`
}



type DynamoRoomRec struct{
        Code        string `json:"code"`                      //mandatory for Dynamo db config table
        ItemType    string `json:"itemType"`               //mandatory for dynamo db config tabble
        Number      int64 `json:"number"`
        Name        string `json:name"`
        Type        string `json:"type"`
        Floor       int32 `json:"floor"`
        Building    int32 `json:"building"`
        BuildingName string `json:"buildingName"`
        Attributes  string `json:"attributes"`
        LastUpdatedBy string `json:"lastUpdatedBy"`
}

type DynamoSellableRec struct{
    Code string `json:"code"`
    ItemType    string `json:"itemType"`               //mandatory for dynamo db config tabble
    Category string `json:"category"`
    Name string `json:"name"`
    Quantity int64 `json:"quantity"`
    InventoryType string `json:"inventoryType"`
    PricePerUnit float64 `json:"pricePerUnit"`
    Description string `json:"description"`
    LastUpdatedBy string `json:"lastUpdatedBy"`
}

type DynamoRoomTypeRec struct{
    Code string `json:"code"`
    ItemType    string `json:"itemType"`               //mandatory for dynamo db config tabble
    Name string `json:"name"`
    Description string `json:"description"`
    LowPrice float64 `json:"lowPrice"`
    MedPrice float64 `json:"medPrice"`
    HighPrice float64 `json:"highPrice"`
    LastUpdatedBy string `json:"lastUpdatedBy"`
}

type DynamoRoomAttributeRec struct {
    Code      string `json:"code"`
    ItemType    string `json:"itemType"`               //mandatory for dynamo db config tabble
    Name string `json:"name"`
    Description      string `json:"description"`
    Category      string `json:"category"`
    PriceImpact float64 `json:"priceImpact"`
    PriceUpdateUnit string `json:"priceUpdateUnit"`
    LastUpdatedBy string `json:"lastUpdatedBy"`
}


//Implementing core.Batchable interface to allow this struc to be written in batch
func  (drr DynamoRoomRec) GetPk() string {
    return drr.Code
}



//Requests and Respone to deserialize teh Json Payload into. 
//ideally this should be a generic struct in teh core package with a specialised body but 
//the design needs more thoughts so we will include it here for the moment
type RqWrapper struct {
	UserInfo   core.User `json:"userInfo"`
    SubFunction   string `json:"subFunction"`
    Id             string `json:"id"` //ID to be used for GET request
    Request    Hotel `json:"request"`
}


type ResWrapper struct {
    Error   core.ResError `json:"error"`
    SubFunction   string `json:"subFunction"`
    Response    []Hotel `json:"response"`

}