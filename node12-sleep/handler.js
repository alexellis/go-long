'use strict'

module.exports = async (event, context) => {
  await wait()
  return context
     .status(200)
     .succeed("OK")
}

async function wait() {
  return new Promise(r => {
   setTimeout(function() {
     r()
   }, 15000)
  })
}

