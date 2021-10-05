'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    API: '"https://596f143a-7437-4f17-8068-22937e868886.mock.pstmn.io/api"'
})