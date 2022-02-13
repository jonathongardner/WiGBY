<template>
  <div>
    <button class='play' :class="{ paused: riding }" @click='toggleRide'></button>
  </div>
</template>

<script>
export default {
  name: 'controls',
  props: {
    riding: {
      type: Boolean,
      required: true
    }
  },
  methods: {
    toggleRide () {
      this.$emit('update:riding', !this.riding)
    }
  }
}
</script>

<style scoped lang="scss">
.play {
  border: 0;
  background: transparent;
  box-sizing: border-box;
  width: 0;
  height: 3.7rem;

  border-color: transparent transparent transparent #202020;
  transition: 100ms all ease;
  cursor: pointer;

  // play state
  border-style: solid;
  border-width: 1.85rem 0 1.85rem 3rem;

  &.paused {
    border-style: double;
    border-width: 0px 0 0px 3rem;
  }

  &:hover {
    border-color: transparent transparent transparent #404040;
  }
}
</style>
