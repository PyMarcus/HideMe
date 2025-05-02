# 🔒 hideme

A simple CLI tool written in Go to **encrypt and decrypt** text files using a password.

## 🚀 Features

- 🔒 Encrypt files securely with a password
- 🔓 Decrypt files back using the same password
- Easy-to-use command-line interface

---

## 📦 Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/PyMarcus/hideme.git
cd hideme
go build -o hideme
```

🔒 Encrypt a file

    ./hideme -f path/to/your/file.txt -e -p yourPassword

🔓 Decrypt a file

    ./hideme -f path/to/your/file.txt -d -p yourPassword
