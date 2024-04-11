![Go Blog API Logo White](https://github.com/Anikate-De/blog-api/assets/40452578/e1c89389-2b6b-459b-9739-a044956f4c55)

# Go Blog API

A RESTful API comprising all endpoints for a Blogging platform. Built with Go!
Check out the [documentation](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL) for all endpoints!


## 🛢 Database Structure
![Go Blog API ER Diagram](https://github.com/Anikate-De/blog-api/assets/40452578/52a391b2-4bce-48de-a17a-ed1ea74029b1)

## 🔗 Endpoints
 - Home
   - [`GET /`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#0c30f6fb-84e4-4b9e-a000-55fb89a9fc5a) Home Ping
 - Users
   - [`POST /signup`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#133ef327-771b-4dc2-be57-47bcadbed7c6) Register New User
   - [`POST /login`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#cd1ac2b7-14de-4f89-926c-672f2e305231) Login User
   - [`GET /users/:uid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#c052d070-733a-455e-93b5-a67ed0a61fcb) Get User by UID
   - [`PUT /users/update`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#ea807659-6903-4418-b4ad-86746169e33b) Update User Profile
   - [`DELETE /unregister`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#bb14c3a3-0c31-4683-b896-9ed4dbcfe638) Delete User Profile
 - Blogs
   - [`POST /blogs`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#00db488f-2cbf-4bf9-ac0e-55d9aa9df709) Create New Blog
   - [`GET /blogs`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#e295c773-a04c-4c97-9650-d2189a5d01a5) Get all Blogs
   - [`GET /blogs/:bid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#e6d074dd-bc6f-43d2-9e4a-4ce125301d84) Get Blog by ID
   - [`PUT /blogs/:bid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#8e7f412c-aba5-4f87-b34c-a5408f266a91) Update Blog
   - [`DELETE /blogs/:bid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#ec88edbb-5f02-4061-94bf-4bf751fb7c32) Delete Blog
   - [`PUT /blogs/:bid/share`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#d6ab6a81-f431-4552-afca-0d63370ce4b3) Share Blog
 - Comments
   - [`POST /blogs/:bid/comments`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#9e0b9367-734d-42d9-a19e-4182638540ea) Add Comment on Blog
   - [`GET /blogs/:bid/comments`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#31a20c1e-5034-4d68-b603-9a2988ad9b46) Get all Comments on Blog
   - [`PUT /blogs/:bid/comment/:cid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#2ce0fc14-cf6a-4f12-bc86-53d62c888d3a) Update Comment
   - [`DELETE /blogs/:bid/comment/:cid`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#e62b9a89-8a5a-4706-a8c9-104e807aff24) Delete Comment
 - Likes
   - [`POST /blogs/:bid/likes`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#a362946f-10ed-4e5d-b969-f91c77065bb9) Like Blog
   - [`GET /blogs/:bid/likes`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#0ac30cd5-29e6-494f-b8de-d119b97f38d9) Get all Likes on Blog
   - [`DELETE /blogs/:bid/likes`](https://documenter.getpostman.com/view/4793515/2sA3BgBFzL#dc240007-cce8-49f2-a1b2-f1c9a75d3501) Unlike Blog

## ❓How to Use

### Pre-requisites

- **Go** is installed and added to `PATH`

### Steps to Follow

- In the terminal, execute:
  
  ```bash
  $ go run .
  ```

- That's it! The Go Blog API can be accessed using http://localhost:8080/

> If prompted, allow the Go Build Outputs access to the network. Ensure that the Firewall doesn't prevent API from listening for connections.

## 🤝 Contributing

Contributions are always welcome!

If you'd like to see a feature get added, create a new issue!
Send in your PRs and improvements/bug fixes

## 📖Lessons Learned

The Go Blog API was a project to strengthen my foundations of Go programming. I wanted to independently build a fully-fledged REST API for a Blogging app, tackling all challenges along the way.

I learnt many things along the way, the most notable ones are mentioned below -

- Gin framework
- Hashing Passwords
- JWTs
- SQLite and Go interconnectivity
- Structs and Bindings
- **and so much more...**

## 💡 Authors

- [@Anikate De](https://www.github.com/Anikate-De)

## 📝 License

Copyright © 2022-present, Anikate De

This project is licensed under [Apache License 2.0](LICENSE)
