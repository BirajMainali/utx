# UTX Command Line Tool

## Overview

The UTX Command Line Tool is a versatile utility designed to perform various operations directly from the command line. It includes functionalities for string manipulation, date handling, network operations, JSON processing, password management, and GUID generation.

## Commands

| Command       | File         | Description                                                                 |
|---------------|--------------|-----------------------------------------------------------------------------|
| String        | `string.go`  | Provides operations for string manipulation including length calculation, comparison, case conversion, trimming, and displaying the current UTC time. |
| Date          | `date.go`    | Handles date and time operations.                                           |
| Network       | `network.go` | Provides functionalities for network-related operations.                    |
| JSON          | `json.go`    | Handles JSON data processing and manipulation.                              |
| Password      | `password.go`| Provides utilities for password generation and management.                  |
| GUID          | `guid.go`    | Generates and manages GUIDs (Globally Unique Identifiers).                  |

## Usage

To use the UTX Command Line Tool, execute the following command:

```bash
utx <command> <operation> [parameters]
```

Replace `<command>` with the desired command category (e.g., `string`, `date`), `<operation>` with the specific operation, and `[parameters]` with any required parameters.

## Conclusion

The UTX Command Line Tool is designed to streamline various command line operations, making it a valuable tool for developers and system administrators. 