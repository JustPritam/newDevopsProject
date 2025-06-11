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

Plugins needed: Docker Pipeline

If this error: **sudo: a terminal is required to read the password; either use the -S option to read from standard input or configure an askpass helper
sudo: a password is required** is shown then follow these:
_Allow Jenkins user to run docker without sudo_
This is the recommended and cleanest way.
1. Add Jenkins user to the docker group:
**sudo usermod -aG docker jenkins**
2. Restart Jenkins and re-login to apply group changes:
**sudo systemctl restart jenkins**

How to use k8s?

1. Install minikube
2. Download minikube and start it by: **minikube start --driver=docker**
3. Create a _pod.yml_ (Defines a single pod manually for testing)
4. Create a _deploy.yml_ (Defines a Deployment which manages pod replicas and handles self-healing)
5. Create a _service.yml_ (Exposes your Deployment so it can be accessed from outside (usually via NodePort or LoadBalancer))
6. To start the pod: **kubectl apply -f pod.yml**
7. To start the deployment: **kubectl apply -f deploy.yml**
8. To start the service: **kubectl apply -f service.yml**
9. Then run: **kubectl get svc** (Lists all services along with their ports and cluster IPs)
10. Run: **minikube service <SERVICE_NAME> --url** and you will get an address (Provides an accessible URL to your service)
11. Use **curl -l <ADDRESS_THAT_YOU_GOT_IN_PT6>** (Checks if the service is up and running)
12. Then run: **sudo nano /etc/hosts** and add the IP-address and add a Name suppose **todoapp.com**. Save and exit. (Add a line like the IP-address todoapp.com to access using a name instead of IP)
13. Run: **curl -l todoapp.com:30007** (Useful for simulating production-like access using domain names)
