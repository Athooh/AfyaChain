package database

import (
	"database/sql"
	"fmt"

	"github.com/Athooh/HealthChain/models"
	_ "github.com/lib/pq"
)

// CreatePatient inserts a new patient into the database
func CreatePatient(db *sql.DB, patient *models.Patient) error {
	query := `INSERT INTO patients (first_name, last_name, dob, gender, email, phone, address) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(query, patient.FirstName, patient.LastName, patient.DOB, patient.Gender,
		patient.Email, patient.Phone, patient.Address).Scan(&patient.ID)
	return err
}

// GetPatient retrieves a patient by ID
func GetPatient(db *sql.DB, id int) (*models.Patient, error) {
	query := `SELECT id, first_name, last_name, dob, gender, email, phone, address, 
                     created_at, updated_at FROM patients WHERE id = $1`
	patient := &models.Patient{}
	err := db.QueryRow(query, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName,
		&patient.DOB, &patient.Gender, &patient.Email,
		&patient.Phone, &patient.Address, &patient.CreatedAt,
		&patient.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Retrieve medical records for the patient
	patient.MedicalRecords, err = GetMedicalRecordsByPatientID(db, patient.ID)
	return patient, err
}

// UpdatePatient updates an existing patient in the database
func UpdatePatient(db *sql.DB, patient *models.Patient) error {
	query := `UPDATE patients SET first_name = $1, last_name = $2, dob = $3, 
              gender = $4, email = $5, phone = $6, address = $7, 
              updated_at = CURRENT_TIMESTAMP WHERE id = $8`
	_, err := db.Exec(query, patient.FirstName, patient.LastName, patient.DOB, patient.Gender,
		patient.Email, patient.Phone, patient.Address, patient.ID)
	return err
}

// DeletePatient removes a patient from the database
func DeletePatient(db *sql.DB, id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

// CreateMedicalRecord inserts a new medical record into the database
func CreateMedicalRecord(db *sql.DB, record *models.MedicalRecord) error {
	query := `INSERT INTO medical_records (patient_id, record_date, condition, treatment, notes) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, record.PatientID, record.RecordDate, record.Condition,
		record.Treatment, record.Notes).Scan(&record.ID)
	return err
}

// GetMedicalRecordsByPatientID retrieves all medical records for a specific patient
func GetMedicalRecordsByPatientID(db *sql.DB, patientID int) ([]models.MedicalRecord, error) {
	query := `SELECT id, patient_id, record_date, condition, treatment, notes 
              FROM medical_records WHERE patient_id = $1`
	rows, err := db.Query(query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.MedicalRecord
	for rows.Next() {
		var record models.MedicalRecord
		if err := rows.Scan(&record.ID, &record.PatientID, &record.RecordDate,
			&record.Condition, &record.Treatment, &record.Notes); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

// UpdateMedicalRecord updates an existing medical record in the database
func UpdateMedicalRecord(db *sql.DB, record *models.MedicalRecord) error {
	query := `UPDATE medical_records SET record_date = $1, condition = $2, 
              treatment = $3, notes = $4 WHERE id = $5`
	_, err := db.Exec(query, record.RecordDate, record.Condition, record.Treatment,
		record.Notes, record.ID)
	return err
}

// DeleteMedicalRecord removes a medical record from the database
func DeleteMedicalRecord(db *sql.DB, id int) error {
	query := `DELETE FROM medical_records WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

// OpenDatabase opens a connection to the PostgreSQL database and creates tables if they don't exist
func OpenDatabase(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Create tables if they don't exist
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// createTables creates the necessary tables in the database if they don't exist
func createTables(db *sql.DB) error {
	// Create patients table
	createPatientsTable := `
    CREATE TABLE IF NOT EXISTS patients (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(50) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        dob DATE NOT NULL,
        gender VARCHAR(10),
        email VARCHAR(100) UNIQUE,
        phone VARCHAR(20),
        address TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(createPatientsTable)
	if err != nil {
		return fmt.Errorf("failed to create patients table: %v", err)
	}

	// Create trigger function for updating the updated_at column
	// createTriggerFunction := `
	// CREATE OR REPLACE FUNCTION update_updated_at_column()
	// RETURNS TRIGGER AS $$
	// BEGIN
	//     NEW.updated_at = CURRENT_TIMESTAMP;
	//     RETURN NEW;
	// END;
	// $$ LANGUAGE plpgsql;`

	// _, err = db.Exec(createTriggerFunction)
	// if err != nil {
	// 	return fmt.Errorf("failed to create trigger function: %v", err)
	// }

	// // Create trigger for the patients table
	// createTrigger := `
	// CREATE TRIGGER update_patient_updated_at
	// BEFORE UPDATE ON patients
	// FOR EACH ROW
	// EXECUTE FUNCTION update_updated_at_column();`

	// _, err = db.Exec(createTrigger)
	// if err != nil {
	// 	return fmt.Errorf("failed to create trigger for patients table: %v", err)
	// }

	// Create medical_records table
	createMedicalRecordsTable := `
    CREATE TABLE IF NOT EXISTS medical_records (
        id SERIAL PRIMARY KEY,
        patient_id INT NOT NULL,
        record_date DATE NOT NULL,
        condition VARCHAR(100),
        treatment VARCHAR(100),
        notes TEXT,
        FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
    );`

	_, err = db.Exec(createMedicalRecordsTable)
	if err != nil {
		return fmt.Errorf("failed to create medical_records table: %v", err)
	}

	return nil
}
