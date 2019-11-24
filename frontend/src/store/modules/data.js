const state = {
  data: []
}

const getters = {}

const actions = {
  getData( {commit} ) {
    fetch(`/backend/data.json`)
      .then( (resp) => resp.json() )
      .then( (json) => commit('setData', json) )
  }
}

const mutations = {
  setData( state, data ) {
    state.data = data;
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
