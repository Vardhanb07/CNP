## title
Design of a Priority-Based Packet Scheduling Mechanism
## description
This project designs a packet scheduling system that assigns different priority levels to network packets.
## tcp/ip stack level
┌─────────────────────────────────────────┐
│ Application Layer (HTTP, VoIP, Video)   │
├─────────────────────────────────────────┤
│ Transport Layer (TCP, UDP)              │
├─────────────────────────────────────────┤
│ Network / Internet Layer (IP, Routing)  │ 
├─────────────────────────────────────────┤
│ ★ TRAFFIC CONTROL / SCHEDULER (QoS) ★  │ ◄── 
├─────────────────────────────────────────┤
│ Data Link / Network Interface (MAC/LLC) │ 
├─────────────────────────────────────────┤
│ Physical Layer (Bits, Wire, Radio)      │
└─────────────────────────────────────────┘
## type
software
