package testhelper

import "net/mail"

func MustParseAddress(address string) mail.Address {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		panic(err)
	}
	return *addr
}

func MustParseAddressPointer(address string) *mail.Address {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		panic(err)
	}
	return addr
}
