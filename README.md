# tcp-vue-go
Server to transfer files between two or more clients using a custom protoco based on TCP protocol.

## 📦 Packages:

- ⚡️ [Vuejs](https://v2.vuejs.org/v2/guide/) - The Progressive JavaScript Framework.
- 💥 [Socket.io](https://socket.io/) - Bidirectional and low-latency communication for every platform.
- 💅 [Vuetify](https://vuetifyjs.com/en/) - Vuetify is a Vue UI Library with beautifully handcrafted Material Components.


## 🚀 Getting Started
```bash
git clone https://github.com/kberlanga/tcp-vue-go.git
```
### 🗄️ Server

Up server:

```bash
cd server
go build && ./server start
```
### 💻 Client

To set clients in mode to receive files:

```bash
cd client
go build && ./client receive -channel 1
```

To send files:

```bash
cd client
go build && ./client send -filepath test.txt -channel 2
```

### 👁️ App

```bash
cd app
npm install
npm run build
```
Go to [http://localhost:3000](http://localhost:3000)
