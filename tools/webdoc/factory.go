package main

import (
	"fmt"
	"rentroll/rlib"
	"rentroll/ws"
)

// WSTypeFactory is a map for creating new data types used by the
// web services routines based on the supplied name.
var WSTypeFactory = map[string]Creator{
	"ColSort":                    NewColSort,
	"GenSearch":                  NewGenSearch,
	"GetRentableResponse":        NewGetRentableResponse,
	"GLAccount":                  NewGLAccount,
	"PrRentableOther":            NewPrRentableOther,
	"RAPeople":                   NewRAPeople,
	"RAPeopleResponse":           NewRAPeopleResponse,
	"RAPets":                     NewRAPets,
	"RAR":                        NewWSRAR,
	"RentalAgr":                  NewRentalAgr,
	"RentalAgreementPet":         NewRentalAgreementPet,
	"RentalAgrSearchResponse":    NewRentalAgrSearchResponse,
	"SearchGLAccountsResponse":   NewSearchGLAccountsResponse,
	"SearchRentablesResponse":    NewSearchRentablesResponse,
	"SearchTransactantsResponse": NewSearchTransactantsResponse,
	"SearchReceiptsResponse":     NewSearchReceiptsResponse,
	"PrReceiptGrid":              NewPrReceiptGrid,
	"SvcStatusResponse":          NewSvcStatusResponse,
	"WebRequest":                 NewWebRequest,
	"GetRentalAgreementResponse": NewGetRentalAgreementResponse,
}

// FactoryNew looks for type t in WSTypeFactory. If it is found,
// it returns a new data object of the type requested. If not
// it returns an error
func FactoryNew(t string) (interface{}, error) {
	fact, ok := WSTypeFactory[t]
	if !ok {
		return nil, fmt.Errorf("**** ERROR **** unrecognized factory type = %s", t)
	}
	return fact(), nil
}

// NewPrReceiptGrid is a factory for PrReceiptGrid structs
func NewPrReceiptGrid() interface{} {
	return new(ws.PrReceiptGrid)
}

// NewSearchReceiptsResponse is a factory for SearchReceiptsResponse structs
func NewSearchReceiptsResponse() interface{} {
	return new(ws.SearchReceiptsResponse)
}

// NewGetRentableResponse is a factory for GetRentableResponse structs
func NewGetRentableResponse() interface{} {
	return new(ws.GetRentableResponse)
}

// NewRAPets is a factory for RAPets structs
func NewRAPets() interface{} {
	return new(ws.RAPets)
}

// NewRAPeopleResponse is a factory for RAPeopleResponse structs
func NewRAPeopleResponse() interface{} {
	return new(ws.RAPeopleResponse)
}

// NewSearchTransactantsResponse is a factory for SearchTransactantsResponse structs
func NewSearchTransactantsResponse() interface{} {
	return new(ws.SearchTransactantsResponse)
}

// NewSearchGLAccountsResponse is a factory for SearchGLAccountsResponse structs
func NewSearchGLAccountsResponse() interface{} {
	return new(ws.SearchGLAccountsResponse)
}

// NewSvcStatusResponse is a factory for SvcStatusResponse structs
func NewSvcStatusResponse() interface{} {
	return new(ws.SvcStatusResponse)
}

// NewGetRentalAgreementResponse is a factory for GetRentalAgreementResponse structs
func NewGetRentalAgreementResponse() interface{} {
	return new(ws.GetRentalAgreementResponse)
}

// NewRentalAgrSearchResponse is a factory for RentalAgrSearchResponse structs
func NewRentalAgrSearchResponse() interface{} {
	return new(ws.RentalAgrSearchResponse)
}

// NewColSort is a factory for ColSort structs
func NewColSort() interface{} {
	return new(ws.ColSort)
}

// NewGenSearch is a factory for GenSearch structs
func NewGenSearch() interface{} {
	return new(ws.GenSearch)
}

// NewRentalAgr is a factory for RentalAgr structs
func NewRentalAgr() interface{} {
	return new(ws.RentalAgr)
}

// NewPrRentableOther is a factory for PrRentableOther structs
func NewPrRentableOther() interface{} {
	return new(ws.PrRentableOther)
}

// NewSearchRentablesResponse is a factory for SearchRentablesResponse structs
func NewSearchRentablesResponse() interface{} {
	return new(ws.SearchRentablesResponse)
}

// NewRentalAgreementPet is a factory for RentalAgreementPet structs
func NewRentalAgreementPet() interface{} {
	return new(rlib.RentalAgreementPet)
}

// NewWSPets is a factory for RAPets structs
func NewWSPets() interface{} {
	return new(ws.RAPets)
}

// NewRAPeople is a factory for RAPeople structs
func NewRAPeople() interface{} {
	return new(ws.RAPeople)
}

// NewWebRequest is a factory for WebRequest structs
func NewWebRequest() interface{} {
	return new(ws.WebRequest)
}

// NewWSRAR is a factory for RAR structs
func NewWSRAR() interface{} {
	return new(ws.RAR)
}

// NewGLAccount is a factory for GLAccount structs
func NewGLAccount() interface{} {
	return new(rlib.GLAccount)
}
