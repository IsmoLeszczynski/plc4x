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

public class NLMSecurityResponse extends NLM implements Message {

  // Accessors for discriminator values.
  public Short getMessageType() {
    return (short) 0x0C;
  }

  // Properties.
  protected final SecurityResponseCode responseCode;
  protected final long originalMessageId;
  protected final long originalTimestamp;
  protected final byte[] variableParameters;

  // Arguments.
  protected final Integer apduLength;

  public NLMSecurityResponse(
      SecurityResponseCode responseCode,
      long originalMessageId,
      long originalTimestamp,
      byte[] variableParameters,
      Integer apduLength) {
    super(apduLength);
    this.responseCode = responseCode;
    this.originalMessageId = originalMessageId;
    this.originalTimestamp = originalTimestamp;
    this.variableParameters = variableParameters;
    this.apduLength = apduLength;
  }

  public SecurityResponseCode getResponseCode() {
    return responseCode;
  }

  public long getOriginalMessageId() {
    return originalMessageId;
  }

  public long getOriginalTimestamp() {
    return originalTimestamp;
  }

  public byte[] getVariableParameters() {
    return variableParameters;
  }

  @Override
  protected void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NLMSecurityResponse");

    // Simple Field (responseCode)
    writeSimpleEnumField(
        "responseCode",
        "SecurityResponseCode",
        responseCode,
        writeEnum(
            SecurityResponseCode::getValue,
            SecurityResponseCode::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (originalMessageId)
    writeSimpleField("originalMessageId", originalMessageId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (originalTimestamp)
    writeSimpleField("originalTimestamp", originalTimestamp, writeUnsignedLong(writeBuffer, 32));

    // Array Field (variableParameters)
    writeByteArrayField("variableParameters", variableParameters, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("NLMSecurityResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NLMSecurityResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (responseCode)
    lengthInBits += 8;

    // Simple field (originalMessageId)
    lengthInBits += 32;

    // Simple field (originalTimestamp)
    lengthInBits += 32;

    // Array field
    if (variableParameters != null) {
      lengthInBits += 8 * variableParameters.length;
    }

    return lengthInBits;
  }

  public static NLMBuilder staticParseNLMBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("NLMSecurityResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SecurityResponseCode responseCode =
        readEnumField(
            "responseCode",
            "SecurityResponseCode",
            readEnum(SecurityResponseCode::enumForValue, readUnsignedShort(readBuffer, 8)));

    long originalMessageId = readSimpleField("originalMessageId", readUnsignedLong(readBuffer, 32));

    long originalTimestamp = readSimpleField("originalTimestamp", readUnsignedLong(readBuffer, 32));

    byte[] variableParameters =
        readBuffer.readByteArray(
            "variableParameters", Math.toIntExact((apduLength) - (((((1) + (1)) + (4)) + (4)))));

    readBuffer.closeContext("NLMSecurityResponse");
    // Create the instance
    return new NLMSecurityResponseBuilderImpl(
        responseCode, originalMessageId, originalTimestamp, variableParameters, apduLength);
  }

  public static class NLMSecurityResponseBuilderImpl implements NLM.NLMBuilder {
    private final SecurityResponseCode responseCode;
    private final long originalMessageId;
    private final long originalTimestamp;
    private final byte[] variableParameters;
    private final Integer apduLength;

    public NLMSecurityResponseBuilderImpl(
        SecurityResponseCode responseCode,
        long originalMessageId,
        long originalTimestamp,
        byte[] variableParameters,
        Integer apduLength) {
      this.responseCode = responseCode;
      this.originalMessageId = originalMessageId;
      this.originalTimestamp = originalTimestamp;
      this.variableParameters = variableParameters;
      this.apduLength = apduLength;
    }

    public NLMSecurityResponse build(Integer apduLength) {

      NLMSecurityResponse nLMSecurityResponse =
          new NLMSecurityResponse(
              responseCode, originalMessageId, originalTimestamp, variableParameters, apduLength);
      return nLMSecurityResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMSecurityResponse)) {
      return false;
    }
    NLMSecurityResponse that = (NLMSecurityResponse) o;
    return (getResponseCode() == that.getResponseCode())
        && (getOriginalMessageId() == that.getOriginalMessageId())
        && (getOriginalTimestamp() == that.getOriginalTimestamp())
        && (getVariableParameters() == that.getVariableParameters())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getResponseCode(),
        getOriginalMessageId(),
        getOriginalTimestamp(),
        getVariableParameters());
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
