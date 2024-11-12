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

// StatusRequestBinaryStateDeprecated is the corresponding interface of StatusRequestBinaryStateDeprecated
type StatusRequestBinaryStateDeprecated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	StatusRequest
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// IsStatusRequestBinaryStateDeprecated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusRequestBinaryStateDeprecated()
	// CreateBuilder creates a StatusRequestBinaryStateDeprecatedBuilder
	CreateStatusRequestBinaryStateDeprecatedBuilder() StatusRequestBinaryStateDeprecatedBuilder
}

// _StatusRequestBinaryStateDeprecated is the data-structure of this message
type _StatusRequestBinaryStateDeprecated struct {
	StatusRequestContract
	Application ApplicationIdContainer
	// Reserved Fields
	reservedField0 *byte
	reservedField1 *byte
}

var _ StatusRequestBinaryStateDeprecated = (*_StatusRequestBinaryStateDeprecated)(nil)
var _ StatusRequestRequirements = (*_StatusRequestBinaryStateDeprecated)(nil)

// NewStatusRequestBinaryStateDeprecated factory function for _StatusRequestBinaryStateDeprecated
func NewStatusRequestBinaryStateDeprecated(statusType byte, application ApplicationIdContainer) *_StatusRequestBinaryStateDeprecated {
	_result := &_StatusRequestBinaryStateDeprecated{
		StatusRequestContract: NewStatusRequest(statusType),
		Application:           application,
	}
	_result.StatusRequestContract.(*_StatusRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StatusRequestBinaryStateDeprecatedBuilder is a builder for StatusRequestBinaryStateDeprecated
type StatusRequestBinaryStateDeprecatedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(application ApplicationIdContainer) StatusRequestBinaryStateDeprecatedBuilder
	// WithApplication adds Application (property field)
	WithApplication(ApplicationIdContainer) StatusRequestBinaryStateDeprecatedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() StatusRequestBuilder
	// Build builds the StatusRequestBinaryStateDeprecated or returns an error if something is wrong
	Build() (StatusRequestBinaryStateDeprecated, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StatusRequestBinaryStateDeprecated
}

// NewStatusRequestBinaryStateDeprecatedBuilder() creates a StatusRequestBinaryStateDeprecatedBuilder
func NewStatusRequestBinaryStateDeprecatedBuilder() StatusRequestBinaryStateDeprecatedBuilder {
	return &_StatusRequestBinaryStateDeprecatedBuilder{_StatusRequestBinaryStateDeprecated: new(_StatusRequestBinaryStateDeprecated)}
}

type _StatusRequestBinaryStateDeprecatedBuilder struct {
	*_StatusRequestBinaryStateDeprecated

	parentBuilder *_StatusRequestBuilder

	err *utils.MultiError
}

var _ (StatusRequestBinaryStateDeprecatedBuilder) = (*_StatusRequestBinaryStateDeprecatedBuilder)(nil)

func (b *_StatusRequestBinaryStateDeprecatedBuilder) setParent(contract StatusRequestContract) {
	b.StatusRequestContract = contract
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) WithMandatoryFields(application ApplicationIdContainer) StatusRequestBinaryStateDeprecatedBuilder {
	return b.WithApplication(application)
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) WithApplication(application ApplicationIdContainer) StatusRequestBinaryStateDeprecatedBuilder {
	b.Application = application
	return b
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) Build() (StatusRequestBinaryStateDeprecated, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StatusRequestBinaryStateDeprecated.deepCopy(), nil
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) MustBuild() StatusRequestBinaryStateDeprecated {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) Done() StatusRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewStatusRequestBuilder().(*_StatusRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) buildForStatusRequest() (StatusRequest, error) {
	return b.Build()
}

func (b *_StatusRequestBinaryStateDeprecatedBuilder) DeepCopy() any {
	_copy := b.CreateStatusRequestBinaryStateDeprecatedBuilder().(*_StatusRequestBinaryStateDeprecatedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStatusRequestBinaryStateDeprecatedBuilder creates a StatusRequestBinaryStateDeprecatedBuilder
func (b *_StatusRequestBinaryStateDeprecated) CreateStatusRequestBinaryStateDeprecatedBuilder() StatusRequestBinaryStateDeprecatedBuilder {
	if b == nil {
		return NewStatusRequestBinaryStateDeprecatedBuilder()
	}
	return &_StatusRequestBinaryStateDeprecatedBuilder{_StatusRequestBinaryStateDeprecated: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusRequestBinaryStateDeprecated) GetParent() StatusRequestContract {
	return m.StatusRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequestBinaryStateDeprecated) GetApplication() ApplicationIdContainer {
	return m.Application
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStatusRequestBinaryStateDeprecated(structType any) StatusRequestBinaryStateDeprecated {
	if casted, ok := structType.(StatusRequestBinaryStateDeprecated); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequestBinaryStateDeprecated); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequestBinaryStateDeprecated) GetTypeName() string {
	return "StatusRequestBinaryStateDeprecated"
}

func (m *_StatusRequestBinaryStateDeprecated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.StatusRequestContract.(*_StatusRequest).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_StatusRequestBinaryStateDeprecated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StatusRequestBinaryStateDeprecated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_StatusRequest) (__statusRequestBinaryStateDeprecated StatusRequestBinaryStateDeprecated, err error) {
	m.StatusRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequestBinaryStateDeprecated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequestBinaryStateDeprecated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0xFA))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	if closeErr := readBuffer.CloseContext("StatusRequestBinaryStateDeprecated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequestBinaryStateDeprecated")
	}

	return m, nil
}

func (m *_StatusRequestBinaryStateDeprecated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusRequestBinaryStateDeprecated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestBinaryStateDeprecated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusRequestBinaryStateDeprecated")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0xFA), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x00), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("StatusRequestBinaryStateDeprecated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusRequestBinaryStateDeprecated")
		}
		return nil
	}
	return m.StatusRequestContract.(*_StatusRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusRequestBinaryStateDeprecated) IsStatusRequestBinaryStateDeprecated() {}

func (m *_StatusRequestBinaryStateDeprecated) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StatusRequestBinaryStateDeprecated) deepCopy() *_StatusRequestBinaryStateDeprecated {
	if m == nil {
		return nil
	}
	_StatusRequestBinaryStateDeprecatedCopy := &_StatusRequestBinaryStateDeprecated{
		m.StatusRequestContract.(*_StatusRequest).deepCopy(),
		m.Application,
		m.reservedField0,
		m.reservedField1,
	}
	m.StatusRequestContract.(*_StatusRequest)._SubType = m
	return _StatusRequestBinaryStateDeprecatedCopy
}

func (m *_StatusRequestBinaryStateDeprecated) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
