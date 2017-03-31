package main

import (
	"fmt"
	"rentroll/rlib"
	"rentroll/ws"
)

// WSTypeFactory is a map for creating new data types used by the
// web services routines based on the supplied name.
var WSTypeFactory = map[string]Creator{
	"ColSort":                      NewColSort,
	"GenSearch":                    NewGenSearch,
	"GetRentableResponse":          NewGetRentableResponse,
	"GLAccount":                    NewGLAccount,
	"PrRentableOther":              NewPrRentableOther,
	"RAPeople":                     NewRAPeople,
	"RAPeopleResponse":             NewRAPeopleResponse,
	"RAPets":                       NewRAPets,
	"RAR":                          NewWSRAR,
	"RentalAgr":                    NewRentalAgr,
	"RentalAgreementPet":           NewRentalAgreementPet,
	"RentalAgrSearchResponse":      NewRentalAgrSearchResponse,
	"SearchGLAccountsResponse":     NewSearchGLAccountsResponse,
	"SearchRentablesResponse":      NewSearchRentablesResponse,
	"SearchTransactantsResponse":   NewSearchTransactantsResponse,
	"SearchReceiptsResponse":       NewSearchReceiptsResponse,
	"PrReceiptGrid":                NewPrReceiptGrid,
	"SvcStatusResponse":            NewSvcStatusResponse,
	"WebGridSearchRequest":         NewWebGridSearchRequest,
	"GetRentalAgreementResponse":   NewGetRentalAgreementResponse,
	"SearchAssessmentsResponse":    NewSearchAssessmentsResponse,
	"AssessmentGrid":               NewAssessmentGrid,
	"GetAssessmentResponse":        NewGetAssessmentResponse,
	"AssessmentSendForm":           NewAssessmentSendForm,
	"SaveAssessmentInput":          NewSaveAssessmentInput,
	"GetReceiptResponse":           NewGetReceiptResponse,
	"ReceiptSendForm":              NewReceiptSendForm,
	"GetTransactantResponse":       NewGetTransactantResponse,
	"SaveReceiptInput":             NewSaveReceiptInput,
	"AssessmentSaveForm":           NewAssessmentSaveForm,
	"TransactantsTypedownResponse": NewTransactantsTypedownResponse,
	"TransactantTypeDown":          NewTransactantTypeDown,
	"WebTypeDownRequest":           NewWebTypeDownRequest,
	"PaymentTypeGetResponse":       NewPaymentTypeGetResponse,
	"PaymentTypeGrid":              NewPaymentTypeGrid,
	"DepositoryGrid":               NewDepositoryGrid,
	"DepositorySearchResponse":     NewDepositorySearchResponse,
	"DepositoryGridSave":           NewDepositoryGridSave,
	"DepositoryGetResponse":        NewDepositoryGetResponse,
	"WebGridDelete":                NewWebGridDelete,
}

// NewWebGridDelete is a factory for WebGridDelete structs
func NewWebGridDelete() interface{} {
	return new(ws.WebGridDelete)
}

// NewDepositoryGetResponse is a factory for DepositoryGetResponse structs
func NewDepositoryGetResponse() interface{} {
	return new(ws.DepositoryGetResponse)
}

// NewDepositoryGridSave is a factory for DepositoryGridSave structs
func NewDepositoryGridSave() interface{} {
	return new(ws.DepositoryGridSave)
}

// NewDepositoryGrid is a factory for DepositoryGrid structs
func NewDepositoryGrid() interface{} {
	return new(ws.DepositoryGrid)
}

// NewDepositorySearchResponse is a factory for DepositorySearchResponse structs
func NewDepositorySearchResponse() interface{} {
	return new(ws.DepositorySearchResponse)
}

// NewPaymentTypeGrid is a factory for PaymentTypeGrid structs
func NewPaymentTypeGrid() interface{} {
	return new(ws.PaymentTypeGrid)
}

// NewPaymentTypeGetResponse is a factory for PaymentTypeGetResponse structs
func NewPaymentTypeGetResponse() interface{} {
	return new(ws.PaymentTypeGetResponse)
}

// NewWebTypeDownRequest is a factory for WebTypeDownRequest structs
func NewWebTypeDownRequest() interface{} {
	return new(ws.WebTypeDownRequest)
}

// NewTransactantTypeDown is a factory for TransactantTypeDown structs
func NewTransactantTypeDown() interface{} {
	return new(rlib.TransactantTypeDown)
}

// NewTransactantsTypedownResponse is a factory for TransactantsTypedownResponse structs
func NewTransactantsTypedownResponse() interface{} {
	return new(ws.TransactantsTypedownResponse)
}

// NewAssessmentSaveForm is a factory for AssessmentSaveForm structs
func NewAssessmentSaveForm() interface{} {
	return new(ws.AssessmentSaveForm)
}

// NewSaveReceiptInput is a factory for SaveReceiptInput structs
func NewSaveReceiptInput() interface{} {
	return new(ws.SaveReceiptInput)
}

// NewGetTransactantResponse is a factory for GetTransactantResponse structs
func NewGetTransactantResponse() interface{} {
	return new(ws.GetTransactantResponse)
}

// NewReceiptSendForm is a factory for ReceiptSendForm structs
func NewReceiptSendForm() interface{} {
	return new(ws.ReceiptSendForm)
}

// NewGetReceiptResponse is a factory for GetReceiptResponse structs
func NewGetReceiptResponse() interface{} {
	return new(ws.GetReceiptResponse)
}

// NewSaveAssessmentInput is a factory for SaveAssessmentInput structs
func NewSaveAssessmentInput() interface{} {
	return new(ws.SaveAssessmentInput)
}

// NewAssessmentSendForm is a factory for AssessmentSendForm structs
func NewAssessmentSendForm() interface{} {
	return new(ws.AssessmentSendForm)
}

// NewGetAssessmentResponse is a factory for GetAssessmentResponse structs
func NewGetAssessmentResponse() interface{} {
	return new(ws.GetAssessmentResponse)
}

// NewAssessmentGrid is a factory for AssessmentGrid structs
func NewAssessmentGrid() interface{} {
	return new(ws.AssessmentGrid)
}

// NewSearchAssessmentsResponse is a factory for SearchAssessmentsResponse structs
func NewSearchAssessmentsResponse() interface{} {
	return new(ws.SearchAssessmentsResponse)
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

// NewWebGridSearchRequest is a factory for WebGridSearchRequest structs
func NewWebGridSearchRequest() interface{} {
	return new(ws.WebGridSearchRequest)
}

// NewWSRAR is a factory for RAR structs
func NewWSRAR() interface{} {
	return new(ws.RAR)
}

// NewGLAccount is a factory for GLAccount structs
func NewGLAccount() interface{} {
	return new(rlib.GLAccount)
}
