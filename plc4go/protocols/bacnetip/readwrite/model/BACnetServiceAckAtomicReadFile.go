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

// BACnetServiceAckAtomicReadFile is the corresponding interface of BACnetServiceAckAtomicReadFile
type BACnetServiceAckAtomicReadFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetEndOfFile returns EndOfFile (property field)
	GetEndOfFile() BACnetApplicationTagBoolean
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() BACnetServiceAckAtomicReadFileStreamOrRecord
}

// BACnetServiceAckAtomicReadFileExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicReadFile.
// This is useful for switch cases.
type BACnetServiceAckAtomicReadFileExactly interface {
	BACnetServiceAckAtomicReadFile
	isBACnetServiceAckAtomicReadFile() bool
}

// _BACnetServiceAckAtomicReadFile is the data-structure of this message
type _BACnetServiceAckAtomicReadFile struct {
	BACnetServiceAckContract
	EndOfFile    BACnetApplicationTagBoolean
	AccessMethod BACnetServiceAckAtomicReadFileStreamOrRecord
}

var _ BACnetServiceAckAtomicReadFile = (*_BACnetServiceAckAtomicReadFile)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckAtomicReadFile)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicReadFile) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFile) GetEndOfFile() BACnetApplicationTagBoolean {
	return m.EndOfFile
}

func (m *_BACnetServiceAckAtomicReadFile) GetAccessMethod() BACnetServiceAckAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFile factory function for _BACnetServiceAckAtomicReadFile
func NewBACnetServiceAckAtomicReadFile(endOfFile BACnetApplicationTagBoolean, accessMethod BACnetServiceAckAtomicReadFileStreamOrRecord, serviceAckLength uint32) *_BACnetServiceAckAtomicReadFile {
	_result := &_BACnetServiceAckAtomicReadFile{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		EndOfFile:                endOfFile,
		AccessMethod:             accessMethod,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFile(structType any) BACnetServiceAckAtomicReadFile {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFile) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFile"
}

func (m *_BACnetServiceAckAtomicReadFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (endOfFile)
	lengthInBits += m.EndOfFile.GetLengthInBits(ctx)

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckAtomicReadFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckAtomicReadFile BACnetServiceAckAtomicReadFile, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	endOfFile, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "endOfFile", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endOfFile' field"))
	}
	m.EndOfFile = endOfFile

	accessMethod, err := ReadSimpleField[BACnetServiceAckAtomicReadFileStreamOrRecord](ctx, "accessMethod", ReadComplex[BACnetServiceAckAtomicReadFileStreamOrRecord](BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessMethod' field"))
	}
	m.AccessMethod = accessMethod

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFile")
	}

	return m, nil
}

func (m *_BACnetServiceAckAtomicReadFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicReadFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFile")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "endOfFile", m.GetEndOfFile(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endOfFile' field")
		}

		if err := WriteSimpleField[BACnetServiceAckAtomicReadFileStreamOrRecord](ctx, "accessMethod", m.GetAccessMethod(), WriteComplex[BACnetServiceAckAtomicReadFileStreamOrRecord](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFile")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicReadFile) isBACnetServiceAckAtomicReadFile() bool {
	return true
}

func (m *_BACnetServiceAckAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
