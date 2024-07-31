package database

import (
	"log"
	"time"

	blockchain "github.com/Athooh/HealthChain/Backend/blockChain"
	"github.com/Athooh/HealthChain/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() (db *gorm.DB, err error) {
	// Database connection
	dsn := "root:12345678@tcp(34.28.233.25:3306)/ehr_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func CreatePatient(firstName, lastName string, gender, email, phone, address string) *models.Patient {
	patient := &models.Patient{
		FirstName: firstName,
		LastName:  lastName,
		Gender:    gender,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	result := db.Create(patient)
	if result.Error != nil {
		log.Printf("Error creating patient: %v", result.Error)
		return nil
	}
	return patient
}

func GetPatient(id int) *models.Patient {
	var patient models.Patient
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	result := db.Preload("MedicalRecords").First(&patient, id)
	if result.Error != nil {
		log.Printf("Error retrieving patient: %v", result.Error)
		return nil
	}
	return &patient
}

func UpdatePatient(id int, firstName, lastName string, gender, email, phone, address string) *models.Patient {
	patient := GetPatient(id)
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	if patient == nil {
		return nil
	}
	patient.FirstName = firstName
	patient.LastName = lastName
	patient.Gender = gender
	patient.Email = email
	patient.Phone = phone
	patient.Address = address
	result := db.Save(patient)
	if result.Error != nil {
		log.Printf("Error updating patient: %v", result.Error)
		return nil
	}
	var userBlockchain blockchain.Blockchain
	lastBlock := userBlockchain.Chain[len(userBlockchain.Chain)-1]
	newBlock := blockchain.Block{
		PatientID:    id,
		UserID:       123,
		Action:       "Update patient information (patient id) by dr(drId)",
		Timestamp:    time.Now(),
		PreviousHash: lastBlock.Hash,
	}

	// Mine the block and calculate its hash
	newBlock.Mine(userBlockchain.Difficulty)
	newBlock.Hash = newBlock.CalculateHash()
	// Validate the new block before adding
	if !userBlockchain.IsValidNewBlock(newBlock, lastBlock) {
		log.Fatal("invalid block")
	}
	userBlockchain.Chain = append(userBlockchain.Chain, newBlock)

	// Save the updated blockchain to the database
	if err := db.Save(&userBlockchain).Error; err != nil {
		log.Fatal(err)
	}
	return patient
}

func DeletePatient(id int) bool {
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return false
	}
	result := db.Delete(&models.Patient{}, id)
	if result.Error != nil {
		log.Printf("Error deleting patient: %v", result.Error)
		return false
	}
	return true
}

// CRUD functions for MedicalRecord

func CreateMedicalRecord(patientID int, recordDate time.Time, condition, treatment, notes string) *models.MedicalRecord {

	medicalRecord := &models.MedicalRecord{
		PatientID:  patientID,
		RecordDate: recordDate,
		Condition:  condition,
		Treatment:  treatment,
		Notes:      notes,
	}
	result := db.Create(medicalRecord)
	if result.Error != nil {
		log.Printf("Error creating medical record: %v", result.Error)
		return nil
	}

	return medicalRecord
}

func GetMedicalRecord(id int) *models.MedicalRecord {
	var medicalRecord models.MedicalRecord
	result := db.First(&medicalRecord, id)
	if result.Error != nil {
		log.Printf("Error retrieving medical record: %v", result.Error)
		return nil
	}
	return &medicalRecord
}

func UpdateMedicalRecord(id, patient_id int, recordDate time.Time, condition, treatment, notes string) *models.MedicalRecord {
	medicalRecord := GetMedicalRecord(id)
	if medicalRecord == nil {
		return nil
	}
	medicalRecord.RecordDate = recordDate
	medicalRecord.Condition = condition
	medicalRecord.Treatment = treatment
	medicalRecord.Notes = notes
	result := db.Save(medicalRecord)
	if result.Error != nil {
		log.Printf("Error updating medical record: %v", result.Error)
		return nil
	}
	var userBlockchain blockchain.Blockchain
	lastBlock := userBlockchain.Chain[len(userBlockchain.Chain)-1]
	newBlock := blockchain.Block{
		PatientID:    patient_id,
		UserID:       123,
		Action:       "Update on patien (patient id) by dr(drId)",
		Timestamp:    time.Now(),
		PreviousHash: lastBlock.Hash,
	}

	// Mine the block and calculate its hash
	newBlock.Mine(userBlockchain.Difficulty)
	newBlock.Hash = newBlock.CalculateHash()
	// Validate the new block before adding
	if !userBlockchain.IsValidNewBlock(newBlock, lastBlock) {
		log.Fatal("invalid block")
	}
	userBlockchain.Chain = append(userBlockchain.Chain, newBlock)

	// Save the updated blockchain to the database
	if err := db.Save(&userBlockchain).Error; err != nil {
		log.Fatal(err)
	}
	return medicalRecord
}

func DeleteMedicalRecord(id int) bool {
	result := db.Delete(&models.MedicalRecord{}, id)
	if result.Error != nil {
		log.Printf("Error deleting medical record: %v", result.Error)
		return false
	}
	var userBlockchain blockchain.Blockchain
	lastBlock := userBlockchain.Chain[len(userBlockchain.Chain)-1]
	newBlock := blockchain.Block{
		PatientID:    id,
		UserID:       123,
		Action:       "delete medical record on patient (patient id) by dr(drId)",
		Timestamp:    time.Now(),
		PreviousHash: lastBlock.Hash,
	}

	// Mine the block and calculate its hash
	newBlock.Mine(userBlockchain.Difficulty)
	newBlock.Hash = newBlock.CalculateHash()
	// Validate the new block before adding
	if !userBlockchain.IsValidNewBlock(newBlock, lastBlock) {
		log.Fatal("invalid block")
	}
	userBlockchain.Chain = append(userBlockchain.Chain, newBlock)

	// Save the updated blockchain to the database
	if err := db.Save(&userBlockchain).Error; err != nil {
		log.Fatal(err)
	}
	return true
}
