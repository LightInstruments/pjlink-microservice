package pjlink

var HumanToRawCommands = map[string]string{
	"power":        "POWR",
	"input-list":   "INST",
	"input":        "INPT",
	"av-mute":      "AVMT",
	"error-status": "ERST",
	"lamp":         "LAMP",
	"name":         "NAME",
	"manufacturer": "INF1",
	"model":        "INF2",
	"version":      "INFO",
}

var RawToHumanCommands = map[string]string{
	"POWR": "power",
	"INST": "input-list",
	"INPT": "input",
	"AVMT": "av-mute",
	"ERST": "error-status",
	"LAMP": "lamp",
	"NAME": "name",
	"INF1": "manufacturer",
	"INF2": "model",
	"INFO": "version",
}

var PowerRequests = map[string]string{
	"query":     "?",
	"power-on":  "1",
	"power-off": "0",
}

var PowerQueryResponses = map[string]string{
	"0":    "power-off (standby)",
	"1":    "power-on (lamp on)",
	"2":    "cooling",
	"3":    "warm-up",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var PowerResponses = map[string]string{
	"OK":   "success, or already current state",
	"ERR2": "out of parameter",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var InputListRequests = map[string]string{
	"query": "?",
}

var InputListQueryResponses = map[string]string{
	// "inputs" to be generated by interpretInputListInputs() function
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var InputRequests = map[string]string{
	"query":    "?",
	"rgb1":     "11",
	"rgb2":     "12",
	"rgb3":     "13",
	"rgb4":     "14",
	"rgb5":     "15",
	"rgb6":     "16",
	"rgb7":     "17",
	"rgb8":     "18",
	"rgb9":     "19",
	"video1":   "21",
	"video2":   "22",
	"video3":   "23",
	"video4":   "24",
	"video5":   "25",
	"video6":   "26",
	"video7":   "27",
	"video8":   "28",
	"video9":   "29",
	"digital1": "31",
	"digital2": "32",
	"digital3": "33",
	"digital4": "34",
	"digital5": "35",
	"digital6": "36",
	"digital7": "37",
	"digital8": "38",
	"digital9": "39",
	"storage1": "41",
	"storage2": "42",
	"storage3": "43",
	"storage4": "44",
	"storage5": "45",
	"storage6": "46",
	"storage7": "47",
	"storage8": "48",
	"storage9": "49",
	"network1": "51",
	"network2": "52",
	"network3": "53",
	"network4": "54",
	"network5": "55",
	"network6": "56",
	"network7": "57",
	"network8": "58",
	"network9": "59",
}

var InputQueryResponses = map[string]string{
	"11":   "rgb1",
	"12":   "rgb2",
	"13":   "rgb3",
	"14":   "rgb4",
	"15":   "rgb5",
	"16":   "rgb6",
	"17":   "rgb7",
	"18":   "rgb8",
	"19":   "rgb9",
	"21":   "video1",
	"22":   "video2",
	"23":   "video3",
	"24":   "video4",
	"25":   "video5",
	"26":   "video6",
	"27":   "video7",
	"28":   "video8",
	"29":   "video9",
	"31":   "digital1",
	"32":   "digital2",
	"33":   "digital3",
	"34":   "digital4",
	"35":   "digital5",
	"36":   "digital6",
	"37":   "digital7",
	"38":   "digital8",
	"39":   "digital9",
	"41":   "storage1",
	"42":   "storage2",
	"43":   "storage3",
	"44":   "storage4",
	"45":   "storage5",
	"46":   "storage6",
	"47":   "storage7",
	"48":   "storage8",
	"49":   "storage9",
	"51":   "network1",
	"52":   "network2",
	"53":   "network3",
	"54":   "network4",
	"55":   "network5",
	"56":   "network6",
	"57":   "network7",
	"58":   "network8",
	"59":   "network9",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var InputResponses = map[string]string{
	"OK":   "OK",
	"ERR2": "nonexistent input source",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var RawToHumanInputs = map[string]string{
	"11": "rgb 1",
	"12": "rgb 2",
	"13": "rgb 3",
	"14": "rgb 4",
	"15": "rgb 5",
	"16": "rgb 6",
	"17": "rgb 7",
	"18": "rgb 8",
	"19": "rgb 9",
	"21": "video 1",
	"22": "video 2",
	"23": "video 3",
	"24": "video 4",
	"25": "video 5",
	"26": "video 6",
	"27": "video 7",
	"28": "video 8",
	"29": "video 9",
	"31": "digital 1",
	"32": "digital 2",
	"33": "digital 3",
	"34": "digital 4",
	"35": "digital 5",
	"36": "digital 6",
	"37": "digital 7",
	"38": "digital 8",
	"39": "digital 9",
	"41": "storage 1",
	"42": "storage 2",
	"43": "storage 3",
	"44": "storage 4",
	"45": "storage 5",
	"46": "storage 6",
	"47": "storage 7",
	"48": "storage 8",
	"49": "storage 9",
	"51": "network 1",
	"52": "network 2",
	"53": "network 3",
	"54": "network 4",
	"55": "network 5",
	"56": "network 6",
	"57": "network 7",
	"58": "network 8",
	"59": "network 9",
}

var AVMuteRequests = map[string]string{
	"query":          "?",
	"video-mute-on":  "11",
	"video-mute-off": "10",
	"audio-mute-on":  "21",
	"audio-mute-off": "20",
	"av-mute-on":     "31",
	"av-mute-off":    "30",
}

var AVMuteQueryResponses = map[string]string{
	"11":   "video mute on, audio mute off",
	"21":   "audio mute on, video mute off",
	"31":   "video and audio mute on",
	"30":   "video and audio mute off",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var AVMuteResponses = map[string]string{
	"OK":   "successful execution, or state already current",
	"ERR2": "out of parameter",
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var ErrorStatusRequests = map[string]string{
	"query": "?",
}

var ErrorStatusQueryResponses = map[string]string{
	//<a><b><c>...<f> - use interpretErrorStatusResponse() function
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var LampRequests = map[string]string{
	"query": "?",
}

var LampQueryResponses = map[string]string{
	//<a> <b> - use interpretLampQueryResponse() function
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var LampStateResponses = map[string]string{
	"0": "off",
	"1": "on",
}

var NameRequests = map[string]string{
	"query": "?",
}

var NameQueryResponses = map[string]string{
	//<hostname> - just pass through
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var ManufacturerRequests = map[string]string{
	"query": "?",
}

var ManufacturerQueryResponses = map[string]string{
	//<manufacturer> - just pass through
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var ModelRequests = map[string]string{
	"query": "?",
}

var ModelQueryResponses = map[string]string{
	//<model> - just pass through
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}

var VersionRequests = map[string]string{
	"query": "?",
}

var VersionQueryResponses = map[string]string{
	//<version> - just pass through
	"ERR3": "unavailable time",
	"ERR4": "device failure",
}
