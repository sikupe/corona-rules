<template>
  <div class="overlay">
    <div class="button close" v-on:click="onClose">
      <img src="../assets/close.svg"/>
    </div>
    <div class="overlay-content">
      <slot>Test</slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "Overlay",
  methods: {
    onClose() {
      this.$emit('on-close');
    }
  }
}
</script>

<style scoped>
.button {
  cursor: pointer;
}

.close {
  position: absolute;
  top: 25px;
  right: 45px;
}

.overlay {
  position: fixed; /* Sit on top of the page content */
  width: 70%; /* Full width (cover the whole page) */
  height: 80%; /* Full height (cover the whole page) */
  top: 5%;
  left: 15%;
  background-color: rgb(255, 255, 255); /* Black background with opacity */
  z-index: 2;
  padding: 25px;
  border-radius: 10px;
  box-shadow: 0 6px 10px rgba(0, 0, 0, .08), 0 0 6px rgba(0, 0, 0, .05);
}

.overlay-content {
  overflow: auto;
  height: 100%;
}
</style>