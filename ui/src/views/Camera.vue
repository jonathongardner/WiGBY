<template>
  <div class="seperated-flex">
    <my-back />
  </div>
</template>

<script>
import MyBack from '@/components/common/my-back.vue'

export default {
  name: 'Camera',
  components: {
    MyBack
  },
}
</script>

<style scoped lang="scss">
.seperated-flex {
  min-height: 100%;
}
#cover {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  box-sizing: border-box;
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}
</style>
