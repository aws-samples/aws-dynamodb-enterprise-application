// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package db

import (
        "fmt"
        "log"
        "reflect"
        core "cloudrack-lambda-core/core"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
     	"github.com/aws/aws-sdk-go/aws/session"
    	"github.com/aws/aws-sdk-go/service/dynamodb"
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

)



type DBConfig struct {
	DbService *dynamodb.DynamoDB
	PrimaryKey string
	SortKey string
	TableName string
}

//init setup teh session and define table name, primary key and sort key
func Init(tn string, pk string, sk string) DBConfig{
		// Initialize a session that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials
		// and region from the shared configuration file ~/.aws/config.
		dbSession := session.Must(session.NewSessionWithOptions(session.Options{
    		SharedConfigState: session.SharedConfigEnable,
		}))
		// Create DynamoDB client
		return DBConfig{
			DbService : dynamodb.New(dbSession),
			PrimaryKey : pk,
			SortKey : sk,
			TableName : tn,
		}
}


func (dbc DBConfig) Save(prop interface{}) (interface{}, error){
		av, err := dynamodbattribute.MarshalMap(prop)
			if err != nil {
	    		fmt.Println("Got error marshalling new property item:")
	    		fmt.Println(err.Error())
			}
		input := &dynamodb.PutItemInput{
    		Item:      av,
    		TableName: aws.String(dbc.TableName),
		}

		_, err = dbc.DbService.PutItem(input)
		if err != nil {
    		fmt.Println("Got error calling PutItem:")
    		fmt.Println(err.Error())
		}
		return prop, err
}

func (dbc DBConfig) Delete(prop interface{}) (interface{}, error){
		av, err := dynamodbattribute.MarshalMap(prop)
			if err != nil {
	    		fmt.Println("Got error marshalling new property item:")
	    		fmt.Println(err.Error())
			}
		
		input := &dynamodb.DeleteItemInput{
    			Key: av,
    			TableName: aws.String(dbc.TableName),
		}
		

		_, err = dbc.DbService.DeleteItem(input)
		if err != nil {
    		fmt.Println("Got error calling DeetItem:")
    		fmt.Println(err.Error())
		}
		return prop, err
}

//TODO: to evaluate th value of this tradeoff: this is probably a little slow but abstract the complexity for all uses of 
//the save many function(and actually any core operation on array of interface)
func InterfaceSlice(slice interface{}) []interface{} {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("InterfaceSlice() given a non-slice type")
    }

    ret := make([]interface{}, s.Len())

    for i:=0; i<s.Len(); i++ {
        ret[i] = s.Index(i).Interface()
    }

    return ret
}

//Writtes many items to a single table
func (dbc DBConfig) SaveMany(data interface{}) error {
	//Dynamo db currently limits batches to 25 items
	batches := core.Chunk(InterfaceSlice(data),25)
	for i, dataArray := range batches {

		log.Printf("DB> Batch %i inserting: %+v",i, dataArray)
		items := make([]*dynamodb.WriteRequest,len(dataArray),len(dataArray))
		for i, item := range dataArray {
			av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
		    		fmt.Println("Got error marshalling new property item:")
		    		fmt.Println(err.Error())
				}
			items[i] = &dynamodb.WriteRequest{
					PutRequest : &dynamodb.PutRequest {
						Item : av,
					},
				}
		}

		bwii := &dynamodb.BatchWriteItemInput {
			RequestItems : map[string][]*dynamodb.WriteRequest{
				dbc.TableName : items,
			},
		}

		_, err := dbc.DbService.BatchWriteItem(bwii)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case dynamodb.ErrCodeResourceNotFoundException:
					fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
					fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
				case dynamodb.ErrCodeRequestLimitExceeded:
					fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return err
		}
	}
	return nil
}		

//Deletes many items to a single table
func (dbc DBConfig) DeleteMany(data interface{}) error {
	//Dynamo db currently limits batches to 25 items
	batches := core.Chunk(InterfaceSlice(data),25)
	for i, dataArray := range batches {

		log.Printf("DB> Batch %i deleting: %+v",i, dataArray)
		items := make([]*dynamodb.WriteRequest,len(dataArray),len(dataArray))
		for i, item := range dataArray {
			av, err := dynamodbattribute.MarshalMap(item)
				if err != nil {
		    		fmt.Println("Got error marshalling new property item:")
		    		fmt.Println(err.Error())
				}
			items[i] = &dynamodb.WriteRequest{
					DeleteRequest : &dynamodb.DeleteRequest {
						Key : av,
					},
				}
		}

		bwii := &dynamodb.BatchWriteItemInput {
			RequestItems : map[string][]*dynamodb.WriteRequest{
				dbc.TableName : items,
			},
		}

		_, err := dbc.DbService.BatchWriteItem(bwii)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case dynamodb.ErrCodeResourceNotFoundException:
					fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
					fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
				case dynamodb.ErrCodeRequestLimitExceeded:
					fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return err
		}
	}
	return nil
}

func (dbc DBConfig) Get(pk string, sk string, data interface{}) error{
	av := map[string]*dynamodb.AttributeValue{
        dbc.PrimaryKey : {
            S: aws.String(pk),
        },
    }
	if(sk != "") {
		av[dbc.SortKey] = &dynamodb.AttributeValue{
			S: aws.String(sk),
		}
	} 
	
	result, err := dbc.DbService.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String(dbc.TableName),
    Key: av,
	})
	if err != nil {
		fmt.Println("NOT FOUND")
	    fmt.Println(err.Error())
	    return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, data)
	if err != nil {
	    panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return err
}


func (dbc DBConfig) FindStartingWith(pk string, value string, data interface{}) error{
	var queryInput = &dynamodb.QueryInput{
		  TableName: aws.String(dbc.TableName),
		  KeyConditions: map[string]*dynamodb.Condition{
		  dbc.PrimaryKey: {
		    ComparisonOperator: aws.String("EQ"),
		    AttributeValueList: []*dynamodb.AttributeValue{
		     {
		      S: aws.String(pk),
		     },
		    },
		   },
		   dbc.SortKey: {
		    ComparisonOperator: aws.String("BEGINS_WITH"),
		    AttributeValueList: []*dynamodb.AttributeValue{
		     {
		      S: aws.String(value),
		     },
		    },
		   },
		  },
		 }

	var result, err = dbc.DbService.Query(queryInput)
	if err != nil {
		fmt.Println("DB:FindStartingWith> NOT FOUND")
	    fmt.Println(err.Error())
	    return  err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
	    panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return err
}



func (dbc DBConfig) FindByGsi(value string, indexName string, indexPk string, data interface{}) error {
	var queryInput = &dynamodb.QueryInput{
		  TableName: aws.String(dbc.TableName),
		  IndexName: aws.String(indexName),
		  KeyConditions: map[string]*dynamodb.Condition{
		   indexPk: {
		    ComparisonOperator: aws.String("EQ"),
		    AttributeValueList: []*dynamodb.AttributeValue{
		     {
		      S: aws.String(value),
		     },
		    },
		   },
		  },
		 }

	var result, err = dbc.DbService.Query(queryInput)
	if err != nil {
		fmt.Println("NOT FOUND")
	    fmt.Println(err.Error())
	    return  err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
	    panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return err
}