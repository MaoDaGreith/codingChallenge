# codingChallenge

## Description
A simple web microservice that can fetch different kind of data.

Available routes:
- GET `/getUser/id`
will return UserID info
- GET `/countActions/id`
will return the total count of action based on userID
- GET `/nextAction?type=action_type`
will return ordered chronologically data
- GET `/referralIndex`
will return the userID:refered_userid

More improvements could be done, with the usage or libraries and other tools.

## How To Run

Use `go run main.go`
