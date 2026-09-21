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
package org.apache.plc4x.java.eip.base;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.ServerSocket;
import java.net.Socket;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;
import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.concurrent.TimeUnit;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

/**
 * The capability probe against devices that are not a Logix controller. The scripted driver
 * testsuites cover the answers a controller gives; these cover a device that has no Connection
 * Manager at all and one that ends the session when asked about it, which is what cpppo does.
 */
class EipHandshakeProbeTest {

    private static final int LIST_SERVICES = 0x0004;
    private static final int REGISTER_SESSION = 0x0065;
    private static final int UNREGISTER_SESSION = 0x0066;
    private static final int SEND_RR_DATA = 0x006F;
    private static final int SESSION_HANDLE = 0x5803DA1F;

    /** How the fake device answers a SendRRData. */
    enum RrDataBehaviour {
        /** A Get_Attribute_All answer listing no classes; anything else is a CIP error. */
        EMPTY_CLASS_LIST,
        /** Every CIP service is answered with status 0x08, Service not supported. */
        CIP_SERVICE_NOT_SUPPORTED,
        /** An encapsulation-level error with an empty body, and the session is ended. */
        ENCAPSULATION_ERROR_AND_CLOSE
    }

    private ServerSocket serverSocket;
    private Thread deviceThread;
    private final List<Integer> receivedCommands = new CopyOnWriteArrayList<>();

    @BeforeEach
    void setUp() throws IOException {
        serverSocket = new ServerSocket();
        serverSocket.bind(new InetSocketAddress("127.0.0.1", 0));
    }

    @AfterEach
    void tearDown() throws Exception {
        serverSocket.close();
        if (deviceThread != null) {
            deviceThread.join(5000);
        }
    }

    @Test
    void aDeviceWithoutAnyClassesConnectsAfterOneProbe() throws Exception {
        startDevice(RrDataBehaviour.EMPTY_CLASS_LIST);

        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection(url())) {
            assertTrue(connection.isConnected());
        }

        // The class list answers everything, so no class is asked about on its own.
        assertEquals(List.of(LIST_SERVICES, REGISTER_SESSION, SEND_RR_DATA, UNREGISTER_SESSION), receivedCommands);
    }

    @Test
    void aDeviceWithoutGetAttributeAllIsAskedAboutEachClass() throws Exception {
        startDevice(RrDataBehaviour.CIP_SERVICE_NOT_SUPPORTED);

        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection(url())) {
            assertTrue(connection.isConnected());
        }

        // Get_Attribute_All, then Get_Attribute_Single for the Message Router and the Connection Manager.
        assertEquals(List.of(LIST_SERVICES, REGISTER_SESSION, SEND_RR_DATA, SEND_RR_DATA, SEND_RR_DATA,
            UNREGISTER_SESSION), receivedCommands);
    }

    @Test
    void aDeviceThatEndsTheSessionOverTheProbeIsNotReportedConnected() {
        startDevice(RrDataBehaviour.ENCAPSULATION_ERROR_AND_CLOSE);

        // Used to come back as a connection whose transport was already gone.
        PlcConnectionException e = assertThrows(PlcConnectionException.class,
            () -> new DefaultPlcDriverManager().getConnection(url()));
        assertTrue(e.getMessage().contains("during the EIP connect handshake"), e.getMessage());
    }

    private String url() {
        return "eip:tcp://127.0.0.1:" + serverSocket.getLocalPort() + "?big-endian=false&request-timeout-ms=2000";
    }

    private void startDevice(RrDataBehaviour behaviour) {
        deviceThread = new Thread(() -> {
            try (Socket socket = serverSocket.accept()) {
                socket.setSoTimeout((int) TimeUnit.SECONDS.toMillis(10));
                serve(socket, behaviour);
            } catch (IOException e) {
                // The driver closed the socket, or the test is over.
            }
        }, "fake-eip-device");
        deviceThread.start();
    }

    private void serve(Socket socket, RrDataBehaviour behaviour) throws IOException {
        InputStream in = socket.getInputStream();
        OutputStream out = socket.getOutputStream();
        while (true) {
            byte[] header = in.readNBytes(24);
            if (header.length < 24) {
                return;
            }
            ByteBuffer h = ByteBuffer.wrap(header).order(ByteOrder.LITTLE_ENDIAN);
            int command = h.getShort(0) & 0xFFFF;
            int length = h.getShort(2) & 0xFFFF;
            byte[] senderContext = new byte[8];
            h.get(12, senderContext);
            byte[] body = in.readNBytes(length);
            receivedCommands.add(command);

            switch (command) {
                case LIST_SERVICES -> out.write(encapsulation(command, 0, 0, senderContext, hex(
                    "0100" + "0001" + "1300" + "0100" + "2000" + "436F6D6D756E69636174696F6E7300")));
                case REGISTER_SESSION -> out.write(encapsulation(command, SESSION_HANDLE, 0, senderContext, hex("01000000")));
                case SEND_RR_DATA -> {
                    switch (behaviour) {
                        case ENCAPSULATION_ERROR_AND_CLOSE -> {
                            out.write(encapsulation(command, SESSION_HANDLE, 0x08, senderContext, new byte[0]));
                            out.flush();
                            return;
                        }
                        case EMPTY_CLASS_LIST -> out.write(encapsulation(command, SESSION_HANDLE, 0, senderContext,
                            rrData(hex("81000000" + "0000" + "0000"))));  // Get_Attribute_All: 0 classes, 0 available
                        case CIP_SERVICE_NOT_SUPPORTED -> {
                            int service = body[16] & 0x7F;  // interface handle 4, timeout 2, count 2, null item 4, item id 2, size 2
                            out.write(encapsulation(command, SESSION_HANDLE, 0, senderContext,
                                rrData(new byte[]{(byte) (0x80 | service), 0, 0x08, 0})));
                        }
                    }
                }
                case UNREGISTER_SESSION -> {
                    return;
                }
                default -> out.write(encapsulation(command, SESSION_HANDLE, 0x01, senderContext, new byte[0]));
            }
            out.flush();
        }
    }

    /** SendRRData body: interface handle, timeout, then a null address item and an unconnected data item. */
    private static byte[] rrData(byte[] cipService) {
        ByteBuffer b = ByteBuffer.allocate(16 + cipService.length).order(ByteOrder.LITTLE_ENDIAN);
        b.putInt(0).putShort((short) 0).putShort((short) 2);
        b.putShort((short) 0x0000).putShort((short) 0);
        b.putShort((short) 0x00B2).putShort((short) cipService.length).put(cipService);
        return b.array();
    }

    private static byte[] encapsulation(int command, int sessionHandle, int status, byte[] senderContext, byte[] body) {
        ByteBuffer b = ByteBuffer.allocate(24 + body.length).order(ByteOrder.LITTLE_ENDIAN);
        b.putShort((short) command).putShort((short) body.length).putInt(sessionHandle).putInt(status)
            .put(senderContext).putInt(0).put(body);
        return b.array();
    }

    private static byte[] hex(String hex) {
        byte[] out = new byte[hex.length() / 2];
        for (int i = 0; i < out.length; i++) {
            out[i] = (byte) Integer.parseInt(hex.substring(2 * i, 2 * i + 2), 16);
        }
        return out;
    }
}
