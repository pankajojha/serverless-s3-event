package util

type Event1 struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"created_at"`
}

type Event struct {
	ID                    string                `json:"id"`
	Description           string                `json:"description"`
	Done                  bool                  `json:"done"`
	CreatedAt             string                `json:"created_at"`
	GatewayEntryPoint     string                `json:"gatewayEntryPoint"`
	Merchant              string                `json:"merchant"`
	AuthorizationResponse AuthorizationResponse `json:"authorizationResponse"`
	Order                 Order                 `json:"order"`
	Response              Response              `json:"response"`
	SourceOfFunds         SourceOfFunds         `json:"sourceOfFunds"`
	TimeOfRecord          string                `json:"timeOfRecord"`
	Transaction           Transaction           `json:"Transaction"`
	Version               string                `json:"version"`
}

type AuthorizationResponse struct {
	CommercialCard          string `json:"commercialCard"`
	CommercialCardIndicator string `json:"commercialCardIndicator"`
	FinancialNetworkCode    bool   `json:"financialNetworkCode"`
	ProcessingCode          string `json:"processingCode"`
	ResponseCode            string `json:"responseCode"`
	Stan                    string `json:"stan"`
	TransactionIdentifier   string `json:"transactionIdentifier"`
	CardSecurityCodeError   string `json:"cardSecurityCodeError"`
}

type Order struct {
	Amount                string     `json:"amount"`
	CreationTime          string     `json:"creationTime"`
	Currency              bool       `json:"currency"`
	ID                    string     `json:"id"`
	MerchantCategoryCode  string     `json:"merchantCategoryCode"`
	NotificationURL       string     `json:"notificationUrl"`
	TotalAuthorizedAmount string     `json:"totalAuthorizedAmount"`
	TotalCapturedAmount   string     `json:"totalCapturedAmount"`
	TotalRefundedAmount   string     `json:"totalRefundedAmount"`
	Chargeback            Chargeback `json:"chargeBack"`
}

type Chargeback struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Response struct {
	AcquirerCode    string `json:"acquirerCode"`
	AcquirerMessage string `json:"acquirerMessage"`
	GatewayCode     string `json:"gatewayCode"`
}

type SourceOfFunds struct {
	Provided Provided `json:"provided"`
}

type Provided struct {
	Card Card `json:"card"`
}
type Card struct {
	Brand         string `json:"acquirerCode"`
	FundingMethod string `json:"acquirerMessage"`
	Issuer        string `json:"issuer"`
	Number        string `json:"number"`
	Scheme        string `json:"scheme"`
	StoredOnFile  string `json:"storedOnFile"`
	Expiry        Expiry `json:"expiry"`
}

type Expiry struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Transaction struct {
	Acquirer          Acquirer `json:"acquirer"`
	Amount            string   `json:"acquirerMessage"`
	AuthorizationCode string   `json:"authorizationCode"`
	Currency          string   `json:"currency"`
	Frequency         string   `json:"frequency"`
	ID                string   `json:"id"`
	Teceipt           Expiry   `json:"receipt"`
	Source            Expiry   `json:"source"`
	Terminal          Expiry   `json:"terminal"`
	Type              Expiry   `json:"receipt"`
}

type Acquirer struct {
	Batch          string `json:"batch"`
	Date           string `json:"date"`
	ID             string `json:"id"`
	MerchantID     string `json:"merchantId"`
	SettlementDate string `json:"settlementDate"`
	TimeZone       string `json:"timeZone"`
	TransactionID  Expiry `json:"transactionId"`
}
