<template>
  <img alt="Vue logo" src="./assets/logo.png" />
</template>

<script>
import SocketioService from "./services/socketio.service";

export default {
  data() {
    return {
      isConnected: false,
      socketMessage: "",
      changedMessage: "",
    };
  },
  created() {
    SocketioService.setupSocketConnection();
    SocketioService.socket.emit("app", "app connected");
    SocketioService.socket.on("connections", (msg) => {
      console.log("connections now", msg);
    });
  },
};
</script>
