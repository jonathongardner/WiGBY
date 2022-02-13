<template>
  <nav class="navbar" :class="{ big: !back }">
    <router-link :to="{ name: 'Home' }">&lsaquo; Back</router-link>
    <span>
      WeGYB
    </span>
  </nav>
</template>

<script>
export default {
  name: 'NavBar',
  computed: {
    back () {
      return this.$route.meta.back
    }
  },
}
</script>

<style scoped lang="scss">
.navbar {
  font-family: "Special Elite", "normal";
  display: block;
  width: 100%;
  text-align: center;
  background: #fff;
  border-bottom: 1.5px solid #eee;
  font-size: 2rem;
  transition: visibility 10s, font-size 0.1s;
  position: relative;
  a {
    position: absolute;
    left: 5px;
  }
}
.big {
  font-size: 3.5rem;
  a {
    visibility: hidden;
  }
}
</style>
