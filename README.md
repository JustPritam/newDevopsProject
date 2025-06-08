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

Another way to push:
1. Go to "Settings" > "Developer Settings"
2. Click on "Create a Personal Access Token" and then after creating copy it.
3. Now you have to set the remote origin URL: git remote set-url origin <YOUR_REPO>
4. Then when you push the commit in username provide your username and in password provide the PAT

How to fix the no space issue in Jenkins in EC2?

1. Go to "Manage Jenkins" > "Nodes" 
2. Click on the settings button in the right most side of the "Built-In Node"
3. In the "Node Properties", under "Disk Space Monitoring Thresholds", change all sizes to 0MB.
4. Click on "Save".
5. Restart Jenkins.

How to create a pipeline?

1. Select "New item".
2. Enter the pipeline name and select "Freestyle Pipeline".
3. Scroll down, under "Build steps" > "Add Build Steps" > select "Execute Shell"
4. Write down the docker commands there:
   a. sudo docker build -t <IMAGE-NAME> -f <LOCATION-OF-THE-DOCKERFILE>/Dockerfile <LOCATION-OF-THE-DOCKERFILE>
   -f specifies the Dockerfile path.
   The second path is the build context (folder where your Go files and go.mod live).
   b. sudo docker run -d -p port:port <IMAGE-NAME>
5. Click on "Save".
6. Click on "Build Now".

Plugins needed: Git client, Docker Pipeline
