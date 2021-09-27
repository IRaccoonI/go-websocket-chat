<template>
  <div class="home">
    <Chat></Chat>
  </div>
</template>

<script>
// @ is an alias to /src
import Chat from "@/components/Chat.vue";

export default {
  name: "Home",
  components: {
    Chat,
  },
};
</script>

<style scoped>
.home {
  height: 100vh;
}
</style>
