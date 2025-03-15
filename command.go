package main

type ICommand interface {
	Execute(args []string)
}
