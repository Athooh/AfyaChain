package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AccessLog represents a single access log entry
type AccessLog struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}

// AccessControlContract contract for managing access logs
type AccessControlContract struct {
	contractapi.Contract
	accessLogs []AccessLog
}

func (c *AccessControlContract) LogAccess(ctx contractapi.TransactionContextInterface, userId string, action string) error {
	logID := strconv.Itoa(len(c.accessLogs) + 1) // Generate a new ID based on the current log count

	// Get the transaction timestamp
	timestampProto, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get transaction timestamp: %v", err)
	}

	// Convert the timestamp to a string
	timestamp := timestampProto.AsTime().String() // Convert to Go time.Time and then to string

	logEntry := AccessLog{
		ID:        logID,
		UserID:    userId,
		Action:    action,
		Timestamp: timestamp,
	}

	c.accessLogs = append(c.accessLogs, logEntry)
	return ctx.GetStub().PutState(logID, []byte(logEntry.ToJSON()))
}

// GetAccessLog retrieves an access log by ID
func (c *AccessControlContract) GetAccessLog(ctx contractapi.TransactionContextInterface, id string) (*AccessLog, error) {
	logData, err := ctx.GetStub().GetState(id)
	if err != nil || logData == nil {
		return nil, fmt.Errorf("access log not found")
	}

	var logEntry AccessLog
	if err := logEntry.FromJSON(string(logData)); err != nil {
		return nil, err
	}
	return &logEntry, nil
}

// ToJSON converts an AccessLog to JSON
func (log *AccessLog) ToJSON() string {
	data, _ := json.Marshal(log)
	return string(data)
}

// FromJSON converts JSON to an AccessLog
func (log *AccessLog) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), log)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(AccessControlContract))
	if err != nil {
		fmt.Printf("Error creating access control contract: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting access control contract: %v", err)
	}
}
