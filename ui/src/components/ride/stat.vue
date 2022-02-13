<template>
  <div class='ride-component'>
    <div class='value'>
      {{ value }}  <span v-if='units'>{{ units }}</span>
    </div>
    <div class='label'>
      {{ label }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'stat',
  props: {
    value: {
      type: String,
      required: true
    },
    label: {
      type: String,
      required: true
    },
    units: {
      type: String,
      required: false,
      default: ''
    }
  }
}
</script>

<style scoped lang="scss">
.ride-component {
  position: relative;
  .value {
    font-size: 6rem;
    span {
      font-size: 2rem;
    }
  }
  .label {
    font-size: 2rem;
    position: absolute;
    bottom: 5px;
    right: 5px;
  }
}
</style>
