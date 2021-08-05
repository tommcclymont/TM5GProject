package models

type Arp struct {
	PriorityLevel uint
	PreemptCap    string
	PreemptVuln   string
}

var ArpData = Arp{
	PriorityLevel: 10,
	PreemptCap:    "NOT_PREEMPT",
	PreemptVuln:   "PREEMPTABLE",
}
