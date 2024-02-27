package gastParams

var (
	FromValue     string // used for transaction(s)
	ToValue       string
	WeiValue      uint64
	TxDataValue   string
	NonceValue    uint64
	PrivKeyValue  string
	TxRpcUrlValue string

	GasLimitValue uint64 // GasLimitValue and RawTxValue used fo raw sig
	RawTxValue    string

	TxHashValue string // used for transaction tracing

	SigValue        string
	SigHashValue    string
	SigAddressValue string
)

// Terminal outputs colours
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)
