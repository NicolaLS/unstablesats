package models

import "time"

// BlinkWebhookPayload represents the payload received from a Blink webhook.
type BlinkWebhookPayload struct {
	AccountID   string      `json:"accountId"`
	EventType   string      `json:"eventType"`
	WalletID    string      `json:"walletId"`
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	CreatedAt               time.Time              `json:"createdAt"`
	ID                      string                 `json:"id"`
	InitiationVia           InitiationVia          `json:"initiationVia"`
	Memo                    *string                `json:"memo"`
	SettlementAmount        int                    `json:"settlementAmount"`
	SettlementCurrency      string                 `json:"settlementCurrency"`
	SettlementDisplayAmount string                 `json:"settlementDisplayAmount"`
	SettlementDisplayFee    string                 `json:"settlementDisplayFee"`
	SettlementDisplayPrice  SettlementDisplayPrice `json:"settlementDisplayPrice"`
	SettlementFee           int                    `json:"settlementFee"`
	SettlementVia           SettlementVia          `json:"settlementVia"`
	Status                  string                 `json:"status"`
	WalletID                string                 `json:"walletId"`
}

type InitiationVia struct {
	PaymentHash string `json:"paymentHash"`
	Pubkey      string `json:"pubkey"`
	Type        string `json:"type"`
}

type SettlementDisplayPrice struct {
	Base            string `json:"base"`
	DisplayCurrency string `json:"displayCurrency"`
	Offset          string `json:"offset"`
	WalletCurrency  string `json:"walletCurrency"`
}

type SettlementVia struct {
	Type string `json:"type"`
}
