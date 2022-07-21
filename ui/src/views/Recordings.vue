<template>
  <div class="seperated-flex">
    <div v-for='recording in recordings' :key='recording.name'>
      <span>
        {{ recording.name }}
      </span>
      <span class="u-pull-right">
        <a :href='recording.downloadPath' :download='recording.name' class="button button-primary">
          Download
        </a>
        <button @click='deleteRecording(recording)' class="button">
          Delete
        </button>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Recordings',
  data () {
    return {
      recordings: []
    }
  },
  methods: {
    getRecordings () {
      fetch('/api/v1/recordings').then(res => res.json())
        .then(rec => this.recordings = rec)
        .catch(err => console.log(err))
    },
    deleteRecording (recording) {
      if (confirm(`Are you sure you want to delete ${recording.name}?`)) {
        fetch(recording.downloadPath, { method: 'DELETE' }).then(() => {
          this.recordings = this.recordings.filter(rec => rec !== recording)
        }).catch(err => console.log(err))
      }
    },
  },
  created () {
    this.getRecordings()
  }
}
</script>

<style scoped lang="scss">
.seperated-flex {
  line-height: 6.5rem;
  width: 100%;
  a, button {
    margin: 5px;
  }
}
</style>
