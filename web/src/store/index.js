import * as Vue from 'vue'
import * as Vuex from 'vuex'
import getters from './getters'
import app from './modules/app'
import permission from './modules/permission'
import settings from './modules/settings'
import user from './modules/user'

const store = Vuex.createStore({
  modules: {
    app,
    permission,
    settings,
    user,
  },
  getters,
})

export default store
