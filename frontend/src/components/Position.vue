<template>
  <div id="position">
    <div id="location-text">{{ district }}</div>
    <div id="locate-button" v-on:click="locate">
      <img src="../assets/location.svg">
    </div>
  </div>
</template>

<script>
export default {
  name: "Position",
  props: {
    district: String,
  },
  methods: {
    locate() {
      navigator.geolocation.getCurrentPosition((location) => {
        this.$emit("position-update", location);
      }, function (errorPosition) {
        alert("Position could not be obtained! " + errorPosition.message);
      })
    }
  }
}
</script>

<style scoped>
#position {
  border-style: solid;
  border-color: gainsboro;
  border-radius: 5px;
  padding-top: 5px;
  width: 400px;

  display: flex;
  flex-wrap: wrap;
  align-content: space-between;
}

#location-text {
  text-align: left;
  padding-left: 10px;
  padding-top: 5px;
  flex: 1;
  color: gray;
  cursor: default;
}

#locate-button {
  flex: 0 0 auto;
  padding-right: 10px;
  cursor: pointer;
}
</style>