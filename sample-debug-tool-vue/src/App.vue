<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <v-list>
        <v-list-group
          v-for="(group, groupIndex) in commandGroups"
          :key="group.name"
          no-action
        >
          <template #activator>
            <v-list-item-title>{{ group.name }}</v-list-item-title>
          </template>
          <v-list-item
            v-for="(command, commandIndex) in group.commands"
            :key="command.name"
            @click="generateForm(groupIndex, commandIndex)"
          >
            <v-list-item-title>{{ command.name }}</v-list-item-title>
          </v-list-item>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <Form :group="currentCommandGroup" :command="currentCommand" />
    </v-main>
  </v-app>
</template>

<script>
import Form from "./components/Form";
import axios from "axios";

export default {
  name: "App",

  components: {
    Form,
  },

  data() {
    return {
      drawer: null,
      commandGroups: [],
      currentCommandGroup: {},
      currentCommand: {},
    };
  },

  created: function () {
    this.initCommands();
  },

  methods: {
    fetchCommands: async function () {
      const response = await axios.get("http://localhost:1323/api/list");
      return response.data;
    },
    initCommands: async function () {
      const ret = await this.fetchCommands();
      this.commandGroups = ret;
      this.generateForm(0, 0);
    },
    generateForm: function (groupIndex, commandIndex) {
      this.currentCommandGroup = this.commandGroups[groupIndex];
      this.currentCommand = this.currentCommandGroup.commands[commandIndex];
    },
  },
};
</script>
