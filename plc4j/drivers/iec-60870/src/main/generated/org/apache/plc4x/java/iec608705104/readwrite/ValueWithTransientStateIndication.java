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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class ValueWithTransientStateIndication implements Message {

  // Properties.
  protected final boolean transientState;
  protected final byte value;

  public ValueWithTransientStateIndication(boolean transientState, byte value) {
    super();
    this.transientState = transientState;
    this.value = value;
  }

  public boolean getTransientState() {
    return transientState;
  }

  public byte getValue() {
    return value;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ValueWithTransientStateIndication");

    // Simple Field (transientState)
    writeSimpleField(
        "transientState",
        transientState,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (value)
    writeSimpleField(
        "value",
        value,
        writeUnsignedByte(writeBuffer, 7),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("ValueWithTransientStateIndication");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ValueWithTransientStateIndication _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (transientState)
    lengthInBits += 1;

    // Simple field (value)
    lengthInBits += 7;

    return lengthInBits;
  }

  public static ValueWithTransientStateIndication staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("ValueWithTransientStateIndication");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean transientState =
        readSimpleField(
            "transientState",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte value =
        readSimpleField(
            "value",
            readUnsignedByte(readBuffer, 7),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("ValueWithTransientStateIndication");
    // Create the instance
    ValueWithTransientStateIndication _valueWithTransientStateIndication;
    _valueWithTransientStateIndication =
        new ValueWithTransientStateIndication(transientState, value);
    return _valueWithTransientStateIndication;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ValueWithTransientStateIndication)) {
      return false;
    }
    ValueWithTransientStateIndication that = (ValueWithTransientStateIndication) o;
    return (getTransientState() == that.getTransientState())
        && (getValue() == that.getValue())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTransientState(), getValue());
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
