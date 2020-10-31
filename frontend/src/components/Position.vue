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
    // onPermissionUpdate(permissionStatus) {
    //   if (permissionStatus.state === 'granted' || permissionStatus.state === 'prompt') {
    //     this.requestLocation();
    //   } else if (permissionStatus.state === 'denied') {
    //     alert("Geolocation permission was denied!")
    //   }
    // },
    locate() {
      if (navigator.geolocation) {
        this.requestLocation();
      } else {
        alert("Browser or device does not support geolocation services!")
      }
    },
    // requestPermissionAndLocate() {
    //   navigator.permissions.query({name: 'geolocation'}).then((result) => {
    //     this.onPermissionUpdate(result);
    //     result.onchange = () => {
    //       this.onPermissionUpdate(result);
    //     }
    //   });
    // },
    requestLocation() {
      navigator.geolocation.getCurrentPosition((location) => {
        this.$emit("position-update", location);
      }, function (errorPosition) {
        alert("Position could not be obtained! " + errorPosition.message);
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