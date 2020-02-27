package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"os"

	"strings"
)

func BuildPayloadByABI(contractAbi, funcName string, args ...interface{}) ([]byte, error) {
	abiObj, _ := abi.JSON(strings.NewReader(contractAbi))
	m, _ := abiObj.Methods[funcName]
	funcSig := m.ID()
	d, err := m.Inputs.PackValues(args)
	if err != nil {
		return nil, err
	}
	return append(funcSig, d...), nil
}

func LoadABI(abiFileAbsPath string) (string, error) {
	fd, err := os.Open(abiFileAbsPath)
	defer fd.Close()

	if err != nil {
		return "", nil
	}

	d, err := ioutil.ReadAll(fd)
	return string(d), nil
}

func BuildPayloadByABIFile(abiFilePath, funcName string, args ...interface{}) ([]byte, error) {
	contractAbi, err := LoadABI(abiFilePath)
	if err != nil {
		return nil, err
	}
	return BuildPayloadByABI(contractAbi, funcName, args...)
}
