<template>
<div>
  <Device
  v-for="device in devices" :key="device.id"
  :name="device.name"
  :id="device.id"
  :adminState="device.adminState"
  :operatingState="device.operatingState"
  :labels="device.labels"
  :commands="device.commands"
  />
</div>
</template>

<script>
import axios from 'axios'
import Device from './Device.vue'
export default {
  components: {
    Device
  },
    mounted() {
        axios.get('/api/v1/device').then(res => this.devices = res.data)
    },
    data() {
        return {
            devices: []
        }
    },
    methods: {
    },
}
</script>

<style>

</style>