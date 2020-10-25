<template>
  <div class="restriction-wrapper">
    <h1 class="restriction-title">{{ title }}</h1>
    <div>
      <ul>
        <li class="restrictions" v-for="rule in description" :key="rule">
          {{ rule }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "Restriction",
  props: {
    title: String,
    description: Array
  }
}
</script>

<style scoped>
.restriction-title {
  font-size: 24pt;
}

.restrictions {
  font-size: 20pt;
}

.restriction-wrapper {
  width: 70%;

  padding-left: 45px;
  padding-right: 25px;
  padding-bottom: 10px;
  margin: 25px;
  text-align: left;
  border-radius: 10px;
  box-shadow: 0 6px 10px rgba(0,0,0,.08), 0 0 6px rgba(0,0,0,.05);
  font-family: 'Nunito', sans-serif;
}
</style>