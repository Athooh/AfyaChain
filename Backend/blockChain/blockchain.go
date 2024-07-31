package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	// Adjust to your actual project path
)

type Block struct {
	ID           uint      `gorm:"primaryKey"`
	PatientID    int       `json:"patient_id"`
	UserID       int       `json:"user_id"`
	Action       string    `json:"action"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	Hash         string    `json:"hash"`
	Pow          int       `json:"pow"`
}

type BlockData struct {
	PatientID int    `json:"patient_id"`
	UserID    int    `json:"user_id"`
	Action    string `json:"action"` // e.g., "view", "edit"
	Timestamp string `json:"timestamp"`
}

type Blockchain struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	UserID     int     `json:"user_id"`
	Difficulty int     `json:"difficulty"`                        // Add difficulty field
	Chain      []Block `gorm:"foreignKey:PatientID" json:"chain"` // Assuming Block has a UserID field
}

// CalculateHash calculates the hash of the block
func (b Block) CalculateHash() string {
	// Marshal only the relevant block data fields for hashing
	blockData := b.PreviousHash + strconv.Itoa(b.PatientID) + strconv.Itoa(b.UserID) + b.Action + b.Timestamp.String()
	h := sha256.New()
	h.Write([]byte(blockData))
	blockHash := h.Sum(nil)
	return hex.EncodeToString(blockHash)
}

// Mine performs proof-of-work to generate a valid hash
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.CalculateHash()
	}
}

// CreateBlockchain initializes a new blockchain with the genesis block
func CreateBlockchain(patientID int, userID int, action string, difficulty int) Blockchain {
	genesisBlock := Block{
		PatientID:    patientID, // Set a default patient ID for the genesis block
		UserID:       userID,    // Set a default user ID for the genesis block
		Action:       action,    // Indicate this is the genesis block
		Timestamp:    time.Now(),
		PreviousHash: "0",
		Pow:          0, // Proof of work for the genesis block
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()

	return Blockchain{
		Chain:      []Block{genesisBlock},
		Difficulty: difficulty,
	}
}

// AddBlock adds a new block to the blockchain
func (b *Blockchain) AddBlock(bl Block) error {

	lastBlock := b.Chain[len(b.Chain)-1]

	newBlock := Block{
		PatientID:    bl.PatientID,
		UserID:       bl.UserID,
		Action:       bl.Action,
		Timestamp:    time.Now(),
		PreviousHash: lastBlock.Hash,
	}

	// Mine the block to generate a valid hash
	newBlock.Mine(b.Difficulty)
	newBlock.Hash = newBlock.CalculateHash()

	// Add block only if it's valid
	if !b.IsValidNewBlock(newBlock, lastBlock) {
		return errors.New("invalid block")
	}

	b.Chain = append(b.Chain, newBlock)
	return nil
}

// IsValid checks the integrity of the blockchain
func (b Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		previousBlock := b.Chain[i]
		currentBlock := b.Chain[i+1]
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

// isValidNewBlock validates the new block before adding it to the blockchain
func (b Blockchain) IsValidNewBlock(newBlock, previousBlock Block) bool {
	if previousBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if newBlock.Hash != newBlock.CalculateHash() {
		return false
	}
	if !strings.HasPrefix(newBlock.Hash, strings.Repeat("0", b.Difficulty)) {
		return false
	}
	return true
}

// SaveToDatabase saves the blockchain to the database
func (b Blockchain) SaveToDatabase(DB *gorm.DB) error {
	for _, block := range b.Chain {
		dbBlock := Block{
			PatientID:    block.PatientID,
			UserID:       block.UserID,
			Action:       block.Action,
			Timestamp:    block.Timestamp,
			PreviousHash: block.PreviousHash,
			Hash:         block.Hash,
			Pow:          block.Pow,
		}

		result := DB.Create(&dbBlock)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// // LoadFromDatabase loads the blockchain from a database
// func LoadFromDatabase() (Blockchain, error) {
// 	var blockchain Blockchain
// 	var blocks []Block

// 	// Load blocks from the database
// 	result := DB.Find(&blocks)
// 	if result.Error != nil {
// 		return blockchain, result.Error
// 	}

// 	// Rebuild the blockchain from the blocks retrieved
// 	blockchain.Chain = make([]Block, len(blocks))
// 	for i, dbBlock := range blocks {
// 		blockchain.Chain[i] = Block{
// 			PatientID:    dbBlock.PatientID,
// 			UserID:       dbBlock.UserID,
// 			Action:       dbBlock.Action,
// 			Timestamp:    dbBlock.Timestamp,
// 			PreviousHash: dbBlock.PreviousHash,
// 			Hash:         dbBlock.Hash,
// 			Pow:          dbBlock.Pow,
// 		}
// 	}

// 	if len(blockchain.Chain) > 0 {
// 		blockchain.GenesisBlock = blockchain.Chain[0]
// 	}

// 	return blockchain, nil
// }
