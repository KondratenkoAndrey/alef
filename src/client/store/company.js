import axios from '~/plugins/axios'

export const state = () => ({
  info: [
    { shortName: 'Алеф' },
    { fullName: 'Компания Алеф' },
    { phone: '+7(789)123-45-67' },
    { email: 'info@example.ru' },
  ],
})

export const actions = {
  async loadInfo({ commit }) {
    const { data } = await axios.get('/company-info')
    commit('setInfo', data)
  },
}

export const mutations = {
  setInfo(state, data) {
    state.info = data
  },
}
