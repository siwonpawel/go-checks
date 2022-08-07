package main

type ValidationRequest struct {
	Value string `json:"value,omitempty"`
}

type ValidationResult struct {
	Country      string `json:"country"`
	Identifier   string `json:"identifier"`
	IsValid      bool   `json:"isValid"`
	Warnings     []string
	Error        string
	CleanedInput string
}

type ValidationResults []ValidationResult
