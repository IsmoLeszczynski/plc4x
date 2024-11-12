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

// BACnetConstructedDataNodeType is the corresponding interface of BACnetConstructedDataNodeType
type BACnetConstructedDataNodeType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNodeType returns NodeType (property field)
	GetNodeType() BACnetNodeTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNodeTypeTagged
	// IsBACnetConstructedDataNodeType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNodeType()
	// CreateBuilder creates a BACnetConstructedDataNodeTypeBuilder
	CreateBACnetConstructedDataNodeTypeBuilder() BACnetConstructedDataNodeTypeBuilder
}

// _BACnetConstructedDataNodeType is the data-structure of this message
type _BACnetConstructedDataNodeType struct {
	BACnetConstructedDataContract
	NodeType BACnetNodeTypeTagged
}

var _ BACnetConstructedDataNodeType = (*_BACnetConstructedDataNodeType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNodeType)(nil)

// NewBACnetConstructedDataNodeType factory function for _BACnetConstructedDataNodeType
func NewBACnetConstructedDataNodeType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, nodeType BACnetNodeTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNodeType {
	if nodeType == nil {
		panic("nodeType of type BACnetNodeTypeTagged for BACnetConstructedDataNodeType must not be nil")
	}
	_result := &_BACnetConstructedDataNodeType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NodeType:                      nodeType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNodeTypeBuilder is a builder for BACnetConstructedDataNodeType
type BACnetConstructedDataNodeTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeType BACnetNodeTypeTagged) BACnetConstructedDataNodeTypeBuilder
	// WithNodeType adds NodeType (property field)
	WithNodeType(BACnetNodeTypeTagged) BACnetConstructedDataNodeTypeBuilder
	// WithNodeTypeBuilder adds NodeType (property field) which is build by the builder
	WithNodeTypeBuilder(func(BACnetNodeTypeTaggedBuilder) BACnetNodeTypeTaggedBuilder) BACnetConstructedDataNodeTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataNodeType or returns an error if something is wrong
	Build() (BACnetConstructedDataNodeType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNodeType
}

// NewBACnetConstructedDataNodeTypeBuilder() creates a BACnetConstructedDataNodeTypeBuilder
func NewBACnetConstructedDataNodeTypeBuilder() BACnetConstructedDataNodeTypeBuilder {
	return &_BACnetConstructedDataNodeTypeBuilder{_BACnetConstructedDataNodeType: new(_BACnetConstructedDataNodeType)}
}

type _BACnetConstructedDataNodeTypeBuilder struct {
	*_BACnetConstructedDataNodeType

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNodeTypeBuilder) = (*_BACnetConstructedDataNodeTypeBuilder)(nil)

func (b *_BACnetConstructedDataNodeTypeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNodeTypeBuilder) WithMandatoryFields(nodeType BACnetNodeTypeTagged) BACnetConstructedDataNodeTypeBuilder {
	return b.WithNodeType(nodeType)
}

func (b *_BACnetConstructedDataNodeTypeBuilder) WithNodeType(nodeType BACnetNodeTypeTagged) BACnetConstructedDataNodeTypeBuilder {
	b.NodeType = nodeType
	return b
}

func (b *_BACnetConstructedDataNodeTypeBuilder) WithNodeTypeBuilder(builderSupplier func(BACnetNodeTypeTaggedBuilder) BACnetNodeTypeTaggedBuilder) BACnetConstructedDataNodeTypeBuilder {
	builder := builderSupplier(b.NodeType.CreateBACnetNodeTypeTaggedBuilder())
	var err error
	b.NodeType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNodeTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNodeTypeBuilder) Build() (BACnetConstructedDataNodeType, error) {
	if b.NodeType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNodeType.deepCopy(), nil
}

func (b *_BACnetConstructedDataNodeTypeBuilder) MustBuild() BACnetConstructedDataNodeType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataNodeTypeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNodeTypeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNodeTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNodeTypeBuilder().(*_BACnetConstructedDataNodeTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNodeTypeBuilder creates a BACnetConstructedDataNodeTypeBuilder
func (b *_BACnetConstructedDataNodeType) CreateBACnetConstructedDataNodeTypeBuilder() BACnetConstructedDataNodeTypeBuilder {
	if b == nil {
		return NewBACnetConstructedDataNodeTypeBuilder()
	}
	return &_BACnetConstructedDataNodeTypeBuilder{_BACnetConstructedDataNodeType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNodeType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NODE_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNodeType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetNodeType() BACnetNodeTypeTagged {
	return m.NodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNodeType) GetActualValue() BACnetNodeTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetNodeTypeTagged(m.GetNodeType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNodeType(structType any) BACnetConstructedDataNodeType {
	if casted, ok := structType.(BACnetConstructedDataNodeType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNodeType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNodeType) GetTypeName() string {
	return "BACnetConstructedDataNodeType"
}

func (m *_BACnetConstructedDataNodeType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (nodeType)
	lengthInBits += m.NodeType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNodeType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNodeType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNodeType BACnetConstructedDataNodeType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNodeType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeType, err := ReadSimpleField[BACnetNodeTypeTagged](ctx, "nodeType", ReadComplex[BACnetNodeTypeTagged](BACnetNodeTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeType' field"))
	}
	m.NodeType = nodeType

	actualValue, err := ReadVirtualField[BACnetNodeTypeTagged](ctx, "actualValue", (*BACnetNodeTypeTagged)(nil), nodeType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNodeType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNodeType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNodeType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNodeType")
		}

		if err := WriteSimpleField[BACnetNodeTypeTagged](ctx, "nodeType", m.GetNodeType(), WriteComplex[BACnetNodeTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNodeType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNodeType) IsBACnetConstructedDataNodeType() {}

func (m *_BACnetConstructedDataNodeType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNodeType) deepCopy() *_BACnetConstructedDataNodeType {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNodeTypeCopy := &_BACnetConstructedDataNodeType{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetNodeTypeTagged](m.NodeType),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNodeTypeCopy
}

func (m *_BACnetConstructedDataNodeType) String() string {
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
