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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetPriorityValueConstructedValue extends BACnetPriorityValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetConstructedData constructedValue;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetPriorityValueConstructedValue(
      BACnetTagHeader peekedTagHeader,
      BACnetConstructedData constructedValue,
      BACnetObjectType objectTypeArgument) {
    super(peekedTagHeader, objectTypeArgument);
    this.constructedValue = constructedValue;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetConstructedData getConstructedValue() {
    return constructedValue;
  }

  @Override
  protected void serializeBACnetPriorityValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPriorityValueConstructedValue");

    // Simple Field (constructedValue)
    writeSimpleField("constructedValue", constructedValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetPriorityValueConstructedValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPriorityValueConstructedValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (constructedValue)
    lengthInBits += constructedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPriorityValueBuilder staticParseBACnetPriorityValueBuilder(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetPriorityValueConstructedValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetConstructedData constructedValue =
        readSimpleField(
            "constructedValue",
            readComplex(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (0),
                        (BACnetObjectType) (objectTypeArgument),
                        (BACnetPropertyIdentifier)
                            (BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE),
                        (BACnetTagPayloadUnsignedInteger) (null)),
                readBuffer));

    readBuffer.closeContext("BACnetPriorityValueConstructedValue");
    // Create the instance
    return new BACnetPriorityValueConstructedValueBuilderImpl(constructedValue, objectTypeArgument);
  }

  public static class BACnetPriorityValueConstructedValueBuilderImpl
      implements BACnetPriorityValue.BACnetPriorityValueBuilder {
    private final BACnetConstructedData constructedValue;
    private final BACnetObjectType objectTypeArgument;

    public BACnetPriorityValueConstructedValueBuilderImpl(
        BACnetConstructedData constructedValue, BACnetObjectType objectTypeArgument) {
      this.constructedValue = constructedValue;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetPriorityValueConstructedValue build(
        BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
      BACnetPriorityValueConstructedValue bACnetPriorityValueConstructedValue =
          new BACnetPriorityValueConstructedValue(
              peekedTagHeader, constructedValue, objectTypeArgument);
      return bACnetPriorityValueConstructedValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPriorityValueConstructedValue)) {
      return false;
    }
    BACnetPriorityValueConstructedValue that = (BACnetPriorityValueConstructedValue) o;
    return (getConstructedValue() == that.getConstructedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getConstructedValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
