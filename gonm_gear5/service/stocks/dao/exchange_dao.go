package dao

type Exchange struct {
	ExchangeCode string `datastore:"-"`
	Name         string
	Mic          string
}
