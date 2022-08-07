package main

import (
	"encoding/json"
	"net/http"

	"github.com/siwonpawel/go-stdnum/stdnum/validation"
)

func validate(w http.ResponseWriter, r *http.Request) {

	var request ValidationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return
	}

	results := validation.ValidateAll(request.Value)

	vr := make(ValidationResults, 0)
	for _, v := range results {
		vr = append(vr, ValidationResult{
			Country:      v.Country,
			Identifier:   v.IdentifierName,
			IsValid:      v.IsValid,
			Warnings:     v.Warnings,
			Error:        v.Error.Error(),
			CleanedInput: v.DebugInfo.CleanedInput,
		})
	}

	err = json.NewEncoder(w).Encode(vr)
	if err != nil {
		return
	}
}
