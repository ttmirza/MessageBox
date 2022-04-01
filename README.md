# Message Box
### A REST API that allows registered users to send & receive individual and group messages.
## Build
Two ways to run application:
1. `make go` at root level of application directory --> API is listening and serving on port 3001 (building using Go, recommended)
2. `make docker` at root level of application directory --> API is listening and serving on port 3001 (building using Docker)
## Data Model
There are essentially five sql tables that define the basic data model for MessageBox.
1. Users - Every time a user is created, a user is stored in the Users table, username must be unique.

2. Groups - Every time a group is created, a group is stored in the Groups table, group name must be unique.

3. GroupUsers - Every time a group is created, an entry linking each group member to group is stored in the GroupUsers table (one entry per username).

4. Messages - A message can be an original message, a reply, a group message, or a user message. The message model is used for all types of messages. Everytime a message is created, a message is stored in the Messages table.

5. Replys - Every time a reply is sent to an original message, an entry in the Replys table is created. The entry links the original message to the reply being crafted. This is used to return all replies for a given message.

## Dependencies/Frameworks
### Gin
- `go get -u github.com/gin-gonic/gin` https://github.com/gin-gonic/gin
- Provides performance that is up to 40 times faster than standard library net/http default mux.
- Out of the box methods allow for ease of adoption & quick delivery, and provide some useful logging.
### Gorm
- `go get -u https://github.com/go-gorm/gorm` https://github.com/go-gorm/gorm
- GORM enabled the use of native programming paradigms (Go) to map data to SQL.
- An ORM transforms database information to Golang objects and vice-versa, instead of writing SQL queries directly, ability to map data cleaner using Golang structs.
### Uuid
- `go get github.com/google/uuid` github.com/google/uuid
- Used as primary_key for MessageID since uuids are effectively unique.
## Database
- The data model has relationships between each other; e.g. group to members, members to messages, messages to replies, user to messages, etc. Hence, a relational database was chosen.
- SQLite3 was chosen for simplicity due to it being a small, fast, self-contained, highly-reliable in memory SQL database.
- Ideally if needed, replace SQL db with more performant SQL, consideration for graph db as well.
## Improvements/Considerations/TODO:
### TODO:
- Add standardized logging (some inherited from gin & gorm packages), logging at appropriate levels
- Add router file to handle all routes instead of in main
- Add Config package to configure environmental variables (db connection string, log level, ports)
- Make connection to DB more secure/safe
- Add feature to utilize LastLoggedIn field in User data model to update each time a user signs in
- Standardize JSON Responses between routes
- Add unit/integration testing --> will need to mock db calls
- Add database indexes to fields that are frequently queried to increase performance
### Considerations:
- For the sake of simplicity basic validations & error handling were done on which improvements can be made
### Improvements:
- Potentially remove Group table. While it does hold all groups, we can extend and utilize the GroupUser table for this use case
- Add more validations & error handling
- Docker Image can be drastically reduced in size
- Dockerfile can take advantage of multistage build to reduce build time
- Reduce the amount of database calls by replacing with Joins between tables, this would also result in less data held in memory
