<template>
  <h1 class="title fly-in">Wo hältst du dich auf?</h1>
  <div class="position-wrapper">
    <Position @position-update="onLocationUpdated" :district="district"/>
  </div>
  <div>
    <Status :state="state" :incidence="incidence"/>
  </div>
  <div class="rules">
    <Restriction v-for="restriction in restrictions" :key="restriction.id" :title="restriction.title"
                 :description="restriction.description"/>
  </div>
  <Footer v-if="restrictions" :name="name" :stand="stand" :source="source"/>
</template>

<script>
import Position from './components/Position.vue'
import Restriction from "@/components/Restriction.vue";
import Status from "@/components/Status.vue";
import Footer from "@/components/Footer.vue";
import json from './assets/restrictions.json';
import axios from 'axios';

export default {
  name: 'App',
  components: {
    Restriction,
    Position,
    Status,
    Footer
  },
  props: {},
  methods: {
    onLocationUpdated(location) {
      const instance = axios.create();
      instance
          .get('/locate/' + location.coords.latitude + '/' + location.coords.longitude)
          .then(response => {
            this.district = response.data.Name;
            this.state = response.data.Land;
            this.incidence = response.data.Incidence;
            for (let land of json) {
              if (land.land === this.state) {
                let marks = Object.keys(land.restrictions);
                marks = marks.sort();
                for (let i = 1; i < marks.length; i++) {
                  const lowerBound = parseInt(marks[i - 1]);
                  const upperBound = parseInt(marks[i]);
                  const incidence = parseFloat(this.incidence);
                  if (lowerBound <= incidence && incidence < upperBound) {
                    this.restrictions = land.restrictions[marks[i - 1]];
                    return;
                  }
                }
                this.restrictions = land.restrictions[marks[marks.length - 1]];
                this.source = land.source;
                this.name = land.name;
                this.stand = land.stand;
              }
            }
          });
    }
  },
  data() {
    return {
      restrictions: undefined,
      district: "Nicht bekannt",
      state: "Nicht bekannt",
      incidence: "Nicht bekannt",
      location: undefined,
      stand: undefined,
      source: undefined,
      name: undefined
    }
  }
}
</script>

<style>
.position-wrapper {
  margin: 0 auto;
  max-width: 100%;
  display: flex;

  justify-content: center;
  flex-wrap: wrap;
}

.title {
  font-size: 32pt;
  color: #484848;
}

@media only screen and (min-width: 768px) {
  .title {
    font-size: 64pt;
  }
}

.rules {
  display: flex;
  margin-top: 25px;
  flex-direction: column;
  align-content: center;
  align-items: center;
}

@keyframes fly-in {

  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(0%);
  }
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
