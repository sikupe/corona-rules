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
      if (navigator.geolocation) {
        this.requestLocation();
      } else {
        alert("Standortfreigabe von Browser oder Gerät nicht unterstützt!")
      }
    },
    requestLocation() {
      navigator.geolocation.getCurrentPosition((location) => {
        this.$emit("position-update", location);
      }, function (errorPosition) {
        alert('Standort konnte nicht abgefragt werrden!\n\n"' + errorPosition.message + '"\n\nWenn du iOS und Safari nutzt, erlaube Safari auf deinen Standort zuzugreifen! (Einstellungen -> Datenschutz -> Ortungsdienste -> Safari -> "Beim Verwenden")');
      });
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