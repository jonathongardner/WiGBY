<template>
  <div class="seperated-flex">
    <my-back />
    <router-link class="close" :to="{ name: 'Home' }">&times;</router-link>
  </div>
</template>

<script>
import MyBack from '@/components/common/my-back.vue'

export default {
  name: 'Camera',
  components: {
    MyBack
  },
  methods: {
    back() {
      console.log('Back')
    }
  }
}
</script>

<style scoped lang="scss">
.seperated-flex {
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  top:0;
  bottom:0;
  left:0;
  right:0;
  background-color: rgba(0,0,0,0.4); /* Black w/ opacity */

  .close:link, .close:visited, .close:hover, .close:active {
    text-decoration: none;
  }
  .close {
    color: white;
    position: absolute;
    right: 15px;
    top: 0px;
    font-size: 1.5em;
    font-weight: bold;
  }

  .close:hover,
  .close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
  }
  img {
    width: 100vw;
    height: 100vh;
    object-fit: contain;
  }
}
</style>
