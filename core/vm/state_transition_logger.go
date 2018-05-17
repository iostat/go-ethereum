package vm

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// StateTransitionLogger melds a StructLogger to be used within
// StateProcessor.ApplyTransition, allowing
// us to have a super initimate view of a tx executing
type StateTransitionLogger struct {
	*StructLogger
	*LogConfig
}

// NewStateTransitionLogger creates a StateTransitionLogger to be
// used for a single state transition
func NewStateTransitionLogger(cfg *LogConfig) *StateTransitionLogger {
	return &StateTransitionLogger{
		LogConfig:    cfg,
		StructLogger: NewStructLogger(cfg),
	}
}

// CaptureStart allows StateTransitionLogger to implement the Tracer interface
func (stl *StateTransitionLogger) CaptureStart(from common.Address, to common.Address, call bool, input []byte, gas uint64, value *big.Int) error {
	return stl.StructLogger.CaptureStart(from, to, call, input, gas, value)
}

// CaptureState allows StateTransitionLogger to implement the Tracer interface
func (stl *StateTransitionLogger) CaptureState(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, contract *Contract, depth int, err error) error {
	return stl.StructLogger.CaptureState(env, pc, op, gas, cost, memory, stack, contract, depth, err)
}

// CaptureFault allows StateTransitionLogger to implement the Tracer interface
func (stl *StateTransitionLogger) CaptureFault(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, contract *Contract, depth int, err error) error {
	return stl.StructLogger.CaptureFault(env, pc, op, gas, cost, memory, stack, contract, depth, err)
}

// CaptureEnd allows StateTransitionLogger to implement the Tracer interface
func (stl *StateTransitionLogger) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	return stl.StructLogger.CaptureEnd(output, gasUsed, t, err)
}

// GetCapturedStorageChanges allows us to extract the captured storage trie changes.
func (stl *StateTransitionLogger) GetCapturedStorageChanges() map[common.Address](map[common.Hash]common.Hash) {
	aliasedMap := make(map[common.Address](map[common.Hash]common.Hash))
	changedValues := stl.StructLogger.changedValues
	for k, v := range changedValues {
		aliasedMap[k] = map[common.Hash]common.Hash(v)
	}
	return aliasedMap
}
