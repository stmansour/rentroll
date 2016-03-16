package main

func deleteJournalAllocations(jid int) {
	_, err := App.prepstmt.deleteJournalAllocations.Exec(jid)
	if err != nil {
		ulog("Error deleting journal allocations for JID = %d, error: %v\n", jid, err)
	}
}

func deleteJournalEntry(jid int) {
	_, err := App.prepstmt.deleteJournalEntry.Exec(jid)
	if err != nil {
		ulog("Error deleting journal entry for JID = %d, error: %v\n", jid, err)
	}
}

func deleteJournalMarker(jmid int) {
	_, err := App.prepstmt.deleteJournalMarker.Exec(jmid)
	if err != nil {
		ulog("Error deleting journal marker for JID = %d, error: %v\n", jmid, err)
	}
}
