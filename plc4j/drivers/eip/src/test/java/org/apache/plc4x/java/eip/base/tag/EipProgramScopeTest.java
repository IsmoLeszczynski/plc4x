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
package org.apache.plc4x.java.eip.base.tag;

import org.apache.plc4x.java.eip.base.tag.EipTag.MemberElement;
import org.apache.plc4x.java.eip.base.tag.EipTag.PathElement;
import org.apache.plc4x.java.eip.base.tag.EipTag.SymbolElement;
import org.apache.plc4x.java.eip.readwrite.CIPDataTypeCode;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

/**
 * Program-scoped Logix tags. A controller addresses one with the symbol {@code Program:<name>},
 * colon included, followed by the tag's own members - so the scope is one path element and must
 * not be split on its colon, nor read as a data type.
 */
class EipProgramScopeTest {

    @Test
    void programScopedTagIsTheScopeThenTheTag() {
        EipTag tag = EipTag.of("Program:Main.Counter");
        assertNotNull(tag);
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Counter")), tag.getPathElements());
        assertEquals(CIPDataTypeCode.DINT, tag.getType());
        assertEquals("Program:Main.Counter", tag.getTag());
        assertEquals(1, tag.getElementNb());
    }

    @Test
    void membersIndicesAndTypeFollowTheScope() {
        EipTag tag = EipTag.of("Program:Main.Recipe[3].SetPoint:REAL");
        assertNotNull(tag);
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Recipe"),
            new MemberElement((short) 3), new SymbolElement("SetPoint")), tag.getPathElements());
        assertEquals(CIPDataTypeCode.REAL, tag.getType());
        assertTrue(tag.getArrayInfo().isEmpty());
    }

    @Test
    void trailingRangeFollowsTheScope() {
        EipTag tag = EipTag.of("Program:Main.Arr[0..3]:DINT");
        assertNotNull(tag);
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Arr")), tag.getPathElements());
        assertEquals(4, tag.getElementNb());
        assertEquals(1, tag.getArrayInfo().size());

        EipTag offset = EipTag.of("Program:Main.Arr[2..3]:DINT");
        assertNotNull(offset);
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Arr"),
            new MemberElement((short) 2)), offset.getPathElements());
        assertEquals(2, offset.getElementNb());
    }

    /** The grammar allows a leading '%' before any address; it is not part of the scope symbol. */
    @Test
    void percentPrefixBeforeTheScope() {
        EipTag tag = EipTag.of("%Program:Main.Counter:INT");
        assertNotNull(tag);
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Counter")), tag.getPathElements());
        assertEquals(CIPDataTypeCode.INT, tag.getType());
    }

    @Test
    void addressStringRoundTrips() {
        for (String address : new String[]{"Program:Main.Counter:DINT", "Program:Main.Recipe[3].SetPoint:REAL",
            "Program:Main.Arr[0..3]:DINT"}) {
            EipTag tag = EipTag.of(address);
            assertNotNull(tag, address);
            assertEquals(address, tag.getAddressString());
            assertEquals(tag.getPathElements(), EipTag.of(tag.getAddressString()).getPathElements());
        }
    }

    /** The older constructors decompose the same way as of(). */
    @Test
    void constructorDecomposesTheScopeToo() {
        assertEquals(List.of(new SymbolElement("Program:Main"), new SymbolElement("Counter")),
            new EipTag("Program:Main.Counter", CIPDataTypeCode.DINT).getPathElements());
    }

    /** Only a whole "Program:<name>." prefix is a scope; nothing else may carry a colon. */
    @Test
    void malformedScopesAreRejected() {
        assertNull(EipTag.of("Program:Main"));
        assertNull(EipTag.of("Program:.Counter"));
        assertNull(EipTag.of("Program:Main:Counter"));
        assertNull(EipTag.of("Program:Main[1].Counter"));
        assertNull(EipTag.of("Main.Program:Sub.Counter"));
        assertNull(EipTag.of("Program:Main.Program:Sub.Counter"));
        assertNull(EipTag.of("Controller:Main.Counter"));
    }

    /** The dotted form is unchanged: three plain symbols, no scope. */
    @Test
    void dottedProgramIsStillThreeSymbols() {
        List<PathElement> expected =
            List.of(new SymbolElement("Program"), new SymbolElement("MainProgram"), new SymbolElement("MyTag"));
        assertEquals(expected, EipTag.of("Program.MainProgram.MyTag:DINT").getPathElements());
    }
}
