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

// BACnetConstructedDataMembers is the corresponding interface of BACnetConstructedDataMembers
type BACnetConstructedDataMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMembers returns Members (property field)
	GetMembers() []BACnetDeviceObjectReference
	// IsBACnetConstructedDataMembers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMembers()
	// CreateBuilder creates a BACnetConstructedDataMembersBuilder
	CreateBACnetConstructedDataMembersBuilder() BACnetConstructedDataMembersBuilder
}

// _BACnetConstructedDataMembers is the data-structure of this message
type _BACnetConstructedDataMembers struct {
	BACnetConstructedDataContract
	Members []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataMembers = (*_BACnetConstructedDataMembers)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMembers)(nil)

// NewBACnetConstructedDataMembers factory function for _BACnetConstructedDataMembers
func NewBACnetConstructedDataMembers(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, members []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMembers {
	_result := &_BACnetConstructedDataMembers{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Members:                       members,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMembersBuilder is a builder for BACnetConstructedDataMembers
type BACnetConstructedDataMembersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(members []BACnetDeviceObjectReference) BACnetConstructedDataMembersBuilder
	// WithMembers adds Members (property field)
	WithMembers(...BACnetDeviceObjectReference) BACnetConstructedDataMembersBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataMembers or returns an error if something is wrong
	Build() (BACnetConstructedDataMembers, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMembers
}

// NewBACnetConstructedDataMembersBuilder() creates a BACnetConstructedDataMembersBuilder
func NewBACnetConstructedDataMembersBuilder() BACnetConstructedDataMembersBuilder {
	return &_BACnetConstructedDataMembersBuilder{_BACnetConstructedDataMembers: new(_BACnetConstructedDataMembers)}
}

type _BACnetConstructedDataMembersBuilder struct {
	*_BACnetConstructedDataMembers

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMembersBuilder) = (*_BACnetConstructedDataMembersBuilder)(nil)

func (b *_BACnetConstructedDataMembersBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataMembersBuilder) WithMandatoryFields(members []BACnetDeviceObjectReference) BACnetConstructedDataMembersBuilder {
	return b.WithMembers(members...)
}

func (b *_BACnetConstructedDataMembersBuilder) WithMembers(members ...BACnetDeviceObjectReference) BACnetConstructedDataMembersBuilder {
	b.Members = members
	return b
}

func (b *_BACnetConstructedDataMembersBuilder) Build() (BACnetConstructedDataMembers, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMembers.deepCopy(), nil
}

func (b *_BACnetConstructedDataMembersBuilder) MustBuild() BACnetConstructedDataMembers {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataMembersBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMembersBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMembersBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMembersBuilder().(*_BACnetConstructedDataMembersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMembersBuilder creates a BACnetConstructedDataMembersBuilder
func (b *_BACnetConstructedDataMembers) CreateBACnetConstructedDataMembersBuilder() BACnetConstructedDataMembersBuilder {
	if b == nil {
		return NewBACnetConstructedDataMembersBuilder()
	}
	return &_BACnetConstructedDataMembersBuilder{_BACnetConstructedDataMembers: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMembers) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMembers) GetMembers() []BACnetDeviceObjectReference {
	return m.Members
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMembers(structType any) BACnetConstructedDataMembers {
	if casted, ok := structType.(BACnetConstructedDataMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMembers) GetTypeName() string {
	return "BACnetConstructedDataMembers"
}

func (m *_BACnetConstructedDataMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMembers) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMembers BACnetConstructedDataMembers, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	members, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "members", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'members' field"))
	}
	m.Members = members

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMembers")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMembers")
		}

		if err := WriteComplexTypeArrayField(ctx, "members", m.GetMembers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'members' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMembers")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMembers) IsBACnetConstructedDataMembers() {}

func (m *_BACnetConstructedDataMembers) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMembers) deepCopy() *_BACnetConstructedDataMembers {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMembersCopy := &_BACnetConstructedDataMembers{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDeviceObjectReference, BACnetDeviceObjectReference](m.Members),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMembersCopy
}

func (m *_BACnetConstructedDataMembers) String() string {
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
