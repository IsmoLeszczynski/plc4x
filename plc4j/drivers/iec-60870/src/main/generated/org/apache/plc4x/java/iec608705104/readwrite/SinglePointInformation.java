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

public class SinglePointInformation implements Message {

  // Properties.
  protected final boolean invalid;
  protected final boolean notTopical;
  protected final boolean substituted;
  protected final boolean blocked;
  protected final boolean stausOn;

  public SinglePointInformation(
      boolean invalid, boolean notTopical, boolean substituted, boolean blocked, boolean stausOn) {
    super();
    this.invalid = invalid;
    this.notTopical = notTopical;
    this.substituted = substituted;
    this.blocked = blocked;
    this.stausOn = stausOn;
  }

  public boolean getInvalid() {
    return invalid;
  }

  public boolean getNotTopical() {
    return notTopical;
  }

  public boolean getSubstituted() {
    return substituted;
  }

  public boolean getBlocked() {
    return blocked;
  }

  public boolean getStausOn() {
    return stausOn;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SinglePointInformation");

    // Simple Field (invalid)
    writeSimpleField(
        "invalid",
        invalid,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (notTopical)
    writeSimpleField(
        "notTopical",
        notTopical,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (substituted)
    writeSimpleField(
        "substituted",
        substituted,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (blocked)
    writeSimpleField(
        "blocked",
        blocked,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        (byte) 0,
        writeUnsignedByte(writeBuffer, 3),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (stausOn)
    writeSimpleField(
        "stausOn",
        stausOn,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("SinglePointInformation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SinglePointInformation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (invalid)
    lengthInBits += 1;

    // Simple field (notTopical)
    lengthInBits += 1;

    // Simple field (substituted)
    lengthInBits += 1;

    // Simple field (blocked)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 3;

    // Simple field (stausOn)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static SinglePointInformation staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SinglePointInformation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean invalid =
        readSimpleField(
            "invalid", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean notTopical =
        readSimpleField(
            "notTopical",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean substituted =
        readSimpleField(
            "substituted",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean blocked =
        readSimpleField(
            "blocked", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    Byte reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedByte(readBuffer, 3),
            (byte) 0,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean stausOn =
        readSimpleField(
            "stausOn", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("SinglePointInformation");
    // Create the instance
    SinglePointInformation _singlePointInformation;
    _singlePointInformation =
        new SinglePointInformation(invalid, notTopical, substituted, blocked, stausOn);
    return _singlePointInformation;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SinglePointInformation)) {
      return false;
    }
    SinglePointInformation that = (SinglePointInformation) o;
    return (getInvalid() == that.getInvalid())
        && (getNotTopical() == that.getNotTopical())
        && (getSubstituted() == that.getSubstituted())
        && (getBlocked() == that.getBlocked())
        && (getStausOn() == that.getStausOn())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getInvalid(), getNotTopical(), getSubstituted(), getBlocked(), getStausOn());
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
