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

// BACnetConstructedDataInProcess is the corresponding interface of BACnetConstructedDataInProcess
type BACnetConstructedDataInProcess interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInProcess returns InProcess (property field)
	GetInProcess() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataInProcess is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInProcess()
	// CreateBuilder creates a BACnetConstructedDataInProcessBuilder
	CreateBACnetConstructedDataInProcessBuilder() BACnetConstructedDataInProcessBuilder
}

// _BACnetConstructedDataInProcess is the data-structure of this message
type _BACnetConstructedDataInProcess struct {
	BACnetConstructedDataContract
	InProcess BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataInProcess = (*_BACnetConstructedDataInProcess)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInProcess)(nil)

// NewBACnetConstructedDataInProcess factory function for _BACnetConstructedDataInProcess
func NewBACnetConstructedDataInProcess(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, inProcess BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInProcess {
	if inProcess == nil {
		panic("inProcess of type BACnetApplicationTagBoolean for BACnetConstructedDataInProcess must not be nil")
	}
	_result := &_BACnetConstructedDataInProcess{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InProcess:                     inProcess,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataInProcessBuilder is a builder for BACnetConstructedDataInProcess
type BACnetConstructedDataInProcessBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(inProcess BACnetApplicationTagBoolean) BACnetConstructedDataInProcessBuilder
	// WithInProcess adds InProcess (property field)
	WithInProcess(BACnetApplicationTagBoolean) BACnetConstructedDataInProcessBuilder
	// WithInProcessBuilder adds InProcess (property field) which is build by the builder
	WithInProcessBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataInProcessBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataInProcess or returns an error if something is wrong
	Build() (BACnetConstructedDataInProcess, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataInProcess
}

// NewBACnetConstructedDataInProcessBuilder() creates a BACnetConstructedDataInProcessBuilder
func NewBACnetConstructedDataInProcessBuilder() BACnetConstructedDataInProcessBuilder {
	return &_BACnetConstructedDataInProcessBuilder{_BACnetConstructedDataInProcess: new(_BACnetConstructedDataInProcess)}
}

type _BACnetConstructedDataInProcessBuilder struct {
	*_BACnetConstructedDataInProcess

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataInProcessBuilder) = (*_BACnetConstructedDataInProcessBuilder)(nil)

func (b *_BACnetConstructedDataInProcessBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataInProcessBuilder) WithMandatoryFields(inProcess BACnetApplicationTagBoolean) BACnetConstructedDataInProcessBuilder {
	return b.WithInProcess(inProcess)
}

func (b *_BACnetConstructedDataInProcessBuilder) WithInProcess(inProcess BACnetApplicationTagBoolean) BACnetConstructedDataInProcessBuilder {
	b.InProcess = inProcess
	return b
}

func (b *_BACnetConstructedDataInProcessBuilder) WithInProcessBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataInProcessBuilder {
	builder := builderSupplier(b.InProcess.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.InProcess, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataInProcessBuilder) Build() (BACnetConstructedDataInProcess, error) {
	if b.InProcess == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'inProcess' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataInProcess.deepCopy(), nil
}

func (b *_BACnetConstructedDataInProcessBuilder) MustBuild() BACnetConstructedDataInProcess {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataInProcessBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataInProcessBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataInProcessBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataInProcessBuilder().(*_BACnetConstructedDataInProcessBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataInProcessBuilder creates a BACnetConstructedDataInProcessBuilder
func (b *_BACnetConstructedDataInProcess) CreateBACnetConstructedDataInProcessBuilder() BACnetConstructedDataInProcessBuilder {
	if b == nil {
		return NewBACnetConstructedDataInProcessBuilder()
	}
	return &_BACnetConstructedDataInProcessBuilder{_BACnetConstructedDataInProcess: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInProcess) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IN_PROCESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInProcess) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetInProcess() BACnetApplicationTagBoolean {
	return m.InProcess
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetInProcess())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInProcess(structType any) BACnetConstructedDataInProcess {
	if casted, ok := structType.(BACnetConstructedDataInProcess); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInProcess); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInProcess) GetTypeName() string {
	return "BACnetConstructedDataInProcess"
}

func (m *_BACnetConstructedDataInProcess) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (inProcess)
	lengthInBits += m.InProcess.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInProcess) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInProcess) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInProcess BACnetConstructedDataInProcess, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInProcess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	inProcess, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "inProcess", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inProcess' field"))
	}
	m.InProcess = inProcess

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), inProcess)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInProcess")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInProcess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInProcess) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInProcess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInProcess")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "inProcess", m.GetInProcess(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'inProcess' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInProcess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInProcess")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInProcess) IsBACnetConstructedDataInProcess() {}

func (m *_BACnetConstructedDataInProcess) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataInProcess) deepCopy() *_BACnetConstructedDataInProcess {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataInProcessCopy := &_BACnetConstructedDataInProcess{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.InProcess),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataInProcessCopy
}

func (m *_BACnetConstructedDataInProcess) String() string {
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
