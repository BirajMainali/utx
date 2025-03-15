# UTX - A Command-Line Utility Tool

UTX is a lightweight and versatile command-line utility providing various functionalities, including date manipulation, GUID generation, JSON processing, networking, password generation, and string operations.

## 📥 Installation

1. Download the `utx` executable.
2. Place it in a directory included in your system's `PATH`. (e.g for windows `C:\Windows\System32\`)
3. Run `utx --help` to check available commands.

---

## 🛠️ Usage

### **1️⃣ Date Utilities**
| Command | Description | Options |
|---------|------------|---------|
| `utx date now` | Returns the current local date and time. | N/A |
| `utx date utc` | Returns the current UTC date and time. | N/A |
| `utx date add <duration>` | Adds a specified duration (e.g., `1d`, `24h`, `30m`) to the current time. | `d`, `h`, `m` |
| `utx date diff <date1> <date2>` | Calculates the difference between two dates. | N/A |

---

### **2️⃣ GUID (UUID) Generation**
| Command | Description | Options |
|---------|------------|---------|
| `utx guid v1` | Generates a time-based UUID (Version 1). | `--upper` (uppercase UUID) |
| `utx guid v4` | Generates a randomly generated UUID (Version 4). | `--upper` (uppercase UUID) |
| `utx guid v5 <namespace> <name>` | Generates a UUID (Version 5) based on a namespace and name. | `--upper` (uppercase UUID) |

---

### **3️⃣ JSON Utilities**
| Command | Description | Options |
|---------|------------|---------|
| `utx json validate <json>` | Validates if the given input is a proper JSON format. | N/A |
| `utx json format <json>` | Formats and pretty-prints a JSON input. | N/A |

---

### **4️⃣ Network Utilities**
| Command | Description | Options |
|---------|------------|---------|
| `utx net local-ip` | Retrieves the local machine's IP address. | N/A |
| `utx net public-ip` | Fetches the machine’s public IP address. | N/A |
| `utx net ping <hostname/IP>` | Sends an ICMP ping request to a target hostname or IP. | `-c <count>` (number of pings) |
| `utx net dns <domain>` | Resolves a domain name to its IP addresses. | N/A |
| `utx net interfaces` | Lists all available network interfaces. | N/A |

---

### **5️⃣ Password Generator**
| Command | Description | Options |
|---------|------------|---------|
| `utx pwd <length>` | Generates a secure password of the specified length. | `--special`, `--numbers`, `--uppercase` |

---

### **6️⃣ String Utilities**
| Command | Description | Options |
|---------|------------|---------|
| `utx str length <string>` | Returns the length of the string. | N/A |
| `utx str compare <str1> <str2>` | Compares two strings and returns if they are equal. | N/A |
| `utx str upper <string>` | Converts the string to uppercase. | N/A |
| `utx str lower <string>` | Converts the string to lowercase. | N/A |
| `utx str trim <string>` | Trims leading and trailing spaces from the string. | N/A |

---

## 🔍 Examples
```sh
utx date now
# Output: 2025-03-15 14:32:10

utx guid v4
# Output: 3f4d2a76-9f36-4c9e-b4a1-57c1db84a6e4

utx json validate '{"name": "John"}'
# Output: Valid JSON

utx net public-ip
# Output: 192.168.1.10

utx pwd 16 --special --numbers
# Output: A@9fL3$z!Qx72Nv#
