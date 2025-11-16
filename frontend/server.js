import { createServer } from "http";
import { Server } from "socket.io";
import next from "next";

const dev = process.env.NODE_ENV !== "production";
const app = next({ dev });
const handler = app.getRequestHandler();

app.prepare().then(() => {
  const httpServer = createServer(handler);

  const io = new Server(httpServer, {
    path: "/socket",
    cors: {
      origin: "*",
    },
  });

  const onlineUsers = new Map();

  io.on("connection", (socket) => {
  
    socket.on("online", (userId) => {
      onlineUsers.set(userId, socket.id);
    });
  
    socket.on("send-message", (msg) => {
      const receiverSocket = onlineUsers.get(msg.receiver);
      if (receiverSocket) {
        io.to(receiverSocket).emit("receive-message", msg);
      }
    });
  
    socket.on("disconnect", () => {
      for (const [id, socketId] of onlineUsers.entries()) {
        if (socketId === socket.id) {
          onlineUsers.delete(id);
        }
      }
    });
  });

  httpServer.listen(3000, () => console.log("SERVER READY"));
});
