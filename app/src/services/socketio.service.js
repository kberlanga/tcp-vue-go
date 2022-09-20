import io from "socket.io-client";

class SocketioService {
  socket;
  constructor() {}

  setupSocketConnection() {
    this.socket = io("http://localhost:3000", { transport: ["websocket"] });
  }
}

export default new SocketioService();
