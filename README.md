# Mouse Phone Server
Mouse Phone Server is a remote mouse which allows you to control your computer's cursor and mouse click by using your smartphone like a wireless touchpad through your local network.

This repository contains the Desktop Agent code (mousephoneserver). This server was written in Go which runs in your computer, listening your commands sent by the mobile application and translate them in mouse actions.

## 🎯 Overview
Ideal for those who give presentations, use their computer as a media center, or are simply looking for a more convenient way to interact with their PC from a distance. The goal is to provide a fast connection, low latency, and an intuitive user experience.

## 🏗️ Architecture
The system is based on a client-server model that communicates over a local network (Wi-Fi) using the WebSocket protocol.

Desktop Agent (This Repository): A Go server that creates a WebSocket endpoint. It is responsible for receiving messages, interpreting them, and using operating system APIs to control the mouse.

Mobile App (Client): An application (currently for Android) that captures user gestures on the screen, formats them into JSON, and sends them to the Desktop Agent.

### Communication Diagram
```
+--------------------------+                            +---------------------------+
|   Smartphone (Android)   |                            |   Desktop (PC/Mac/Linux)  |
|                          |                            |                           |
| +----------------------+ |    WebSocket via Wi-Fi     | +-----------------------+ |
| |   App "Remote Mouse"   | |<-----------------------> | |    Desktop Agent      | |
| |                      | |     (Local Network)        | |   (mousephoneserver)  | |
| | +------------------+ | |                            | | +-------------------+ | |
| | |   UI (Touchpad)  | | |                            | | |    WebSocket      | | |
| | +------------------+ | |                            | | |     Server        | | |
| | | Client WebSocket | | |                            | | +-------------------+ | |
| | +------------------+ | |                            | | |  Mouse Controller | | |
| +----------------------+ |                            | | |     (OS API)      | | |
|                          |                            | +-----------------------+ |
+--------------------------+                            +---------------------------+
```

## 📡 Communication Protocol
Communication is handled with simple JSON messages over a WebSocket connection (ws://).

| Action      | Example JSON Message                                |
| ----------- | ----------------------------------------------------|
| Mouse Move  | `{"type": "move", "dx": 10, "dy": -5}`              |
| Click       | `{"type": "click", "button": "left"}`               |
| Scroll      | `{"type": "scroll", "delta": -120}`                 |

## 🚀 How It Works
1. Start the Desktop Agent on your computer.
2. The agent will display the local IP address and a QR code for connection.
3. Open the app on your smartphone and scan the QR code (or type the IP).
4. The WebSocket connection is established.
5. Move your finger on the phone's screen to control the computer's mouse!
