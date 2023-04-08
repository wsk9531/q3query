package q3query

// TODO: Externalize Configuration
//
// CONSTANT:
// Hardcoded list of master servers to query
func MASTER_SERVERS() []string {
	return []string{
		"master.quake3arena.com:27950",
		"master.ioquake3.org:27950",
		"master.maverickservers.com:27950",
		"master.quakeservers.net:27950",
		"master.qtracker.com:27900",
	}
}

// Quake 3 Game Protocol
const PROTOCOL = "68"

// OOB - Out Of Bounds Packet Header
const OOB = "\xFF\xFF\xFF\xFF"

const SEP = '\u005c'

const MSG_GETSERVERS = OOB + "getservers " + PROTOCOL + " empty full\n"
const MSG_EOT = "\\EOT\x00\x00\x00"
