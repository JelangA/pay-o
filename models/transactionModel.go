package models

import "gorm.io/gorm"

type Transaction struct {
   gorm.Model
   ID               int    `gorm:"primaryKey"`
   Total            int
   Status           string
   CustomerName     string
   CustomerEmail    string
   SnapToken        string
   SnapRedirectURL  string
   PaymentMethod    string
   TransactionItem  string
}
