<template>
  <div id="position">
    <Select id="location-text" :elements="suggestions" @selected="onSelected" :nuts="nuts"></Select>
    <div id="locate-button" v-on:click="locate">
      <img src="../assets/location.svg">
    </div>
  </div>
</template>

<script>
import json from '../assets/kreise.json';
import Select from "@/components/Select";

export default {
  name: "Position",
  components: {Select},
  props: {
    nuts: String,
  },
  methods: {
    locate() {
      this.$emit('on-location-requested');
    },
    onSelected(nuts) {
      this.$emit('selected', nuts);
    }
  },
  data() {
    return {
      suggestions: json
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