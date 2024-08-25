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

// ServiceCounterDataType is the corresponding interface of ServiceCounterDataType
type ServiceCounterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTotalCount returns TotalCount (property field)
	GetTotalCount() uint32
	// GetErrorCount returns ErrorCount (property field)
	GetErrorCount() uint32
}

// ServiceCounterDataTypeExactly can be used when we want exactly this type and not a type which fulfills ServiceCounterDataType.
// This is useful for switch cases.
type ServiceCounterDataTypeExactly interface {
	ServiceCounterDataType
	isServiceCounterDataType() bool
}

// _ServiceCounterDataType is the data-structure of this message
type _ServiceCounterDataType struct {
	*_ExtensionObjectDefinition
	TotalCount uint32
	ErrorCount uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServiceCounterDataType) GetIdentifier() string {
	return "873"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServiceCounterDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ServiceCounterDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServiceCounterDataType) GetTotalCount() uint32 {
	return m.TotalCount
}

func (m *_ServiceCounterDataType) GetErrorCount() uint32 {
	return m.ErrorCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServiceCounterDataType factory function for _ServiceCounterDataType
func NewServiceCounterDataType(totalCount uint32, errorCount uint32) *_ServiceCounterDataType {
	_result := &_ServiceCounterDataType{
		TotalCount:                 totalCount,
		ErrorCount:                 errorCount,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServiceCounterDataType(structType any) ServiceCounterDataType {
	if casted, ok := structType.(ServiceCounterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServiceCounterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServiceCounterDataType) GetTypeName() string {
	return "ServiceCounterDataType"
}

func (m *_ServiceCounterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (totalCount)
	lengthInBits += 32

	// Simple field (errorCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ServiceCounterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServiceCounterDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ServiceCounterDataType, error) {
	return ServiceCounterDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ServiceCounterDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ServiceCounterDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ServiceCounterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServiceCounterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (totalCount)
	_totalCount, _totalCountErr := readBuffer.ReadUint32("totalCount", 32)
	if _totalCountErr != nil {
		return nil, errors.Wrap(_totalCountErr, "Error parsing 'totalCount' field of ServiceCounterDataType")
	}
	totalCount := _totalCount

	// Simple Field (errorCount)
	_errorCount, _errorCountErr := readBuffer.ReadUint32("errorCount", 32)
	if _errorCountErr != nil {
		return nil, errors.Wrap(_errorCountErr, "Error parsing 'errorCount' field of ServiceCounterDataType")
	}
	errorCount := _errorCount

	if closeErr := readBuffer.CloseContext("ServiceCounterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServiceCounterDataType")
	}

	// Create a partially initialized instance
	_child := &_ServiceCounterDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		TotalCount:                 totalCount,
		ErrorCount:                 errorCount,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ServiceCounterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServiceCounterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServiceCounterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServiceCounterDataType")
		}

		// Simple Field (totalCount)
		totalCount := uint32(m.GetTotalCount())
		_totalCountErr := writeBuffer.WriteUint32("totalCount", 32, uint32((totalCount)))
		if _totalCountErr != nil {
			return errors.Wrap(_totalCountErr, "Error serializing 'totalCount' field")
		}

		// Simple Field (errorCount)
		errorCount := uint32(m.GetErrorCount())
		_errorCountErr := writeBuffer.WriteUint32("errorCount", 32, uint32((errorCount)))
		if _errorCountErr != nil {
			return errors.Wrap(_errorCountErr, "Error serializing 'errorCount' field")
		}

		if popErr := writeBuffer.PopContext("ServiceCounterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServiceCounterDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServiceCounterDataType) isServiceCounterDataType() bool {
	return true
}

func (m *_ServiceCounterDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
