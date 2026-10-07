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
package org.apache.plc4x.java.eip.base;

import org.apache.plc4x.java.eip.base.tag.EipTag;
import org.apache.plc4x.java.eip.readwrite.CipReadRequest;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.bytebased.WithByteBasedOption;
import org.apache.plc4x.java.spi.buffers.bytebased.WriteBufferByteBased;
import org.junit.jupiter.api.Test;

import java.io.ByteArrayOutputStream;
import java.nio.charset.StandardCharsets;

import static org.junit.jupiter.api.Assertions.*;

/**
 * The request path of a program-scoped tag: an ANSI extended symbol segment
 * {@code Program:<name>}, colon included, then one segment per member - which is how a Logix
 * controller resolves a program-scoped tag.
 */
class EipProgramScopeEncodingTest {

    @Test
    void readRequestCarriesTheScopeAsOneSymbolSegment() throws Exception {
        EipTag tag = EipTag.of("Program:Main.Counter");
        CipReadRequest request = new CipReadRequest(
            EipTcpConnection.toAnsi(tag, EipProgramScopeEncodingTest::buffer), tag.getElementNb());

        WriteBufferByteBased out = buffer(request.getLengthInBytes());
        request.serialize(out);

        ByteArrayOutputStream expected = new ByteArrayOutputStream();
        expected.write(0x4C);                                // Read Tag service
        expected.write(12);                                  // request path size, in words
        expected.write(0x91);                                // ANSI extended symbol segment
        expected.write(12);
        expected.writeBytes("Program:Main".getBytes(StandardCharsets.US_ASCII));  // even: no pad
        expected.write(0x91);
        expected.write(7);
        expected.writeBytes("Counter".getBytes(StandardCharsets.US_ASCII));
        expected.write(0x00);                                // pad to an even length
        expected.write(0x01);                                // element count, little-endian
        expected.write(0x00);

        assertArrayEquals(expected.toByteArray(), out.getBytes());
    }

    @Test
    void oddLengthScopeIsPadded() throws Exception {
        byte[] path = EipTcpConnection.toAnsi(EipTag.of("Program:Mix.A"), EipProgramScopeEncodingTest::buffer);

        ByteArrayOutputStream expected = new ByteArrayOutputStream();
        expected.write(0x91);
        expected.write(11);
        expected.writeBytes("Program:Mix".getBytes(StandardCharsets.US_ASCII));
        expected.write(0x00);
        expected.write(0x91);
        expected.write(1);
        expected.write('A');
        expected.write(0x00);

        assertArrayEquals(expected.toByteArray(), path);
    }

    private static WriteBufferByteBased buffer(int size) {
        return new WriteBufferByteBased(new byte[size],
            WithByteBasedOption.WithByteOrder("LITTLE_ENDIAN"),
            WithOption.WithUnsignedIntegerEncoding("unsigned-binary"),
            WithOption.WithSignedIntegerEncoding("twos-complement"),
            WithOption.WithStringEncoding("UTF8"));
    }
}
