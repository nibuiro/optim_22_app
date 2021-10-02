'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    API: '"https://3b454541-17d7-43b7-9e25-982b18a5cc77.mock.pstmn.io/api"'
})