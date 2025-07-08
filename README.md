# Learning

Programming languages - C/C++, JavaScript, Typescript, Solidity, Go 

---

# 🚀 Git & GitHub Commands Cheat Sheet

> 📚 A quick and essential reference for learning Git and GitHub step by step.

---

## 🔧 Setup Commands

```bash
git config --global user.name "Your Name"         # 👤 Set your Git username
git config --global user.email "you@example.com"  # 📧 Set your Git email
git config --list                             # 📃 Show current config

📁 Repository Basics


git init                                          # 🆕 Create a new local Git repo
git clone <repo-url>                              # 📥 Clone an existing repo

💾 Staging & Committing Changes


git status                                        # 🔍 Check status of files
git add .                                         # ➕ Stage all changed files
git add <file>                                    # 🎯 Stage specific file
git commit -m "your message"                      # 💬 Commit with message

🌿 Branching & Merging

bash
Copy
Edit
git branch                                        # 🌴 List all branches
git branch <branch-name>                          # 🌱 Create a new branch
git checkout <branch-name>                        # 🔁 Switch to a branch
git checkout -b <branch-name>                     # 🌟 Create and switch to new branch
git merge <branch-name>                           # 🔗 Merge a branch into current
⬆️ Push / ⬇️ Pull / 🔄 Sync with GitHub
bash
Copy
Edit
git remote add origin <repo-url>                  # 🔗 Link local repo to GitHub
git push -u origin main                           # 🚀 Push for the first time
git push                                          # 🔼 Push commits to GitHub
git pull                                          # 🔽 Pull latest changes
git fetch                                         # 📡 Fetch without merging
🧹 Undo Mistakes
bash
Copy
Edit
git restore <file>                                # ↩️ Undo changes to a file
git reset --soft HEAD~1                           # ⏪ Undo last commit (keep changes)
git reset --hard HEAD~1                           # 🧨 Undo commit and changes
git clean -fd                                     # 🧹 Delete untracked files/folders
📂 Viewing History & Diffs
bash
Copy
Edit
git log                                           # 🕓 Show commit history
git log --oneline                                 # 🧾 Compact log
git diff                                          # 🔍 Show file differences
🧑‍💻 GitHub CLI (Optional)
bash
Copy
Edit
gh auth login                                     # 🔐 Login to GitHub via CLI
gh repo create                                    # 📦 Create a new repo
gh repo clone <user>/<repo>                       # 📥 Clone via GitHub CLI
🧠 Pro Tips
bash
Copy
Edit
git reflog                                        # 🧠 See all HEAD changes (recover lost commits)
✅ Best Practices
📝 Write meaningful commit messages

⚠️ Always git pull before pushing

🧾 Use .gitignore to avoid tracking sensitive or irrelevant files

🌲 Keep your branches clean and organized

