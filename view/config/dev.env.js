'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    API: '"https://05efb1e3-21f2-4031-8cd8-c5f6cc9dcc3d.mock.pstmn.io/api"'
})