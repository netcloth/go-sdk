package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func BuildPayloadByABI(contractAbi, funcName string, args ...interface{}) ([]byte, error) {
	abiObj, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}
	m, ok := abiObj.Methods[funcName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("function name [%s] not found", funcName))
	}
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

func UnpackValuesByABI(contractAbi, funcName string, d []byte) ([]interface{}, error) {
	abiObj, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}
	m, ok := abiObj.Methods[funcName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("function name [%s] not found", funcName))
	}
	return m.Outputs.UnpackValues(d)
}

func UnpackValuesByABI2(contractAbi, funcName string, d []byte, isInput bool) ([]interface{}, error) {
	abiObj, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}
	m, ok := abiObj.Methods[funcName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("function name [%s] not found", funcName))
	}
	if isInput {
		return m.Inputs.UnpackValues(d)
	}
	return m.Outputs.UnpackValues(d)
}

func UnpackEventValuesByABI(contractAbi, eventName string, d []byte) ([]interface{}, error) {
	abiObj, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	m, ok := abiObj.Events[eventName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("event name [%s] not found", eventName))
	}

	return m.Inputs.UnpackValues(d)
}

func UnpackValuesByABIFile(abiFilePath, funcName string, d []byte) ([]interface{}, error) {
	contractAbi, err := LoadABI(abiFilePath)
	if err != nil {
		return nil, err
	}

	return UnpackValuesByABI(contractAbi, funcName, d)
}

func UnpackValuesByABIFile2(abiFilePath, funcName string, d []byte, isInput bool) ([]interface{}, error) {
	contractAbi, err := LoadABI(abiFilePath)
	if err != nil {
		return nil, err
	}

	return UnpackValuesByABI2(contractAbi, funcName, d, isInput)
}

func UnpackEventValuesByABIFile(abiFilePath, funcName string, d []byte) ([]interface{}, error) {
	contractAbi, err := LoadABI(abiFilePath)
	if err != nil {
		return nil, err
	}

	return UnpackEventValuesByABI(contractAbi, funcName, d)
}
