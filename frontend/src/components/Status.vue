<template>
  <div class="bar">
    <div class="item">Bundesland: <b>{{ state }}</b></div>
    <div class="item">7-Tage-Inzidenz: <b>{{ incidenceFormatted }}</b></div>
  </div>
</template>

<script>
export default {
  name: "Status",
  props: {
    state: String,
    incidence: Number,
  },
  computed: {
    incidenceFormatted() {
      if (isNaN(this.incidence)) {
        return this.incidence;
      }
      return this.incidence.toFixed(2);
    }
  }
}
</script>

<style scoped>
.bar {
  padding-top: 15px;

  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  align-content: center;
}

.item {
  padding: 15px;
  font-size: 16pt;
}
</style>