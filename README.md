# UTX - A Command-Line Utility Tool

UTX is a lightweight and versatile command-line utility providing various functionalities, including date manipulation, GUID generation, JSON processing, networking, password generation, and string operations.

## 📥 Installation

1. Download the `utx` executable.
2. Place it in a directory included in your system's `PATH`.
3. Run `utx --help` to check available commands.

---

## 🛠️ Usage

### **1️⃣ Date Utilities**
| Command | Description |
|---------|------------|
| `utx date now` | Returns the current local date and time. |
| `utx date utc` | Returns the current UTC date and time. |
| `utx date add <duration>` | Adds a specified duration (e.g., `1d`, `24h`, `30m`) to the current time. |
| `utx date diff <date1> <date2>` | Calculates the difference between two dates. |

---

### **2️⃣ GUID (UUID) Generation**
| Command | Description |
|---------|------------|
| `utx guid v1` | Generates a time-based UUID (Version 1). |
| `utx guid v4` | Generates a randomly generated UUID (Version 4). |
| `utx guid v5 <namespace> <name>` | Generates a UUID (Version 5) based on a namespace and name. |
| `utx guid v4 --upper` | Generates an uppercase UUID. |

---

### **3️⃣ JSON Utilities**
| Command | Description |
|---------|------------|
| `utx json validate <json>` | Validates if the given input is a proper JSON format. |
| `utx json format <json>` | Formats and pretty-prints a JSON input. |

---

### **4️⃣ Network Utilities**
| Command | Description |
|---------|------------|
| `utx net local-ip` | Retrieves the local machine's IP address. |
| `utx net public-ip` | Fetches the machine’s public IP address. |
| `utx net ping <hostname/IP>` | Sends an ICMP ping request to a target hostname or IP. |
| `utx net dns <domain>` | Resolves a domain name to its IP addresses. |
| `utx net interfaces` | Lists all available network interfaces. |

---

### **5️⃣ Password Generator**
| Command | Description |
|---------|------------|
| `utx pwd <length>` | Generates a secure password of the specified length. |
| `utx pwd <length> --special` | Includes special characters in the generated password. |
| `utx pwd <length> --numbers` | Includes numbers in the generated password. |
| `utx pwd <length> --uppercase` | Includes uppercase letters in the generated password. |

---

### **6️⃣ String Utilities**
| Command | Description |
|---------|------------|
| `utx str length <string>` | Returns the length of the string. |
| `utx str compare <str1> <str2>` | Compares two strings and returns if they are equal. |
| `utx str upper <string>` | Converts the string to uppercase. |
| `utx str lower <string>` | Converts the string to lowercase. |
| `utx str trim <string>` | Trims leading and trailing spaces from the string. |

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
