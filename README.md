# `openai-ft-validate` CLI Tool

## Overview

`openai-ft-validate` is a Command-Line Interface (CLI) tool written in Go, designed to validate the structure of JSONL files used in OpenAI fine-tuning processes. This tool ensures that the input files adhere to OpenAI's formatting requirements, helping users avoid errors during fine-tuning.

---

## Features

- Validates JSONL files for structural correctness.
- Checks for required fields and their data types.
- Provides clear error messages for invalid entries.
- Ensures compatibility with OpenAI fine-tuning guidelines.
- Lightweight and easy to use.

---

## Installation

To install `openai-ft-validate`, ensure you have Go installed, then clone the repository and build the binary:

```bash
# Clone the repository
git clone https://github.com/Siddhesh-Agarwal/openai-ft-validate
cd openai-ft-validate

# Build the binary
go build -o openai-ft-validate
```

Alternatively, you can download a prebuilt binary (if available) from the [releases page](https://github.com/Siddhesh-Agarwal/openai-ft-validate/releases).

---

## Usage

The tool can be invoked directly from the command line. Below are the usage details:

### Basic Command

```bash
./openai-ft-validate <path-to-jsonl-file>
```

### Example

To validate a JSONL file named `training_data.jsonl`:

```bash
./openai-ft-validate training_data.jsonl
```

---

## Output

- **Success:** If the file is valid, you will see a message like:

    ```sh
    The JSONL is valid 🎉
    ```

- **Failure:** If the file contains errors, the tool will output details, such as:

    ```sh
    Line 3: Missing required field "prompt".
    Line 7: Invalid "role" programmer found
    The JSONL is invalid ❌
    ```

---

## Requirements

- Go 1.16+

---

## Development

For contributors, clone the repository and build the project:

```bash
git clone https://github.com/Siddhesh-Agarwal/openai-ft-validate
cd openai-ft-validate

# Build the binary
go build
```

Run tests to ensure functionality:

---

## License

This tool is licensed under [MIT License](https://github.com/Siddhesh-Agarwal/openai-ft-validate/blob/main/LICENSE).

---

## Contributions

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request. Ensure your code follows the established coding standards and includes tests for any new features.

---

## Contact

For any questions or issues, please open an issue on the [GitHub repository](https://github.com/Siddhesh-Agarwal/openai-ft-validate) or contact [Siddhesh Agarwal](mailto:siddhesh.agarwal@gmail.com).
