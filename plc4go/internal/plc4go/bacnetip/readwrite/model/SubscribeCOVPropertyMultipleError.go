/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SubscribeCOVPropertyMultipleError is the data-structure of this message
type SubscribeCOVPropertyMultipleError struct {
	*BACnetError
	ErrorType               *ErrorEnclosed
	FirstFailedSubscription *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

// ISubscribeCOVPropertyMultipleError is the corresponding interface of SubscribeCOVPropertyMultipleError
type ISubscribeCOVPropertyMultipleError interface {
	IBACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() *ErrorEnclosed
	// GetFirstFailedSubscription returns FirstFailedSubscription (property field)
	GetFirstFailedSubscription() *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *SubscribeCOVPropertyMultipleError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SubscribeCOVPropertyMultipleError) InitializeParent(parent *BACnetError) {}

func (m *SubscribeCOVPropertyMultipleError) GetParent() *BACnetError {
	return m.BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SubscribeCOVPropertyMultipleError) GetErrorType() *ErrorEnclosed {
	return m.ErrorType
}

func (m *SubscribeCOVPropertyMultipleError) GetFirstFailedSubscription() *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return m.FirstFailedSubscription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleError factory function for SubscribeCOVPropertyMultipleError
func NewSubscribeCOVPropertyMultipleError(errorType *ErrorEnclosed, firstFailedSubscription *SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) *SubscribeCOVPropertyMultipleError {
	_result := &SubscribeCOVPropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
		BACnetError:             NewBACnetError(),
	}
	_result.Child = _result
	return _result
}

func CastSubscribeCOVPropertyMultipleError(structType interface{}) *SubscribeCOVPropertyMultipleError {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleError); ok {
		return &casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleError); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastSubscribeCOVPropertyMultipleError(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastSubscribeCOVPropertyMultipleError(casted.Child)
	}
	return nil
}

func (m *SubscribeCOVPropertyMultipleError) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleError"
}

func (m *SubscribeCOVPropertyMultipleError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SubscribeCOVPropertyMultipleError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (firstFailedSubscription)
	lengthInBits += m.FirstFailedSubscription.GetLengthInBits()

	return lengthInBits
}

func (m *SubscribeCOVPropertyMultipleError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SubscribeCOVPropertyMultipleErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (*SubscribeCOVPropertyMultipleError, error) {
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleError"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, pullErr
	}
	_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer, uint8(uint8(0)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field")
	}
	errorType := CastErrorEnclosed(_errorType)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (firstFailedSubscription)
	if pullErr := readBuffer.PullContext("firstFailedSubscription"); pullErr != nil {
		return nil, pullErr
	}
	_firstFailedSubscription, _firstFailedSubscriptionErr := SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(readBuffer, uint8(uint8(1)))
	if _firstFailedSubscriptionErr != nil {
		return nil, errors.Wrap(_firstFailedSubscriptionErr, "Error parsing 'firstFailedSubscription' field")
	}
	firstFailedSubscription := CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(_firstFailedSubscription)
	if closeErr := readBuffer.CloseContext("firstFailedSubscription"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleError"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SubscribeCOVPropertyMultipleError{
		ErrorType:               CastErrorEnclosed(errorType),
		FirstFailedSubscription: CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(firstFailedSubscription),
		BACnetError:             &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *SubscribeCOVPropertyMultipleError) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleError"); pushErr != nil {
			return pushErr
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return pushErr
		}
		_errorTypeErr := m.ErrorType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return popErr
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (firstFailedSubscription)
		if pushErr := writeBuffer.PushContext("firstFailedSubscription"); pushErr != nil {
			return pushErr
		}
		_firstFailedSubscriptionErr := m.FirstFailedSubscription.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("firstFailedSubscription"); popErr != nil {
			return popErr
		}
		if _firstFailedSubscriptionErr != nil {
			return errors.Wrap(_firstFailedSubscriptionErr, "Error serializing 'firstFailedSubscription' field")
		}

		if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleError"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SubscribeCOVPropertyMultipleError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
