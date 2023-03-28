# CodeBox 1.0

![GitHub go.mod Go version](https://img.shields.io/badge/Golang-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![RabbitMQ](https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white)


## About
CodeBox is an online code executor and judging platform designed to provide a user-friendly environment for developers and programmers as practice platform for data structures and algorithms (DSA). It is similar to Leetcode , compile and execute code online.

![SC1](https://user-images.githubusercontent.com/42716731/228065518-baa1ffac-a9ee-4470-9b5e-8c0c776953e5.png)

## Features

+ Based on Docker; One-click deployment.
+ Separated backend and frontend; Modular programming; Micro service.
+ Multi-language support: `C++`, `Java`, `Python3`
+ Sandbox for running quick snipts.
+ Provides insights of result.
+ Maintains history of code problems tried.

## Working

![Blank diagram - Page 1(1)](https://user-images.githubusercontent.com/42716731/228130954-a7f6b212-75fd-4234-a48c-8613cb158645.png)

+ When user hit Run on the screen, client makes a packs the code written and problem id this payload is called `Task`, this is sent to te api server which stores it in DB. 
+ Api server after successfully stroing it in DB adds it to the RMQ which delivers the `task` aka message to availble CodeRunner.
+ CodeRunner is a Sandbox env, It accepts the task mark it `Running` in DB. CR compiles the code and runs it with standard inputs. If code is DSA type it matches the code output with expected output and produces the result. Now if there is malicious code it would just crash the docker container, for that enable restart on crash in docker. There is additional measure which timeouts the code execution if it exceeds X amount of seconds/minutes and results TLE.  


## Installation
1. Install the necessary dependencies

    #### Fedora
    ```bash
    sudo dnf groupinstall c-development
    sudo dnf install python3
    pip3 install --upgrade pip
    sudo dnf install java-1.8.0-openjdk.x86_64
    ```
     #### Ubuntu
    ```bash
    sudo apt-get update
    sudo apt install build-essential
    sudo apt-get install -y vim python3-pip curl git
    pip3 install --upgrade pip
    sudo apt install default-jdk
    ```
    
2. Install docker and docker compose from [Here](https://docs.docker.com/engine/install/)
3. Create a postgres database server using docker. Set a user and password for the db which we will use. 
    ```bash 
    docker run --name postgres-db -e POSTGRES_USER=vishvajeet POSTGRES_PASSWORD=xyzPassword -p 5432:5432 -d postgres
    ```
4. Now we need to start a rabbit mq instance which will be the message queue for our application.
    
    ```bash
    docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management
    ```
5. cd into the root of the project and create a `.env` file to store our db and queue credentials.
    ```
    PORT=8080
    
    RM_QUEUE=codeBox
    RMQ_URL=amqps://xxxuserxxx:xxxpassxxx@xxxhostxxx/xxxportxxx
    RMQ_EXCHANGE_NAME=CodeBoxExchange
    RMQ_QUEUE_NAME=CodeBoxQueue
    
    DB_URL=postgresql://<user>:<Db Pass>@<host add>:<port>/<db>
    ```
6. From the root of the project, start the service
    ```bash
    cd CodeBox
    docker compose up
    ``` 
   
## Screenshots

### Home:

![SC1](https://user-images.githubusercontent.com/42716731/228065518-baa1ffac-a9ee-4470-9b5e-8c0c776953e5.png)

### When solution is accepted: 
![SS2](https://user-images.githubusercontent.com/42716731/228065935-1bfc8826-d6a5-48e7-a18c-1782dde53911.png)

### If solution is rejected due to wrong answer or TLE:
![SS3](https://user-images.githubusercontent.com/42716731/228066123-c05f2d08-22de-4243-8a9f-fbdab28adc94.png)

### Sandbox to run small snippets:
![SS4](https://user-images.githubusercontent.com/42716731/228066377-6d9a2787-946e-478b-8049-ab7908c9d6d1.png)

### History tab:
![ss5](https://user-images.githubusercontent.com/42716731/228066529-a107cd5d-86af-453b-a430-b595f4e58d5b.png)

## Thanks

+ Project is yet in developement, it is right now in the bare minimum working codition. I would try imporve uppon the code quality if time permits. 
+ Special thanks to [Charan Vasu](https://github.com/charan1973) and [Gaurav](https://github.com/darkfusion90), for resolving minor challenges that arose during the project 

