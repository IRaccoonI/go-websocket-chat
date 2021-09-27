<template>
  <div class="chat">Hello</div>
  <div class="footer">
    <textarea
      type="text"
      multiple
      v-model="newMessage"
      @keydown.enter.prevent="newMessageInputEnter"
    />
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
  data: () => ({
    newMessage: "",
    ws: null,
  }),
  methods: {
    newMessageInputEnter: function (ev) {
      if (ev.shiftKey) {
        this.newMessage += "\n";
      } else {
        this.$socket.send(JSON.stringify({ qwe: "qwe" }));
        this.newMessage = "";
      }
    },
  },
  mounted: function () {
    this.$options.sockets.onmessage += (msg) =>
      console.log(JSON.parse(msg.data));
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.chat {
  height: 100%;
  width: 100%;
  position: relative;
}
.footer {
  height: 80px;
  width: 100%;
  position: absolute;
  bottom: 0px;
}
.footer textarea {
  height: 100%;
  width: 100%;
}
</style>
