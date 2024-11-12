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

// BACnetConstructedDataDoorStatus is the corresponding interface of BACnetConstructedDataDoorStatus
type BACnetConstructedDataDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorStatusTagged
	// IsBACnetConstructedDataDoorStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDoorStatus()
	// CreateBuilder creates a BACnetConstructedDataDoorStatusBuilder
	CreateBACnetConstructedDataDoorStatusBuilder() BACnetConstructedDataDoorStatusBuilder
}

// _BACnetConstructedDataDoorStatus is the data-structure of this message
type _BACnetConstructedDataDoorStatus struct {
	BACnetConstructedDataContract
	DoorStatus BACnetDoorStatusTagged
}

var _ BACnetConstructedDataDoorStatus = (*_BACnetConstructedDataDoorStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDoorStatus)(nil)

// NewBACnetConstructedDataDoorStatus factory function for _BACnetConstructedDataDoorStatus
func NewBACnetConstructedDataDoorStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doorStatus BACnetDoorStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorStatus {
	if doorStatus == nil {
		panic("doorStatus of type BACnetDoorStatusTagged for BACnetConstructedDataDoorStatus must not be nil")
	}
	_result := &_BACnetConstructedDataDoorStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoorStatus:                    doorStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDoorStatusBuilder is a builder for BACnetConstructedDataDoorStatus
type BACnetConstructedDataDoorStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doorStatus BACnetDoorStatusTagged) BACnetConstructedDataDoorStatusBuilder
	// WithDoorStatus adds DoorStatus (property field)
	WithDoorStatus(BACnetDoorStatusTagged) BACnetConstructedDataDoorStatusBuilder
	// WithDoorStatusBuilder adds DoorStatus (property field) which is build by the builder
	WithDoorStatusBuilder(func(BACnetDoorStatusTaggedBuilder) BACnetDoorStatusTaggedBuilder) BACnetConstructedDataDoorStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDoorStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataDoorStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDoorStatus
}

// NewBACnetConstructedDataDoorStatusBuilder() creates a BACnetConstructedDataDoorStatusBuilder
func NewBACnetConstructedDataDoorStatusBuilder() BACnetConstructedDataDoorStatusBuilder {
	return &_BACnetConstructedDataDoorStatusBuilder{_BACnetConstructedDataDoorStatus: new(_BACnetConstructedDataDoorStatus)}
}

type _BACnetConstructedDataDoorStatusBuilder struct {
	*_BACnetConstructedDataDoorStatus

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDoorStatusBuilder) = (*_BACnetConstructedDataDoorStatusBuilder)(nil)

func (b *_BACnetConstructedDataDoorStatusBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDoorStatusBuilder) WithMandatoryFields(doorStatus BACnetDoorStatusTagged) BACnetConstructedDataDoorStatusBuilder {
	return b.WithDoorStatus(doorStatus)
}

func (b *_BACnetConstructedDataDoorStatusBuilder) WithDoorStatus(doorStatus BACnetDoorStatusTagged) BACnetConstructedDataDoorStatusBuilder {
	b.DoorStatus = doorStatus
	return b
}

func (b *_BACnetConstructedDataDoorStatusBuilder) WithDoorStatusBuilder(builderSupplier func(BACnetDoorStatusTaggedBuilder) BACnetDoorStatusTaggedBuilder) BACnetConstructedDataDoorStatusBuilder {
	builder := builderSupplier(b.DoorStatus.CreateBACnetDoorStatusTaggedBuilder())
	var err error
	b.DoorStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDoorStatusTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDoorStatusBuilder) Build() (BACnetConstructedDataDoorStatus, error) {
	if b.DoorStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doorStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDoorStatus.deepCopy(), nil
}

func (b *_BACnetConstructedDataDoorStatusBuilder) MustBuild() BACnetConstructedDataDoorStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDoorStatusBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDoorStatusBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDoorStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDoorStatusBuilder().(*_BACnetConstructedDataDoorStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDoorStatusBuilder creates a BACnetConstructedDataDoorStatusBuilder
func (b *_BACnetConstructedDataDoorStatus) CreateBACnetConstructedDataDoorStatusBuilder() BACnetConstructedDataDoorStatusBuilder {
	if b == nil {
		return NewBACnetConstructedDataDoorStatusBuilder()
	}
	return &_BACnetConstructedDataDoorStatusBuilder{_BACnetConstructedDataDoorStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorStatus) GetActualValue() BACnetDoorStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorStatusTagged(m.GetDoorStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorStatus(structType any) BACnetConstructedDataDoorStatus {
	if casted, ok := structType.(BACnetConstructedDataDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorStatus) GetTypeName() string {
	return "BACnetConstructedDataDoorStatus"
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoorStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDoorStatus BACnetConstructedDataDoorStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorStatus, err := ReadSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", ReadComplex[BACnetDoorStatusTagged](BACnetDoorStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorStatus' field"))
	}
	m.DoorStatus = doorStatus

	actualValue, err := ReadVirtualField[BACnetDoorStatusTagged](ctx, "actualValue", (*BACnetDoorStatusTagged)(nil), doorStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorStatus")
		}

		if err := WriteSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", m.GetDoorStatus(), WriteComplex[BACnetDoorStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorStatus) IsBACnetConstructedDataDoorStatus() {}

func (m *_BACnetConstructedDataDoorStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDoorStatus) deepCopy() *_BACnetConstructedDataDoorStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDoorStatusCopy := &_BACnetConstructedDataDoorStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDoorStatusTagged](m.DoorStatus),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDoorStatusCopy
}

func (m *_BACnetConstructedDataDoorStatus) String() string {
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
