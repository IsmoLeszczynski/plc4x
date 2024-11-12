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

// BACnetConfirmedServiceRequestAtomicReadFileRecord is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileRecord
type BACnetConfirmedServiceRequestAtomicReadFileRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartRecord returns FileStartRecord (property field)
	GetFileStartRecord() BACnetApplicationTagSignedInteger
	// GetRequestRecordCount returns RequestRecordCount (property field)
	GetRequestRecordCount() BACnetApplicationTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestAtomicReadFileRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAtomicReadFileRecord()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	CreateBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
}

// _BACnetConfirmedServiceRequestAtomicReadFileRecord is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileRecord struct {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
	FileStartRecord    BACnetApplicationTagSignedInteger
	RequestRecordCount BACnetApplicationTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestAtomicReadFileRecord = (*_BACnetConfirmedServiceRequestAtomicReadFileRecord)(nil)
var _ BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordRequirements = (*_BACnetConfirmedServiceRequestAtomicReadFileRecord)(nil)

// NewBACnetConfirmedServiceRequestAtomicReadFileRecord factory function for _BACnetConfirmedServiceRequestAtomicReadFileRecord
func NewBACnetConfirmedServiceRequestAtomicReadFileRecord(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag, fileStartRecord BACnetApplicationTagSignedInteger, requestRecordCount BACnetApplicationTagUnsignedInteger) *_BACnetConfirmedServiceRequestAtomicReadFileRecord {
	if fileStartRecord == nil {
		panic("fileStartRecord of type BACnetApplicationTagSignedInteger for BACnetConfirmedServiceRequestAtomicReadFileRecord must not be nil")
	}
	if requestRecordCount == nil {
		panic("requestRecordCount of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestAtomicReadFileRecord must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileRecord{
		BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartRecord:    fileStartRecord,
		RequestRecordCount: requestRecordCount,
	}
	_result.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder is a builder for BACnetConfirmedServiceRequestAtomicReadFileRecord
type BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fileStartRecord BACnetApplicationTagSignedInteger, requestRecordCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	// WithFileStartRecord adds FileStartRecord (property field)
	WithFileStartRecord(BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	// WithFileStartRecordBuilder adds FileStartRecord (property field) which is build by the builder
	WithFileStartRecordBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	// WithRequestRecordCount adds RequestRecordCount (property field)
	WithRequestRecordCount(BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	// WithRequestRecordCountBuilder adds RequestRecordCount (property field) which is build by the builder
	WithRequestRecordCountBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder
	// Build builds the BACnetConfirmedServiceRequestAtomicReadFileRecord or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAtomicReadFileRecord, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAtomicReadFileRecord
}

// NewBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder() creates a BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
func NewBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	return &_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder{_BACnetConfirmedServiceRequestAtomicReadFileRecord: new(_BACnetConfirmedServiceRequestAtomicReadFileRecord)}
}

type _BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder struct {
	*_BACnetConfirmedServiceRequestAtomicReadFileRecord

	parentBuilder *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) = (*_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) setParent(contract BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract) {
	b.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = contract
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) WithMandatoryFields(fileStartRecord BACnetApplicationTagSignedInteger, requestRecordCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	return b.WithFileStartRecord(fileStartRecord).WithRequestRecordCount(requestRecordCount)
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) WithFileStartRecord(fileStartRecord BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	b.FileStartRecord = fileStartRecord
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) WithFileStartRecordBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	builder := builderSupplier(b.FileStartRecord.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.FileStartRecord, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) WithRequestRecordCount(requestRecordCount BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	b.RequestRecordCount = requestRecordCount
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) WithRequestRecordCountBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	builder := builderSupplier(b.RequestRecordCount.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RequestRecordCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) Build() (BACnetConfirmedServiceRequestAtomicReadFileRecord, error) {
	if b.FileStartRecord == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fileStartRecord' not set"))
	}
	if b.RequestRecordCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestRecordCount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestAtomicReadFileRecord.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) MustBuild() BACnetConfirmedServiceRequestAtomicReadFileRecord {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) Done() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder().(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) buildForBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord() (BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder().(*_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder creates a BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder
func (b *_BACnetConfirmedServiceRequestAtomicReadFileRecord) CreateBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder() BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestAtomicReadFileRecordBuilder()
	}
	return &_BACnetConfirmedServiceRequestAtomicReadFileRecordBuilder{_BACnetConfirmedServiceRequestAtomicReadFileRecord: b.deepCopy()}
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

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract {
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetFileStartRecord() BACnetApplicationTagSignedInteger {
	return m.FileStartRecord
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetRequestRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileRecord(structType any) BACnetConfirmedServiceRequestAtomicReadFileRecord {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileRecord"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartRecord)
	lengthInBits += m.FileStartRecord.GetLengthInBits(ctx)

	// Simple field (requestRecordCount)
	lengthInBits += m.RequestRecordCount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord) (__bACnetConfirmedServiceRequestAtomicReadFileRecord BACnetConfirmedServiceRequestAtomicReadFileRecord, err error) {
	m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartRecord, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartRecord", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartRecord' field"))
	}
	m.FileStartRecord = fileStartRecord

	requestRecordCount, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestRecordCount", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestRecordCount' field"))
	}
	m.RequestRecordCount = requestRecordCount

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartRecord", m.GetFileStartRecord(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartRecord' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestRecordCount", m.GetRequestRecordCount(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestRecordCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) IsBACnetConfirmedServiceRequestAtomicReadFileRecord() {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) deepCopy() *_BACnetConfirmedServiceRequestAtomicReadFileRecord {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAtomicReadFileRecordCopy := &_BACnetConfirmedServiceRequestAtomicReadFileRecord{
		m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.FileStartRecord),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RequestRecordCount),
	}
	m.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordContract.(*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)._SubType = m
	return _BACnetConfirmedServiceRequestAtomicReadFileRecordCopy
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) String() string {
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
