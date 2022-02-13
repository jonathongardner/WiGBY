const getSeconds = () => {
  return (new Date()).getTime() / 1000
}

const toHHMMSS = (secs) => {
  const hours   = Math.floor(secs / 3600)
  const minutes = Math.floor(secs / 60) % 60
  const seconds = secs % 60

  return [hours, minutes, seconds].map(zeroPad).join(":")
}
const zeroPad = (v) => {
  return v < 10 ? "0" + v : "" + v
}

const format = (v) => {
  const value = Math.floor(v)
  const decimal = Math.round(v * 100) % 100
  return zeroPad(value) + '.' + zeroPad(decimal)
}
const distance = (prevCoords, coords) => {
  const { longitude: lon1, latitude: lat1 } = prevCoords
  const { longitude: lon2, latitude: lat2 } = coords
  const R = 6371 // Radius of the earth in km
  const dLat = toRad(lat2 - lat1)
  const dLon = toRad(lon2 - lon1)
  const a = Math.sin(dLat/2) * Math.sin(dLat/2) +
          Math.cos(toRad(lat1)) * Math.cos(toRad(lat2)) *
          Math.sin(dLon/2) * Math.sin(dLon/2)
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a))
  const d = R * c // Distance in km
  return d
}
const toRad = (v) => v * Math.PI / 180

export default class Tracker {
  constructor(callbackMethod) {
    this.callbackMethod = callbackMethod
  }
  start() {
    if (this.callback) {
      return
    }
    this.startTime = getSeconds()
    this.distance = 0
    this.speed = 0
    this.avgSpeed = 0
    this.callback = setInterval(() => {
      this.callbackWithStats()
    }, 200)
    navigator.geolocation.watchPosition(this.updateDistance)
  }
  stop() {
    clearInterval(this.callback)
    navigator.geolocation.clearWatch(this.callback)
    this.callback = null
  }
  updateDistance = (loc) => {
    if (loc.coords.accuracy > 10000) {
      return
    }

    const prevCoords = this.prevCoords
    this.prevCoords = loc.coords
    if (!prevCoords) {
      return // if previous location wasnt set yet
    }
    const seconds = (loc.timestamp / 1000) - this.startTime

    this.distance += distance(prevCoords, loc.coords)
    this.speed = loc.speed || 0
    this.avgSpeed = seconds > 0  ? (60 * 60 * this.distance / seconds) : 0
  }
  callbackWithStats() {
    const seconds = Math.round(getSeconds() - this.startTime)

    return this.callbackMethod(
      {
        seconds: seconds,
        time: toHHMMSS(seconds),
        distance: format(this.distance),
        speed: format(this.speed),
        avgSpeed: format(this.avgSpeed)
      }
    )
  }
}
