package main

import (
	"errors"
	"fmt"
)

type Account struct {
	numBank  string
	banlance float64
}
type Accounts struct {
	accs map[string]*Account
}

func (a *Accounts) add(numBank string, num float64) {
	account, ok := a.accs[numBank]
	if !ok {
		a.accs[numBank] = &Account{numBank, num}
	} else {
		account.banlance += num
	}
}
func (a *Accounts) sub(numBank string, num float64) (float64, error) {
	account, ok := a.accs[numBank]
	if !ok {
		return 0, errors.New("账户不存在")
	} else {
		if account.banlance < num {
			return account.banlance, errors.New("余额不足")
		} else {
			account.banlance -= num
			return account.banlance, nil
		}

	}
}

func (a *Accounts) findBanlance(numBank string) (float64, error) {
	acc, ok := a.accs[numBank]
	if ok {
		return acc.banlance, nil
	} else {
		return 0, fmt.Errorf("no account found for %s", numBank)
	}
}
func main() {
	acc := Accounts{make(map[string]*Account)}
	acc.add("1", 100)
	acc.add("2", 100)
	bal, err := acc.findBanlance("1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1", bal)
	}
	_, err = acc.sub("1", 10)
	if err != nil {
		fmt.Println(err)
	}
	bal, err = acc.findBanlance("1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1", bal)
	}
}
