'use strict'

module.exports = async (event, context) => {

  let duration = 15000
  if(event.headers["x-duration-ms"]) {
    duration = parseInt(event.headers["x-duration-ms"])
  }

  await wait(duration)
  return context
     .status(200)
     .succeed("OK")
}

async function wait(duration) {
  return new Promise(r => {
   setTimeout(function() {
     r()
   }, duration)
  })
}

