<template>
  <div id='the-app'>
    <nav-bar id='nav' />
    <router-view id='body'/>
  </div>
</template>

<script>
import NavBar from '@/components/layout/nav-bar'
export default {
  name: 'App',
  components: {
    NavBar
  },
  // created() {
  //   this.$router.options.routes.forEach(route => {
  //     console.log({ name: route.name , path: route.path })
  //   })
  // }
}
</script>

<style lang="scss">
#the-app {
  min-height: 100vh;
  max-width: 960px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  position: relative;
  #body {
    flex: 1;
  }
}
</style>
