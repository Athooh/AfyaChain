package models

import "time"

type MedicalRecord struct {
	ID         int       `json:"id"`
	PatientID  int       `json:"patient_id"` // Foreign key reference
	RecordDate time.Time `json:"record_date"`
	Condition  string    `json:"condition"`
	Treatment  string    `json:"treatment"`
	Notes      string    `json:"notes"`
}

type Patient struct {
	ID             int             `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	DOB            time.Time       `json:"dob"`
	Gender         string          `json:"gender"`
	Email          string          `json:"email"`
	Phone          string          `json:"phone"`
	Address        string          `json:"address"`
	MedicalRecords []MedicalRecord `json:"medical_records"` // Slice of medical records
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
