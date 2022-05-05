<template>
  <div class="seperated-flex">
    <div>
      <router-link :to="{ name: 'Ride' }">Ride</router-link>
    </div>
    <div>
      <router-link :to="{ name: 'Camera' }">Camera</router-link>
    </div>
    <div>
      <router-link :to="{ name: 'Settings' }">Settings</router-link>
    </div>
    <div>
      <router-link :to="{ name: 'Help' }">Help</router-link>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'Home',
  components: {
    // HelloWorld
  }
}
</script>

<style scoped lang="scss">
.seperated-flex {
  a {
    padding: 0px 15px;
    text-transform: uppercase;
    font-weight: 600;
    letter-spacing: .2rem;
    text-decoration: none;
    line-height: 6.5rem;
    color: #222;
    display:inline-block;
    width: 100%;
    box-sizing: border-box;
  }
}
</style>
