# Pentest-Tools (Educational Lab)

This repository contains a collection of proof-of-concept scripts and networking tools developed for educational purposes. The primary goal of this project is to explore **Go concurrency (Goroutines/Channels)**, **Network programming (TCP/UDP)**, and **Python automation** within the context of cybersecurity and systems research.

> [!IMPORTANT]  
> **Disclaimer:** These tools are for educational and ethical testing purposes only. They are not intended for professional use or illegal activities.

---

## 🛠 Project Overview

### 1. Bind-Shell (`netcat.go`)
A basic implementation of a TCP bind shell. It demonstrates how to handle network connections and redirect standard input/output/error to a shell process (`/bin/bash` or `cmd.exe`).
* **Key concepts:** `net` package, `os/exec` command execution, I/O piping.

### 2. Goroutine-Stress (`stress.go`)
A performance-testing utility designed to observe Go's scheduler behavior under heavy load. It spawns hundreds of thousands of Goroutines to test system resource management.
* **Key concepts:** Lightweight concurrency, synchronization, `time.Timer`.

### 3. MD5-Brute-Force (`brute-force.go`)
A recursive MD5 hash cracker that demonstrates a brute-force approach across a specific charset.
* **Key concepts:** Recursion, `crypto/md5` hashing, string manipulation in Go.

### 4. Net-Guard (`NetGuard.py`)
A Python-based monitoring script designed to ensure connection privacy. It tracks network status, IPv6 leaks, and VPN/Proxychains connectivity, using audio alerts (via Pygame) for status changes.
* **Key concepts:** Subprocess management, threading, network automation, system-level monitoring.

### 5. Port-Scanner (`port-scanner.go`)
A high-performance, concurrent TCP port scanner. It utilizes a worker-pool pattern with channels to efficiently scan the most common ports.
* **Key concepts:** Worker pools, buffered channels, `net.Dial` timeouts, sorting algorithms.

### 6. Reconnaissance-Tool (`reconnaissance.go` & `server.go`)
A client-server architecture for system information gathering.
* **Client (`reconnaissance.go`):** Collects host info (OS, Arch, IP, User) and transmits it via HTTP POST in JSON format.
* **Server (`server.go`):** A REST API endpoint that receives, decodes, and logs the incoming system data.
* **Key concepts:** JSON Marshalling/Unmarshalling, HTTP Client/Server, `runtime` and `os/user` packages.

### 7. TCP-Proxy (`tcp-proxy.go`)
A simple TCP stream redirector that acts as a middleman between a local port and a remote destination.
* **Key concepts:** `io.Copy`, full-duplex communication, bidirectional data flow.

---

## 🚀 Technical Goals
* **Concurrency Mastery:** Implementing efficient multi-threaded operations using Go's primitives.
* **Protocol Research:** Understanding the handshake and data exchange process in TCP/HTTP.
* **System Integration:** Interacting with OS-level commands and environment variables.

## 📝 Requirements
* **Go:** 1.18+ (for Go-based tools)
* **Python:** 3.x (with `pygame`, `termcolor` and `argparse`)
* **OS:** Primarily tested on Linux (some tools require `/bin/bash` or `NetworkManager`).

---
*Created as part of a personal learning journey in Cybersecurity and Go Development.*
