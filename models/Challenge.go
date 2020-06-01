package models

import (
	"../websiteData"
)

type Challenge struct {
	Link    string
	Name    string
	Website websiteData.Website
}
