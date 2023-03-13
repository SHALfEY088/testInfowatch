package main

import (
	s "github.com/SHALfEY088/testInfowatch/countSymbols"
	f "github.com/SHALfEY088/testInfowatch/createFiles"
)

func main() {
	f.CreateFiles()
	s.CountSymbolsInFiles()
}
