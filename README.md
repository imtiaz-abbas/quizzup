# quizzup


### Installing Docker and Docker Compose

#### For Linux

1.  Visit the [Docker Website](https://docs.docker.com/install/linux/docker-ce/ubuntu/) to learn how to install it.
2.  Also install Docker Compose using instructions [here](https://docs.docker.com/compose/install/)

#### For MacOS

1.  Install Docker for Mac from [here](https://docs.docker.com/docker-for-mac/)
2.  Also install Docker Compose using instructions [here](https://docs.docker.com/compose/install/)


### Setting up the environment
 
 Create a file named ```.env```, copy contents from the file ```Sample.env``` to ```.env```

### Pulling up the application

To pull up the complete application you can run:

```
# To Start the server
$ docker-compose up -d

# To see logs
$ docker-compose logs -f

#To stop the server 
$ docker-compose down

```

Learn more about how to use docker-compose from [here](https://docs.docker.com/compose/reference/).

#### Access to services

The services from this compose file can be accessed with the following urls:

Server Url: http://localhost:8080

