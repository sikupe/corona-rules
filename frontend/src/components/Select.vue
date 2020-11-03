<template>
  <div class="select" v-click-outside="onFocusOut">
    <label for="lk-input"></label><input id="lk-input" type="text" v-model="input" v-on:focusin="showDropdown">
    <div class="select-wrapper" v-bind:class="{hidden: isDropdownHidden}">
      <div v-for="suggestion in filtered" :key="suggestion.code" class="select-item"
           v-on:click="onSelected(suggestion.code)">
        {{ suggestion.label }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Select",
  props: {
    elements: Array,
    nuts: String,
  },
  methods: {
    onFocusOut() {
      this.hideDropdown();
    },
    showDropdown() {
      this.isDropdownHidden = false;
    },
    hideDropdown() {
      this.isDropdownHidden = true;
    },
    onSelected(code) {
      this.selected = code;
      this.input = this.getLabelBy(code);
      this.hideDropdown();
      this.$emit('selected', code);
    },
    getLabelBy(code) {
      for (let element of this.elements) {
        if (element.code === code) {
          return element.label;
        }
      }
      return undefined;
    }
  },
  data() {
    return {
      selected: undefined,
      input: "",
      isDropdownHidden: true
    };
  },
  computed: {
    filtered() {
      const filtered = []
      for (let element of this.elements) {
        const elementLabel = element.label.toLowerCase();
        const input = this.input ? this.input.toLowerCase() : "";
        const index = elementLabel.indexOf(input);
        if (index !== -1) {
          filtered.push(element);
        }
      }
      return filtered;
    }
  },
  watch: {
    nuts(newVal) {
      this.onSelected(newVal);
    }
  }
}
</script>

<style scoped>
#lk-input {
  outline: none;
  border-style: none;

  width: 100%;

  text-align: left;
  font-size: 14pt;
}

#lk-input:focus {
  border-style: none;
}

.select {
  position: relative;
}

.hidden {
  display: none;
}

.select-wrapper {
  position: absolute;
  overflow: scroll;
  width: 100%;
  height: 150px;

  border-style: solid;
  border-color: gainsboro;
  border-radius: 5px;
  z-index: 1;

  background-color: white;
}

.select-item {
  text-decoration: none;
  display: block;
  padding: 5px;
  background-color: #f6f6f6;
  color: black;
}

.select-item:focus {
  background-color: #f1f1f1;
}

.select-item:hover {
  background-color: #f1f1f1;
}
</style>