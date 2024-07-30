package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data         BlockData `json:"data"`
	Hash         string    `json:"hash"`
	PreviousHash string    `json:"previous_hash"`
	Timestamp    time.Time `json:"timestamp"`
	Pow          int       `json:"pow"`
}

type Blockchain struct {
	GenesisBlock Block   `json:"genesis_block"`
	Chain        []Block `json:"chain"`
	Difficulty   int     `json:"difficulty"`
}
type BlockData struct {
	PatientID int    `json:"patient_id"`
	UserID    int    `json:"user_id"`
	Action    string `json:"action"` // e.g., "view", "edit"
	Timestamp string `json:"timestamp"`
}

func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Pow)
	h := sha256.New()
	h.Write([]byte(blockData))
	blockHash := h.Sum(nil)
	return hex.EncodeToString(blockHash)
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.CalculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}

	return Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []Block{genesisBlock},
		Difficulty:   difficulty,
	}
}

func (b *Blockchain) AddBlock(patientID int, userID int, action string) {
	blockData := BlockData{
		PatientID: patientID,
		UserID:    userID,
		Action:    action,
		Timestamp: time.Now().String(),
	}
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := Block{
		Data:         blockData,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}
	newBlock.Mine(b.Difficulty)
	b.Chain = append(b.Chain, newBlock)
}

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
