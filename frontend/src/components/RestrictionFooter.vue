<template>
  <div class="bar">
    <div class="item">Quelle Inzidenz: <a
        href="https://npgeo-corona-npgeo-de.hub.arcgis.com/datasets/917fc37a709542548cc3be077a786c17_0">RKI</a></div>
    <div class="item">Stand der Regeln: {{ stand }}</div>
    <div class="item">Quelle der Regeln: <a :href="source">{{ name }}</a></div>
  </div>
  <div>Alle Angaben ohne Gewähr!</div>
</template>

<script>
export default {
  name: "Status",
  props: {
    stand: String,
    source: String,
    name: String
  },
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
}
</style>