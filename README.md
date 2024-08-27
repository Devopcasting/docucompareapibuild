# DocCompareService

**DocCompareService** is a Windows service developed in Go that starts a Python FastAPI application for document comparison. This service runs in the background, allowing you to utilize the FastAPI application without manual intervention.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Features

- Runs a FastAPI application as a Windows service.
- Automatic startup and management of the Python application.
- Simple configuration through an `.ini` file.

## Requirements

- Go (1.16 or later)
- Python (3.12 or later)
- FastAPI framework and dependencies installed in your Python environment.

## Installation

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd DocCompareService
    ```
2. **Build the Go service:**

    ```bash
    go build -o doc_compare_service.exe main.go
    ```
3. **Create a Windows service:**

    ```bash
    sc create DocCompareService binPath= "C:\path\to\doc_compare_service.exe"
    ```
## Configuration

Create a project.ini file in the same directory as the executable with the following format:

```bash
[doc_compare_service]
python_path = C:\Users\xyz\AppData\Local\Programs\Python\Python312\python.exe
main_py_path = main.py
working_dir = D:\FastAPI-Document-Comparison-Project
```
- **python_path**: The full path to the Python executable.
- **main_py_path**: The main Python script to run (e.g., main.py).
- **working_dir**: The directory where the Python application resides.

## Usage

1. **Start Service**

    ```bash
    sc start DocCompareService
    ```

2. **Stop Service**

    ```bash
    sc stop DocCompareService
    ```

3. **Delete the service**

    ```bash
    sc delete DocCompareService
    ```

