export const state = () => ({
  info: [
    { shortName: 'Алеф' },
    { fullName: 'Компания Алеф' },
    { phone: '+7(789)123-45-67' },
    { email: 'info@example.ru' },
  ],
})

export const actions = {
  async storeDispatchFunc({ commit }) {
    const { data } = await this.$axios.get('http://localhost:8080/company-info')
    commit('SET_DATA', data)
  },
}

export const mutations = {
  SET_DATA(state, data) {
    state.info = data
  },
}
