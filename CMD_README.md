# 🖥️ Essential CMD Commands for Software Developers

This guide lists useful Windows Command Prompt (CMD) commands every developer should know for file navigation, automation, networking, system tasks, and more.

---

## 📁 1. Directory Navigation

| Command             | Description                           |
|---------------------|---------------------------------------|
| `cd`               | Show current directory                |
| `cd foldername`    | Enter a folder                        |
| `cd ..`            | Go up one level                       |
| `cd /`             | Go to root directory                  |
| `dir`              | List files and folders                |
| `tree`             | Display folder structure              |
| `cls`              | Clear the screen                      |

---

## 📄 2. File & Folder Management

| Command                           | Description                              |
|-----------------------------------|------------------------------------------|
| `mkdir foldername`               | Create a new folder                      |
| `rmdir foldername /s /q`         | Delete folder (and contents) silently    |
| `del filename.ext`               | Delete a file                            |
| `copy file1.txt file2.txt`       | Copy a file                              |
| `xcopy /s source dest`           | Copy folder with subfolders              |
| `move file.txt foldername`       | Move file to folder                      |
| `ren oldname.txt newname.txt`    | Rename a file                            |
| `type filename.txt`              | Display content of a text file           |
| `echo Hello > file.txt`          | Write text to a file (overwrite)         |
| `echo More >> file.txt`          | Append text to a file                    |

---

## 🧠 3. System Information

| Command            | Description                        |
|--------------------|------------------------------------|
| `systeminfo`      | Detailed system information        |
| `hostname`        | Display computer name              |
| `set`             | List environment variables         |
| `echo %USERNAME%` | Show current user name             |
| `echo %CD%`       | Show current directory             |
| `ver`             | Display Windows version            |
| `title MyWindow`  | Change CMD window title            |

---

## 🧰 4. Task Management

| Command                      | Description                        |
|------------------------------|------------------------------------|
| `tasklist`                  | List running processes             |
| `taskkill /IM app.exe /F`   | Kill a process by name             |
| `taskkill /PID 1234 /F`     | Kill a process by PID              |

---

## 🌐 5. Network Commands

| Command                    | Description                             |
|----------------------------|-----------------------------------------|
| `ipconfig`                | View IP address and network info        |
| `ipconfig /flushdns`      | Clear DNS cache                         |
| `ping example.com`        | Test connection to domain or IP         |
| `tracert example.com`     | Trace route to a host                   |
| `netstat -an`             | Show open ports and connections         |
| `nslookup example.com`    | Get DNS info for a domain               |

---

## ⚙️ 6. Automation & Utilities

| Command                | Description                              |
|------------------------|------------------------------------------|
| `start filename.ext`  | Open file in default program             |
| `start .`             | Open current folder in File Explorer     |
| `pause`               | Pause script execution                   |
| `exit`                | Close Command Prompt                     |
| `help`                | List all available commands              |
| `command /?`          | Show help for a specific command         |

---

## 📝 7. Batch Scripting Basics

Use `.bat` files to automate CMD tasks.

```bat
@echo off
echo Running my project...
cd C:\MyApp
start .
pause
