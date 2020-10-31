<template>
  <Overlay class="impress">
    <div>
      <h1>Impressum</h1>
      <div class="content">
        <p> Simon Petrik </p>
        <p> Willy-Andreas-Allee 7 </p>
        <p> 76131 Karlsruhe </p>
        <p> simon.petrik@web.de </p>
      </div>
    </div>
  </Overlay>
</template>

<script>
import Overlay from "@/components/Overlay";

export default {
  name: "Impress",
  components: {Overlay}
}
</script>

<style scoped>
.impress {
  width: 30%;
  height: 30%;
  top: 35%;
  left: 35%;
}
.content {
  font-size: 16pt;
}
</style>