
# done
![Deployment Strategies in Kubernetes 1](img/1.png)
![Deployment Strategies in Kubernetes 2](img/2.png)


# You Can Run the app 
    go mod init github.com/freddymac124/upwork-go

    go mod tidy

[then]

    go run cmd/main.go


# Or You Can Run with Docker Composer 

    docker image rm ahmedkhalaf666/gogolfapp:latest

    docker-compose up -d

[to-stop-docker-compser]
    
    docker-compose down

# Open The API URL 
# http://localhost Or http://localhost:80

# -------------------------------  NOTIFICATION --------------------------------------------
 # use POSTMAN Program
 # Select WebSocket Request

  ![3](img/3.png)



ws://localhost:8080/notification/userid
# if user id = 1

ws://localhost:8080/notification/1

# if user id = 2
ws://localhost:8080/notification/2

  ![4](img/4.png)

.. [Note] To Craete User First Before you start Connect To the Notification Service
# if You Created Remember to login first

  ![4](img/4.png)
  ![5](img/5.png)
# ☑ auth // varyfied profile account
  ![6](img/6.png)
# ☑ friends // send & add ew friend = > recever user
  ![7](img/7.png)


etc...
[notification-messages]
# ☑ friends // Accept the Friend Request = > sender user
# ☑ post // when some one add comment to the user post
# ☑ post // when some one like user post
# ☑ game // notifiy user friends when user start new game
# ☑ game // notifiy user friends wehn user finish the game
# ☑ score // notify user when some one attested added score




now we need to change score round to hole
and adding game Course Name
righ ? is there is any thing else