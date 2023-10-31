# BankPoc

[![Build Status](https)]
[![License](https:)[If any]

## Introduction

BankPoc is a Golang-based application that demonstrates API creation using CRUD operations. It is designed as a microservices-based project with mailing capabilities.

## Features

- Feature 1
- Feature 2
- ...

## Getting Started

### Installation

To install BankPoc, follow these steps:

1. Clone this repository.
2. Install the required dependencies.
3. Configure the settings (if necessary).
4. Build the application using `go build`. It creates an executable file without any extension in macOS and Linux, but creates an executable with `.exe` in Windows.
5. Run the application.

## Getting Started

### Installation

To install BankPoc, follow these steps:

1. **Clone this Repository**:

   ```sh
   git clone https://github.com/yourusername/BankPoc.git
   cd BankPoc
   ```

2. **Install Required Dependencies**:

   Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

3. **Configuration (if Necessary)**:

   If your application requires any configuration, specify where users should add configuration files or environment variables. For example:
   - Create a `.env` file in the project root.
   - Add your configuration settings to the `.env` file.

4. **Build the Application**:

   - On macOS and Linux:

     ```sh
     go build  -o myapp
     #go build  -o myapp.exe
     ```

   - On Windows:

     ```sh
     go build -o myapp.exe
     ```

   This command will create an executable file named `myapp` (without an extension on macOS and Linux) or `myapp.exe` (on Windows).

5. **Run the Application**:

   - On macOS and Linux:

     ```sh
     ./myapp
     ```

   - On Windows:

     ```sh
     .\myapp.exe
     ```

By following these steps, users should be able to clone, set up, build, and run your Go application with ease.

### Keycloak
Step 1: Create a user. You can take its carl bash from developer console network tab fetch/XHR tab by pressing F12, and paste it in Postman to generate the URL endpoint automatically.

Step 2: Create password for the created user (non-temporary password).

Step 3: Create a client of "OpenID Connect" type and keep track of its secret found in dashboard.

Step 4: To fetch token for login, hit the Keycloak API endpoint:
             http://localhost:8080/realms/bankpoc/protocol/openid-connect/token
            
             with body parameters ( x-www-form-urlencoded ) as:
             client_id=<to be taken from Keycloak API> #bankpoc
             username=<username>
             password=<password>
             grant_type=password #Value is literal "password" here.
             client_secret=<value> # Taken from the Keycloak dashboard.
         #  Body parameters ( x-www-form-urlencoded ) while hitting the keycloak API: Key values in params but were to be passed in body section x-www-form-urlencoded subsection, then only they work

Step 5: Take the token from the response generated in the above step and use in hitting the Keycloak API:
            http://localhost:8080/realms/bankpoc/protocol/openid-connect/token/introspect
          with body parameters ( x-www-form-urlencoded ) as:
             client_secret=<value> # Taken from the Keycloak dashboard.
             token=<token from above token api response>
             client_id=<taken from Keycloak dashboard> #bankpoc
            
          #  Body parameters ( x-www-form-urlencoded ) while hitting the keycloak API: Key values in params but were to be passed in body section x-www-form-urlencoded subsection, then only they work.



### Usage

...

## Configuration

If your application requires any configuration, specify where users should add configuration files or environment variables. For example:

1. Create a `.env` file in the project root.
2. Add your configuration settings to the `.env` file.

## API Documentation

...

## Contribution Guidelines

...

## License

[GitHub Repository](https://github.com/manmohansharma21/BankPoc)


### Rebasing
 Use 'esc' key to come out of 'insert' mode and type ':wq' to save and exit.




 #ADD

  TO RUN docker-compose on MAC in case of docker command not found error:

ls -l /usr/local/bin | grep docker-compose
ls /usr/local/bin | grep docker-compose

export PATH="/usr/local/bin:$PATH"

/usr/local/bin/docker-compose up

source ~/.bashrc

It looks like you were able to resolve the issue by sourcing your .bashrc file. This likely refreshed your environment variables and made the docker-compose command available. The Docker Compose command should now work as expected. Run ut after setting the environment, then it works.
The source command, when used in a shell like Bash or Zsh, is used to execute commands from a file within the current shell session. In your case, when you run source ~/.bashrc, it executes the commands found in the ~/.bashrc file, which is typically used to set up environment variables, configure your shell, and define various aliases and functions.

docker-compose up
docker-compose down
docker-compose -f mysql-service.yaml up -d
docker-compose -f docker-compose.yaml -f mysql-service.yaml up -d
docker-compose -f docker-compose.yaml -f mysql-service.yaml down


If modules not found then create build file and then try to run the same.


* If the docker desktop is not running then, this will give this output:
Manmohans-MacBook-Air:docker manmohansharma$ source ~/.bashrc
Manmohans-MacBook-Air:docker manmohansharma$ docker-compose up
Cannot connect to the Docker daemon at unix:///Users/manmohansharma/.docker/run/docker.sock. Is the docker daemon running?
