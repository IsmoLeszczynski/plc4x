/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableTransaction is an enum
type OpcuaNodeIdServicesVariableTransaction int32

type IOpcuaNodeIdServicesVariableTransaction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_StartTime                 OpcuaNodeIdServicesVariableTransaction = 32287
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_EndTime                   OpcuaNodeIdServicesVariableTransaction = 32288
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Result                    OpcuaNodeIdServicesVariableTransaction = 32289
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedTrustLists        OpcuaNodeIdServicesVariableTransaction = 32290
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedCertificateGroups OpcuaNodeIdServicesVariableTransaction = 32291
	OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Errors                    OpcuaNodeIdServicesVariableTransaction = 32292
)

var OpcuaNodeIdServicesVariableTransactionValues []OpcuaNodeIdServicesVariableTransaction

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTransactionValues = []OpcuaNodeIdServicesVariableTransaction{
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_StartTime,
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_EndTime,
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Result,
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedTrustLists,
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedCertificateGroups,
		OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Errors,
	}
}

func OpcuaNodeIdServicesVariableTransactionByValue(value int32) (enum OpcuaNodeIdServicesVariableTransaction, ok bool) {
	switch value {
	case 32287:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_StartTime, true
	case 32288:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_EndTime, true
	case 32289:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Result, true
	case 32290:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedTrustLists, true
	case 32291:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedCertificateGroups, true
	case 32292:
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Errors, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTransactionByName(value string) (enum OpcuaNodeIdServicesVariableTransaction, ok bool) {
	switch value {
	case "TransactionDiagnosticsType_StartTime":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_StartTime, true
	case "TransactionDiagnosticsType_EndTime":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_EndTime, true
	case "TransactionDiagnosticsType_Result":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Result, true
	case "TransactionDiagnosticsType_AffectedTrustLists":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedTrustLists, true
	case "TransactionDiagnosticsType_AffectedCertificateGroups":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedCertificateGroups, true
	case "TransactionDiagnosticsType_Errors":
		return OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Errors, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTransactionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTransactionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTransaction(structType any) OpcuaNodeIdServicesVariableTransaction {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTransaction {
		if sOpcuaNodeIdServicesVariableTransaction, ok := typ.(OpcuaNodeIdServicesVariableTransaction); ok {
			return sOpcuaNodeIdServicesVariableTransaction
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTransaction) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTransaction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTransactionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTransaction, error) {
	return OpcuaNodeIdServicesVariableTransactionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTransactionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTransaction, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTransaction", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTransaction")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTransactionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTransaction")
		return OpcuaNodeIdServicesVariableTransaction(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTransaction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTransaction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTransaction", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTransaction) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_StartTime:
		return "TransactionDiagnosticsType_StartTime"
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_EndTime:
		return "TransactionDiagnosticsType_EndTime"
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Result:
		return "TransactionDiagnosticsType_Result"
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedTrustLists:
		return "TransactionDiagnosticsType_AffectedTrustLists"
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_AffectedCertificateGroups:
		return "TransactionDiagnosticsType_AffectedCertificateGroups"
	case OpcuaNodeIdServicesVariableTransaction_TransactionDiagnosticsType_Errors:
		return "TransactionDiagnosticsType_Errors"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTransaction) String() string {
	return e.PLC4XEnumName()
}
