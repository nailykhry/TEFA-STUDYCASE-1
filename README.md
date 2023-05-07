# TEFA-STUDYCASE-1

### A. STUDY CASE 1
This project is created to fulfill the TEFA task. The API project is implemented using Golang with the Fiber framework, MongoDB as the database, JWT for authentication, and implementing the SOLID principle. 

### B. Folder Structure

```
📦TEFA-STUDYCASE-1
 ┣ 📂controllers
 ┃ ┣ 📜auth.go
 ┃ ┣ 📜content.go
 ┃ ┣ 📜task.go
 ┃ ┣ 📜userChat.go
 ┃ ┗ 📜userSub.go
 ┣ 📂database
 ┃ ┗ 📜db.go
 ┣ 📂middleware
 ┃ ┗ 📜Authentication.go
 ┣ 📂models
 ┃ ┣ 📜content.go
 ┃ ┣ 📜task.go
 ┃ ┣ 📜user.go
 ┃ ┣ 📜userChat.go
 ┃ ┗ 📜userSub.go
 ┣ 📂repository
 ┃ ┣ 📜contens.go
 ┃ ┣ 📜task.go
 ┃ ┣ 📜userChat.go
 ┃ ┣ 📜users.go
 ┃ ┗ 📜userSub.go
 ┣ 📂routes
 ┃ ┣ 📜auth.go
 ┃ ┣ 📜content.go
 ┃ ┣ 📜task.go
 ┃ ┣ 📜userChat.go
 ┃ ┗ 📜userSub.go
 ┣ 📂security
 ┃ ┣ 📜password.go
 ┃ ┣ 📜token.go
 ┃ ┗ 📜token_test.go
 ┣ 📂util
 ┃ ┣ 📜errors.go
 ┃ ┗ 📜util.go
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜main.go

 ```


 ### C. INSTALLATIONS
 1. Clone the repository using
 `git clone https://github.com/nailykhry/TEFA-STUDYCASE-1.git`

 2. Run `go mod tidy` in terminal
 3. Create `database` and `users` collections in MongoDB
 4. Run program using `go run main.go`

 ### D. GROUPS
 - Naily Khairiya
 - Helsa Nesta 
 
 ### DEMO 
 
 ### DOCUMENTS