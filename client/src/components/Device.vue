<template>
  <v-card>
      <p> Name: {{name}}</p>
       <p> ID: {{id}}</p>
       <p> Admin State: {{adminState}}</p>
       <p> Operating State: {{operatingState}}</p>
       <p> Labels: {{labels}}</p>
       <v-btn @click="showCommands = !showCommands">Show Commands</v-btn>
       <div v-if="showCommands">
        <p v-for="command in commands" :key="command.id">
            Command ID: {{ command.id }}
            <br>
            Command name: {{ command.name}}
            <br>
            Command created: {{ command.created }}
            <br>
            Command modified: {{ command.modified }}
            <br>
            command Get URL: {{ command.get.url }}
            <br>
            Command get path: {{ command.get.path }}
            <br>
            Command get responses: {{ command.get.responses }}
        </p>
      </div>
     </v-card>
</template>

<script>
export default {
    props: ['name', 'id', 'adminState', 'operatingState', 'labels', 'commands'],
    data(){
        return {
            showCommands: false
        }
    }
  }
</script>

<style>

</style>