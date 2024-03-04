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
	SigMsgValue     string
	SigAddressValue string

	DirValue string
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

// NetworkExplorers maps network IDs to their respective explorers.
var NetworkExplorers = map[uint64]string{
	0x01:     "https://etherscan.io/",                          // Ethereum Mainnet
	0x05:     "https://goerli.etherscan.io/",                   // Goerli Testnet
	0xAA36A7: "https://sepolia.etherscan.io/",                  // Sepolia Testnet
	0x89:     "https://polygonscan.com/",                       // Polygon Mainnet
	0x13881:  "https://mumbai.polygonscan.com/",                // Polygon Mumbai Testnet
	0x0A:     "https://optimistic.etherscan.io/",               // Optimism Mainnet
	0x1A4:    "https://goerli-optimism.etherscan.io/",          // Optimism Goerli Testnet
	0xA4B1:   "https://arbiscan.io/",                           // Arbitrum One Mainnet
	0x66EEE:  "https://sepolia.arbiscan.io/",                   // Arbitrum Sepolia Testnet
	0x38:     "https://bscscan.com/",                           // Binance Smart Chain Mainnet
	0x61:     "https://testnet.bscscan.com/",                   // Binance Smart Chain Testnet
	0x421611: "https://explorer.celo.org/",                     // Celo Mainnet
	0xA4EC:   "https://alfajores-blockscout.celo-testnet.org/", // Celo Alfajores Testnet
	0x2105:   "https://basescan.org/",                          // Base Mainnet
	0xE708:   "https://lineascan.build/",                       // Linea Mainnet
	0x144:    "https://explorer.zksync.io/",                    // zkSync Mainnet
}
