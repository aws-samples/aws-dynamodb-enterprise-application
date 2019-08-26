// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package core

import (
	"hash/fnv"
	"strconv"
	"time"
)

func Hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}


func Chunk(array []interface{}, chunkSize int) [][]interface{}{
	var divided [][]interface{}
	for i := 0; i < len(array); i += chunkSize {
	    end := i + chunkSize
	    if end > len(array) {
	        end = len(array)
	    }
	    divided = append(divided, array[i:end])
	}
	return divided
}


//Generats a usinque ID based timestamp
func GeneratUniqueId() string {
	return strconv.FormatInt(int64(Hash(time.Now().Format("20060102150405"))),10)
}