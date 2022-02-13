
export const promiseTimeout = (timeout) => {
  return new Promise((resolve) => {
    setTimeout(() => {
        resolve()
    }, timeout)
  })
}
