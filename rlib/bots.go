package rlib

import "fmt"

// WorkerAsmt et al., are UIDs and Designators for automated processes.
// The negative number space is reserved for automated processes.
const (
	WorkerAsmt        = int64(-1)
	TaskManual        = int64(-2)
	RARBcacheBot      = int64(-3)
	SecDepCacheBot    = int64(-4)
	AcctSliceCacheBot = int64(-5)
	ARSliceCacheBot   = int64(-6)
	TLReportBot       = int64(-7)
	TLInstanceBot     = int64(-8)
	LastBotUID        = int64(-8) // set this to the uid of the last bot
)

// BotRegistryEntry is a struct to associate a bot's id with its name and
// designation
type BotRegistryEntry struct {
	BotID      int64
	Designator string
	Name       string
}

// BotReg is a map indexed by the BotID that provides the BotRegistryEntry
// containing the bot's name and designation. It could not be kept
// in rentroll/worker/ because this info is needed by other libraries that are
// compiled before worker.
//-----------------------------------------------------------------------------
var BotReg = map[int64]BotRegistryEntry{
	WorkerAsmt:        {WorkerAsmt, "AssessmentBot", "Assessment Instance Bot"},
	TaskManual:        {TaskManual, "ManualTaskBot", "Manual Task Bot"},
	RARBcacheBot:      {RARBcacheBot, "RARBcacheBot", "Clean RARBalance Cache"},
	SecDepCacheBot:    {SecDepCacheBot, "SecDepCacheBot", "Clean Security Deposit Cache"},
	AcctSliceCacheBot: {AcctSliceCacheBot, "AcctSliceCacheBot", "Clean Accounts Slice Cache"},
	ARSliceCacheBot:   {ARSliceCacheBot, "ARSliceCacheBot", "Clean Account Rules Slice Cache"},
	TLReportBot:       {TLReportBot, "TLReportBot", "TaskList Report Bot"},
	TLInstanceBot:     {TLInstanceBot, "TLInstanceBot", "TaskList Instance Bot"},
}

// BotName finds and returns the name associated with the bot uid.
//
// INPUTS
//     uid - uid of the bot. These are negative integers
//
// RETURNS
//     the bot entry - it will be an empty string if no match was made on uid
//     any error encountered
//-----------------------------------------------------------------------------
func BotName(uid int64) (string, error) {
	e, err := BotEntry(uid)
	return e.Name, err
}

// BotDesignator finds and returns the designator (i.e., the short name)
// associated with the bot uid.
//
// INPUTS
//     uid - uid of the bot. These are negative integers
//
// RETURNS
//     the bot designator - it will be an empty string if no match was made on uid
//     any error encountered
//-----------------------------------------------------------------------------
func BotDesignator(uid int64) (string, error) {
	e, err := BotEntry(uid)
	return e.Designator, err
}

// BotEntry returns the BotResistryEntry for the bot with the supplied UID or
// an error if it does not find the entry.
//
// INPUTS
//     uid - uid of the bot. These are negative integers
//
// RETURNS
//     the bot entry - it will be an empty struct if no match was made on uid
//     any error encountered
//-----------------------------------------------------------------------------
func BotEntry(uid int64) (BotRegistryEntry, error) {
	for _, v := range BotReg {
		if v.BotID == uid {
			return v, nil
		}
	}
	var e BotRegistryEntry
	return e, fmt.Errorf("No bot registered with UID = %d", uid)
}
