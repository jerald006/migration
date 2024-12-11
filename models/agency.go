package models

type MongoAgency struct {
	_id string `json:"name"`
}

type CockroachDBAgency struct {
	_id string `json:"name"`
}

type CockroachDBAgencyPolicy struct {
	_id string `json:"name"`
}
