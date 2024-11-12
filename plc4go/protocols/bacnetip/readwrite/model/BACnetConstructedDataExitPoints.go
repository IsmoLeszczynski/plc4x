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

// BACnetConstructedDataExitPoints is the corresponding interface of BACnetConstructedDataExitPoints
type BACnetConstructedDataExitPoints interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetExitPoints returns ExitPoints (property field)
	GetExitPoints() []BACnetDeviceObjectReference
	// IsBACnetConstructedDataExitPoints is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataExitPoints()
	// CreateBuilder creates a BACnetConstructedDataExitPointsBuilder
	CreateBACnetConstructedDataExitPointsBuilder() BACnetConstructedDataExitPointsBuilder
}

// _BACnetConstructedDataExitPoints is the data-structure of this message
type _BACnetConstructedDataExitPoints struct {
	BACnetConstructedDataContract
	ExitPoints []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataExitPoints = (*_BACnetConstructedDataExitPoints)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataExitPoints)(nil)

// NewBACnetConstructedDataExitPoints factory function for _BACnetConstructedDataExitPoints
func NewBACnetConstructedDataExitPoints(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, exitPoints []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExitPoints {
	_result := &_BACnetConstructedDataExitPoints{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ExitPoints:                    exitPoints,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataExitPointsBuilder is a builder for BACnetConstructedDataExitPoints
type BACnetConstructedDataExitPointsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(exitPoints []BACnetDeviceObjectReference) BACnetConstructedDataExitPointsBuilder
	// WithExitPoints adds ExitPoints (property field)
	WithExitPoints(...BACnetDeviceObjectReference) BACnetConstructedDataExitPointsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataExitPoints or returns an error if something is wrong
	Build() (BACnetConstructedDataExitPoints, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataExitPoints
}

// NewBACnetConstructedDataExitPointsBuilder() creates a BACnetConstructedDataExitPointsBuilder
func NewBACnetConstructedDataExitPointsBuilder() BACnetConstructedDataExitPointsBuilder {
	return &_BACnetConstructedDataExitPointsBuilder{_BACnetConstructedDataExitPoints: new(_BACnetConstructedDataExitPoints)}
}

type _BACnetConstructedDataExitPointsBuilder struct {
	*_BACnetConstructedDataExitPoints

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataExitPointsBuilder) = (*_BACnetConstructedDataExitPointsBuilder)(nil)

func (b *_BACnetConstructedDataExitPointsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataExitPointsBuilder) WithMandatoryFields(exitPoints []BACnetDeviceObjectReference) BACnetConstructedDataExitPointsBuilder {
	return b.WithExitPoints(exitPoints...)
}

func (b *_BACnetConstructedDataExitPointsBuilder) WithExitPoints(exitPoints ...BACnetDeviceObjectReference) BACnetConstructedDataExitPointsBuilder {
	b.ExitPoints = exitPoints
	return b
}

func (b *_BACnetConstructedDataExitPointsBuilder) Build() (BACnetConstructedDataExitPoints, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataExitPoints.deepCopy(), nil
}

func (b *_BACnetConstructedDataExitPointsBuilder) MustBuild() BACnetConstructedDataExitPoints {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataExitPointsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataExitPointsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataExitPointsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataExitPointsBuilder().(*_BACnetConstructedDataExitPointsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataExitPointsBuilder creates a BACnetConstructedDataExitPointsBuilder
func (b *_BACnetConstructedDataExitPoints) CreateBACnetConstructedDataExitPointsBuilder() BACnetConstructedDataExitPointsBuilder {
	if b == nil {
		return NewBACnetConstructedDataExitPointsBuilder()
	}
	return &_BACnetConstructedDataExitPointsBuilder{_BACnetConstructedDataExitPoints: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExitPoints) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExitPoints) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXIT_POINTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExitPoints) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExitPoints) GetExitPoints() []BACnetDeviceObjectReference {
	return m.ExitPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExitPoints(structType any) BACnetConstructedDataExitPoints {
	if casted, ok := structType.(BACnetConstructedDataExitPoints); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExitPoints); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExitPoints) GetTypeName() string {
	return "BACnetConstructedDataExitPoints"
}

func (m *_BACnetConstructedDataExitPoints) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.ExitPoints) > 0 {
		for _, element := range m.ExitPoints {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataExitPoints) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataExitPoints) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataExitPoints BACnetConstructedDataExitPoints, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExitPoints"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExitPoints")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	exitPoints, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "exitPoints", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exitPoints' field"))
	}
	m.ExitPoints = exitPoints

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExitPoints"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExitPoints")
	}

	return m, nil
}

func (m *_BACnetConstructedDataExitPoints) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExitPoints) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExitPoints"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExitPoints")
		}

		if err := WriteComplexTypeArrayField(ctx, "exitPoints", m.GetExitPoints(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'exitPoints' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExitPoints"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExitPoints")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExitPoints) IsBACnetConstructedDataExitPoints() {}

func (m *_BACnetConstructedDataExitPoints) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataExitPoints) deepCopy() *_BACnetConstructedDataExitPoints {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataExitPointsCopy := &_BACnetConstructedDataExitPoints{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDeviceObjectReference, BACnetDeviceObjectReference](m.ExitPoints),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataExitPointsCopy
}

func (m *_BACnetConstructedDataExitPoints) String() string {
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
