package go-gestpay

package entities

import (
	"context"
	"encoding/xml"
	"sync"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

var UIC = map[string]string{
	"EUR": "242",
	"AED": "187",
	"ARS": "216",
	"AUD": "109",
	"BGN": "262",
	"BOB": "74",
	"BRL": "234",
	"CAD": "12",
	"CHF": "3",
	"CLP": "29",
	"CNY": "144",
	"COP": "40",
	"CRC": "77",
	"CZK": "223",
	"DKK": "7",
	"DOP": "116",
	"GBP": "2",
	"GTQ": "78",
	"HKD": "103",
	"HNL": "118",
	"HUF": "153",
	"ILS": "203",
	"INR": "31",
	"ISK": "62",
	"JPY": "71",
	"KRW": "119",
	"MYR": "55",
	"MXN": "222",
	"NIO": "120",
	"NOK": "8",
	"NZD": "113",
	"PEN": "201",
	"PLN": "237",
	"PYG": "101",
	"QAR": "189",
	"RON": "270",
	"RUB": "244",
	"SAR": "75",
	"SEK": "9",
	"SGD": "124",
	"TRY": "267",
	"TWD": "143",
	"UAH": "241",
	"USD": "1",
}

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type CallRefundS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callRefundS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	UicCode string `xml:"uicCode,omitempty" json:"uicCode,omitempty"`

	Amount string `xml:"amount,omitempty" json:"amount,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	BankTransactionId string `xml:"bankTransactionId,omitempty" json:"bankTransactionId,omitempty"`

	SettlementID string `xml:"settlementID,omitempty" json:"settlementID,omitempty"`

	OrderDetail *RefundProductDetail `xml:"OrderDetail,omitempty" json:"OrderDetail,omitempty"`

	RefundReason string `xml:"RefundReason,omitempty" json:"RefundReason,omitempty"`

	ChargeBackFraud string `xml:"chargeBackFraud,omitempty" json:"chargeBackFraud,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`

	PaymentDeviceDetails *PaymentDeviceDetails `xml:"paymentDeviceDetails,omitempty" json:"paymentDeviceDetails,omitempty"`
}

type CallRefundS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callRefundS2SResponse"`

	CallRefundS2SResult struct {
		GestPayS2S struct {
			TransactionType   string `xml:"TransactionType,omitempty" json:"TransactionType,omitempty"`
			TransactionResult string `xml:"TransactionResult,omitempty" json:"TransactionResult,omitempty"`
			ErrorCode         int    `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`
			ErrorDescription  string `xml:"ErrorDescription,omitempty" json:"ErrorDescription,omitempty"`
			ShopTransactionID string `xml:"ShopTransactionID,omitempty" json:"ShopTransactionID,omitempty"`
			BankTransactionID string `xml:"BankTransactionID,omitempty" json:"BankTransactionID,omitempty"`
		} `xml:"GestPayS2S,omitempty" json:"GestPayS2S,omitempty"`
	} `xml:"callRefundS2SResult,omitempty" json:"callRefundS2SResult,omitempty"`
}

type CallReadTrxS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callReadTrxS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	BankTransactionId string `xml:"bankTransactionId,omitempty" json:"bankTransactionId,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`

	PaymentDeviceDetails *PaymentDeviceDetails `xml:"paymentDeviceDetails,omitempty" json:"paymentDeviceDetails,omitempty"`
}

type CallReadTrxS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callReadTrxS2SResponse"`

	CallReadTrxS2SResult struct {
	} `xml:"callReadTrxS2SResult,omitempty" json:"callReadTrxS2SResult,omitempty"`
}

type CallPagamS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callPagamS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	UicCode string `xml:"uicCode,omitempty" json:"uicCode,omitempty"`

	Amount string `xml:"amount,omitempty" json:"amount,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	CardNumber string `xml:"cardNumber,omitempty" json:"cardNumber,omitempty"`

	ExpiryMonth string `xml:"expiryMonth,omitempty" json:"expiryMonth,omitempty"`

	ExpiryYear string `xml:"expiryYear,omitempty" json:"expiryYear,omitempty"`

	BuyerName string `xml:"buyerName,omitempty" json:"buyerName,omitempty"`

	BuyerEmail string `xml:"buyerEmail,omitempty" json:"buyerEmail,omitempty"`

	LanguageId string `xml:"languageId,omitempty" json:"languageId,omitempty"`

	Cvv string `xml:"cvv,omitempty" json:"cvv,omitempty"`

	Cof string `xml:"cof,omitempty" json:"cof,omitempty"`

	Min string `xml:"min,omitempty" json:"min,omitempty"`

	TransKey string `xml:"transKey,omitempty" json:"transKey,omitempty"`

	PARes string `xml:"PARes,omitempty" json:"PARes,omitempty"`

	CustomInfo string `xml:"customInfo,omitempty" json:"customInfo,omitempty"`

	IDEA string `xml:"IDEA,omitempty" json:"IDEA,omitempty"`

	RequestToken string `xml:"requestToken,omitempty" json:"requestToken,omitempty"`

	TokenValue string `xml:"tokenValue,omitempty" json:"tokenValue,omitempty"`

	ExecutionDate string `xml:"executionDate,omitempty" json:"executionDate,omitempty"`

	ClientIP string `xml:"clientIP,omitempty" json:"clientIP,omitempty"`

	ItemType string `xml:"itemType,omitempty" json:"itemType,omitempty"`

	Recurrent string `xml:"recurrent,omitempty" json:"recurrent,omitempty"`

	SubMerchantId string `xml:"subMerchantId,omitempty" json:"subMerchantId,omitempty"`

	PpSellerProtection *string `xml:"ppSellerProtection,omitempty" json:"ppSellerProtection,omitempty"`

	PaymentTypes *PaymentTypes `xml:"paymentTypes,omitempty" json:"paymentTypes,omitempty"`

	PaymentDeviceDetails *PaymentDeviceDetails `xml:"paymentDeviceDetails,omitempty" json:"paymentDeviceDetails,omitempty"`

	ShippingDetails *ShippingDetails `xml:"shippingDetails,omitempty" json:"shippingDetails,omitempty"`

	RedFraudPrevention string `xml:"redFraudPrevention,omitempty" json:"redFraudPrevention,omitempty"`

	Red_CustomerInfo *RedCustomerInfo `xml:"Red_CustomerInfo,omitempty" json:"Red_CustomerInfo,omitempty"`

	Red_ShippingInfo *RedShippingInfo `xml:"Red_ShippingInfo,omitempty" json:"Red_ShippingInfo,omitempty"`

	Red_BillingInfo *RedBillingInfo `xml:"Red_BillingInfo,omitempty" json:"Red_BillingInfo,omitempty"`

	Red_CustomerData *RedCustomerData `xml:"Red_CustomerData,omitempty" json:"Red_CustomerData,omitempty"`

	Red_CustomInfo *RedCustomInfo `xml:"Red_CustomInfo,omitempty" json:"Red_CustomInfo,omitempty"`

	Red_Items *RedItems `xml:"Red_Items,omitempty" json:"Red_Items,omitempty"`

	ApplePay *ApplePayRequest `xml:"applePay,omitempty" json:"applePay,omitempty"`

	CardReaderPayments *CardReaderPayments `xml:"cardReaderPayments,omitempty" json:"cardReaderPayments,omitempty"`

	GooglePay *GooglePayRequest `xml:"googlePay,omitempty" json:"googlePay,omitempty"`

	OrderDetails *EcommGestpayPaymentDetails `xml:"OrderDetails,omitempty" json:"OrderDetails,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`

	TransDetails *TransDetails `xml:"transDetails,omitempty" json:"transDetails,omitempty"`

	MarketPlace *Marketplace `xml:"marketPlace,omitempty" json:"marketPlace,omitempty"`
}

type CallPagamS2SResponse struct {
	XMLName            xml.Name `xml:"https://ecomms2s.sella.it/ callPagamS2SResponse"`
	CallPagamS2SResult struct {
		GestPayS2S struct {
			TransactionType   string `xml:"TransactionType,omitempty" json:"TransactionType,omitempty"`
			TransactionResult string `xml:"TransactionResult,omitempty" json:"TransactionResult,omitempty"`
			ErrorCode         int    `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`
			ErrorDescription  string `xml:"ErrorDescription,omitempty" json:"ErrorDescription,omitempty"`
			TransactionKey    string `xml:"TransactionKey,omitempty" json:"TransactionKey,omitempty"`
			BankTransactionID string `xml:"BankTransactionID,omitempty" json:"BankTransactionID,omitempty"`
			AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty"`
			Currency          string `xml:"Currency,omitempty" json:"Currency,omitempty"`
			Country           string `xml:"Country,omitempty" json:"Country,omitempty"`
		} `xml:"GestPayS2S,omitempty" json:"GestPayS2S,omitempty"`
	} `xml:"callPagamS2SResult,omitempty" json:"callPagamS2SResult,omitempty"`
}

type CallDeleteS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeleteS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	BankTransactionId string `xml:"bankTransactionId,omitempty" json:"bankTransactionId,omitempty"`

	CancelReason string `xml:"CancelReason,omitempty" json:"CancelReason,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`

	PaymentDeviceDetails *PaymentDeviceDetails `xml:"paymentDeviceDetails,omitempty" json:"paymentDeviceDetails,omitempty"`
}

type CallDeleteS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeleteS2SResponse"`

	CallDeleteS2SResult struct {
	} `xml:"callDeleteS2SResult,omitempty" json:"callDeleteS2SResult,omitempty"`
}

type CallSettleS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callSettleS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	UicCode string `xml:"uicCode,omitempty" json:"uicCode,omitempty"`

	Amount string `xml:"amount,omitempty" json:"amount,omitempty"`

	ShopTransID string `xml:"shopTransID,omitempty" json:"shopTransID,omitempty"`

	BankTransID string `xml:"bankTransID,omitempty" json:"bankTransID,omitempty"`

	SettlementID string `xml:"settlementID,omitempty" json:"settlementID,omitempty"`

	FullFillment *FullFillmentDetails `xml:"FullFillment,omitempty" json:"FullFillment,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`

	PaymentDeviceDetails *PaymentDeviceDetails `xml:"paymentDeviceDetails,omitempty" json:"paymentDeviceDetails,omitempty"`
}

type CallSettleS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callSettleS2SResponse"`

	CallSettleS2SResult struct {
	} `xml:"callSettleS2SResult,omitempty" json:"callSettleS2SResult,omitempty"`
}

type CallVerifycardS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callVerifycardS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	CardNumber string `xml:"cardNumber,omitempty" json:"cardNumber,omitempty"`

	ExpMonth string `xml:"expMonth,omitempty" json:"expMonth,omitempty"`

	ExpYear string `xml:"expYear,omitempty" json:"expYear,omitempty"`

	CVV2 string `xml:"CVV2,omitempty" json:"CVV2,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallVerifycardS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callVerifycardS2SResponse"`

	CallVerifycardS2SResult struct {
	} `xml:"callVerifycardS2SResult,omitempty" json:"callVerifycardS2SResult,omitempty"`
}

type CallCheckCartaS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callCheckCartaS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	CardNumber string `xml:"cardNumber,omitempty" json:"cardNumber,omitempty"`

	ExpMonth string `xml:"expMonth,omitempty" json:"expMonth,omitempty"`

	ExpYear string `xml:"expYear,omitempty" json:"expYear,omitempty"`

	CVV2 string `xml:"CVV2,omitempty" json:"CVV2,omitempty"`

	WithAuth string `xml:"withAuth,omitempty" json:"withAuth,omitempty"`

	TokenValue string `xml:"tokenValue,omitempty" json:"tokenValue,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallCheckCartaS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callCheckCartaS2SResponse"`

	CallCheckCartaS2SResult struct {
	} `xml:"callCheckCartaS2SResult,omitempty" json:"callCheckCartaS2SResult,omitempty"`
}

type CallRenounce struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callRenounce"`
}

type CallRenounceResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callRenounceResponse"`

	CallRenounceResult struct {
	} `xml:"callRenounceResult,omitempty" json:"callRenounceResult,omitempty"`
}

type CallRequestTokenS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallRequestTokenS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	RequestToken string `xml:"requestToken,omitempty" json:"requestToken,omitempty"`

	CardNumber string `xml:"cardNumber,omitempty" json:"cardNumber,omitempty"`

	ExpiryMonth string `xml:"expiryMonth,omitempty" json:"expiryMonth,omitempty"`

	ExpiryYear string `xml:"expiryYear,omitempty" json:"expiryYear,omitempty"`

	Cvv string `xml:"cvv,omitempty" json:"cvv,omitempty"`

	WithAuth string `xml:"withAuth,omitempty" json:"withAuth,omitempty"`

	OriginalId string `xml:"originalId,omitempty" json:"originalId,omitempty"`

	Description string `xml:"description,omitempty" json:"description,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallRequestTokenS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallRequestTokenS2SResponse"`

	CallRequestTokenS2SResult struct {
	} `xml:"CallRequestTokenS2SResult,omitempty" json:"CallRequestTokenS2SResult,omitempty"`
}

type CallDeleteTokenS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeleteTokenS2S"`

	TokenValue string `xml:"tokenValue,omitempty" json:"tokenValue,omitempty"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallDeleteTokenS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeleteTokenS2SResponse"`

	CallDeleteTokenS2SResult struct {
	} `xml:"callDeleteTokenS2SResult,omitempty" json:"callDeleteTokenS2SResult,omitempty"`
}

type CallUpdateTokenS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateTokenS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	Token string `xml:"token,omitempty" json:"token,omitempty"`

	ExpiryMonth string `xml:"expiryMonth,omitempty" json:"expiryMonth,omitempty"`

	ExpiryYear string `xml:"expiryYear,omitempty" json:"expiryYear,omitempty"`

	WithAut string `xml:"withAut,omitempty" json:"withAut,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallUpdateTokenS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateTokenS2SResponse"`

	CallUpdateTokenS2SResult struct {
	} `xml:"CallUpdateTokenS2SResult,omitempty" json:"CallUpdateTokenS2SResult,omitempty"`
}

type CallIdealListS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallIdealListS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallIdealListS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallIdealListS2SResponse"`

	CallIdealListS2SResult struct {
	} `xml:"CallIdealListS2SResult,omitempty" json:"CallIdealListS2SResult,omitempty"`
}

type CallMyBankListS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallMyBankListS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	LanguageId string `xml:"languageId,omitempty" json:"languageId,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallMyBankListS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallMyBankListS2SResponse"`

	CallMyBankListS2SResult struct {
	} `xml:"CallMyBankListS2SResult,omitempty" json:"CallMyBankListS2SResult,omitempty"`
}

type CallUpdateCustomParamS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateCustomParamS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	BankTransactionId string `xml:"bankTransactionId,omitempty" json:"bankTransactionId,omitempty"`

	ParamName string `xml:"paramName,omitempty" json:"paramName,omitempty"`

	ParamValue string `xml:"paramValue,omitempty" json:"paramValue,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallUpdateCustomParamS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateCustomParamS2SResponse"`

	CallUpdateCustomParamS2SResult struct {
	} `xml:"CallUpdateCustomParamS2SResult,omitempty" json:"CallUpdateCustomParamS2SResult,omitempty"`
}

type CallUpdateOrderS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateOrderS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	ShopTransactionId string `xml:"shopTransactionId,omitempty" json:"shopTransactionId,omitempty"`

	BankTransactionId string `xml:"bankTransactionId,omitempty" json:"bankTransactionId,omitempty"`

	OrderDetails *EcommGestpayPaymentDetails `xml:"OrderDetails,omitempty" json:"OrderDetails,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallUpdateOrderS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ CallUpdateOrderS2SResponse"`

	CallUpdateOrderS2SResult struct {
	} `xml:"CallUpdateOrderS2SResult,omitempty" json:"CallUpdateOrderS2SResult,omitempty"`
}

type CallExchangeRateS2S struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callExchangeRateS2S"`

	ShopLogin string `xml:"shopLogin,omitempty" json:"shopLogin,omitempty"`

	UicCode string `xml:"uicCode,omitempty" json:"uicCode,omitempty"`

	IsoCode string `xml:"isoCode,omitempty" json:"isoCode,omitempty"`

	NumericIsoCode string `xml:"numericIsoCode,omitempty" json:"numericIsoCode,omitempty"`

	Apikey string `xml:"apikey,omitempty" json:"apikey,omitempty"`
}

type CallExchangeRateS2SResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callExchangeRateS2SResponse"`

	CallExchangeRateS2SResult struct {
	} `xml:"callExchangeRateS2SResult,omitempty" json:"callExchangeRateS2SResult,omitempty"`
}

type CallDeviceActivation struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeviceActivation"`

	TerminalId string `xml:"terminalId,omitempty" json:"terminalId,omitempty"`
}

type CallDeviceActivationResponse struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ callDeviceActivationResponse"`

	CallDeviceActivationResult struct {
	} `xml:"callDeviceActivationResult,omitempty" json:"callDeviceActivationResult,omitempty"`
}

type RefundProductDetail struct {
	ProductDetails *ArrayOfProductDetail `xml:"ProductDetails,omitempty" json:"ProductDetails,omitempty"`
}

type ArrayOfProductDetail struct {
	ProductDetail []*ProductDetail `xml:"ProductDetail,omitempty" json:"ProductDetail,omitempty"`
}

type ProductDetail struct {
	ProductCode string `xml:"ProductCode,omitempty" json:"ProductCode,omitempty"`

	SKU string `xml:"SKU,omitempty" json:"SKU,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	Quantity string `xml:"Quantity,omitempty" json:"Quantity,omitempty"`

	Price string `xml:"Price,omitempty" json:"Price,omitempty"`

	UnitPrice string `xml:"UnitPrice,omitempty" json:"UnitPrice,omitempty"`

	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	TaxType string `xml:"TaxType,omitempty" json:"TaxType,omitempty"`

	Vat string `xml:"Vat,omitempty" json:"Vat,omitempty"`

	Discount string `xml:"Discount,omitempty" json:"Discount,omitempty"`

	RequiresShipping string `xml:"RequiresShipping,omitempty" json:"RequiresShipping,omitempty"`

	Condition string `xml:"Condition,omitempty" json:"Condition,omitempty"`

	Seller string `xml:"Seller,omitempty" json:"Seller,omitempty"`

	Category string `xml:"Category,omitempty" json:"Category,omitempty"`

	SubCategory string `xml:"SubCategory,omitempty" json:"SubCategory,omitempty"`

	Brand string `xml:"Brand,omitempty" json:"Brand,omitempty"`

	DeliveryAt string `xml:"DeliveryAt,omitempty" json:"DeliveryAt,omitempty"`

	DigitalGiftCardDetails *DigitalGiftCardDetails `xml:"DigitalGiftCardDetails,omitempty" json:"DigitalGiftCardDetails,omitempty"`
}

type DigitalGiftCardDetails struct {
	SenderName string `xml:"SenderName,omitempty" json:"SenderName,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty" json:"DisplayName,omitempty"`

	GreetingMessage string `xml:"GreetingMessage,omitempty" json:"GreetingMessage,omitempty"`

	Recipient *Recipient `xml:"Recipient,omitempty" json:"Recipient,omitempty"`
}

type Recipient struct {
	Email string `xml:"Email,omitempty" json:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty"`

	Social string `xml:"Social,omitempty" json:"Social,omitempty"`
}

type PaymentDeviceDetails struct {
	TerminalId string `xml:"TerminalId,omitempty" json:"TerminalId,omitempty"`
}

type PaymentTypes struct {
	PaymentType []string `xml:"paymentType,omitempty" json:"paymentType,omitempty"`
}

type ShippingDetails struct {
	ShipToName string `xml:"shipToName,omitempty" json:"shipToName,omitempty"`

	ShipToStreet string `xml:"shipToStreet,omitempty" json:"shipToStreet,omitempty"`

	ShipToCity string `xml:"shipToCity,omitempty" json:"shipToCity,omitempty"`

	ShipToState string `xml:"shipToState,omitempty" json:"shipToState,omitempty"`

	ShipToCountryCode string `xml:"shipToCountryCode,omitempty" json:"shipToCountryCode,omitempty"`

	ShipToZip string `xml:"shipToZip,omitempty" json:"shipToZip,omitempty"`

	ShipToStreet2 string `xml:"shipToStreet2,omitempty" json:"shipToStreet2,omitempty"`
}

type RedCustomerInfo struct {
	Customer_Title string `xml:"Customer_Title,omitempty" json:"Customer_Title,omitempty"`

	Customer_Name string `xml:"Customer_Name,omitempty" json:"Customer_Name,omitempty"`

	Customer_Surname string `xml:"Customer_Surname,omitempty" json:"Customer_Surname,omitempty"`

	Customer_Email string `xml:"Customer_Email,omitempty" json:"Customer_Email,omitempty"`

	Customer_Address string `xml:"Customer_Address,omitempty" json:"Customer_Address,omitempty"`

	Customer_Address2 string `xml:"Customer_Address2,omitempty" json:"Customer_Address2,omitempty"`

	Customer_City string `xml:"Customer_City,omitempty" json:"Customer_City,omitempty"`

	Customer_StateCode string `xml:"Customer_StateCode,omitempty" json:"Customer_StateCode,omitempty"`

	Customer_Country string `xml:"Customer_Country,omitempty" json:"Customer_Country,omitempty"`

	Customer_PostalCode string `xml:"Customer_PostalCode,omitempty" json:"Customer_PostalCode,omitempty"`

	Customer_Phone string `xml:"Customer_Phone,omitempty" json:"Customer_Phone,omitempty"`
}

type RedShippingInfo struct {
	Shipping_Name string `xml:"Shipping_Name,omitempty" json:"Shipping_Name,omitempty"`

	Shipping_Surname string `xml:"Shipping_Surname,omitempty" json:"Shipping_Surname,omitempty"`

	Shipping_Email string `xml:"Shipping_Email,omitempty" json:"Shipping_Email,omitempty"`

	Shipping_Address string `xml:"Shipping_Address,omitempty" json:"Shipping_Address,omitempty"`

	Shipping_Address2 string `xml:"Shipping_Address2,omitempty" json:"Shipping_Address2,omitempty"`

	Shipping_City string `xml:"Shipping_City,omitempty" json:"Shipping_City,omitempty"`

	Shipping_StateCode string `xml:"Shipping_StateCode,omitempty" json:"Shipping_StateCode,omitempty"`

	Shipping_Country string `xml:"Shipping_Country,omitempty" json:"Shipping_Country,omitempty"`

	Shipping_PostalCode string `xml:"Shipping_PostalCode,omitempty" json:"Shipping_PostalCode,omitempty"`

	Shipping_HomePhone string `xml:"Shipping_HomePhone,omitempty" json:"Shipping_HomePhone,omitempty"`

	Shipping_MobilePhone string `xml:"Shipping_MobilePhone,omitempty" json:"Shipping_MobilePhone,omitempty"`

	Shipping_FaxPhone string `xml:"Shipping_FaxPhone,omitempty" json:"Shipping_FaxPhone,omitempty"`

	Shipping_TimeToDeparture string `xml:"Shipping_TimeToDeparture,omitempty" json:"Shipping_TimeToDeparture,omitempty"`
}

type RedBillingInfo struct {
	Billing_Id string `xml:"Billing_Id,omitempty" json:"Billing_Id,omitempty"`

	Billing_Name string `xml:"Billing_Name,omitempty" json:"Billing_Name,omitempty"`

	Billing_Surname string `xml:"Billing_Surname,omitempty" json:"Billing_Surname,omitempty"`

	Billing_DateOfBirth string `xml:"Billing_DateOfBirth,omitempty" json:"Billing_DateOfBirth,omitempty"`

	Billing_Email string `xml:"Billing_Email,omitempty" json:"Billing_Email,omitempty"`

	Billing_Address string `xml:"Billing_Address,omitempty" json:"Billing_Address,omitempty"`

	Billing_Address2 string `xml:"Billing_Address2,omitempty" json:"Billing_Address2,omitempty"`

	Billing_City string `xml:"Billing_City,omitempty" json:"Billing_City,omitempty"`

	Billing_StateCode string `xml:"Billing_StateCode,omitempty" json:"Billing_StateCode,omitempty"`

	Billing_Country string `xml:"Billing_Country,omitempty" json:"Billing_Country,omitempty"`

	Billing_PostalCode string `xml:"Billing_PostalCode,omitempty" json:"Billing_PostalCode,omitempty"`

	Billing_HomePhone string `xml:"Billing_HomePhone,omitempty" json:"Billing_HomePhone,omitempty"`

	Billing_WorkPhone string `xml:"Billing_WorkPhone,omitempty" json:"Billing_WorkPhone,omitempty"`

	Billing_MobilePhone string `xml:"Billing_MobilePhone,omitempty" json:"Billing_MobilePhone,omitempty"`
}

type RedCustomerData struct {
	MerchantWebSite string `xml:"MerchantWebSite,omitempty" json:"MerchantWebSite,omitempty"`

	Customer_IPAddress string `xml:"Customer_IPAddress,omitempty" json:"Customer_IPAddress,omitempty"`

	PC_FingerPrint string `xml:"PC_FingerPrint,omitempty" json:"PC_FingerPrint,omitempty"`

	PreviousCustomer string `xml:"PreviousCustomer,omitempty" json:"PreviousCustomer,omitempty"`

	Red_Merchant_ID string `xml:"Red_Merchant_ID,omitempty" json:"Red_Merchant_ID,omitempty"`

	Red_ServiceType string `xml:"Red_ServiceType,omitempty" json:"Red_ServiceType,omitempty"`
}

type RedCustomInfo struct {
	UserCustomData []string `xml:"UserCustomData,omitempty" json:"UserCustomData,omitempty"`
}

type RedItems struct {
	NumberOfItems string `xml:"NumberOfItems,omitempty" json:"NumberOfItems,omitempty"`

	Red_Item []*RedItem `xml:"Red_Item,omitempty" json:"Red_Item,omitempty"`
}

type RedItem struct {
	Item_ProductCode string `xml:"Item_ProductCode,omitempty" json:"Item_ProductCode,omitempty"`

	Item_StockKeepingUnit string `xml:"Item_StockKeepingUnit,omitempty" json:"Item_StockKeepingUnit,omitempty"`

	Item_Description string `xml:"Item_Description,omitempty" json:"Item_Description,omitempty"`

	Item_Quantity string `xml:"Item_Quantity,omitempty" json:"Item_Quantity,omitempty"`

	Item_UnitCost string `xml:"Item_UnitCost,omitempty" json:"Item_UnitCost,omitempty"`

	Item_TotalCost string `xml:"Item_TotalCost,omitempty" json:"Item_TotalCost,omitempty"`

	Item_ShippingNumber string `xml:"Item_ShippingNumber,omitempty" json:"Item_ShippingNumber,omitempty"`

	Item_GiftMessage string `xml:"Item_GiftMessage,omitempty" json:"Item_GiftMessage,omitempty"`

	Item_PartEAN_Number string `xml:"Item_PartEAN_Number,omitempty" json:"Item_PartEAN_Number,omitempty"`

	Item_ShippingComments string `xml:"Item_ShippingComments,omitempty" json:"Item_ShippingComments,omitempty"`
}

type ApplePayRequest struct {
	ApplePayPKPaymentToken string `xml:"applePayPKPaymentToken,omitempty" json:"applePayPKPaymentToken,omitempty"`

	PKPaymentToken struct {
	} `xml:"PKPaymentToken,omitempty" json:"PKPaymentToken,omitempty"`

	OnlinePaymentCryptogram string `xml:"onlinePaymentCryptogram,omitempty" json:"onlinePaymentCryptogram,omitempty"`

	EciIndicator string `xml:"eciIndicator,omitempty" json:"eciIndicator,omitempty"`

	RequestToken string `xml:"requestToken,omitempty" json:"requestToken,omitempty"`

	Token string `xml:"token,omitempty" json:"token,omitempty"`
}

type CardReaderPayments struct {
	DeviceName string `xml:"deviceName,omitempty" json:"deviceName,omitempty"`

	DeviceProvidedInfo *DeviceProvidedInfo `xml:"deviceProvidedInfo,omitempty" json:"deviceProvidedInfo,omitempty"`

	BuyerSignature string `xml:"buyerSignature,omitempty" json:"buyerSignature,omitempty"`
}

type DeviceProvidedInfo struct {
	DeviceID string `xml:"deviceID,omitempty" json:"deviceID,omitempty"`

	Ksn string `xml:"ksn,omitempty" json:"ksn,omitempty"`

	CardData *CardData `xml:"cardData,omitempty" json:"cardData,omitempty"`
}

type CardData struct {
	Track1 string `xml:"track1,omitempty" json:"track1,omitempty"`

	Track2 string `xml:"track2,omitempty" json:"track2,omitempty"`

	CardHolderName string `xml:"cardHolderName,omitempty" json:"cardHolderName,omitempty"`
}

type GooglePayRequest struct {
	TokenizationData *TokenizationData `xml:"tokenizationData,omitempty" json:"tokenizationData,omitempty"`
}

type TokenizationData struct {
	Token string `xml:"token,omitempty" json:"token,omitempty"`
}

type EcommGestpayPaymentDetails struct {
	FraudPrevention *FraudPrevention `xml:"FraudPrevention,omitempty" json:"FraudPrevention,omitempty"`

	CustomerDetail *CustomerDetail `xml:"CustomerDetail,omitempty" json:"CustomerDetail,omitempty"`

	ShippingAddress *ShippingAddress `xml:"ShippingAddress,omitempty" json:"ShippingAddress,omitempty"`

	BillingAddress *BillingAddress `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty"`

	ProductDetails *ArrayOfProductDetail `xml:"ProductDetails,omitempty" json:"ProductDetails,omitempty"`

	DiscountCodes *ArrayOfDiscountCode `xml:"DiscountCodes,omitempty" json:"DiscountCodes,omitempty"`

	ShippingLines *ArrayOfShippingLine `xml:"ShippingLines,omitempty" json:"ShippingLines,omitempty"`

	AccommodationDetails *ArrayOfAccommodationDetail `xml:"AccommodationDetails,omitempty" json:"AccommodationDetails,omitempty"`

	TravelTicketDetails *ArrayOfTravelTicketDetail `xml:"TravelTicketDetails,omitempty" json:"TravelTicketDetails,omitempty"`

	PassengerDetails *ArrayOfPassengerDetail `xml:"PassengerDetails,omitempty" json:"PassengerDetails,omitempty"`
}

type FraudPrevention struct {
	SubmitForReview string `xml:"SubmitForReview,omitempty" json:"SubmitForReview,omitempty"`

	OrderDateTime string `xml:"OrderDateTime,omitempty" json:"OrderDateTime,omitempty"`

	OrderNote string `xml:"OrderNote,omitempty" json:"OrderNote,omitempty"`

	Source string `xml:"Source,omitempty" json:"Source,omitempty"`

	SubmissionReason string `xml:"SubmissionReason,omitempty" json:"SubmissionReason,omitempty"`

	BeaconSessionID string `xml:"BeaconSessionID,omitempty" json:"BeaconSessionID,omitempty"`

	VendorID string `xml:"VendorID,omitempty" json:"VendorID,omitempty"`

	VendorName string `xml:"VendorName,omitempty" json:"VendorName,omitempty"`
}

type CustomerDetail struct {
	ProfileID string `xml:"ProfileID,omitempty" json:"ProfileID,omitempty"`

	MerchantCustomerID string `xml:"MerchantCustomerID,omitempty" json:"MerchantCustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty"`

	Lastname string `xml:"Lastname,omitempty" json:"Lastname,omitempty"`

	PrimaryEmail string `xml:"PrimaryEmail,omitempty" json:"PrimaryEmail,omitempty"`

	SecondaryEmail string `xml:"SecondaryEmail,omitempty" json:"SecondaryEmail,omitempty"`

	PrimaryPhone string `xml:"PrimaryPhone,omitempty" json:"PrimaryPhone,omitempty"`

	SecondaryPhone string `xml:"SecondaryPhone,omitempty" json:"SecondaryPhone,omitempty"`

	DateOfBirth string `xml:"DateOfBirth,omitempty" json:"DateOfBirth,omitempty"`

	Gender string `xml:"Gender,omitempty" json:"Gender,omitempty"`

	SocialSecurityNumber string `xml:"SocialSecurityNumber,omitempty" json:"SocialSecurityNumber,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty"`

	CreatedAtDate string `xml:"CreatedAtDate,omitempty" json:"CreatedAtDate,omitempty"`

	VerifiedEmail string `xml:"VerifiedEmail,omitempty" json:"VerifiedEmail,omitempty"`

	AccountType string `xml:"AccountType,omitempty" json:"AccountType,omitempty"`

	Social *CustomerSocial `xml:"Social,omitempty" json:"Social,omitempty"`
}

type CustomerSocial struct {
	Network string `xml:"Network,omitempty" json:"Network,omitempty"`

	PublicUsername string `xml:"PublicUsername,omitempty" json:"PublicUsername,omitempty"`

	CommunityScore string `xml:"CommunityScore,omitempty" json:"CommunityScore,omitempty"`

	ProfilePicture string `xml:"ProfilePicture,omitempty" json:"ProfilePicture,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty"`

	Bio string `xml:"Bio,omitempty" json:"Bio,omitempty"`

	AccountUrl string `xml:"AccountUrl,omitempty" json:"AccountUrl,omitempty"`

	Following string `xml:"Following,omitempty" json:"Following,omitempty"`

	Followed string `xml:"Followed,omitempty" json:"Followed,omitempty"`

	Posts string `xml:"Posts,omitempty" json:"Posts,omitempty"`

	Id string `xml:"Id,omitempty" json:"Id,omitempty"`

	AuthToken string `xml:"AuthToken,omitempty" json:"AuthToken,omitempty"`

	SocialData string `xml:"SocialData,omitempty" json:"SocialData,omitempty"`
}

type ShippingAddress struct {
	ProfileID string `xml:"ProfileID,omitempty" json:"ProfileID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty"`

	Lastname string `xml:"Lastname,omitempty" json:"Lastname,omitempty"`

	StreetName string `xml:"StreetName,omitempty" json:"StreetName,omitempty"`

	Streetname2 string `xml:"Streetname2,omitempty" json:"Streetname2,omitempty"`

	HouseNumber string `xml:"HouseNumber,omitempty" json:"HouseNumber,omitempty"`

	HouseExtention string `xml:"HouseExtention,omitempty" json:"HouseExtention,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty"`

	ZipCode string `xml:"ZipCode,omitempty" json:"ZipCode,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty"`

	PrimaryPhone string `xml:"PrimaryPhone,omitempty" json:"PrimaryPhone,omitempty"`

	SecondaryPhone string `xml:"SecondaryPhone,omitempty" json:"SecondaryPhone,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty"`

	StateCode string `xml:"StateCode,omitempty" json:"StateCode,omitempty"`
}

type BillingAddress struct {
	ProfileID string `xml:"ProfileID,omitempty" json:"ProfileID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty"`

	Lastname string `xml:"Lastname,omitempty" json:"Lastname,omitempty"`

	StreetNumber string `xml:"StreetNumber,omitempty" json:"StreetNumber,omitempty"`

	StreetName string `xml:"StreetName,omitempty" json:"StreetName,omitempty"`

	Streetname2 string `xml:"Streetname2,omitempty" json:"Streetname2,omitempty"`

	HouseNumber string `xml:"HouseNumber,omitempty" json:"HouseNumber,omitempty"`

	HouseExtention string `xml:"HouseExtention,omitempty" json:"HouseExtention,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty"`

	ZipCode string `xml:"ZipCode,omitempty" json:"ZipCode,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty"`

	PrimaryPhone string `xml:"PrimaryPhone,omitempty" json:"PrimaryPhone,omitempty"`

	SecondaryPhone string `xml:"SecondaryPhone,omitempty" json:"SecondaryPhone,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty"`

	StateCode string `xml:"StateCode,omitempty" json:"StateCode,omitempty"`
}

type ArrayOfDiscountCode struct {
	DiscountCode []*DiscountCode `xml:"DiscountCode,omitempty" json:"DiscountCode,omitempty"`
}

type DiscountCode struct {
	Amount string `xml:"Amount,omitempty" json:"Amount,omitempty"`

	Code string `xml:"Code,omitempty" json:"Code,omitempty"`
}

type ArrayOfShippingLine struct {
	ShippingLine []*ShippingLine `xml:"ShippingLine,omitempty" json:"ShippingLine,omitempty"`
}

type ShippingLine struct {
	Price string `xml:"Price,omitempty" json:"Price,omitempty"`

	Title string `xml:"Title,omitempty" json:"Title,omitempty"`

	Code string `xml:"Code,omitempty" json:"Code,omitempty"`
}

type ArrayOfAccommodationDetail struct {
	AccommodationDetail []*AccommodationDetail `xml:"AccommodationDetail,omitempty" json:"AccommodationDetail,omitempty"`
}

type AccommodationDetail struct {
	City string `xml:"City,omitempty" json:"City,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty"`

	Title string `xml:"Title,omitempty" json:"Title,omitempty"`

	CheckInDate string `xml:"CheckInDate,omitempty" json:"CheckInDate,omitempty"`

	CheckOutDate string `xml:"CheckOutDate,omitempty" json:"CheckOutDate,omitempty"`

	ProductId string `xml:"ProductId,omitempty" json:"ProductId,omitempty"`

	Rating string `xml:"Rating,omitempty" json:"Rating,omitempty"`

	NumberOfGuests string `xml:"NumberOfGuests,omitempty" json:"NumberOfGuests,omitempty"`

	AccommodationType string `xml:"AccommodationType,omitempty" json:"AccommodationType,omitempty"`

	RoomType string `xml:"RoomType,omitempty" json:"RoomType,omitempty"`

	Price string `xml:"Price,omitempty" json:"Price,omitempty"`

	Quantity string `xml:"Quantity,omitempty" json:"Quantity,omitempty"`

	CancellationPolicy string `xml:"CancellationPolicy,omitempty" json:"CancellationPolicy,omitempty"`
}

type ArrayOfTravelTicketDetail struct {
	TravelTicketDetail []*TravelTicketDetail `xml:"TravelTicketDetail,omitempty" json:"TravelTicketDetail,omitempty"`
}

type TravelTicketDetail struct {
	ArrivalCity string `xml:"ArrivalCity,omitempty" json:"ArrivalCity,omitempty"`

	ArrivalCountryCode string `xml:"ArrivalCountryCode,omitempty" json:"ArrivalCountryCode,omitempty"`

	ArrivalDate string `xml:"ArrivalDate,omitempty" json:"ArrivalDate,omitempty"`

	ArrivalPortCode string `xml:"ArrivalPortCode,omitempty" json:"ArrivalPortCode,omitempty"`

	CarrierCode string `xml:"CarrierCode,omitempty" json:"CarrierCode,omitempty"`

	CarrierName string `xml:"CarrierName,omitempty" json:"CarrierName,omitempty"`

	DepartureCity string `xml:"DepartureCity,omitempty" json:"DepartureCity,omitempty"`

	DepartureCountryCode string `xml:"DepartureCountryCode,omitempty" json:"DepartureCountryCode,omitempty"`

	DepartureDate string `xml:"DepartureDate,omitempty" json:"DepartureDate,omitempty"`

	DeparturePortCode string `xml:"DeparturePortCode,omitempty" json:"DeparturePortCode,omitempty"`

	LegId string `xml:"LegId,omitempty" json:"LegId,omitempty"`

	LegIndex string `xml:"LegIndex,omitempty" json:"LegIndex,omitempty"`

	RouteIndex string `xml:"RouteIndex,omitempty" json:"RouteIndex,omitempty"`

	TicketClass string `xml:"TicketClass,omitempty" json:"TicketClass,omitempty"`

	Price string `xml:"Price,omitempty" json:"Price,omitempty"`

	Quantity string `xml:"Quantity,omitempty" json:"Quantity,omitempty"`

	Title string `xml:"Title,omitempty" json:"Title,omitempty"`
}

type ArrayOfPassengerDetail struct {
	PassengerDetail []*PassengerDetail `xml:"PassengerDetail,omitempty" json:"PassengerDetail,omitempty"`
}

type PassengerDetail struct {
	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty"`

	DateOfBirth string `xml:"DateOfBirth,omitempty" json:"DateOfBirth,omitempty"`

	NationalityCode string `xml:"NationalityCode,omitempty" json:"NationalityCode,omitempty"`

	InsuranceType string `xml:"InsuranceType,omitempty" json:"InsuranceType,omitempty"`

	InsurancePrice string `xml:"InsurancePrice,omitempty" json:"InsurancePrice,omitempty"`

	DocumentNumber string `xml:"DocumentNumber,omitempty" json:"DocumentNumber,omitempty"`

	DocumentType string `xml:"DocumentType,omitempty" json:"DocumentType,omitempty"`

	DocumentIssueDate string `xml:"DocumentIssueDate,omitempty" json:"DocumentIssueDate,omitempty"`

	DocumentExpirationDate string `xml:"DocumentExpirationDate,omitempty" json:"DocumentExpirationDate,omitempty"`

	PassengerType string `xml:"PassengerType,omitempty" json:"PassengerType,omitempty"`
}

type TransDetails struct {
	Type string `xml:"type,omitempty" json:"type,omitempty"`

	AuthenticationAmount string `xml:"authenticationAmount,omitempty" json:"authenticationAmount,omitempty"`

	ThreeDSAuthResult *ThreeDSAuthResult `xml:"threeDSAuthResult,omitempty" json:"threeDSAuthResult,omitempty"`

	ThreeDsContainer *ThreeDsContainer `xml:"threeDsContainer,omitempty" json:"threeDsContainer,omitempty"`

	RecurringTransaction *RecurringTransaction `xml:"recurringTransaction,omitempty" json:"recurringTransaction,omitempty"`

	PreviousTransDetails *PreviousTransDetails `xml:"previousTransDetails,omitempty" json:"previousTransDetails,omitempty"`
}

type ThreeDSAuthResult struct {
	AuthenticationLevel string `xml:"authenticationLevel,omitempty" json:"authenticationLevel,omitempty"`

	AuthenticationStatus string `xml:"authenticationStatus,omitempty" json:"authenticationStatus,omitempty"`

	XID string `xml:"XID,omitempty" json:"XID,omitempty"`

	AV string `xml:"AV,omitempty" json:"AV,omitempty"`

	ECI string `xml:"ECI,omitempty" json:"ECI,omitempty"`

	AVAlgorithm string `xml:"AVAlgorithm,omitempty" json:"AVAlgorithm,omitempty"`

	ThreeDsVersion string `xml:"threeDsVersion,omitempty" json:"threeDsVersion,omitempty"`
}

type ThreeDsContainer struct {
	TransTypeReq string `xml:"transTypeReq,omitempty" json:"transTypeReq,omitempty"`

	AcquirerBIN string `xml:"acquirerBIN,omitempty" json:"acquirerBIN,omitempty"`

	AcquirerMerchantID string `xml:"acquirerMerchantID,omitempty" json:"acquirerMerchantID,omitempty"`

	Exemption string `xml:"exemption,omitempty" json:"exemption,omitempty"`

	BuyerDetails *BuyerDetails `xml:"buyerDetails,omitempty" json:"buyerDetails,omitempty"`

	MerchantRiskIndicator *MerchantRiskIndicator `xml:"merchantRiskIndicator,omitempty" json:"merchantRiskIndicator,omitempty"`

	SDKDetails string `xml:"SDKDetails,omitempty" json:"SDKDetails,omitempty"`
}

type BuyerDetails struct {
	ProfileDetails *BuyerProfileDetails `xml:"profileDetails,omitempty" json:"profileDetails,omitempty"`

	BillingAddress *ThreeDSBillingAddress `xml:"billingAddress,omitempty" json:"billingAddress,omitempty"`

	ShippingAddress *ThreeDSShippingAddress `xml:"shippingAddress,omitempty" json:"shippingAddress,omitempty"`

	AddrMatch string `xml:"addrMatch,omitempty" json:"addrMatch,omitempty"`

	CardHolder *CardHolder `xml:"cardHolder,omitempty" json:"cardHolder,omitempty"`

	AccInfo *AccInfo `xml:"accInfo,omitempty" json:"accInfo,omitempty"`
}

type BuyerProfileDetails struct {
	CardHolderID string `xml:"cardHolderID,omitempty" json:"cardHolderID,omitempty"`

	AuthData string `xml:"authData,omitempty" json:"authData,omitempty"`

	AuthMethod string `xml:"authMethod,omitempty" json:"authMethod,omitempty"`

	AuthTimestamp string `xml:"authTimestamp,omitempty" json:"authTimestamp,omitempty"`
}

type ThreeDSBillingAddress struct {
	City string `xml:"city,omitempty" json:"city,omitempty"`

	Country string `xml:"country,omitempty" json:"country,omitempty"`

	Line1 string `xml:"line1,omitempty" json:"line1,omitempty"`

	Line2 string `xml:"line2,omitempty" json:"line2,omitempty"`

	Line3 string `xml:"line3,omitempty" json:"line3,omitempty"`

	PostCode string `xml:"postCode,omitempty" json:"postCode,omitempty"`

	State string `xml:"state,omitempty" json:"state,omitempty"`
}

type ThreeDSShippingAddress struct {
	City string `xml:"city,omitempty" json:"city,omitempty"`

	Country string `xml:"country,omitempty" json:"country,omitempty"`

	Line1 string `xml:"line1,omitempty" json:"line1,omitempty"`

	Line2 string `xml:"line2,omitempty" json:"line2,omitempty"`

	Line3 string `xml:"line3,omitempty" json:"line3,omitempty"`

	PostCode string `xml:"postCode,omitempty" json:"postCode,omitempty"`

	State string `xml:"state,omitempty" json:"state,omitempty"`
}

type CardHolder struct {
	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Email string `xml:"email,omitempty" json:"email,omitempty"`

	HomePhone_cc string `xml:"homePhone_cc,omitempty" json:"homePhone_cc,omitempty"`

	HomePhone_num string `xml:"homePhone_num,omitempty" json:"homePhone_num,omitempty"`

	MobilePhone_cc string `xml:"mobilePhone_cc,omitempty" json:"mobilePhone_cc,omitempty"`

	MobilePhone_num string `xml:"mobilePhone_num,omitempty" json:"mobilePhone_num,omitempty"`

	WorkPhone_cc string `xml:"workPhone_cc,omitempty" json:"workPhone_cc,omitempty"`

	WorkPhone_num string `xml:"workPhone_num,omitempty" json:"workPhone_num,omitempty"`
}

type AccInfo struct {
	ChAccAgeInd string `xml:"chAccAgeInd,omitempty" json:"chAccAgeInd,omitempty"`

	ChAccChange string `xml:"chAccChange,omitempty" json:"chAccChange,omitempty"`

	ChAccChangeInd string `xml:"chAccChangeInd,omitempty" json:"chAccChangeInd,omitempty"`

	ChAccDate string `xml:"chAccDate,omitempty" json:"chAccDate,omitempty"`

	ChAccPwChange string `xml:"chAccPwChange,omitempty" json:"chAccPwChange,omitempty"`

	ChAccPwChangeInd string `xml:"chAccPwChangeInd,omitempty" json:"chAccPwChangeInd,omitempty"`

	NbPurchaseAccount string `xml:"nbPurchaseAccount,omitempty" json:"nbPurchaseAccount,omitempty"`

	ProvisionAttemptsDay string `xml:"provisionAttemptsDay,omitempty" json:"provisionAttemptsDay,omitempty"`

	TxnActivityDay string `xml:"txnActivityDay,omitempty" json:"txnActivityDay,omitempty"`

	TxnActivityYear string `xml:"txnActivityYear,omitempty" json:"txnActivityYear,omitempty"`

	PaymentAccAge string `xml:"paymentAccAge,omitempty" json:"paymentAccAge,omitempty"`

	PaymentAccInd string `xml:"paymentAccInd,omitempty" json:"paymentAccInd,omitempty"`

	ShipAddressUsage string `xml:"shipAddressUsage,omitempty" json:"shipAddressUsage,omitempty"`

	ShipAddressUsageInd string `xml:"shipAddressUsageInd,omitempty" json:"shipAddressUsageInd,omitempty"`

	ShipNameIndicator string `xml:"shipNameIndicator,omitempty" json:"shipNameIndicator,omitempty"`

	SuspiciousAccActivity string `xml:"suspiciousAccActivity,omitempty" json:"suspiciousAccActivity,omitempty"`
}

type MerchantRiskIndicator struct {
	DeliveryEmailAddress string `xml:"deliveryEmailAddress,omitempty" json:"deliveryEmailAddress,omitempty"`

	DeliveryTimeframe string `xml:"deliveryTimeframe,omitempty" json:"deliveryTimeframe,omitempty"`

	GiftCardAmount string `xml:"giftCardAmount,omitempty" json:"giftCardAmount,omitempty"`

	GiftCardCount string `xml:"giftCardCount,omitempty" json:"giftCardCount,omitempty"`

	GiftCardCurr string `xml:"giftCardCurr,omitempty" json:"giftCardCurr,omitempty"`

	PreOrderDate string `xml:"preOrderDate,omitempty" json:"preOrderDate,omitempty"`

	PreOrderPurchaseInd string `xml:"preOrderPurchaseInd,omitempty" json:"preOrderPurchaseInd,omitempty"`

	ReorderItemsInd string `xml:"reorderItemsInd,omitempty" json:"reorderItemsInd,omitempty"`

	ShipIndicator string `xml:"shipIndicator,omitempty" json:"shipIndicator,omitempty"`
}

type RecurringTransaction struct {
	InstallNo string `xml:"installNo,omitempty" json:"installNo,omitempty"`

	Expiry string `xml:"expiry,omitempty" json:"expiry,omitempty"`

	Frequency string `xml:"frequency,omitempty" json:"frequency,omitempty"`
}

type PreviousTransDetails struct {
	AuthData string `xml:"authData,omitempty" json:"authData,omitempty"`

	AuthMethod string `xml:"authMethod,omitempty" json:"authMethod,omitempty"`

	AuthTimestamp string `xml:"authTimestamp,omitempty" json:"authTimestamp,omitempty"`

	AcsID string `xml:"acsID,omitempty" json:"acsID,omitempty"`

	BankTransactionID string `xml:"bankTransactionID,omitempty" json:"bankTransactionID,omitempty"`

	XID string `xml:"XID,omitempty" json:"XID,omitempty"`
}

type Marketplace struct {
	MkpType string `xml:"mkpType,omitempty" json:"mkpType,omitempty"`

	MkpRefMaster *RefMaster `xml:"mkpRefMaster,omitempty" json:"mkpRefMaster,omitempty"`

	MkpSplits *ArrayOfMkpSplit `xml:"mkpSplits,omitempty" json:"mkpSplits,omitempty"`
}

type RefMaster struct {
	MkpBankId string `xml:"mkpBankId,omitempty" json:"mkpBankId,omitempty"`

	MkpPaymentId string `xml:"mkpPaymentId,omitempty" json:"mkpPaymentId,omitempty"`
}

type ArrayOfMkpSplit struct {
	MkpSplit []*MkpSplit `xml:"mkpSplit,omitempty" json:"mkpSplit,omitempty"`
}

type MkpSplit struct {
	XMLName xml.Name `xml:"https://ecomms2s.sella.it/ mkpSplit"`

	MkpSubMerchant string `xml:"mkpSubMerchant,omitempty" json:"mkpSubMerchant,omitempty"`
}

type FullFillmentDetails struct {
	FullFillmentlst *ArrayOfFullfillmentDetail `xml:"FullFillmentlst,omitempty" json:"FullFillmentlst,omitempty"`
}

type ArrayOfFullfillmentDetail struct {
	FullfillmentDetail []*FullfillmentDetail `xml:"FullfillmentDetail,omitempty" json:"FullfillmentDetail,omitempty"`
}

type FullfillmentDetail struct {
	Status string `xml:"Status,omitempty" json:"Status,omitempty"`

	ProductDetails []*ProductDetail `xml:"ProductDetails,omitempty" json:"ProductDetails,omitempty"`

	TrackingCompany string `xml:"TrackingCompany,omitempty" json:"TrackingCompany,omitempty"`

	TrackingNumbers string `xml:"TrackingNumbers,omitempty" json:"TrackingNumbers,omitempty"`

	TrackingUrls string `xml:"TrackingUrls,omitempty" json:"TrackingUrls,omitempty"`

	Message string `xml:"Message,omitempty" json:"Message,omitempty"`

	Receipt string `xml:"Receipt,omitempty" json:"Receipt,omitempty"`
}

type WSs2sSoap interface {
	CallRefundS2S(request *CallRefundS2S) (*CallRefundS2SResponse, error)

	CallRefundS2SContext(ctx context.Context, request *CallRefundS2S) (*CallRefundS2SResponse, error)

	CallReadTrxS2S(request *CallReadTrxS2S) (*CallReadTrxS2SResponse, error)

	CallReadTrxS2SContext(ctx context.Context, request *CallReadTrxS2S) (*CallReadTrxS2SResponse, error)

	CallPagamS2S(request *CallPagamS2S) (*CallPagamS2SResponse, error)

	CallPagamS2SContext(ctx context.Context, request *CallPagamS2S) (*CallPagamS2SResponse, error)

	CallDeleteS2S(request *CallDeleteS2S) (*CallDeleteS2SResponse, error)

	CallDeleteS2SContext(ctx context.Context, request *CallDeleteS2S) (*CallDeleteS2SResponse, error)

	CallSettleS2S(request *CallSettleS2S) (*CallSettleS2SResponse, error)

	CallSettleS2SContext(ctx context.Context, request *CallSettleS2S) (*CallSettleS2SResponse, error)

	CallVerifycardS2S(request *CallVerifycardS2S) (*CallVerifycardS2SResponse, error)

	CallVerifycardS2SContext(ctx context.Context, request *CallVerifycardS2S) (*CallVerifycardS2SResponse, error)

	CallCheckCartaS2S(request *CallCheckCartaS2S) (*CallCheckCartaS2SResponse, error)

	CallCheckCartaS2SContext(ctx context.Context, request *CallCheckCartaS2S) (*CallCheckCartaS2SResponse, error)

	CallRenounce(request *CallRenounce) (*CallRenounceResponse, error)

	CallRenounceContext(ctx context.Context, request *CallRenounce) (*CallRenounceResponse, error)

	CallRequestTokenS2S(request *CallRequestTokenS2S) (*CallRequestTokenS2SResponse, error)

	CallRequestTokenS2SContext(ctx context.Context, request *CallRequestTokenS2S) (*CallRequestTokenS2SResponse, error)

	CallDeleteTokenS2S(request *CallDeleteTokenS2S) (*CallDeleteTokenS2SResponse, error)

	CallDeleteTokenS2SContext(ctx context.Context, request *CallDeleteTokenS2S) (*CallDeleteTokenS2SResponse, error)

	CallUpdateTokenS2S(request *CallUpdateTokenS2S) (*CallUpdateTokenS2SResponse, error)

	CallUpdateTokenS2SContext(ctx context.Context, request *CallUpdateTokenS2S) (*CallUpdateTokenS2SResponse, error)

	CallIdealListS2S(request *CallIdealListS2S) (*CallIdealListS2SResponse, error)

	CallIdealListS2SContext(ctx context.Context, request *CallIdealListS2S) (*CallIdealListS2SResponse, error)

	CallMyBankListS2S(request *CallMyBankListS2S) (*CallMyBankListS2SResponse, error)

	CallMyBankListS2SContext(ctx context.Context, request *CallMyBankListS2S) (*CallMyBankListS2SResponse, error)

	CallUpdateCustomParamS2S(request *CallUpdateCustomParamS2S) (*CallUpdateCustomParamS2SResponse, error)

	CallUpdateCustomParamS2SContext(ctx context.Context, request *CallUpdateCustomParamS2S) (*CallUpdateCustomParamS2SResponse, error)

	CallUpdateOrderS2S(request *CallUpdateOrderS2S) (*CallUpdateOrderS2SResponse, error)

	CallUpdateOrderS2SContext(ctx context.Context, request *CallUpdateOrderS2S) (*CallUpdateOrderS2SResponse, error)

	CallExchangeRateS2S(request *CallExchangeRateS2S) (*CallExchangeRateS2SResponse, error)

	CallExchangeRateS2SContext(ctx context.Context, request *CallExchangeRateS2S) (*CallExchangeRateS2SResponse, error)

	CallDeviceActivation(request *CallDeviceActivation) (*CallDeviceActivationResponse, error)

	CallDeviceActivationContext(ctx context.Context, request *CallDeviceActivation) (*CallDeviceActivationResponse, error)
}

type wSs2sSoap struct {
	Mu     sync.Mutex
	client *soap.Client
}

func NewWSs2sSoap(client *soap.Client) WSs2sSoap {
	return &wSs2sSoap{
		client: client,
	}
}

var WSClient wSs2sSoap

//var ClientSoap *soap.Client

// neeed to be finishi
func InitSoap(ctx context.Context) {
	WSClient = wSs2sSoap{
		Mu: sync.Mutex{},
	}

	WSClient.Mu.Lock()
	defer WSClient.Mu.Unlock()
}

func (service *wSs2sSoap) CallRefundS2SContext(ctx context.Context, request *CallRefundS2S) (*CallRefundS2SResponse, error) {
	response := new(CallRefundS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callRefundS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallRefundS2S(request *CallRefundS2S) (*CallRefundS2SResponse, error) {
	return service.CallRefundS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallReadTrxS2SContext(ctx context.Context, request *CallReadTrxS2S) (*CallReadTrxS2SResponse, error) {
	response := new(CallReadTrxS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callReadTrxS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallReadTrxS2S(request *CallReadTrxS2S) (*CallReadTrxS2SResponse, error) {
	return service.CallReadTrxS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallPagamS2SContext(ctx context.Context, request *CallPagamS2S) (*CallPagamS2SResponse, error) {
	response := new(CallPagamS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callPagamS2S", request, response)
	//err := service.client.CallContext(ctx, "https://sandbox.gestpay.net/callPagamS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallPagamS2S(request *CallPagamS2S) (*CallPagamS2SResponse, error) {
	return service.CallPagamS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallDeleteS2SContext(ctx context.Context, request *CallDeleteS2S) (*CallDeleteS2SResponse, error) {
	response := new(CallDeleteS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callDeleteS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallDeleteS2S(request *CallDeleteS2S) (*CallDeleteS2SResponse, error) {
	return service.CallDeleteS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallSettleS2SContext(ctx context.Context, request *CallSettleS2S) (*CallSettleS2SResponse, error) {
	response := new(CallSettleS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callSettleS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallSettleS2S(request *CallSettleS2S) (*CallSettleS2SResponse, error) {
	return service.CallSettleS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallVerifycardS2SContext(ctx context.Context, request *CallVerifycardS2S) (*CallVerifycardS2SResponse, error) {
	response := new(CallVerifycardS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callVerifycardS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallVerifycardS2S(request *CallVerifycardS2S) (*CallVerifycardS2SResponse, error) {
	return service.CallVerifycardS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallCheckCartaS2SContext(ctx context.Context, request *CallCheckCartaS2S) (*CallCheckCartaS2SResponse, error) {
	response := new(CallCheckCartaS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callCheckCartaS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallCheckCartaS2S(request *CallCheckCartaS2S) (*CallCheckCartaS2SResponse, error) {
	return service.CallCheckCartaS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallRenounceContext(ctx context.Context, request *CallRenounce) (*CallRenounceResponse, error) {
	response := new(CallRenounceResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callRenounce", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallRenounce(request *CallRenounce) (*CallRenounceResponse, error) {
	return service.CallRenounceContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallRequestTokenS2SContext(ctx context.Context, request *CallRequestTokenS2S) (*CallRequestTokenS2SResponse, error) {
	response := new(CallRequestTokenS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallRequestTokenS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallRequestTokenS2S(request *CallRequestTokenS2S) (*CallRequestTokenS2SResponse, error) {
	return service.CallRequestTokenS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallDeleteTokenS2SContext(ctx context.Context, request *CallDeleteTokenS2S) (*CallDeleteTokenS2SResponse, error) {
	response := new(CallDeleteTokenS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callDeleteTokenS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallDeleteTokenS2S(request *CallDeleteTokenS2S) (*CallDeleteTokenS2SResponse, error) {
	return service.CallDeleteTokenS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallUpdateTokenS2SContext(ctx context.Context, request *CallUpdateTokenS2S) (*CallUpdateTokenS2SResponse, error) {
	response := new(CallUpdateTokenS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallUpdateTokenS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallUpdateTokenS2S(request *CallUpdateTokenS2S) (*CallUpdateTokenS2SResponse, error) {
	return service.CallUpdateTokenS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallIdealListS2SContext(ctx context.Context, request *CallIdealListS2S) (*CallIdealListS2SResponse, error) {
	response := new(CallIdealListS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallIdealListS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallIdealListS2S(request *CallIdealListS2S) (*CallIdealListS2SResponse, error) {
	return service.CallIdealListS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallMyBankListS2SContext(ctx context.Context, request *CallMyBankListS2S) (*CallMyBankListS2SResponse, error) {
	response := new(CallMyBankListS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallMyBankListS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallMyBankListS2S(request *CallMyBankListS2S) (*CallMyBankListS2SResponse, error) {
	return service.CallMyBankListS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallUpdateCustomParamS2SContext(ctx context.Context, request *CallUpdateCustomParamS2S) (*CallUpdateCustomParamS2SResponse, error) {
	response := new(CallUpdateCustomParamS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallUpdateCustomParamS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallUpdateCustomParamS2S(request *CallUpdateCustomParamS2S) (*CallUpdateCustomParamS2SResponse, error) {
	return service.CallUpdateCustomParamS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallUpdateOrderS2SContext(ctx context.Context, request *CallUpdateOrderS2S) (*CallUpdateOrderS2SResponse, error) {
	response := new(CallUpdateOrderS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/CallUpdateOrderS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallUpdateOrderS2S(request *CallUpdateOrderS2S) (*CallUpdateOrderS2SResponse, error) {
	return service.CallUpdateOrderS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallExchangeRateS2SContext(ctx context.Context, request *CallExchangeRateS2S) (*CallExchangeRateS2SResponse, error) {
	response := new(CallExchangeRateS2SResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callExchangeRateS2S", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallExchangeRateS2S(request *CallExchangeRateS2S) (*CallExchangeRateS2SResponse, error) {
	return service.CallExchangeRateS2SContext(
		context.Background(),
		request,
	)
}

func (service *wSs2sSoap) CallDeviceActivationContext(ctx context.Context, request *CallDeviceActivation) (*CallDeviceActivationResponse, error) {
	response := new(CallDeviceActivationResponse)
	err := service.client.CallContext(ctx, "https://ecomms2s.sella.it/callDeviceActivation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *wSs2sSoap) CallDeviceActivation(request *CallDeviceActivation) (*CallDeviceActivationResponse, error) {
	return service.CallDeviceActivationContext(
		context.Background(),
		request,
	)
}
