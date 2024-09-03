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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ContentFilterElementResult is the corresponding interface of ContentFilterElementResult
type ContentFilterElementResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfOperandStatusCodes returns NoOfOperandStatusCodes (property field)
	GetNoOfOperandStatusCodes() int32
	// GetOperandStatusCodes returns OperandStatusCodes (property field)
	GetOperandStatusCodes() []StatusCode
	// GetNoOfOperandDiagnosticInfos returns NoOfOperandDiagnosticInfos (property field)
	GetNoOfOperandDiagnosticInfos() int32
	// GetOperandDiagnosticInfos returns OperandDiagnosticInfos (property field)
	GetOperandDiagnosticInfos() []DiagnosticInfo
	// IsContentFilterElementResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsContentFilterElementResult()
}

// _ContentFilterElementResult is the data-structure of this message
type _ContentFilterElementResult struct {
	ExtensionObjectDefinitionContract
	StatusCode                 StatusCode
	NoOfOperandStatusCodes     int32
	OperandStatusCodes         []StatusCode
	NoOfOperandDiagnosticInfos int32
	OperandDiagnosticInfos     []DiagnosticInfo
}

var _ ContentFilterElementResult = (*_ContentFilterElementResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ContentFilterElementResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilterElementResult) GetIdentifier() string {
	return "606"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilterElementResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilterElementResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_ContentFilterElementResult) GetNoOfOperandStatusCodes() int32 {
	return m.NoOfOperandStatusCodes
}

func (m *_ContentFilterElementResult) GetOperandStatusCodes() []StatusCode {
	return m.OperandStatusCodes
}

func (m *_ContentFilterElementResult) GetNoOfOperandDiagnosticInfos() int32 {
	return m.NoOfOperandDiagnosticInfos
}

func (m *_ContentFilterElementResult) GetOperandDiagnosticInfos() []DiagnosticInfo {
	return m.OperandDiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewContentFilterElementResult factory function for _ContentFilterElementResult
func NewContentFilterElementResult(statusCode StatusCode, noOfOperandStatusCodes int32, operandStatusCodes []StatusCode, noOfOperandDiagnosticInfos int32, operandDiagnosticInfos []DiagnosticInfo) *_ContentFilterElementResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for ContentFilterElementResult must not be nil")
	}
	_result := &_ContentFilterElementResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfOperandStatusCodes:            noOfOperandStatusCodes,
		OperandStatusCodes:                operandStatusCodes,
		NoOfOperandDiagnosticInfos:        noOfOperandDiagnosticInfos,
		OperandDiagnosticInfos:            operandDiagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastContentFilterElementResult(structType any) ContentFilterElementResult {
	if casted, ok := structType.(ContentFilterElementResult); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilterElementResult); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilterElementResult) GetTypeName() string {
	return "ContentFilterElementResult"
}

func (m *_ContentFilterElementResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfOperandStatusCodes)
	lengthInBits += 32

	// Array field
	if len(m.OperandStatusCodes) > 0 {
		for _curItem, element := range m.OperandStatusCodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperandStatusCodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfOperandDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.OperandDiagnosticInfos) > 0 {
		for _curItem, element := range m.OperandDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperandDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilterElementResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ContentFilterElementResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__contentFilterElementResult ContentFilterElementResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContentFilterElementResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilterElementResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfOperandStatusCodes, err := ReadSimpleField(ctx, "noOfOperandStatusCodes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOperandStatusCodes' field"))
	}
	m.NoOfOperandStatusCodes = noOfOperandStatusCodes

	operandStatusCodes, err := ReadCountArrayField[StatusCode](ctx, "operandStatusCodes", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfOperandStatusCodes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operandStatusCodes' field"))
	}
	m.OperandStatusCodes = operandStatusCodes

	noOfOperandDiagnosticInfos, err := ReadSimpleField(ctx, "noOfOperandDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOperandDiagnosticInfos' field"))
	}
	m.NoOfOperandDiagnosticInfos = noOfOperandDiagnosticInfos

	operandDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "operandDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfOperandDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operandDiagnosticInfos' field"))
	}
	m.OperandDiagnosticInfos = operandDiagnosticInfos

	if closeErr := readBuffer.CloseContext("ContentFilterElementResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilterElementResult")
	}

	return m, nil
}

func (m *_ContentFilterElementResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilterElementResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilterElementResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilterElementResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOperandStatusCodes", m.GetNoOfOperandStatusCodes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOperandStatusCodes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "operandStatusCodes", m.GetOperandStatusCodes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'operandStatusCodes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOperandDiagnosticInfos", m.GetNoOfOperandDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOperandDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "operandDiagnosticInfos", m.GetOperandDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'operandDiagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ContentFilterElementResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilterElementResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilterElementResult) IsContentFilterElementResult() {}

func (m *_ContentFilterElementResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
