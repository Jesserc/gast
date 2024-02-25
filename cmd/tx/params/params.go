package params

var (
	RawTx    string
	PrivKey  string
	TxData   string
	To       string
	From     string
	TxRpcUrl string
	TxHash   string
	GasLimit uint64
	Nonce    uint64
	Wei      uint64
)

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
