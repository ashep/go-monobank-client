package monobank

type CurrencyRate struct {
	CurrencyA int     `json:"currencyCodeA"`
	CurrencyB int     `json:"currencyCodeB"`
	Date      int64   `json:"date"`
	Sell      float64 `json:"rateSell"`
	Buy       float64 `json:"rateBuy"`
	Cross     float64 `json:"rateCross"`
}

type Account struct {
	ID           string   `json:"id"`
	SendID       string   `json:"sendId"`
	CurrencyCode int      `json:"currencyCode"`
	CashbackType string   `json:"cashbackType"`
	Balance      int64    `json:"balance"`
	CreditLimit  int64    `json:"creditLimit"`
	IBAN         string   `json:"iban"`
	MaskedPAN    []string `json:"maskedPan"`
	Type         string   `json:"type"`
}

type JarInfo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Currency    int    `json:"currencyCode"`
	Balance     int64  `json:"balance"`
	Goal        int64  `json:"goal"`
}

type Customer struct {
	ID          string    `json:"clientId"`
	Name        string    `json:"name"`
	WebHookURL  string    `json:"webHookUrl"`
	Accounts    []Account `json:"accounts"`
	Jars        []JarInfo `json:"jars"`
	Permissions string    `json:"permissions"`
}

type StatementItem struct {
	ID              string `json:"id"`
	Time            int64  `json:"time"`
	Description     string `json:"description"`
	MCC             int    `json:"mcc"`
	OriginalMCC     int    `json:"originalMcc"`
	Hold            bool   `json:"hold"`
	Amount          int64  `json:"amount"`
	OperationAmount int64  `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int64  `json:"commissionRate"`
	CashbackAmount  int64  `json:"cashbackAmount"`
	Balance         int64  `json:"balance"`
	Comment         string `json:"comment"`
	ReceiptID       string `json:"receiptId"`
	InvoiceID       string `json:"invoiceId"`
	CounterEDRPOU   string `json:"counterEdrpou"`
	CounterIBAN     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}

type Statement []StatementItem
