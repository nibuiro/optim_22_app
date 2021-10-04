'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    API: '"https://d90a64fd-0de1-4060-8ad7-51e22f0d78d8.mock.pstmn.io/api"'
})