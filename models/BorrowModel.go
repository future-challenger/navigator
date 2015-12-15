package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

const (
	borrowCollectionName = "rule"
)

type BorrowModel struct {
	BorrowedByWho string `json: "borrowedByWho"`
	BorrowedBook  string `json: "borrowedBook"`
	BorrowedAt    string `json: "borrowedAt"`
	BorrowStatus  string `json: "borrowStatus"` // 0: borrowed, 1: returned, 2: over due, 3: available
}
