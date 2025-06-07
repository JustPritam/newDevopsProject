How to push?

1. generate an SSH key:
ssh-keygen -t ed25519 -C "your_email@example.com"
2. Copy the SSH key to your clipboard:
cat ~/.ssh/id_ed25519.pub
Go to GitHub, navigate to Settings > SSH and GPG keys, and click New SSH key. Paste your key and save.
3. Update Remote URL:
Change the remote URL to use SSH:
git remote set-url origin git@github.com:username/repository.git
4. To test the SSH Connection:
ssh -T git@github.com
5. To commit and push:
git add .
git commit -m "Test commit"
git push origin main
