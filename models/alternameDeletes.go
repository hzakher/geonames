package models

type AlternameDeletes struct {
	Id        int    `csv:"alternateNameId"`
	GeonameId int    `csv:"geonameId"`
	Name      string `csv:"name"`
	Comment   string `csv:"comment"`
}
