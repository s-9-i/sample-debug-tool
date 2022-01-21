<template>
  <div class="panel-body">
    <div class="form">
      <h2>{{ group.name }} / {{ command.name }}</h2>
      {{ command.description }}
      <form>
        <vue-form-generator
          :schema="command.schema"
          :model="command.model"
          :options="formOptions"
        >
        </vue-form-generator>
        <v-btn elevation="2" class="ma-3" @click="send">実行</v-btn>
      </form>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import VueFormGenerator from "vue-form-generator";
import "vue-form-generator/dist/vfg.css";
import axios from "axios";

Vue.use(VueFormGenerator);

export default {
  data() {
    return {
      formOptions: {
        validateAfterLoad: false,
        validateAfterChanged: true,
        validateAsync: true,
      },
    };
  },

  props: {
    group: Object,
    command: Object,
  },

  methods: {
    send: async function () {
      try {
        const apiUrl =
          "http://localhost:1323" + this.group.url + this.command.url;
        const response = await axios.post(apiUrl, this.command.model);
        console.log(response);
        alert("success");
      } catch (error) {
        alert(error);
      }
    },
  },
};
</script>
