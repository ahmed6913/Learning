
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
to create a file use '<echo>' file name

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

```

# 🚀 Git & GitHub Commands Cheat Sheet

 📚 A quick and essential reference for learning Git and GitHub step by step.

---

## 🔧 Setup Commands
```bash
git --version                                     # To check if git is installed/ version 

git config --global user.name "Your Name"         # 👤 Set your Git username
git config --global user.email "you@example.com"  # 📧 Set your Git email
git config --list                                 # 📃 Show current config
```
📁 Repository Basics
```bash
git init                                          # 🆕 Create a new local Git repo
git clone <repo-url>                              # 📥 Clone an existing repo
```
💾 Staging & Committing Changes
```bash
git status                                        # 🔍 Check status of files
git add .                                         # ➕ Stage all changed files
git add <file>                                    # 🎯 Stage specific file
git commit -m "your message"                      # 💬 Commit with message
```
🌿 Branching & Merging
```bash
git branch                                        # 🌴 List all branches
git branch <branch-name>                          # 🌱 Create a new branch
git checkout <branch-name>                        # 🔁 Switch to a branch
git checkout -b <branch-name>                     # 🌟 Create and switch to new branch
git merge <branch-name>                           # 🔗 Merge a branch into current
⬆️ Push / ⬇️ Pull / 🔄 Sync with GitHub

git remote add origin <repo-url>                  # 🔗 Link local repo to GitHub
git push -u origin main                           # 🚀 Push for the first time
git push                                          # 🔼 Push commits to GitHub
git pull                                          # 🔽 Pull latest changes
git fetch                                         # 📡 Fetch without merging
```
🧹 Undo Mistakes
```bash
git restore <file>                                # ↩️ Undo changes to a file
git reset --soft HEAD~1                           # ⏪ Undo last commit (keep changes)
git reset --hard HEAD~1                           # 🧨 Undo commit and changes
git clean -fd                                     # 🧹 Delete untracked files/folders
```
📂 Viewing History & Diffs
```bash
git log                                           # 🕓 Show commit history
git log --oneline                                 # 🧾 Compact log
git diff                                          # 🔍 Show file differences
```
🧑‍💻 GitHub CLI (Optional)
```bash
gh auth login                                     # 🔐 Login to GitHub via CLI
gh repo create                                    # 📦 Create a new repo
gh repo clone <user>/<repo>                       # 📥 Clone via GitHub CLI
```
🧠 Pro Tips
```bash
git reflog                                        # 🧠 See all HEAD changes (recover lost commits)
```
✅ Best Practices

- 📝 Write meaningful commit messages
- ⚠️ Always git pull before pushing
- 🧾 Use .gitignore to avoid tracking sensitive or irrelevant files
- 🌲 Keep your branches clean and organized

