import Hammer from 'hammerjs'

class TouchEvents {
  constructor(el) {
    const tap = new Hammer.Tap({ event: 'tap', taps: 1 })
    const dblTap = new Hammer.Tap({ event: 'dbltap', taps: 2 })

    dblTap.recognizeWith(tap)
    tap.requireFailure(dblTap)

    const mc = new Hammer.Manager(el)

    mc.add([dblTap, tap])

    mc.on('tap', this.trigger)
    mc.on('dbltap', this.trigger)

    this.events = {}
  }

  on = (eventName, fn) => {
    this.events[eventName] = fn
  }

  off = (eventName) => {
    delete this.events[eventName]
  }

  trigger = (event) => {
    if (event.target.matches('button')) { // , button *
      return
    }

    if (this.events[event.type]) {
      this.events[event.type]()
    }
  }
}

export default TouchEvents
