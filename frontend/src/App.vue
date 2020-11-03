<template>
  <h1 class="title fly-in">Wo hältst du dich auf?</h1>
  <div class="position-wrapper">
    <Position @on-location-requested="requestLocation" :nuts="nuts" @selected="onSelected"/>
  </div>
  <div>
    <Status :state="state" :incidence="incidence"/>
  </div>
  <div class="rules">
    <Restriction v-for="restriction in restrictions" :key="restriction.id" :title="restriction.title"
                 :description="restriction.description"/>
  </div>
  <RestrictionFooter v-if="areRestrictionsAvailable()" :name="name" :stand="stand" :source="source"/>

  <div class="center-horizontally">
    <DataProtection v-if="showDataProtection" @on-close="onDataProtectionClose"/>
    <Impress v-if="showImpress" @on-close="onImpressClose"/>
  </div>

  <div class="bar">
    <a class="item" href="#" v-on:click="onDataProtectionOpen">Datenschutzerklärung</a>
    <a class="item" href="#" v-on:click="onImpressOpen">Impressum</a>
  </div>
</template>

<script>
import Position from './components/Position.vue'
import Restriction from "@/components/Restriction.vue";
import Status from "@/components/Status.vue";
import RestrictionFooter from "@/components/RestrictionFooter.vue";
import json from './assets/restrictions.json';
import axios from 'axios';
import DataProtection from "@/components/DataProtection";
import Impress from "@/components/Impress";

const unknown = "Nicht bekannt";

export default {
  name: 'App',
  components: {
    Impress,
    DataProtection,
    Restriction,
    Position,
    Status,
    RestrictionFooter
  },
  props: {},
  methods: {
    onImpressOpen() {
      this.showImpress = true;
    },
    onImpressClose() {
      this.showImpress = false;
    },
    onDataProtectionOpen() {
      this.showDataProtection = true;
    },
    onDataProtectionClose() {
      this.showDataProtection = false;
    },
    requestLocation() {
      navigator.geolocation.getCurrentPosition((location) => {
        this.onLocationUpdated(location);
      }, function (errorPosition) {
        alert('Standort konnte nicht abgefragt werrden!\n\n"' + errorPosition.message + '"\n\nWenn du iOS und Safari nutzt, erlaube Safari auf deinen Standort zuzugreifen!\n(Einstellungen -> Datenschutz -> Ortungsdienste -> Safari -> "Beim Verwenden")\n(Einstellungen -> Safari -> Standort -> Fragen)');
      });
    },
    onSelected(nuts) {
      const instance = axios.create();
      instance
          .get('/nuts/' + nuts)
          .then(response => {
            this.onServerResult(response);
          })
          .catch(this.onServerError);
    },
    onLocationUpdated(location) {
      const instance = axios.create();
      instance
          .get('/locate/' + location.coords.latitude + '/' + location.coords.longitude)
          .then(response => {
            this.onServerResult(response);
          })
          .catch(this.onServerError);
    },
    areRestrictionsAvailable() {
      return this.restrictions & this.restrictions.length > 0;
    },
    onServerResult(response) {
      if (response.status === 200) {
        this.district = response.data.Name;
        this.state = response.data.Land;
        this.incidence = response.data.Incidence;
        this.nuts = response.data.Nuts;
        this.restrictions = [];
        for (let land of json) {
          if (land.land === this.state) {
            let marks = Object.keys(land.restrictions);
            marks = marks.map((s) => parseInt(s)).sort((a, b) => a - b);
            for (let i = 1; i < marks.length; i++) {
              const lowerBound = marks[i - 1];
              const upperBound = marks[i];
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
      } else {
        this.onServerError();
      }
    },
    onServerError(error) {
      this.district = unknown;
      this.state = unknown;
      this.incidence = unknown;
      this.nuts = undefined;
      this.restrictions = [];
      this.source = undefined;
      this.name = undefined;
      this.stand = undefined;
      if (error) {
        alert(error);
      }
    }
  },
  data() {
    return {
      restrictions: [],
      district: unknown,
      state: unknown,
      incidence: unknown,
      location: undefined,
      stand: undefined,
      source: undefined,
      name: undefined,
      nuts: undefined,
      showDataProtection: false,
      showImpress: false,
    }
  },
  mounted() {
    let counter = 0;
    for (let land of json) {
      let marks = Object.keys(land.restrictions);
      for (let mark of marks) {
        const restrictions = land.restrictions[mark];
        for (let restriction of restrictions) {
          restriction.id = counter;
          counter++;
        }
      }
    }
    this.requestLocation();
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

.center-horizontally {
  display: flex;
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
