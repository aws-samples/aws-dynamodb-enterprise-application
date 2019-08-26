// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package usercase

import (
	core "cloudrack-lambda-core/core"
	db "cloudrack-lambda-core/db"
	model "cloudrack-lambda-core/config/model"
	dbadapter "cloudrack-lambda-core/config/dbadapter"
	"log"
	"errors"
	"strconv"
)


func GetHotel(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	dynRes := []model.DynamoHotelRec{}
	dbr := dbadapter.IdToDynamo(wrapper.Id)
	err := configDb.FindStartingWith(dbr.Code, dbadapter.ITEM_TYPE_CONFIG_PREFIX, &dynRes)
	if err != nil {
		log.Printf("Cannot din property with ID", wrapper.Id)
	}
	return model.ResWrapper{Response: []model.Hotel{ dbadapter.DynamoListToBom(dynRes)}}, err
}

func ListHotel(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	results := []model.Hotel{}
	err := configDb.FindByGsi(wrapper.UserInfo.Username, "user-index", "user", &results)
	log.Printf("List Properties for user %+v", results)
	return model.ResWrapper{Response: results}, err
}

func SaveHotel(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
		dbr := dbadapter.BomToDynamo(wrapper.Request, wrapper.UserInfo)
		dbr.ItemType = dbadapter.ITEM_TYPE_CONFIG_GENERAL
		if(dbr.Code == "") {
		dbr.Code = GeneratePropertyCode(wrapper.UserInfo,dbr.Name)
		if Exists(dbr,configDb) {
				return model.ResWrapper{}, errors.New("object already exists")
		}
		log.Printf("No Property code provided: Creating property %+v", wrapper.Request)
		} else {
			log.Printf("Updating Property %+v", wrapper.Request)
		}
		_, err := configDb.Save(dbr)
		res := wrapper.Request
		res.Code = dbr.Code
		return model.ResWrapper{Response: []model.Hotel{res}}, err
}

func Exists(obj model.DynamoHotelRec,configDb db.DBConfig) bool {
	log.Printf("Exists>Checking existend of hotel with code  %+v", obj.Code)
	hotel := model.Hotel{}
	configDb.Get(obj.Code, obj.ItemType, &hotel)
	log.Printf("Exists>DB returned %+v", hotel)
	return hotel.Code != ""
}

func GeneratePropertyCode(userInfo core.User, propertyName string) string{
	return strconv.FormatInt(int64(core.Hash(userInfo.Username + propertyName)), 10)
}

func AddRoomType(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
		dbr := dbadapter.BomRoomTypeToDynamo(wrapper.Request, wrapper.UserInfo)
		_, err := configDb.Save(dbr)
		return model.ResWrapper{Response: []model.Hotel{wrapper.Request}},err
}

func DeleteRoomType(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	dbr := dbadapter.BomRoomTypeToDynamoRecord(wrapper.Request)
	_, err := configDb.Delete(dbr)
	return model.ResWrapper{Response: []model.Hotel{wrapper.Request}},err
}

func AddSellable(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func AddTag(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}
func DeleteSellable(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func SaveRooms(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func AddSellablePicture(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}
func AddRoomTypePicture(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func DeleteSellablePicture(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func DeleteRoomTypePicture(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func DeleteTag(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func AddPictures(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func DeletePicture(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}

func PublishChanges(wrapper model.RqWrapper,configDb db.DBConfig) (model.ResWrapper, error) {
	//TODO: handle this use case
	return model.ResWrapper{}, nil
}