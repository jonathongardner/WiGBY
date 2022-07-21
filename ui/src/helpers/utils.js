
export const promiseTimeout = (timeout) => {
  return new Promise((resolve) => {
    setTimeout(() => {
        resolve()
    }, timeout)
  })
}

export const debounce = function(func, timeout = 500) {
  let timer
  return function() {
    clearTimeout(timer)
    timer = setTimeout(() => { func.apply(this) }, timeout)
  }
}
