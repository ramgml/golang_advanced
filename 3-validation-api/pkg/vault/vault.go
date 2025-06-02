package vault

import (
	"encoding/json"
	"fmt"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Account struct {
	Email string
	Key string
}

func NewAccount(email string, key string) *Account {
	newAccount := &Account{
		Email: email,
		Key: key,
	}
	return newAccount
}

type Vault struct {
	Accounts []Account `json:"accounts"`
}

type VaultWithDb struct {
	Vault
	Db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
			},
			Db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println("Не удалось разобрать файл")
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
			},
			Db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		Db: db,
	}
}

func (vault *VaultWithDb) GetAccountByEmail(email string) *Account {
	for _, account := range vault.Accounts {
		if account.Email == email {
			return &account
		}
	}
	return nil
}

func (vault *VaultWithDb) GetAccountByKey(key string) *Account {
	for _, account := range vault.Accounts {
		if account.Key == key {
			return &account
		}
	}
	return nil
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.Save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vault *VaultWithDb) Save() {
	data, err := vault.Vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать")
	}
	vault.Db.Write(data)
}